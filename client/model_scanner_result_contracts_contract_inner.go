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

// checks if the ScannerResultContractsContractInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannerResultContractsContractInner{}

// ScannerResultContractsContractInner struct for ScannerResultContractsContractInner
type ScannerResultContractsContractInner struct {
	InScanTime *string `json:"inScanTime,omitempty"`
	Distance   *int32  `json:"distance,omitempty"`
	ContractID *int32  `json:"contractID,omitempty"`
}

// NewScannerResultContractsContractInner instantiates a new ScannerResultContractsContractInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannerResultContractsContractInner() *ScannerResultContractsContractInner {
	this := ScannerResultContractsContractInner{}
	return &this
}

// NewScannerResultContractsContractInnerWithDefaults instantiates a new ScannerResultContractsContractInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannerResultContractsContractInnerWithDefaults() *ScannerResultContractsContractInner {
	this := ScannerResultContractsContractInner{}
	return &this
}

// GetInScanTime returns the InScanTime field value if set, zero value otherwise.
func (o *ScannerResultContractsContractInner) GetInScanTime() string {
	if o == nil || IsNil(o.InScanTime) {
		var ret string
		return ret
	}
	return *o.InScanTime
}

// GetInScanTimeOk returns a tuple with the InScanTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerResultContractsContractInner) GetInScanTimeOk() (*string, bool) {
	if o == nil || IsNil(o.InScanTime) {
		return nil, false
	}
	return o.InScanTime, true
}

// HasInScanTime returns a boolean if a field has been set.
func (o *ScannerResultContractsContractInner) HasInScanTime() bool {
	if o != nil && !IsNil(o.InScanTime) {
		return true
	}

	return false
}

// SetInScanTime gets a reference to the given string and assigns it to the InScanTime field.
func (o *ScannerResultContractsContractInner) SetInScanTime(v string) {
	o.InScanTime = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *ScannerResultContractsContractInner) GetDistance() int32 {
	if o == nil || IsNil(o.Distance) {
		var ret int32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerResultContractsContractInner) GetDistanceOk() (*int32, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *ScannerResultContractsContractInner) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int32 and assigns it to the Distance field.
func (o *ScannerResultContractsContractInner) SetDistance(v int32) {
	o.Distance = &v
}

// GetContractID returns the ContractID field value if set, zero value otherwise.
func (o *ScannerResultContractsContractInner) GetContractID() int32 {
	if o == nil || IsNil(o.ContractID) {
		var ret int32
		return ret
	}
	return *o.ContractID
}

// GetContractIDOk returns a tuple with the ContractID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerResultContractsContractInner) GetContractIDOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractID) {
		return nil, false
	}
	return o.ContractID, true
}

// HasContractID returns a boolean if a field has been set.
func (o *ScannerResultContractsContractInner) HasContractID() bool {
	if o != nil && !IsNil(o.ContractID) {
		return true
	}

	return false
}

// SetContractID gets a reference to the given int32 and assigns it to the ContractID field.
func (o *ScannerResultContractsContractInner) SetContractID(v int32) {
	o.ContractID = &v
}

func (o ScannerResultContractsContractInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannerResultContractsContractInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InScanTime) {
		toSerialize["inScanTime"] = o.InScanTime
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if !IsNil(o.ContractID) {
		toSerialize["contractID"] = o.ContractID
	}
	return toSerialize, nil
}

type NullableScannerResultContractsContractInner struct {
	value *ScannerResultContractsContractInner
	isSet bool
}

func (v NullableScannerResultContractsContractInner) Get() *ScannerResultContractsContractInner {
	return v.value
}

func (v *NullableScannerResultContractsContractInner) Set(val *ScannerResultContractsContractInner) {
	v.value = val
	v.isSet = true
}

func (v NullableScannerResultContractsContractInner) IsSet() bool {
	return v.isSet
}

func (v *NullableScannerResultContractsContractInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannerResultContractsContractInner(val *ScannerResultContractsContractInner) *NullableScannerResultContractsContractInner {
	return &NullableScannerResultContractsContractInner{value: val, isSet: true}
}

func (v NullableScannerResultContractsContractInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannerResultContractsContractInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
