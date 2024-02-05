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

// checks if the IserverMarketdataUnsubscribeallGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IserverMarketdataUnsubscribeallGet200Response{}

// IserverMarketdataUnsubscribeallGet200Response struct for IserverMarketdataUnsubscribeallGet200Response
type IserverMarketdataUnsubscribeallGet200Response struct {
	// true means market data is cancelled, false means it is not.
	Confirmed *bool `json:"confirmed,omitempty"`
}

// NewIserverMarketdataUnsubscribeallGet200Response instantiates a new IserverMarketdataUnsubscribeallGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIserverMarketdataUnsubscribeallGet200Response() *IserverMarketdataUnsubscribeallGet200Response {
	this := IserverMarketdataUnsubscribeallGet200Response{}
	return &this
}

// NewIserverMarketdataUnsubscribeallGet200ResponseWithDefaults instantiates a new IserverMarketdataUnsubscribeallGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIserverMarketdataUnsubscribeallGet200ResponseWithDefaults() *IserverMarketdataUnsubscribeallGet200Response {
	this := IserverMarketdataUnsubscribeallGet200Response{}
	return &this
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *IserverMarketdataUnsubscribeallGet200Response) GetConfirmed() bool {
	if o == nil || IsNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverMarketdataUnsubscribeallGet200Response) GetConfirmedOk() (*bool, bool) {
	if o == nil || IsNil(o.Confirmed) {
		return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *IserverMarketdataUnsubscribeallGet200Response) HasConfirmed() bool {
	if o != nil && !IsNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *IserverMarketdataUnsubscribeallGet200Response) SetConfirmed(v bool) {
	o.Confirmed = &v
}

func (o IserverMarketdataUnsubscribeallGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IserverMarketdataUnsubscribeallGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	return toSerialize, nil
}

type NullableIserverMarketdataUnsubscribeallGet200Response struct {
	value *IserverMarketdataUnsubscribeallGet200Response
	isSet bool
}

func (v NullableIserverMarketdataUnsubscribeallGet200Response) Get() *IserverMarketdataUnsubscribeallGet200Response {
	return v.value
}

func (v *NullableIserverMarketdataUnsubscribeallGet200Response) Set(val *IserverMarketdataUnsubscribeallGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIserverMarketdataUnsubscribeallGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIserverMarketdataUnsubscribeallGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIserverMarketdataUnsubscribeallGet200Response(val *IserverMarketdataUnsubscribeallGet200Response) *NullableIserverMarketdataUnsubscribeallGet200Response {
	return &NullableIserverMarketdataUnsubscribeallGet200Response{value: val, isSet: true}
}

func (v NullableIserverMarketdataUnsubscribeallGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIserverMarketdataUnsubscribeallGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
