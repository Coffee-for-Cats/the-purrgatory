package game

import "fmt"

// GameObject is a object, wich it's reference implements step()
type GameObject interface {
	step(*GameRoom)
}

type vector struct {
	X    int     `json:"x"`
	Y    int     `json:"y"`
	VelX float32 `json:"vel_x"`
	VelY float32 `json:"vel_y"`
}

type Player struct {
	Name     string
	controls PlayerControls
	vector
}

type PlayerControls struct {
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
}

func (p *Player) UpdateControls(newControls PlayerControls) {
	p.controls = newControls
}

var playerAcc float32 = 25.0
var airRes float32 = 0.8

func (p *Player) step(g *GameRoom) {

	p.X += int(p.VelX)
	p.Y += int(p.VelY)

	p.VelX *= 1 - airRes
	p.VelY *= 1 - airRes

	if p.controls.ArrowRight {
		p.VelX += playerAcc
	}
	if p.controls.ArrowLeft {
		p.VelX -= playerAcc
	}
	if p.controls.ArrowUp {
		p.VelY += playerAcc
	}
	if p.controls.ArrowDown {
		p.VelY -= playerAcc
	}
}

type Enemy struct {
	enemyType string
	vector
}

func (e *Enemy) step(g *GameRoom) {
	fmt.Println("Enemy moved!")
}
