package game

import "purrgatory/config"

type Tree struct {
	Physical
}

func (t *Tree) step(g *GameRoom) {

}

func (t *Tree) setup() {
	t.Type = "tree"
	t.Solid = true
	t.Width = 48 * config.ZOOM_FACTOR
	t.Height = 96 * config.ZOOM_FACTOR
}
