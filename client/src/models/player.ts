import { ZOOM_FACTOR } from '../configs'
import { UpdatePosition } from '../logic/movement'
import { animate } from '../painting/animations'
import { ctx, Source } from '../painting/painting'
import type { GameObject } from './models'

const src = Source('mage')

export function player(player: GameObject) {
  UpdatePosition(player)

  if (player.vel_x > 5) player.animationData.flipped = false
  else if (player.vel_x < -5) player.animationData.flipped = true

  if (player.animationData.flipped) {
    ctx.save()
    ctx.scale(-1, 1)
    ctx.translate(-player.x * 2, 0)
    animate(player, src, 4, 5)
    ctx.restore()
  } else {
    animate(player, src, 4, 5)
  }
}
