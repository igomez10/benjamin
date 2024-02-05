/*
Client Portal Web API

Client Poral Web API

API version: 1.0.0
Contact: e@e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CalendarRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalendarRequest{}

// CalendarRequest struct for CalendarRequest
type CalendarRequest struct {
	Date    *CalendarRequestDate    `json:"date,omitempty"`
	Filters *CalendarRequestFilters `json:"filters,omitempty"`
}

// NewCalendarRequest instantiates a new CalendarRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarRequest() *CalendarRequest {
	this := CalendarRequest{}
	return &this
}

// NewCalendarRequestWithDefaults instantiates a new CalendarRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarRequestWithDefaults() *CalendarRequest {
	this := CalendarRequest{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *CalendarRequest) GetDate() CalendarRequestDate {
	if o == nil || IsNil(o.Date) {
		var ret CalendarRequestDate
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarRequest) GetDateOk() (*CalendarRequestDate, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *CalendarRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given CalendarRequestDate and assigns it to the Date field.
func (o *CalendarRequest) SetDate(v CalendarRequestDate) {
	o.Date = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *CalendarRequest) GetFilters() CalendarRequestFilters {
	if o == nil || IsNil(o.Filters) {
		var ret CalendarRequestFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarRequest) GetFiltersOk() (*CalendarRequestFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *CalendarRequest) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given CalendarRequestFilters and assigns it to the Filters field.
func (o *CalendarRequest) SetFilters(v CalendarRequestFilters) {
	o.Filters = &v
}

func (o CalendarRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalendarRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	return toSerialize, nil
}

type NullableCalendarRequest struct {
	value *CalendarRequest
	isSet bool
}

func (v NullableCalendarRequest) Get() *CalendarRequest {
	return v.value
}

func (v *NullableCalendarRequest) Set(val *CalendarRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarRequest(val *CalendarRequest) *NullableCalendarRequest {
	return &NullableCalendarRequest{value: val, isSet: true}
}

func (v NullableCalendarRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}