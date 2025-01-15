package game

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

// called in a new thread
func (GameRoom *GameRoom) Start() {
	// 10 TPS
	gameLoop := time.NewTicker(100 * time.Millisecond)
	ticksBeforeClose := 5

	// start game loop
	for range gameLoop.C {
		// do the entire game logic =)
		for _, object := range GameRoom.GameMap {
			object.step(GameRoom)
		}

		// sends a signal for everything "listening", just for syncronization
		for _, c := range *GameRoom.Sync {
			c <- true
		}

		// closes after 5 ticks with no listeners
		if len(*GameRoom.Sync) == 0 {
			ticksBeforeClose--
			if ticksBeforeClose < 0 {
				gameLoop.Stop()
				break
			}
		}
	}
	fmt.Println("A session was closed")
}

func (GameRoom *GameRoom) Append(object GameObject) (id string) {
	object.setup()

	if object.about().Type == "" {
		panic("You missed to set the Type field for an object added into the map.")
	}

	// generates a random id
	randBytes := make([]byte, 8)
	rand.Read(randBytes)
	randomId := base64.StdEncoding.EncodeToString(randBytes)

	// no idea if this is going to work
	GameRoom.GameMap[randomId] = object
	return randomId
}
