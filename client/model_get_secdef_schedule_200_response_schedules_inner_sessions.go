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

// checks if the GetSecdefSchedule200ResponseSchedulesInnerSessions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSecdefSchedule200ResponseSchedulesInnerSessions{}

// GetSecdefSchedule200ResponseSchedulesInnerSessions If the LIQUID hours differs from the total trading day then a separate 'session' tag is returned.
type GetSecdefSchedule200ResponseSchedulesInnerSessions struct {
	OpeningTime *int32 `json:"openingTime,omitempty"`
	ClosingTime *int32 `json:"closingTime,omitempty"`
	// If the whole trading day is considered LIQUID then the value 'LIQUID' is returned.
	Prop *string `json:"prop,omitempty"`
}

// NewGetSecdefSchedule200ResponseSchedulesInnerSessions instantiates a new GetSecdefSchedule200ResponseSchedulesInnerSessions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSecdefSchedule200ResponseSchedulesInnerSessions() *GetSecdefSchedule200ResponseSchedulesInnerSessions {
	this := GetSecdefSchedule200ResponseSchedulesInnerSessions{}
	return &this
}

// NewGetSecdefSchedule200ResponseSchedulesInnerSessionsWithDefaults instantiates a new GetSecdefSchedule200ResponseSchedulesInnerSessions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSecdefSchedule200ResponseSchedulesInnerSessionsWithDefaults() *GetSecdefSchedule200ResponseSchedulesInnerSessions {
	this := GetSecdefSchedule200ResponseSchedulesInnerSessions{}
	return &this
}

// GetOpeningTime returns the OpeningTime field value if set, zero value otherwise.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) GetOpeningTime() int32 {
	if o == nil || IsNil(o.OpeningTime) {
		var ret int32
		return ret
	}
	return *o.OpeningTime
}

// GetOpeningTimeOk returns a tuple with the OpeningTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) GetOpeningTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.OpeningTime) {
		return nil, false
	}
	return o.OpeningTime, true
}

// HasOpeningTime returns a boolean if a field has been set.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) HasOpeningTime() bool {
	if o != nil && !IsNil(o.OpeningTime) {
		return true
	}

	return false
}

// SetOpeningTime gets a reference to the given int32 and assigns it to the OpeningTime field.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) SetOpeningTime(v int32) {
	o.OpeningTime = &v
}

// GetClosingTime returns the ClosingTime field value if set, zero value otherwise.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) GetClosingTime() int32 {
	if o == nil || IsNil(o.ClosingTime) {
		var ret int32
		return ret
	}
	return *o.ClosingTime
}

// GetClosingTimeOk returns a tuple with the ClosingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) GetClosingTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ClosingTime) {
		return nil, false
	}
	return o.ClosingTime, true
}

// HasClosingTime returns a boolean if a field has been set.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) HasClosingTime() bool {
	if o != nil && !IsNil(o.ClosingTime) {
		return true
	}

	return false
}

// SetClosingTime gets a reference to the given int32 and assigns it to the ClosingTime field.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) SetClosingTime(v int32) {
	o.ClosingTime = &v
}

// GetProp returns the Prop field value if set, zero value otherwise.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) GetProp() string {
	if o == nil || IsNil(o.Prop) {
		var ret string
		return ret
	}
	return *o.Prop
}

// GetPropOk returns a tuple with the Prop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) GetPropOk() (*string, bool) {
	if o == nil || IsNil(o.Prop) {
		return nil, false
	}
	return o.Prop, true
}

// HasProp returns a boolean if a field has been set.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) HasProp() bool {
	if o != nil && !IsNil(o.Prop) {
		return true
	}

	return false
}

// SetProp gets a reference to the given string and assigns it to the Prop field.
func (o *GetSecdefSchedule200ResponseSchedulesInnerSessions) SetProp(v string) {
	o.Prop = &v
}

func (o GetSecdefSchedule200ResponseSchedulesInnerSessions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSecdefSchedule200ResponseSchedulesInnerSessions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OpeningTime) {
		toSerialize["openingTime"] = o.OpeningTime
	}
	if !IsNil(o.ClosingTime) {
		toSerialize["closingTime"] = o.ClosingTime
	}
	if !IsNil(o.Prop) {
		toSerialize["prop"] = o.Prop
	}
	return toSerialize, nil
}

type NullableGetSecdefSchedule200ResponseSchedulesInnerSessions struct {
	value *GetSecdefSchedule200ResponseSchedulesInnerSessions
	isSet bool
}

func (v NullableGetSecdefSchedule200ResponseSchedulesInnerSessions) Get() *GetSecdefSchedule200ResponseSchedulesInnerSessions {
	return v.value
}

func (v *NullableGetSecdefSchedule200ResponseSchedulesInnerSessions) Set(val *GetSecdefSchedule200ResponseSchedulesInnerSessions) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSecdefSchedule200ResponseSchedulesInnerSessions) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSecdefSchedule200ResponseSchedulesInnerSessions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSecdefSchedule200ResponseSchedulesInnerSessions(val *GetSecdefSchedule200ResponseSchedulesInnerSessions) *NullableGetSecdefSchedule200ResponseSchedulesInnerSessions {
	return &NullableGetSecdefSchedule200ResponseSchedulesInnerSessions{value: val, isSet: true}
}

func (v NullableGetSecdefSchedule200ResponseSchedulesInnerSessions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSecdefSchedule200ResponseSchedulesInnerSessions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}