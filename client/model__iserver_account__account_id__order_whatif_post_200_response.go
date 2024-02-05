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

// checks if the IserverAccountAccountIdOrderWhatifPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IserverAccountAccountIdOrderWhatifPost200Response{}

// IserverAccountAccountIdOrderWhatifPost200Response struct for IserverAccountAccountIdOrderWhatifPost200Response
type IserverAccountAccountIdOrderWhatifPost200Response struct {
	Amount      *IserverAccountAccountIdOrderWhatifPost200ResponseAmount `json:"amount,omitempty"`
	Equity      *IserverAccountAccountIdOrderWhatifPost200ResponseEquity `json:"equity,omitempty"`
	Initial     *IserverAccountAccountIdOrderWhatifPost200ResponseEquity `json:"initial,omitempty"`
	Maintenance *IserverAccountAccountIdOrderWhatifPost200ResponseEquity `json:"maintenance,omitempty"`
	Warn        *string                                                  `json:"warn,omitempty"`
	Error       *string                                                  `json:"error,omitempty"`
}

// NewIserverAccountAccountIdOrderWhatifPost200Response instantiates a new IserverAccountAccountIdOrderWhatifPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIserverAccountAccountIdOrderWhatifPost200Response() *IserverAccountAccountIdOrderWhatifPost200Response {
	this := IserverAccountAccountIdOrderWhatifPost200Response{}
	return &this
}

// NewIserverAccountAccountIdOrderWhatifPost200ResponseWithDefaults instantiates a new IserverAccountAccountIdOrderWhatifPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIserverAccountAccountIdOrderWhatifPost200ResponseWithDefaults() *IserverAccountAccountIdOrderWhatifPost200Response {
	this := IserverAccountAccountIdOrderWhatifPost200Response{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetAmount() IserverAccountAccountIdOrderWhatifPost200ResponseAmount {
	if o == nil || IsNil(o.Amount) {
		var ret IserverAccountAccountIdOrderWhatifPost200ResponseAmount
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetAmountOk() (*IserverAccountAccountIdOrderWhatifPost200ResponseAmount, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given IserverAccountAccountIdOrderWhatifPost200ResponseAmount and assigns it to the Amount field.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) SetAmount(v IserverAccountAccountIdOrderWhatifPost200ResponseAmount) {
	o.Amount = &v
}

// GetEquity returns the Equity field value if set, zero value otherwise.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetEquity() IserverAccountAccountIdOrderWhatifPost200ResponseEquity {
	if o == nil || IsNil(o.Equity) {
		var ret IserverAccountAccountIdOrderWhatifPost200ResponseEquity
		return ret
	}
	return *o.Equity
}

// GetEquityOk returns a tuple with the Equity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetEquityOk() (*IserverAccountAccountIdOrderWhatifPost200ResponseEquity, bool) {
	if o == nil || IsNil(o.Equity) {
		return nil, false
	}
	return o.Equity, true
}

// HasEquity returns a boolean if a field has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) HasEquity() bool {
	if o != nil && !IsNil(o.Equity) {
		return true
	}

	return false
}

// SetEquity gets a reference to the given IserverAccountAccountIdOrderWhatifPost200ResponseEquity and assigns it to the Equity field.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) SetEquity(v IserverAccountAccountIdOrderWhatifPost200ResponseEquity) {
	o.Equity = &v
}

// GetInitial returns the Initial field value if set, zero value otherwise.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetInitial() IserverAccountAccountIdOrderWhatifPost200ResponseEquity {
	if o == nil || IsNil(o.Initial) {
		var ret IserverAccountAccountIdOrderWhatifPost200ResponseEquity
		return ret
	}
	return *o.Initial
}

// GetInitialOk returns a tuple with the Initial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetInitialOk() (*IserverAccountAccountIdOrderWhatifPost200ResponseEquity, bool) {
	if o == nil || IsNil(o.Initial) {
		return nil, false
	}
	return o.Initial, true
}

// HasInitial returns a boolean if a field has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) HasInitial() bool {
	if o != nil && !IsNil(o.Initial) {
		return true
	}

	return false
}

// SetInitial gets a reference to the given IserverAccountAccountIdOrderWhatifPost200ResponseEquity and assigns it to the Initial field.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) SetInitial(v IserverAccountAccountIdOrderWhatifPost200ResponseEquity) {
	o.Initial = &v
}

// GetMaintenance returns the Maintenance field value if set, zero value otherwise.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetMaintenance() IserverAccountAccountIdOrderWhatifPost200ResponseEquity {
	if o == nil || IsNil(o.Maintenance) {
		var ret IserverAccountAccountIdOrderWhatifPost200ResponseEquity
		return ret
	}
	return *o.Maintenance
}

// GetMaintenanceOk returns a tuple with the Maintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetMaintenanceOk() (*IserverAccountAccountIdOrderWhatifPost200ResponseEquity, bool) {
	if o == nil || IsNil(o.Maintenance) {
		return nil, false
	}
	return o.Maintenance, true
}

// HasMaintenance returns a boolean if a field has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) HasMaintenance() bool {
	if o != nil && !IsNil(o.Maintenance) {
		return true
	}

	return false
}

// SetMaintenance gets a reference to the given IserverAccountAccountIdOrderWhatifPost200ResponseEquity and assigns it to the Maintenance field.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) SetMaintenance(v IserverAccountAccountIdOrderWhatifPost200ResponseEquity) {
	o.Maintenance = &v
}

// GetWarn returns the Warn field value if set, zero value otherwise.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetWarn() string {
	if o == nil || IsNil(o.Warn) {
		var ret string
		return ret
	}
	return *o.Warn
}

// GetWarnOk returns a tuple with the Warn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetWarnOk() (*string, bool) {
	if o == nil || IsNil(o.Warn) {
		return nil, false
	}
	return o.Warn, true
}

// HasWarn returns a boolean if a field has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) HasWarn() bool {
	if o != nil && !IsNil(o.Warn) {
		return true
	}

	return false
}

// SetWarn gets a reference to the given string and assigns it to the Warn field.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) SetWarn(v string) {
	o.Warn = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *IserverAccountAccountIdOrderWhatifPost200Response) SetError(v string) {
	o.Error = &v
}

func (o IserverAccountAccountIdOrderWhatifPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IserverAccountAccountIdOrderWhatifPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Equity) {
		toSerialize["equity"] = o.Equity
	}
	if !IsNil(o.Initial) {
		toSerialize["initial"] = o.Initial
	}
	if !IsNil(o.Maintenance) {
		toSerialize["maintenance"] = o.Maintenance
	}
	if !IsNil(o.Warn) {
		toSerialize["warn"] = o.Warn
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableIserverAccountAccountIdOrderWhatifPost200Response struct {
	value *IserverAccountAccountIdOrderWhatifPost200Response
	isSet bool
}

func (v NullableIserverAccountAccountIdOrderWhatifPost200Response) Get() *IserverAccountAccountIdOrderWhatifPost200Response {
	return v.value
}

func (v *NullableIserverAccountAccountIdOrderWhatifPost200Response) Set(val *IserverAccountAccountIdOrderWhatifPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableIserverAccountAccountIdOrderWhatifPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableIserverAccountAccountIdOrderWhatifPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIserverAccountAccountIdOrderWhatifPost200Response(val *IserverAccountAccountIdOrderWhatifPost200Response) *NullableIserverAccountAccountIdOrderWhatifPost200Response {
	return &NullableIserverAccountAccountIdOrderWhatifPost200Response{value: val, isSet: true}
}

func (v NullableIserverAccountAccountIdOrderWhatifPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIserverAccountAccountIdOrderWhatifPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}