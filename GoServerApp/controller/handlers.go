package controller

import (
	"encoding/json"
	"io"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"net/http"
	"sort"
	"strconv"
	"time"
	"fmt"

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
			v1.GET("/get-user/:username", handleGetUser)
			v1.POST("/createAccount", handleCreateAccount)
			v1.POST("/login", handleLogin)
			v1.POST("/logout", handleLogout)
			v1.GET("/getTeam/:id", handleGetTeam)
			v1.GET("/getComments/:id", handleGetComments)
			v1.GET("/getTeams", handleGetTeams)
			v1.GET("/getTeamPlayers/:id", handleGetTeamPlayers)
			v1.GET("/getPlayoffs", handleGetPlayoffs)
			v1.GET("/getMatch/:id", handleGetMatch)

			// endpoints requiring user authentication
			userAuth := v1.Group("")
			userAuth.Use(userAuthMiddleware)
			{
				userAuth.POST("/createTeamRequest", handleCreateTeamNotification)
				userAuth.POST("/updatePersonalInfo", handleUpdateUserPersonalInfo)
				userAuth.POST("/createPhoto", handleCreatePhoto)
				userAuth.POST("/createComment", handleCreateComment)
				userAuth.POST("/postLikes/:id", handlePostLikes)
				userAuth.POST("/postDislikes/:id", handlePostDislikes)
				

				playerAuth := userAuth.Group("")
				playerAuth.Use(playerAuthMiddleware)
				{
					playerAuth.POST("/create-promotion-to-manager-request", handleCreatePromotionToManagerRequest)
				}

				// endpoints requiring user to be a manager
				managerAuth := userAuth.Group("")
				managerAuth.Use(managerAuthMiddleware)
				{
					managerAuth.POST("/createTeam", handleCreateTeam)
					managerAuth.POST("/acceptPlayer", handleAcceptPlayer)
				}

				// endpoints requiring user to be an admin
				adminAuth := userAuth.Group("")
				adminAuth.Use(adminAuthMiddleware)
				{
					adminAuth.GET("/promotion-to-manager-requests", handleGetPromotionToManagerRequests)
					adminAuth.GET("/getUserTeamData", handleGetUserTeamData)
					adminAuth.POST("/start-match", handleStartMatch)
					adminAuth.POST("/finish-match", handleFinishMatch)
					adminAuth.POST("/changeRoster", handleChangeRoster)
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

	_, err = model.AuthenticateUser(user.Username, user.Password)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	setSessionUser(ctx, user.Username, user.Password)
	ctx.Status(http.StatusAccepted)
}

// Logs out current user by clearing the current session. Do not need to validate
// user or any permissions. Responds with HTTP reset content status code.
func handleLogout(ctx *gin.Context) {
	clearSession(ctx)

	ctx.Status(http.StatusResetContent)
}

// Receives UserPersonalInfo JSON from the client and updates it in the DB through
// the model. Responds with HTTP Status Accepted on success. Responds HTTP Status
// Bad Request if JSON cannot be bound. Responds HTTP Status Internal Server Error
// if DB cannot be updated.
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
// in the DB. Returns HTTP Status Accepted on success. Responds HTTP Status
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
// Responds HTTP Status Accepted on success. Responds HTTP Status Conflict if
// manager already has a team, HTTP Bad Request if JSON cannot be bound, and
// HTTP Status Internal Server Error if Team cannot be created in DB or manager
// TeamID cannot be updated.
func handleCreateTeam(ctx *gin.Context) {
	user := ctx.Keys[USER_KEY].(*model.User)

	if user.TeamID != 0 {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

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

	err = model.UpdateUserTeam(user, team.ID)
	if err != nil {
		model.DeleteTeam(team)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// Receives TeamNotification as JSON and passes it to model to be
// created in the DB. Returns HTTP Status Accepted on success. Returns
// HTTP Status Bad Request if JSON cannot be bound or notification
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

// Receives a username in URL params and responds with JSON of the public info
// about that user. Returns HTTP Status Found on success. Aborts with Not Found
// if user doesn't exist and Internal Server Error if the user's personal
// info or team cannot be found.
func handleGetUser(ctx *gin.Context) {
	usernameToGet := ctx.Param("username")

	userFound, err := model.GetUserByUsername(usernameToGet)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	userFound.Password = ""

	userInfoFound, err := model.GetUserPersonalInfoByID(userFound.ID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	teamName := ""
	if userFound.TeamID != 0 {
		userTeamFound, err := model.GetTeamByID(userFound.TeamID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		teamName = userTeamFound.Name
	}

	userData := &userDataWrapper{
		User:         *userFound,
		PersonalInfo: *userInfoFound,
		TeamName:     teamName,
	}

	ctx.JSON(http.StatusFound, userData)
}

// Receives a PromotionToManagerRequest JSON and passes to model to be created.
// Responds HTTP Status Accepted on success. Returns HTTP Status Conflict if user
// already has an open request, HTTP Status Bad Request if JSON cannot be bound, and
// HTTP Internal Server Error if request cannot be created in DB.
func handleCreatePromotionToManagerRequest(ctx *gin.Context) {
	user := ctx.Keys[USER_KEY].(*model.User)

	// make sure user has no open requests
	_, err := model.GetPromoToManReqBySendUsername(user.Username)
	if err == nil {
		ctx.AbortWithStatus(http.StatusConflict)
		return
	}

	request := &model.PromotionToManagerRequest{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	request.SenderID = user.ID
	request.SenderUsername = user.Username

	err = model.CreatePromotionToManagerRequest(request)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// Responds with JSON of all PromotionToManagerRequest's in the DB for an admin's
// notifications. Responds HTTP Status Found on success and Internal Sever Error
// if for some reason there is an error retrieved the requests.
func handleGetPromotionToManagerRequests(ctx *gin.Context) {
	requests, err := model.GetAllPromotionToManagerRequests()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusFound, requests)
}

// Recieves a team ID and returns a JSON containing the id, name, and location of that team
func handleGetTeam(ctx *gin.Context) {
	teamID := ctx.Param("id")
	teamIDInt, err := strconv.Atoi(teamID)

	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	team, err2 := model.GetTeamByID(uint(teamIDInt))

	if err2 != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	manager, err3 := model.GetManagerByTeamID(uint(teamIDInt))

	if err3 != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"id":           team.ID,
		"name":         team.Name,
		"location":     team.TeamLocation,
		"manager_id":   manager.ID,
		"manager_name": manager.Username,
	})

}

// Return the teams from the database
func handleGetTeams(ctx *gin.Context) {
	teams, err := model.GetTeams()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusAccepted, teams)
}

// Receives player's username and manager's username
// Updates the player's teamID to match the manager's
func handleAcceptPlayer(ctx *gin.Context) {
	data := &model.AcceptData{}
	println(data.ManagerName)
	println(data.PlayerName)

	err := ctx.BindJSON(&data)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	playerName := data.PlayerName
	managerName := data.ManagerName

	player, err := model.GetUserByUsername(playerName)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	manager, err := model.GetUserByUsername(managerName)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	//change player team ID to match manager team ID
	err = model.UpdateUserTeam(player, manager.TeamID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	}

	ctx.Status(http.StatusAccepted)
}

// Takes a team ID in the body and returns a list of players on that team
func handleGetTeamPlayers(ctx *gin.Context) {
	teamID := ctx.Param("id")
	teamIDInt, err := strconv.Atoi(teamID)

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	players, err2 := model.GetPlayersByTeamID(uint(teamIDInt))

	if err2 != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusAccepted, players)
}

// Return the current top 8 teams (teams in the playoffs) in order
func handleGetPlayoffs(ctx *gin.Context) {
	matches, err := model.GetMatchesThisSeason()

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	teams := make(map[uint]int)

	for _, match := range matches {
		if match.HomeTeamScore > match.AwayTeamScore {
			teams[match.HomeTeamID] += 1
			teams[match.AwayTeamID] -= 1
		}
		if match.AwayTeamScore > match.HomeTeamScore {
			teams[match.AwayTeamID] += 1
			teams[match.HomeTeamID] -= 1
		}
	}

	keys := make([]uint, 0)

	for k := range teams {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return teams[keys[i]] > teams[keys[j]]
	})

	team_names := make([]string, 0)

	for k := range keys {
		team, err := model.GetTeamByID(uint(keys[k]))
		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		team_names = append(team_names, team.Name)
	}

	ctx.JSON(http.StatusAccepted, team_names[0:8])
}

// Receives match wrapper struct in JSON and creates match in the DB, setting InProgress = true. Responds
// HTTP Status Created on success. Responds HTTP Status Bad Request if JSON cannot be bound or
// if one or both of the teams don't exist, and HTTP Status Internal Server Error if the match
// cannot be created in the DB.
func handleStartMatch(ctx *gin.Context) {
	wrapper := &startMatchWrapper{}

	err := ctx.BindJSON(wrapper)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	names := []string{wrapper.HomeTeamName, wrapper.AwayTeamName}
	ids := make([]uint, 2)

	for i := range names {
		team, err := model.GetTeamByName(names[i])
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ids[i] = team.ID
	}

	match := &model.Match{
		HomeTeamID: ids[0],
		AwayTeamID: ids[1],
		Location:   wrapper.Location,

		InProgress: true,
		StartTime:  time.Now(),
	}

	err = model.CreateMatch(match)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusCreated)
}

// Receives match ID in JSON and marks it as finished in the database.
func handleFinishMatch(ctx *gin.Context) {
	match := &model.Match{}

	err := ctx.BindJSON(match)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = model.FinishMatch(match.ID)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusAccepted)
}

// called when the frontend tries to retrieve a match by id
func handleGetMatch(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	match, err := model.GetMatchById(uint(idInt))

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusAccepted, match)
}

func handleGetComments(ctx *gin.Context){
	id := ctx.Param("id")
	int_id, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return 
	}

	comments, err := model.GetCommentsById(uint(int_id))


	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return 
	}

	ctx.JSON(http.StatusAccepted, comments)
}

func handleCreateComment(ctx *gin.Context){
	comment := model.Comment{}
	err := ctx.BindJSON(&comment)
	username, _, _ := getSessionUser(ctx)
	fmt.Println(username)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	
	comment.Username = username
	err = model.CreateComment(&comment)

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return 
	}
}
// Handles getting user/team data from db
func handleGetUserTeamData(ctx *gin.Context) {
	//get all the users
	users, err := model.GetUsers()

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	var userTeamData []model.UserTeamData

	//create array of userTeamData from users
	for i := 0; i < len(users); i++ {
		if users[i].Role == model.PLAYER || users[i].Role == model.MANAGER {
			userTeamData = append(userTeamData, model.GatherUserTeamData(&users[i]))
		}

	}

	ctx.JSON(http.StatusAccepted, userTeamData)
}

func handleChangeRoster(ctx *gin.Context) {
	info := &model.UserTeamReturnData{}

	err := ctx.BindJSON(&info)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Get team
	team, err := model.GetTeamByName(info.Teamname)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	//Get user
	user, err1 := model.GetUserbyID(info.UserId)
	if err1 != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//Update user's team
	err2 := model.UpdateUserTeam(user, team.ID)
	if err2 != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.Status(http.StatusAccepted)
}

func handlePostLikes(ctx *gin.Context){
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	_, err = model.GetMatchById(uint(idInt))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	model.AddLike(uint(idInt))

	ctx.Status(http.StatusAccepted)
}

func handlePostDislikes(ctx *gin.Context){
	fmt.Println("got here")
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	_, err = model.GetMatchById(uint(idInt))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	model.AddDislike(uint(idInt))
	ctx.Status(http.StatusAccepted)
}
