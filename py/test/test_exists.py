# WorldTime SDK exists test

import pytest
from worldtime_sdk import WorldTimeSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = WorldTimeSDK.test(None, None)
        assert testsdk is not None
