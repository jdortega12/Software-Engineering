package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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
