# frozen_string_literal: true

# Typed models for the WorldTime SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Ipn entity data model.
class Ipn
end

# Ipn2 entity data model.
#
# @!attribute [rw] abbreviation
#   @return [String, nil]
#
# @!attribute [rw] client_ip
#   @return [String, nil]
#
# @!attribute [rw] datetime
#   @return [String, nil]
#
# @!attribute [rw] day_of_week
#   @return [Integer, nil]
#
# @!attribute [rw] day_of_year
#   @return [Integer, nil]
#
# @!attribute [rw] dst
#   @return [Boolean, nil]
#
# @!attribute [rw] dst_from
#   @return [String, nil]
#
# @!attribute [rw] dst_offset
#   @return [Integer, nil]
#
# @!attribute [rw] dst_until
#   @return [String, nil]
#
# @!attribute [rw] raw_offset
#   @return [Integer, nil]
#
# @!attribute [rw] timezone
#   @return [String, nil]
#
# @!attribute [rw] unixtime
#   @return [Integer, nil]
#
# @!attribute [rw] utc_datetime
#   @return [String, nil]
#
# @!attribute [rw] utc_offset
#   @return [String, nil]
#
# @!attribute [rw] week_number
#   @return [Integer, nil]
Ipn2 = Struct.new(
  :abbreviation,
  :client_ip,
  :datetime,
  :day_of_week,
  :day_of_year,
  :dst,
  :dst_from,
  :dst_offset,
  :dst_until,
  :raw_offset,
  :timezone,
  :unixtime,
  :utc_datetime,
  :utc_offset,
  :week_number,
  keyword_init: true
)

# Request payload for Ipn2#load.
#
# @!attribute [rw] ipv4
#   @return [String]
Ipn2LoadMatch = Struct.new(
  :ipv4,
  keyword_init: true
)

# Timezone entity data model.
#
# @!attribute [rw] abbreviation
#   @return [String, nil]
#
# @!attribute [rw] client_ip
#   @return [String, nil]
#
# @!attribute [rw] datetime
#   @return [String, nil]
#
# @!attribute [rw] day_of_week
#   @return [Integer, nil]
#
# @!attribute [rw] day_of_year
#   @return [Integer, nil]
#
# @!attribute [rw] dst
#   @return [Boolean, nil]
#
# @!attribute [rw] dst_from
#   @return [String, nil]
#
# @!attribute [rw] dst_offset
#   @return [Integer, nil]
#
# @!attribute [rw] dst_until
#   @return [String, nil]
#
# @!attribute [rw] raw_offset
#   @return [Integer, nil]
#
# @!attribute [rw] timezone
#   @return [String, nil]
#
# @!attribute [rw] unixtime
#   @return [Integer, nil]
#
# @!attribute [rw] utc_datetime
#   @return [String, nil]
#
# @!attribute [rw] utc_offset
#   @return [String, nil]
#
# @!attribute [rw] week_number
#   @return [Integer, nil]
Timezone = Struct.new(
  :abbreviation,
  :client_ip,
  :datetime,
  :day_of_week,
  :day_of_year,
  :dst,
  :dst_from,
  :dst_offset,
  :dst_until,
  :raw_offset,
  :timezone,
  :unixtime,
  :utc_datetime,
  :utc_offset,
  :week_number,
  keyword_init: true
)

# Request payload for Timezone#load.
#
# @!attribute [rw] area
#   @return [String]
#
# @!attribute [rw] location
#   @return [String]
#
# @!attribute [rw] id
#   @return [String]
TimezoneLoadMatch = Struct.new(
  :area,
  :location,
  :id,
  keyword_init: true
)

# Match filter for Timezone#list (any subset of Timezone fields).
#
# @!attribute [rw] abbreviation
#   @return [String, nil]
#
# @!attribute [rw] client_ip
#   @return [String, nil]
#
# @!attribute [rw] datetime
#   @return [String, nil]
#
# @!attribute [rw] day_of_week
#   @return [Integer, nil]
#
# @!attribute [rw] day_of_year
#   @return [Integer, nil]
#
# @!attribute [rw] dst
#   @return [Boolean, nil]
#
# @!attribute [rw] dst_from
#   @return [String, nil]
#
# @!attribute [rw] dst_offset
#   @return [Integer, nil]
#
# @!attribute [rw] dst_until
#   @return [String, nil]
#
# @!attribute [rw] raw_offset
#   @return [Integer, nil]
#
# @!attribute [rw] timezone
#   @return [String, nil]
#
# @!attribute [rw] unixtime
#   @return [Integer, nil]
#
# @!attribute [rw] utc_datetime
#   @return [String, nil]
#
# @!attribute [rw] utc_offset
#   @return [String, nil]
#
# @!attribute [rw] week_number
#   @return [Integer, nil]
TimezoneListMatch = Struct.new(
  :abbreviation,
  :client_ip,
  :datetime,
  :day_of_week,
  :day_of_year,
  :dst,
  :dst_from,
  :dst_offset,
  :dst_until,
  :raw_offset,
  :timezone,
  :unixtime,
  :utc_datetime,
  :utc_offset,
  :week_number,
  keyword_init: true
)

