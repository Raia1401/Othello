package service

import (
	"backend/service/board"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type TestService struct{}

// func (TestService) GetUserTest() bool {
// 	user := model.User{UserId: 1}
// 	fmt.Println(engine)
// 	// result, err := engine.Where("user_id = ?", 1).Get(&user)
// 	result, err := engine.Get(&user)
// 	// fmt.Println(err)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if !result {
// 		log.Fatal("Not Found")
// 	}
// 	fmt.Println("user:", result)
// 	return result
// }

func (TestService) BoardTest() {
	board := board.NewBoard()
	fmt.Println(board.StonesPos)
	//fmt.Println(board.PlaceabilityPos)
	//fmt.Println(board.PlaceabilityDir)

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
