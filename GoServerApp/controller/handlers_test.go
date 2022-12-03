package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

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

	//cleanUpDB()
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

// Tests get promotion to manager requests endpoint responds correct status
// code when all conditions are met
func Test_handleGetPromotionToManagerRequests_Valid(t *testing.T) {
	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.DBConn.Create(admin)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/promotion-to-manager-requests", nil, router)

	if w.Code != http.StatusFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusFound)
	}

	cleanUpDB()
}

// Makes sure get promotion to manager requests endpoint correctly rejects a non-admin.
func Test_handleGetPromotionToManagerRequests_NotAdmin(t *testing.T) {
	notAdmin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.MANAGER,
	}
	model.DBConn.Create(notAdmin)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, notAdmin.Username, notAdmin.Password)
	})

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/promotion-to-manager-requests", nil, router)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusUnauthorized)
	}

	cleanUpDB()
}

// Test Retrieving a team with call to getTeam
func TestHandleGetTeam(t *testing.T) {
	// Insert team into DB
	team := model.Team{
		Name:         "Halal Knots",
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

	// Insert manager with that team into DB so that we don't have an error
	user := model.User{
		Username: "Joe Douglas",
		Password: "allgasnobreaks",
		Role:     model.MANAGER,
		TeamID:   teamQuery.ID,
	}

	model.DBConn.Create(&user)

	// Now, we send the get request and check if we can retrieve the team with the ID we inserted
	router := setupTestRouter()
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getTeam/"+strconv.FormatUint(uint64(teamQuery.ID), 10), nil, router)

	if w.Code != http.StatusAccepted {
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

// send a request to getTeam for a team that does not have a manager
func TestGetTeamNoManager(t *testing.T) {
	// Insert team into DB
	team := model.Team{
		Name:         "Halal Knots",
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
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getTeam/"+strconv.FormatUint(uint64(teamQuery.ID), 10), nil, router)

	if w.Code != http.StatusNotFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusNotFound)
	}

	cleanUpDB()
}

// Handle get teams valid
func TestHandleGetTeams(t *testing.T) {
	team1 := model.Team{
		Name: "Desk",
	}

	team2 := model.Team{
		Name: "Chair",
	}

	model.DBConn.Create(&team1)
	model.DBConn.Create(&team2)

	router := setupTestRouter()
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getTeams", nil, router)

	if w.Code != http.StatusAccepted {
		t.Error("Unable to retrieve teams")
	}

	cleanUpDB()
}

// Handle Get teams when there are no teams
func TestHandleGetTeamsInvalid(t *testing.T) {
	router := setupTestRouter()
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getTeams", nil, router)

	if w.Code != http.StatusAccepted {
		t.Error("Unable to retrieve teams")
	}

	cleanUpDB()
}

// sends data to handler and receives status code
// Invalid entries are caught by user.go error checking
func TestAcceptPlayerValid(t *testing.T) {
	//Add player and manager to DB
	player := &model.User{
		Username: "player",
		Password: "crap",
		Role:     model.PLAYER,
	}
	result := model.DBConn.Create(player)
	if result.Error != nil {
		t.Error(result.Error)
	}

	manager := &model.User{
		Username: "manager",
		TeamID:   1,
		Password: "crapp",
		Role:     model.MANAGER,
	}
	result = model.DBConn.Create(manager)
	if result.Error != nil {
		t.Error(result.Error)
	}

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "manager", "crapp")
		ctx.Next()
	})

	data := &model.AcceptData{
		PlayerName:  "player",
		ManagerName: "manager",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Errorf("could not marshal json: %s\n", err)
	}

	// mock request
	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/acceptPlayer", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

// Insert Players into the DB and check if they can be retrieved by TeamID in the frontend
func TestHandleGetTeamPlayers(t *testing.T) {
	player := model.User{
		Username: "saucegardner",
		Password: "allgasnobrakes",
		TeamID:   1,
	}

	player2 := model.User{
		Username: "zachwilson",
		Password: "gocougars",
		TeamID:   1,
	}

	model.DBConn.Create(&player)
	model.DBConn.Create(&player2)

	router := setupTestRouter()
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getTeamPlayers/1", nil, router)

	if w.Code != http.StatusAccepted {
		t.Error("Unable to retrieve players")
	}

	cleanUpDB()
}

// See if we can retrieve playoff teams after we insert teams and matches
func TestHandleGetPlayoffs(t *testing.T) {
	// first we must insert teams
	team_names := []string{"gormgobblers", "gormfighters", "gormlovers", "gormhaters", "gorm4lyfe", "earthgorm", "gormstormers", "gormgorm", "outofgormstuff", "lastgorm"}

	for team := range team_names {
		temp := model.Team{
			Name: team_names[team],
		}

		model.DBConn.Create(&temp)
	}

	// Everything that can be the same between matches
	match_type := model.REGULAR
	location := "GORM Stadium"
	start_time := time.Now()

	// Only scores and ID's vary
	home_scores := []uint{10, 9, 8, 7, 6}
	away_scores := []uint{11, 12, 13, 14, 15}
	home_ids := []uint{1, 2, 3, 4, 5}
	away_ids := []uint{6, 7, 8, 9, 10}

	// Now, we insert the matches
	for index := range home_scores {
		temp := model.Match{
			MatchType:     match_type,
			Location:      location,
			StartTime:     start_time,
			HomeTeamScore: home_scores[index],
			AwayTeamScore: away_scores[index],
			HomeTeamID:    home_ids[index],
			AwayTeamID:    away_ids[index],
		}

		model.DBConn.Create(&temp)
	}

	// Finally, we extract the playoff teams and see if they are the correct ones
	router := setupTestRouter()
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getPlayoffs", nil, router)

	if w.Code != http.StatusAccepted {
		t.Error("something went wrong")
	}

	responseData := make([]string, 0)
	err := json.NewDecoder(w.Result().Body).Decode(&responseData)

	if err != nil {
		t.Error(err)
	}

	// check for correct teams. It should be the last 8 names
	for i := range [5]int{} {
		target := team_names[len(team_names)-i-1]
		for j := range responseData {
			if responseData[j] == target {
				break
			}
			if j == len(responseData)-1 {
				t.Error("we were supposed to find this")
			}
		}

	}

	cleanUpDB()
}

// Makes sure that correct status code is returned when all conditions are met.
func Test_handleStartMatch_Valid(t *testing.T) {
	defer cleanUpDB()

	// create test admin
	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.DBConn.Create(admin)

	// two test teams
	teams := make([]*model.Team, 2)

	// create two test teams
	for i := range teams {
		teams[i] = &model.Team{
			Name:         "team " + fmt.Sprint(i),
			TeamLocation: "location " + fmt.Sprint(i),
		}
		model.DBConn.Create(teams[i])
	}

	// add middleware to set session to admin
	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})

	wrapper := &startMatchWrapper{
		HomeTeamName: teams[0].Name,
		AwayTeamName: teams[1].Name,
	}

	jsonData, _ := json.Marshal(wrapper)
	buffer := bytes.NewBuffer(jsonData)

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/start-match", buffer, router)
	if w.Code != http.StatusCreated {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusCreated)
	}
}

// Test that correct status is returned when JSON is nil.
func Test_handleStartMatch_BadJSON(t *testing.T) {
	defer cleanUpDB()

	// create test admin
	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.DBConn.Create(admin)

	// add middleware to set session to admin
	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/start-match", nil, router)
	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}
}

// Test correct status is returned when one of the teams doesn't exist.
func Test_handleStartMatch_BadTeam(t *testing.T) {
	defer cleanUpDB()

	// create test admin
	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.DBConn.Create(admin)

	// create only one team
	team := &model.Team{
		Name:         "team 1",
		TeamLocation: "location 1",
	}
	model.DBConn.Create(team)

	// add middleware to set session to admin
	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})

	wrapper := &startMatchWrapper{
		HomeTeamName: team.Name,
		AwayTeamName: team.Name + "gabagool", // doesn't exist
	}

	jsonData, _ := json.Marshal(wrapper)
	buffer := bytes.NewBuffer(jsonData)

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/start-match", buffer, router)
	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}
}

func Test_handleFinishMatch_Valid(t *testing.T) {
	defer cleanUpDB()

	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.DBConn.Create(admin)

	match := &model.Match{}
	model.DBConn.Create(match)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})

	jsonData, _ := json.Marshal(&model.Match{
		ID: match.ID,
	})
	body := bytes.NewBuffer(jsonData)

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/finish-match", body, router)
	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}
}

func Test_handleFinishMatch_DoesntExist(t *testing.T) {
	defer cleanUpDB()

	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.DBConn.Create(admin)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})

	jsonData, _ := json.Marshal(&model.Match{
		ID: 1,
	})
	body := bytes.NewBuffer(jsonData)

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/finish-match", body, router)
	if w.Code != http.StatusNotFound {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusNotFound)
	}
}

func Test_handleFinishMatch_BadJSON(t *testing.T) {
	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.DBConn.Create(admin)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/finish-match", nil, router)
	if w.Code != http.StatusBadRequest {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusBadRequest)
	}
}

func TestHandleGetMatch(t *testing.T) {
	cleanUpDB()
	match := model.Match{
		MatchType:     model.REGULAR,
		Location:      "Knott Hall",
		StartTime:     time.Now(),
		InProgress:    true,
		Quarter:       uint(1),
		QuarterTime:   time.Date(0, 0, 0, 0, 15, 0, 0, time.FixedZone("UTC-7", 0)),
		HomeTeamID:    1,
		AwayTeamID:    2,
		HomeTeamScore: 0,
		AwayTeamScore: 0,
		Likes:         0,
		Dislikes:      0,
	}

	model.DBConn.Create(&match)
	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getMatch/1", nil, router)

	if w.Code == http.StatusBadRequest {
		t.Error("Bad status")
	}

	cleanUpDB()
}

func TestGetMatchInvalid(t *testing.T) {
	cleanUpDB()
	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getMatch/1", nil, router)

	if w.Code != http.StatusNotFound {
		t.Error("Bad status")
	}
}

func TestFindComments(t *testing.T){
	comment := model.Comment {
		Message: "Hello",
		MatchID: "12",
		UserID: 1,
	}
	model.DBConn.Create(&comment)
	comment2 := model.Comment {
		Message: "World",
		MatchID: "12",
		UserID: 2,
	}
	model.DBConn.Create(&comment2)
	router := setupTestRouter()

	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getComments/12", nil, router)

	if w.Code == http.StatusBadRequest {
		t.Error("Bad status")
	}
	
	cleanUpDB()
}

func TestCreateCommentValid(t *testing.T){
	model.DBConn.Create(&model.User{
		Username: "kevin",
		Password: "wasspord",
		ID: 5,
		Role: model.PLAYER,
	})

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")
		ctx.Next()
	})

	comment := model.Comment {
		Message: "Hello",
		MatchID: "12",
		UserID: 1,
	}

	jsonData, _ := json.Marshal(comment)
	
	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createComment", bytes.NewBuffer(jsonData), router)

	if w.Code != 200 {
		t.Error("wrong code")
	}
}

func Test_HandleGetUserTeamData(t *testing.T) {

	team := &model.Team{
		ID:   1,
		Name: "Jaymins",
	}
	model.CreateTeam(team)

	user := &model.User{
		ID:       1,
		TeamID:   1,
		Username: "jaymin",
		Password: "123",
	}
	model.CreateUser(user)

	personalInfo := &model.UserPersonalInfo{
		ID:        1,
		Firstname: "Jaymin",
		Lastname:  "Ortega",
	}
	model.UpdateUserPersonalInfo(personalInfo)

	//Create admin for session login
	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.CreateUser(admin)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})
	w := sendMockHTTPRequest(http.MethodGet, "/api/v1/getUserTeamData", nil, router)

	if w.Code != http.StatusAccepted {
		t.Error("Unable to retrieve user team data")
	}

	cleanUpDB()

}

// test valid roster change post
func Test_HandleChangeRoster(t *testing.T) {

	team := &model.Team{
		ID:   1,
		Name: "Jaymins",
	}
	model.CreateTeam(team)

	team2 := &model.Team{
		ID:   2,
		Name: "Ortegas",
	}
	model.CreateTeam(team2)

	user := &model.User{
		ID:       1,
		TeamID:   1,
		Username: "jaymin",
		Password: "123",
	}
	model.CreateUser(user)

	personalInfo := &model.UserPersonalInfo{
		ID:        1,
		Firstname: "Jaymin",
		Lastname:  "Ortega",
	}
	model.UpdateUserPersonalInfo(personalInfo)

	returnData := &model.UserTeamReturnData{
		UserId:   1,
		Teamname: "Ortegas",
	}

	jsonData, err := json.Marshal(returnData)
	if err != nil {
		t.Errorf("could not marshal json: %s\n", err)
	}

	//Create admin for session login
	admin := &model.User{
		Username: "jaluhrman",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.CreateUser(admin)
	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, admin.Username, admin.Password)
	})

	// mock request
	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/changeRoster", bytes.NewBuffer(jsonData), router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

func TestCreateCommentInvalid(t *testing.T){
	comment := model.Comment {
		Message: "Hello",
		MatchID: "12",
		UserID: 1,
	}

	router := setupTestRouter()

	jsonData, _ := json.Marshal(comment)
	
	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/createComment", bytes.NewBuffer(jsonData), router)

	if w.Code != 401 {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

func TestAddLikeValid(t *testing.T){
	model.DBConn.Create(&model.User{
		Username: "kevin",
		Password: "wasspord",
		ID: 5,
		Role: model.PLAYER,
	})

	match := model.Match {
		MatchType: model.REGULAR, 
		Location: "Knott Hall",
		StartTime: time.Now(),
		InProgress: true,
		Quarter: uint(1),
		QuarterTime: time.Date(0, 0, 0, 0, 15, 0, 0, time.FixedZone("UTC-7", 0)),
		HomeTeamID: 1,
		AwayTeamID: 2,
		HomeTeamScore: 0,
		AwayTeamScore: 0,
		Likes: 0,
		Dislikes: 0,
	}

	model.DBConn.Create(&match)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/postLikes/1", nil, router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}

func TestAddDislikeValid(t *testing.T){
	model.DBConn.Create(&model.User{
		Username: "kevin",
		Password: "wasspord",
		ID: 5,
		Role: model.PLAYER,
	})

	match := model.Match {
		MatchType: model.REGULAR, 
		Location: "Knott Hall",
		StartTime: time.Now(),
		InProgress: true,
		Quarter: uint(1),
		QuarterTime: time.Date(0, 0, 0, 0, 15, 0, 0, time.FixedZone("UTC-7", 0)),
		HomeTeamID: 1,
		AwayTeamID: 2,
		HomeTeamScore: 0,
		AwayTeamScore: 0,
		Likes: 0,
		Dislikes: 0,
	}

	model.DBConn.Create(&match)

	router := setupTestRouter(func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")
		ctx.Next()
	})

	w := sendMockHTTPRequest(http.MethodPost, "/api/v1/postDislikes/1", nil, router)

	if w.Code != http.StatusAccepted {
		t.Errorf("code was %d, should have been %d", w.Code, http.StatusAccepted)
	}

	cleanUpDB()
}
