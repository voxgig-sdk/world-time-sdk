import { WorldTimeEntityBase } from '../WorldTimeEntityBase';
import type { WorldTimeSDK } from '../WorldTimeSDK';
import type { Control } from '../types';
import type { Timezone, TimezoneLoadMatch, TimezoneListMatch } from '../WorldTimeTypes';
declare class TimezoneEntity extends WorldTimeEntityBase<Timezone> {
    constructor(client: WorldTimeSDK, entopts: any);
    make(this: TimezoneEntity): TimezoneEntity;
    load(this: any, reqmatch?: TimezoneLoadMatch, ctrl?: Control): Promise<TimezoneEntity>;
    list(this: any, reqmatch?: TimezoneListMatch, ctrl?: Control): Promise<TimezoneEntity[]>;
}
export { TimezoneEntity };
