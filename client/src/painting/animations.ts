import { ZOOM_FACTOR } from '../configs'
import type { GameObject } from '../models/models'
import { ctx } from './painting'

export function animate(
  obj: GameObject,
  src: HTMLImageElement,
  steps: number,
  waits: number,
) {
  if (!obj.animationData) obj.animationData = { countToStep: 0, actualStep: 0 }

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
  const src_x = sWidth * obj.animationData.actualStep
  const src_y = 0
  const x = obj.x - dWidth / 2
  const y = -obj.y - dHeight / 2

  ctx.drawImage(src, src_x, src_y, sWidth, sHeigth, x, y, dWidth, dHeight)
}
