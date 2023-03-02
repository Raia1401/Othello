package controller

import (
	"backend/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//POST method
func CreateGameMatch(c *gin.Context) {

	params, err := checkFormValue(c, "user_id")

	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}
	userId := params[0]

	gameMatchService := service.GameMatchService{}
	err = gameMatchService.CreateGameMatch(int64(userId))
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

	boardId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "board_id should be sended"})
		return
	}
	gameMatchService := service.GameMatchService{}
	gameMatch, err := gameMatchService.GetGameMatch(int64(boardId))
	if err != nil {
		println(err)
		c.JSONP(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    gameMatch,
	})
}

//PUT method
func PutDownStone(c *gin.Context) {

	boardId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "board_id should be sended"})
		return
	}

	params, err := checkFormValue(c, "x", "y", "color")
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "x,y,color should be sended"})
		return
	}

	x, y, color := params[0], params[1], params[2]

	gameMatchService := service.GameMatchService{}
	err = gameMatchService.PutDownStone(int64(boardId), x, y, color)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "you can't put down the stone"})
		return
	}

	c.JSONP(http.StatusNoContent, gin.H{})
}

//PUT method
func PutDownStoneByOpponent(c *gin.Context) {

	boardId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "board_id should be sended"})
		return
	}
	params, err := checkFormValue(c, "color")
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	color := params[0]

	gameMatchService := service.GameMatchService{}
	err = gameMatchService.PutDownStoneByOpponent(int64(boardId), color)

	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "opponent can't put down the stone"})
		return
	}
	c.JSONP(http.StatusNoContent, gin.H{})

}

//get form value in request data
func checkFormValue(c *gin.Context, formList ...string) ([]int, error) {
	var formValueList []int
	for _, formValueName := range formList {
		formValue, err := strconv.Atoi(c.Request.FormValue(formValueName))
		if err != nil {
			err := errors.New("queryString is not correct")
			return formValueList, err
		}
		formValueList = append(formValueList, formValue)
	}
	return formValueList, nil
}
