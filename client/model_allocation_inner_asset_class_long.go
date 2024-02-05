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

// checks if the AllocationInnerAssetClassLong type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllocationInnerAssetClassLong{}

// AllocationInnerAssetClassLong long positions allocation
type AllocationInnerAssetClassLong struct {
	STK  *float32 `json:"STK,omitempty"`
	OPT  *float32 `json:"OPT,omitempty"`
	FUT  *float32 `json:"FUT,omitempty"`
	WAR  *float32 `json:"WAR,omitempty"`
	BOND *float32 `json:"BOND,omitempty"`
	CASH *float32 `json:"CASH,omitempty"`
}

// NewAllocationInnerAssetClassLong instantiates a new AllocationInnerAssetClassLong object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocationInnerAssetClassLong() *AllocationInnerAssetClassLong {
	this := AllocationInnerAssetClassLong{}
	return &this
}

// NewAllocationInnerAssetClassLongWithDefaults instantiates a new AllocationInnerAssetClassLong object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocationInnerAssetClassLongWithDefaults() *AllocationInnerAssetClassLong {
	this := AllocationInnerAssetClassLong{}
	return &this
}

// GetSTK returns the STK field value if set, zero value otherwise.
func (o *AllocationInnerAssetClassLong) GetSTK() float32 {
	if o == nil || IsNil(o.STK) {
		var ret float32
		return ret
	}
	return *o.STK
}

// GetSTKOk returns a tuple with the STK field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationInnerAssetClassLong) GetSTKOk() (*float32, bool) {
	if o == nil || IsNil(o.STK) {
		return nil, false
	}
	return o.STK, true
}

// HasSTK returns a boolean if a field has been set.
func (o *AllocationInnerAssetClassLong) HasSTK() bool {
	if o != nil && !IsNil(o.STK) {
		return true
	}

	return false
}

// SetSTK gets a reference to the given float32 and assigns it to the STK field.
func (o *AllocationInnerAssetClassLong) SetSTK(v float32) {
	o.STK = &v
}

// GetOPT returns the OPT field value if set, zero value otherwise.
func (o *AllocationInnerAssetClassLong) GetOPT() float32 {
	if o == nil || IsNil(o.OPT) {
		var ret float32
		return ret
	}
	return *o.OPT
}

// GetOPTOk returns a tuple with the OPT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationInnerAssetClassLong) GetOPTOk() (*float32, bool) {
	if o == nil || IsNil(o.OPT) {
		return nil, false
	}
	return o.OPT, true
}

// HasOPT returns a boolean if a field has been set.
func (o *AllocationInnerAssetClassLong) HasOPT() bool {
	if o != nil && !IsNil(o.OPT) {
		return true
	}

	return false
}

// SetOPT gets a reference to the given float32 and assigns it to the OPT field.
func (o *AllocationInnerAssetClassLong) SetOPT(v float32) {
	o.OPT = &v
}

// GetFUT returns the FUT field value if set, zero value otherwise.
func (o *AllocationInnerAssetClassLong) GetFUT() float32 {
	if o == nil || IsNil(o.FUT) {
		var ret float32
		return ret
	}
	return *o.FUT
}

// GetFUTOk returns a tuple with the FUT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationInnerAssetClassLong) GetFUTOk() (*float32, bool) {
	if o == nil || IsNil(o.FUT) {
		return nil, false
	}
	return o.FUT, true
}

// HasFUT returns a boolean if a field has been set.
func (o *AllocationInnerAssetClassLong) HasFUT() bool {
	if o != nil && !IsNil(o.FUT) {
		return true
	}

	return false
}

// SetFUT gets a reference to the given float32 and assigns it to the FUT field.
func (o *AllocationInnerAssetClassLong) SetFUT(v float32) {
	o.FUT = &v
}

// GetWAR returns the WAR field value if set, zero value otherwise.
func (o *AllocationInnerAssetClassLong) GetWAR() float32 {
	if o == nil || IsNil(o.WAR) {
		var ret float32
		return ret
	}
	return *o.WAR
}

// GetWAROk returns a tuple with the WAR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationInnerAssetClassLong) GetWAROk() (*float32, bool) {
	if o == nil || IsNil(o.WAR) {
		return nil, false
	}
	return o.WAR, true
}

// HasWAR returns a boolean if a field has been set.
func (o *AllocationInnerAssetClassLong) HasWAR() bool {
	if o != nil && !IsNil(o.WAR) {
		return true
	}

	return false
}

// SetWAR gets a reference to the given float32 and assigns it to the WAR field.
func (o *AllocationInnerAssetClassLong) SetWAR(v float32) {
	o.WAR = &v
}

// GetBOND returns the BOND field value if set, zero value otherwise.
func (o *AllocationInnerAssetClassLong) GetBOND() float32 {
	if o == nil || IsNil(o.BOND) {
		var ret float32
		return ret
	}
	return *o.BOND
}

// GetBONDOk returns a tuple with the BOND field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationInnerAssetClassLong) GetBONDOk() (*float32, bool) {
	if o == nil || IsNil(o.BOND) {
		return nil, false
	}
	return o.BOND, true
}

// HasBOND returns a boolean if a field has been set.
func (o *AllocationInnerAssetClassLong) HasBOND() bool {
	if o != nil && !IsNil(o.BOND) {
		return true
	}

	return false
}

// SetBOND gets a reference to the given float32 and assigns it to the BOND field.
func (o *AllocationInnerAssetClassLong) SetBOND(v float32) {
	o.BOND = &v
}

// GetCASH returns the CASH field value if set, zero value otherwise.
func (o *AllocationInnerAssetClassLong) GetCASH() float32 {
	if o == nil || IsNil(o.CASH) {
		var ret float32
		return ret
	}
	return *o.CASH
}

// GetCASHOk returns a tuple with the CASH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocationInnerAssetClassLong) GetCASHOk() (*float32, bool) {
	if o == nil || IsNil(o.CASH) {
		return nil, false
	}
	return o.CASH, true
}

// HasCASH returns a boolean if a field has been set.
func (o *AllocationInnerAssetClassLong) HasCASH() bool {
	if o != nil && !IsNil(o.CASH) {
		return true
	}

	return false
}

// SetCASH gets a reference to the given float32 and assigns it to the CASH field.
func (o *AllocationInnerAssetClassLong) SetCASH(v float32) {
	o.CASH = &v
}

func (o AllocationInnerAssetClassLong) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllocationInnerAssetClassLong) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.STK) {
		toSerialize["STK"] = o.STK
	}
	if !IsNil(o.OPT) {
		toSerialize["OPT"] = o.OPT
	}
	if !IsNil(o.FUT) {
		toSerialize["FUT"] = o.FUT
	}
	if !IsNil(o.WAR) {
		toSerialize["WAR"] = o.WAR
	}
	if !IsNil(o.BOND) {
		toSerialize["BOND"] = o.BOND
	}
	if !IsNil(o.CASH) {
		toSerialize["CASH"] = o.CASH
	}
	return toSerialize, nil
}

type NullableAllocationInnerAssetClassLong struct {
	value *AllocationInnerAssetClassLong
	isSet bool
}

func (v NullableAllocationInnerAssetClassLong) Get() *AllocationInnerAssetClassLong {
	return v.value
}

func (v *NullableAllocationInnerAssetClassLong) Set(val *AllocationInnerAssetClassLong) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocationInnerAssetClassLong) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocationInnerAssetClassLong) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocationInnerAssetClassLong(val *AllocationInnerAssetClassLong) *NullableAllocationInnerAssetClassLong {
	return &NullableAllocationInnerAssetClassLong{value: val, isSet: true}
}

func (v NullableAllocationInnerAssetClassLong) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocationInnerAssetClassLong) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
