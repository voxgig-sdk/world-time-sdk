<?php
declare(strict_types=1);

// WorldTime SDK utility: feature_add

class WorldTimeFeatureAdd
{
    public static function call(WorldTimeContext $ctx, mixed $f): void
    {
        $ctx->client->features[] = $f;
    }
}
