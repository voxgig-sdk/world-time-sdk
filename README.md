# WorldTime SDK

Get the current time and related timezone data via a simple JSON/plain-text API

> TypeScript, Python, PHP, Golang, Ruby, Lua SDKs, a CLI, an interactive REPL, and an MCP server for AI agents — all generated from one OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).

## About World Time API

[World Time API](https://worldtimeapi.org/) is a free, public service that returns the current time and a handful of useful timezone metadata for a requested location. It is operated as an open community project and serves responses as JSON or plain text.

What you typically get from a response:
- Current local date/time and UTC offset for the requested timezone
- Day-of-week, day-of-year, and week number
- Whether daylight saving time is currently in effect
- The IANA timezone name and abbreviation

Operational notes:
- CORS is disabled, so calls must be made server-side or via a proxy
- No API key or authentication is required
- Reliability varies by endpoint; the per-timezone endpoint is the most consistently available

## Try it

**TypeScript**
```bash
npm install world-time
```

**Python**
```bash
pip install world-time-sdk
```

**PHP**
```bash
composer require voxgig/world-time-sdk
```

**Golang**
```bash
go get github.com/voxgig-sdk/world-time-sdk/go
```

**Ruby**
```bash
gem install world-time-sdk
```

**Lua**
```bash
luarocks install world-time-sdk
```

## 30-second quickstart

### TypeScript

```ts
import { WorldTimeSDK } from 'world-time'

const client = new WorldTimeSDK({})

```

See the [TypeScript README](ts/README.md) for the
full guide, or scroll down for the same example in other languages.

## What's in the box

| Surface | Use it for | Path |
| --- | --- | --- |
| **SDK** (TypeScript, Python, PHP, Golang, Ruby, Lua) | App integration | `ts/` `py/` `php/` `go/` `rb/` `lua/` |
| **CLI** | Scripts, CI, ops, one-off API calls | `go-cli/` |
| **MCP server** | AI agents (Claude, Cursor, Cline) | `go-mcp/` |

## Use it from an AI agent (MCP)

The generated MCP server exposes every operation in this SDK as an
[MCP](https://modelcontextprotocol.io) tool that Claude, Cursor or Cline
can call directly. Build and register it:

```bash
cd go-mcp && go build -o world-time-mcp .
```

Then add it to your agent's MCP config (Claude Desktop, Cursor, etc.):

```json
{
  "mcpServers": {
    "world-time": {
      "command": "/abs/path/to/world-time-mcp"
    }
  }
}
```

## Entities

The API exposes 3 entities:

| Entity | Description | API path |
| --- | --- | --- |
| **Ipn** | Lookup of timezone information based on the caller's IP address (or a supplied IPv4), backed by `/ip` and `/ip/{ipv4}`. | `` |
| **Ipn2** | Alternate IP-based timezone lookup grouping exposed by this SDK, covering the same `/ip` family of endpoints. | `/ip/{ipv4}` |
| **Timezone** | Current time and DST/offset metadata for a named IANA timezone, served from paths like `/timezone/{area}` and `/timezone/{area}/{location}` (e.g. `/timezone/America/Argentina/Salta`). | `/timezone` |

Each entity supports the following operations where available: **load**,
**list**, **create**, **update**, and **remove**.

## Quickstart in other languages

### Python

```python
from worldtime_sdk import WorldTimeSDK

client = WorldTimeSDK({})

```

### PHP

```php
<?php
require_once 'worldtime_sdk.php';

$client = new WorldTimeSDK([]);

```

### Golang

```go
import sdk "github.com/voxgig-sdk/world-time-sdk/go"

client := sdk.NewWorldTimeSDK(map[string]any{})

```

### Ruby

```ruby
require_relative "WorldTime_sdk"

client = WorldTimeSDK.new({})

```

### Lua

```lua
local sdk = require("world-time_sdk")

local client = sdk.new({})

```

## Unit testing in offline mode

Every SDK ships a test mode that swaps the HTTP transport for an
in-memory mock, so unit tests run offline.

### TypeScript

```ts
const client = WorldTimeSDK.test()
const result = await client.Ipn().load({ id: 'test01' })
// result.ok === true, result.data contains mock data
```

### Python

```python
client = WorldTimeSDK.test(None, None)
result, err = client.Ipn(None).load(
    {"id": "test01"}, None
)
```

### PHP

```php
$client = WorldTimeSDK::test(null, null);
[$result, $err] = $client->Ipn(null)->load(
    ["id" => "test01"], null
);
```

### Golang

```go
client := sdk.TestSDK(nil, nil)
result, err := client.Ipn(nil).Load(
    map[string]any{"id": "test01"}, nil,
)
```

### Ruby

```ruby
client = WorldTimeSDK.test(nil, nil)
result, err = client.Ipn(nil).load(
  { "id" => "test01" }, nil
)
```

### Lua

```lua
local client = sdk.test(nil, nil)
local result, err = client:Ipn(nil):load(
  { id = "test01" }, nil
)
```

## How it works

Every SDK call runs the same five-stage pipeline:

1. **Point** — resolve the API endpoint from the operation definition.
2. **Spec** — build the HTTP specification (URL, method, headers, body).
3. **Request** — send the HTTP request.
4. **Response** — receive and parse the response.
5. **Result** — extract the result data for the caller.

A feature hook fires at each stage (e.g. `PrePoint`, `PreSpec`,
`PreRequest`), so features can inspect or modify the pipeline without
forking the SDK.

### Features

| Feature | Purpose |
| --- | --- |
| **TestFeature** | In-memory mock transport for testing without a live server |

Pass custom features via the `extend` option at construction time.

### Direct and Prepare

For endpoints the entity model doesn't cover, use the low-level methods:

- **`direct(fetchargs)`** — build and send an HTTP request in one step.
- **`prepare(fetchargs)`** — build the request without sending it.

Both accept a map with `path`, `method`, `params`, `query`,
`headers`, and `body`. See the [How-to guides](#how-to-guides) below.

## How-to guides

### Make a direct API call

When the entity interface does not cover an endpoint, use `direct`:

**TypeScript:**
```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})
console.log(result.data)
```

**Python:**
```python
result, err = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})
```

**PHP:**
```php
[$result, $err] = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);
```

**Go:**
```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
```

**Ruby:**
```ruby
result, err = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})
```

**Lua:**
```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
```

## Per-language documentation

- [TypeScript](ts/README.md)
- [Python](py/README.md)
- [PHP](php/README.md)
- [Golang](go/README.md)
- [Ruby](rb/README.md)
- [Lua](lua/README.md)

## Using the World Time API

- Upstream: [https://worldtimeapi.org/](https://worldtimeapi.org/)
- API docs: [https://worldtimeapi.org/pages/examples](https://worldtimeapi.org/pages/examples)

- Released into the Public Domain by the project authors
- No attribution required, but a courtesy link back is appreciated
- The service is provided as-is with no uptime guarantee

---

Generated from the World Time API OpenAPI spec by [@voxgig/sdkgen](https://github.com/voxgig/sdkgen).
