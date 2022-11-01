package controller

import (
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// middlewares.go -> middlewares for api endpoints

// Middleware for authenticating a user. Responds with HTTP Status Unauthorized
// if session not set or user cannot be authenticated. Sets user in context with
// USER_KEY to be accessed by next handlers in chain.
func userAuthMiddleware(ctx *gin.Context) {
	username, password, sessionExists := getSessionUser(ctx)
	if !sessionExists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := model.AuthenticateUser(username, password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set(USER_KEY, user)

	ctx.Next()
}

// Middleware for authenticating that a user is a player. Aborts with HTTP
// Status Unauthorized if not.
func playerAuthMiddleware(ctx *gin.Context) {
	user := ctx.Keys[USER_KEY].(*model.User)

	if user.Role != model.PLAYER {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}

// Middleware for authenticating that a user is a manager. Aborts with
// HTTP Status Unauthorized if user is not manager.
func managerAuthMiddleware(ctx *gin.Context) {
	user := ctx.Keys[USER_KEY].(*model.User)

	if user.Role != model.MANAGER {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}

// Middleware for authenticating that a user is an admin. Aborts with
// HTTP Status Unauthorized if user is not admin.
func adminAuthMiddleware(ctx *gin.Context) {
	user := ctx.Keys[USER_KEY].(*model.User)

	if user.Role != model.ADMIN {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
