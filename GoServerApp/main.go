package main

import (
	"jdortega12/Software-Engineering/GoServerApp/controller"
	"jdortega12/Software-Engineering/GoServerApp/model"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	PORT    = ":8080"
	DB_PATH = "database.db"
)

// Initializes DB and router, runs server.
func main() {
	err := model.InitDB(DB_PATH)
	if err != nil {
		panic(err)
	}

	team1 := model.Team{
		ID:    1,
		Name:  "Desk",
		Wins:  8,
		Loses: 0,
	}
	model.CreateTeam(&team1)

	team2 := model.Team{
		ID:    2,
		Name:  "Chair",
		Wins:  6,
		Loses: 2,
	}
	model.CreateTeam(&team2)

	team3 := model.Team{
		ID:    3,
		Name:  "Hair",
		Wins:  5,
		Loses: 3,
	}
	model.CreateTeam(&team3)

	team4 := model.Team{
		ID:    4,
		Name:  "Stair",
		Wins:  4,
		Loses: 4,
	}
	model.CreateTeam(&team4)

	team5 := model.Team{
		ID:    5,
		Name:  "Lair",
		Wins:  2,
		Loses: 6,
	}
	model.CreateTeam(&team5)

	team6 := model.Team{
		ID:    6,
		Name:  "Flair",
		Wins:  1,
		Loses: 7,
	}
	model.CreateTeam(&team6)

	team7 := model.Team{
		ID:    7,
		Name:  "Bear",
		Wins:  1,
		Loses: 7,
	}
	model.CreateTeam(&team7)

	team8 := model.Team{
		ID:    8,
		Name:  "Fair",
		Wins:  0,
		Loses: 8,
	}
	model.CreateTeam(&team8)

	user := &model.User{
		ID:       1,
		TeamID:   1,
		Username: "jaymin",
		Password: "123",
	}
	model.CreateUser(user)

	personalInfo := &model.UserPersonalInfo{
		ID:        1,
		Firstname: "Jaymin",
		Lastname:  "Ortega",
	}
	model.UpdateUserPersonalInfo(personalInfo)

	manager1 := &model.User{
		ID:       2,
		TeamID:   1,
		Username: "manager1",
		Password: "123",
		Role:     model.MANAGER,
	}

	personalInfo1 := &model.UserPersonalInfo{
		ID:        2,
		Firstname: "First",
		Lastname:  "Manager",
	}
	model.UpdateUserPersonalInfo(personalInfo)

	manager2 := &model.User{
		ID:       3,
		TeamID:   2,
		Username: "manager2",
		Password: "123",
		Role:     model.MANAGER,
	}
	personalInfo2 := &model.UserPersonalInfo{
		ID:        3,
		Firstname: "Second",
		Lastname:  "Manager",
	}

	model.CreateUser(manager1)
	model.CreateUser(manager2)
	model.UpdateUserPersonalInfo(personalInfo1)
	model.UpdateUserPersonalInfo(personalInfo2)

	admin := &model.User{
		Username: "jortega",
		Password: "123",
		Role:     model.ADMIN,
	}
	model.CreateUser(admin)

	match := model.Match{
		ID:            1,
		MatchType:     model.REGULAR,
		Location:      "Knott Hall",
		StartTime:     time.Now(),
		InProgress:    true,
		Quarter:       uint(1),
		QuarterTime:   time.Date(0, 0, 0, 0, 15, 0, 0, time.FixedZone("UTC-7", 0)),
		HomeTeamID:    1,
		AwayTeamID:    2,
		HomeTeamScore: 0,
		AwayTeamScore: 0,
		Likes:         0,
		Dislikes:      0,
	}
	model.DBConn.Create(&match)

	// session store must be set up right after router is initialized
	router := gin.Default()
	store := cookie.NewStore([]byte("placeholder"))
	router.Use(sessions.Sessions("session", store))

	controller.SetupHandlers(router)

	router.Run(PORT)
}
