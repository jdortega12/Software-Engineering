package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"
	"strconv"

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
			v1.POST("/createPhoto", CreatePhoto)
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

	user := &model.User{}

	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//validate current user
	_, _, err = model.ValidateUser(user.Username, user.Password)

	if err != nil {
		ctx.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	//set session user
	setSessionUser(ctx, user.Username, user.Password)
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

	username, _, sessionExists := getSessionUser(ctx)

	if !sessionExists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	data["SenderID"] = username

	sender, err2 := model.GetUserId(data["SenderID"])

	if err2 != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	receiver, err3 := model.GetUserId(data["ReceiverID"])

	if err3 != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	data["SenderID"] = strconv.Itoa(int(sender))
	data["ReceiverID"] = strconv.Itoa(int(receiver))

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
		return
	}

	// validate that user
	userID, _, err := model.ValidateUser(username, password)
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

	userPersInfo.UserPersonalInfoID = userID

	err = model.UpdateUserPersonalInfo(&userPersInfo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// Take JSON of user info, transfers to user struct
// Creates user
func CreateAccountHandler(ctx *gin.Context) {

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
func CreatePhoto(ctx *gin.Context) {
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

func CreateTeamHandler(ctx *gin.Context) {
	username, password, sessExists := getSessionUser(ctx)
	if !sessExists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	_, role, err := model.ValidateUser(username, password) // _ is userID (not needed)
	if err != nil || role != model.MANAGER {
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
