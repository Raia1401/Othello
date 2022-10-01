package agent

import (
	"backend/service/board"

	_ "github.com/go-sql-driver/mysql"
)

func PositionEvalueation() [board.BOARD_SIZE + 2][board.BOARD_SIZE + 2]float64 {
	positionEvalueation := [board.BOARD_SIZE + 2][board.BOARD_SIZE + 2]float64{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 30, -12, 0, -1, -1, 0, -12, 30, 0},
		{0, -12, -15, -3, -3, -3, -3, -15, -12, 0},
		{0, 0, -3, 0, -1, -1, 0, -3, 0, 0},
		{0, -1, -3, -1, -1, -1, -1, -3, -1, 0},
		{0, -1, -3, -1, -1, -1, -1, -3, -1, 0},
		{0, 0, -3, 0, -1, -1, 0, -3, 0, 0},
		{0, -12, -15, -3, -3, -3, -3, -15, -12, 0},
		{0, 30, -12, 0, -1, -1, 0, -12, 30, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	return positionEvalueation
}
