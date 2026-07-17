-- WorldTime SDK exists test

local sdk = require("world-time_sdk")

describe("WorldTimeSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
