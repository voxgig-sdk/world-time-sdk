
import { inspect } from 'node:util'

import { WorldTimeEntityBase } from '../WorldTimeEntityBase'

import type {
  WorldTimeSDK,
} from '../WorldTimeSDK'


import type {
  Operation,
  Context,
  Control,
} from '../types'

import type {
  Ipn,
} from '../WorldTimeTypes'

// TODO: needs Entity superclass
class IpnEntity extends WorldTimeEntityBase<Ipn> {

  constructor(client: WorldTimeSDK, entopts: any) {
    super(client, entopts)
    this.name = 'ipn'
    this.name_ = 'ipn'
    this.Name = 'Ipn'
  }


  make(this: IpnEntity) {
    return new IpnEntity(this._client, this.entopts())
  }







}


export {
  IpnEntity
}
