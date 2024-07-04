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
			{"🔵", "⬜", "🔵", "⬜", "⬛", "⬜", "🔵", "⬜"},
			{"⬜", "⬛", "⬜", "🔵", "⬜", "🔵", "⬜", "🔵"},
			{"⬛", "⬜", "⬛", "⬜", "⬛", "⬜", "⬛", "⬜"},
			{"⬜", "🔵", "⬜", "⬛", "⬜", "⬛", "⬜", "⬛"},
			{"🔴", "⬜", "🔴", "⬜", "🔴", "⬜", "🔴", "⬜"},
			{"⬜", "🔴", "⬜", "🔴", "⬜", "🔴", "⬜", "🔴"},
			{"🔴", "⬜", "🔴", "⬜", "🔴", "⬜", "🔴", "⬜"},
		},
	}
}

func (b *Board) Display() {
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 8; j++ {
			fmt.Print(b.board[i][j])
		}
		fmt.Println()
	}
	fmt.Print("   ")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

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
	hasPiece := func(x, y int) bool {
		return !emptyCell(x, y) && b.board[xFrom][yFrom] != b.board[x][y]
	}

	xLap := [4]int{2, 2, -2, -2}
	xIn := [4]int{1, 1, -1, -1}
	yLap := [4]int{2, -2, 2, -2}
	yIn := [4]int{1, -1, 1, -1}

	used := [8][8]bool{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			used[i][j] = false
		}
	}

	correctMove := false

	var checkTakes func(x, y int)

	checkTakes = func(x, y int) {
		used[x][y] = true
		cnt := 0
		for k := 0; k < 4; k++ {
			dx := x + xLap[k]
			dy := y + yLap[k]
			dxIn := x + xIn[k]
			dyIn := y + yIn[k]
			if !outOfBoundaries(dx) && !outOfBoundaries(dy) && hasPiece(dxIn, dyIn) && emptyCell(dx, dy) && !used[dx][dy] {
				piece := b.board[dxIn][dyIn]
				if (dxIn+dyIn)%2 == 0 {
					b.board[dxIn][dyIn] = "⬜"
				} else {
					b.board[dxIn][dyIn] = "⬛"
				}
				checkTakes(dx, dy)
				if correctMove {
					return
				}
				cnt++
				b.board[dxIn][dyIn] = piece
			}
		}
		if cnt == 0 {
			fmt.Println(x, y)
			if x == xTo && y == yTo {
				correctMove = true
			}
		}
	}

	if outOfBoundaries(xFrom) || outOfBoundaries(yFrom) || outOfBoundaries(xTo) || outOfBoundaries(yTo) {
		return false
	}
	if emptyCell(xFrom, yFrom) || !emptyCell(xTo, yTo) {
		return false
	}

	checkTakes(xFrom, yFrom)
	if correctMove {
		updateBoard(xFrom, yFrom, xTo, yTo)
		return true
	}

	if invalidMove(xFrom, yFrom, xTo, yTo) {
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
