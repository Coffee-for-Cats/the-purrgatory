import { FRAMES_PER_TICK } from './configs'
import type { GameObject } from './main'

type Map = { [index: string]: GameObject }
export const GameMap: Map = {}

export function Listen(Socket: WebSocket) {
  Socket.onmessage = (ev) => {
    const newState = JSON.parse(ev.data) as Map

    // creates the object if it doesn't exist
    for (const key of Object.keys(newState)) {
      if (!GameMap[key]) GameMap.key = newState[key]
    }

    for (const key of Object.keys(newState)) {
      console.log(key)
      if (!newState[key] && GameMap[key]) {
        // should not exist and exists
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
}
