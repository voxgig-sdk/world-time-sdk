# WorldTime PHP SDK Reference

Complete API reference for the WorldTime PHP SDK.


## WorldTimeSDK

### Constructor

```php
require_once __DIR__ . '/worldtime_sdk.php';

$client = new WorldTimeSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `WorldTimeSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = WorldTimeSDK::test();
```


### Instance Methods

#### `Ipn($data = null)`

Create a new `IpnEntity` instance. Pass `null` for no initial data.

#### `Timezone($data = null)`

Create a new `TimezoneEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): WorldTimeUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## IpnEntity

```php
$ipn = $client->Ipn();
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

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Ipn()->load(["ipv4" => "ipv4"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IpnEntity`

Create a new `IpnEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## TimezoneEntity

```php
$timezone = $client->Timezone();
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

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Timezone()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Timezone()->load(["id" => "timezone_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): TimezoneEntity`

Create a new `TimezoneEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new WorldTimeSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

