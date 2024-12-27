package sessions

import (
	"purrgatory/game"

	"github.com/gorilla/websocket"
)

type session struct {
	Connection *websocket.Conn
	GameRoom   *game.GameRoom
	PlayerId   string
	SyncChan   *chan bool
}

func Connect(conn *websocket.Conn, GameRoom *game.GameRoom, PlayerId string) session {
	newSync := make(chan bool, 1)
	*GameRoom.Sync = append(*GameRoom.Sync, newSync)

	newSession := session{
		Connection: conn,
		GameRoom:   GameRoom,
		PlayerId:   PlayerId,
		SyncChan:   &newSync,
	}

	return newSession
}

func (session session) Disconnect() {
	delete(session.GameRoom.GameMap, session.PlayerId)

	// finds and "deletes" the sync channel
	index := -1
	for i, c := range *session.GameRoom.Sync {
		if c == *session.SyncChan {
			index = i
		}
	}
	sync := *session.GameRoom.Sync
	if index < 0 || index > len(sync) {
		panic("Session disconnection gone wrong!")
	}
	sync[index] = sync[len(sync)-1]
	*session.GameRoom.Sync = sync[:len(sync)-1]
}
