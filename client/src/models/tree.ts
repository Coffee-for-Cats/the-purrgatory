import { PaintObject, Source } from '../painting/painting'
import type { GameObject } from './models'

const src = Source('tree1')
export function tree(tree: GameObject) {
  PaintObject(tree, src)
}
