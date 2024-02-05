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

// checks if the ScannerResultContracts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannerResultContracts{}

// ScannerResultContracts Contains list of contracts matching the scanner query
type ScannerResultContracts struct {
	Contract []ScannerResultContractsContractInner `json:"Contract,omitempty"`
}

// NewScannerResultContracts instantiates a new ScannerResultContracts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannerResultContracts() *ScannerResultContracts {
	this := ScannerResultContracts{}
	return &this
}

// NewScannerResultContractsWithDefaults instantiates a new ScannerResultContracts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannerResultContractsWithDefaults() *ScannerResultContracts {
	this := ScannerResultContracts{}
	return &this
}

// GetContract returns the Contract field value if set, zero value otherwise.
func (o *ScannerResultContracts) GetContract() []ScannerResultContractsContractInner {
	if o == nil || IsNil(o.Contract) {
		var ret []ScannerResultContractsContractInner
		return ret
	}
	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannerResultContracts) GetContractOk() ([]ScannerResultContractsContractInner, bool) {
	if o == nil || IsNil(o.Contract) {
		return nil, false
	}
	return o.Contract, true
}

// HasContract returns a boolean if a field has been set.
func (o *ScannerResultContracts) HasContract() bool {
	if o != nil && !IsNil(o.Contract) {
		return true
	}

	return false
}

// SetContract gets a reference to the given []ScannerResultContractsContractInner and assigns it to the Contract field.
func (o *ScannerResultContracts) SetContract(v []ScannerResultContractsContractInner) {
	o.Contract = v
}

func (o ScannerResultContracts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannerResultContracts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contract) {
		toSerialize["Contract"] = o.Contract
	}
	return toSerialize, nil
}

type NullableScannerResultContracts struct {
	value *ScannerResultContracts
	isSet bool
}

func (v NullableScannerResultContracts) Get() *ScannerResultContracts {
	return v.value
}

func (v *NullableScannerResultContracts) Set(val *ScannerResultContracts) {
	v.value = val
	v.isSet = true
}

func (v NullableScannerResultContracts) IsSet() bool {
	return v.isSet
}

func (v *NullableScannerResultContracts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannerResultContracts(val *ScannerResultContracts) *NullableScannerResultContracts {
	return &NullableScannerResultContracts{value: val, isSet: true}
}

func (v NullableScannerResultContracts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannerResultContracts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}