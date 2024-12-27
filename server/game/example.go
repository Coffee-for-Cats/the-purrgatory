package game

type Example Generic[struct {
	some int
}]

func (e *Example) step(room *GameRoom) {}

func (e *Example) typeSet() {
	e.TypeName = "example"
}
