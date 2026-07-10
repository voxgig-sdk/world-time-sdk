# WorldTime TypeScript SDK



The TypeScript SDK for the WorldTime API — a type-safe, entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Ipn()` — each with a small set of operations (`list`, `load`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to npm. Install it from the GitHub
release tag (`ts/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/world-time-sdk/releases](https://github.com/voxgig-sdk/world-time-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ts
import { WorldTimeSDK } from '@voxgig-sdk/world-time'

const client = new WorldTimeSDK()
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const ipn2 = await client.Ipn2().load()
  console.log(ipn2)
} catch (err) {
  console.error('load failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```ts
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```ts
const client = WorldTimeSDK.test()

const ipn2 = await client.Ipn2().load()
// ipn2 is a bare entity populated with mock response data
console.log(ipn2)
```

You can also use the instance method:

```ts
const client = new WorldTimeSDK()
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```ts
const entity = client.Ipn2()

// First call runs the operation and stores its result
await entity.load()

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```ts
const logger = {
  hooks: {
    PreRequest: (ctx: any) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx: any) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new WorldTimeSDK({
  extend: [logger],
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
WORLD_TIME_TEST_LIVE=TRUE
```

Then run:

```bash
cd ts && npm test
```


## Reference

### WorldTimeSDK

#### Constructor

```ts
new WorldTimeSDK(options?: {
  base?: string
  prefix?: string
  suffix?: string
  feature?: Record<string, { active: boolean }>
  extend?: Feature[]
})
```

| Option | Type | Description |
| --- | --- | --- |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Ipn(data?)` | `IpnEntity` | Create an Ipn entity instance. |
| `Ipn2(data?)` | `Ipn2Entity` | Create an Ipn2 entity instance. |
| `Timezone(data?)` | `TimezoneEntity` | Create a Timezone entity instance. |
| `tester(testopts?, sdkopts?)` | `WorldTimeSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `WorldTimeSDK.test(testopts?, sdkopts?)` | `WorldTimeSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): WorldTimeSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load` resolves to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```ts
{
  ok: boolean
  status: number
  headers: object
  data: any
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```ts
{
  url: string
  method: string
  headers: Record<string, string>
  body?: any
}
```

### Entities

#### Ipn

| Field | Description |
| --- | --- |

Operations: .

API path: ``

#### Ipn2

| Field | Description |
| --- | --- |
| `abbreviation` |  |
| `client_ip` |  |
| `datetime` |  |
| `day_of_week` |  |
| `day_of_year` |  |
| `dst` |  |
| `dst_from` |  |
| `dst_offset` |  |
| `dst_until` |  |
| `raw_offset` |  |
| `timezone` |  |
| `unixtime` |  |
| `utc_datetime` |  |
| `utc_offset` |  |
| `week_number` |  |

Operations: load.

API path: `/ip/{ipv4}`

#### Timezone

| Field | Description |
| --- | --- |
| `abbreviation` |  |
| `client_ip` |  |
| `datetime` |  |
| `day_of_week` |  |
| `day_of_year` |  |
| `dst` |  |
| `dst_from` |  |
| `dst_offset` |  |
| `dst_until` |  |
| `raw_offset` |  |
| `timezone` |  |
| `unixtime` |  |
| `utc_datetime` |  |
| `utc_offset` |  |
| `week_number` |  |

Operations: list, load.

API path: `/timezone`



## Entities


### Ipn

Create an instance: `const ipn = client.Ipn()`


### Ipn2

Create an instance: `const ipn2 = client.Ipn2()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `abbreviation` | `string` |  |
| `client_ip` | `string` |  |
| `datetime` | `string` |  |
| `day_of_week` | `number` |  |
| `day_of_year` | `number` |  |
| `dst` | `boolean` |  |
| `dst_from` | `string` |  |
| `dst_offset` | `number` |  |
| `dst_until` | `string` |  |
| `raw_offset` | `number` |  |
| `timezone` | `string` |  |
| `unixtime` | `number` |  |
| `utc_datetime` | `string` |  |
| `utc_offset` | `string` |  |
| `week_number` | `number` |  |

#### Example: Load

```ts
const ipn2 = await client.Ipn2().load()
```


### Timezone

Create an instance: `const timezone = client.Timezone()`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `abbreviation` | `string` |  |
| `client_ip` | `string` |  |
| `datetime` | `string` |  |
| `day_of_week` | `number` |  |
| `day_of_year` | `number` |  |
| `dst` | `boolean` |  |
| `dst_from` | `string` |  |
| `dst_offset` | `number` |  |
| `dst_until` | `string` |  |
| `raw_offset` | `number` |  |
| `timezone` | `string` |  |
| `unixtime` | `number` |  |
| `utc_datetime` | `string` |  |
| `utc_offset` | `string` |  |
| `week_number` | `number` |  |

#### Example: Load

```ts
const timezone = await client.Timezone().load({ id: 'timezone_id' })
```

#### Example: List

```ts
const timezones = await client.Timezone().list()
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
world-time/
├── src/
│   ├── WorldTimeSDK.ts        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
├── test/                   # Test suites
└── dist/                   # Compiled output
```

Import the SDK from the package root:

```ts
import { WorldTimeSDK } from '@voxgig-sdk/world-time'
```

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const ipn2 = client.Ipn2()
await ipn2.load()

// ipn2.data() now returns the ipn2 data from the last `load`
// ipn2.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
