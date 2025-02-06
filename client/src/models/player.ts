import { UpdatePosition } from '../logic/movement'
import { animate } from '../painting/animations'
import { Flipped, Source } from '../painting/painting'
import type { GameObject } from './models'

const src = Source('mage')

export function player(player: GameObject) {
  UpdatePosition(player)

  if (player.vel_x > 5) player.animationData.flipped = false
  else if (player.vel_x < -5) player.animationData.flipped = true

  if (player.animationData.flipped) {
    Flipped(player, () => {
      animate(player, src, 4, 5)
    })
  } else {
    animate(player, src, 4, 5)
  }
}
