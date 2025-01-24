import { ZOOM_FACTOR } from '../configs'
import type { GameObject } from '../models/models'

const GameCanvas = document.getElementById('game-canvas') as HTMLCanvasElement
GameCanvas.width = GameCanvas.clientWidth
GameCanvas.height = GameCanvas.clientHeight
export const ctx = GameCanvas.getContext('2d') as CanvasRenderingContext2D
ctx.imageSmoothingEnabled = false

export function ClearCanvas() {
  // paint everything here =)
  ctx.clearRect(0, 0, GameCanvas.width, GameCanvas.height)
}

export function PaintObject(obj: GameObject, src: HTMLImageElement) {
  const width = src.width * ZOOM_FACTOR
  const height = src.height * ZOOM_FACTOR

  const x = obj.x - width / 2
  const y = -obj.y - height / 2

  ctx.drawImage(src, x, y, width, height)
}

// gets the filename and returns an image with the src set to /sprites/<name>.png
export function Source(url: string): HTMLCanvasElement {
  const img = document.createElement('img')
  img.src = `./sprites/${url}.png`

  const canvas = document.createElement('canvas')

  img.onload = () => {
    canvas.width = img.width
    canvas.height = img.height
    canvas.getContext('2d')?.drawImage(img, 0, 0)
  }

  return canvas
}
