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

// checks if the Ledger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ledger{}

// Ledger struct for Ledger
type Ledger struct {
	Commoditymarketvalue      *float32 `json:"commoditymarketvalue,omitempty"`
	Futuremarketvalue         *float32 `json:"futuremarketvalue,omitempty"`
	Settledcash               *float32 `json:"settledcash,omitempty"`
	Exchangerate              *float32 `json:"exchangerate,omitempty"`
	Sessionid                 *int32   `json:"sessionid,omitempty"`
	Cashbalance               *float32 `json:"cashbalance,omitempty"`
	Corporatebondsmarketvalue *float32 `json:"corporatebondsmarketvalue,omitempty"`
	Warrantsmarketvalue       *float32 `json:"warrantsmarketvalue,omitempty"`
	Netliquidationvalue       *float32 `json:"netliquidationvalue,omitempty"`
	Interest                  *float32 `json:"interest,omitempty"`
	Unrealizedpnl             *float32 `json:"unrealizedpnl,omitempty"`
	Stockmarketvalue          *float32 `json:"stockmarketvalue,omitempty"`
	Moneyfunds                *float32 `json:"moneyfunds,omitempty"`
	Currency                  *string  `json:"currency,omitempty"`
	Realizedpnl               *float32 `json:"realizedpnl,omitempty"`
	Funds                     *float32 `json:"funds,omitempty"`
	Acctcode                  *string  `json:"acctcode,omitempty"`
	Issueroptionsmarketvalue  *float32 `json:"issueroptionsmarketvalue,omitempty"`
	Key                       *string  `json:"key,omitempty"`
	Timestamp                 *int32   `json:"timestamp,omitempty"`
	Severity                  *int32   `json:"severity,omitempty"`
}

// NewLedger instantiates a new Ledger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedger() *Ledger {
	this := Ledger{}
	return &this
}

// NewLedgerWithDefaults instantiates a new Ledger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerWithDefaults() *Ledger {
	this := Ledger{}
	return &this
}

// GetCommoditymarketvalue returns the Commoditymarketvalue field value if set, zero value otherwise.
func (o *Ledger) GetCommoditymarketvalue() float32 {
	if o == nil || IsNil(o.Commoditymarketvalue) {
		var ret float32
		return ret
	}
	return *o.Commoditymarketvalue
}

// GetCommoditymarketvalueOk returns a tuple with the Commoditymarketvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetCommoditymarketvalueOk() (*float32, bool) {
	if o == nil || IsNil(o.Commoditymarketvalue) {
		return nil, false
	}
	return o.Commoditymarketvalue, true
}

// HasCommoditymarketvalue returns a boolean if a field has been set.
func (o *Ledger) HasCommoditymarketvalue() bool {
	if o != nil && !IsNil(o.Commoditymarketvalue) {
		return true
	}

	return false
}

// SetCommoditymarketvalue gets a reference to the given float32 and assigns it to the Commoditymarketvalue field.
func (o *Ledger) SetCommoditymarketvalue(v float32) {
	o.Commoditymarketvalue = &v
}

// GetFuturemarketvalue returns the Futuremarketvalue field value if set, zero value otherwise.
func (o *Ledger) GetFuturemarketvalue() float32 {
	if o == nil || IsNil(o.Futuremarketvalue) {
		var ret float32
		return ret
	}
	return *o.Futuremarketvalue
}

// GetFuturemarketvalueOk returns a tuple with the Futuremarketvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetFuturemarketvalueOk() (*float32, bool) {
	if o == nil || IsNil(o.Futuremarketvalue) {
		return nil, false
	}
	return o.Futuremarketvalue, true
}

// HasFuturemarketvalue returns a boolean if a field has been set.
func (o *Ledger) HasFuturemarketvalue() bool {
	if o != nil && !IsNil(o.Futuremarketvalue) {
		return true
	}

	return false
}

// SetFuturemarketvalue gets a reference to the given float32 and assigns it to the Futuremarketvalue field.
func (o *Ledger) SetFuturemarketvalue(v float32) {
	o.Futuremarketvalue = &v
}

// GetSettledcash returns the Settledcash field value if set, zero value otherwise.
func (o *Ledger) GetSettledcash() float32 {
	if o == nil || IsNil(o.Settledcash) {
		var ret float32
		return ret
	}
	return *o.Settledcash
}

// GetSettledcashOk returns a tuple with the Settledcash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetSettledcashOk() (*float32, bool) {
	if o == nil || IsNil(o.Settledcash) {
		return nil, false
	}
	return o.Settledcash, true
}

// HasSettledcash returns a boolean if a field has been set.
func (o *Ledger) HasSettledcash() bool {
	if o != nil && !IsNil(o.Settledcash) {
		return true
	}

	return false
}

// SetSettledcash gets a reference to the given float32 and assigns it to the Settledcash field.
func (o *Ledger) SetSettledcash(v float32) {
	o.Settledcash = &v
}

// GetExchangerate returns the Exchangerate field value if set, zero value otherwise.
func (o *Ledger) GetExchangerate() float32 {
	if o == nil || IsNil(o.Exchangerate) {
		var ret float32
		return ret
	}
	return *o.Exchangerate
}

// GetExchangerateOk returns a tuple with the Exchangerate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetExchangerateOk() (*float32, bool) {
	if o == nil || IsNil(o.Exchangerate) {
		return nil, false
	}
	return o.Exchangerate, true
}

// HasExchangerate returns a boolean if a field has been set.
func (o *Ledger) HasExchangerate() bool {
	if o != nil && !IsNil(o.Exchangerate) {
		return true
	}

	return false
}

// SetExchangerate gets a reference to the given float32 and assigns it to the Exchangerate field.
func (o *Ledger) SetExchangerate(v float32) {
	o.Exchangerate = &v
}

// GetSessionid returns the Sessionid field value if set, zero value otherwise.
func (o *Ledger) GetSessionid() int32 {
	if o == nil || IsNil(o.Sessionid) {
		var ret int32
		return ret
	}
	return *o.Sessionid
}

// GetSessionidOk returns a tuple with the Sessionid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetSessionidOk() (*int32, bool) {
	if o == nil || IsNil(o.Sessionid) {
		return nil, false
	}
	return o.Sessionid, true
}

// HasSessionid returns a boolean if a field has been set.
func (o *Ledger) HasSessionid() bool {
	if o != nil && !IsNil(o.Sessionid) {
		return true
	}

	return false
}

// SetSessionid gets a reference to the given int32 and assigns it to the Sessionid field.
func (o *Ledger) SetSessionid(v int32) {
	o.Sessionid = &v
}

// GetCashbalance returns the Cashbalance field value if set, zero value otherwise.
func (o *Ledger) GetCashbalance() float32 {
	if o == nil || IsNil(o.Cashbalance) {
		var ret float32
		return ret
	}
	return *o.Cashbalance
}

// GetCashbalanceOk returns a tuple with the Cashbalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetCashbalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Cashbalance) {
		return nil, false
	}
	return o.Cashbalance, true
}

// HasCashbalance returns a boolean if a field has been set.
func (o *Ledger) HasCashbalance() bool {
	if o != nil && !IsNil(o.Cashbalance) {
		return true
	}

	return false
}

// SetCashbalance gets a reference to the given float32 and assigns it to the Cashbalance field.
func (o *Ledger) SetCashbalance(v float32) {
	o.Cashbalance = &v
}

// GetCorporatebondsmarketvalue returns the Corporatebondsmarketvalue field value if set, zero value otherwise.
func (o *Ledger) GetCorporatebondsmarketvalue() float32 {
	if o == nil || IsNil(o.Corporatebondsmarketvalue) {
		var ret float32
		return ret
	}
	return *o.Corporatebondsmarketvalue
}

// GetCorporatebondsmarketvalueOk returns a tuple with the Corporatebondsmarketvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetCorporatebondsmarketvalueOk() (*float32, bool) {
	if o == nil || IsNil(o.Corporatebondsmarketvalue) {
		return nil, false
	}
	return o.Corporatebondsmarketvalue, true
}

// HasCorporatebondsmarketvalue returns a boolean if a field has been set.
func (o *Ledger) HasCorporatebondsmarketvalue() bool {
	if o != nil && !IsNil(o.Corporatebondsmarketvalue) {
		return true
	}

	return false
}

// SetCorporatebondsmarketvalue gets a reference to the given float32 and assigns it to the Corporatebondsmarketvalue field.
func (o *Ledger) SetCorporatebondsmarketvalue(v float32) {
	o.Corporatebondsmarketvalue = &v
}

// GetWarrantsmarketvalue returns the Warrantsmarketvalue field value if set, zero value otherwise.
func (o *Ledger) GetWarrantsmarketvalue() float32 {
	if o == nil || IsNil(o.Warrantsmarketvalue) {
		var ret float32
		return ret
	}
	return *o.Warrantsmarketvalue
}

// GetWarrantsmarketvalueOk returns a tuple with the Warrantsmarketvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetWarrantsmarketvalueOk() (*float32, bool) {
	if o == nil || IsNil(o.Warrantsmarketvalue) {
		return nil, false
	}
	return o.Warrantsmarketvalue, true
}

// HasWarrantsmarketvalue returns a boolean if a field has been set.
func (o *Ledger) HasWarrantsmarketvalue() bool {
	if o != nil && !IsNil(o.Warrantsmarketvalue) {
		return true
	}

	return false
}

// SetWarrantsmarketvalue gets a reference to the given float32 and assigns it to the Warrantsmarketvalue field.
func (o *Ledger) SetWarrantsmarketvalue(v float32) {
	o.Warrantsmarketvalue = &v
}

// GetNetliquidationvalue returns the Netliquidationvalue field value if set, zero value otherwise.
func (o *Ledger) GetNetliquidationvalue() float32 {
	if o == nil || IsNil(o.Netliquidationvalue) {
		var ret float32
		return ret
	}
	return *o.Netliquidationvalue
}

// GetNetliquidationvalueOk returns a tuple with the Netliquidationvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetNetliquidationvalueOk() (*float32, bool) {
	if o == nil || IsNil(o.Netliquidationvalue) {
		return nil, false
	}
	return o.Netliquidationvalue, true
}

// HasNetliquidationvalue returns a boolean if a field has been set.
func (o *Ledger) HasNetliquidationvalue() bool {
	if o != nil && !IsNil(o.Netliquidationvalue) {
		return true
	}

	return false
}

// SetNetliquidationvalue gets a reference to the given float32 and assigns it to the Netliquidationvalue field.
func (o *Ledger) SetNetliquidationvalue(v float32) {
	o.Netliquidationvalue = &v
}

// GetInterest returns the Interest field value if set, zero value otherwise.
func (o *Ledger) GetInterest() float32 {
	if o == nil || IsNil(o.Interest) {
		var ret float32
		return ret
	}
	return *o.Interest
}

// GetInterestOk returns a tuple with the Interest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetInterestOk() (*float32, bool) {
	if o == nil || IsNil(o.Interest) {
		return nil, false
	}
	return o.Interest, true
}

// HasInterest returns a boolean if a field has been set.
func (o *Ledger) HasInterest() bool {
	if o != nil && !IsNil(o.Interest) {
		return true
	}

	return false
}

// SetInterest gets a reference to the given float32 and assigns it to the Interest field.
func (o *Ledger) SetInterest(v float32) {
	o.Interest = &v
}

// GetUnrealizedpnl returns the Unrealizedpnl field value if set, zero value otherwise.
func (o *Ledger) GetUnrealizedpnl() float32 {
	if o == nil || IsNil(o.Unrealizedpnl) {
		var ret float32
		return ret
	}
	return *o.Unrealizedpnl
}

// GetUnrealizedpnlOk returns a tuple with the Unrealizedpnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetUnrealizedpnlOk() (*float32, bool) {
	if o == nil || IsNil(o.Unrealizedpnl) {
		return nil, false
	}
	return o.Unrealizedpnl, true
}

// HasUnrealizedpnl returns a boolean if a field has been set.
func (o *Ledger) HasUnrealizedpnl() bool {
	if o != nil && !IsNil(o.Unrealizedpnl) {
		return true
	}

	return false
}

// SetUnrealizedpnl gets a reference to the given float32 and assigns it to the Unrealizedpnl field.
func (o *Ledger) SetUnrealizedpnl(v float32) {
	o.Unrealizedpnl = &v
}

// GetStockmarketvalue returns the Stockmarketvalue field value if set, zero value otherwise.
func (o *Ledger) GetStockmarketvalue() float32 {
	if o == nil || IsNil(o.Stockmarketvalue) {
		var ret float32
		return ret
	}
	return *o.Stockmarketvalue
}

// GetStockmarketvalueOk returns a tuple with the Stockmarketvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetStockmarketvalueOk() (*float32, bool) {
	if o == nil || IsNil(o.Stockmarketvalue) {
		return nil, false
	}
	return o.Stockmarketvalue, true
}

// HasStockmarketvalue returns a boolean if a field has been set.
func (o *Ledger) HasStockmarketvalue() bool {
	if o != nil && !IsNil(o.Stockmarketvalue) {
		return true
	}

	return false
}

// SetStockmarketvalue gets a reference to the given float32 and assigns it to the Stockmarketvalue field.
func (o *Ledger) SetStockmarketvalue(v float32) {
	o.Stockmarketvalue = &v
}

// GetMoneyfunds returns the Moneyfunds field value if set, zero value otherwise.
func (o *Ledger) GetMoneyfunds() float32 {
	if o == nil || IsNil(o.Moneyfunds) {
		var ret float32
		return ret
	}
	return *o.Moneyfunds
}

// GetMoneyfundsOk returns a tuple with the Moneyfunds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetMoneyfundsOk() (*float32, bool) {
	if o == nil || IsNil(o.Moneyfunds) {
		return nil, false
	}
	return o.Moneyfunds, true
}

// HasMoneyfunds returns a boolean if a field has been set.
func (o *Ledger) HasMoneyfunds() bool {
	if o != nil && !IsNil(o.Moneyfunds) {
		return true
	}

	return false
}

// SetMoneyfunds gets a reference to the given float32 and assigns it to the Moneyfunds field.
func (o *Ledger) SetMoneyfunds(v float32) {
	o.Moneyfunds = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Ledger) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *Ledger) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Ledger) SetCurrency(v string) {
	o.Currency = &v
}

// GetRealizedpnl returns the Realizedpnl field value if set, zero value otherwise.
func (o *Ledger) GetRealizedpnl() float32 {
	if o == nil || IsNil(o.Realizedpnl) {
		var ret float32
		return ret
	}
	return *o.Realizedpnl
}

// GetRealizedpnlOk returns a tuple with the Realizedpnl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetRealizedpnlOk() (*float32, bool) {
	if o == nil || IsNil(o.Realizedpnl) {
		return nil, false
	}
	return o.Realizedpnl, true
}

// HasRealizedpnl returns a boolean if a field has been set.
func (o *Ledger) HasRealizedpnl() bool {
	if o != nil && !IsNil(o.Realizedpnl) {
		return true
	}

	return false
}

// SetRealizedpnl gets a reference to the given float32 and assigns it to the Realizedpnl field.
func (o *Ledger) SetRealizedpnl(v float32) {
	o.Realizedpnl = &v
}

// GetFunds returns the Funds field value if set, zero value otherwise.
func (o *Ledger) GetFunds() float32 {
	if o == nil || IsNil(o.Funds) {
		var ret float32
		return ret
	}
	return *o.Funds
}

// GetFundsOk returns a tuple with the Funds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetFundsOk() (*float32, bool) {
	if o == nil || IsNil(o.Funds) {
		return nil, false
	}
	return o.Funds, true
}

// HasFunds returns a boolean if a field has been set.
func (o *Ledger) HasFunds() bool {
	if o != nil && !IsNil(o.Funds) {
		return true
	}

	return false
}

// SetFunds gets a reference to the given float32 and assigns it to the Funds field.
func (o *Ledger) SetFunds(v float32) {
	o.Funds = &v
}

// GetAcctcode returns the Acctcode field value if set, zero value otherwise.
func (o *Ledger) GetAcctcode() string {
	if o == nil || IsNil(o.Acctcode) {
		var ret string
		return ret
	}
	return *o.Acctcode
}

// GetAcctcodeOk returns a tuple with the Acctcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetAcctcodeOk() (*string, bool) {
	if o == nil || IsNil(o.Acctcode) {
		return nil, false
	}
	return o.Acctcode, true
}

// HasAcctcode returns a boolean if a field has been set.
func (o *Ledger) HasAcctcode() bool {
	if o != nil && !IsNil(o.Acctcode) {
		return true
	}

	return false
}

// SetAcctcode gets a reference to the given string and assigns it to the Acctcode field.
func (o *Ledger) SetAcctcode(v string) {
	o.Acctcode = &v
}

// GetIssueroptionsmarketvalue returns the Issueroptionsmarketvalue field value if set, zero value otherwise.
func (o *Ledger) GetIssueroptionsmarketvalue() float32 {
	if o == nil || IsNil(o.Issueroptionsmarketvalue) {
		var ret float32
		return ret
	}
	return *o.Issueroptionsmarketvalue
}

// GetIssueroptionsmarketvalueOk returns a tuple with the Issueroptionsmarketvalue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetIssueroptionsmarketvalueOk() (*float32, bool) {
	if o == nil || IsNil(o.Issueroptionsmarketvalue) {
		return nil, false
	}
	return o.Issueroptionsmarketvalue, true
}

// HasIssueroptionsmarketvalue returns a boolean if a field has been set.
func (o *Ledger) HasIssueroptionsmarketvalue() bool {
	if o != nil && !IsNil(o.Issueroptionsmarketvalue) {
		return true
	}

	return false
}

// SetIssueroptionsmarketvalue gets a reference to the given float32 and assigns it to the Issueroptionsmarketvalue field.
func (o *Ledger) SetIssueroptionsmarketvalue(v float32) {
	o.Issueroptionsmarketvalue = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Ledger) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Ledger) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Ledger) SetKey(v string) {
	o.Key = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Ledger) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Ledger) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *Ledger) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *Ledger) GetSeverity() int32 {
	if o == nil || IsNil(o.Severity) {
		var ret int32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ledger) GetSeverityOk() (*int32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *Ledger) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given int32 and assigns it to the Severity field.
func (o *Ledger) SetSeverity(v int32) {
	o.Severity = &v
}

func (o Ledger) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ledger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Commoditymarketvalue) {
		toSerialize["commoditymarketvalue"] = o.Commoditymarketvalue
	}
	if !IsNil(o.Futuremarketvalue) {
		toSerialize["futuremarketvalue"] = o.Futuremarketvalue
	}
	if !IsNil(o.Settledcash) {
		toSerialize["settledcash"] = o.Settledcash
	}
	if !IsNil(o.Exchangerate) {
		toSerialize["exchangerate"] = o.Exchangerate
	}
	if !IsNil(o.Sessionid) {
		toSerialize["sessionid"] = o.Sessionid
	}
	if !IsNil(o.Cashbalance) {
		toSerialize["cashbalance"] = o.Cashbalance
	}
	if !IsNil(o.Corporatebondsmarketvalue) {
		toSerialize["corporatebondsmarketvalue"] = o.Corporatebondsmarketvalue
	}
	if !IsNil(o.Warrantsmarketvalue) {
		toSerialize["warrantsmarketvalue"] = o.Warrantsmarketvalue
	}
	if !IsNil(o.Netliquidationvalue) {
		toSerialize["netliquidationvalue"] = o.Netliquidationvalue
	}
	if !IsNil(o.Interest) {
		toSerialize["interest"] = o.Interest
	}
	if !IsNil(o.Unrealizedpnl) {
		toSerialize["unrealizedpnl"] = o.Unrealizedpnl
	}
	if !IsNil(o.Stockmarketvalue) {
		toSerialize["stockmarketvalue"] = o.Stockmarketvalue
	}
	if !IsNil(o.Moneyfunds) {
		toSerialize["moneyfunds"] = o.Moneyfunds
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Realizedpnl) {
		toSerialize["realizedpnl"] = o.Realizedpnl
	}
	if !IsNil(o.Funds) {
		toSerialize["funds"] = o.Funds
	}
	if !IsNil(o.Acctcode) {
		toSerialize["acctcode"] = o.Acctcode
	}
	if !IsNil(o.Issueroptionsmarketvalue) {
		toSerialize["issueroptionsmarketvalue"] = o.Issueroptionsmarketvalue
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	return toSerialize, nil
}

type NullableLedger struct {
	value *Ledger
	isSet bool
}

func (v NullableLedger) Get() *Ledger {
	return v.value
}

func (v *NullableLedger) Set(val *Ledger) {
	v.value = val
	v.isSet = true
}

func (v NullableLedger) IsSet() bool {
	return v.isSet
}

func (v *NullableLedger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedger(val *Ledger) *NullableLedger {
	return &NullableLedger{value: val, isSet: true}
}

func (v NullableLedger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
