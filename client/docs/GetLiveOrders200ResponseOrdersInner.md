# GetLiveOrders200ResponseOrdersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acct** | Pointer to **string** | Account number | [optional] 
**Conidex** | Pointer to **string** | conid and exchange. Format supports conid or conid@exchange | [optional] 
**Conid** | Pointer to **float32** | Contract identifier | [optional] 
**OrderId** | Pointer to **string** | Order identifier | [optional] 
**CashCcy** | Pointer to **string** | Cash currency | [optional] 
**SizeAndFills** | Pointer to **string** | Quantity outstanding and total quantity concatenated with forward slash separator | [optional] 
**OrderDesc** | Pointer to **string** | Order description | [optional] 
**Description1** | Pointer to **string** | Formatted ticker description | [optional] 
**Ticker** | Pointer to **string** | Underlying symbol | [optional] 
**SecType** | Pointer to **string** | Asset class | [optional] 
**ListingExchange** | Pointer to **string** | Listing Exchange | [optional] 
**RemainingQuantity** | Pointer to **float32** | Quantity remaining | [optional] 
**FilledQuantity** | Pointer to **float32** | Quantity filled | [optional] 
**CompanyName** | Pointer to **string** | Company Name | [optional] 
**Status** | Pointer to **string** | Status of the order | [optional] 
**OrigOrderType** | Pointer to **string** | Original order type | [optional] 
**SupportsTaxOpt** | Pointer to **float32** | Supports Tax Optimization with 0 for no and 1 for yes | [optional] 
**LastExecutionTime** | Pointer to **float32** | Last status update in format YYMMDDhhmms based in GMT | [optional] 
**LastExecutionTimeR** | Pointer to **float32** | Last status update unix time in ms | [optional] 
**OrderType** | Pointer to **string** | Order type | [optional] 
**OrderRef** | Pointer to **string** | Order reference | [optional] 
**Side** | Pointer to **string** | The side of the market of the order.  * BUY: Buy contract near posted ask price  * SELL: Sell contract near posted bid price  * ASSN: Option Assignment, if BUYSELL&#x3D;BUY and OptionType&#x3D;PUT or BUYSELL&#x3D;SELL and OptionType&#x3D;CALL  * EXER: Option Exercise, if BUYSELL&#x3D;SELL and OptionType&#x3D;PUT or BUYSELL&#x3D;BUY and OptionType&#x3D;CALL  | [optional] 
**TimeInForce** | Pointer to **string** | Time in force | [optional] 
**Price** | Pointer to **float32** | Price of order | [optional] 
**BgColor** | Pointer to **string** | Background color in hex format | [optional] 
**FgColor** | Pointer to **string** | Foreground color in hex format | [optional] 

## Methods

### NewGetLiveOrders200ResponseOrdersInner

`func NewGetLiveOrders200ResponseOrdersInner() *GetLiveOrders200ResponseOrdersInner`

NewGetLiveOrders200ResponseOrdersInner instantiates a new GetLiveOrders200ResponseOrdersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLiveOrders200ResponseOrdersInnerWithDefaults

`func NewGetLiveOrders200ResponseOrdersInnerWithDefaults() *GetLiveOrders200ResponseOrdersInner`

NewGetLiveOrders200ResponseOrdersInnerWithDefaults instantiates a new GetLiveOrders200ResponseOrdersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcct

`func (o *GetLiveOrders200ResponseOrdersInner) GetAcct() string`

GetAcct returns the Acct field if non-nil, zero value otherwise.

### GetAcctOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetAcctOk() (*string, bool)`

GetAcctOk returns a tuple with the Acct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcct

`func (o *GetLiveOrders200ResponseOrdersInner) SetAcct(v string)`

SetAcct sets Acct field to given value.

### HasAcct

`func (o *GetLiveOrders200ResponseOrdersInner) HasAcct() bool`

HasAcct returns a boolean if a field has been set.

### GetConidex

`func (o *GetLiveOrders200ResponseOrdersInner) GetConidex() string`

GetConidex returns the Conidex field if non-nil, zero value otherwise.

### GetConidexOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetConidexOk() (*string, bool)`

GetConidexOk returns a tuple with the Conidex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConidex

`func (o *GetLiveOrders200ResponseOrdersInner) SetConidex(v string)`

SetConidex sets Conidex field to given value.

### HasConidex

`func (o *GetLiveOrders200ResponseOrdersInner) HasConidex() bool`

HasConidex returns a boolean if a field has been set.

### GetConid

`func (o *GetLiveOrders200ResponseOrdersInner) GetConid() float32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetConidOk() (*float32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *GetLiveOrders200ResponseOrdersInner) SetConid(v float32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *GetLiveOrders200ResponseOrdersInner) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetOrderId

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetLiveOrders200ResponseOrdersInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *GetLiveOrders200ResponseOrdersInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCashCcy

`func (o *GetLiveOrders200ResponseOrdersInner) GetCashCcy() string`

GetCashCcy returns the CashCcy field if non-nil, zero value otherwise.

### GetCashCcyOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetCashCcyOk() (*string, bool)`

GetCashCcyOk returns a tuple with the CashCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashCcy

`func (o *GetLiveOrders200ResponseOrdersInner) SetCashCcy(v string)`

SetCashCcy sets CashCcy field to given value.

### HasCashCcy

`func (o *GetLiveOrders200ResponseOrdersInner) HasCashCcy() bool`

HasCashCcy returns a boolean if a field has been set.

### GetSizeAndFills

`func (o *GetLiveOrders200ResponseOrdersInner) GetSizeAndFills() string`

GetSizeAndFills returns the SizeAndFills field if non-nil, zero value otherwise.

### GetSizeAndFillsOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetSizeAndFillsOk() (*string, bool)`

GetSizeAndFillsOk returns a tuple with the SizeAndFills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeAndFills

`func (o *GetLiveOrders200ResponseOrdersInner) SetSizeAndFills(v string)`

SetSizeAndFills sets SizeAndFills field to given value.

### HasSizeAndFills

`func (o *GetLiveOrders200ResponseOrdersInner) HasSizeAndFills() bool`

HasSizeAndFills returns a boolean if a field has been set.

### GetOrderDesc

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrderDesc() string`

GetOrderDesc returns the OrderDesc field if non-nil, zero value otherwise.

### GetOrderDescOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrderDescOk() (*string, bool)`

GetOrderDescOk returns a tuple with the OrderDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDesc

`func (o *GetLiveOrders200ResponseOrdersInner) SetOrderDesc(v string)`

SetOrderDesc sets OrderDesc field to given value.

### HasOrderDesc

`func (o *GetLiveOrders200ResponseOrdersInner) HasOrderDesc() bool`

HasOrderDesc returns a boolean if a field has been set.

### GetDescription1

`func (o *GetLiveOrders200ResponseOrdersInner) GetDescription1() string`

GetDescription1 returns the Description1 field if non-nil, zero value otherwise.

### GetDescription1Ok

`func (o *GetLiveOrders200ResponseOrdersInner) GetDescription1Ok() (*string, bool)`

GetDescription1Ok returns a tuple with the Description1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription1

`func (o *GetLiveOrders200ResponseOrdersInner) SetDescription1(v string)`

SetDescription1 sets Description1 field to given value.

### HasDescription1

`func (o *GetLiveOrders200ResponseOrdersInner) HasDescription1() bool`

HasDescription1 returns a boolean if a field has been set.

### GetTicker

`func (o *GetLiveOrders200ResponseOrdersInner) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *GetLiveOrders200ResponseOrdersInner) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *GetLiveOrders200ResponseOrdersInner) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetSecType

`func (o *GetLiveOrders200ResponseOrdersInner) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *GetLiveOrders200ResponseOrdersInner) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *GetLiveOrders200ResponseOrdersInner) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetListingExchange

`func (o *GetLiveOrders200ResponseOrdersInner) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *GetLiveOrders200ResponseOrdersInner) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *GetLiveOrders200ResponseOrdersInner) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *GetLiveOrders200ResponseOrdersInner) GetRemainingQuantity() float32`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetRemainingQuantityOk() (*float32, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *GetLiveOrders200ResponseOrdersInner) SetRemainingQuantity(v float32)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *GetLiveOrders200ResponseOrdersInner) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *GetLiveOrders200ResponseOrdersInner) GetFilledQuantity() float32`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetFilledQuantityOk() (*float32, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *GetLiveOrders200ResponseOrdersInner) SetFilledQuantity(v float32)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *GetLiveOrders200ResponseOrdersInner) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetCompanyName

`func (o *GetLiveOrders200ResponseOrdersInner) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *GetLiveOrders200ResponseOrdersInner) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *GetLiveOrders200ResponseOrdersInner) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetStatus

`func (o *GetLiveOrders200ResponseOrdersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLiveOrders200ResponseOrdersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLiveOrders200ResponseOrdersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrigOrderType

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrigOrderType() string`

GetOrigOrderType returns the OrigOrderType field if non-nil, zero value otherwise.

### GetOrigOrderTypeOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrigOrderTypeOk() (*string, bool)`

GetOrigOrderTypeOk returns a tuple with the OrigOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigOrderType

`func (o *GetLiveOrders200ResponseOrdersInner) SetOrigOrderType(v string)`

SetOrigOrderType sets OrigOrderType field to given value.

### HasOrigOrderType

`func (o *GetLiveOrders200ResponseOrdersInner) HasOrigOrderType() bool`

HasOrigOrderType returns a boolean if a field has been set.

### GetSupportsTaxOpt

`func (o *GetLiveOrders200ResponseOrdersInner) GetSupportsTaxOpt() float32`

GetSupportsTaxOpt returns the SupportsTaxOpt field if non-nil, zero value otherwise.

### GetSupportsTaxOptOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetSupportsTaxOptOk() (*float32, bool)`

GetSupportsTaxOptOk returns a tuple with the SupportsTaxOpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTaxOpt

`func (o *GetLiveOrders200ResponseOrdersInner) SetSupportsTaxOpt(v float32)`

SetSupportsTaxOpt sets SupportsTaxOpt field to given value.

### HasSupportsTaxOpt

`func (o *GetLiveOrders200ResponseOrdersInner) HasSupportsTaxOpt() bool`

HasSupportsTaxOpt returns a boolean if a field has been set.

### GetLastExecutionTime

`func (o *GetLiveOrders200ResponseOrdersInner) GetLastExecutionTime() float32`

GetLastExecutionTime returns the LastExecutionTime field if non-nil, zero value otherwise.

### GetLastExecutionTimeOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetLastExecutionTimeOk() (*float32, bool)`

GetLastExecutionTimeOk returns a tuple with the LastExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionTime

`func (o *GetLiveOrders200ResponseOrdersInner) SetLastExecutionTime(v float32)`

SetLastExecutionTime sets LastExecutionTime field to given value.

### HasLastExecutionTime

`func (o *GetLiveOrders200ResponseOrdersInner) HasLastExecutionTime() bool`

HasLastExecutionTime returns a boolean if a field has been set.

### GetLastExecutionTimeR

`func (o *GetLiveOrders200ResponseOrdersInner) GetLastExecutionTimeR() float32`

GetLastExecutionTimeR returns the LastExecutionTimeR field if non-nil, zero value otherwise.

### GetLastExecutionTimeROk

`func (o *GetLiveOrders200ResponseOrdersInner) GetLastExecutionTimeROk() (*float32, bool)`

GetLastExecutionTimeROk returns a tuple with the LastExecutionTimeR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExecutionTimeR

`func (o *GetLiveOrders200ResponseOrdersInner) SetLastExecutionTimeR(v float32)`

SetLastExecutionTimeR sets LastExecutionTimeR field to given value.

### HasLastExecutionTimeR

`func (o *GetLiveOrders200ResponseOrdersInner) HasLastExecutionTimeR() bool`

HasLastExecutionTimeR returns a boolean if a field has been set.

### GetOrderType

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *GetLiveOrders200ResponseOrdersInner) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *GetLiveOrders200ResponseOrdersInner) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrderRef

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrderRef() string`

GetOrderRef returns the OrderRef field if non-nil, zero value otherwise.

### GetOrderRefOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetOrderRefOk() (*string, bool)`

GetOrderRefOk returns a tuple with the OrderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRef

`func (o *GetLiveOrders200ResponseOrdersInner) SetOrderRef(v string)`

SetOrderRef sets OrderRef field to given value.

### HasOrderRef

`func (o *GetLiveOrders200ResponseOrdersInner) HasOrderRef() bool`

HasOrderRef returns a boolean if a field has been set.

### GetSide

`func (o *GetLiveOrders200ResponseOrdersInner) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *GetLiveOrders200ResponseOrdersInner) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *GetLiveOrders200ResponseOrdersInner) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetTimeInForce

`func (o *GetLiveOrders200ResponseOrdersInner) GetTimeInForce() string`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetTimeInForceOk() (*string, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *GetLiveOrders200ResponseOrdersInner) SetTimeInForce(v string)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *GetLiveOrders200ResponseOrdersInner) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetPrice

`func (o *GetLiveOrders200ResponseOrdersInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetLiveOrders200ResponseOrdersInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetLiveOrders200ResponseOrdersInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetBgColor

`func (o *GetLiveOrders200ResponseOrdersInner) GetBgColor() string`

GetBgColor returns the BgColor field if non-nil, zero value otherwise.

### GetBgColorOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetBgColorOk() (*string, bool)`

GetBgColorOk returns a tuple with the BgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgColor

`func (o *GetLiveOrders200ResponseOrdersInner) SetBgColor(v string)`

SetBgColor sets BgColor field to given value.

### HasBgColor

`func (o *GetLiveOrders200ResponseOrdersInner) HasBgColor() bool`

HasBgColor returns a boolean if a field has been set.

### GetFgColor

`func (o *GetLiveOrders200ResponseOrdersInner) GetFgColor() string`

GetFgColor returns the FgColor field if non-nil, zero value otherwise.

### GetFgColorOk

`func (o *GetLiveOrders200ResponseOrdersInner) GetFgColorOk() (*string, bool)`

GetFgColorOk returns a tuple with the FgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFgColor

`func (o *GetLiveOrders200ResponseOrdersInner) SetFgColor(v string)`

SetFgColor sets FgColor field to given value.

### HasFgColor

`func (o *GetLiveOrders200ResponseOrdersInner) HasFgColor() bool`

HasFgColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


