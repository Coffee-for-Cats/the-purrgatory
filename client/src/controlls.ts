export const Controls = {
  ArrowDown: false,
  ArrowUp: false,
  ArrowRight: false,
  ArrowLeft: false,
}

export function Update(Socket: WebSocket) {
  // set up the listeners for keyboard interaction
  window.addEventListener('keydown', (e) => {
    if (!Controls[e.key] && Controls[e.key] !== undefined) {
      Controls[e.key] = true
      updateControls(Socket)
    }
  })
  window.addEventListener('keyup', (e) => {
    if (Controls[e.key] && Controls[e.key] !== undefined) {
      Controls[e.key] = false
      updateControls(Socket)
    }
  })
}

function updateControls(Socket: WebSocket) {
  Socket.send(JSON.stringify(Controls))
}
