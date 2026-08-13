package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewIpnEntityFunc func(client *WorldTimeSDK, entopts map[string]any) WorldTimeEntity

var NewTimezoneEntityFunc func(client *WorldTimeSDK, entopts map[string]any) WorldTimeEntity

