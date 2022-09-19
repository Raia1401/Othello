package board

import (
	"errors"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//参考
//https://ruby-de-free.net/wp/how-to-make-othello-program-using-python-4/

type Board struct {
	StonesPos       [BOARD_SIZE + 2][BOARD_SIZE + 2]int  //石の置いてある場所
	PlaceabilityPos [BOARD_SIZE + 2][BOARD_SIZE + 2]bool //これから石が置ける場所
	PlaceabilityDir [BOARD_SIZE + 2][BOARD_SIZE + 2]int  //石がひっくり返る方向
}

func NewBoard() *Board {
	board := new(Board)
	//初期配置
	board.StonesPos[4][4] = BLACK_STONE
	board.StonesPos[4][5] = WHITE_STONE
	board.StonesPos[5][5] = BLACK_STONE
	board.StonesPos[5][4] = WHITE_STONE
	//壁の設定
	for i := 0; i < BOARD_SIZE+2; i += 1 {
		for j := 0; j < BOARD_SIZE+2; j += 1 {
			if i == 0 || i == BOARD_SIZE+1 || j == 0 || j == BOARD_SIZE+1 {
				board.StonesPos[i][j] = WALL
			}
		}
	}
	return board
}

func (b *Board) GetStonesPos() string {
	stonePlace := ""
	for i := 0; i < BOARD_SIZE+2; i += 1 {
		for j := 0; j < BOARD_SIZE+2; j += 1 {
			stonePlace += strconv.Itoa(b.StonesPos[i][j])
		}
	}
	return stonePlace
}

func (b *Board) SetStonesPos(StonesPos string) *Board {
	for i := 0; i < BOARD_SIZE+2; i += 1 {
		for j := 0; j < BOARD_SIZE+2; j += 1 {
			stone, err := strconv.Atoi(StonesPos[i*8+j : i*8+j+1])
			if err != nil {
				log.Fatal(err)
			}
			b.StonesPos[i][j] = stone
		}
	}
	return b
}

//石を置く処理
func (b *Board) PutDownStone(x int, y int, color int) (bool, error) {
	err := errors.New("you can't put down the stone here")
	if x < 1 || BOARD_SIZE < x {
		return false, err
	}
	if y < 1 || BOARD_SIZE < y {
		return false, err
	}
	if b.StonesPos[x][y] != 0 {
		return false, err
	}
	if !b.PlaceabilityPos[x][y] {
		return false, err
	}
	b.ReverseStones(x, y, color)
	b.InitPlaceability(color)
	return true, nil

}

//石を置くとどの方向にひっくり返せるか確認
func (b *Board) checkPlaceability(x int, y int, color int) int {
	var sum = 0
	if b.StonesPos[x][y] != 0 {
		return sum
	}
	dirList := DirList()
	for _, dir := range dirList {
		if b.checkPlaceabilityVector(x, y, dir.DirectionVecX, dir.DirectionVecY, color) {
			sum = sum | dir.DirectionNumb
		}
	}
	return sum
}

//特定の方向(xVec,yVec)に対して石をひっくり返せるか確認
func (b *Board) checkPlaceabilityVector(x int, y int, xVec int, yVec int, color int) bool {
	if b.StonesPos[x+xVec][y+yVec] == -color {
		x_tmp := x + xVec*2
		y_tmp := y + yVec*2
		for b.StonesPos[x_tmp][y_tmp] == -color {
			x_tmp += xVec
			y_tmp += yVec
		}
		if b.StonesPos[x_tmp][y_tmp] == color {
			return true
		} else {
			return false
		}
	}
	return false
}

//PlaceabilityPosとPlaceabilityDirを更新
func (b *Board) InitPlaceability(color int) {
	for i := 1; i < BOARD_SIZE+2; i++ {
		for j := 1; j < BOARD_SIZE+2; j++ {
			dir := b.checkPlaceability(i, j, color)
			b.PlaceabilityDir[i][j] = dir
			if dir != 0 {
				b.PlaceabilityPos[i][j] = true
			}
		}
	}

}

//石をひっくり返す処理
func (b *Board) ReverseStones(x int, y int, color int) {

	b.StonesPos[x][y] = color
	placeabilityDir := b.checkPlaceability(x, y, color)

	dirList := DirList()
	for _, dir := range dirList {
		canReverse := placeabilityDir & dir.DirectionNumb
		if canReverse != 0 {
			x_tmp := x + dir.DirectionVecX
			y_tmp := y + dir.DirectionVecY

			for b.StonesPos[x_tmp][y_tmp] == -color {
				b.StonesPos[x_tmp][y_tmp] = color
				x_tmp += dir.DirectionVecX
				y_tmp += dir.DirectionVecY
			}
		}

	}

}
