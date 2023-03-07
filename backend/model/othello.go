package model

import (
	"time"
)

type Board struct {
	BoardId    int64 `xorm:"pk autoincr unique 'board_id'"`
	Board      string
	CreatedAt  time.Time `xorm:"created 'created_at'"`
	IsMatchEnd bool      `xorm:"default false 'is_match_end'" `
	IsMyTurn   bool      `xorm:"default true 'is_my_turn'" `
	DeleteFlag bool      `xorm:"default false 'delete_flag'"`
}
