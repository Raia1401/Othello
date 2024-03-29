package controller

import (
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//POST method
func CreateBoard(c *gin.Context) {

	othelloService := service.OthelloService{}
	othello, err := othelloService.CreateBoard()
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": othello})
}

//GET method
func GetBoard(c *gin.Context) {

	boardId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "invalid boardId"})
		return
	}
	othelloService := service.OthelloService{}
	othello, err := othelloService.GetBoard(int64(boardId))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSONP(http.StatusOK, gin.H{"message": "ok", "data": othello})
}

//PUT method
func UpdateBoard(c *gin.Context) {

	boardId, err := strconv.Atoi(c.Param("board_id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "invalid boardId"})
		return
	}

	isMyTurn, err := strconv.ParseBool(c.Request.FormValue("is_my_turn"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{"message": "invalid param: is_my_turn"})
		return
	}

	if isMyTurn {
		//　自分が石を置くターンの場合の処理

		x, err1 := strconv.Atoi(c.Request.FormValue("x"))
		y, err2 := strconv.Atoi(c.Request.FormValue("y"))
		myStoneColor, err3 := strconv.Atoi(c.Request.FormValue("my_stone_color"))
		if err1 != nil || err2 != nil || err3 != nil {
			c.JSONP(http.StatusBadRequest, gin.H{"message": "invalid param: x ,y, my_stone_color"})
			return
		}

		othelloService := service.OthelloService{}
		othello, err := othelloService.PutDownStone(int64(boardId), x, y, myStoneColor)
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{"message": err.Error()})
			return

		}

		c.JSONP(http.StatusOK, gin.H{"message": "ok", "data": othello})

	} else {
		// 相手が石を置くターンの場合の処理

		myStoneColor, err := strconv.Atoi(c.Request.FormValue("my_stone_color"))
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{"message": "invalid param: my_stone_color"})
			return
		}

		othelloService := service.OthelloService{}
		othello, err := othelloService.PutDownStoneByOpponent(int64(boardId), myStoneColor)
		if err != nil {
			c.JSONP(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSONP(http.StatusOK, gin.H{"message": "ok", "data": othello})

	}
}
