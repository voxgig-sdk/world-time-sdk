# Typed models for the WorldTime SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.

from __future__ import annotations

from dataclasses import dataclass
from typing import Optional, Any


@dataclass
class Ipn:
    pass


@dataclass
class Ipn2:
    abbreviation: Optional[str] = None
    client_ip: Optional[str] = None
    datetime: Optional[str] = None
    day_of_week: Optional[int] = None
    day_of_year: Optional[int] = None
    dst: Optional[bool] = None
    dst_from: Optional[str] = None
    dst_offset: Optional[int] = None
    dst_until: Optional[str] = None
    raw_offset: Optional[int] = None
    timezone: Optional[str] = None
    unixtime: Optional[int] = None
    utc_datetime: Optional[str] = None
    utc_offset: Optional[str] = None
    week_number: Optional[int] = None


@dataclass
class Ipn2LoadMatch:
    ipv4: str


@dataclass
class Timezone:
    abbreviation: Optional[str] = None
    client_ip: Optional[str] = None
    datetime: Optional[str] = None
    day_of_week: Optional[int] = None
    day_of_year: Optional[int] = None
    dst: Optional[bool] = None
    dst_from: Optional[str] = None
    dst_offset: Optional[int] = None
    dst_until: Optional[str] = None
    raw_offset: Optional[int] = None
    timezone: Optional[str] = None
    unixtime: Optional[int] = None
    utc_datetime: Optional[str] = None
    utc_offset: Optional[str] = None
    week_number: Optional[int] = None


@dataclass
class TimezoneLoadMatch:
    area: str
    location: str
    id: str


@dataclass
class TimezoneListMatch:
    abbreviation: Optional[str] = None
    client_ip: Optional[str] = None
    datetime: Optional[str] = None
    day_of_week: Optional[int] = None
    day_of_year: Optional[int] = None
    dst: Optional[bool] = None
    dst_from: Optional[str] = None
    dst_offset: Optional[int] = None
    dst_until: Optional[str] = None
    raw_offset: Optional[int] = None
    timezone: Optional[str] = None
    unixtime: Optional[int] = None
    utc_datetime: Optional[str] = None
    utc_offset: Optional[str] = None
    week_number: Optional[int] = None

