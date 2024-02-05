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

// checks if the IserverAccountAccountIdAlertActivatePost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IserverAccountAccountIdAlertActivatePost200Response{}

// IserverAccountAccountIdAlertActivatePost200Response struct for IserverAccountAccountIdAlertActivatePost200Response
type IserverAccountAccountIdAlertActivatePost200Response struct {
	RequestId   *int32  `json:"request_id,omitempty"`
	OrderId     *int32  `json:"order_id,omitempty"`
	Success     *bool   `json:"success,omitempty"`
	Text        *string `json:"text,omitempty"`
	OrderStatus *string `json:"order_status,omitempty"`
	FailureList *string `json:"failure_list,omitempty"`
}

// NewIserverAccountAccountIdAlertActivatePost200Response instantiates a new IserverAccountAccountIdAlertActivatePost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIserverAccountAccountIdAlertActivatePost200Response() *IserverAccountAccountIdAlertActivatePost200Response {
	this := IserverAccountAccountIdAlertActivatePost200Response{}
	return &this
}

// NewIserverAccountAccountIdAlertActivatePost200ResponseWithDefaults instantiates a new IserverAccountAccountIdAlertActivatePost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIserverAccountAccountIdAlertActivatePost200ResponseWithDefaults() *IserverAccountAccountIdAlertActivatePost200Response {
	this := IserverAccountAccountIdAlertActivatePost200Response{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetRequestId() int32 {
	if o == nil || IsNil(o.RequestId) {
		var ret int32
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetRequestIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given int32 and assigns it to the RequestId field.
func (o *IserverAccountAccountIdAlertActivatePost200Response) SetRequestId(v int32) {
	o.RequestId = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetOrderId() int32 {
	if o == nil || IsNil(o.OrderId) {
		var ret int32
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetOrderIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given int32 and assigns it to the OrderId field.
func (o *IserverAccountAccountIdAlertActivatePost200Response) SetOrderId(v int32) {
	o.OrderId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *IserverAccountAccountIdAlertActivatePost200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *IserverAccountAccountIdAlertActivatePost200Response) SetText(v string) {
	o.Text = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *IserverAccountAccountIdAlertActivatePost200Response) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetFailureList returns the FailureList field value if set, zero value otherwise.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetFailureList() string {
	if o == nil || IsNil(o.FailureList) {
		var ret string
		return ret
	}
	return *o.FailureList
}

// GetFailureListOk returns a tuple with the FailureList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) GetFailureListOk() (*string, bool) {
	if o == nil || IsNil(o.FailureList) {
		return nil, false
	}
	return o.FailureList, true
}

// HasFailureList returns a boolean if a field has been set.
func (o *IserverAccountAccountIdAlertActivatePost200Response) HasFailureList() bool {
	if o != nil && !IsNil(o.FailureList) {
		return true
	}

	return false
}

// SetFailureList gets a reference to the given string and assigns it to the FailureList field.
func (o *IserverAccountAccountIdAlertActivatePost200Response) SetFailureList(v string) {
	o.FailureList = &v
}

func (o IserverAccountAccountIdAlertActivatePost200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IserverAccountAccountIdAlertActivatePost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["order_status"] = o.OrderStatus
	}
	if !IsNil(o.FailureList) {
		toSerialize["failure_list"] = o.FailureList
	}
	return toSerialize, nil
}

type NullableIserverAccountAccountIdAlertActivatePost200Response struct {
	value *IserverAccountAccountIdAlertActivatePost200Response
	isSet bool
}

func (v NullableIserverAccountAccountIdAlertActivatePost200Response) Get() *IserverAccountAccountIdAlertActivatePost200Response {
	return v.value
}

func (v *NullableIserverAccountAccountIdAlertActivatePost200Response) Set(val *IserverAccountAccountIdAlertActivatePost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIserverAccountAccountIdAlertActivatePost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIserverAccountAccountIdAlertActivatePost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIserverAccountAccountIdAlertActivatePost200Response(val *IserverAccountAccountIdAlertActivatePost200Response) *NullableIserverAccountAccountIdAlertActivatePost200Response {
	return &NullableIserverAccountAccountIdAlertActivatePost200Response{value: val, isSet: true}
}

func (v NullableIserverAccountAccountIdAlertActivatePost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIserverAccountAccountIdAlertActivatePost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}