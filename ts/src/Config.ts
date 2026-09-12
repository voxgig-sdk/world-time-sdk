
import { BaseFeature } from './feature/base/BaseFeature'
import { TestFeature } from './feature/test/TestFeature'



const FEATURE_CLASS: Record<string, typeof BaseFeature> = {
   test: TestFeature,

}


// Per-feature plugin DEFINITIONS (voxgig/plugin `Definition` values), from
// the model's active plugin groups. A feature that takes a `plugins` option
// (secrets over sekreto) reads its own entry; a feature with no plugins has
// none. Named imports above make each definition statically reachable, so
// an SDK carries exactly the plugin modules its model selects — the same
// leanness the old side-effect registry imports bought, without a registry.
const FEATURE_PLUGINS: Record<string, any[]> = {
  
}


class Config {

  makeFeature(this: any, fn: string) {
    const fc = FEATURE_CLASS[fn]
    const fi = new fc()
    // TODO: errors etc
    return fi
  }

  // False for a feature added at runtime via options.extend (station's
  // adopt path) - the constructor uses this to skip makeFeature for names
  // no generated class backs.
  hasFeature(this: any, fn: string) {
    return null != FEATURE_CLASS[fn]
  }


  main = {
    name: 'WorldTime',
        slug: "world-time",
    version: "0.0.1",
    target: "ts",

  }


  feature = {
     test:     {
      "options": {
        "active": false
      },
      "transport": "base"
    },

  }


  options = {
    base: "https://worldtimeapi.org/api",

    headers: {
      "content-type": "application/json"
    },

    entity: {
      
      ipn: {
      },

      timezone: {
      },

    }
  }


  entity = {
    "ipn": {
      "fields": [
        {
          "name": "abbreviation",
          "short": "The abbreviated name of the timezone",
          "type": "`$STRING`"
        },
        {
          "name": "client_ip",
          "short": "The IP address of the client",
          "type": "`$STRING`"
        },
        {
          "format": "date-time",
          "name": "datetime",
          "short": "The current datetime in ISO 8601 format",
          "type": "`$STRING`"
        },
        {
          "name": "day_of_week",
          "short": "The day of the week (0-6, where 0 is Sunday)",
          "type": "`$INTEGER`"
        },
        {
          "name": "day_of_year",
          "short": "The day of the year (1-365/366)",
          "type": "`$INTEGER`"
        },
        {
          "name": "dst",
          "short": "Whether daylight saving time is currently in effect",
          "type": "`$BOOLEAN`"
        },
        {
          "format": "date-time",
          "name": "dst_from",
          "short": "The datetime when DST starts",
          "type": "`$STRING`"
        },
        {
          "name": "dst_offset",
          "short": "The DST offset in seconds",
          "type": "`$INTEGER`"
        },
        {
          "format": "date-time",
          "name": "dst_until",
          "short": "The datetime when DST ends",
          "type": "`$STRING`"
        },
        {
          "name": "raw_offset",
          "short": "The raw offset from UTC in seconds",
          "type": "`$INTEGER`"
        },
        {
          "name": "timezone",
          "short": "The IANA timezone identifier",
          "type": "`$STRING`"
        },
        {
          "name": "unixtime",
          "short": "The current Unix timestamp",
          "type": "`$INTEGER`"
        },
        {
          "format": "date-time",
          "name": "utc_datetime",
          "short": "The current UTC datetime in ISO 8601 format",
          "type": "`$STRING`"
        },
        {
          "name": "utc_offset",
          "short": "The UTC offset in ±HH:MM format",
          "type": "`$STRING`"
        },
        {
          "name": "week_number",
          "short": "The ISO week number of the year",
          "type": "`$INTEGER`"
        }
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
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/ip/{ipv4}",
              "segments": [
                {
                  "lit": "ip"
                },
                {
                  "var": "ipv4"
                }
              ],
              "select": {
                "exist": [
                  "ipv4"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "ip",
                "{ipv4}"
              ]
            },
            {
              "args": {},
              "kind": "http",
              "method": "GET",
              "orig": "/ip",
              "segments": [
                {
                  "lit": "ip"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "ip"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "ip"
          ]
        ]
      }
    },
    "timezone": {
      "fields": [
        {
          "name": "abbreviation",
          "short": "The abbreviated name of the timezone",
          "type": "`$STRING`"
        },
        {
          "name": "client_ip",
          "short": "The IP address of the client",
          "type": "`$STRING`"
        },
        {
          "format": "date-time",
          "name": "datetime",
          "short": "The current datetime in ISO 8601 format",
          "type": "`$STRING`"
        },
        {
          "name": "day_of_week",
          "short": "The day of the week (0-6, where 0 is Sunday)",
          "type": "`$INTEGER`"
        },
        {
          "name": "day_of_year",
          "short": "The day of the year (1-365/366)",
          "type": "`$INTEGER`"
        },
        {
          "name": "dst",
          "short": "Whether daylight saving time is currently in effect",
          "type": "`$BOOLEAN`"
        },
        {
          "format": "date-time",
          "name": "dst_from",
          "short": "The datetime when DST starts",
          "type": "`$STRING`"
        },
        {
          "name": "dst_offset",
          "short": "The DST offset in seconds",
          "type": "`$INTEGER`"
        },
        {
          "format": "date-time",
          "name": "dst_until",
          "short": "The datetime when DST ends",
          "type": "`$STRING`"
        },
        {
          "name": "id",
          "type": "`$STRING`"
        },
        {
          "name": "raw_offset",
          "short": "The raw offset from UTC in seconds",
          "type": "`$INTEGER`"
        },
        {
          "name": "timezone",
          "short": "The IANA timezone identifier",
          "type": "`$STRING`"
        },
        {
          "name": "unixtime",
          "short": "The current Unix timestamp",
          "type": "`$INTEGER`"
        },
        {
          "format": "date-time",
          "name": "utc_datetime",
          "short": "The current UTC datetime in ISO 8601 format",
          "type": "`$STRING`"
        },
        {
          "name": "utc_offset",
          "short": "The UTC offset in ±HH:MM format",
          "type": "`$STRING`"
        },
        {
          "name": "week_number",
          "short": "The ISO week number of the year",
          "type": "`$INTEGER`"
        }
      ],
      "id": {
        "field": "id",
        "name": "id"
      },
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
              "segments": [
                {
                  "lit": "timezone"
                }
              ],
              "select": {},
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "timezone"
              ]
            }
          ]
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
                    "reqd": true,
                    "type": "`$STRING`"
                  },
                  {
                    "example": "New_York",
                    "kind": "param",
                    "name": "location",
                    "orig": "location",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/timezone/{area}/{location}",
              "segments": [
                {
                  "lit": "timezone"
                },
                {
                  "var": "area"
                },
                {
                  "var": "location"
                }
              ],
              "select": {
                "exist": [
                  "area",
                  "location"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "timezone",
                "{area}",
                "{location}"
              ]
            },
            {
              "args": {
                "params": [
                  {
                    "example": "Europe",
                    "kind": "param",
                    "name": "id",
                    "orig": "area",
                    "reqd": true,
                    "type": "`$STRING`"
                  }
                ]
              },
              "kind": "http",
              "method": "GET",
              "orig": "/timezone/{area}",
              "rename": {
                "param": {
                  "area": "id"
                }
              },
              "segments": [
                {
                  "lit": "timezone"
                },
                {
                  "var": "id"
                }
              ],
              "select": {
                "exist": [
                  "id"
                ]
              },
              "transform": {
                "req": "`reqdata`",
                "res": "`body`"
              },
              "parts": [
                "timezone",
                "{id}"
              ]
            }
          ]
        }
      },
      "relations": {
        "ancestors": [
          [
            "timezone"
          ]
        ]
      }
    }
  }
}


const config = new Config()

export {
  config,
  FEATURE_PLUGINS,
}

