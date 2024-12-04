package main

import (
	"net/http"
)

func main() {

	// this route directly connects to a "session"
	http.HandleFunc("/new", CreateGame)
	http.HandleFunc("/{gameId}", EnterGame)

	http.ListenAndServe(":8080", nil)
}
