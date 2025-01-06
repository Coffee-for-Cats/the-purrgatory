package game

type GameMap map[string]GameObject

type GameRoom struct {
	GameMap GameMap
	Sync    *[]chan bool
}

// GameObject is a object, wich it's reference implements step()
type GameObject interface {
	step(*GameRoom)
	typeSet()
}

type Generic[T any] struct {
	Data     T
	TypeName string  `json:"type"`
	X        int     `json:"x"`
	Y        int     `json:"y"`
	VelX     float32 `json:"vel_x"`
	VelY     float32 `json:"vel_y"`
}
