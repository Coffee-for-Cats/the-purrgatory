package game

import "fmt"

// GameObject is a object, wich it's reference implements step()
type GameObject interface {
	step(*GameRoom)
}

type vector struct {
	X    int `json:"x"`
	Y    int `json:"y"`
	VelX int `json:"vel_x"`
	VelY int `json:"vel_y"`
}

type Player struct {
	Name string
	vector
}

func (p *Player) step(g *GameRoom) {
	p.X++
	// fmt.Println("Player x: ", p.X)
}

type Enemy struct {
	enemyType string
	vector
}

func (e *Enemy) step(g *GameRoom) {
	fmt.Println("Enemy moved!")
}
