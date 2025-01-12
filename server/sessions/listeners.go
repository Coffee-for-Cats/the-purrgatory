package sessions

import (
	"purrgatory/game"
)

// encapsulates player communication logic using session.
func PlayerHandler(session session) {
	go controlsUpdater(session)
	playerUpdater(session)
}

// blocking op
func controlsUpdater(session session) {
	player, ok := session.GameRoom.GameMap[session.PlayerId].(*game.Player)
	if !ok {
		session.Disconnect()
	}

	for {
		swap := game.PlayerControls{}
		err := session.Connection.ReadJSON(&swap)

		// player disconnected
		if err != nil {
			session.Disconnect()
			return
		}
		player.UpdateControls(swap)
	}
}

// blocking op
func playerUpdater(session session) {
	for {
		<-*session.SyncChan
		err := session.Connection.WriteJSON(session.GameRoom.GameMap)

		// player disconnected
		if err != nil {
			session.Disconnect()
			return
		}
	}
}
