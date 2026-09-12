import { Context } from './Context';
declare class WorldTimeError extends Error {
    isWorldTimeError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { WorldTimeError };
