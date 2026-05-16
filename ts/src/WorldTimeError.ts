
import { Context } from './Context'


class WorldTimeError extends Error {

  isWorldTimeError = true

  sdk = 'WorldTime'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  WorldTimeError
}

