import { player } from './player'
import { tree } from './tree'

export interface GameObject {
  type: string
}

export const Models = {
  player,
  tree,
}
