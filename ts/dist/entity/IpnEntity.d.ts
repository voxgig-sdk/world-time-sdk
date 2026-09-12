import { WorldTimeEntityBase } from '../WorldTimeEntityBase';
import type { WorldTimeSDK } from '../WorldTimeSDK';
import type { Control } from '../types';
import type { Ipn, IpnLoadMatch } from '../WorldTimeTypes';
declare class IpnEntity extends WorldTimeEntityBase<Ipn> {
    constructor(client: WorldTimeSDK, entopts: any);
    make(this: IpnEntity): IpnEntity;
    load(this: any, reqmatch?: IpnLoadMatch, ctrl?: Control): Promise<IpnEntity>;
}
export { IpnEntity };
