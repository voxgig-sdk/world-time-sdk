package sdktest

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/world-time-sdk/go"
	"github.com/voxgig-sdk/world-time-sdk/go/core"

	vs "github.com/voxgig-sdk/world-time-sdk/go/utility/struct"
)

func TestIpn2Entity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Ipn2(nil)
		if ent == nil {
			t.Fatal("expected non-nil Ipn2Entity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := ipn2BasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "ipn2." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set WORLDTIME_TEST_IPN__ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		ipn2Ref01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.ipn2", setup.data)))
		var ipn2Ref01Data map[string]any
		if len(ipn2Ref01DataRaw) > 0 {
			ipn2Ref01Data = core.ToMapAny(ipn2Ref01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = ipn2Ref01Data

		// LOAD
		ipn2Ref01Ent := client.Ipn2(nil)
		ipn2Ref01MatchDt0 := map[string]any{}
		ipn2Ref01DataDt0Loaded, err := ipn2Ref01Ent.Load(ipn2Ref01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if ipn2Ref01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func ipn2BasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "ipn2", "Ipn2TestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read ipn2 test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse ipn2 test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"ipn201", "ipn202", "ipn203", "ip01", "ip02", "ip03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("WORLDTIME_TEST_IPN__ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"WORLDTIME_TEST_IPN__ENTID": idmap,
		"WORLDTIME_TEST_LIVE":      "FALSE",
		"WORLDTIME_TEST_EXPLAIN":   "FALSE",
		"WORLDTIME_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["WORLDTIME_TEST_IPN__ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}

	if env["WORLDTIME_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["WORLDTIME_APIKEY"],
			},
			extra,
		})
		client = sdk.NewWorldTimeSDK(core.ToMapAny(mergedOpts))
	}

	live := env["WORLDTIME_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["WORLDTIME_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
