package game

type Example struct {
	Physical
	//some int
}

func (e *Example) step(room *GameRoom) {}

func (e *Example) setup() {
	e.Type = "example"
}

// you don't need to overwrite neither the about function, but you can
// func (e *Example) about() Physical {
// 	return About{
// 		X: e.X + e.some,
// 	}
// }
