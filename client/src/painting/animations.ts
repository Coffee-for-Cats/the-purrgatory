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

  // sWidth needs to only be 1 frame of the image, not it full
  const Width = src.width / steps
  const Height = src.height
  // shifts the image source x to match actual step.
  const sx = Width * obj.animationData.actualStep
  const sy = 0
  const dx = obj.x - Width / 2
  const dy = obj.y + Height / 2

  ctx.drawImage(src, sx, sy, Width, Height, dx, -dy, Width, Height)
}
