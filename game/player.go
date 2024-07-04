package game

import "fmt"

type Player struct {
	name  string
	color string
}

func NewPlayer(name, color string) *Player {
	return &Player{
		name:  name,
		color: color,
	}
}

func (p *Player) GetMove() (int, int, int, int) {
	var xFrom, yFrom, xTo, yTo int
	fmt.Printf("%s, choose a piece you want to move: ", p.name)
	fmt.Scanf("%d %d", &xFrom, &yFrom)
	fmt.Print("Now choose a cell you want to move to: ")
	fmt.Scanf("%d %d", &xTo, &yTo)
	return xFrom, yFrom, xTo, yTo
}
