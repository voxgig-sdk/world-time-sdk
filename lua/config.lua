-- WorldTime SDK configuration

-- Build a fresh, fully materialised config table. Every call rebuilds the
-- whole structure, so prefer require("config_shared") unless you need a
-- private copy you intend to mutate.
local function make_config()
  return {
    main = {
      name = "WorldTime",
    },
    feature = {
      ["test"] = {
        ["options"] = {
          ["active"] = false,
        },
      },
    },
    options = {
      base = "https://worldtimeapi.org/api",
      headers = {
        ["content-type"] = "application/json",
      },
      entity = {
        ["ipn"] = {},
        ["timezone"] = {},
      },
    },
    entity = {
      ["ipn"] = {
        ["fields"] = {
          {
            ["name"] = "abbreviation",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "client_ip",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "datetime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "day_of_week",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "day_of_year",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "dst",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "dst_from",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "dst_offset",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "dst_until",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "raw_offset",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "timezone",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "unixtime",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "utc_datetime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "utc_offset",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "week_number",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "ipn",
        ["op"] = {
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["example"] = "8.8.8.8",
                      ["kind"] = "param",
                      ["name"] = "ipv4",
                      ["orig"] = "ipv4",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/ip/{ipv4}",
                ["parts"] = {
                  "ip",
                  "{ipv4}",
                },
                ["select"] = {
                  ["exist"] = {
                    "ipv4",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/ip",
                ["parts"] = {
                  "ip",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "ip",
            },
          },
        },
      },
      ["timezone"] = {
        ["fields"] = {
          {
            ["name"] = "abbreviation",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "client_ip",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "datetime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "day_of_week",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "day_of_year",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "dst",
            ["type"] = "`$BOOLEAN`",
          },
          {
            ["name"] = "dst_from",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "dst_offset",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "dst_until",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "raw_offset",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "timezone",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "unixtime",
            ["type"] = "`$INTEGER`",
          },
          {
            ["name"] = "utc_datetime",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "utc_offset",
            ["type"] = "`$STRING`",
          },
          {
            ["name"] = "week_number",
            ["type"] = "`$INTEGER`",
          },
        },
        ["name"] = "timezone",
        ["op"] = {
          ["list"] = {
            ["input"] = "data",
            ["name"] = "list",
            ["points"] = {
              {
                ["args"] = {},
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/timezone",
                ["parts"] = {
                  "timezone",
                },
                ["select"] = {},
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
          ["load"] = {
            ["input"] = "data",
            ["name"] = "load",
            ["points"] = {
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["example"] = "America",
                      ["kind"] = "param",
                      ["name"] = "area",
                      ["orig"] = "area",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                    {
                      ["example"] = "New_York",
                      ["kind"] = "param",
                      ["name"] = "location",
                      ["orig"] = "location",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/timezone/{area}/{location}",
                ["parts"] = {
                  "timezone",
                  "{area}",
                  "{location}",
                },
                ["select"] = {
                  ["exist"] = {
                    "area",
                    "location",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
              {
                ["args"] = {
                  ["params"] = {
                    {
                      ["example"] = "Europe",
                      ["kind"] = "param",
                      ["name"] = "id",
                      ["orig"] = "area",
                      ["reqd"] = true,
                      ["type"] = "`$STRING`",
                    },
                  },
                },
                ["kind"] = "http",
                ["method"] = "GET",
                ["orig"] = "/timezone/{area}",
                ["parts"] = {
                  "timezone",
                  "{id}",
                },
                ["rename"] = {
                  ["param"] = {
                    ["area"] = "id",
                  },
                },
                ["select"] = {
                  ["exist"] = {
                    "id",
                  },
                },
                ["transform"] = {
                  ["req"] = "`reqdata`",
                  ["res"] = "`body`",
                },
              },
            },
          },
        },
        ["relations"] = {
          ["ancestors"] = {
            {
              "timezone",
            },
          },
        },
      },
    },
  }
end


local function make_feature(name)
  local features = require("features")
  local factory = features[name]
  if factory ~= nil then
    return factory()
  end
  return features.base()
end


-- Attach make_feature to the SDK class
local function setup_sdk(SDK)
  SDK._make_feature = make_feature
end


return make_config
