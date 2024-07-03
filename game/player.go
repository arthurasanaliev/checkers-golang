package game

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
