package game

type Tree struct {
	Entity
	Physical
}

func (t *Tree) step(g *GameRoom) {

}

func (t *Tree) setup() {
	t.Type = "tree"
}

// func (t *Tree) about() About {
// 	t.Type = "tree"

// 	return About{
// 		Type:   &"tree",
// 		Solid:  &true,
// 		X:      300,
// 		Y:      -400,
// 		Width:  40,
// 		Height: 60,
// 	}
// }
