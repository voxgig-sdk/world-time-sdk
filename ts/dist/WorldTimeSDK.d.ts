import { IpnEntity } from './entity/IpnEntity';
import { TimezoneEntity } from './entity/TimezoneEntity';
export type * from './WorldTimeTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { WorldTimeEntityBase } from './WorldTimeEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class WorldTimeSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    Ipn(entopts?: Record<string, any>): IpnEntity;
    Timezone(entopts?: Record<string, any>): TimezoneEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): WorldTimeSDK;
    tester(testopts?: any, sdkopts?: any): WorldTimeSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof WorldTimeSDK;
export { stdutil, config, BaseFeature, WorldTimeEntityBase, WorldTimeSDK, SDK, };
