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

export function PaintObject(obj: GameObject, src: HTMLCanvasElement) {
  const x = obj.x - src.width / 2
  const y = obj.y + src.height / 2

  ctx.drawImage(src, x, -y)
}

// gets the filename and returns an image with the src set to /sprites/<name>.png
export function Source(url: string): HTMLCanvasElement {
  const img = document.createElement('img')
  img.src = `./sprites/${url}.png`

  const canvas = document.createElement('canvas')

  img.onload = () => {
    canvas.width = img.width * ZOOM_FACTOR
    canvas.height = img.height * ZOOM_FACTOR
    const ctx = canvas.getContext('2d')
    if (!ctx) throw new Error('Failed finding ctx for image')
    ctx.imageSmoothingEnabled = false
    ctx.drawImage(img, 0, 0, canvas.width, canvas.height)
  }

  return canvas
}
