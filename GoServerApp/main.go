package main

import "github.com/gin-gonic/gin"

const (
	PORT = ":8080"
)

func main() {
	// initialize gin router
	router := gin.Default()

	// run router
	router.Run(PORT)
}
