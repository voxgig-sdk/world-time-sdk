// Typed models for the WorldTime SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/world-time-sdk/go/core"
)

// Ipn is the typed data model for the ipn entity.
type Ipn struct {
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

// IpnLoadMatch is the typed request payload for Ipn.LoadTyped.
type IpnLoadMatch struct {
	Ipv4 *string `json:"ipv4,omitempty"`
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
	Area *string `json:"area,omitempty"`
	Location *string `json:"location,omitempty"`
	Id *string `json:"id,omitempty"`
}

// TimezoneListMatch is the typed request payload for Timezone.ListTyped.
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

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
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

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
