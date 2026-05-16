# WorldTime SDK configuration


def make_config():
    return {
        "main": {
            "name": "WorldTime",
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
            "auth": {
                "prefix": "Bearer",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "ipn": {},
                "ipn2": {},
                "timezone": {},
            },
        },
        "entity": {
      "ipn": {
        "fields": [],
        "name": "ipn",
        "op": {},
        "relations": {
          "ancestors": [],
        },
      },
      "ipn2": {
        "fields": [
          {
            "name": "abbreviation",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 0,
          },
          {
            "name": "client_ip",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 1,
          },
          {
            "name": "datetime",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 2,
          },
          {
            "name": "day_of_week",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 3,
          },
          {
            "name": "day_of_year",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 4,
          },
          {
            "name": "dst",
            "req": False,
            "type": "`$BOOLEAN`",
            "active": True,
            "index$": 5,
          },
          {
            "name": "dst_from",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 6,
          },
          {
            "name": "dst_offset",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 7,
          },
          {
            "name": "dst_until",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 8,
          },
          {
            "name": "raw_offset",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 9,
          },
          {
            "name": "timezone",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 10,
          },
          {
            "name": "unixtime",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 11,
          },
          {
            "name": "utc_datetime",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 12,
          },
          {
            "name": "utc_offset",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 13,
          },
          {
            "name": "week_number",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 14,
          },
        ],
        "name": "ipn2",
        "op": {
          "load": {
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
                      "active": True,
                    },
                  ],
                },
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
                "active": True,
                "index$": 0,
              },
              {
                "method": "GET",
                "orig": "/ip",
                "parts": [
                  "ip",
                ],
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "active": True,
                "args": {},
                "select": {},
                "index$": 1,
              },
            ],
            "input": "data",
            "key$": "load",
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
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 0,
          },
          {
            "name": "client_ip",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 1,
          },
          {
            "name": "datetime",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 2,
          },
          {
            "name": "day_of_week",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 3,
          },
          {
            "name": "day_of_year",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 4,
          },
          {
            "name": "dst",
            "req": False,
            "type": "`$BOOLEAN`",
            "active": True,
            "index$": 5,
          },
          {
            "name": "dst_from",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 6,
          },
          {
            "name": "dst_offset",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 7,
          },
          {
            "name": "dst_until",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 8,
          },
          {
            "name": "raw_offset",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 9,
          },
          {
            "name": "timezone",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 10,
          },
          {
            "name": "unixtime",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 11,
          },
          {
            "name": "utc_datetime",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 12,
          },
          {
            "name": "utc_offset",
            "req": False,
            "type": "`$STRING`",
            "active": True,
            "index$": 13,
          },
          {
            "name": "week_number",
            "req": False,
            "type": "`$INTEGER`",
            "active": True,
            "index$": 14,
          },
        ],
        "name": "timezone",
        "op": {
          "list": {
            "name": "list",
            "points": [
              {
                "method": "GET",
                "orig": "/timezone",
                "parts": [
                  "timezone",
                ],
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "active": True,
                "args": {},
                "select": {},
                "index$": 0,
              },
            ],
            "input": "data",
            "key$": "list",
          },
          "load": {
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
                      "active": True,
                    },
                    {
                      "example": "New_York",
                      "kind": "param",
                      "name": "location",
                      "orig": "location",
                      "reqd": True,
                      "type": "`$STRING`",
                      "active": True,
                    },
                  ],
                },
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
                "active": True,
                "index$": 0,
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
                      "active": True,
                    },
                  ],
                },
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
                "active": True,
                "index$": 1,
              },
            ],
            "input": "data",
            "key$": "load",
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
