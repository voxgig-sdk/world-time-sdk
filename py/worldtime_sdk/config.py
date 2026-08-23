# WorldTime SDK configuration


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "WorldTime",
            "slug": "world-time",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
      },
        },
        "options": {
            "base": "https://worldtimeapi.org/api",
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "ipn": {},
                "timezone": {},
            },
        },
        "entity": {
      "ipn": {
        "fields": [
          {
            "name": "abbreviation",
            "short": "The abbreviated name of the timezone",
            "type": "`$STRING`",
          },
          {
            "name": "client_ip",
            "short": "The IP address of the client",
            "type": "`$STRING`",
          },
          {
            "name": "datetime",
            "short": "The current datetime in ISO 8601 format",
            "type": "`$STRING`",
          },
          {
            "name": "day_of_week",
            "short": "The day of the week (0-6, where 0 is Sunday)",
            "type": "`$INTEGER`",
          },
          {
            "name": "day_of_year",
            "short": "The day of the year (1-365/366)",
            "type": "`$INTEGER`",
          },
          {
            "name": "dst",
            "short": "Whether daylight saving time is currently in effect",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "dst_from",
            "short": "The datetime when DST starts",
            "type": "`$STRING`",
          },
          {
            "name": "dst_offset",
            "short": "The DST offset in seconds",
            "type": "`$INTEGER`",
          },
          {
            "name": "dst_until",
            "short": "The datetime when DST ends",
            "type": "`$STRING`",
          },
          {
            "name": "raw_offset",
            "short": "The raw offset from UTC in seconds",
            "type": "`$INTEGER`",
          },
          {
            "name": "timezone",
            "short": "The IANA timezone identifier",
            "type": "`$STRING`",
          },
          {
            "name": "unixtime",
            "short": "The current Unix timestamp",
            "type": "`$INTEGER`",
          },
          {
            "name": "utc_datetime",
            "short": "The current UTC datetime in ISO 8601 format",
            "type": "`$STRING`",
          },
          {
            "name": "utc_offset",
            "short": "The UTC offset in ±HH:MM format",
            "type": "`$STRING`",
          },
          {
            "name": "week_number",
            "short": "The ISO week number of the year",
            "type": "`$INTEGER`",
          },
        ],
        "name": "ipn",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "8.8.8.8",
                      "kind": "param",
                      "name": "ipv4",
                      "orig": "ipv4",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/ip/{ipv4}",
                "parts": [
                  "ip",
                  "{ipv4}",
                ],
                "select": {
                  "exist": [
                    "ipv4",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
              {
                "args": {},
                "kind": "http",
                "method": "GET",
                "orig": "/ip",
                "parts": [
                  "ip",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "ip",
            ],
          ],
        },
      },
      "timezone": {
        "fields": [
          {
            "name": "abbreviation",
            "short": "The abbreviated name of the timezone",
            "type": "`$STRING`",
          },
          {
            "name": "client_ip",
            "short": "The IP address of the client",
            "type": "`$STRING`",
          },
          {
            "name": "datetime",
            "short": "The current datetime in ISO 8601 format",
            "type": "`$STRING`",
          },
          {
            "name": "day_of_week",
            "short": "The day of the week (0-6, where 0 is Sunday)",
            "type": "`$INTEGER`",
          },
          {
            "name": "day_of_year",
            "short": "The day of the year (1-365/366)",
            "type": "`$INTEGER`",
          },
          {
            "name": "dst",
            "short": "Whether daylight saving time is currently in effect",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "dst_from",
            "short": "The datetime when DST starts",
            "type": "`$STRING`",
          },
          {
            "name": "dst_offset",
            "short": "The DST offset in seconds",
            "type": "`$INTEGER`",
          },
          {
            "name": "dst_until",
            "short": "The datetime when DST ends",
            "type": "`$STRING`",
          },
          {
            "name": "raw_offset",
            "short": "The raw offset from UTC in seconds",
            "type": "`$INTEGER`",
          },
          {
            "name": "timezone",
            "short": "The IANA timezone identifier",
            "type": "`$STRING`",
          },
          {
            "name": "unixtime",
            "short": "The current Unix timestamp",
            "type": "`$INTEGER`",
          },
          {
            "name": "utc_datetime",
            "short": "The current UTC datetime in ISO 8601 format",
            "type": "`$STRING`",
          },
          {
            "name": "utc_offset",
            "short": "The UTC offset in ±HH:MM format",
            "type": "`$STRING`",
          },
          {
            "name": "week_number",
            "short": "The ISO week number of the year",
            "type": "`$INTEGER`",
          },
        ],
        "name": "timezone",
        "op": {
          "list": {
            "input": "data",
            "name": "list",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "GET",
                "orig": "/timezone",
                "parts": [
                  "timezone",
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": "America",
                      "kind": "param",
                      "name": "area",
                      "orig": "area",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "New_York",
                      "kind": "param",
                      "name": "location",
                      "orig": "location",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/timezone/{area}/{location}",
                "parts": [
                  "timezone",
                  "{area}",
                  "{location}",
                ],
                "select": {
                  "exist": [
                    "area",
                    "location",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
              {
                "args": {
                  "params": [
                    {
                      "example": "Europe",
                      "kind": "param",
                      "name": "id",
                      "orig": "area",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/timezone/{area}",
                "parts": [
                  "timezone",
                  "{id}",
                ],
                "rename": {
                  "param": {
                    "area": "id",
                  },
                },
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "timezone",
            ],
          ],
        },
      },
    },
    }
