package main

import (
	"jdortega12/Software-Engineering/GoServerApp/controller"
	"jdortega12/Software-Engineering/GoServerApp/model"

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

	// session store must be set up right after router is initialized
	router := gin.Default()
	store := cookie.NewStore([]byte("placeholder"))
	router.Use(sessions.Sessions("session", store))

	controller.SetupHandlers(router)

	router.Run(PORT)
}
