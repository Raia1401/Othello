package test

import (
	"backend/service/board"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard(t *testing.T) {
	board := board.NewBoard()

	//石を置ける場所(3,5)に置く
	ok, _ := board.PutDownStone(3, 5, 1)
	assert.Equal(t, true, ok)

	//石を既に置いてある場所(4,4)に置く
	ok, _ = board.PutDownStone(4, 4, 1)
	assert.Equal(t, false, ok)

	//石を置けない場所(8,8)に置く）
	ok, _ = board.PutDownStone(8, 8, 1)
	assert.Equal(t, false, ok)

}
