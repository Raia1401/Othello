package service

import (
	"fmt"
	"os"
	"time"
    "../model"
	"github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init() {
    
    // Create engine
    engine, err = xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
    if err != nil {
        log.Fatal(err)
        return
    }
    
    // Create new table
    err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("init Database success")

}