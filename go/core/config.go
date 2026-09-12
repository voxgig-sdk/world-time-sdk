package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "WorldTime",
			"slug": "world-time",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
				"transport": "base",
			},
		},
		"options": map[string]any{
			"base": "https://worldtimeapi.org/api",
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"ipn": map[string]any{},
				"timezone": map[string]any{},
			},
		},
		"entity": map[string]any{
			"ipn": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "abbreviation",
						"short": "The abbreviated name of the timezone",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "client_ip",
						"short": "The IP address of the client",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "datetime",
						"short": "The current datetime in ISO 8601 format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "day_of_week",
						"short": "The day of the week (0-6, where 0 is Sunday)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "day_of_year",
						"short": "The day of the year (1-365/366)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "dst",
						"short": "Whether daylight saving time is currently in effect",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"format": "date-time",
						"name": "dst_from",
						"short": "The datetime when DST starts",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dst_offset",
						"short": "The DST offset in seconds",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"format": "date-time",
						"name": "dst_until",
						"short": "The datetime when DST ends",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "raw_offset",
						"short": "The raw offset from UTC in seconds",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "timezone",
						"short": "The IANA timezone identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "unixtime",
						"short": "The current Unix timestamp",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"format": "date-time",
						"name": "utc_datetime",
						"short": "The current UTC datetime in ISO 8601 format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "utc_offset",
						"short": "The UTC offset in ±HH:MM format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "week_number",
						"short": "The ISO week number of the year",
						"type": "`$INTEGER`",
					},
				},
				"name": "ipn",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "8.8.8.8",
											"kind": "param",
											"name": "ipv4",
											"orig": "ipv4",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/ip/{ipv4}",
								"segments": []any{
									map[string]any{
										"lit": "ip",
									},
									map[string]any{
										"var": "ipv4",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"ipv4",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"ip",
									"{ipv4}",
								},
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/ip",
								"segments": []any{
									map[string]any{
										"lit": "ip",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"ip",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"ip",
						},
					},
				},
			},
			"timezone": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "abbreviation",
						"short": "The abbreviated name of the timezone",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "client_ip",
						"short": "The IP address of the client",
						"type": "`$STRING`",
					},
					map[string]any{
						"format": "date-time",
						"name": "datetime",
						"short": "The current datetime in ISO 8601 format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "day_of_week",
						"short": "The day of the week (0-6, where 0 is Sunday)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "day_of_year",
						"short": "The day of the year (1-365/366)",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "dst",
						"short": "Whether daylight saving time is currently in effect",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"format": "date-time",
						"name": "dst_from",
						"short": "The datetime when DST starts",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dst_offset",
						"short": "The DST offset in seconds",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"format": "date-time",
						"name": "dst_until",
						"short": "The datetime when DST ends",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "id",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "raw_offset",
						"short": "The raw offset from UTC in seconds",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "timezone",
						"short": "The IANA timezone identifier",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "unixtime",
						"short": "The current Unix timestamp",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"format": "date-time",
						"name": "utc_datetime",
						"short": "The current UTC datetime in ISO 8601 format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "utc_offset",
						"short": "The UTC offset in ±HH:MM format",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "week_number",
						"short": "The ISO week number of the year",
						"type": "`$INTEGER`",
					},
				},
				"id": map[string]any{
					"field": "id",
					"name": "id",
				},
				"name": "timezone",
				"op": map[string]any{
					"list": map[string]any{
						"input": "data",
						"name": "list",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/timezone",
								"segments": []any{
									map[string]any{
										"lit": "timezone",
									},
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"timezone",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "America",
											"kind": "param",
											"name": "area",
											"orig": "area",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "New_York",
											"kind": "param",
											"name": "location",
											"orig": "location",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/timezone/{area}/{location}",
								"segments": []any{
									map[string]any{
										"lit": "timezone",
									},
									map[string]any{
										"var": "area",
									},
									map[string]any{
										"var": "location",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"area",
										"location",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"timezone",
									"{area}",
									"{location}",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "Europe",
											"kind": "param",
											"name": "id",
											"orig": "area",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/timezone/{area}",
								"rename": map[string]any{
									"param": map[string]any{
										"area": "id",
									},
								},
								"segments": []any{
									map[string]any{
										"lit": "timezone",
									},
									map[string]any{
										"var": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
								"parts": []any{
									"timezone",
									"{id}",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"timezone",
						},
					},
				},
			},
		},
	}
}

// The plugin definitions the model selected per feature, as []any so a
// feature package can consume them without core naming its types. Empty
// when no active feature declares active plugin groups for this target.
var featurePlugins = map[string][]any{
}

// FeaturePlugins is the definitions list for one feature's chain.
func FeaturePlugins(name string) []any {
	return featurePlugins[name]
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
