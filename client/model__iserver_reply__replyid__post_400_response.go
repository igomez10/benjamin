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

// checks if the IserverReplyReplyidPost400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IserverReplyReplyidPost400Response{}

// IserverReplyReplyidPost400Response struct for IserverReplyReplyidPost400Response
type IserverReplyReplyidPost400Response struct {
	// for example-order not confirmed
	Error      *string `json:"error,omitempty"`
	StatusCode *int32  `json:"statusCode,omitempty"`
}

// NewIserverReplyReplyidPost400Response instantiates a new IserverReplyReplyidPost400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIserverReplyReplyidPost400Response() *IserverReplyReplyidPost400Response {
	this := IserverReplyReplyidPost400Response{}
	return &this
}

// NewIserverReplyReplyidPost400ResponseWithDefaults instantiates a new IserverReplyReplyidPost400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIserverReplyReplyidPost400ResponseWithDefaults() *IserverReplyReplyidPost400Response {
	this := IserverReplyReplyidPost400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *IserverReplyReplyidPost400Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverReplyReplyidPost400Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *IserverReplyReplyidPost400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *IserverReplyReplyidPost400Response) SetError(v string) {
	o.Error = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *IserverReplyReplyidPost400Response) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverReplyReplyidPost400Response) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *IserverReplyReplyidPost400Response) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *IserverReplyReplyidPost400Response) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o IserverReplyReplyidPost400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IserverReplyReplyidPost400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableIserverReplyReplyidPost400Response struct {
	value *IserverReplyReplyidPost400Response
	isSet bool
}

func (v NullableIserverReplyReplyidPost400Response) Get() *IserverReplyReplyidPost400Response {
	return v.value
}

func (v *NullableIserverReplyReplyidPost400Response) Set(val *IserverReplyReplyidPost400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIserverReplyReplyidPost400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIserverReplyReplyidPost400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIserverReplyReplyidPost400Response(val *IserverReplyReplyidPost400Response) *NullableIserverReplyReplyidPost400Response {
	return &NullableIserverReplyReplyidPost400Response{value: val, isSet: true}
}

func (v NullableIserverReplyReplyidPost400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIserverReplyReplyidPost400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
