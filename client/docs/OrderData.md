# OrderData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientOrderId** | Pointer to **string** |  | [optional] 
**ExecId** | Pointer to **string** |  | [optional] 
**ExecType** | Pointer to **string** |  | [optional] 
**OrderType** | Pointer to **string** |  | [optional] 
**OrderStatus** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** | Underlying symbol for contract | [optional] 
**OrderQty** | Pointer to **string** | Quantity of active order | [optional] 
**Price** | Pointer to **string** | Price of active order | [optional] 
**LastShares** | Pointer to **string** | Quantity of the last partial fill | [optional] 
**LastPrice** | Pointer to **string** | Price of the last partial fill | [optional] 
**CumQty** | Pointer to **string** | Cumulative fill quantity | [optional] 
**LeavesQty** | Pointer to **string** | Remaining quantity to be filled | [optional] 
**AvgPrice** | Pointer to **string** | Average fill price | [optional] 
**Side** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** | Order identifier | [optional] 
**Account** | Pointer to **string** | Account number | [optional] 
**SecType** | Pointer to **string** | Contracts asset class | [optional] 
**TxTime** | Pointer to **string** | Time of transaction in GMT, format YYYYMMDD-hh:m:ss | [optional] 
**RcptTime** | Pointer to **string** | Time of receipt in GMT, format YYYYMMDD-hh:mm:ss | [optional] 
**Tif** | Pointer to **string** | Time in Force | [optional] 
**Conid** | Pointer to **string** | Contract identifier from IBKR&#39;s database. | [optional] 
**Currency** | Pointer to **string** | Trading currency | [optional] 
**Exchange** | Pointer to **string** | Exchange or venue | [optional] 
**ListingExchange** | Pointer to **string** | Listing Exchange | [optional] 
**Text** | Pointer to **float32** | error message | [optional] 
**Warnings** | Pointer to [**OrderDataWarnings**](OrderDataWarnings.md) |  | [optional] 
**CommCurr** | Pointer to **string** | Commission currency | [optional] 
**Comms** | Pointer to **string** | Commissions | [optional] 
**RealizedPnl** | Pointer to **string** | Realized PnL | [optional] 

## Methods

### NewOrderData

`func NewOrderData() *OrderData`

NewOrderData instantiates a new OrderData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDataWithDefaults

`func NewOrderDataWithDefaults() *OrderData`

NewOrderDataWithDefaults instantiates a new OrderData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientOrderId

`func (o *OrderData) GetClientOrderId() string`

GetClientOrderId returns the ClientOrderId field if non-nil, zero value otherwise.

### GetClientOrderIdOk

`func (o *OrderData) GetClientOrderIdOk() (*string, bool)`

GetClientOrderIdOk returns a tuple with the ClientOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientOrderId

`func (o *OrderData) SetClientOrderId(v string)`

SetClientOrderId sets ClientOrderId field to given value.

### HasClientOrderId

`func (o *OrderData) HasClientOrderId() bool`

HasClientOrderId returns a boolean if a field has been set.

### GetExecId

`func (o *OrderData) GetExecId() string`

GetExecId returns the ExecId field if non-nil, zero value otherwise.

### GetExecIdOk

`func (o *OrderData) GetExecIdOk() (*string, bool)`

GetExecIdOk returns a tuple with the ExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecId

`func (o *OrderData) SetExecId(v string)`

SetExecId sets ExecId field to given value.

### HasExecId

`func (o *OrderData) HasExecId() bool`

HasExecId returns a boolean if a field has been set.

### GetExecType

`func (o *OrderData) GetExecType() string`

GetExecType returns the ExecType field if non-nil, zero value otherwise.

### GetExecTypeOk

`func (o *OrderData) GetExecTypeOk() (*string, bool)`

GetExecTypeOk returns a tuple with the ExecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecType

`func (o *OrderData) SetExecType(v string)`

SetExecType sets ExecType field to given value.

### HasExecType

`func (o *OrderData) HasExecType() bool`

HasExecType returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderData) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderData) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderData) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderData) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderData) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderData) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderData) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderData) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetSymbol

`func (o *OrderData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetOrderQty

`func (o *OrderData) GetOrderQty() string`

GetOrderQty returns the OrderQty field if non-nil, zero value otherwise.

### GetOrderQtyOk

`func (o *OrderData) GetOrderQtyOk() (*string, bool)`

GetOrderQtyOk returns a tuple with the OrderQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderQty

`func (o *OrderData) SetOrderQty(v string)`

SetOrderQty sets OrderQty field to given value.

### HasOrderQty

`func (o *OrderData) HasOrderQty() bool`

HasOrderQty returns a boolean if a field has been set.

### GetPrice

`func (o *OrderData) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderData) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderData) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderData) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetLastShares

`func (o *OrderData) GetLastShares() string`

GetLastShares returns the LastShares field if non-nil, zero value otherwise.

### GetLastSharesOk

`func (o *OrderData) GetLastSharesOk() (*string, bool)`

GetLastSharesOk returns a tuple with the LastShares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastShares

`func (o *OrderData) SetLastShares(v string)`

SetLastShares sets LastShares field to given value.

### HasLastShares

`func (o *OrderData) HasLastShares() bool`

HasLastShares returns a boolean if a field has been set.

### GetLastPrice

`func (o *OrderData) GetLastPrice() string`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *OrderData) GetLastPriceOk() (*string, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *OrderData) SetLastPrice(v string)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *OrderData) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetCumQty

`func (o *OrderData) GetCumQty() string`

GetCumQty returns the CumQty field if non-nil, zero value otherwise.

### GetCumQtyOk

`func (o *OrderData) GetCumQtyOk() (*string, bool)`

GetCumQtyOk returns a tuple with the CumQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumQty

`func (o *OrderData) SetCumQty(v string)`

SetCumQty sets CumQty field to given value.

### HasCumQty

`func (o *OrderData) HasCumQty() bool`

HasCumQty returns a boolean if a field has been set.

### GetLeavesQty

`func (o *OrderData) GetLeavesQty() string`

GetLeavesQty returns the LeavesQty field if non-nil, zero value otherwise.

### GetLeavesQtyOk

`func (o *OrderData) GetLeavesQtyOk() (*string, bool)`

GetLeavesQtyOk returns a tuple with the LeavesQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeavesQty

`func (o *OrderData) SetLeavesQty(v string)`

SetLeavesQty sets LeavesQty field to given value.

### HasLeavesQty

`func (o *OrderData) HasLeavesQty() bool`

HasLeavesQty returns a boolean if a field has been set.

### GetAvgPrice

`func (o *OrderData) GetAvgPrice() string`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *OrderData) GetAvgPriceOk() (*string, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *OrderData) SetAvgPrice(v string)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *OrderData) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetSide

`func (o *OrderData) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderData) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderData) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderData) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderData) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAccount

`func (o *OrderData) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OrderData) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OrderData) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OrderData) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSecType

`func (o *OrderData) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *OrderData) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *OrderData) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *OrderData) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetTxTime

`func (o *OrderData) GetTxTime() string`

GetTxTime returns the TxTime field if non-nil, zero value otherwise.

### GetTxTimeOk

`func (o *OrderData) GetTxTimeOk() (*string, bool)`

GetTxTimeOk returns a tuple with the TxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxTime

`func (o *OrderData) SetTxTime(v string)`

SetTxTime sets TxTime field to given value.

### HasTxTime

`func (o *OrderData) HasTxTime() bool`

HasTxTime returns a boolean if a field has been set.

### GetRcptTime

`func (o *OrderData) GetRcptTime() string`

GetRcptTime returns the RcptTime field if non-nil, zero value otherwise.

### GetRcptTimeOk

`func (o *OrderData) GetRcptTimeOk() (*string, bool)`

GetRcptTimeOk returns a tuple with the RcptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcptTime

`func (o *OrderData) SetRcptTime(v string)`

SetRcptTime sets RcptTime field to given value.

### HasRcptTime

`func (o *OrderData) HasRcptTime() bool`

HasRcptTime returns a boolean if a field has been set.

### GetTif

`func (o *OrderData) GetTif() string`

GetTif returns the Tif field if non-nil, zero value otherwise.

### GetTifOk

`func (o *OrderData) GetTifOk() (*string, bool)`

GetTifOk returns a tuple with the Tif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTif

`func (o *OrderData) SetTif(v string)`

SetTif sets Tif field to given value.

### HasTif

`func (o *OrderData) HasTif() bool`

HasTif returns a boolean if a field has been set.

### GetConid

`func (o *OrderData) GetConid() string`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *OrderData) GetConidOk() (*string, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *OrderData) SetConid(v string)`

SetConid sets Conid field to given value.

### HasConid

`func (o *OrderData) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderData) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderData) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderData) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderData) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetExchange

`func (o *OrderData) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *OrderData) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *OrderData) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *OrderData) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetListingExchange

`func (o *OrderData) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *OrderData) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *OrderData) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *OrderData) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetText

`func (o *OrderData) GetText() float32`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OrderData) GetTextOk() (*float32, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OrderData) SetText(v float32)`

SetText sets Text field to given value.

### HasText

`func (o *OrderData) HasText() bool`

HasText returns a boolean if a field has been set.

### GetWarnings

`func (o *OrderData) GetWarnings() OrderDataWarnings`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *OrderData) GetWarningsOk() (*OrderDataWarnings, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *OrderData) SetWarnings(v OrderDataWarnings)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *OrderData) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetCommCurr

`func (o *OrderData) GetCommCurr() string`

GetCommCurr returns the CommCurr field if non-nil, zero value otherwise.

### GetCommCurrOk

`func (o *OrderData) GetCommCurrOk() (*string, bool)`

GetCommCurrOk returns a tuple with the CommCurr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommCurr

`func (o *OrderData) SetCommCurr(v string)`

SetCommCurr sets CommCurr field to given value.

### HasCommCurr

`func (o *OrderData) HasCommCurr() bool`

HasCommCurr returns a boolean if a field has been set.

### GetComms

`func (o *OrderData) GetComms() string`

GetComms returns the Comms field if non-nil, zero value otherwise.

### GetCommsOk

`func (o *OrderData) GetCommsOk() (*string, bool)`

GetCommsOk returns a tuple with the Comms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComms

`func (o *OrderData) SetComms(v string)`

SetComms sets Comms field to given value.

### HasComms

`func (o *OrderData) HasComms() bool`

HasComms returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *OrderData) GetRealizedPnl() string`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *OrderData) GetRealizedPnlOk() (*string, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *OrderData) SetRealizedPnl(v string)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *OrderData) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


