# PurrGatory
  A web-based, vampire-survivors-like multiplayer game.
  It's low on production dependencies: Golang & Typescript + BiomeJS

## server
  The server has 2 routes, one for creating and other for connecting to a room.
  A Room doesn't contain connections, it is a separate goroutine for isolation and performance; the player object is a struct and it's refence is saved in said goroutine.
The room contains a GameMap which is a hashmap of GameObject interfaces. You can make your gameobject structs embeeding "Physical". The type is a necessary field to be defined at setup.
  setup() is like a constructor function, and it is useful to share behavior.
  obj.about() return the Physical's fields as references if you need to modify objects generically.

## client
The frontend is built using vite, typescript and BiomeJS for linting.
The client updates the speed to match the server's, and always adds a small compensation if their positions diverge.
This ensures smooth position updating and is lightweight.

### how to run
`bash dev.sh` or 
`go run .` in the /server and `npm run dev` in the /client
