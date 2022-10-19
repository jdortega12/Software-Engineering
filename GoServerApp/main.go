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
func setupEndpoints(router *gin.Engine) {
	// 404 handler
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}

func main() {
	// initialize the db
	err := model.InitDB(DB_PATH)
	if err != nil {
		panic(err)
	}

	// initialize gin router
	router := gin.Default()
	setupEndpoints(router)

	// run router
	router.Run(PORT)
}
