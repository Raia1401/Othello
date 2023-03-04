package service

import (
	"backend/model"
	"backend/service/agent"
	"backend/service/board"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type GameMatchService struct{}

func (GameMatchService) CreateGameMatch() (model.Board, error) {

	board := board.NewBoard()
	stonesPlace := board.GetStonesPos()
	boardModel := model.Board{Board: stonesPlace, IsMyTurn: true, IsMatchEnd: false, DeleteFlag: false}
	_, err := engine.Insert(&boardModel)
	if err != nil {
		return boardModel, err
	}
	return boardModel, nil
}

func (GameMatchService) GetGameMatch(boardId int64) (model.Board, error) {
	boardModel := model.Board{}
	has, err := engine.Where("board_id=?", boardId).Get(&boardModel)
	if err != nil {
		return boardModel, err
	}
	if !has {
		err := errors.New("id is not correct")
		return boardModel, err
	}
	return boardModel, nil
}

func (GameMatchService) PutDownStone(boardId int64, putStoneX int, putStoneY int, color int) (model.Board, error) {
	boardModel := model.Board{}
	_, err := engine.Where("board_id=?", boardId).Get(&boardModel)
	if err != nil {
		return boardModel, err
	}
	if !boardModel.IsMyTurn {
		err := errors.New("its not your turn")
		return boardModel, err
	}

	board := board.NewBoard()
	board = board.SetStonesPos(boardModel.Board)
	_, err = board.PutDownStone(putStoneX, putStoneY, color)
	if err != nil {
		return boardModel, err
	}

	boardModel.Board = board.GetStonesPos()
	boardModel.IsMyTurn = false
	boardModel.DeleteFlag = true
	_, err = engine.Where("board_id=?", boardId).Update(boardModel)
	if err != nil {
		return boardModel, err
	}
	fmt.Println(boardModel)

	return boardModel, nil
}

func (GameMatchService) PutDownStoneByOpponent(boardId int64, color int) (model.Board, error) {
	boardModel := model.Board{}
	_, err := engine.Where("board_id=?", boardId).Get(&boardModel)
	if err != nil {
		return boardModel, err
	}

	fmt.Println(boardModel)

	if boardModel.IsMyTurn {
		err := errors.New("its not opponents turn")
		return boardModel, err
	}

	board := board.NewBoard()
	board = board.SetStonesPos(boardModel.Board)

	opponent := agent.Opponent{}
	x, y := opponent.FindPosToPutDown(board, (color ^ 3))
	_, err = board.PutDownStone(x, y, (color ^ 3))
	if err != nil {
		return boardModel, err
	}

	boardModel.Board = board.GetStonesPos()
	boardModel.IsMyTurn = true

	_, err = engine.Where("board_id=?", boardId).Update(&boardModel)
	if err != nil {
		return boardModel, err
	}

	return boardModel, nil
}
