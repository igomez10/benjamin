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

// checks if the StocksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StocksInner{}

// StocksInner future contract information
type StocksInner struct {
	// company name
	Name *string `json:"name,omitempty"`
	// company name in Chinese
	ChineseName *string `json:"chineseName,omitempty"`
	AssetClass  *string `json:"assetClass,omitempty"`
	// array of contracts from different exchanges
	Contracts []StocksInnerContractsInner `json:"contracts,omitempty"`
}

// NewStocksInner instantiates a new StocksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStocksInner() *StocksInner {
	this := StocksInner{}
	return &this
}

// NewStocksInnerWithDefaults instantiates a new StocksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStocksInnerWithDefaults() *StocksInner {
	this := StocksInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StocksInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StocksInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StocksInner) SetName(v string) {
	o.Name = &v
}

// GetChineseName returns the ChineseName field value if set, zero value otherwise.
func (o *StocksInner) GetChineseName() string {
	if o == nil || IsNil(o.ChineseName) {
		var ret string
		return ret
	}
	return *o.ChineseName
}

// GetChineseNameOk returns a tuple with the ChineseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksInner) GetChineseNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChineseName) {
		return nil, false
	}
	return o.ChineseName, true
}

// HasChineseName returns a boolean if a field has been set.
func (o *StocksInner) HasChineseName() bool {
	if o != nil && !IsNil(o.ChineseName) {
		return true
	}

	return false
}

// SetChineseName gets a reference to the given string and assigns it to the ChineseName field.
func (o *StocksInner) SetChineseName(v string) {
	o.ChineseName = &v
}

// GetAssetClass returns the AssetClass field value if set, zero value otherwise.
func (o *StocksInner) GetAssetClass() string {
	if o == nil || IsNil(o.AssetClass) {
		var ret string
		return ret
	}
	return *o.AssetClass
}

// GetAssetClassOk returns a tuple with the AssetClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksInner) GetAssetClassOk() (*string, bool) {
	if o == nil || IsNil(o.AssetClass) {
		return nil, false
	}
	return o.AssetClass, true
}

// HasAssetClass returns a boolean if a field has been set.
func (o *StocksInner) HasAssetClass() bool {
	if o != nil && !IsNil(o.AssetClass) {
		return true
	}

	return false
}

// SetAssetClass gets a reference to the given string and assigns it to the AssetClass field.
func (o *StocksInner) SetAssetClass(v string) {
	o.AssetClass = &v
}

// GetContracts returns the Contracts field value if set, zero value otherwise.
func (o *StocksInner) GetContracts() []StocksInnerContractsInner {
	if o == nil || IsNil(o.Contracts) {
		var ret []StocksInnerContractsInner
		return ret
	}
	return o.Contracts
}

// GetContractsOk returns a tuple with the Contracts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StocksInner) GetContractsOk() ([]StocksInnerContractsInner, bool) {
	if o == nil || IsNil(o.Contracts) {
		return nil, false
	}
	return o.Contracts, true
}

// HasContracts returns a boolean if a field has been set.
func (o *StocksInner) HasContracts() bool {
	if o != nil && !IsNil(o.Contracts) {
		return true
	}

	return false
}

// SetContracts gets a reference to the given []StocksInnerContractsInner and assigns it to the Contracts field.
func (o *StocksInner) SetContracts(v []StocksInnerContractsInner) {
	o.Contracts = v
}

func (o StocksInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StocksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ChineseName) {
		toSerialize["chineseName"] = o.ChineseName
	}
	if !IsNil(o.AssetClass) {
		toSerialize["assetClass"] = o.AssetClass
	}
	if !IsNil(o.Contracts) {
		toSerialize["contracts"] = o.Contracts
	}
	return toSerialize, nil
}

type NullableStocksInner struct {
	value *StocksInner
	isSet bool
}

func (v NullableStocksInner) Get() *StocksInner {
	return v.value
}

func (v *NullableStocksInner) Set(val *StocksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStocksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStocksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStocksInner(val *StocksInner) *NullableStocksInner {
	return &NullableStocksInner{value: val, isSet: true}
}

func (v NullableStocksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStocksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}