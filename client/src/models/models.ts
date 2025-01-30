import { ZOOM_FACTOR } from '../configs'
import { player } from './player'
import { tree } from './tree'

export class GameObject {
  type: string
  x: number
  y: number
  width: number
  height: number
  vel_x: number
  vel_y: number
  animationData: {
    countToStep: number
    actualStep: number
    flipped: boolean
  }

  // slightly problematic, but there should be no harm
  constructor(contents) {
    Object.assign(this, contents)

    this.x = contents.x * ZOOM_FACTOR
    this.y = contents.y * ZOOM_FACTOR
    this.vel_x = contents.vel_x * ZOOM_FACTOR
    this.vel_y = contents.vel_y * ZOOM_FACTOR
    this.width = contents.width * ZOOM_FACTOR
    this.height = contents.height * ZOOM_FACTOR

    this.animationData = {
      countToStep: 0,
      actualStep: 0,
      flipped: false,
    }
  }
}

// biome-ignore lint/complexity/noBannedTypes: <explanation>
export const Models: { [key: string]: Function } = {
  player,
  tree,
}
