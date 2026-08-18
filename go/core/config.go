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
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "client_ip",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "datetime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "day_of_week",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "day_of_year",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "dst",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "dst_from",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dst_offset",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "dst_until",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "raw_offset",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "timezone",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "unixtime",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "utc_datetime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "utc_offset",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "week_number",
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
								"parts": []any{
									"ip",
									"{ipv4}",
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
							},
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/ip",
								"parts": []any{
									"ip",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
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
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "client_ip",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "datetime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "day_of_week",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "day_of_year",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "dst",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "dst_from",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "dst_offset",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "dst_until",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "raw_offset",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "timezone",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "unixtime",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "utc_datetime",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "utc_offset",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "week_number",
						"type": "`$INTEGER`",
					},
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
								"parts": []any{
									"timezone",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
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
								"parts": []any{
									"timezone",
									"{area}",
									"{location}",
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
								"parts": []any{
									"timezone",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"area": "id",
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
