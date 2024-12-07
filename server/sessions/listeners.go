package sessions

import (
	"purrgatory/game"
)

func PlayerHandler(session Session) {
	// listens conn.ReadMessage, and parses into s
	// controls must be acessed from outside, pretty tricky

	// maybe a channel based would be fancier
	go ControlsUpdater(session)
	PlayerUpdater(session)
}

// blocking, should be called with the go keyword
func ControlsUpdater(session Session) {
	player := session.GameRoom.GameMap[session.PlayerId].(*game.Player)
	for {
		session.Connection.ReadJSON(&player.Controls)
	}
}

// blocking, should be called with the go keyword
func PlayerUpdater(session Session) {
	for {
		<-*session.SyncChan
		session.Connection.WriteJSON((session.GameRoom.GameMap))
	}
}
