package controller

import (
	"backend/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userId = int64(1)

//POST method
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

//GET method
func GetGameMatch(c *gin.Context) {
	gameMatchService := service.GameMatchService{}
	gameMatch, err := gameMatchService.GetGameMatch(userId)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    gameMatch,
	})

	// for i := 0; i < 100; {
	// 	fmt.Println(gameMatch.Board[i : i+10])
	// 	i += 10
	// }
}

//PUT method
func PutDownStone(c *gin.Context) {

	params, err := queryCheck(c, "userId", "x", "y", "color")
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	userId, x, y, color := params[0], params[1], params[2], params[3]
	gameMatchService := service.GameMatchService{}
	err = gameMatchService.PutDownStone(int64(userId), x, y, color)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "you can't put down the stone",
		})
		return
	}

	c.JSONP(http.StatusNoContent, gin.H{})
}

//PUT method
func PutDownStoneByOpponent(c *gin.Context) {
	params, err := queryCheck(c, "userId", "color")
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	userId, color := params[0], params[1]
	gameMatchService := service.GameMatchService{}
	err = gameMatchService.PutDownStoneByOpponent(int64(userId), color)

	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "opponent can't put down the stone",
		})
		return
	}
	c.JSONP(http.StatusNoContent, gin.H{})

}

//get query string in request data
func queryCheck(c *gin.Context, paramList ...string) ([]int, error) {
	var params []int
	for _, paramName := range paramList {
		param, err := strconv.Atoi(c.Query(paramName))
		if err != nil {
			err := errors.New("param is not correct")
			return params, err
		}
		params = append(params, param)
	}

	return params, nil
}
