<?php
declare(strict_types=1);

// Typed models for the WorldTime SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Ipn entity data model. */
class Ipn
{
    public ?string $abbreviation = null;
    public ?string $client_ip = null;
    public ?string $datetime = null;
    public ?int $day_of_week = null;
    public ?int $day_of_year = null;
    public ?bool $dst = null;
    public ?string $dst_from = null;
    public ?int $dst_offset = null;
    public ?string $dst_until = null;
    public ?int $raw_offset = null;
    public ?string $timezone = null;
    public ?int $unixtime = null;
    public ?string $utc_datetime = null;
    public ?string $utc_offset = null;
    public ?int $week_number = null;
}

/** Request payload for Ipn#load. */
class IpnLoadMatch
{
    public string $ipv4;
}

/** Timezone entity data model. */
class Timezone
{
    public ?string $abbreviation = null;
    public ?string $client_ip = null;
    public ?string $datetime = null;
    public ?int $day_of_week = null;
    public ?int $day_of_year = null;
    public ?bool $dst = null;
    public ?string $dst_from = null;
    public ?int $dst_offset = null;
    public ?string $dst_until = null;
    public ?int $raw_offset = null;
    public ?string $timezone = null;
    public ?int $unixtime = null;
    public ?string $utc_datetime = null;
    public ?string $utc_offset = null;
    public ?int $week_number = null;
}

/** Request payload for Timezone#load. */
class TimezoneLoadMatch
{
    public string $id;
}

/** Request payload for Timezone#list. */
class TimezoneListMatch
{
    public ?string $abbreviation = null;
    public ?string $client_ip = null;
    public ?string $datetime = null;
    public ?int $day_of_week = null;
    public ?int $day_of_year = null;
    public ?bool $dst = null;
    public ?string $dst_from = null;
    public ?int $dst_offset = null;
    public ?string $dst_until = null;
    public ?int $raw_offset = null;
    public ?string $timezone = null;
    public ?int $unixtime = null;
    public ?string $utc_datetime = null;
    public ?string $utc_offset = null;
    public ?int $week_number = null;
}

