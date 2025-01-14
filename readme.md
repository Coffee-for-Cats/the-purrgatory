# PurrGatory
A web-based, vampire-survivors-like multiplayer game.

## Sever
The server is written in golang, you probably don't need anything other than go in your machine to run.
The server has 2 routes, one for creating and other for connecting to a room.
A Room doesn't contain players, it is a separate goroutine for isolation and performance.
The room contains a map, each gameObject must have a "type" json property setted by setType, and must implement a step() function.
This part of the backend is to ensure your types are correct and are parsed to the frontend.

## client
The client recieves the positions of every object in the screen.
The game then updates the speed, and adds a velocity to match the position for the next tick.
This ensures smooth position updating and is lightweight.
The frontend is built using vite.

### how to run
`bash dev.sh` or `go run .` in the /server and `npm run dev` in the /client
