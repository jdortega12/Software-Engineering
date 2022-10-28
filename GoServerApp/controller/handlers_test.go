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
func TestHandleLogout(t *testing.T) {
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

// Test login for user that actually exists.
func TestHandleLoginGoodCredentials(t *testing.T) {
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

// Test Login for user that doesn't exist.
func TestHandleLoginBadCredentials(t *testing.T) {
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

// Tests that login returns HTTP Status Bad Request
// when JSON is not correct.
func TestHandleLoginBadJSON(t *testing.T) {
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

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer([]byte("bad data")))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	fmt.Println(w.Code)

	if w.Code != http.StatusBadRequest {
		t.FailNow()
	}
}

// Tests that CreateTeamRequest is able to properly recieve a team request
// Test if function can succesfully call model to insert and send JSON indicating success to frontend
func TestHandleCreateTeamRequest(t *testing.T) {
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
		handleCreateTeamRequest(ctx)
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
func TestHandleCreateTeamRequestBad(t *testing.T) {
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
		handleCreateTeamRequest(ctx)
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
func TestHandleUpdateUserPersonalInfo(t *testing.T) {
	// generic user for this test
	model.DBConn, _ = model.InitDB(TEST_DB_PATH)
	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
	}
	model.DBConn.Create(user)
	defer model.DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

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

// Tests that update info handler returns HTTP Status Unauthorized
// when session doesn't exist.
func TestHandleUpdateUserPersonalInfoNilSession(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	testStore := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", testStore))

	SetupHandlers(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/updatePersonalInfo", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.FailNow()
	}
}

// Make sure update info handler returns correct status code for bad request.
func TestHandleUpdateUserPersonalInfoBadJSON(t *testing.T) {
	// generic user for this test
	model.DBConn, _ = model.InitDB(TEST_DB_PATH)
	user := &model.User{
		Username: "jaluhrman",
		Password: "ween",
	}
	model.DBConn.Create(user)
	defer model.DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

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

func TestHandleCreateAccount(t *testing.T) {
	var err error
	model.DBConn, err = model.InitDB(TEST_DB_PATH)
	if err != nil {
		panic(err)
	}
	defer model.DBConn.Exec("DELETE FROM users")
	defer model.DBConn.Exec("DELETE FROM user_personal_infos")

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	SetupHandlers(router)

	testUser := &model.User{
		Username: "jdo",
		Email:    "jdo@gmail.com",
		Password: "123",
	}

	// format as JSON
	reader, err := json.Marshal(testUser)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	//Mock request
	req, _ := http.NewRequest("POST", "/api/v1/createAccount", bytes.NewBuffer(reader))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.FailNow()
	}
}

// Test whether a Photo can be inserted into the database
func TestHandleCreatePhoto(t *testing.T) {
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

// Test whether the correct response is given for an invalid call to CreatePhoto
func TestHandleCreatePhotoInvalid(t *testing.T) {
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

func TestHandleCreateTeam(t *testing.T) {
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
	model.DBConn.Create(user)
	defer model.DBConn.Unscoped().Where("id = ?", user.ID).Delete(user)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	SetupHandlers(router)

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
