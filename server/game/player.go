package game

type PlayerControls struct {
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
}

type Player struct {
	Physical
	controls PlayerControls
}

func (p *Player) UpdateControls(newControls PlayerControls) {
	p.controls = newControls
}

var playerAcc float32 = 12.0
var airRes float32 = 0.8

func (p *Player) step(g *GameRoom) {

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

	new_pos := occupySpace(p, g)
	p.X = *new_pos.X
	p.Y = *new_pos.Y
}

func (p *Player) setup() {
	p.Type = "player"
	p.Solid = true
	p.X = 20
	p.Y = -30
	p.Width = 64
	p.Height = 64
}
