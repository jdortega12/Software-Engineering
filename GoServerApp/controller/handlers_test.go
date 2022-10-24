package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"encoding/json"
	"fmt"
	"jdortega12/Software-Engineering/GoServerApp/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	TDBConn, err := gorm.Open(sqlite.Open("../test.db"))

	if err != nil {
		panic(err)
	}
	
	model.InitDBTest(TDBConn)

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
	TDBConn, err := gorm.Open(sqlite.Open("../test.db"))
	if err != nil {
		panic(err)
	}
	model.InitDBTest(TDBConn)

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
