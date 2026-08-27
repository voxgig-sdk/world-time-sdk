-- Typed models for the WorldTime SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Ipn
---@field abbreviation? string
---@field client_ip? string
---@field datetime? string
---@field day_of_week? number
---@field day_of_year? number
---@field dst? boolean
---@field dst_from? string
---@field dst_offset? number
---@field dst_until? string
---@field raw_offset? number
---@field timezone? string
---@field unixtime? number
---@field utc_datetime? string
---@field utc_offset? string
---@field week_number? number

---@class IpnLoadMatch
---@field ipv4 string

---@class Timezone
---@field abbreviation? string
---@field client_ip? string
---@field datetime? string
---@field day_of_week? number
---@field day_of_year? number
---@field dst? boolean
---@field dst_from? string
---@field dst_offset? number
---@field dst_until? string
---@field id? string
---@field raw_offset? number
---@field timezone? string
---@field unixtime? number
---@field utc_datetime? string
---@field utc_offset? string
---@field week_number? number

---@class TimezoneLoadMatch
---@field id string

---@class TimezoneListMatch
---@field abbreviation? string
---@field client_ip? string
---@field datetime? string
---@field day_of_week? number
---@field day_of_year? number
---@field dst? boolean
---@field dst_from? string
---@field dst_offset? number
---@field dst_until? string
---@field id? string
---@field raw_offset? number
---@field timezone? string
---@field unixtime? number
---@field utc_datetime? string
---@field utc_offset? string
---@field week_number? number

local M = {}

return M
