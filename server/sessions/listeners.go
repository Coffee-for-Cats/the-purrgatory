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
		swap := game.PlayerControls{}
		err := session.Connection.ReadJSON(&swap)
		if err != nil {
			// removes the player
			delete(session.GameRoom.GameMap, session.PlayerId)
			// closes the connection and breaks the for loop
			return
		}
		player.UpdateControls(swap)
	}
}

// blocking, should be called with the go keyword
func PlayerUpdater(session Session) {
	for {
		<-*session.SyncChan
		session.Connection.WriteJSON((session.GameRoom.GameMap))
	}
}
