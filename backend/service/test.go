package service

import (
	"backend/service/board"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TestService struct{}

func (TestService) BoardTest() {
	board := board.NewBoard()
	fmt.Println(board.StonesPos)

	//石を(3,4)に置く
	ok, err := board.PutDownStone(3, 4, 1)
	fmt.Println(ok)
	fmt.Println(err)
	fmt.Println(board.StonesPos)

	//石を(3,5)に置く
	ok, err = board.PutDownStone(3, 5, 1)
	fmt.Println(ok)
	fmt.Println(err)
	fmt.Println(board.StonesPos)

}
