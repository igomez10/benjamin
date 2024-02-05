# SecdefInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conid** | Pointer to **float32** | IBKR contract identifier | [optional] 
**Symbol** | Pointer to **string** | Underlying symbol | [optional] 
**SecType** | Pointer to **string** | Security type | [optional] 
**Exchange** | Pointer to **string** | Primary Exchange, Routing or Trading Venue | [optional] 
**ListingExchange** | Pointer to **string** | Main Trading Venue | [optional] 
**Right** | Pointer to **string** | Put or Call of the option. C &#x3D; Call Option, P &#x3D; Put Option | [optional] 
**Strike** | Pointer to **float32** | Set price at which a derivative contract can be bought or sold. The strike price also known as exercise price. | [optional] 
**Currency** | Pointer to **string** | Currency the contract trades in | [optional] 
**Cusip** | Pointer to **string** | Committee on Uniform Securities Identification Procedures number | [optional] 
**Coupon** | Pointer to **string** | Annual interest rate paid on a bond | [optional] 
**Desc1** | Pointer to **string** | Currency pairs for Forex e.g. EUR.AUD, EUR.CAD, EUR.CHF etc. | [optional] 
**Desc2** | Pointer to **string** | Formatted expiration, strike and right | [optional] 
**MaturityDate** | Pointer to **float32** | Format YYYYMMDD, the date on which the underlying transaction settles if the option is exercised | [optional] 
**Multiplier** | Pointer to **string** | Multiplier for total premium paid or received for derivative contract. | [optional] 
**TradingClass** | Pointer to **string** | Designation of the contract. | [optional] 
**ValidExchanges** | Pointer to **string** | Comma separated list of exchanges or trading venues. | [optional] 

## Methods

### NewSecdefInfo

`func NewSecdefInfo() *SecdefInfo`

NewSecdefInfo instantiates a new SecdefInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecdefInfoWithDefaults

`func NewSecdefInfoWithDefaults() *SecdefInfo`

NewSecdefInfoWithDefaults instantiates a new SecdefInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConid

`func (o *SecdefInfo) GetConid() float32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *SecdefInfo) GetConidOk() (*float32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *SecdefInfo) SetConid(v float32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *SecdefInfo) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetSymbol

`func (o *SecdefInfo) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SecdefInfo) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SecdefInfo) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SecdefInfo) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSecType

`func (o *SecdefInfo) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *SecdefInfo) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *SecdefInfo) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *SecdefInfo) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetExchange

`func (o *SecdefInfo) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *SecdefInfo) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *SecdefInfo) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *SecdefInfo) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetListingExchange

`func (o *SecdefInfo) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *SecdefInfo) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *SecdefInfo) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *SecdefInfo) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetRight

`func (o *SecdefInfo) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *SecdefInfo) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *SecdefInfo) SetRight(v string)`

SetRight sets Right field to given value.

### HasRight

`func (o *SecdefInfo) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetStrike

`func (o *SecdefInfo) GetStrike() float32`

GetStrike returns the Strike field if non-nil, zero value otherwise.

### GetStrikeOk

`func (o *SecdefInfo) GetStrikeOk() (*float32, bool)`

GetStrikeOk returns a tuple with the Strike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrike

`func (o *SecdefInfo) SetStrike(v float32)`

SetStrike sets Strike field to given value.

### HasStrike

`func (o *SecdefInfo) HasStrike() bool`

HasStrike returns a boolean if a field has been set.

### GetCurrency

`func (o *SecdefInfo) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SecdefInfo) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SecdefInfo) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SecdefInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCusip

`func (o *SecdefInfo) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *SecdefInfo) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *SecdefInfo) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *SecdefInfo) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetCoupon

`func (o *SecdefInfo) GetCoupon() string`

GetCoupon returns the Coupon field if non-nil, zero value otherwise.

### GetCouponOk

`func (o *SecdefInfo) GetCouponOk() (*string, bool)`

GetCouponOk returns a tuple with the Coupon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoupon

`func (o *SecdefInfo) SetCoupon(v string)`

SetCoupon sets Coupon field to given value.

### HasCoupon

`func (o *SecdefInfo) HasCoupon() bool`

HasCoupon returns a boolean if a field has been set.

### GetDesc1

`func (o *SecdefInfo) GetDesc1() string`

GetDesc1 returns the Desc1 field if non-nil, zero value otherwise.

### GetDesc1Ok

`func (o *SecdefInfo) GetDesc1Ok() (*string, bool)`

GetDesc1Ok returns a tuple with the Desc1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc1

`func (o *SecdefInfo) SetDesc1(v string)`

SetDesc1 sets Desc1 field to given value.

### HasDesc1

`func (o *SecdefInfo) HasDesc1() bool`

HasDesc1 returns a boolean if a field has been set.

### GetDesc2

`func (o *SecdefInfo) GetDesc2() string`

GetDesc2 returns the Desc2 field if non-nil, zero value otherwise.

### GetDesc2Ok

`func (o *SecdefInfo) GetDesc2Ok() (*string, bool)`

GetDesc2Ok returns a tuple with the Desc2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc2

`func (o *SecdefInfo) SetDesc2(v string)`

SetDesc2 sets Desc2 field to given value.

### HasDesc2

`func (o *SecdefInfo) HasDesc2() bool`

HasDesc2 returns a boolean if a field has been set.

### GetMaturityDate

`func (o *SecdefInfo) GetMaturityDate() float32`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *SecdefInfo) GetMaturityDateOk() (*float32, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *SecdefInfo) SetMaturityDate(v float32)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *SecdefInfo) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetMultiplier

`func (o *SecdefInfo) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *SecdefInfo) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *SecdefInfo) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *SecdefInfo) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetTradingClass

`func (o *SecdefInfo) GetTradingClass() string`

GetTradingClass returns the TradingClass field if non-nil, zero value otherwise.

### GetTradingClassOk

`func (o *SecdefInfo) GetTradingClassOk() (*string, bool)`

GetTradingClassOk returns a tuple with the TradingClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingClass

`func (o *SecdefInfo) SetTradingClass(v string)`

SetTradingClass sets TradingClass field to given value.

### HasTradingClass

`func (o *SecdefInfo) HasTradingClass() bool`

HasTradingClass returns a boolean if a field has been set.

### GetValidExchanges

`func (o *SecdefInfo) GetValidExchanges() string`

GetValidExchanges returns the ValidExchanges field if non-nil, zero value otherwise.

### GetValidExchangesOk

`func (o *SecdefInfo) GetValidExchangesOk() (*string, bool)`

GetValidExchangesOk returns a tuple with the ValidExchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidExchanges

`func (o *SecdefInfo) SetValidExchanges(v string)`

SetValidExchanges sets ValidExchanges field to given value.

### HasValidExchanges

`func (o *SecdefInfo) HasValidExchanges() bool`

HasValidExchanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

