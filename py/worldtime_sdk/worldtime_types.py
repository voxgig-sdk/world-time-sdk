# Typed models for the WorldTime SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Ipn(TypedDict, total=False):
    abbreviation: str
    client_ip: str
    datetime: str
    day_of_week: int
    day_of_year: int
    dst: bool
    dst_from: str
    dst_offset: int
    dst_until: str
    raw_offset: int
    timezone: str
    unixtime: int
    utc_datetime: str
    utc_offset: str
    week_number: int


class IpnLoadMatch(TypedDict):
    ipv4: str


class Timezone(TypedDict, total=False):
    abbreviation: str
    client_ip: str
    datetime: str
    day_of_week: int
    day_of_year: int
    dst: bool
    dst_from: str
    dst_offset: int
    dst_until: str
    id: str
    raw_offset: int
    timezone: str
    unixtime: int
    utc_datetime: str
    utc_offset: str
    week_number: int


class TimezoneLoadMatch(TypedDict):
    id: str


class TimezoneListMatch(TypedDict, total=False):
    abbreviation: str
    client_ip: str
    datetime: str
    day_of_week: int
    day_of_year: int
    dst: bool
    dst_from: str
    dst_offset: int
    dst_until: str
    id: str
    raw_offset: int
    timezone: str
    unixtime: int
    utc_datetime: str
    utc_offset: str
    week_number: int
