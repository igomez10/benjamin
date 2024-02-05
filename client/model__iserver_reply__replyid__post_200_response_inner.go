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

// checks if the IserverReplyReplyidPost200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IserverReplyReplyidPost200ResponseInner{}

// IserverReplyReplyidPost200ResponseInner struct for IserverReplyReplyidPost200ResponseInner
type IserverReplyReplyidPost200ResponseInner struct {
	OrderId      *string `json:"order_id,omitempty"`
	OrderStatus  *string `json:"order_status,omitempty"`
	LocalOrderId *string `json:"local_order_id,omitempty"`
}

// NewIserverReplyReplyidPost200ResponseInner instantiates a new IserverReplyReplyidPost200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIserverReplyReplyidPost200ResponseInner() *IserverReplyReplyidPost200ResponseInner {
	this := IserverReplyReplyidPost200ResponseInner{}
	return &this
}

// NewIserverReplyReplyidPost200ResponseInnerWithDefaults instantiates a new IserverReplyReplyidPost200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIserverReplyReplyidPost200ResponseInnerWithDefaults() *IserverReplyReplyidPost200ResponseInner {
	this := IserverReplyReplyidPost200ResponseInner{}
	return &this
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *IserverReplyReplyidPost200ResponseInner) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverReplyReplyidPost200ResponseInner) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *IserverReplyReplyidPost200ResponseInner) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *IserverReplyReplyidPost200ResponseInner) SetOrderId(v string) {
	o.OrderId = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *IserverReplyReplyidPost200ResponseInner) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverReplyReplyidPost200ResponseInner) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *IserverReplyReplyidPost200ResponseInner) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *IserverReplyReplyidPost200ResponseInner) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetLocalOrderId returns the LocalOrderId field value if set, zero value otherwise.
func (o *IserverReplyReplyidPost200ResponseInner) GetLocalOrderId() string {
	if o == nil || IsNil(o.LocalOrderId) {
		var ret string
		return ret
	}
	return *o.LocalOrderId
}

// GetLocalOrderIdOk returns a tuple with the LocalOrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverReplyReplyidPost200ResponseInner) GetLocalOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocalOrderId) {
		return nil, false
	}
	return o.LocalOrderId, true
}

// HasLocalOrderId returns a boolean if a field has been set.
func (o *IserverReplyReplyidPost200ResponseInner) HasLocalOrderId() bool {
	if o != nil && !IsNil(o.LocalOrderId) {
		return true
	}

	return false
}

// SetLocalOrderId gets a reference to the given string and assigns it to the LocalOrderId field.
func (o *IserverReplyReplyidPost200ResponseInner) SetLocalOrderId(v string) {
	o.LocalOrderId = &v
}

func (o IserverReplyReplyidPost200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IserverReplyReplyidPost200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	if !IsNil(o.LocalOrderId) {
		toSerialize["local_order_id"] = o.LocalOrderId
	}
	return toSerialize, nil
}

type NullableIserverReplyReplyidPost200ResponseInner struct {
	value *IserverReplyReplyidPost200ResponseInner
	isSet bool
}

func (v NullableIserverReplyReplyidPost200ResponseInner) Get() *IserverReplyReplyidPost200ResponseInner {
	return v.value
}

func (v *NullableIserverReplyReplyidPost200ResponseInner) Set(val *IserverReplyReplyidPost200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIserverReplyReplyidPost200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIserverReplyReplyidPost200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIserverReplyReplyidPost200ResponseInner(val *IserverReplyReplyidPost200ResponseInner) *NullableIserverReplyReplyidPost200ResponseInner {
	return &NullableIserverReplyReplyidPost200ResponseInner{value: val, isSet: true}
}

func (v NullableIserverReplyReplyidPost200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIserverReplyReplyidPost200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
