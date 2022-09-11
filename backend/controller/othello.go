package controller

import (
	"backend/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userId = int64(1)

func CreateGameMatch(c *gin.Context) {
	// board := service.NewBoard()
	// stonesPlace := board.GetStonesPlace()
	// gameMatch := model.GameMatch{MatchUserId: 1, Board: stonesPlace, IsMatchEnd: false, DeleteFlag: false}
	// err := c.Bind(&gameMatch)
	// if err != nil{
	//     c.String(http.StatusBadRequest, "Bad request")
	//     return
	// }

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

func GetGameMatch(c *gin.Context) {
	gameMatchService := service.GameMatchService{}
	gameMatch := gameMatchService.GetGameMatch(userId)
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    gameMatch,
	})
}

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
