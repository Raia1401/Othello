package service

import (
	"backend/model"
	"errors"
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

	// Create new Board table
	err = engine.Sync2(new(model.Board))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("init Database succeed")
}
