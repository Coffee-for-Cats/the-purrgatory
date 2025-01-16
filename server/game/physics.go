package game

import (
	"math"
)

// automatically updates the position of the player.
// Should be called before modifies in speed and acceleration
func occupySpace(primary GameObject, g *GameRoom) (modified About) {
	p := primary.about()
	if !*p.Solid {
		return
	}

	// for Y-axis movement limitation
	for _, object := range g.GameMap {
		obj := object.about()
		if !*obj.Solid || obj == p {
			continue
		}

		// this bad as shit
		distX := math.Abs(float64(*p.X-*obj.X)) - (float64(*p.Width)/2 + float64(*obj.Width)/2)
		distY := math.Abs(float64(*p.Y-*obj.Y)+float64(*p.VelY)) - (float64(*p.Height)/2 + float64(*obj.Height)/2)

		if distX < 0 && distY < 0 {
			sign := 1.0
			if math.Signbit(float64(*p.VelY)) {
				sign = -1.0
			}
			*p.VelY += float32(distY * sign)
		}
	}

	// for X-axis movement limitation
	for _, object := range g.GameMap {
		obj := object.about()
		if !*obj.Solid || obj == p {
			continue
		}

		// this bad as shit
		distX := math.Abs(float64(*p.X-*obj.X)+float64(*p.VelX)) - (float64(*p.Width)/2 + float64(*obj.Width)/2)
		distY := math.Abs(float64(*p.Y-*obj.Y)) - (float64(*p.Height)/2 + float64(*obj.Height)/2)

		if distX < 0 && distY < 0 {
			sign := 1.0
			if math.Signbit(float64(*p.VelX)) {
				sign = -1.0
			}
			*p.VelX += float32(distX * sign)
		}
	}

	return p
}
