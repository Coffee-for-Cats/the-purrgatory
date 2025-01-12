import { UpdatePosition } from '../logic/movement'
import { type Movable, PaintObject, Source } from '../painting/animations'
import type { GameObject } from './models'

const src = Source('bluemage')
export function player(player: GameObject) {
  UpdatePosition(player as Movable)
  PaintObject(player as Movable, src)
}
