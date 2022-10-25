package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// session_util.go -> functionality for setting/getting/modifying session,
// also interactions with model to validate user permissions. Funcs should
// not be publicly visible to other packages (lowercase).

// Sets the current session with given username and password for
// validation of user for the rest of their session.
func setSession(ctx *gin.Context, username string, password string) {
	session := sessions.Default(ctx)
	session.Set("username", username)
	session.Set("password", password)

	session.Save()
}

// Completely clears the current session.
func clearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Clear()
	session.Save()
}
