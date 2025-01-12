import { Update } from './logic/controlls'
import { GameMap, Listen } from './logic/map'
import { Models } from './models/models'
import { ClearCanvas } from './painting/animations'

let roomId = prompt('RoomId: ')
if (!roomId) {
  try {
    const request = await fetch('http://localhost:8080/new')
    roomId = await request.text()
  } catch {
    alert('Error finding server.')
  }

  // this line replaces the URL field without redirecting.
  // this is not :8080 port, but 5500. Thats because frontend and backend are separate
  // window.history.pushState(null, '', `http://127.0.0.1:5500/${roomId}`)
  // TODO: make a redirect in the hosting to redirect any link to this page
}

export const Socket = new WebSocket(`http://localhost:8080/${roomId}`)
Socket.onerror = (e) => console.error(e)

// listen for server updates
Listen(Socket)
// updates the server on key press
Update(Socket)

window.requestAnimationFrame(Step)
function Step(timestamp: number) {
  ClearCanvas()
  for (const key in GameMap) {
    const obj = GameMap[key]
    const typeFunc = Models[obj.type]
    if (!typeFunc) throw new Error(`Unkown type returned: ${obj.typeName}`)
    Models[obj.type](obj)
  }
  requestAnimationFrame(Step)
}
