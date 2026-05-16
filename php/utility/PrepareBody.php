<?php
declare(strict_types=1);

// WorldTime SDK utility: prepare_body

class WorldTimePrepareBody
{
    public static function call(WorldTimeContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
