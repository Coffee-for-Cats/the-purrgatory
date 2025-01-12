import type { GameObject } from '../models/models'

export type GameMapType = Map<string, GameObject>
export const GameMap: GameMapType = new Map()

let socketRef: WebSocket | null = null
export function Listen(Socket: WebSocket) {
  socketRef = Socket
  Socket.onmessage = defaultMapUpdater
}

function defaultMapUpdater(ev: MessageEvent) {
  const newState: GameMapType = JSON.parse(ev.data)

  // creates the object if it doesn't exist
  for (const key of Object.keys(newState)) {
    if (!GameMap[key]) {
      GameMap[key] = newState[key]
    }
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
      // no need for "/ FRAMES_PER_TICK" because the position is already updated
      // considering the frames.
      GameMap[key].vel_x = newState[key].vel_x + dif_x
      GameMap[key].vel_y = newState[key].vel_y + dif_y
    }
  }
}

// updates the page instantly in the next frame when the page was offscreen
addEventListener('visibilitychange', () => {
  if (document.visibilityState !== 'hidden' && socketRef)
    socketRef.onmessage = fasterMapUpdater
})

// lasts 1 tick, updates the screen instantly and resets to the normal execution
function fasterMapUpdater(ev: MessageEvent) {
  if (!socketRef) return
  const newState: GameMapType = JSON.parse(ev.data)
  // not directly assinging to keep GameMap as a ES6 Map instead of Object.
  for (const key of Object.keys(newState)) {
    GameMap[key] = newState[key]
  }
  socketRef.onmessage = defaultMapUpdater
}
