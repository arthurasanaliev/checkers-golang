package game

import (
	"fmt"
)

type Board struct {
	board [8][8]string
}

func NewBoard() *Board {
	return &Board{
		[8][8]string{
			{"⬜", "🔵", "⬜", "🔵", "⬜", "🔵", "⬜", "🔵"},
			{"🔵", "⬜", "🔵", "⬜", "🔵", "⬜", "🔵", "⬜"},
			{"⬜", "🔵", "⬜", "🔵", "⬜", "🔵", "⬜", "🔵"},
			{"⬛", "⬜", "⬛", "⬜", "⬛", "⬜", "⬛", "⬜"},
			{"⬜", "⬛", "⬜", "⬛", "⬜", "⬛", "⬜", "⬛"},
			{"🔴", "⬜", "🔴", "⬜", "🔴", "⬜", "🔴", "⬜"},
			{"⬜", "🔴", "⬜", "🔴", "⬜", "🔴", "⬜", "🔴"},
			{"🔴", "⬜", "🔴", "⬜", "🔴", "⬜", "🔴", "⬜"},
		},
	}
}

func (b *Board) Display() {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print(b.board[i][j])
		}
		fmt.Println()
	}
}

// only for common pieces for now
func (b *Board) CheckMove(xFrom, yFrom, xTo, yTo int) bool {
	absInt := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	outOfBoundaries := func(x int) bool {
		return x < 0 || x >= 8
	}
	invalidMove := func(x1, y1, x2, y2 int) bool {
		return absInt(x1-x2) != 1 || absInt(y1-y2) != 1
	}
	emptyCell := func(x, y int) bool {
		return b.board[x][y] == "⬜" || b.board[x][y] == "⬛"
	}
	updateBoard := func(x1, y1, x2, y2 int) {
		b.board[x1][y1], b.board[x2][y2] = b.board[x2][y2], b.board[x1][y1]
	}

	if outOfBoundaries(xFrom) || outOfBoundaries(yFrom) || outOfBoundaries(xTo) || outOfBoundaries(yTo) {
		return false
	}
	if mustTake() {

	}
	if invalidMove(xFrom, yFrom, xTo, yTo) || emptyCell(xFrom, yFrom) || !emptyCell(xTo, yTo) {
		return false
	}

	updateBoard(xFrom, yFrom, xTo, yTo)

	return true
}

func (b *Board) CheckWin() bool {
	redPieces := 0
	bluePieces := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.board[i][j] == "🔴" {
				redPieces++
			}
			if b.board[i][j] == "🔵" {
				bluePieces++
			}
		}
	}
	return redPieces == 0 || bluePieces == 0
}
