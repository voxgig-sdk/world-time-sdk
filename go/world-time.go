package voxgigworldtimesdk

import (
	"github.com/voxgig-sdk/world-time-sdk/go/core"
	"github.com/voxgig-sdk/world-time-sdk/go/entity"
	"github.com/voxgig-sdk/world-time-sdk/go/feature"
	_ "github.com/voxgig-sdk/world-time-sdk/go/utility"
)

// Type aliases preserve external API.
type WorldTimeSDK = core.WorldTimeSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type WorldTimeEntity = core.WorldTimeEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type WorldTimeError = core.WorldTimeError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewIpnEntityFunc = func(client *core.WorldTimeSDK, entopts map[string]any) core.WorldTimeEntity {
		return entity.NewIpnEntity(client, entopts)
	}
	core.NewTimezoneEntityFunc = func(client *core.WorldTimeSDK, entopts map[string]any) core.WorldTimeEntity {
		return entity.NewTimezoneEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewWorldTimeSDK = core.NewWorldTimeSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewWorldTimeSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *WorldTimeSDK  { return NewWorldTimeSDK(nil) }
func Test() *WorldTimeSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
