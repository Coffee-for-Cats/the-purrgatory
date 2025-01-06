import { UpdatePos } from '../logic/movement'
import type { GameObject } from '../main'
import { PaintObject, Source } from '../painting/animations'

const src = Source('bluemage')
export function player(player: GameObject) {
  UpdatePos(player)
  PaintObject(player, src)
}
