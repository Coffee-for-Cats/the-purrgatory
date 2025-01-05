import { FRAMES_PER_TICK } from './configs'
import type { GameObject } from './main'

const GameCanvas = document.getElementById('game-canvas') as HTMLCanvasElement
GameCanvas.width = GameCanvas.clientWidth
GameCanvas.height = GameCanvas.clientHeight
const ctx = GameCanvas.getContext('2d') as CanvasRenderingContext2D
ctx.imageSmoothingEnabled = false

const blueMage = document.createElement('img')
blueMage.src = './sprites/bluemage.png'

export function ClearCanvas() {
  // paint everything here =)
  ctx.clearRect(0, 0, GameCanvas.width, GameCanvas.height)
}

export function Paint(obj: GameObject) {
  obj.x += obj.vel_x / FRAMES_PER_TICK
  obj.y += obj.vel_y / FRAMES_PER_TICK
  ctx.drawImage(blueMage, obj.x, -obj.y, 64, 64)
}
