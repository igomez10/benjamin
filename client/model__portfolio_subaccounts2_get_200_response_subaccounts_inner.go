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

// checks if the PortfolioSubaccounts2Get200ResponseSubaccountsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortfolioSubaccounts2Get200ResponseSubaccountsInner{}

// PortfolioSubaccounts2Get200ResponseSubaccountsInner Account information
type PortfolioSubaccounts2Get200ResponseSubaccountsInner struct {
	// The account identification value
	Id *string `json:"id,omitempty"`
	// The account number
	AccountId *string `json:"accountId,omitempty"`
	// The accountAlias
	AccountVan *string `json:"accountVan,omitempty"`
	// Title of the account
	AccountTitle *string `json:"accountTitle,omitempty"`
	// Whichever value is not null in this priority
	DisplayName *string `json:"displayName,omitempty"`
	// User customizable account alias. Refer to [Configure Account Alias](https://guides.interactivebrokers.com/cp/cp.htm#am/settings/accountalias.htm) for details.
	AccountAlias *string `json:"accountAlias,omitempty"`
	// When the account was opened in unix time.
	AccountStatus *float32 `json:"accountStatus,omitempty"`
	// Base currency of the account.
	Currency *string `json:"currency,omitempty"`
	// Account Type
	Type *string `json:"type,omitempty"`
	// UNI - Deprecated property
	TradingType *string `json:"tradingType,omitempty"`
	// If an account is a sub-account to a Financial Advisor.
	Faclient *bool `json:"faclient,omitempty"`
	// Status of the Account   * O = Open   * P or N = Pending   * A = Abandoned   * R = Rejected   * C = Closed   covestor:     type: boolean     description: Is a Covestor Account   parent:     type: object     properties:       mmc:         type: array         items:           type: string           description: Money Manager Client (MMC) Account       accountId:         type: string         description: Account Number for Money Manager Client       isMParent:         type: boolean         description: Is MM a Parent Account       isMChild:         type: boolean         description: Is MM a Child Account       isMultiplex:         type: boolean         description: Is a Multiplex Account. These are account models with individual account being parent and managed account being child.   desc:     type: string     description: Formatted \"accountId - accountAlias\"
	ClearingStatus *string `json:"clearingStatus,omitempty"`
}

// NewPortfolioSubaccounts2Get200ResponseSubaccountsInner instantiates a new PortfolioSubaccounts2Get200ResponseSubaccountsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortfolioSubaccounts2Get200ResponseSubaccountsInner() *PortfolioSubaccounts2Get200ResponseSubaccountsInner {
	this := PortfolioSubaccounts2Get200ResponseSubaccountsInner{}
	return &this
}

// NewPortfolioSubaccounts2Get200ResponseSubaccountsInnerWithDefaults instantiates a new PortfolioSubaccounts2Get200ResponseSubaccountsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortfolioSubaccounts2Get200ResponseSubaccountsInnerWithDefaults() *PortfolioSubaccounts2Get200ResponseSubaccountsInner {
	this := PortfolioSubaccounts2Get200ResponseSubaccountsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetId(v string) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountVan returns the AccountVan field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountVan() string {
	if o == nil || IsNil(o.AccountVan) {
		var ret string
		return ret
	}
	return *o.AccountVan
}

// GetAccountVanOk returns a tuple with the AccountVan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountVanOk() (*string, bool) {
	if o == nil || IsNil(o.AccountVan) {
		return nil, false
	}
	return o.AccountVan, true
}

// HasAccountVan returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountVan() bool {
	if o != nil && !IsNil(o.AccountVan) {
		return true
	}

	return false
}

// SetAccountVan gets a reference to the given string and assigns it to the AccountVan field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountVan(v string) {
	o.AccountVan = &v
}

// GetAccountTitle returns the AccountTitle field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountTitle() string {
	if o == nil || IsNil(o.AccountTitle) {
		var ret string
		return ret
	}
	return *o.AccountTitle
}

// GetAccountTitleOk returns a tuple with the AccountTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountTitleOk() (*string, bool) {
	if o == nil || IsNil(o.AccountTitle) {
		return nil, false
	}
	return o.AccountTitle, true
}

// HasAccountTitle returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountTitle() bool {
	if o != nil && !IsNil(o.AccountTitle) {
		return true
	}

	return false
}

// SetAccountTitle gets a reference to the given string and assigns it to the AccountTitle field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountTitle(v string) {
	o.AccountTitle = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAccountAlias returns the AccountAlias field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountAlias() string {
	if o == nil || IsNil(o.AccountAlias) {
		var ret string
		return ret
	}
	return *o.AccountAlias
}

// GetAccountAliasOk returns a tuple with the AccountAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountAliasOk() (*string, bool) {
	if o == nil || IsNil(o.AccountAlias) {
		return nil, false
	}
	return o.AccountAlias, true
}

// HasAccountAlias returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountAlias() bool {
	if o != nil && !IsNil(o.AccountAlias) {
		return true
	}

	return false
}

// SetAccountAlias gets a reference to the given string and assigns it to the AccountAlias field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountAlias(v string) {
	o.AccountAlias = &v
}

// GetAccountStatus returns the AccountStatus field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountStatus() float32 {
	if o == nil || IsNil(o.AccountStatus) {
		var ret float32
		return ret
	}
	return *o.AccountStatus
}

// GetAccountStatusOk returns a tuple with the AccountStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountStatusOk() (*float32, bool) {
	if o == nil || IsNil(o.AccountStatus) {
		return nil, false
	}
	return o.AccountStatus, true
}

// HasAccountStatus returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountStatus() bool {
	if o != nil && !IsNil(o.AccountStatus) {
		return true
	}

	return false
}

// SetAccountStatus gets a reference to the given float32 and assigns it to the AccountStatus field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountStatus(v float32) {
	o.AccountStatus = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetType(v string) {
	o.Type = &v
}

// GetTradingType returns the TradingType field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetTradingType() string {
	if o == nil || IsNil(o.TradingType) {
		var ret string
		return ret
	}
	return *o.TradingType
}

// GetTradingTypeOk returns a tuple with the TradingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetTradingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TradingType) {
		return nil, false
	}
	return o.TradingType, true
}

// HasTradingType returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasTradingType() bool {
	if o != nil && !IsNil(o.TradingType) {
		return true
	}

	return false
}

// SetTradingType gets a reference to the given string and assigns it to the TradingType field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetTradingType(v string) {
	o.TradingType = &v
}

// GetFaclient returns the Faclient field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetFaclient() bool {
	if o == nil || IsNil(o.Faclient) {
		var ret bool
		return ret
	}
	return *o.Faclient
}

// GetFaclientOk returns a tuple with the Faclient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetFaclientOk() (*bool, bool) {
	if o == nil || IsNil(o.Faclient) {
		return nil, false
	}
	return o.Faclient, true
}

// HasFaclient returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasFaclient() bool {
	if o != nil && !IsNil(o.Faclient) {
		return true
	}

	return false
}

// SetFaclient gets a reference to the given bool and assigns it to the Faclient field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetFaclient(v bool) {
	o.Faclient = &v
}

// GetClearingStatus returns the ClearingStatus field value if set, zero value otherwise.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetClearingStatus() string {
	if o == nil || IsNil(o.ClearingStatus) {
		var ret string
		return ret
	}
	return *o.ClearingStatus
}

// GetClearingStatusOk returns a tuple with the ClearingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetClearingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ClearingStatus) {
		return nil, false
	}
	return o.ClearingStatus, true
}

// HasClearingStatus returns a boolean if a field has been set.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasClearingStatus() bool {
	if o != nil && !IsNil(o.ClearingStatus) {
		return true
	}

	return false
}

// SetClearingStatus gets a reference to the given string and assigns it to the ClearingStatus field.
func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetClearingStatus(v string) {
	o.ClearingStatus = &v
}

func (o PortfolioSubaccounts2Get200ResponseSubaccountsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortfolioSubaccounts2Get200ResponseSubaccountsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.AccountVan) {
		toSerialize["accountVan"] = o.AccountVan
	}
	if !IsNil(o.AccountTitle) {
		toSerialize["accountTitle"] = o.AccountTitle
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.AccountAlias) {
		toSerialize["accountAlias"] = o.AccountAlias
	}
	if !IsNil(o.AccountStatus) {
		toSerialize["accountStatus"] = o.AccountStatus
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TradingType) {
		toSerialize["tradingType"] = o.TradingType
	}
	if !IsNil(o.Faclient) {
		toSerialize["faclient"] = o.Faclient
	}
	if !IsNil(o.ClearingStatus) {
		toSerialize["clearingStatus"] = o.ClearingStatus
	}
	return toSerialize, nil
}

type NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner struct {
	value *PortfolioSubaccounts2Get200ResponseSubaccountsInner
	isSet bool
}

func (v NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner) Get() *PortfolioSubaccounts2Get200ResponseSubaccountsInner {
	return v.value
}

func (v *NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner) Set(val *PortfolioSubaccounts2Get200ResponseSubaccountsInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortfolioSubaccounts2Get200ResponseSubaccountsInner(val *PortfolioSubaccounts2Get200ResponseSubaccountsInner) *NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner {
	return &NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner{value: val, isSet: true}
}

func (v NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortfolioSubaccounts2Get200ResponseSubaccountsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}