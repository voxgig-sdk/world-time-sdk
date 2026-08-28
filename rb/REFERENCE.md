# WorldTime Ruby SDK Reference

Complete API reference for the WorldTime Ruby SDK.


## WorldTimeSDK

### Constructor

```ruby
require_relative 'WorldTime_sdk'

client = WorldTimeSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `WorldTimeSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = WorldTimeSDK.test
```


### Instance Methods

#### `Ipn(data = nil)`

Create a new `Ipn` entity instance. Pass `nil` for no initial data.

#### `Timezone(data = nil)`

Create a new `Timezone` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## IpnEntity

```ruby
ipn = client.Ipn
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `String` | No | The abbreviated name of the timezone |
| `client_ip` | `String` | No | The IP address of the client |
| `datetime` | `String` | No | The current datetime in ISO 8601 format |
| `day_of_week` | `Integer` | No | The day of the week (0-6, where 0 is Sunday) |
| `day_of_year` | `Integer` | No | The day of the year (1-365/366) |
| `dst` | `Boolean` | No | Whether daylight saving time is currently in effect |
| `dst_from` | `String` | No | The datetime when DST starts |
| `dst_offset` | `Integer` | No | The DST offset in seconds |
| `dst_until` | `String` | No | The datetime when DST ends |
| `raw_offset` | `Integer` | No | The raw offset from UTC in seconds |
| `timezone` | `String` | No | The IANA timezone identifier |
| `unixtime` | `Integer` | No | The current Unix timestamp |
| `utc_datetime` | `String` | No | The current UTC datetime in ISO 8601 format |
| `utc_offset` | `String` | No | The UTC offset in ±HH:MM format |
| `week_number` | `Integer` | No | The ISO week number of the year |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Ipn.load({ "ipv4" => "ipv4" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IpnEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## TimezoneEntity

```ruby
timezone = client.Timezone
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `String` | No | The abbreviated name of the timezone |
| `client_ip` | `String` | No | The IP address of the client |
| `datetime` | `String` | No | The current datetime in ISO 8601 format |
| `day_of_week` | `Integer` | No | The day of the week (0-6, where 0 is Sunday) |
| `day_of_year` | `Integer` | No | The day of the year (1-365/366) |
| `dst` | `Boolean` | No | Whether daylight saving time is currently in effect |
| `dst_from` | `String` | No | The datetime when DST starts |
| `dst_offset` | `Integer` | No | The DST offset in seconds |
| `dst_until` | `String` | No | The datetime when DST ends |
| `id` | `String` | No |  |
| `raw_offset` | `Integer` | No | The raw offset from UTC in seconds |
| `timezone` | `String` | No | The IANA timezone identifier |
| `unixtime` | `Integer` | No | The current Unix timestamp |
| `utc_datetime` | `String` | No | The current UTC datetime in ISO 8601 format |
| `utc_offset` | `String` | No | The UTC offset in ±HH:MM format |
| `week_number` | `Integer` | No | The ISO week number of the year |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Timezone.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Timezone.load({ "id" => "timezone_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `TimezoneEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = WorldTimeSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```


### Configuring features

Each feature is inactive until switched on, and an SDK with no feature
configured does no feature work at all. Every option below keeps its default
unless you name it.

The array form of \`feature\` is significant: several features wrap the
transport, and the order you list them in is the order they nest.

#### `test`

In-memory mock transport for testing without a live server.

**Configuration**

| Option | Default |
|---|---|
| `active` | `false` |

Options above are those the model carries a default for. A feature may
also accept callback options — a `sink` to receive each record, for
instance — which have no default and are covered in the full feature
reference.

**Usage**

Set `feature.test.active` to true in the client options, and override any option above in the same entry. Every option keeps
its default unless you name it.

**Considerations**

- Attaches to pipeline hooks, not the transport, so activation order does
  not change what it observes.
- Installs the BASE transport that the wrapping features wrap, so it must be
  activated before them.
- Inactive by default: leaving it out costs nothing at runtime.

