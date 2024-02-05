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

// checks if the WagersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WagersInner{}

// WagersInner List of wagers
type WagersInner struct {
	Conid *float32 `json:"conid,omitempty"`
	Curr  *string  `json:"curr,omitempty"`
	Desc  *string  `json:"desc,omitempty"`
	Part  *string  `json:"part,omitempty"`
}

// NewWagersInner instantiates a new WagersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWagersInner() *WagersInner {
	this := WagersInner{}
	return &this
}

// NewWagersInnerWithDefaults instantiates a new WagersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWagersInnerWithDefaults() *WagersInner {
	this := WagersInner{}
	return &this
}

// GetConid returns the Conid field value if set, zero value otherwise.
func (o *WagersInner) GetConid() float32 {
	if o == nil || IsNil(o.Conid) {
		var ret float32
		return ret
	}
	return *o.Conid
}

// GetConidOk returns a tuple with the Conid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WagersInner) GetConidOk() (*float32, bool) {
	if o == nil || IsNil(o.Conid) {
		return nil, false
	}
	return o.Conid, true
}

// HasConid returns a boolean if a field has been set.
func (o *WagersInner) HasConid() bool {
	if o != nil && !IsNil(o.Conid) {
		return true
	}

	return false
}

// SetConid gets a reference to the given float32 and assigns it to the Conid field.
func (o *WagersInner) SetConid(v float32) {
	o.Conid = &v
}

// GetCurr returns the Curr field value if set, zero value otherwise.
func (o *WagersInner) GetCurr() string {
	if o == nil || IsNil(o.Curr) {
		var ret string
		return ret
	}
	return *o.Curr
}

// GetCurrOk returns a tuple with the Curr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WagersInner) GetCurrOk() (*string, bool) {
	if o == nil || IsNil(o.Curr) {
		return nil, false
	}
	return o.Curr, true
}

// HasCurr returns a boolean if a field has been set.
func (o *WagersInner) HasCurr() bool {
	if o != nil && !IsNil(o.Curr) {
		return true
	}

	return false
}

// SetCurr gets a reference to the given string and assigns it to the Curr field.
func (o *WagersInner) SetCurr(v string) {
	o.Curr = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *WagersInner) GetDesc() string {
	if o == nil || IsNil(o.Desc) {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WagersInner) GetDescOk() (*string, bool) {
	if o == nil || IsNil(o.Desc) {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *WagersInner) HasDesc() bool {
	if o != nil && !IsNil(o.Desc) {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *WagersInner) SetDesc(v string) {
	o.Desc = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *WagersInner) GetPart() string {
	if o == nil || IsNil(o.Part) {
		var ret string
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WagersInner) GetPartOk() (*string, bool) {
	if o == nil || IsNil(o.Part) {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *WagersInner) HasPart() bool {
	if o != nil && !IsNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given string and assigns it to the Part field.
func (o *WagersInner) SetPart(v string) {
	o.Part = &v
}

func (o WagersInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WagersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conid) {
		toSerialize["conid"] = o.Conid
	}
	if !IsNil(o.Curr) {
		toSerialize["curr"] = o.Curr
	}
	if !IsNil(o.Desc) {
		toSerialize["desc"] = o.Desc
	}
	if !IsNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	return toSerialize, nil
}

type NullableWagersInner struct {
	value *WagersInner
	isSet bool
}

func (v NullableWagersInner) Get() *WagersInner {
	return v.value
}

func (v *NullableWagersInner) Set(val *WagersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableWagersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableWagersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWagersInner(val *WagersInner) *NullableWagersInner {
	return &NullableWagersInner{value: val, isSet: true}
}

func (v NullableWagersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWagersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
