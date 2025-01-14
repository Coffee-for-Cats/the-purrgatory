package game

type PlayerControls struct {
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
}

type Player Generic[struct {
	Name     string
	controls PlayerControls
}]

func (p *Player) UpdateControls(newControls PlayerControls) {
	p.Data.controls = newControls
}

var playerAcc float32 = 25.0
var airRes float32 = 0.8

func (p *Player) step(g *GameRoom) {

	p.X += int(p.VelX)
	p.Y += int(p.VelY)

	p.VelX *= 1 - airRes
	p.VelY *= 1 - airRes

	if p.Data.controls.ArrowRight {
		p.VelX += playerAcc
	}
	if p.Data.controls.ArrowLeft {
		p.VelX -= playerAcc
	}
	if p.Data.controls.ArrowUp {
		p.VelY += playerAcc
	}
	if p.Data.controls.ArrowDown {
		p.VelY -= playerAcc
	}
}

func (p *Player) typeSet() {
	p.TypeName = "player"
}
