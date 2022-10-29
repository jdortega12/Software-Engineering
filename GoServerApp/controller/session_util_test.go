package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Tests whether clearSession() properly clears an arbitrary session.
func Test_clearSession(t *testing.T) {
	router := setupTestRouter()

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

// Test that setSessionUser() sets the session correctly.
func Test_setSessionUser(t *testing.T) {
	router := setupTestRouter()

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

// Check that getSessionUser() works when session has been set correctly.
func Test_getSessionUser_Valid(t *testing.T) {
	router := setupTestRouter()

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

// Test that getSessionUser() works in the case that the session has not been set.
func Test_getSessionUser_NilSession(t *testing.T) {
	router := setupTestRouter()

	router.GET("/test", func(ctx *gin.Context) {
		username, password, exists := getSessionUser(ctx)

		if exists == true {
			t.Error("getSessionUser should return false when session doesn't exist")
		}

		if username != "" || password != "" {
			t.Error("getSessionUser should retun username and password as empty" +
				"strings when session doesn't exist")
		}
	})

	// mock request
	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
}
