const request = await fetch("http://localhost:8080/new");
const roomId = await request.text();

const socket = new WebSocket(`http://localhost:8080/${roomId}`)
socket.onerror = (e) => {
  console.error(e)
}
socket.onmessage = (ev) => {
  console.log(ev.data)
}