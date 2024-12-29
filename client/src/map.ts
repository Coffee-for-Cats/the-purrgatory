import { FRAMES_PER_TICK } from './configs'
import type { GameObject } from './main'

type Map = { [index: string]: GameObject }
export let GameMap: Map = {}

let socketRef: WebSocket | null = null
export function Listen(Socket: WebSocket) {
  socketRef = Socket
  Socket.onmessage = defaultMapUpdater
}

function defaultMapUpdater(ev: MessageEvent) {
  const newState = JSON.parse(ev.data) as Map

  // creates the object if it doesn't exist
  for (const key of Object.keys(newState)) {
    if (!GameMap[key]) GameMap[key] = newState[key]
  }

  for (const key of Object.keys(GameMap)) {
    // should not exist and exists
    if (!newState[key] && GameMap[key]) {
      delete GameMap[key]
    } else {
      // adds a little acceleration to compensate the difference in the next tick update.
      // difference between "should" position and "is" position
      const dif_x = newState[key].x - GameMap[key].x
      const dif_y = newState[key].y - GameMap[key].y
      GameMap[key].vel_x = newState[key].vel_x + dif_x / FRAMES_PER_TICK
      GameMap[key].vel_y = newState[key].vel_y + dif_y / FRAMES_PER_TICK
    }
  }
}

// updates the page instantly in the next frame when the page was offscreen
addEventListener('visibilitychange', () => {
  if (document.visibilityState === 'hidden' || !socketRef) return
  socketRef.onmessage = fasterMapUpdater
})

// lasts 1 tick, updates the screen instantly and resets to the normal execution
function fasterMapUpdater(ev: MessageEvent) {
  if (!socketRef) return
  GameMap = JSON.parse(ev.data) as Map
  socketRef.onmessage = defaultMapUpdater
}
