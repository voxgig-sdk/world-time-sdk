<?php
declare(strict_types=1);

// WorldTime SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class WorldTimeMakeContext
{
    public static function call(array $ctxmap, ?WorldTimeContext $basectx): WorldTimeContext
    {
        return new WorldTimeContext($ctxmap, $basectx);
    }
}
