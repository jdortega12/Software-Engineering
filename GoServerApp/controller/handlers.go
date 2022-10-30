package controller

import (
	"encoding/json"
	"io"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	USER_KEY = "user"
)

// handlers.go -> endpoint setup and funcs bound to them

// Sets up the handling functions bound to the Engine's api endpoints.
func SetupHandlers(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// all handlers should go inside v1 (besides NoRoute)
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// endpoints requiring no authentication
			v1.POST("/createAccount", handleCreateAccount)
			v1.POST("/login", handleLogin)
			v1.POST("/logout", handleLogout)

			// endpoints requiring user authentication
			userAuth := v1.Group("")
			userAuth.Use(userAuthMiddleware)
			{
				userAuth.POST("/createTeamRequest", handleCreateTeamNotification)
				userAuth.POST("/updatePersonalInfo", handleUpdateUserPersonalInfo)
				userAuth.POST("/createPhoto", handleCreatePhoto)

				// endpoints requiring user to be a manager
				managerAuth := userAuth.Group("")
				managerAuth.Use(managerAuthMiddleware)
				{
					managerAuth.POST("/createTeam", handleCreateTeam)
				}

				// endpoints requiring user to be an admin
				adminAuth := userAuth.Group("")
				adminAuth.Use(adminAuthMiddleware)
				{

				}
			}
		}
	}
}

// Receives User JSON in request body and passes it to model to be created in DB.
// Responds HTTP Status Accepted on success. If JSON cannot be bound, responds
// HTTP Status Bad Request. If user cannot be created, responds HTTP Status Conflict.
func handleCreateAccount(ctx *gin.Context) {
	user := &model.User{}

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user.Role = model.PLAYER

	err = model.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// Logs in a user using username and password sent by client as JSON.
// Validates the credentials, sets current session with them, and
// responds with HTTP Status Accepted. If JSON cannot be bound, aborts
// with HTTP Status Bad Request. If user cannot be validated, aborts with
// HTTP Status Unauthorized.
func handleLogin(ctx *gin.Context) {
	user := &model.User{}

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err = model.ValidateUser(user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	setSessionUser(ctx, user.Username, user.Password)
	ctx.Status(http.StatusAccepted)
}

// Logs out current user by clearing the current session.
// Do not need to validate user or any permissions. Responds
// with HTTP reset content status code.
func handleLogout(ctx *gin.Context) {
	clearSession(ctx)

	ctx.Status(http.StatusResetContent)
}

// Receives UserPersonalInfo JSON from the client and updates it in the DB
// through the model. Responds with HTTP Status Accepted on success. Responds
// HTTP Status Unauthorized if session not set or user cannot be validated.
// Responds HTTP Status Bad Request if JSON cannot be bound. Responds HTTP
// Status Internal Server Error if DB cannot be updated.
func handleUpdateUserPersonalInfo(ctx *gin.Context) {
	user := ctx.Keys[USER_KEY].(*model.User)

	var userPersInfo model.UserPersonalInfo

	err := ctx.BindJSON(&userPersInfo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userPersInfo.ID = user.ID

	err = model.UpdateUserPersonalInfo(&userPersInfo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// Receives image string in request body and updates current user's photo
// in the DB. Returns HTTP Status Accepted on success. Returns
// HTTP Unprocessable Entity if session doesn't exist. Responds HTTP Status
// Bad Request if image cannot be processed.
func handleCreatePhoto(ctx *gin.Context) {
	user := ctx.Keys[USER_KEY].(*model.User)

	body := ctx.Request.Body

	value, err := io.ReadAll(body)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	var data map[string]string

	if err := json.Unmarshal(value, &data); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	photo, checkphoto := data["photo"]
	if !checkphoto {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	err = model.UpdateUserPhoto(photo, user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	ctx.Status(http.StatusAccepted)
}

// Receives Team JSON in response body and passes to model to be created in DB.
// Responds HTTP Status Accepted on success. Returns HTTP Status Unauthorized
// if session doesn't exist or user cannot be validated. Responds HTTP Bad Request
// if JSON cannot be bound and HTTP Status Internal Server Error if Team cannot
// be created in DB.
func handleCreateTeam(ctx *gin.Context) {
	team := &model.Team{}
	err := ctx.BindJSON(team)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = model.CreateTeam(team)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// Receives TeamNotification as JSON and passes it to model to be
// created in the DB. Returns HTTP Status Accepted on success. Returns
// HTTP Status Unauthorized if session not set or user credentials invalid.
// Returns HTTP Status Bad Request if JSON cannot be bound or notification
// cannot be created in DB.
func handleCreateTeamNotification(ctx *gin.Context) {
	user := ctx.Keys[USER_KEY].(*model.User)

	teamNotification := &model.TeamNotification{}

	err := ctx.BindJSON(teamNotification)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	teamNotification.SenderUsername = user.Username

	err = model.CreateTeamNotification(teamNotification)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusAccepted)
}
