package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// session_util.go -> functionality for setting/getting/modifying session,
// also interactions with model to validate user permissions. Funcs should
// not be publicly visible to other packages (lowercase).

// Completely clears the current session.
func clearSession(ctx *gin.Context) {
	session := sessions.Default(ctx)

	session.Clear()
	session.Save()
}
