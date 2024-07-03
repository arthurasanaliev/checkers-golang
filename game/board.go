package game

import "fmt"

type Board struct {
	board [8][8]string
}

func NewBoard() *Board {
	return &Board{
		[8][8]string{
			{"⬜", "🔵", "⬜", "🔵", "⬜", "🔵", "⬜"},
			{"🔵", "⬜", "🔵", "⬜", "🔵", "⬜", "🔵"},
			{"⬜", "🔵", "⬜", "🔵", "⬜", "🔵", "⬜"},
			{"⬛", "⬜", "⬛", "⬜", "⬛", "⬜", "⬛"},
			{"⬜", "⬛", "⬜", "⬛", "⬜", "⬛", "⬜"},
			{"🔴", "⬜", "🔴", "⬜", "🔴", "⬜", "🔴"},
			{"⬜", "🔴", "⬜", "🔴", "⬜", "🔴", "⬜"},
			{"🔴", "⬜", "🔴", "⬜", "🔴", "⬜", "🔴"},
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

func (b *Board) IsWin() bool {
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
