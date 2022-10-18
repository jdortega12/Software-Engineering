package main

import "github.com/gin-gonic/gin"

const (
	PORT = ":8080"
)

// Sets up the routers api endpoints.
func setupEndpoints(*gin.Engine) {
}

func main() {
	// initialize gin router
	router := gin.Default()

	// run router
	router.Run(PORT)
}
