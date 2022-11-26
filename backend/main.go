package main

import (
	"backend/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Origin",
		},
		// ExposeHeaders: []string{
		// 	"Content-Length",
		// 	"Content-Type",
		// },
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	othelloRouter := router.Group("/othello")
	{
		v1 := othelloRouter.Group("/v1")

		//v1.POST("/users",controller.CreateUser)

		v1.GET("/gamematch/:board_id", controller.GetGameMatch)
		v1.POST("/gamematch", controller.CreateGameMatch)

		v1.PUT("/gamematch/myself/:board_id", controller.PutDownStone)
		v1.PUT("/gamematch/opponent/:board_id", controller.PutDownStoneByOpponent)

	}

	router.Run(":8080")
}
