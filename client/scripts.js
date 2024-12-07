const request = await fetch("http://localhost:8080/new");
const roomId = await request.text();

const socket = new WebSocket(`http://localhost:8080/${roomId}`)
socket.onerror = (e) => {
  console.error(e)
}
socket.onmessage = (ev) => {
  console.log(ev.data)
}


const Controls = {
  ArrowDown: false,
  ArrowUp: false,
  ArrowRight: false,
  ArrowLeft: false,
}

// send controls to the server
document.addEventListener("keydown", (e) => {
  if(Controls[e.key] !== undefined) {
    Controls[e.key] = true;
  }
  socket.send(JSON.stringify(Controls))
})

document.addEventListener("keyup", (e) => {
  if(Controls[e.key] && Controls[e.key] !== undefined) {
    Controls[e.key] = false;
  }
  socket.send(JSON.stringify(Controls))
})