package sessions

func PlayerHandler(session Session) {
	// listens conn.ReadMessage, and parses into controlls
	// controlls must be acessed from outside, pretty tricky
	for {
		// TODO: select { ... }
		// controlls

		// awaits for sync channel and updates the player
		<-*session.SyncChan
		session.Connection.WriteJSON(session.GameRoom.GameMap)
	}
}
