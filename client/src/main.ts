import { ClearCanvas, Paint } from './animations'
import { Update } from './controlls'
import { GameMap, Listen } from './map'

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

console.log(`Logged into session ${roomId}`)

export const Socket = new WebSocket(`http://localhost:8080/${roomId}`)
Socket.onerror = (e) => console.error(e)

// listen for server updates
Listen(Socket)
// updates the server on key press
Update(Socket)

export function Step(timestamp: number) {
  ClearCanvas()
  for (const obj of Object.values(GameMap)) {
    console.log(obj)
    Paint(obj as GameObject)
  }
  requestAnimationFrame(Step)
}
window.requestAnimationFrame(Step)

export interface GameObject {
  x: number
  y: number
  vel_x: number
  vel_y: number
  type: string
}
