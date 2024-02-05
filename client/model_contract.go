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

// checks if the Contract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Contract{}

// Contract Contains all details of the contract, including rules you can use when placing orders
type Contract struct {
	// true means you can trade outside RTH(regular trading hours)
	RTH *bool `json:"r_t_h,omitempty"`
	// same as that in request
	ConId *string `json:"con_id,omitempty"`
	// Contracts company name
	CompanyName *string `json:"company_name,omitempty"`
	Exchange    *string `json:"exchange,omitempty"`
	// for exmple FB
	LocalSymbol *string `json:"local_symbol,omitempty"`
	// for example STK
	InstrumentType *string        `json:"instrument_type,omitempty"`
	Currency       *string        `json:"currency,omitempty"`
	Category       *string        `json:"category,omitempty"`
	Industry       *string        `json:"industry,omitempty"`
	Rules          *ContractRules `json:"rules,omitempty"`
}

// NewContract instantiates a new Contract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContract() *Contract {
	this := Contract{}
	return &this
}

// NewContractWithDefaults instantiates a new Contract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractWithDefaults() *Contract {
	this := Contract{}
	return &this
}

// GetRTH returns the RTH field value if set, zero value otherwise.
func (o *Contract) GetRTH() bool {
	if o == nil || IsNil(o.RTH) {
		var ret bool
		return ret
	}
	return *o.RTH
}

// GetRTHOk returns a tuple with the RTH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetRTHOk() (*bool, bool) {
	if o == nil || IsNil(o.RTH) {
		return nil, false
	}
	return o.RTH, true
}

// HasRTH returns a boolean if a field has been set.
func (o *Contract) HasRTH() bool {
	if o != nil && !IsNil(o.RTH) {
		return true
	}

	return false
}

// SetRTH gets a reference to the given bool and assigns it to the RTH field.
func (o *Contract) SetRTH(v bool) {
	o.RTH = &v
}

// GetConId returns the ConId field value if set, zero value otherwise.
func (o *Contract) GetConId() string {
	if o == nil || IsNil(o.ConId) {
		var ret string
		return ret
	}
	return *o.ConId
}

// GetConIdOk returns a tuple with the ConId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetConIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConId) {
		return nil, false
	}
	return o.ConId, true
}

// HasConId returns a boolean if a field has been set.
func (o *Contract) HasConId() bool {
	if o != nil && !IsNil(o.ConId) {
		return true
	}

	return false
}

// SetConId gets a reference to the given string and assigns it to the ConId field.
func (o *Contract) SetConId(v string) {
	o.ConId = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *Contract) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *Contract) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *Contract) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *Contract) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *Contract) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *Contract) SetExchange(v string) {
	o.Exchange = &v
}

// GetLocalSymbol returns the LocalSymbol field value if set, zero value otherwise.
func (o *Contract) GetLocalSymbol() string {
	if o == nil || IsNil(o.LocalSymbol) {
		var ret string
		return ret
	}
	return *o.LocalSymbol
}

// GetLocalSymbolOk returns a tuple with the LocalSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetLocalSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.LocalSymbol) {
		return nil, false
	}
	return o.LocalSymbol, true
}

// HasLocalSymbol returns a boolean if a field has been set.
func (o *Contract) HasLocalSymbol() bool {
	if o != nil && !IsNil(o.LocalSymbol) {
		return true
	}

	return false
}

// SetLocalSymbol gets a reference to the given string and assigns it to the LocalSymbol field.
func (o *Contract) SetLocalSymbol(v string) {
	o.LocalSymbol = &v
}

// GetInstrumentType returns the InstrumentType field value if set, zero value otherwise.
func (o *Contract) GetInstrumentType() string {
	if o == nil || IsNil(o.InstrumentType) {
		var ret string
		return ret
	}
	return *o.InstrumentType
}

// GetInstrumentTypeOk returns a tuple with the InstrumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetInstrumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.InstrumentType) {
		return nil, false
	}
	return o.InstrumentType, true
}

// HasInstrumentType returns a boolean if a field has been set.
func (o *Contract) HasInstrumentType() bool {
	if o != nil && !IsNil(o.InstrumentType) {
		return true
	}

	return false
}

// SetInstrumentType gets a reference to the given string and assigns it to the InstrumentType field.
func (o *Contract) SetInstrumentType(v string) {
	o.InstrumentType = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Contract) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Contract) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Contract) SetCurrency(v string) {
	o.Currency = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *Contract) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *Contract) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *Contract) SetCategory(v string) {
	o.Category = &v
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
func (o *Contract) GetIndustry() string {
	if o == nil || IsNil(o.Industry) {
		var ret string
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetIndustryOk() (*string, bool) {
	if o == nil || IsNil(o.Industry) {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *Contract) HasIndustry() bool {
	if o != nil && !IsNil(o.Industry) {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given string and assigns it to the Industry field.
func (o *Contract) SetIndustry(v string) {
	o.Industry = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *Contract) GetRules() ContractRules {
	if o == nil || IsNil(o.Rules) {
		var ret ContractRules
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetRulesOk() (*ContractRules, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *Contract) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given ContractRules and assigns it to the Rules field.
func (o *Contract) SetRules(v ContractRules) {
	o.Rules = &v
}

func (o Contract) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Contract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RTH) {
		toSerialize["r_t_h"] = o.RTH
	}
	if !IsNil(o.ConId) {
		toSerialize["con_id"] = o.ConId
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.LocalSymbol) {
		toSerialize["local_symbol"] = o.LocalSymbol
	}
	if !IsNil(o.InstrumentType) {
		toSerialize["instrument_type"] = o.InstrumentType
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Industry) {
		toSerialize["industry"] = o.Industry
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableContract struct {
	value *Contract
	isSet bool
}

func (v NullableContract) Get() *Contract {
	return v.value
}

func (v *NullableContract) Set(val *Contract) {
	v.value = val
	v.isSet = true
}

func (v NullableContract) IsSet() bool {
	return v.isSet
}

func (v *NullableContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContract(val *Contract) *NullableContract {
	return &NullableContract{value: val, isSet: true}
}

func (v NullableContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}