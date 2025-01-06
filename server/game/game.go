package game

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

func (GameRoom *GameRoom) Start() {
	// 10 TPS
	gameLoop := time.NewTicker(100 * time.Millisecond)

	// start game loop
	go func() {
		for range gameLoop.C {
			// do the entire game logic =)
			for _, object := range GameRoom.GameMap {
				object.step(GameRoom)
			}

			for _, c := range *GameRoom.Sync {
				c <- true
			}
		}
	}()
}

func (GameRoom *GameRoom) Append(object GameObject) (id string) {
	object.typeSet() // just to have a reference to the object type

	// generates a random id
	randBytes := make([]byte, 8)
	rand.Read(randBytes)
	randomId := base64.StdEncoding.EncodeToString(randBytes)

	// no idea if this is going to work
	GameRoom.GameMap[randomId] = object
	return randomId
}
