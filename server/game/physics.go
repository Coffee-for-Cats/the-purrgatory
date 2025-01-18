package game

import (
	"math"
)

// for correct player movement, the game should update the axis individually.
// here, we check for every object on screen, checking first how much space the object
// needs in that axis after moving, and adjusting it's speed to match said space.
// Then, we consider that it already moved (even tho the p.Y did not changed), because
// the object could ignore other objects that are in the diagonal to it's movement.
// It's still crude, but works well enough.

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

		x1, x2 := float64(*p.X), float64(*obj.X)
		y1, y2 := float64(*p.Y), float64(*obj.Y)
		w1, w2 := float64(*p.Width), float64(*obj.Width)
		h1, h2 := float64(*p.Height), float64(*obj.Height)

		distX := math.Abs(x1-x2) - (w1/2 + w2/2)
		distYnext := math.Abs(y1-y2+float64(*p.VelY)) - (h1/2 + h2/2)

		if distX < 0 && distYnext < 0 {
			*p.VelY += float32(distYnext * signFrom(*p.VelY))
		}

		// temporary Y position to catch diagonal contact
		tempY := y1 + float64(*p.VelY)

		distY := math.Abs(tempY-y2) - (h1/2 + h2/2)
		distXnext := math.Abs(x1-x2+float64(*p.VelX)) - (w1/2 + w2/2)

		if distXnext < 0 && distY < 0 {
			*p.VelX += float32(distXnext * signFrom(*p.VelX))
		}
	}

	return p
}

func signFrom(num float32) float64 {
	if num < 0 {
		return -1.0
	}
	return 1.0
}
