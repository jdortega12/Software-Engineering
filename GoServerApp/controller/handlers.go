package controller

import (
	"encoding/json"
	"io"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handlers.go -> endpoint setup and funcs bound to them

// Sets up the handling functions bound to the Engine's api endpoints.
func SetupHandlers(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// all handlers should go inside v1 (besides NoRoute)
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/logout", handleLogout)
			v1.POST("/login", handleLogin)
			v1.POST("/createTeamRequest", handleCreateTeamNotification)
			v1.POST("/updatePersonalInfo", handleUpdateUserPersonalInfo)
			v1.POST("/createAccount", handleCreateAccount)
			v1.POST("/createPhoto", handleCreatePhoto)
		}
	}
}

// Logs out current user by clearing the current session.
// Do not need to validate user or any permissions. Responds
// with HTTP reset content status code.
func handleLogout(ctx *gin.Context) {
	clearSession(ctx)

	ctx.JSON(http.StatusResetContent, gin.H{})
}

// Logs in a user using username and password sent
// by client as JSON. Validates the credentials, sets
// current session with them, and responds with HTTP
// Status Accepted. If JSON cannot be bound, aborts
// with HTTP Status Bad Request. If user cannot be
// validated, aborts with HTTP Status Unauthorized.
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

// Receives TeamNotification as JSON from client and passes it to
// model to be created in the DB. Returns HTTP Status Accepted on success.
// Returns HTTP Status Unauthorized if session not set or user credentials
// invalid. Returns HTTP Status Bad Request if JSON cannot be bound, SenderUsername
// is not the same as the logged in user, or notification cannot be created in DB.
func handleCreateTeamNotification(ctx *gin.Context) {
	username, password, sessionExists := getSessionUser(ctx)
	if !sessionExists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	_, err := model.ValidateUser(username, password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	teamNotification := &model.TeamNotification{}

	err = ctx.BindJSON(teamNotification)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if username != teamNotification.SenderUsername {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = model.CreateTeamNotification(teamNotification)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// Receives UserPersonalInfo JSON from the client and updates
// it in the DB through the model. Validates that a user is
// logged in before doing anything.
func handleUpdateUserPersonalInfo(ctx *gin.Context) {
	// get user session and check not null
	username, password, sessionExists := getSessionUser(ctx)
	if !sessionExists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// validate that user
	user, err := model.ValidateUser(username, password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var userPersInfo model.UserPersonalInfo

	err = ctx.BindJSON(&userPersInfo)
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

// Take JSON of user info, transfers to user struct
// Creates user
func handleCreateAccount(ctx *gin.Context) {
	user := &model.User{}
	//bind JSON with User struct
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Create user with struct
	err = model.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}
	ctx.Status(http.StatusAccepted)
}

// Take JSON with base64 of image, image filetype, and user id as parameters, insert into DB
func handleCreatePhoto(ctx *gin.Context) {
	username, password, sessionExists := getSessionUser(ctx)
	if !sessionExists {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
	}
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

	err = model.UpdateUserPhoto(photo, username, password)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

	ctx.Status(http.StatusAccepted)
}

func handleCreateTeam(ctx *gin.Context) {
	username, password, sessExists := getSessionUser(ctx)
	if !sessExists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	user, err := model.ValidateUser(username, password) // _ is userID (not needed)
	if err != nil || user.Role != model.MANAGER {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	team := &model.Team{}
	err = ctx.BindJSON(team)
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
