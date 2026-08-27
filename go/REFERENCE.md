# WorldTime Golang SDK Reference

Complete API reference for the WorldTime Golang SDK.


## WorldTimeSDK

### Constructor

```go
func NewWorldTimeSDK(options map[string]any) *WorldTimeSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *WorldTimeSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *WorldTimeSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Ipn(data map[string]any) WorldTimeEntity`

Create a new `Ipn` entity instance. Pass `nil` for no initial data.

#### `Timezone(data map[string]any) WorldTimeEntity`

Create a new `Timezone` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## IpnEntity

```go
ipn := client.Ipn(nil)
fmt.Println(ipn.GetName()) // "ipn"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `string` | No | The abbreviated name of the timezone |
| `client_ip` | `string` | No | The IP address of the client |
| `datetime` | `string` | No | The current datetime in ISO 8601 format |
| `day_of_week` | `int` | No | The day of the week (0-6, where 0 is Sunday) |
| `day_of_year` | `int` | No | The day of the year (1-365/366) |
| `dst` | `bool` | No | Whether daylight saving time is currently in effect |
| `dst_from` | `string` | No | The datetime when DST starts |
| `dst_offset` | `int` | No | The DST offset in seconds |
| `dst_until` | `string` | No | The datetime when DST ends |
| `raw_offset` | `int` | No | The raw offset from UTC in seconds |
| `timezone` | `string` | No | The IANA timezone identifier |
| `unixtime` | `int` | No | The current Unix timestamp |
| `utc_datetime` | `string` | No | The current UTC datetime in ISO 8601 format |
| `utc_offset` | `string` | No | The UTC offset in ±HH:MM format |
| `week_number` | `int` | No | The ISO week number of the year |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Ipn(nil).Load(map[string]any{"ipv4": "ipv4"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IpnEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## TimezoneEntity

```go
timezone := client.Timezone(nil)
fmt.Println(timezone.GetName()) // "timezone"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `string` | No | The abbreviated name of the timezone |
| `client_ip` | `string` | No | The IP address of the client |
| `datetime` | `string` | No | The current datetime in ISO 8601 format |
| `day_of_week` | `int` | No | The day of the week (0-6, where 0 is Sunday) |
| `day_of_year` | `int` | No | The day of the year (1-365/366) |
| `dst` | `bool` | No | Whether daylight saving time is currently in effect |
| `dst_from` | `string` | No | The datetime when DST starts |
| `dst_offset` | `int` | No | The DST offset in seconds |
| `dst_until` | `string` | No | The datetime when DST ends |
| `id` | `string` | No |  |
| `raw_offset` | `int` | No | The raw offset from UTC in seconds |
| `timezone` | `string` | No | The IANA timezone identifier |
| `unixtime` | `int` | No | The current Unix timestamp |
| `utc_datetime` | `string` | No | The current UTC datetime in ISO 8601 format |
| `utc_offset` | `string` | No | The UTC offset in ±HH:MM format |
| `week_number` | `int` | No | The ISO week number of the year |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Timezone(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Timezone(nil).Load(map[string]any{"id": "timezone_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `TimezoneEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewWorldTimeSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

