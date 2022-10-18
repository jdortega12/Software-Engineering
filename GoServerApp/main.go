package main

import (
	"jdortega12/Software-Engineering/GoServerApp/model"

	"github.com/gin-gonic/gin"
)

const (
	PORT    = ":8080"
	DB_PATH = "database.db"
)

// Sets up the routers api endpoints.
func setupEndpoints(*gin.Engine) {
}

func main() {
	// initialize the db
	err := model.InitDB(DB_PATH)
	if err != nil {
		panic(err)
	}

	// initialize gin router
	router := gin.Default()

	// run router
	router.Run(PORT)
}
