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

// checks if the AllocationInnerGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllocationInnerGroup{}

// AllocationInnerGroup portfolio allocation by group
type AllocationInnerGroup struct {
	Long  *AllocationInnerGroupLong  `json:"long,omitempty"`
	Short *AllocationInnerGroupShort `json:"short,omitempty"`
}

// NewAllocationInnerGroup instantiates a new AllocationInnerGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocationInnerGroup() *AllocationInnerGroup {
	this := AllocationInnerGroup{}
	return &this
}

// NewAllocationInnerGroupWithDefaults instantiates a new AllocationInnerGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocationInnerGroupWithDefaults() *AllocationInnerGroup {
	this := AllocationInnerGroup{}
	return &this
}

// GetLong returns the Long field value if set, zero value otherwise.
func (o *AllocationInnerGroup) GetLong() AllocationInnerGroupLong {
	if o == nil || IsNil(o.Long) {
		var ret AllocationInnerGroupLong
		return ret
	}
	return *o.Long
}

// GetLongOk returns a tuple with the Long field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationInnerGroup) GetLongOk() (*AllocationInnerGroupLong, bool) {
	if o == nil || IsNil(o.Long) {
		return nil, false
	}
	return o.Long, true
}

// HasLong returns a boolean if a field has been set.
func (o *AllocationInnerGroup) HasLong() bool {
	if o != nil && !IsNil(o.Long) {
		return true
	}

	return false
}

// SetLong gets a reference to the given AllocationInnerGroupLong and assigns it to the Long field.
func (o *AllocationInnerGroup) SetLong(v AllocationInnerGroupLong) {
	o.Long = &v
}

// GetShort returns the Short field value if set, zero value otherwise.
func (o *AllocationInnerGroup) GetShort() AllocationInnerGroupShort {
	if o == nil || IsNil(o.Short) {
		var ret AllocationInnerGroupShort
		return ret
	}
	return *o.Short
}

// GetShortOk returns a tuple with the Short field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationInnerGroup) GetShortOk() (*AllocationInnerGroupShort, bool) {
	if o == nil || IsNil(o.Short) {
		return nil, false
	}
	return o.Short, true
}

// HasShort returns a boolean if a field has been set.
func (o *AllocationInnerGroup) HasShort() bool {
	if o != nil && !IsNil(o.Short) {
		return true
	}

	return false
}

// SetShort gets a reference to the given AllocationInnerGroupShort and assigns it to the Short field.
func (o *AllocationInnerGroup) SetShort(v AllocationInnerGroupShort) {
	o.Short = &v
}

func (o AllocationInnerGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllocationInnerGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Long) {
		toSerialize["long"] = o.Long
	}
	if !IsNil(o.Short) {
		toSerialize["short"] = o.Short
	}
	return toSerialize, nil
}

type NullableAllocationInnerGroup struct {
	value *AllocationInnerGroup
	isSet bool
}

func (v NullableAllocationInnerGroup) Get() *AllocationInnerGroup {
	return v.value
}

func (v *NullableAllocationInnerGroup) Set(val *AllocationInnerGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocationInnerGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocationInnerGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocationInnerGroup(val *AllocationInnerGroup) *NullableAllocationInnerGroup {
	return &NullableAllocationInnerGroup{value: val, isSet: true}
}

func (v NullableAllocationInnerGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocationInnerGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
