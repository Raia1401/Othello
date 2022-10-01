package service

import (
	"backend/model"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {

	// Create engine
	err := errors.New("")
	engine, err = xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/othello?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create new User table
	err = engine.Sync2(new(model.User))
	if err != nil {
		log.Fatal(err)
	}

	// Create new GameMatch table
	err = engine.Sync2(new(model.GameMatch))
	if err != nil {
		log.Fatal(err)
	}

	user := model.User{UserId: 1, UserName: "taro", Password: "password"}
	_, err = engine.Table("user").Insert(user)

	fmt.Println("init Database succeed")
}
