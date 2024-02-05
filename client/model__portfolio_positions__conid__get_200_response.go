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

// checks if the PortfolioPositionsConidGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioPositionsConidGet200Response{}

// PortfolioPositionsConidGet200Response struct for PortfolioPositionsConidGet200Response
type PortfolioPositionsConidGet200Response struct {
	ACCTID []PositionInner `json:"ACCTID,omitempty"`
}

// NewPortfolioPositionsConidGet200Response instantiates a new PortfolioPositionsConidGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioPositionsConidGet200Response() *PortfolioPositionsConidGet200Response {
	this := PortfolioPositionsConidGet200Response{}
	return &this
}

// NewPortfolioPositionsConidGet200ResponseWithDefaults instantiates a new PortfolioPositionsConidGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioPositionsConidGet200ResponseWithDefaults() *PortfolioPositionsConidGet200Response {
	this := PortfolioPositionsConidGet200Response{}
	return &this
}

// GetACCTID returns the ACCTID field value if set, zero value otherwise.
func (o *PortfolioPositionsConidGet200Response) GetACCTID() []PositionInner {
	if o == nil || IsNil(o.ACCTID) {
		var ret []PositionInner
		return ret
	}
	return o.ACCTID
}

// GetACCTIDOk returns a tuple with the ACCTID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioPositionsConidGet200Response) GetACCTIDOk() ([]PositionInner, bool) {
	if o == nil || IsNil(o.ACCTID) {
		return nil, false
	}
	return o.ACCTID, true
}

// HasACCTID returns a boolean if a field has been set.
func (o *PortfolioPositionsConidGet200Response) HasACCTID() bool {
	if o != nil && !IsNil(o.ACCTID) {
		return true
	}

	return false
}

// SetACCTID gets a reference to the given []PositionInner and assigns it to the ACCTID field.
func (o *PortfolioPositionsConidGet200Response) SetACCTID(v []PositionInner) {
	o.ACCTID = v
}

func (o PortfolioPositionsConidGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioPositionsConidGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ACCTID) {
		toSerialize["ACCTID"] = o.ACCTID
	}
	return toSerialize, nil
}

type NullablePortfolioPositionsConidGet200Response struct {
	value *PortfolioPositionsConidGet200Response
	isSet bool
}

func (v NullablePortfolioPositionsConidGet200Response) Get() *PortfolioPositionsConidGet200Response {
	return v.value
}

func (v *NullablePortfolioPositionsConidGet200Response) Set(val *PortfolioPositionsConidGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioPositionsConidGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioPositionsConidGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioPositionsConidGet200Response(val *PortfolioPositionsConidGet200Response) *NullablePortfolioPositionsConidGet200Response {
	return &NullablePortfolioPositionsConidGet200Response{value: val, isSet: true}
}

func (v NullablePortfolioPositionsConidGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioPositionsConidGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
