package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"purrgatory/game"
	"purrgatory/sessions"
	"unicode/utf8"

	"github.com/gorilla/websocket"
)

var RunningGames = make(map[string]game.GameRoom)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func EnterGame(w http.ResponseWriter, req *http.Request) {
	roomPath := req.URL.Path
	_, i := utf8.DecodeRuneInString(roomPath) // i = size of first character
	gameId := roomPath[i:]                    // slice it and skips the i bytes
	fmt.Println("Player connecting to " + gameId)
	connectPlayer(gameId, w, req)
}

func CreateGame(w http.ResponseWriter, req *http.Request) {
	randBytes := make([]byte, 6)
	rand.Read(randBytes)
	randomRoomId := base64.URLEncoding.EncodeToString(randBytes)

	// no idea why this can't be inline

	// creates a game room
	channel := make([]chan bool, 0)
	room := game.GameRoom{
		GameMap: make(map[string]game.GameObject),
		Sync:    &channel,
	}
	// and saves it
	RunningGames[randomRoomId] = room

	// I just wasted 50 minutes because of this "go"
	go room.Start()
	fmt.Println("Started room " + randomRoomId)

	// TODO: replace this line:
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%s", randomRoomId)
}

func connectPlayer(gameId string, w http.ResponseWriter, req *http.Request) {
	room, ok := RunningGames[gameId]
	if !ok {
		http.Error(w, "Room doesn't exist!", 404)
		fmt.Println("Room not found:" + gameId)
		return
	}

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		panic(err)
	}

	// adds the player to the game
	player := game.Player{}
	player.Data.Name = "Arthur"

	id := room.Append(&player)

	// creates a channel, saves on the GameMap, and passes a
	// reference to Session in question.

	newSession := sessions.Connect(conn, &room, id)

	// set up a listener for the new player
	// updates him sometimes and recieves control updates too
	go sessions.PlayerHandler(newSession)
}
