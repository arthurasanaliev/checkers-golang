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
	// helper variables
	absInt := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	moveX := [4]int{2, 2, -2, -2}
	moveY := [4]int{2, -2, 2, -2}
	takeX := [4]int{1, 1, -1, -1}
	takeY := [4]int{1, -1, 1, -1}

	// logic funtions
	printMoveStop := func() {
		fmt.Println("Another player moves")
	}
	printInvalidMove := func() {
		fmt.Println("Invalid move")
	}
	printMoveGo := func() {
		fmt.Println("Continue moving")
	}
	outOfBounds := func(x int) bool {
		return x < 0 || x >= 8
	}
	emptyCell := func(x, y int) bool {
		return b.board[x][y] == "⬜" || b.board[x][y] != "⬛"
	}
	redPiece := func(x, y int) bool {
		return b.board[x][y] == "🔴"
	}
	bluePiece := func(x, y int) bool {
		return b.board[x][y] == "🔵"
	}
	difPieces := func(x1, y1, x2, y2 int) bool {
		return (redPiece(x1, y1) && bluePiece(x2, y2) || 
						bluePiece(x1, y1) && redPiece(x2, y2))
	}
	validMove := func(x, y int) (bool, int, int) {
		for k := 0; k < 4; k++ {
			r, c := x + moveX[k], y + moveY[k]
			rTake, cTake := x + takeX[k], y + takeY[k]
			if !outOfBounds(r) && !outOfBounds(c) && difPieces(x, y, rTake, cTake) {
				if r == xTo && c == yTo {
					return true, rTake, cTake
				}
			}
		}
		return false, -1, -1
	}
	switchCells := func(x1, y1, x2, y2 int) {
		b.board[x1][y1], b.board[x2][y2] = b.board[x2][y2], b.board[x1][y1]
	}

	if outOfBounds(xFrom) || outOfBounds(yFrom) || outOfBounds(xTo) || outOfBounds(yTo) {
		printInvalidMove()
		return false
	}
	if emptyCell(xFrom, yFrom) && !emptyCell(xTo, yTo) {
		printInvalidMove()
		return false
	}
	good, r, c := validMove(xFrom, yFrom)
	if good {
		switchCells(xFrom, yFrom, xTo, yTo)

	}
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
