package game

import "fmt"

// GameObject is a object, wich it's reference implements step()
type GameObject interface {
	step(*GameRoom)
}

type vector struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"vel_x"`
	VelY int `json:"vel_y"`
}

type Player struct {
	Name     string
	Controls PlayerControls
	vector
}

type PlayerControls struct {
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
}

func (p *Player) step(g *GameRoom) {
	// yay =)
	switch {
	case p.Controls.ArrowDown:
		p.Y--
	case p.Controls.ArrowUp:
		p.Y++
	case p.Controls.ArrowRight:
		p.X++
	case p.Controls.ArrowLeft:
		p.X--
	}
}

type Enemy struct {
	enemyType string
	vector
}

func (e *Enemy) step(g *GameRoom) {
	fmt.Println("Enemy moved!")
}
