// Typed models for the WorldTime SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Ipn is the typed data model for the ipn entity.
type Ipn struct {
}

// Ipn2 is the typed data model for the ipn2 entity.
type Ipn2 struct {
	Abbreviation *string `json:"abbreviation,omitempty"`
	ClientIp *string `json:"client_ip,omitempty"`
	Datetime *string `json:"datetime,omitempty"`
	DayOfWeek *int `json:"day_of_week,omitempty"`
	DayOfYear *int `json:"day_of_year,omitempty"`
	Dst *bool `json:"dst,omitempty"`
	DstFrom *string `json:"dst_from,omitempty"`
	DstOffset *int `json:"dst_offset,omitempty"`
	DstUntil *string `json:"dst_until,omitempty"`
	RawOffset *int `json:"raw_offset,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Unixtime *int `json:"unixtime,omitempty"`
	UtcDatetime *string `json:"utc_datetime,omitempty"`
	UtcOffset *string `json:"utc_offset,omitempty"`
	WeekNumber *int `json:"week_number,omitempty"`
}

// Ipn2LoadMatch is the typed request payload for Ipn2.LoadTyped.
type Ipn2LoadMatch struct {
	Ipv4 string `json:"ipv4"`
}

// Timezone is the typed data model for the timezone entity.
type Timezone struct {
	Abbreviation *string `json:"abbreviation,omitempty"`
	ClientIp *string `json:"client_ip,omitempty"`
	Datetime *string `json:"datetime,omitempty"`
	DayOfWeek *int `json:"day_of_week,omitempty"`
	DayOfYear *int `json:"day_of_year,omitempty"`
	Dst *bool `json:"dst,omitempty"`
	DstFrom *string `json:"dst_from,omitempty"`
	DstOffset *int `json:"dst_offset,omitempty"`
	DstUntil *string `json:"dst_until,omitempty"`
	RawOffset *int `json:"raw_offset,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Unixtime *int `json:"unixtime,omitempty"`
	UtcDatetime *string `json:"utc_datetime,omitempty"`
	UtcOffset *string `json:"utc_offset,omitempty"`
	WeekNumber *int `json:"week_number,omitempty"`
}

// TimezoneLoadMatch is the typed request payload for Timezone.LoadTyped.
type TimezoneLoadMatch struct {
	Area string `json:"area"`
	Location string `json:"location"`
	Id string `json:"id"`
}

// TimezoneListMatch mirrors the timezone fields as an all-optional match
// filter (Go analog of Partial<Timezone>).
type TimezoneListMatch struct {
	Abbreviation *string `json:"abbreviation,omitempty"`
	ClientIp *string `json:"client_ip,omitempty"`
	Datetime *string `json:"datetime,omitempty"`
	DayOfWeek *int `json:"day_of_week,omitempty"`
	DayOfYear *int `json:"day_of_year,omitempty"`
	Dst *bool `json:"dst,omitempty"`
	DstFrom *string `json:"dst_from,omitempty"`
	DstOffset *int `json:"dst_offset,omitempty"`
	DstUntil *string `json:"dst_until,omitempty"`
	RawOffset *int `json:"raw_offset,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Unixtime *int `json:"unixtime,omitempty"`
	UtcDatetime *string `json:"utc_datetime,omitempty"`
	UtcOffset *string `json:"utc_offset,omitempty"`
	WeekNumber *int `json:"week_number,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
