# WorldTime TypeScript SDK Reference

Complete API reference for the WorldTime TypeScript SDK.


## WorldTimeSDK

### Constructor

```ts
new WorldTimeSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `WorldTimeSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = WorldTimeSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `WorldTimeSDK` instance in test mode.


### Instance Methods

#### `Ipn(data?: object)`

Create a new `Ipn` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IpnEntity` instance.

#### `Timezone(data?: object)`

Create a new `Timezone` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `TimezoneEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `WorldTimeSDK.test()`.

**Returns:** `WorldTimeSDK` instance in test mode.


---

## IpnEntity

```ts
const ipn = client.Ipn()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `string` | No |  |
| `client_ip` | `string` | No |  |
| `datetime` | `string` | No |  |
| `day_of_week` | `number` | No |  |
| `day_of_year` | `number` | No |  |
| `dst` | `boolean` | No |  |
| `dst_from` | `string` | No |  |
| `dst_offset` | `number` | No |  |
| `dst_until` | `string` | No |  |
| `raw_offset` | `number` | No |  |
| `timezone` | `string` | No |  |
| `unixtime` | `number` | No |  |
| `utc_datetime` | `string` | No |  |
| `utc_offset` | `string` | No |  |
| `week_number` | `number` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Ipn().load()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IpnEntity` instance with the same client and
options.

#### `client()`

Return the parent `WorldTimeSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## TimezoneEntity

```ts
const timezone = client.Timezone()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `abbreviation` | `string` | No |  |
| `client_ip` | `string` | No |  |
| `datetime` | `string` | No |  |
| `day_of_week` | `number` | No |  |
| `day_of_year` | `number` | No |  |
| `dst` | `boolean` | No |  |
| `dst_from` | `string` | No |  |
| `dst_offset` | `number` | No |  |
| `dst_until` | `string` | No |  |
| `raw_offset` | `number` | No |  |
| `timezone` | `string` | No |  |
| `unixtime` | `number` | No |  |
| `utc_datetime` | `string` | No |  |
| `utc_offset` | `string` | No |  |
| `week_number` | `number` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Timezone().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Timezone().load({ id: 'timezone_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `TimezoneEntity` instance with the same client and
options.

#### `client()`

Return the parent `WorldTimeSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new WorldTimeSDK({
  feature: {
    test: { active: true },
  }
})
```

