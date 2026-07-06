# WorldTime Python SDK Reference

Complete API reference for the WorldTime Python SDK.


## WorldTimeSDK

### Constructor

```python
from worldtime_sdk import WorldTimeSDK

client = WorldTimeSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `WorldTimeSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = WorldTimeSDK.test()
```


### Instance Methods

#### `Ipn(data=None)`

Create a new `IpnEntity` instance. Pass `None` for no initial data.

#### `Ipn2(data=None)`

Create a new `Ipn2Entity` instance. Pass `None` for no initial data.

#### `Timezone(data=None)`

Create a new `TimezoneEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## IpnEntity

```python
ipn = client.Ipn()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IpnEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Ipn2Entity

```python
ipn2 = client.Ipn2()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `str` | No |  |
| `client_ip` | `str` | No |  |
| `datetime` | `str` | No |  |
| `day_of_week` | `int` | No |  |
| `day_of_year` | `int` | No |  |
| `dst` | `bool` | No |  |
| `dst_from` | `str` | No |  |
| `dst_offset` | `int` | No |  |
| `dst_until` | `str` | No |  |
| `raw_offset` | `int` | No |  |
| `timezone` | `str` | No |  |
| `unixtime` | `int` | No |  |
| `utc_datetime` | `str` | No |  |
| `utc_offset` | `str` | No |  |
| `week_number` | `int` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Ipn2().load()
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `Ipn2Entity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## TimezoneEntity

```python
timezone = client.Timezone()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `str` | No |  |
| `client_ip` | `str` | No |  |
| `datetime` | `str` | No |  |
| `day_of_week` | `int` | No |  |
| `day_of_year` | `int` | No |  |
| `dst` | `bool` | No |  |
| `dst_from` | `str` | No |  |
| `dst_offset` | `int` | No |  |
| `dst_until` | `str` | No |  |
| `raw_offset` | `int` | No |  |
| `timezone` | `str` | No |  |
| `unixtime` | `int` | No |  |
| `utc_datetime` | `str` | No |  |
| `utc_offset` | `str` | No |  |
| `week_number` | `int` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Timezone().list()
for timezone in results:
    print(timezone)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Timezone().load({"id": "timezone_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `TimezoneEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = WorldTimeSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

