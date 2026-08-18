// Typed models for the WorldTime SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Ipn {
  abbreviation?: string
  client_ip?: string
  datetime?: string
  day_of_week?: number
  day_of_year?: number
  dst?: boolean
  dst_from?: string
  dst_offset?: number
  dst_until?: string
  raw_offset?: number
  timezone?: string
  unixtime?: number
  utc_datetime?: string
  utc_offset?: string
  week_number?: number
}

export interface IpnLoadMatch {
  ipv4: string
}

export interface Timezone {
  abbreviation?: string
  client_ip?: string
  datetime?: string
  day_of_week?: number
  day_of_year?: number
  dst?: boolean
  dst_from?: string
  dst_offset?: number
  dst_until?: string
  raw_offset?: number
  timezone?: string
  unixtime?: number
  utc_datetime?: string
  utc_offset?: string
  week_number?: number
}

export interface TimezoneLoadMatch {
  id: string
}

export interface TimezoneListMatch {
  abbreviation?: string
  client_ip?: string
  datetime?: string
  day_of_week?: number
  day_of_year?: number
  dst?: boolean
  dst_from?: string
  dst_offset?: number
  dst_until?: string
  raw_offset?: number
  timezone?: string
  unixtime?: number
  utc_datetime?: string
  utc_offset?: string
  week_number?: number
}

