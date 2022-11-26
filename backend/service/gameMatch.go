package service

import (
	"backend/model"
	"backend/service/agent"
	"backend/service/board"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type GameMatchService struct{}

func (GameMatchService) CreateGameMatch(userId int64) error {

	board := board.NewBoard()
	stonesPlace := board.GetStonesPos()
	boardModel := model.Board{BoardUserId: userId, Board: stonesPlace, IsMyTurn: true, IsMatchEnd: false, DeleteFlag: false}
	_, err := engine.Insert(&boardModel)
	if err != nil {
		return err
	}
	return nil
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

func (GameMatchService) PutDownStone(boardId int64, putStoneX int, putStoneY int, color int) error {
	boardModel := model.Board{}
	_, err := engine.Where("board_id=?", boardId).Get(&boardModel)
	if err != nil {
		return err
	}
	board := board.NewBoard()
	board = board.SetStonesPos(boardModel.Board)
	_, err = board.PutDownStone(putStoneX, putStoneY, color)
	if err != nil {
		return err
	}

	boardModel = model.Board{Board: board.GetStonesPos()}
	_, err = engine.Where("board_id=?", boardId).Update(&boardModel)
	if err != nil {
		return err
	}

	return nil
}

func (GameMatchService) PutDownStoneByOpponent(boardId int64, color int) error {
	boardModel := model.Board{}
	_, err := engine.Where("board_id=?", boardId).Get(&boardModel)
	if err != nil {
		return err
	}
	board := board.NewBoard()
	board = board.SetStonesPos(boardModel.Board)

	opponent := agent.Opponent{}
	x, y := opponent.FindPosToPutDown(board, (color ^ 3))
	_, err = board.PutDownStone(x, y, (color ^ 3))

	if err != nil {
		return err
	}

	boardModel = model.Board{Board: board.GetStonesPos()}
	_, err = engine.Where("board_id=?", boardId).Update(&boardModel)
	if err != nil {
		return err
	}

	return nil
}
