import { FRAMES_PER_TICK } from '../configs'
import type { GameObject } from '../models/models'

type Movable = {
  x: number
  y: number
  vel_x: number
  vel_y: number
}

export function UpdatePosition(obj: GameObject) {
  const mov = obj as unknown as Movable
  mov.x += mov.vel_x / FRAMES_PER_TICK
  mov.y += mov.vel_y / FRAMES_PER_TICK
}
