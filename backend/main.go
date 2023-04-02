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
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	othelloRouter := router.Group("/othello")
	{
		v1 := othelloRouter.Group("/v1")

		v1.POST("/", controller.CreateBoard)
		v1.GET("/:board_id", controller.GetBoard)
		v1.PUT("/:board_id", controller.UpdateBoard)
		// v1.PUT("/gamematch/:board_id", controller.PutDownStoneByOpponent)

	}

	router.Run(":8080")
}
