package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	TEST_DB_PATH = "file::memory:?cache=shared"
)

// Just a wrapper for InitDB() and error handling to save
// space in tests. Also declared in model because golang.
func initTestDB() {
	err := model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}
}

// Sets up a router for testing purposes. You can pass any number of custom
// middlewares depending on how you want to set up your test.
func setupTestRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	testStore := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", testStore))

	for _, middleware := range middlewares {
		router.Use(middleware)
	}

	SetupHandlers(router)

	return router
}

// Sends a mock HTTP request and returns *httptest.ResponseRecorder. Panics if error encountered.
func sendMockHTTPRequest(method string, endpoint string, data *bytes.Buffer, router *gin.Engine) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()

	if data == nil {
		data = &bytes.Buffer{}
	}

	req, err := http.NewRequest(method, endpoint, data)
	if err != nil {
		panic(err)
	}

	router.ServeHTTP(w, req)

	return w
}

// Cleans up all DB tables.
func cleanUpDB() {
	model.DBConn.Exec("DELETE FROM users")
	model.DBConn.Exec("DELETE FROM user_personal_infos")
	model.DBConn.Exec("DELETE FROM teams")
	model.DBConn.Exec("DELETE FROM team_notifications")
	model.DBConn.Exec("DELETE FROM matches")
	model.DBConn.Exec("DELETE FROM promotion_to_manager_requests")
}

// Initializes db then runs all tests in controller
func TestMain(m *testing.M) {
	initTestDB()
	rc := m.Run()
	os.Exit(rc)
}

// Tests that handleLogout() responds HTTP Status Reset Content.
func Test_handleLogout(t *testing.T) {
	router := setupTestRouter()

	// mock request
	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/logout", nil, router)

	if w.Code != http.StatusResetContent {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusResetContent)
	}
}

// Tests handleLogin() for user that actually exists. Should respond
// HTTP Status Accepted.
func Test_handleLogin_GoodCredentials(t *testing.T) {
	router := setupTestRouter()

	// Add user to DB
	user := &model.User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}
	model.DBConn.Create(user)

	// create mock login credentials
	userCpy := &model.User{
		Username: "jdo",
		Password: "123",
	}

	// format as JSON
	reader, err := json.Marshal(userCpy)
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(reader), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("Response code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

// Test handleLogin() when credentials are invalid. Should
// respond HTTP Status Unauthorized.
func Test_handleLogin_BadCredentials(t *testing.T) {
	router := setupTestRouter()

	fakeUser := &model.User{
		Username: "doesnt_matter",
		Password: "yup",
	}

	//format as JSON
	reader, err := json.Marshal(fakeUser)
	if err != nil {
		t.Errorf("Error: %s\n", err)
	}

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(reader), router)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Response code was %d, should have been %d", w.Code, http.StatusUnauthorized)
	}
}

// Tests that handleLogin() responds with HTTP Status Bad Request
// when JSON is not correct.
func Test_handleLogin_BadJSON(t *testing.T) {
	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer([]byte("bad data")), router)

	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}
}

// Tests handleCreateTeamNotification() when a valid teamNotification
// JSON is sent over and all success conditions are met. Should return
// HTTP Status Accepted.
func Test_handleCreateTeamNotification_ValidInvite(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
		Role:     model.MANAGER,
	})
	model.DBConn.Create(&model.User{
		Username: "colbert",
		Role:     model.PLAYER,
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	notification := &model.TeamNotification{
		SenderUsername:   "jaluhrman",
		ReceiverUsername: "colbert",
	}

	jsonData, _ := json.Marshal(notification)

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createTeamRequest", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

// Tests case that handleCreateTeamNotification() endpoint is
// requested when a user is not logged in/there is no valid
// session. Should respond HTTP Status Unauthorized.
func Test_handleCreateTeamNotification_NoSession(t *testing.T) {
	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createTeamRequest", nil, router)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusUnauthorized)
	}
}

// Tests case that handleCreateTeamNotification() is called
// when a session exists but the user is not valid. Should respond
// HTTP Status Unauthorized.
func Test_handleCreateTeamNotification_InvalidUser(t *testing.T) {
	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createTeamRequest", nil, router)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusUnauthorized)
	}
}

// Tests case that handleCreateTeamNotification() is called without a valid
// JSON request body in context. Should respond HTTP Status Bad Request.
func Test_handleCreateTeamNotification_BadJSON(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createTeamRequest", nil, router)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusBadRequest)
	}

	cleanUpDB()
}

// Tests that handleUpdataeUserPersonalInfo() responds HTTP Status Accepted
// when everything is valid.
func Test_handleUpdateUserPersonalInfo_Valid(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	info := &model.UserPersonalInfo{
		Firstname: "Joe",
		Lastname:  "Luhrman",
		Height:    50,
		Weight:    240,
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		t.Error(err)
	}

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/updatePersonalInfo", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

// Tests that handleUpdataeUserPersonalInfo() responds HTTP Status Unauthorized
// when session doesn't exist.
func Test_handleUpdateUserPersonalInfo_NilSession(t *testing.T) {
	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/updatePersonalInfo", nil, router)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusUnauthorized)
	}
}

// Tests handleUpdataeUserPersonalInfo() responds HTTP Status Bad Request
// when the JSON is not correct.
func Test_handleUpdateUserPersonalInfo_BadJSON(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	// test endpoint just to set session for this test
	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/updatePersonalInfo", nil, router)

	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}

	cleanUpDB()
}

// Tests handleCreateAccount() responds HTTP Status Accepted when
// all conditions are met.
func Test_handleCreateAccount_Valid(t *testing.T) {
	router := setupTestRouter()

	testUser := &model.User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	jsonData, err := json.Marshal(testUser)
	if err != nil {
		t.Errorf("could not marshal json: %s\n", err)
	}

	// mock request
	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createAccount", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

// Tests handleCreatePhoto() returns HTTP Status Accepted
// when the request is correct.
func Test_handleCreatePhoto_Valid(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	info := gin.H{
		"photo": "dfjekjfcks",
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createPhoto", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

// Tests that handleCreatePhoto() responds with HTTP Status Bad
// Request when the request doesn't have the correct key value pair.
func Test_handleCreatePhoto_Invalid(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	// test endpoint just to set session for this test
	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	info := gin.H{
		"thisfieldisntrelevant": "dfjekjfcks",
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		t.Error(err)
	}

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createPhoto", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}

	cleanUpDB()
}

// Tests that create team endpoint responds with HTTP Status Accepted
// when all conditions are correct.
func Test_handleCreateTeam_Valid(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "kevin",
		Password: "wasspord",
		Role:     model.MANAGER,
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")
		ctx.Next()
	})

	info := &model.Team{
		Name:         "Greyhounds",
		TeamLocation: "Baltimore",
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		t.Error(err)
	}

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createTeam", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

// Tests that the create team endpoint correctly rejects a non-manager.
func Test_handleCreateTeam_NotMan(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "kevin",
		Password: "wasspord",
		Role:     model.PLAYER,
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")
		ctx.Next()
	})

	info := &model.Team{
		Name:         "Greyhounds",
		TeamLocation: "Baltimore",
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		t.Error(err)
	}

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createTeam", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusUnauthorized)
	}

	cleanUpDB()
}

// Tests that create team endpoint responds Status Conflict
// if manager already has a team.
func Test_handleCreateTeam_AlreadyHasTeam(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "kevin",
		Password: "wasspord",
		Role:     model.MANAGER,
		TeamID:   1,
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createTeam", nil, router)

	if w.Code != http.StatusConflict {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusConflict)
	}

	cleanUpDB()
}

// Tests that create team endpoint responds Status Bad Request
// when JSON is not correct.
func Test_handleCreateTeam_BadJSON(t *testing.T) {
	model.DBConn.Create(&model.User{
		Username: "kevin",
		Password: "wasspord",
		Role:     model.MANAGER,
		TeamID:   0,
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createTeam", nil, router)

	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}

	cleanUpDB()
}

// Tests that get-user endpoint responds http found and correct
// data when all conditions are met.
func Test_handleGetUser_Valid(t *testing.T) {
	userToGet := &model.User{
		ID:       5,
		TeamID:   3,
		Username: "Jimmy",
		Email:    "j@j.com",
		Password: "shouldn't matter",
		Role:     model.PLAYER,
		Position: model.QUARTERBACK,
		Photo:    "asdfasdfasdf",
	}
	model.DBConn.Create(userToGet)

	userInfoToGet := &model.UserPersonalInfo{
		ID:        5,
		Firstname: "Joe",
		Lastname:  "Luhrman",
		Height:    50,
		Weight:    50,
	}
	model.DBConn.Create(userInfoToGet)

	teamUserBelongsTo := &model.Team{
		ID:   3,
		Name: "Jumbos",
	}
	model.DBConn.Create(teamUserBelongsTo)

	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/get-user/Jimmy", nil, router)

	if w.Code != http.StatusFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusFound)
	}

	responseData := &userDataWrapper{}
	err := json.NewDecoder(w.Result().Body).Decode(responseData)
	if err != nil {
		t.Error(err)
	}

	originalData := &userDataWrapper{
		User:         *userToGet,
		PersonalInfo: *userInfoToGet,
		TeamName:     teamUserBelongsTo.Name,
	}
	originalData.User.Password = ""

	// marshall and unmarshall the data to remove non-json fields before
	// final comparison with response data
	originalDataJSON, _ := json.Marshal(originalData)

	originalDataUnMarsh := &userDataWrapper{}
	err = json.Unmarshal(originalDataJSON, originalDataUnMarsh)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(responseData)
	fmt.Println(originalDataUnMarsh)

	if *responseData != *originalDataUnMarsh {
		t.Error("structs are not equal")
	}

	cleanUpDB()
}

// Tests that get-user endpoint returns found and correct data
// when user doesn't belong to a team.
func Test_handleGetUser_NoTeam(t *testing.T) {
	userToGet := &model.User{
		ID:       5,
		Username: "Jimmy",
		Email:    "j@j.com",
		Password: "shouldn't matter",
		Role:     model.PLAYER,
		Position: model.QUARTERBACK,
		Photo:    "asdfasdfasdf",
	}
	model.DBConn.Create(userToGet)

	userInfoToGet := &model.UserPersonalInfo{
		ID:        5,
		Firstname: "Joe",
		Lastname:  "Luhrman",
		Height:    50,
		Weight:    50,
	}
	model.DBConn.Create(userInfoToGet)

	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/get-user/Jimmy", nil, router)

	if w.Code != http.StatusFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusFound)
	}

	responseData := &userDataWrapper{}
	err := json.NewDecoder(w.Result().Body).Decode(responseData)
	if err != nil {
		t.Error(err)
	}

	originalData := &userDataWrapper{
		User:         *userToGet,
		PersonalInfo: *userInfoToGet,
	}
	originalData.User.Password = ""

	// marshall and unmarshall the data to remove non-json fields before
	// final comparison with response data
	originalDataJSON, _ := json.Marshal(originalData)

	originalDataUnMarsh := &userDataWrapper{}
	err = json.Unmarshal(originalDataJSON, originalDataUnMarsh)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(responseData)
	fmt.Println(originalDataUnMarsh)

	if *responseData != *originalDataUnMarsh {
		t.Error("structs are not equal")
	}

	cleanUpDB()
}

// Tests that get-user endpoint returns bad request when the json isn't correct.
func Test_handleGetUser_BadURL(t *testing.T) {
	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/get-user", nil, router)

	if w.Code != http.StatusNotFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusNotFound)
	}
}

// Makes sure get-user endpoint responds http not found when user doesn't exist.
func Test_handleGetUser_NoUser(t *testing.T) {
	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/get-user/jasdasd", nil, router)

	if w.Code != http.StatusNotFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusNotFound)
	}
}

// Tests that create-promotion-to-manager-request responds Status Accepted when everything
// is valid.
func Test_handleCreatePromotionToManagerRequest_Valid(t *testing.T) {
	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
		Role:     model.PLAYER,
	}
	model.DBConn.Create(user)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	request := &model.PromotionToManagerRequest{
		Message: "UH OH",
	}

	jsonData, _ := json.Marshal(request)

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/create-promotion-to-manager-request",
		bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

// Tests that create-promotion-to-manager-request endpoint responds status conflict
// when user already has an open request.
func Test_handleCreatePromotionToManagerRequest_Conflict(t *testing.T) {
	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
		Role:     model.PLAYER,
	}
	model.DBConn.Create(user)

	conflictReq := &model.PromotionToManagerRequest{
		SenderID:       user.ID,
		SenderUsername: user.Username,
	}
	model.DBConn.Create(conflictReq)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/create-promotion-to-manager-request", nil, router)

	if w.Code != http.StatusConflict {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusConflict)
	}

	cleanUpDB()
}

// Makes sure create-promotion-to-manager-request endpoint responds bad request
// when the JSON is invalid.
func Test_handleCreatePromotionToManagerRequest_BadJSON(t *testing.T) {
	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
		Role:     model.PLAYER,
	}
	model.DBConn.Create(user)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/create-promotion-to-manager-request", nil, router)

	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}

	cleanUpDB()
}

// Test Retrieving a team with call to getTeam
func TestHandleGetTeam(t *testing.T) {
	// Insert team into DB
	team := model.Team {
		Name: "Halal Knots",
		TeamLocation: "Knott Hall",
	}

	result := model.DBConn.Create(&team)

	if result.Error != nil {
		t.Error(result.Error)
	}

	// Find team based on its name so we know what ID it has
	teamQuery := model.Team{}
	err := model.DBConn.Where("name = ?", "Halal Knots").First(&teamQuery).Error

	if err != nil {
		t.Error(err)
	}

	// Now, we send the get request and check if we can retrieve the team with the ID we inserted
	router := setupTestRouter()
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getTeam/" + strconv.FormatUint(uint64(teamQuery.ID), 10), nil, router)

	if w.Code != http.StatusFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusFound)
	}

	cleanUpDB()
}

// Send a request to getTeam with an ID that doesn't exist
func TestGetTeamInvalidId(t *testing.T) {
	// We send the get request and check if we get the proper error for an invalid ID
	router := setupTestRouter()
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getTeam/69", nil, router)

	if w.Code != http.StatusNotFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusNotFound)
	}

	cleanUpDB()
}

// Send a request to getTeam with an ID that is formatted incorrectly
func TestGetTeamBadFormat(t *testing.T) {
	// We send the get request and check if we get the proper error for an invalid ID
	router := setupTestRouter()
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getTeam/sixtynine", nil, router)

	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}

	cleanUpDB()
}
