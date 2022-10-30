package controller

import (
	"bytes"
	"encoding/json"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"
	"net/http/httptest"
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
	model.DBConn.Exec("DELETE FROM users_personal_infos")
	model.DBConn.Exec("DELETE FROM teams")
	model.DBConn.Exec("DELETE FROM team_notifications")
	model.DBConn.Exec("DELETE FROM matches")
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
	initTestDB()
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
	initTestDB()
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
	initTestDB()
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
	initTestDB()

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
	initTestDB()
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
	initTestDB()

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
	initTestDB()

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
	initTestDB()

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
	initTestDB()
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
	initTestDB()

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
	initTestDB()

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

// Tests that handleCreateTeam() responds with HTTP Status Accepted
// when all conditions are correct.
func Test_handleCreateTeam_Valid(t *testing.T) {
	initTestDB()

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
