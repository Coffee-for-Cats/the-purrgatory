import { ResetAnimationFrame } from './animate'

// biome-ignore lint/suspicious/noExplicitAny: <explanation>
export let GameMap: { [index: string]: any } = {}

export function Listen(Socket: WebSocket) {
  Socket.onmessage = (ev) => {
    GameMap = JSON.parse(ev.data)

    ResetAnimationFrame()
  }
}
