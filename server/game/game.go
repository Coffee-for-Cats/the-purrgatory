package game

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type GameMap map[string]GameObject

type GameRoom struct {
	GameMap GameMap
	Sync    *[]chan bool
}

func (GameRoom *GameRoom) Start() {
	gameLoop := time.NewTicker(200 * time.Millisecond)

	// start game loop
	go func() {
		for range gameLoop.C {
			// do the entire game logic =)
			for _, object := range GameRoom.GameMap {
				object.step(GameRoom)
			}

			// a list of channels to sync the gameloop with
			// listeners like session.PlayerHandler()
			for _, c := range *GameRoom.Sync {
				c <- true
			}
		}
	}()
}

func (GameRoom *GameRoom) Append(object GameObject) (id string) {
	// generates a random id
	randBytes := make([]byte, 8)
	rand.Read(randBytes)
	randomId := base64.StdEncoding.EncodeToString(randBytes)

	// no idea if this is going to work
	GameRoom.GameMap[randomId] = object
	return randomId
}
