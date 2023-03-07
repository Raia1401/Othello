package agent

import (
	"backend/service/board"
	"math"

	_ "github.com/go-sql-driver/mysql"
)

type Agent struct {
}

func (a Agent) FindPosToPutDown(b *board.Board, color int) (int, int) {
	b.InitPlaceability(color)
	placeabilityPos := b.PlaceabilityPos
	bestX := 0
	bestY := 0
	bestEvaluetion := math.Inf(-1)

	for x, _ := range placeabilityPos {
		for y, _ := range placeabilityPos[0] {
			if placeabilityPos[x][y] {
				evaluetion := a.getEvaluation(b, x, y, color)
				if evaluetion > bestEvaluetion {
					bestEvaluetion = evaluetion
					bestX = x
					bestY = y
				}
			}
		}
	}
	return bestX, bestY
}

func (Agent) getEvaluation(b *board.Board, x int, y int, color int) float64 {
	var board_copy *board.Board = &board.Board{}
	*board_copy = *b
	board_copy.PutDownStone(x, y, color)

	evaluetionSum := 0.0
	positionEvaluetion := PositionEvalueation()

	for i := 1; i < board.BOARD_SIZE+1; i++ {
		for j := 1; j < board.BOARD_SIZE+1; j++ {
			if board_copy.StonesPos[i][j] == color {
				evaluetionSum += positionEvaluetion[i][j]
			}
		}
	}
	return evaluetionSum
}
