# PurrGatory
  A web-based, vampire-survivors-like multiplayer game.
  It's low on production dependencies:
    The frontend is built in bare Typescript (no TSX) with Vite as build tool.
    The backend is pure Go, with Air for hot-reloading; gorillax library was needed to use websockets.
  Frontend is linted with BiomeJS and backend with the default go formatter.

## server
  The server has 2 routes, one for creating and other for connecting to a room.
  A Room doesn't contain connections, it's a goroutine call. Players are not special types of object, but a goroutine keeps track of the player for as long as it exists.
  The room contains a GameMap which is a hashmap of GameObject interfaces. You can make your gameobject structs embeeding "Physical". The `type` is a necessary field to be defined at setup.
  setup() is like a constructor function, and it is useful to share behavior.
  obj.about() return the Physical's fields as references if you need to modify objects generically.

## client
  The client updates the speed to match the server's, and always adds a small compensation if their positions diverge.
  This ensures smooth position updating and is lightweight.

### how to run
`bash dev.sh` or 
`go run .` in the /server and `npm run dev` in the /client

You will need to put assets in /client/sprites
