package game

type Example struct {
	Entity
	some int
}

func (e *Example) step(room *GameRoom) {}

func (e *Example) typeSet() {
	e.TypeName = "example"
	e.some++
}
