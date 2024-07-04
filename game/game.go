package game

import "fmt"

type Game struct {
	board         *Board
	player1       *Player
	player2       *Player
	currentPlayer *Player
}

func NewGame() *Game {
	board := NewBoard()
	player1 := NewPlayer("Red Player", "red")
	player2 := NewPlayer("Blue Player", "blue")

	return &Game{
		board:         board,
		player1:       player1,
		player2:       player2,
		currentPlayer: player1,
	}
}

func (g *Game) Start() {
	for {
		g.board.Display()
		xFrom, yFrom, xTo, yTo := g.currentPlayer.GetMove()
		if !g.board.CheckMove(xFrom, yFrom, xTo, yTo) {
			continue
		}
		if g.board.CheckWin() {
			g.board.Display()
			fmt.Printf("%s wins!\n", g.currentPlayer.name)
			return
		}
		g.switchPlayers()
	}
}

func (g *Game) switchPlayers() {
	if g.currentPlayer == g.player1 {
		g.currentPlayer = g.player2
	} else {
		g.currentPlayer = g.player1
	}
}
