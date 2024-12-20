import './animate'
import { Update } from './controlls'
import { Listen } from './map'

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

alert(`Logged into session ${roomId}`)

export const Socket = new WebSocket(`http://localhost:8080/${roomId}`)
Socket.onerror = (e) => console.error(e)
Listen(Socket)
Update(Socket)
