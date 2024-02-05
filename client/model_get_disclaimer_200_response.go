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

// checks if the GetDisclaimer200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDisclaimer200Response{}

// GetDisclaimer200Response struct for GetDisclaimer200Response
type GetDisclaimer200Response struct {
	// disclaimer message
	DT *string `json:"DT,omitempty"`
	// fyi code
	FC *string `json:"FC,omitempty"`
}

// NewGetDisclaimer200Response instantiates a new GetDisclaimer200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDisclaimer200Response() *GetDisclaimer200Response {
	this := GetDisclaimer200Response{}
	return &this
}

// NewGetDisclaimer200ResponseWithDefaults instantiates a new GetDisclaimer200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDisclaimer200ResponseWithDefaults() *GetDisclaimer200Response {
	this := GetDisclaimer200Response{}
	return &this
}

// GetDT returns the DT field value if set, zero value otherwise.
func (o *GetDisclaimer200Response) GetDT() string {
	if o == nil || IsNil(o.DT) {
		var ret string
		return ret
	}
	return *o.DT
}

// GetDTOk returns a tuple with the DT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDisclaimer200Response) GetDTOk() (*string, bool) {
	if o == nil || IsNil(o.DT) {
		return nil, false
	}
	return o.DT, true
}

// HasDT returns a boolean if a field has been set.
func (o *GetDisclaimer200Response) HasDT() bool {
	if o != nil && !IsNil(o.DT) {
		return true
	}

	return false
}

// SetDT gets a reference to the given string and assigns it to the DT field.
func (o *GetDisclaimer200Response) SetDT(v string) {
	o.DT = &v
}

// GetFC returns the FC field value if set, zero value otherwise.
func (o *GetDisclaimer200Response) GetFC() string {
	if o == nil || IsNil(o.FC) {
		var ret string
		return ret
	}
	return *o.FC
}

// GetFCOk returns a tuple with the FC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDisclaimer200Response) GetFCOk() (*string, bool) {
	if o == nil || IsNil(o.FC) {
		return nil, false
	}
	return o.FC, true
}

// HasFC returns a boolean if a field has been set.
func (o *GetDisclaimer200Response) HasFC() bool {
	if o != nil && !IsNil(o.FC) {
		return true
	}

	return false
}

// SetFC gets a reference to the given string and assigns it to the FC field.
func (o *GetDisclaimer200Response) SetFC(v string) {
	o.FC = &v
}

func (o GetDisclaimer200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDisclaimer200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DT) {
		toSerialize["DT"] = o.DT
	}
	if !IsNil(o.FC) {
		toSerialize["FC"] = o.FC
	}
	return toSerialize, nil
}

type NullableGetDisclaimer200Response struct {
	value *GetDisclaimer200Response
	isSet bool
}

func (v NullableGetDisclaimer200Response) Get() *GetDisclaimer200Response {
	return v.value
}

func (v *NullableGetDisclaimer200Response) Set(val *GetDisclaimer200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDisclaimer200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDisclaimer200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDisclaimer200Response(val *GetDisclaimer200Response) *NullableGetDisclaimer200Response {
	return &NullableGetDisclaimer200Response{value: val, isSet: true}
}

func (v NullableGetDisclaimer200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDisclaimer200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}