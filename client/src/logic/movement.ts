import { FRAMES_PER_TICK } from '../configs'
import type { GameObject } from '../main'

export function UpdatePosition(obj: GameObject) {
  obj.x += obj.vel_x / FRAMES_PER_TICK
  obj.y += obj.vel_y / FRAMES_PER_TICK
}
