package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
    "./controller"
)

func main() {
    router := gin.Default()

    othelloRouter := router.Group("/othello")
    {
        v1 := othelloRouter.Group("/v1")
        // v1.GET("/get-stone")
        v1.POST("/make-board")
        v1.PUT("/update-stone")
    }

    router.Run(":8080")
}