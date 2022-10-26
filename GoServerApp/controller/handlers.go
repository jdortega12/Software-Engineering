package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"

	"github.com/gin-gonic/gin"
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
			v1.POST("/login", Login)
			v1.POST("/createTeamRequest", CreateTeamRequest)
			v1.POST("/updatePersonalInfo", UpdateUserPersonalInfoHandler)
			v1.POST("/createAccount", CreateAccountHandler)
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

// Logins in a user
func Login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	//validate current user
	_, _, err := model.ValidateUser(username, password)

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	//set session user
	setSessionUser(ctx, username, password)
	ctx.Status(http.StatusAccepted)

}

// Takes a POST request with Team request information
// Adds the request to the database
func CreateTeamRequest(ctx *gin.Context) {
	body := ctx.Request.Body
	value, err := io.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	var data map[string]string

	if err := json.Unmarshal(value, &data); err != nil {
		panic(err)
	}

	if model.InsertTeamNotification(data) == 0 {
		ctx.JSON(201, gin.H{"Created-notification": "true"})
	} else {
		ctx.JSON(201, gin.H{"Created-notification": "false"})
	}

}

// Receives UserPersonalInfo JSON from the client and updates
// it in the DB through the model. Validates that a user is
// logged in before doing anything.
func UpdateUserPersonalInfoHandler(ctx *gin.Context) {
	// get user session and check not null
	username, password, sessionExists := getSessionUser(ctx)
	if !sessionExists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// validate that user
	userID, _, err := model.ValidateUser(username, password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	var userPersInfo model.UserPersonalInfo

	err = ctx.BindJSON(&userPersInfo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	userPersInfo.UserPersonalInfoID = userID

	err = model.UpdateUserPersonalInfo(&userPersInfo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}

	ctx.Status(http.StatusAccepted)
}

func CreateAccountHandler(ctx *gin.Context) {
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if username == "" || email == "" || password == "" {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: password,
	}

	err := model.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	ctx.Status(http.StatusAccepted)
}
