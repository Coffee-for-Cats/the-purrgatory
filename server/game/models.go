package game

type GameMap map[string]GameObject

type GameRoom struct {
	GameMap GameMap
	Sync    *[]chan bool
}

// GameObject is a object, wich it's reference implements step()
type GameObject interface {
	step(*GameRoom)
	setup()
	// comes from Physical and Entity
	about() About
}

type Entity struct {
	Type string `json:"type"`
}

type Physical struct {
	Entity
	Solid  bool
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int
	Height int
	VelX   float32 `json:"vel_x"`
	VelY   float32 `json:"vel_y"`
}

func (p *Physical) about() About {
	return About{
		Type:   p.Type,
		Solid:  p.Solid,
		X:      p.X,
		Y:      p.Y,
		Width:  p.Width,
		Height: p.Height,
	}
}

type About struct {
	Type   string
	Solid  bool
	X      int
	Y      int
	Width  int
	Height int
}
