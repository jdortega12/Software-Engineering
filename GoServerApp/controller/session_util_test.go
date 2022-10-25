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
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

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

func TestSetSession(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	test_store := cookie.NewStore([]byte("test"))
	router.Use(sessions.Sessions("test_session", test_store))

	router.GET("/test", func(ctx *gin.Context) {
		setSession(ctx, "test_username", "test_password")

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
