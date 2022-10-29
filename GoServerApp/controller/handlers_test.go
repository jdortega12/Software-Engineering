package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
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

// Sets up a router for testing purposes.
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.Default()

	testStore := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", testStore))

	SetupHandlers(router)

	return router
}

// Tests that handleLogout() responds HTTP Status Reset Content.
func Test_handleLogout(t *testing.T) {
	router := setupTestRouter()

	// mock request
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/logout", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusResetContent {
		t.FailNow()
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
	defer model.DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	// create mock login credentials
	userCpy := &model.User{
		Username: "jdo",
		Password: "123",
	}

	// format as JSON
	reader, err := json.Marshal(userCpy)
	if err != nil {
		t.Errorf("Error: %s\n", err)
		return
	}

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(reader))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}

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
		return
	}

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(reader))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Println(w.Code)

	if w.Code != http.StatusUnauthorized {
		t.FailNow()
	}
}

// Tests that handleLogin() responds with HTTP Status Bad Request
// when JSON is not correct.
func Test_handleLogin_BadJSON(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer([]byte("bad data")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Println(w.Code)

	if w.Code != http.StatusBadRequest {
		t.FailNow()
	}
}

// Tests handleCreateTeamNotification() when a valid teamNotification
// JSON is sent over and all success conditions are met. Should return
// HTTP Status Accepted.
func Test_handleCreateTeamNotification_ValidInvite(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
		Role:     model.MANAGER,
	})
	model.DBConn.Create(&model.User{
		Username: "colbert",
		Role:     model.PLAYER,
	})

	router.POST("/test", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		handleCreateTeamNotification(ctx)
	})

	notification := &model.TeamNotification{
		SenderUsername:   "jaluhrman",
		ReceiverUsername: "colbert",
	}

	jsonData, _ := json.Marshal(notification)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusAccepted)
	}

	model.DBConn.Exec("DELETE FROM users")
	model.DBConn.Exec("DELETE FROM notifications")
}

// Tests case that handleCreateTeamNotification() endpoint is
// requested when a user is not logged in/there is no valid
// session. Should respond HTTP Status Unauthorized.
func Test_handleCreateTeamNotification_NoSession(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/createTeamRequest", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusUnauthorized)
	}
}

// Tests case that handleCreateTeamNotification() is called
// when a session exists but the user is not valid. Should respond
// HTTP Status Unauthorized.
func Test_handleCreateTeamNotification_InvalidUser(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	router.POST("/test", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		handleCreateTeamNotification(ctx)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusUnauthorized)
	}
}

// Tests case that handleCreateTeamNotification() is called without a valid
// JSON request body in context. Should respond HTTP Status Bad Request.
func Test_handleCreateTeamNotification_BadJSON(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	router.POST("/test", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		handleCreateTeamNotification(ctx)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusBadRequest)
	}

	model.DBConn.Exec("DELETE FROM users")
}

// Tests case that handleCreateTeamNotification() is called when the
// SenderUsername of the TeamNotification is not the same as the logged
// in user. Should respond HTTP Status Bad Request.
func Test_handleCreateTeamNotification_BadSender(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	router.POST("/test", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		handleCreateTeamNotification(ctx)
	})

	notification := &model.TeamNotification{
		SenderUsername:   "not_jaluhrman",
		ReceiverUsername: "colbert",
	}

	jsonData, _ := json.Marshal(notification)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Status code: %d, should have been %d", w.Code, http.StatusBadRequest)
	}

	model.DBConn.Exec("DELETE FROM users")
}

// Tests that handleUpdataeUserPersonalInfo() responds HTTP Status Accepted
// when everything is valid.
func Test_handleUpdateUserPersonalInfo_Valid(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
	}
	model.DBConn.Create(user)
	defer model.DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	// test endpoint just to set session for this test
	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		handleUpdateUserPersonalInfo(ctx)
	})

	info := &model.UserPersonalInfo{
		Firstname: "Joe",
		Lastname:  "Luhrman",
		Height:    50,
		Weight:    240,
	}
	defer model.DBConn.Unscoped().Where("id = ?", user.ID).Delete(info)

	json_info, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testWrapper", bytes.NewBuffer(json_info))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}
}

// Tests that handleUpdataeUserPersonalInfo() responds HTTP Status Unauthorized
// when session doesn't exist.
func Test_handleUpdateUserPersonalInfo_NilSession(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/updatePersonalInfo", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.FailNow()
	}
}

// Tests handleUpdataeUserPersonalInfo() responds HTTP Status Bad Request
// when the JSON is not correct.
func Test_handleUpdateUserPersonalInfo_BadJSON(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
	}
	model.DBConn.Create(user)
	defer model.DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	// test endpoint just to set session for this test
	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		handleUpdateUserPersonalInfo(ctx)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testWrapper", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.FailNow()
	}
}

// Tests handleCreateAccount() responds HTTP Status Accepted when
// all conditions are met.
func Test_handleCreateAccount_Valid(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	defer model.DBConn.Exec("DELETE FROM users")
	defer model.DBConn.Exec("DELETE FROM user_personal_infos")

	testUser := &model.User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	// format as JSON
	reader, err := json.Marshal(testUser)
	if err != nil {
		t.Errorf("could not marshal json: %s\n", err)
		return
	}

	// mock request
	req, _ := http.NewRequest("POST", "/api/v1/createAccount", bytes.NewBuffer(reader))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}
}

// Tests handleCreatePhoto() returns HTTP Status Accepted
// when the request is correct.
func Test_handleCreatePhoto_Valid(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	// test endpoint just to set session for this test
	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		handleCreatePhoto(ctx)
	})

	info := gin.H{
		"photo": "dfjekjfcks",
	}

	json_info, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testWrapper", bytes.NewBuffer(json_info))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}
}

// Tests that handleCreatePhoto() responds with HTTP Status Bad
// Request when the request doesn't have the correct key value pair.
func Test_handleCreatePhoto_Invalid(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	model.DBConn.Create(&model.User{
		Username: "jaluhrman2",
		Password: "ween",
	})

	// test endpoint just to set session for this test
	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		handleCreatePhoto(ctx)
	})

	info := gin.H{
		"thisfieldisntrelevant": "dfjekjfcks",
	}

	json_info, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testWrapper", bytes.NewBuffer(json_info))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.FailNow()
	}
}

// Tests that handleCreateTeam() responds with HTTP Status Accepted
// when all conditions are correct.
func Test_handleCreateTeam_Valid(t *testing.T) {
	initTestDB()
	router := setupTestRouter()

	user := &model.User{
		Username: "kevin",
		Password: "wasspord",
		Role:     model.MANAGER,
	}
	model.DBConn.Create(user)
	defer model.DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")

		handleCreateTeam(ctx)
	})

	info := &model.Team{
		Name:         "Greyhounds",
		TeamLocation: "Baltimore",
	}
	defer model.DBConn.Unscoped().Where("name = ?", "Greyhounds").Delete(info)

	json_info, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testWrapper", bytes.NewBuffer(json_info))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		fmt.Println("Http status failed")
		t.FailNow()
	}
}
