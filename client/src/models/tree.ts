import { Movable, PaintObject, Source } from "../painting/animations";
import { GameObject } from "./models";

const src = Source("tree1")
export function tree(tree: GameObject) {
  PaintObject(tree as Movable, src)
}