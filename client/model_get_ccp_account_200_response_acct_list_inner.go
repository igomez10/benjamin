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

// checks if the GetCCPAccount200ResponseAcctListInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCCPAccount200ResponseAcctListInner{}

// GetCCPAccount200ResponseAcctListInner struct for GetCCPAccount200ResponseAcctListInner
type GetCCPAccount200ResponseAcctListInner struct {
	// For multi-account structures each trading account will numbered from 0 to ...
	Var0 *string `json:"0,omitempty"`
}

// NewGetCCPAccount200ResponseAcctListInner instantiates a new GetCCPAccount200ResponseAcctListInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCCPAccount200ResponseAcctListInner() *GetCCPAccount200ResponseAcctListInner {
	this := GetCCPAccount200ResponseAcctListInner{}
	return &this
}

// NewGetCCPAccount200ResponseAcctListInnerWithDefaults instantiates a new GetCCPAccount200ResponseAcctListInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCCPAccount200ResponseAcctListInnerWithDefaults() *GetCCPAccount200ResponseAcctListInner {
	this := GetCCPAccount200ResponseAcctListInner{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *GetCCPAccount200ResponseAcctListInner) GetVar0() string {
	if o == nil || IsNil(o.Var0) {
		var ret string
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCCPAccount200ResponseAcctListInner) GetVar0Ok() (*string, bool) {
	if o == nil || IsNil(o.Var0) {
		return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *GetCCPAccount200ResponseAcctListInner) HasVar0() bool {
	if o != nil && !IsNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given string and assigns it to the Var0 field.
func (o *GetCCPAccount200ResponseAcctListInner) SetVar0(v string) {
	o.Var0 = &v
}

func (o GetCCPAccount200ResponseAcctListInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCCPAccount200ResponseAcctListInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var0) {
		toSerialize["0"] = o.Var0
	}
	return toSerialize, nil
}

type NullableGetCCPAccount200ResponseAcctListInner struct {
	value *GetCCPAccount200ResponseAcctListInner
	isSet bool
}

func (v NullableGetCCPAccount200ResponseAcctListInner) Get() *GetCCPAccount200ResponseAcctListInner {
	return v.value
}

func (v *NullableGetCCPAccount200ResponseAcctListInner) Set(val *GetCCPAccount200ResponseAcctListInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCCPAccount200ResponseAcctListInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCCPAccount200ResponseAcctListInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCCPAccount200ResponseAcctListInner(val *GetCCPAccount200ResponseAcctListInner) *NullableGetCCPAccount200ResponseAcctListInner {
	return &NullableGetCCPAccount200ResponseAcctListInner{value: val, isSet: true}
}

func (v NullableGetCCPAccount200ResponseAcctListInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCCPAccount200ResponseAcctListInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
