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

/*
func TestUpdatePersonalInfoHandler(t *testing.T) {
	// generic user for this test
	DBConn, _ := model.InitDB(TEST_DB_PATH)
	DBConn.Create(&model.User{
		Username: "jaluhrman",
		Password: "ween",
	})

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	SetupHandlers(router)

	// test endpoint just to set session for this test
	router.POST("/setTestSession", func(ctx *gin.Context) {
		setSessionUser(ctx, "jaluhrman", "ween")

		// mock request to actually test endpoint
		var json = []byte(`{
		"firstname": "Joe",
		"lastname": "Luhrman",
		"height": "50",
		"weight": "240",
		}`)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/updatePersonalInfo", bytes.NewBuffer(json))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.FailNow()
		}
	})

	// mock request to set session
	req, _ := http.NewRequest(http.MethodPost, "/setTestSession", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	http.
}
*/
