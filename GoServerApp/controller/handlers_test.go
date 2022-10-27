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

	//set test session
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

	//create data
	data := map[string]interface{}{
		"email":    "jdo@gmail.com",
		"password": "123",
	}

	//format as JSON
	reader, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	//Login
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(reader))
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

	//create data
	data := map[string]interface{}{
		"email":    "WRONG",
		"password": "WRONG",
	}

	//format as JSON
	reader, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(reader))
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

	user := &model.User{
		Username: "timba11",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	user2 := &model.User{
		Username: "paulba11",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	model.CreateUser(user)
	model.CreateUser(user2)

	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "timba11", "123")
		CreateTeamRequest(ctx)
	})

	// mock request
	var jsonStr = []byte(`{"Message": "Hello World", "ReceiverID": "paulba11"}`)
	req, _ := http.NewRequest(http.MethodPost, "/testWrapper", bytes.NewBuffer(jsonStr))

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

	if err != nil {
		panic(err)
	}

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	SetupHandlers(router)
	store := cookie.NewStore([]byte("placeholder"))
	router.Use(sessions.Sessions("session", store))

	user := &model.User{
		Username: "tom",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	user2 := &model.User{
		Username: "john",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	model.CreateUser(user)
	model.CreateUser(user2)

	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "tom", "123")
		CreateTeamRequest(ctx)
	})

	// mock request
	var jsonStr = []byte(`{"Message": "Hello World", "ReceiverID": "l"}`)
	req, _ := http.NewRequest(http.MethodPost, "/testWrapper", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var data map[string]string

	if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
		fmt.Println("got here")
		t.FailNow()
	}

	fmt.Println(data)

	if w.Code != http.StatusUnauthorized {
		t.FailNow()
	}
}

// Make sure update info handler returns correct status code.
func TestUpdatePersonalInfoHandler(t *testing.T) {
	// generic user for this test
	model.DBConn, _ = model.InitDB(TEST_DB_PATH)
	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
	}
	model.DBConn.Create(user)
	defer model.DBConn.Unscoped().Where("user_id = ?", user.UserID).Delete(user)

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
	defer model.DBConn.Unscoped().Where("user_personal_info_id = ?", user.UserID).Delete(info)

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
	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
	}
	model.DBConn.Create(user)
	defer model.DBConn.Unscoped().Where("user_id = ?", user.UserID).Delete(user)

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
	data := map[string]interface{}{
		"username": "jdo",
		"email":    "jdo@gmail.com",
		"password": "123",
	}

	//format as JSON
	reader, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	//Mock request
	req, _ := http.NewRequest("POST", "/api/v1/createAccount", bytes.NewBuffer(reader))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Clean up
	model.DBConn.Unscoped().Where("username = ?", "jdo").Delete(&model.User{})

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}
}

// Test whether a Photo can be inserted into the database
func TestCreatePhoto(t *testing.T) {
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
		CreatePhoto(ctx)
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

// Test whether the correct response is given for an invalid call to CreatePhoto
func TestCreatePhotoInvalid(t *testing.T) {
	// generic user for this test
	model.DBConn, _ = model.InitDB(TEST_DB_PATH)
	model.DBConn.Create(&model.User{
		Username: "jaluhrman2",
		Password: "ween",
	})

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	// test endpoint just to set session for this test
	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")
		CreatePhoto(ctx)
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

func TestGoodCreateTeam(t *testing.T) {
	var err error
	model.DBConn, err = model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}
	user := &model.User{
		Username: "kevin",
		Password: "wasspord",
		Role:     model.MANAGER,
	}
	defer model.DBConn.Unscoped().Where("user_id = ?", user.UserID).Delete(user)
	model.DBConn.Create(user)
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))
	SetupHandlers(router)
	router.POST("/testWrapper", func(ctx *gin.Context) {
		setSessionUser(ctx, "kevin", "wasspord")

		CreateTeamHandler(ctx)
	})

	info := &model.Team{
		Name:         "Greyhounds",
		TeamLocation: "Baltimore",
	}
	json_info, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	defer model.DBConn.Unscoped().Where("team_id = ?", info.TeamID).Delete(&model.Team{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/testWrapper", bytes.NewBuffer(json_info))
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		fmt.Println("Http status failed")
		t.FailNow()
	}

}
