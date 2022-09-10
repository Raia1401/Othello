package model

type User struct {
    UserId int64 `xorm:"pk autoincr int(64)" form:"user_id"`
    UserName string `xorm:"varchar(40)" form:"user_name"`
    Password string `xorm:"varchar(40)"`
}

type GameMatch struct {
    MatchId int64 `xorm:"pk autoincr" form:"match_id"` 
    MatchUserId int64 `xorm:"match_user_id"`
    Board int64 
    CreatedAt time.Time `xorm:"created_at"`
    IsMatchEnd bool `xorm:"is_match_end"`
    DeleteFlag bool `xorm:"delete_flag"`
}