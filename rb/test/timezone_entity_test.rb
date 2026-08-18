# Timezone entity test

require "minitest/autorun"
require "json"
require_relative "../WorldTime_sdk"
require_relative "runner"

class TimezoneEntityTest < Minitest::Test
  def test_create_instance
    testsdk = WorldTimeSDK.test(nil, nil)
    ent = testsdk.Timezone(nil)
    assert !ent.nil?
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "timezone" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = WorldTimeSDK.test(seed, nil)
    seen = base.Timezone(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = WorldTimeConfig.shared_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = WorldTimeSDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.Timezone(nil).stream("list", nil, nil).each do |item|
        if item.is_a?(Array)
          got.concat(item)
        else
          got << item
        end
      end
      assert_equal 3, got.length
    end
  end

  def test_basic_flow
    setup = timezone_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "timezone." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set WORLD_TIME_TEST_TIMEZONE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    timezone_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.timezone")))
    timezone_ref01_data = nil
    if timezone_ref01_data_raw.length > 0
      timezone_ref01_data = Helpers.to_map(timezone_ref01_data_raw[0][1])
    end

    # LIST
    timezone_ref01_ent = client.Timezone(nil)
    timezone_ref01_match = {}

    timezone_ref01_list_result = timezone_ref01_ent.list(timezone_ref01_match, nil)
    assert timezone_ref01_list_result.is_a?(Array)

    # LOAD
    timezone_ref01_match_dt0 = {}
    timezone_ref01_data_dt0_loaded = timezone_ref01_ent.load(timezone_ref01_match_dt0, nil)
    assert !timezone_ref01_data_dt0_loaded.nil?

  end
end

def timezone_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "timezone", "TimezoneTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = WorldTimeSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["timezone01", "timezone02", "timezone03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["WORLD_TIME_TEST_TIMEZONE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "WORLD_TIME_TEST_TIMEZONE_ENTID" => idmap,
    "WORLD_TIME_TEST_LIVE" => "FALSE",
    "WORLD_TIME_TEST_EXPLAIN" => "FALSE",
  })

  idmap_resolved = Helpers.to_map(
    env["WORLD_TIME_TEST_TIMEZONE_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["WORLD_TIME_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
      },
      extra || {},
    ])
    client = WorldTimeSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["WORLD_TIME_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["WORLD_TIME_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
