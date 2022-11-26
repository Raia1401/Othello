package model

import (
	"time"
)

type User struct {
	UserId   int64  `xorm:"pk autoincr int(64) unique" form:"user_id"`
	UserName string `xorm:"varchar(40)" form:"user_name"`
	Password string `xorm:"varchar(40)"`
}

type Board struct {
	BoardId     int64 `xorm:"pk autoincr unique 'board_id'"`
	BoardUserId int64 `xorm:"board_user_id"`
	Board       string
	CreatedAt   time.Time `xorm:"created 'created_at'"`
	IsMatchEnd  bool      `xorm:"default false 'is_match_end'" `
	IsMyTurn    bool      `xorm:"default true 'is_my_turn'" `
	DeleteFlag  bool      `xorm:"default false 'delete_flag'"`
}
