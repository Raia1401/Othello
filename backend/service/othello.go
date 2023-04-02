package service

import (
	"backend/model"
	"backend/service/agent"
	"backend/service/board"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type OthelloService struct{}

func (OthelloService) CreateBoard() (model.Board, error) {

	board := board.NewBoard()
	stonesPlace := board.GetStonesPos()
	boardModel := model.Board{Board: stonesPlace, IsMyTurn: true, IsMatchEnd: false, DeleteFlag: false}
	_, err := engine.Insert(&boardModel)
	if err != nil {
		return boardModel, err
	}
	return boardModel, nil
}

func (OthelloService) GetBoard(boardId int64) (model.Board, error) {
	boardModel := model.Board{}
	has, err := engine.Where("board_id=?", boardId).Get(&boardModel)
	if err != nil {
		return boardModel, err
	}
	if !has {
		err := errors.New("id is not correct")
		return boardModel, err
	}
	if IsMatchEnd(boardModel) {
		fmt.Println("マッチが終わった")
		boardModel.IsMatchEnd = true
	}
	return boardModel, nil
}

func (OthelloService) PutDownStone(boardId int64, putStoneX int, putStoneY int, myStoneColor int) (model.Board, error) {
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
	_, err = board.PutDownStone(putStoneX, putStoneY, myStoneColor)
	if err != nil {
		return boardModel, err
	}

	boardModel.Board = board.GetStonesPos()
	boardModel.IsMyTurn = false

	_, err = engine.Where("board_id=?", boardId).AllCols().Update(boardModel)

	if err != nil {
		return boardModel, err
	}

	return boardModel, nil

}

func (OthelloService) PutDownStoneByOpponent(boardId int64, myStoneColor int) (model.Board, error) {

	opponentStoneColor := (myStoneColor ^ 3)
	boardModel := model.Board{}
	_, err := engine.Where("board_id=?", boardId).Get(&boardModel)
	if err != nil {
		return boardModel, err
	}

	if boardModel.IsMyTurn {
		err := errors.New("its not opponents turn")
		return boardModel, err
	}

	board := board.NewBoard()
	board = board.SetStonesPos(boardModel.Board)

	opponent := agent.Agent{}
	x, y := opponent.FindPosToPutDown(board, opponentStoneColor)
	_, err = board.PutDownStone(x, y, opponentStoneColor)
	if err != nil {
		return boardModel, err
	}

	boardModel.Board = board.GetStonesPos()
	boardModel.IsMyTurn = true

	_, err = engine.Where("board_id=?", boardId).AllCols().Update(&boardModel)
	if err != nil {
		return boardModel, err
	}

	return boardModel, nil
}

func IsMatchEnd(boardModel model.Board) bool {

	b := board.NewBoard()
	b = b.SetStonesPos(boardModel.Board)

	myagent := agent.Agent{}
	x, y := myagent.FindPosToPutDown(b, board.BLACK_STONE)
	fmt.Println("my")
	fmt.Println(x, y)
	if x == 0 && y == 0 {
		opponent := agent.Agent{}
		x, y := opponent.FindPosToPutDown(b, board.WHITE_STONE)
		fmt.Println("opponent", x, y)
		if x == 0 && y == 0 {
			return true
		}
	}
	return false
}
