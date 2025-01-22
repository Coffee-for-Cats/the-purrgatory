import { UpdatePosition } from '../logic/movement'
import { animate } from '../painting/animations'
import { Source } from '../painting/painting'
import type { GameObject } from './models'

const src = Source('mage')
export function player(player: GameObject) {
  UpdatePosition(player)
  // PaintObject(player, src)
  animate(player, src, 4, 5)
}
