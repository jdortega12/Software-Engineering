package main

import (
	"jdortega12/Software-Engineering/GoServerApp/controller"
	"jdortega12/Software-Engineering/GoServerApp/model"

	"github.com/gin-gonic/gin"
)

const (
	PORT    = ":8080"
	DB_PATH = "database.db"
)

func main() {
	err := model.InitDB(DB_PATH)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	controller.SetupHandlers(router)

	router.Run(PORT)
}
