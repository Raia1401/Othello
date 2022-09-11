package main

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	othelloRouter := router.Group("/othello")
	{
		v1 := othelloRouter.Group("/v1")

		//v1.POST("/users",controller.CreateUser)
		v1.GET("/gamematch", controller.GetGameMatch)
		v1.POST("/gamematch", controller.CreateGameMatch)
		v1.PUT("/gamematch")

		// To test
		// v1.GET("/test",controller.Test)
	}

	router.Run(":8080")
}
