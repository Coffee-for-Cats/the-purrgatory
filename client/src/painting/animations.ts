import type { GameObject } from '../main'

const GameCanvas = document.getElementById('game-canvas') as HTMLCanvasElement
GameCanvas.width = GameCanvas.clientWidth
GameCanvas.height = GameCanvas.clientHeight
const ctx = GameCanvas.getContext('2d') as CanvasRenderingContext2D
ctx.imageSmoothingEnabled = false

export function ClearCanvas() {
  // paint everything here =)
  ctx.clearRect(0, 0, GameCanvas.width, GameCanvas.height)
}

export function PaintObject(obj: GameObject, src: HTMLImageElement) {
  ctx.drawImage(src, obj.x, -obj.y, 64, 64)
}

// gets the filename and returns an image with the src set to /sprites/<name>.png
export function Source(url: string) {
  const img = document.createElement('img')
  img.src = `./sprites/${url}.png`
  return img
}
