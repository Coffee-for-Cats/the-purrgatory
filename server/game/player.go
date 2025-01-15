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

func (p *Player) setup() {
	p.Solid = true
	p.Type = "player"
	p.X = 20
	p.Y = -30
}
