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

// checks if the IserverContractRulesPost200ResponseRulesInnerCqtTypesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IserverContractRulesPost200ResponseRulesInnerCqtTypesInner{}

// IserverContractRulesPost200ResponseRulesInnerCqtTypesInner struct for IserverContractRulesPost200ResponseRulesInnerCqtTypesInner
type IserverContractRulesPost200ResponseRulesInnerCqtTypesInner struct {
	// order types that support cash quantity trades
	Var0 *string `json:"0,omitempty"`
}

// NewIserverContractRulesPost200ResponseRulesInnerCqtTypesInner instantiates a new IserverContractRulesPost200ResponseRulesInnerCqtTypesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIserverContractRulesPost200ResponseRulesInnerCqtTypesInner() *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner {
	this := IserverContractRulesPost200ResponseRulesInnerCqtTypesInner{}
	return &this
}

// NewIserverContractRulesPost200ResponseRulesInnerCqtTypesInnerWithDefaults instantiates a new IserverContractRulesPost200ResponseRulesInnerCqtTypesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIserverContractRulesPost200ResponseRulesInnerCqtTypesInnerWithDefaults() *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner {
	this := IserverContractRulesPost200ResponseRulesInnerCqtTypesInner{}
	return &this
}

// GetVar0 returns the Var0 field value if set, zero value otherwise.
func (o *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner) GetVar0() string {
	if o == nil || IsNil(o.Var0) {
		var ret string
		return ret
	}
	return *o.Var0
}

// GetVar0Ok returns a tuple with the Var0 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner) GetVar0Ok() (*string, bool) {
	if o == nil || IsNil(o.Var0) {
		return nil, false
	}
	return o.Var0, true
}

// HasVar0 returns a boolean if a field has been set.
func (o *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner) HasVar0() bool {
	if o != nil && !IsNil(o.Var0) {
		return true
	}

	return false
}

// SetVar0 gets a reference to the given string and assigns it to the Var0 field.
func (o *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner) SetVar0(v string) {
	o.Var0 = &v
}

func (o IserverContractRulesPost200ResponseRulesInnerCqtTypesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IserverContractRulesPost200ResponseRulesInnerCqtTypesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var0) {
		toSerialize["0"] = o.Var0
	}
	return toSerialize, nil
}

type NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner struct {
	value *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner
	isSet bool
}

func (v NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner) Get() *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner {
	return v.value
}

func (v *NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner) Set(val *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner(val *IserverContractRulesPost200ResponseRulesInnerCqtTypesInner) *NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner {
	return &NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner{value: val, isSet: true}
}

func (v NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIserverContractRulesPost200ResponseRulesInnerCqtTypesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}