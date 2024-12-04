package sessions

import (
	"purrgatory/game"

	"github.com/gorilla/websocket"
)

type Session struct {
	Connection *websocket.Conn
	GameRoom *game.GameRoom
	PlayerId string
	SyncChan *chan bool
}