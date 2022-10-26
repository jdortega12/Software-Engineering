package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	TEST_DB_PATH = "../test.db"
)

// Tests that Logout() returns correct status code.
func TestLogout(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	SetupHandlers(router)

	// mock request
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/logout", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusResetContent {
		t.FailNow()
	}
}

// Test login for existing user with correct credentials
func TestLogin(t *testing.T) {
	var err error
	model.DBConn, err = model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	SetupHandlers(router)

	//Add User to DB
	user := &model.User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}
	model.CreateUser(user)
	defer model.DBConn.Unscoped().Where("user_id = ?", user.UserID).Delete(user)

	//Login
	reader := strings.NewReader("username=jdo&password=123")
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}

}

// Test Login with improper username and password
func TestImproperLogin(t *testing.T) {
	var err error
	model.DBConn, err = model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	SetupHandlers(router)

	//Add User to DB
	user := &model.User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}
	model.CreateUser(user)
	defer model.DBConn.Unscoped().Where("user_id = ?", user.UserID).Delete(user)

	reader := strings.NewReader("username=Wrong&password=Wrong")
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Println(w.Code)

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}
}

// Tests that CreateTeamRequest is able to properly recieve a team request
// Test if function can succesfully call model to insert and send JSON indicating success to frontend
func TestGoodRequestInsert(t *testing.T) {
	var err error
	model.DBConn, err = model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	SetupHandlers(router)
	store := cookie.NewStore([]byte("placeholder"))
	router.Use(sessions.Sessions("session", store))

	// mock request
	var jsonStr = []byte(`{"Message": "Hello World", "SenderID" : "1", "ReceiverID": "2"}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/createTeamRequest", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var data map[string]string

	if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
		fmt.Println("got here")
		t.FailNow()
	}

	fmt.Println(data)

	if data["Created-notification"] != "true" {
		t.FailNow()
	}
}

// Tests that CreateTeamRequest is able to handle in incorrect request
// Test if function can succesfully call model to insert and send JSON indicating failure to frontend
func TestBadRequestInsert(t *testing.T) {
	var err error
	model.DBConn, err = model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	SetupHandlers(router)

	// mock request
	var jsonStr = []byte(`{"hi": "hello"}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/createTeamRequest", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var data map[string]string

	if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
		fmt.Println("got here")
		t.FailNow()
	}

	fmt.Println(data)

	if data["Created-notification"] != "false" {
		t.FailNow()
	}
}

// Make sure update info handler returns correct status code.
func TestUpdatePersonalInfoHandler(t *testing.T) {
	// generic user for this test
	model.DBConn, _ = model.InitDB(TEST_DB_PATH)
	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	// test endpoint just to set session for this test
	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		UpdateUserPersonalInfoHandler(ctx)
	})

	info := &model.UserPersonalInfo{
		Firstname: "Joe",
		Lastname:  "Luhrman",
		Height:    50,
		Weight:    240,
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

// Make sure update info handler returns correct status code for bad request.
func TestUpdatePersonalInfoHandlerBadJSON(t *testing.T) {
	// generic user for this test
	model.DBConn, _ = model.InitDB(TEST_DB_PATH)
	model.DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	// test endpoint just to set session for this test
	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		UpdateUserPersonalInfoHandler(ctx)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testWrapper", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.FailNow()
	}
}

func TestCreateAccount(t *testing.T) {
	var err error
	model.DBConn, err = model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	SetupHandlers(router)

	//create data
	reader := strings.NewReader("username=jdo&email=jdo@gmail.com&password=123")
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/createAccount", reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	model.DBConn.Unscoped().Where("username = ?", "jdo").Delete(&model.User{})

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}
}
