<?php
declare(strict_types=1);

// WorldTime SDK exists test

require_once __DIR__ . '/../worldtime_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = WorldTimeSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
