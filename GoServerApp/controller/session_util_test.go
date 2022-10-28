package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Tests whether clearSession() properly clears an arbitrary
// session inside a default *gin.Context struct. Uses anonymous
// func endpoint and a mock GET request which is necessary to
// set and clear a session.
func TestClearSession(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	testStore := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", testStore))

	router.GET("/test", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Set("test_val_1", "doesnt matter")
		session.Set("test_val_2", "also doesnt matter")
		session.Save()

		clearSession(ctx)

		if session.Get("test_val_1") != nil || session.Get("test_val_2") != nil {
			t.FailNow()
		}
	})

	// mock request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
}

// Test that setSession sets the session correctly.
func TestSetSession(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	testStore := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", testStore))

	router.GET("/test", func(ctx *gin.Context) {
		setSessionUser(ctx, "test_username", "test_password")

		test_session := sessions.Default(ctx)

		if test_session.Get("username") != "test_username" ||
			test_session.Get("password") != "test_password" {
			t.FailNow()
		}
	})

	// mock request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
}

// Check that getSessionUser works when session has been
// set correctly.
func TestGetSessionUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	testStore := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", testStore))

	router.GET("/test", func(ctx *gin.Context) {
		testSession := sessions.Default(ctx)
		testSession.Set("username", "ween")
		testSession.Set("password", "feen")

		username, password, exists := getSessionUser(ctx)

		if exists != true {
			t.FailNow()
		}

		if username != "ween" || password != "feen" {
			t.FailNow()
		}
	})

	// mock request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
}

// Test that getSessionUser works in the case that
// the session has not been set/doesn't exist.
func TestGetSessionUserNil(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	testStore := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", testStore))

	router.GET("/test", func(ctx *gin.Context) {
		username, password, exists := getSessionUser(ctx)

		if exists == true {
			t.Error("getSessionUser should return false when session doesn't exist")
			t.FailNow()
		}

		if username != "" || password != "" {
			t.Error("getSessionUser should retun username and password as empty" +
				"strings when session doesn't exist")
			t.FailNow()
		}
	})

	// mock request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
}
