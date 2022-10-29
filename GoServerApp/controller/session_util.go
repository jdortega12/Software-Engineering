package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// session_util.go -> functionality for setting/getting/modifying session.

// Sets the current session with given username and password for
// validation of user for the rest of their session.
func setSessionUser(ctx *gin.Context, username string, password string) {
	session := sessions.Default(ctx)
	session.Set("username", username)
	session.Set("password", password)

	session.Save()
}

// Returns username and password of current session user as well
// as a bool for whether the session user key values exist or not.
func getSessionUser(ctx *gin.Context) (string, string, bool) {
	sessionExists := false

	session := sessions.Default(ctx)

	username := session.Get("username")
	password := session.Get("password")

	if username == nil || password == nil {
		return "", "", sessionExists
	}

	sessionExists = true
	return username.(string), password.(string), sessionExists
}

// Completely clears the current session.
func clearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Clear()
	session.Save()
}
