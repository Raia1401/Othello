package service

import (
	"backend/model"
	"backend/service/board"

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

func (GameMatchService) GetGameMatch(userId int64) model.GameMatch {
	gameMatch := model.GameMatch{}
	_, err := engine.Where("match_user_id=?", userId).Get(&gameMatch)
	if err != nil {
		panic(err)
	}
	return gameMatch
}

func (GameMatchService) UpdateGameMatch(userId int64, putStoneX int, putStoneY int, color int) (model.GameMatch, error) {
	gameMatch := model.GameMatch{}
	_, err := engine.Where("match_user_id=?", userId).Get(&gameMatch)
	if err != nil {
		panic(err)
	}
	board := board.NewBoard()
	board = board.SetStonesPos(gameMatch.Board)
	_, err = board.PutDownStone(putStoneX, putStoneY, color)
	if err != nil {
		return gameMatch, err
	}

	return gameMatch, nil
}
