<?php
declare(strict_types=1);

// WorldTime SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class WorldTimeFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new WorldTimeBaseFeature();
            case "test":
                return new WorldTimeTestFeature();
            default:
                return new WorldTimeBaseFeature();
        }
    }
}
