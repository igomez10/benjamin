# OrderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubType** | Pointer to **string** | order sub-type | [optional] 
**RequestId** | Pointer to **string** | order request id | [optional] 
**OrderId** | Pointer to **int32** | system generated order id, unique per account | [optional] 
**Conidex** | Pointer to **string** | conid and exchange. Format supports conid or conid@exchange | [optional] 
**Symbol** | Pointer to **string** | Underlying symbol | [optional] 
**Side** | Pointer to **string** | The side of the market of the order.   * B - Buy contract near posted ask price   * S - Sell contract near posted bid price   * X - Option expired  | [optional] 
**ContractDescription1** | Pointer to **string** | Format contract name | [optional] 
**ListingExchange** | Pointer to **string** | Trading Exchange or Venue | [optional] 
**OptionAcct** | Pointer to **string** |  | [optional] 
**CompanyName** | Pointer to **string** | Contracts company name | [optional] 
**Size** | Pointer to **string** | Quantity updated | [optional] 
**TotalSize** | Pointer to **string** | Total quantity | [optional] 
**Currency** | Pointer to **string** | Contract traded currency | [optional] 
**Account** | Pointer to **string** | account id | [optional] 
**OrderType** | Pointer to **string** | Types of orders | [optional] 
**LimitPrice** | Pointer to **string** | Limit price | [optional] 
**StopPrice** | Pointer to **string** | Stop price | [optional] 
**CumFill** | Pointer to **string** | Cumulative fill | [optional] 
**OrderStatus** | Pointer to **string** | *  PendingSubmit - Indicates the order was sent, but confirmation has not been received that it has been received by the destination.                    Occurs most commonly if an exchange is closed. *  PendingCancel - Indicates that a request has been sent to cancel an order but confirmation has not been received of its cancellation. *  PreSubmitted - Indicates that a simulated order type has been accepted by the IBKR system and that this order has yet to be elected.                   The order is held in the IBKR system until the election criteria are met. At that time the order is transmitted to the order destination as specified. *  Submitted - Indicates that the order has been accepted at the order destination and is working. *  Cancelled - Indicates that the balance of the order has been confirmed cancelled by the IB system.                This could occur unexpectedly when IB or the destination has rejected the order. *  Filled - Indicates that the order has been completely filled. *  Inactive - Indicates the order is not working, for instance if the order was invalid and triggered an error message,               or if the order was to short a security and shares have not yet been located.  | [optional] 
**OrderStatusDescription** | Pointer to **string** | Description of the order status | [optional] 
**Tif** | Pointer to **string** | Time-in-Force - length of time order will continue working before it is canceled. | [optional] 
**FgColor** | Pointer to **string** | Foreground color in hex format | [optional] 
**BgColor** | Pointer to **string** | Background color in hex format | [optional] 
**OrderNotEditable** | Pointer to **bool** | If true not allowed to modify order | [optional] 
**EditableFields** | Pointer to **string** | Fields that can be edited in escaped unicode characters | [optional] 
**CannotCancelOrder** | Pointer to **bool** | If true not allowed to cancel order | [optional] 
**OutsideRth** | Pointer to **bool** | If true order trades outside regular trading hours | [optional] 
**DeactivateOrder** | Pointer to **bool** | If true order is de-activated | [optional] 
**UsePriceMgmtAlgo** | Pointer to **bool** | If true price management algo is enabled, refer to https://www.interactivebrokers.com/en/index.php?f&#x3D;43423 | [optional] 
**SecType** | Pointer to **string** | Asset class | [optional] 
**AvailableChartPeriods** | Pointer to **string** | List of available chart periods | [optional] 
**OrderDescription** | Pointer to **string** | Format description of order | [optional] 
**OrderDescriptionWithContract** | Pointer to **string** | order_description with the symbol | [optional] 
**AlertActive** | Pointer to **int32** |  | [optional] 
**ChildOrderType** | Pointer to **string** | type of the child order | [optional] 
**SizeAndFills** | Pointer to **string** | Format fillQuantity\\totalQuantity | [optional] 
**ExitStrategyDisplayPrice** | Pointer to **string** | Position display price | [optional] 
**ExitStrategyChartDescription** | Pointer to **string** | Position description to display on chart | [optional] 
**ExitStrategyToolAvailability** | Pointer to **string** | * 1: If your account has position or order for contract * 0: If your account has no position or order for contract  | [optional] 
**AllowedDuplicateOpposite** | Pointer to **bool** | Returns true if contract supports duplicate/opposite side order. | [optional] 
**OrderTime** | Pointer to **string** | Time of status update in unix time | [optional] 
**OcaGroupId** | Pointer to **string** | only exists for oca orders, oca orders in same group will have same id | [optional] 

## Methods

### NewOrderStatus

`func NewOrderStatus() *OrderStatus`

NewOrderStatus instantiates a new OrderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStatusWithDefaults

`func NewOrderStatusWithDefaults() *OrderStatus`

NewOrderStatusWithDefaults instantiates a new OrderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubType

`func (o *OrderStatus) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *OrderStatus) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *OrderStatus) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *OrderStatus) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetRequestId

`func (o *OrderStatus) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *OrderStatus) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *OrderStatus) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *OrderStatus) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderStatus) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderStatus) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderStatus) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderStatus) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetConidex

`func (o *OrderStatus) GetConidex() string`

GetConidex returns the Conidex field if non-nil, zero value otherwise.

### GetConidexOk

`func (o *OrderStatus) GetConidexOk() (*string, bool)`

GetConidexOk returns a tuple with the Conidex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConidex

`func (o *OrderStatus) SetConidex(v string)`

SetConidex sets Conidex field to given value.

### HasConidex

`func (o *OrderStatus) HasConidex() bool`

HasConidex returns a boolean if a field has been set.

### GetSymbol

`func (o *OrderStatus) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *OrderStatus) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *OrderStatus) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *OrderStatus) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *OrderStatus) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderStatus) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderStatus) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderStatus) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetContractDescription1

`func (o *OrderStatus) GetContractDescription1() string`

GetContractDescription1 returns the ContractDescription1 field if non-nil, zero value otherwise.

### GetContractDescription1Ok

`func (o *OrderStatus) GetContractDescription1Ok() (*string, bool)`

GetContractDescription1Ok returns a tuple with the ContractDescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDescription1

`func (o *OrderStatus) SetContractDescription1(v string)`

SetContractDescription1 sets ContractDescription1 field to given value.

### HasContractDescription1

`func (o *OrderStatus) HasContractDescription1() bool`

HasContractDescription1 returns a boolean if a field has been set.

### GetListingExchange

`func (o *OrderStatus) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *OrderStatus) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *OrderStatus) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *OrderStatus) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetOptionAcct

`func (o *OrderStatus) GetOptionAcct() string`

GetOptionAcct returns the OptionAcct field if non-nil, zero value otherwise.

### GetOptionAcctOk

`func (o *OrderStatus) GetOptionAcctOk() (*string, bool)`

GetOptionAcctOk returns a tuple with the OptionAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionAcct

`func (o *OrderStatus) SetOptionAcct(v string)`

SetOptionAcct sets OptionAcct field to given value.

### HasOptionAcct

`func (o *OrderStatus) HasOptionAcct() bool`

HasOptionAcct returns a boolean if a field has been set.

### GetCompanyName

`func (o *OrderStatus) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *OrderStatus) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *OrderStatus) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *OrderStatus) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetSize

`func (o *OrderStatus) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *OrderStatus) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *OrderStatus) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *OrderStatus) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotalSize

`func (o *OrderStatus) GetTotalSize() string`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *OrderStatus) GetTotalSizeOk() (*string, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *OrderStatus) SetTotalSize(v string)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *OrderStatus) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetCurrency

`func (o *OrderStatus) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *OrderStatus) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *OrderStatus) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *OrderStatus) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAccount

`func (o *OrderStatus) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OrderStatus) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OrderStatus) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OrderStatus) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderStatus) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderStatus) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderStatus) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderStatus) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetLimitPrice

`func (o *OrderStatus) GetLimitPrice() string`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *OrderStatus) GetLimitPriceOk() (*string, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *OrderStatus) SetLimitPrice(v string)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *OrderStatus) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderStatus) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderStatus) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderStatus) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderStatus) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetCumFill

`func (o *OrderStatus) GetCumFill() string`

GetCumFill returns the CumFill field if non-nil, zero value otherwise.

### GetCumFillOk

`func (o *OrderStatus) GetCumFillOk() (*string, bool)`

GetCumFillOk returns a tuple with the CumFill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumFill

`func (o *OrderStatus) SetCumFill(v string)`

SetCumFill sets CumFill field to given value.

### HasCumFill

`func (o *OrderStatus) HasCumFill() bool`

HasCumFill returns a boolean if a field has been set.

### GetOrderStatus

`func (o *OrderStatus) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *OrderStatus) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *OrderStatus) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *OrderStatus) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOrderStatusDescription

`func (o *OrderStatus) GetOrderStatusDescription() string`

GetOrderStatusDescription returns the OrderStatusDescription field if non-nil, zero value otherwise.

### GetOrderStatusDescriptionOk

`func (o *OrderStatus) GetOrderStatusDescriptionOk() (*string, bool)`

GetOrderStatusDescriptionOk returns a tuple with the OrderStatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatusDescription

`func (o *OrderStatus) SetOrderStatusDescription(v string)`

SetOrderStatusDescription sets OrderStatusDescription field to given value.

### HasOrderStatusDescription

`func (o *OrderStatus) HasOrderStatusDescription() bool`

HasOrderStatusDescription returns a boolean if a field has been set.

### GetTif

`func (o *OrderStatus) GetTif() string`

GetTif returns the Tif field if non-nil, zero value otherwise.

### GetTifOk

`func (o *OrderStatus) GetTifOk() (*string, bool)`

GetTifOk returns a tuple with the Tif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTif

`func (o *OrderStatus) SetTif(v string)`

SetTif sets Tif field to given value.

### HasTif

`func (o *OrderStatus) HasTif() bool`

HasTif returns a boolean if a field has been set.

### GetFgColor

`func (o *OrderStatus) GetFgColor() string`

GetFgColor returns the FgColor field if non-nil, zero value otherwise.

### GetFgColorOk

`func (o *OrderStatus) GetFgColorOk() (*string, bool)`

GetFgColorOk returns a tuple with the FgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFgColor

`func (o *OrderStatus) SetFgColor(v string)`

SetFgColor sets FgColor field to given value.

### HasFgColor

`func (o *OrderStatus) HasFgColor() bool`

HasFgColor returns a boolean if a field has been set.

### GetBgColor

`func (o *OrderStatus) GetBgColor() string`

GetBgColor returns the BgColor field if non-nil, zero value otherwise.

### GetBgColorOk

`func (o *OrderStatus) GetBgColorOk() (*string, bool)`

GetBgColorOk returns a tuple with the BgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgColor

`func (o *OrderStatus) SetBgColor(v string)`

SetBgColor sets BgColor field to given value.

### HasBgColor

`func (o *OrderStatus) HasBgColor() bool`

HasBgColor returns a boolean if a field has been set.

### GetOrderNotEditable

`func (o *OrderStatus) GetOrderNotEditable() bool`

GetOrderNotEditable returns the OrderNotEditable field if non-nil, zero value otherwise.

### GetOrderNotEditableOk

`func (o *OrderStatus) GetOrderNotEditableOk() (*bool, bool)`

GetOrderNotEditableOk returns a tuple with the OrderNotEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNotEditable

`func (o *OrderStatus) SetOrderNotEditable(v bool)`

SetOrderNotEditable sets OrderNotEditable field to given value.

### HasOrderNotEditable

`func (o *OrderStatus) HasOrderNotEditable() bool`

HasOrderNotEditable returns a boolean if a field has been set.

### GetEditableFields

`func (o *OrderStatus) GetEditableFields() string`

GetEditableFields returns the EditableFields field if non-nil, zero value otherwise.

### GetEditableFieldsOk

`func (o *OrderStatus) GetEditableFieldsOk() (*string, bool)`

GetEditableFieldsOk returns a tuple with the EditableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableFields

`func (o *OrderStatus) SetEditableFields(v string)`

SetEditableFields sets EditableFields field to given value.

### HasEditableFields

`func (o *OrderStatus) HasEditableFields() bool`

HasEditableFields returns a boolean if a field has been set.

### GetCannotCancelOrder

`func (o *OrderStatus) GetCannotCancelOrder() bool`

GetCannotCancelOrder returns the CannotCancelOrder field if non-nil, zero value otherwise.

### GetCannotCancelOrderOk

`func (o *OrderStatus) GetCannotCancelOrderOk() (*bool, bool)`

GetCannotCancelOrderOk returns a tuple with the CannotCancelOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCannotCancelOrder

`func (o *OrderStatus) SetCannotCancelOrder(v bool)`

SetCannotCancelOrder sets CannotCancelOrder field to given value.

### HasCannotCancelOrder

`func (o *OrderStatus) HasCannotCancelOrder() bool`

HasCannotCancelOrder returns a boolean if a field has been set.

### GetOutsideRth

`func (o *OrderStatus) GetOutsideRth() bool`

GetOutsideRth returns the OutsideRth field if non-nil, zero value otherwise.

### GetOutsideRthOk

`func (o *OrderStatus) GetOutsideRthOk() (*bool, bool)`

GetOutsideRthOk returns a tuple with the OutsideRth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRth

`func (o *OrderStatus) SetOutsideRth(v bool)`

SetOutsideRth sets OutsideRth field to given value.

### HasOutsideRth

`func (o *OrderStatus) HasOutsideRth() bool`

HasOutsideRth returns a boolean if a field has been set.

### GetDeactivateOrder

`func (o *OrderStatus) GetDeactivateOrder() bool`

GetDeactivateOrder returns the DeactivateOrder field if non-nil, zero value otherwise.

### GetDeactivateOrderOk

`func (o *OrderStatus) GetDeactivateOrderOk() (*bool, bool)`

GetDeactivateOrderOk returns a tuple with the DeactivateOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivateOrder

`func (o *OrderStatus) SetDeactivateOrder(v bool)`

SetDeactivateOrder sets DeactivateOrder field to given value.

### HasDeactivateOrder

`func (o *OrderStatus) HasDeactivateOrder() bool`

HasDeactivateOrder returns a boolean if a field has been set.

### GetUsePriceMgmtAlgo

`func (o *OrderStatus) GetUsePriceMgmtAlgo() bool`

GetUsePriceMgmtAlgo returns the UsePriceMgmtAlgo field if non-nil, zero value otherwise.

### GetUsePriceMgmtAlgoOk

`func (o *OrderStatus) GetUsePriceMgmtAlgoOk() (*bool, bool)`

GetUsePriceMgmtAlgoOk returns a tuple with the UsePriceMgmtAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePriceMgmtAlgo

`func (o *OrderStatus) SetUsePriceMgmtAlgo(v bool)`

SetUsePriceMgmtAlgo sets UsePriceMgmtAlgo field to given value.

### HasUsePriceMgmtAlgo

`func (o *OrderStatus) HasUsePriceMgmtAlgo() bool`

HasUsePriceMgmtAlgo returns a boolean if a field has been set.

### GetSecType

`func (o *OrderStatus) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *OrderStatus) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *OrderStatus) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *OrderStatus) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetAvailableChartPeriods

`func (o *OrderStatus) GetAvailableChartPeriods() string`

GetAvailableChartPeriods returns the AvailableChartPeriods field if non-nil, zero value otherwise.

### GetAvailableChartPeriodsOk

`func (o *OrderStatus) GetAvailableChartPeriodsOk() (*string, bool)`

GetAvailableChartPeriodsOk returns a tuple with the AvailableChartPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableChartPeriods

`func (o *OrderStatus) SetAvailableChartPeriods(v string)`

SetAvailableChartPeriods sets AvailableChartPeriods field to given value.

### HasAvailableChartPeriods

`func (o *OrderStatus) HasAvailableChartPeriods() bool`

HasAvailableChartPeriods returns a boolean if a field has been set.

### GetOrderDescription

`func (o *OrderStatus) GetOrderDescription() string`

GetOrderDescription returns the OrderDescription field if non-nil, zero value otherwise.

### GetOrderDescriptionOk

`func (o *OrderStatus) GetOrderDescriptionOk() (*string, bool)`

GetOrderDescriptionOk returns a tuple with the OrderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDescription

`func (o *OrderStatus) SetOrderDescription(v string)`

SetOrderDescription sets OrderDescription field to given value.

### HasOrderDescription

`func (o *OrderStatus) HasOrderDescription() bool`

HasOrderDescription returns a boolean if a field has been set.

### GetOrderDescriptionWithContract

`func (o *OrderStatus) GetOrderDescriptionWithContract() string`

GetOrderDescriptionWithContract returns the OrderDescriptionWithContract field if non-nil, zero value otherwise.

### GetOrderDescriptionWithContractOk

`func (o *OrderStatus) GetOrderDescriptionWithContractOk() (*string, bool)`

GetOrderDescriptionWithContractOk returns a tuple with the OrderDescriptionWithContract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDescriptionWithContract

`func (o *OrderStatus) SetOrderDescriptionWithContract(v string)`

SetOrderDescriptionWithContract sets OrderDescriptionWithContract field to given value.

### HasOrderDescriptionWithContract

`func (o *OrderStatus) HasOrderDescriptionWithContract() bool`

HasOrderDescriptionWithContract returns a boolean if a field has been set.

### GetAlertActive

`func (o *OrderStatus) GetAlertActive() int32`

GetAlertActive returns the AlertActive field if non-nil, zero value otherwise.

### GetAlertActiveOk

`func (o *OrderStatus) GetAlertActiveOk() (*int32, bool)`

GetAlertActiveOk returns a tuple with the AlertActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertActive

`func (o *OrderStatus) SetAlertActive(v int32)`

SetAlertActive sets AlertActive field to given value.

### HasAlertActive

`func (o *OrderStatus) HasAlertActive() bool`

HasAlertActive returns a boolean if a field has been set.

### GetChildOrderType

`func (o *OrderStatus) GetChildOrderType() string`

GetChildOrderType returns the ChildOrderType field if non-nil, zero value otherwise.

### GetChildOrderTypeOk

`func (o *OrderStatus) GetChildOrderTypeOk() (*string, bool)`

GetChildOrderTypeOk returns a tuple with the ChildOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOrderType

`func (o *OrderStatus) SetChildOrderType(v string)`

SetChildOrderType sets ChildOrderType field to given value.

### HasChildOrderType

`func (o *OrderStatus) HasChildOrderType() bool`

HasChildOrderType returns a boolean if a field has been set.

### GetSizeAndFills

`func (o *OrderStatus) GetSizeAndFills() string`

GetSizeAndFills returns the SizeAndFills field if non-nil, zero value otherwise.

### GetSizeAndFillsOk

`func (o *OrderStatus) GetSizeAndFillsOk() (*string, bool)`

GetSizeAndFillsOk returns a tuple with the SizeAndFills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeAndFills

`func (o *OrderStatus) SetSizeAndFills(v string)`

SetSizeAndFills sets SizeAndFills field to given value.

### HasSizeAndFills

`func (o *OrderStatus) HasSizeAndFills() bool`

HasSizeAndFills returns a boolean if a field has been set.

### GetExitStrategyDisplayPrice

`func (o *OrderStatus) GetExitStrategyDisplayPrice() string`

GetExitStrategyDisplayPrice returns the ExitStrategyDisplayPrice field if non-nil, zero value otherwise.

### GetExitStrategyDisplayPriceOk

`func (o *OrderStatus) GetExitStrategyDisplayPriceOk() (*string, bool)`

GetExitStrategyDisplayPriceOk returns a tuple with the ExitStrategyDisplayPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitStrategyDisplayPrice

`func (o *OrderStatus) SetExitStrategyDisplayPrice(v string)`

SetExitStrategyDisplayPrice sets ExitStrategyDisplayPrice field to given value.

### HasExitStrategyDisplayPrice

`func (o *OrderStatus) HasExitStrategyDisplayPrice() bool`

HasExitStrategyDisplayPrice returns a boolean if a field has been set.

### GetExitStrategyChartDescription

`func (o *OrderStatus) GetExitStrategyChartDescription() string`

GetExitStrategyChartDescription returns the ExitStrategyChartDescription field if non-nil, zero value otherwise.

### GetExitStrategyChartDescriptionOk

`func (o *OrderStatus) GetExitStrategyChartDescriptionOk() (*string, bool)`

GetExitStrategyChartDescriptionOk returns a tuple with the ExitStrategyChartDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitStrategyChartDescription

`func (o *OrderStatus) SetExitStrategyChartDescription(v string)`

SetExitStrategyChartDescription sets ExitStrategyChartDescription field to given value.

### HasExitStrategyChartDescription

`func (o *OrderStatus) HasExitStrategyChartDescription() bool`

HasExitStrategyChartDescription returns a boolean if a field has been set.

### GetExitStrategyToolAvailability

`func (o *OrderStatus) GetExitStrategyToolAvailability() string`

GetExitStrategyToolAvailability returns the ExitStrategyToolAvailability field if non-nil, zero value otherwise.

### GetExitStrategyToolAvailabilityOk

`func (o *OrderStatus) GetExitStrategyToolAvailabilityOk() (*string, bool)`

GetExitStrategyToolAvailabilityOk returns a tuple with the ExitStrategyToolAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitStrategyToolAvailability

`func (o *OrderStatus) SetExitStrategyToolAvailability(v string)`

SetExitStrategyToolAvailability sets ExitStrategyToolAvailability field to given value.

### HasExitStrategyToolAvailability

`func (o *OrderStatus) HasExitStrategyToolAvailability() bool`

HasExitStrategyToolAvailability returns a boolean if a field has been set.

### GetAllowedDuplicateOpposite

`func (o *OrderStatus) GetAllowedDuplicateOpposite() bool`

GetAllowedDuplicateOpposite returns the AllowedDuplicateOpposite field if non-nil, zero value otherwise.

### GetAllowedDuplicateOppositeOk

`func (o *OrderStatus) GetAllowedDuplicateOppositeOk() (*bool, bool)`

GetAllowedDuplicateOppositeOk returns a tuple with the AllowedDuplicateOpposite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDuplicateOpposite

`func (o *OrderStatus) SetAllowedDuplicateOpposite(v bool)`

SetAllowedDuplicateOpposite sets AllowedDuplicateOpposite field to given value.

### HasAllowedDuplicateOpposite

`func (o *OrderStatus) HasAllowedDuplicateOpposite() bool`

HasAllowedDuplicateOpposite returns a boolean if a field has been set.

### GetOrderTime

`func (o *OrderStatus) GetOrderTime() string`

GetOrderTime returns the OrderTime field if non-nil, zero value otherwise.

### GetOrderTimeOk

`func (o *OrderStatus) GetOrderTimeOk() (*string, bool)`

GetOrderTimeOk returns a tuple with the OrderTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTime

`func (o *OrderStatus) SetOrderTime(v string)`

SetOrderTime sets OrderTime field to given value.

### HasOrderTime

`func (o *OrderStatus) HasOrderTime() bool`

HasOrderTime returns a boolean if a field has been set.

### GetOcaGroupId

`func (o *OrderStatus) GetOcaGroupId() string`

GetOcaGroupId returns the OcaGroupId field if non-nil, zero value otherwise.

### GetOcaGroupIdOk

`func (o *OrderStatus) GetOcaGroupIdOk() (*string, bool)`

GetOcaGroupIdOk returns a tuple with the OcaGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcaGroupId

`func (o *OrderStatus) SetOcaGroupId(v string)`

SetOcaGroupId sets OcaGroupId field to given value.

### HasOcaGroupId

`func (o *OrderStatus) HasOcaGroupId() bool`

HasOcaGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


