package game

import "purrgatory/config"

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
	p.Y += int(p.VelY)
	p.X += int(p.VelX)

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

	occupySpace(p, g)
}

func (p *Player) setup() {
	p.Type = "player"
	p.Solid = true
	p.X = 20
	p.Y = -30
	p.Width = 16 * config.ZOOM_FACTOR
	p.Height = 16 * config.ZOOM_FACTOR
}
