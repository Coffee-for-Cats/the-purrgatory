import { player } from './player'
import { tree } from './tree'

export class GameObject {
  type: string
  x: number
  y: number
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
