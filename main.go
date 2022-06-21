package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/api/v1/positions", getPositions)

	router.Run()
}
