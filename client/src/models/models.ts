import { player } from './player'
import { tree } from './tree'

export class GameObject {
  type: string
  x: number
  y: number
  vel_x: number
  vel_y: number
  animationData:  undefined | {
    countToStep: number,
    actualStep: number
  }
}

// biome-ignore lint/complexity/noBannedTypes: <explanation>
export const Models: { [key: string]: Function } = {
  player,
  tree,
}
