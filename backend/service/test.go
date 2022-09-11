package service

import (
	"fmt"
    "log"
    "backend/model"
    _ "github.com/go-sql-driver/mysql"
)

type TestService struct {}

func (TestService) GetUser() bool {
	user := model.User{UserId:1}
    fmt.Println(engine)
	// result, err := engine.Where("user_id = ?", 1).Get(&user)
    result,err := engine.Get(&user)
    // fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	if !result {
		log.Fatal("Not Found")
	}
	fmt.Println("user:", result)
    return result
}
