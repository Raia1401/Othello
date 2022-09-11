package service

import (
	"fmt"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Board struct {
	StonesPlace [8][8]int
}

func NewBoard() *Board {
	board := new(Board)
	return board
}

func (b *Board) GetStonesPlace() string {
	stonePlace := ""
	for i := 0; i < 8; i += 1 {
		for j := 0; j < 8; j += 1 {
			stonePlace += strconv.Itoa(b.StonesPlace[i][j])
		}
	}
	// stone, err := strconv.Atoi(stonePlace)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return stonePlace
}

func (b *Board) SetStonesPlace(stonesPlace string) *Board {
	//stonePlace := strconv.Itoa(stonesPlace)
	// board := new(Board)
	for i := 0; i < 8; i += 1 {
		for j := 0; j < 8; j += 1 {
			stone, err := strconv.Atoi(stonesPlace[i*8+j : i*8+j+1])
			if err != nil {
				log.Fatal(err)
			}
			b.StonesPlace[i][j] = stone
		}
	}
	return b
}

func (Board) PutDownStone() {
	fmt.Println("")
}

func (Board) ReverseStones() {
	fmt.Println("")
}
