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

// checks if the GetStocks200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStocks200Response{}

// GetStocks200Response struct for GetStocks200Response
type GetStocks200Response struct {
	// This is an array of object(s), there could be multiple results under same symbol
	Symbol []StocksInner `json:"symbol,omitempty"`
}

// NewGetStocks200Response instantiates a new GetStocks200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStocks200Response() *GetStocks200Response {
	this := GetStocks200Response{}
	return &this
}

// NewGetStocks200ResponseWithDefaults instantiates a new GetStocks200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStocks200ResponseWithDefaults() *GetStocks200Response {
	this := GetStocks200Response{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *GetStocks200Response) GetSymbol() []StocksInner {
	if o == nil || IsNil(o.Symbol) {
		var ret []StocksInner
		return ret
	}
	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStocks200Response) GetSymbolOk() ([]StocksInner, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *GetStocks200Response) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given []StocksInner and assigns it to the Symbol field.
func (o *GetStocks200Response) SetSymbol(v []StocksInner) {
	o.Symbol = v
}

func (o GetStocks200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStocks200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	return toSerialize, nil
}

type NullableGetStocks200Response struct {
	value *GetStocks200Response
	isSet bool
}

func (v NullableGetStocks200Response) Get() *GetStocks200Response {
	return v.value
}

func (v *NullableGetStocks200Response) Set(val *GetStocks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStocks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStocks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStocks200Response(val *GetStocks200Response) *NullableGetStocks200Response {
	return &NullableGetStocks200Response{value: val, isSet: true}
}

func (v NullableGetStocks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStocks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}