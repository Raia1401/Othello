package controller

import (
	"backend/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userId = int64(1)

//POST
func CreateGameMatch(c *gin.Context) {
	gameMatchService := service.GameMatchService{}
	err := gameMatchService.CreateGameMatch(userId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

//GET
func GetGameMatch(c *gin.Context) {
	gameMatchService := service.GameMatchService{}
	gameMatch := gameMatchService.GetGameMatch(userId)
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    gameMatch,
	})
}

//PATCH
func UpdateGameMatch(c *gin.Context) {
	gameMatchService := service.GameMatchService{}
	userId, err := strconv.Atoi(c.Param("userId"))
	x, err := strconv.Atoi(c.Param("x"))
	y, err := strconv.Atoi(c.Param("y"))
	color, err := strconv.Atoi(c.Param("color"))

	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
	}

	_, err = gameMatchService.UpdateGameMatch(int64(userId), x, y, color)

	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "you cant put down the stone",
		})
	}
	c.JSONP(http.StatusNoContent, gin.H{})
}

//to test
func Test(c *gin.Context) {
	testService := service.TestService{}
	User := testService.GetUser()
	fmt.Println(User)
}

// func CreateUser(c *gin.Context) {
//     user := model.User{}
//     err := c.Bind(&user)
//     if err != nil{
//         c.String(http.StatusBadRequest, "Bad request")
//         return
//     }
//    bookService :=service.BookService{}
//    err = bookService.SetBook(&book)
//    if err != nil{
//        c.String(http.StatusInternalServerError, "Server Error")
//        return
//    }
//    c.JSON(http.StatusCreated, gin.H{
//        "status": "ok",
//    })
// }
