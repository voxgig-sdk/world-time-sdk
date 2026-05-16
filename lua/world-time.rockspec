package = "voxgig-sdk-world-time"
version = "0.0-1"
source = {
  url = "git://github.com/voxgig-sdk/world-time-sdk.git"
}
description = {
  summary = "WorldTime SDK for Lua",
  license = "MIT"
}
dependencies = {
  "lua >= 5.3",
  "dkjson >= 2.5",
  "dkjson >= 2.5",
}
build = {
  type = "builtin",
  modules = {
    ["world-time_sdk"] = "world-time_sdk.lua",
    ["config"] = "config.lua",
    ["features"] = "features.lua",
  }
}
