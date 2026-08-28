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
| `abbreviation` | `string` | No | The abbreviated name of the timezone |
| `client_ip` | `string` | No | The IP address of the client |
| `datetime` | `string` | No | The current datetime in ISO 8601 format |
| `day_of_week` | `number` | No | The day of the week (0-6, where 0 is Sunday) |
| `day_of_year` | `number` | No | The day of the year (1-365/366) |
| `dst` | `boolean` | No | Whether daylight saving time is currently in effect |
| `dst_from` | `string` | No | The datetime when DST starts |
| `dst_offset` | `number` | No | The DST offset in seconds |
| `dst_until` | `string` | No | The datetime when DST ends |
| `raw_offset` | `number` | No | The raw offset from UTC in seconds |
| `timezone` | `string` | No | The IANA timezone identifier |
| `unixtime` | `number` | No | The current Unix timestamp |
| `utc_datetime` | `string` | No | The current UTC datetime in ISO 8601 format |
| `utc_offset` | `string` | No | The UTC offset in ±HH:MM format |
| `week_number` | `number` | No | The ISO week number of the year |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Ipn().load({ ipv4: 'ipv4' })
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
| `abbreviation` | `string` | No | The abbreviated name of the timezone |
| `client_ip` | `string` | No | The IP address of the client |
| `datetime` | `string` | No | The current datetime in ISO 8601 format |
| `day_of_week` | `number` | No | The day of the week (0-6, where 0 is Sunday) |
| `day_of_year` | `number` | No | The day of the year (1-365/366) |
| `dst` | `boolean` | No | Whether daylight saving time is currently in effect |
| `dst_from` | `string` | No | The datetime when DST starts |
| `dst_offset` | `number` | No | The DST offset in seconds |
| `dst_until` | `string` | No | The datetime when DST ends |
| `id` | `string` | No |  |
| `raw_offset` | `number` | No | The raw offset from UTC in seconds |
| `timezone` | `string` | No | The IANA timezone identifier |
| `unixtime` | `number` | No | The current Unix timestamp |
| `utc_datetime` | `string` | No | The current UTC datetime in ISO 8601 format |
| `utc_offset` | `string` | No | The UTC offset in ±HH:MM format |
| `week_number` | `number` | No | The ISO week number of the year |

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

