package game

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

func (GameRoom *GameRoom) Start() {
	// 10 TPS
	gameLoop := time.NewTicker(100 * time.Millisecond)
	ticksBeforeClose := 5

	// start game loop
	func() {
		for range gameLoop.C {
			// do the entire game logic =)
			for _, object := range GameRoom.GameMap {
				object.step(GameRoom)
			}
			// breaks if nothing is listening to the server tick
			if len(*GameRoom.Sync) == 0 {
				ticksBeforeClose--
				if ticksBeforeClose < 0 {
					gameLoop.Stop()
					return
				}
			}
			fmt.Println("Tick")

			// sends a signal for everything "listening", just for syncronization
			for _, c := range *GameRoom.Sync {
				c <- true
			}
		}
	}()
	fmt.Println("A session was closed")
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
