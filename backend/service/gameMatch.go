package service

import (
	"backend/model"

	_ "github.com/go-sql-driver/mysql"
)

type GameMatchService struct {
}

func (GameMatchService) CreateGameMatch(userId int64) error {

	board := NewBoard()
	stonesPlace := board.GetStonesPlace()
	gameMatch := model.GameMatch{MatchUserId: userId, Board: stonesPlace, IsMatchEnd: false, DeleteFlag: false}

	_, err := engine.Insert(&gameMatch)
	if err != nil {
		return err
	}
	return nil
}

func (GameMatchService) GetGameMatch(userId int64) model.GameMatch {
	gameMatch := model.GameMatch{}
	_, err := engine.Where("UserId=?", userId).Get(&gameMatch)
	if err != nil {
		panic(err)
	}
	return gameMatch
}
