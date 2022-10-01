package board

//盤の大きさや石の設定
const BOARD_SIZE = 8
const BLACK_STONE = 1 //2 01
const WHITE_STONE = 2 //4 10
const WALL = 3

//方向の設定
func DirList() map[string]Direction {
	direction := map[string]Direction{
		"LEFT":        Direction{1, -1, 0},
		"UPPER_LEFT":  Direction{2, -1, -1},
		"UPPER":       Direction{4, 0, -1},
		"UPPER_RIGHT": Direction{8, 1, -1},
		"RIGHT":       Direction{16, 1, 0},
		"LOWER_RIGHT": Direction{32, 1, 1},
		"LOWER":       Direction{64, 0, 1},
		"LOWER_LEFT":  Direction{128, -1, 1}}
	return direction
}

type Direction struct {
	DirectionNumb int
	DirectionVecX int
	DirectionVecY int
}
