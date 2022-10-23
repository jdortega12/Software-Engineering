package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"io"
	"fmt"
)

// handlers.go -> funcs bound to the router's endpoints

// Sets up the routers api endpoints.
func SetupHandlers(router *gin.Engine) {
	// 404 handler
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// all handlers should go inside v1 (besides NoRoute)
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/logout", Logout)
			v1.POST("/createTeamRequest", CreateTeamRequest)
		}
	}
}

// Logs out current user by clearing the current session.
// Do not need to validate user or any permissions. Responds
// with HTTP reset content status code.
func Logout(ctx *gin.Context) {
	clearSession(ctx)

	ctx.JSON(http.StatusResetContent, gin.H{})
}

// Takes a POST request with Team request information 
// Adds the request the the database 
func CreateTeamRequest(ctx *gin.Context) {
	body := ctx.Request.Body 
	value, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(value))
}
