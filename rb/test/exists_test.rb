# WorldTime SDK exists test

require "minitest/autorun"
require_relative "../WorldTime_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = WorldTimeSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
