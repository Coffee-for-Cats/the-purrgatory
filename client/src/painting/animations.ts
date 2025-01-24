import { ZOOM_FACTOR } from '../configs'
import type { GameObject } from '../models/models'
import { ctx } from './painting'

export function animate(
  obj: GameObject,
  src: HTMLCanvasElement,
  steps: number,
  waits: number,
) {
  if (++obj.animationData.countToStep >= waits) {
    obj.animationData.countToStep = 0
    obj.animationData.actualStep++
  }
  if (obj.animationData.actualStep >= steps) {
    obj.animationData.actualStep = 0
  }

  const sWidth = src.width / steps
  const sHeigth = src.height
  const dWidth = sWidth * ZOOM_FACTOR
  const dHeight = sHeigth * ZOOM_FACTOR
  // shifts the image source x to match actual step.
  const sx = sWidth * obj.animationData.actualStep
  const sy = 0
  const dx = obj.x - dWidth / 2
  const dy = -obj.y - dHeight / 2

  ctx.drawImage(src, sx, sy, sWidth, sHeigth, dx, dy, dWidth, dHeight)
}
