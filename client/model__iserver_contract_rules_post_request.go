/*
Client Portal Web API

Client Poral Web API

API version: 1.0.0
Contact: e@e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the IserverContractRulesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IserverContractRulesPostRequest{}

// IserverContractRulesPostRequest struct for IserverContractRulesPostRequest
type IserverContractRulesPostRequest struct {
	// IBKR contract identifier
	Conid string `json:"conid"`
	// Side of the market rules apply too. Set to **true** for Buy Orders, set to **false** for Sell Orders
	IsBuy bool `json:"isBuy"`
}

type _IserverContractRulesPostRequest IserverContractRulesPostRequest

// NewIserverContractRulesPostRequest instantiates a new IserverContractRulesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIserverContractRulesPostRequest(conid string, isBuy bool) *IserverContractRulesPostRequest {
	this := IserverContractRulesPostRequest{}
	this.Conid = conid
	this.IsBuy = isBuy
	return &this
}

// NewIserverContractRulesPostRequestWithDefaults instantiates a new IserverContractRulesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIserverContractRulesPostRequestWithDefaults() *IserverContractRulesPostRequest {
	this := IserverContractRulesPostRequest{}
	return &this
}

// GetConid returns the Conid field value
func (o *IserverContractRulesPostRequest) GetConid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Conid
}

// GetConidOk returns a tuple with the Conid field value
// and a boolean to check if the value has been set.
func (o *IserverContractRulesPostRequest) GetConidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conid, true
}

// SetConid sets field value
func (o *IserverContractRulesPostRequest) SetConid(v string) {
	o.Conid = v
}

// GetIsBuy returns the IsBuy field value
func (o *IserverContractRulesPostRequest) GetIsBuy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBuy
}

// GetIsBuyOk returns a tuple with the IsBuy field value
// and a boolean to check if the value has been set.
func (o *IserverContractRulesPostRequest) GetIsBuyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBuy, true
}

// SetIsBuy sets field value
func (o *IserverContractRulesPostRequest) SetIsBuy(v bool) {
	o.IsBuy = v
}

func (o IserverContractRulesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IserverContractRulesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conid"] = o.Conid
	toSerialize["isBuy"] = o.IsBuy
	return toSerialize, nil
}

func (o *IserverContractRulesPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conid",
		"isBuy",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIserverContractRulesPostRequest := _IserverContractRulesPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIserverContractRulesPostRequest)

	if err != nil {
		return err
	}

	*o = IserverContractRulesPostRequest(varIserverContractRulesPostRequest)

	return err
}

type NullableIserverContractRulesPostRequest struct {
	value *IserverContractRulesPostRequest
	isSet bool
}

func (v NullableIserverContractRulesPostRequest) Get() *IserverContractRulesPostRequest {
	return v.value
}

func (v *NullableIserverContractRulesPostRequest) Set(val *IserverContractRulesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIserverContractRulesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIserverContractRulesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIserverContractRulesPostRequest(val *IserverContractRulesPostRequest) *NullableIserverContractRulesPostRequest {
	return &NullableIserverContractRulesPostRequest{value: val, isSet: true}
}

func (v NullableIserverContractRulesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIserverContractRulesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
