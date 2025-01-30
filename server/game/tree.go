package game

type Tree struct {
	Physical
}

func (t *Tree) step(g *GameRoom) {

}

func (t *Tree) setup() {
	t.Type = "tree"
	t.Solid = true
	t.Width = 48
	t.Height = 96
}
