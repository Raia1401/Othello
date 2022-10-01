package service

import (
	"backend/model"
	"backend/service/agent"
	"backend/service/board"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type GameMatchService struct {
}

func (GameMatchService) CreateGameMatch(userId int64) error {

	board := board.NewBoard()
	stonesPlace := board.GetStonesPos()
	gameMatch := model.GameMatch{MatchUserId: userId, Board: stonesPlace, IsMatchEnd: false, DeleteFlag: false}
	_, err := engine.Insert(&gameMatch)
	if err != nil {
		return err
	}
	return nil
}

func (GameMatchService) GetGameMatch(userId int64) (model.GameMatch, error) {
	gameMatch := model.GameMatch{}
	_, err := engine.Where("match_user_id=?", userId).Get(&gameMatch)
	if err != nil {
		// panic(err)
		return gameMatch, err
	}
	return gameMatch, nil
}

func (GameMatchService) PutDownStone(userId int64, putStoneX int, putStoneY int, color int) error {
	gameMatch := model.GameMatch{}
	_, err := engine.Where("match_user_id=?", userId).Get(&gameMatch)
	if err != nil {
		return err
	}
	board := board.NewBoard()
	board = board.SetStonesPos(gameMatch.Board)
	_, err = board.PutDownStone(putStoneX, putStoneY, color)
	if err != nil {
		return err
	}

	gameMatch = model.GameMatch{Board: board.GetStonesPos()}
	_, err = engine.Where("match_user_id=?", userId).Update(&gameMatch)
	if err != nil {
		return err
	}

	return nil
}

func (GameMatchService) PutDownStoneByOpponent(userId int64, color int) error {
	gameMatch := model.GameMatch{}
	_, err := engine.Where("match_user_id=?", userId).Get(&gameMatch)
	if err != nil {
		return err
	}
	board := board.NewBoard()
	board = board.SetStonesPos(gameMatch.Board)

	opponent := agent.Opponent{}
	x, y := opponent.FindPosToPutDown(board, (color ^ 3))
	fmt.Println(x, y)
	_, err = board.PutDownStone(x, y, (color ^ 3))

	if err != nil {
		return err
	}

	gameMatch = model.GameMatch{Board: board.GetStonesPos()}
	_, err = engine.Where("match_user_id=?", userId).Update(&gameMatch)
	if err != nil {
		return err
	}

	return nil
}
