# OrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctId** | Pointer to **string** | acctId is optional. It should be one of the accounts returned by /iserver/accounts. If not passed, the first one in the list is selected.  | [optional] 
**Conid** | Pointer to **int32** | conid is the identifier of the security you want to trade, you can find the conid with /iserver/secdef/search.  | [optional] 
**Conidex** | Pointer to **string** | Conid and Exchange - Can be used instead of conid when specifying the contract identifier of a security.  | [optional] 
**SecType** | Pointer to **string** | The contract-identifier (conid) and security type (type) specified as a concatenated value, conid:type | [optional] 
**COID** | Pointer to **string** | Customer Order ID. An arbitrary string that can be used to identify the order, e.g \&quot;my-fb-order\&quot;. The value must be unique for a 24h span. Please do not set this value for child orders when placing a bracket order.  | [optional] 
**ParentId** | Pointer to **string** | Only specify for child orders when placing bracket orders. The parentId for the child order(s) must be equal to the cOId (customer order id) of the parent.  | [optional] 
**OrderType** | Pointer to **string** | The order-type determines what type of order you want to send.   * LMT - A limit order is an order to buy or sell at the specified price or better.   * MKT - A market order is an order to buy or sell at the markets current NBBO.   * STP - A stop order becomes a market order once the specified stop price is attained or penetrated.   * STOP_LIMIT - A stop limit order becomes a limit order once the specified stop price is attained or penetrated.   * MIDPRICE - A MidPrice order attempts to fill at the current midpoint of the NBBO or better.   * TRAIL - A sell trailing stop order sets the stop price at a fixed amount below the market price with an attached \&quot;trailing\&quot; amount. See more details here: https://ndcdyn.interactivebrokers.com/en/index.php?f&#x3D;605   * TRAILLMT - A trailing stop limit order is designed to allow an investor to specify a limit on the maximum possible loss, without setting a limit on the maximum possible gain.     See more details here: https://ndcdyn.interactivebrokers.com/en/index.php?f&#x3D;606  | [optional] 
**ListingExchange** | Pointer to **string** | listingExchange is optional. By default we use \&quot;SMART\&quot; routing. Possible values are available via the endpoint: /iserver/contract/{conid}/info, see **valid_exchange** e.g: SMART,AMEX,NYSE,CBOE,ISE,CHX,ARCA,ISLAND,DRCTEDGE,BEX,BATS,EDGEA,CSFBALGO,JE FFALGO,BYX,IEX,FOXRIVER,TPLUS1,NYSENAT,PSX  | [optional] 
**IsSingleGroup** | Pointer to **bool** | set to true if you want to place a single group orders(OCA)  | [optional] 
**OutsideRTH** | Pointer to **bool** | set to true if the order can be executed outside regular trading hours.  | [optional] 
**Price** | Pointer to **float32** | optional if order is LMT, or STOP_LIMIT, this is the limit price. For STP|TRAIL this is the stop price. For MIDPRICE this is the option price cap.  | [optional] 
**AuxPrice** | Pointer to **map[string]interface{}** | optional if order is STOP_LIMIT|TRAILLMT, this is the stop price. You must specify both price and auxPrice for STOP_LIMIT|TRAILLMT orders.  | [optional] 
**Side** | Pointer to **string** | SELL or BUY | [optional] 
**Ticker** | Pointer to **string** | This is the  underlying symbol for the contract.  | [optional] 
**Tif** | Pointer to **string** | The Time-In-Force determines how long the order remains active on the market.   * GTC - use Good-Till-Cancel for orders to remain active until it executes or cancelled.   * OPG - use Open-Price-Guarantee for Limit-On-Open (LOO) or Market-On-Open (MOO) orders.   * DAY - if not executed a Day order will automatically cancel at the end of the markets regular trading hours.   * IOC - any portion of an Immediate-or-Cancel order that is not filled as soon as it becomes available in the market is cancelled.  | [optional] 
**TrailingAmt** | Pointer to **float32** | optional if order is TRAIL, or TRAILLMT. When trailingType is amt, this is the trailing amount, when trailingType is %, it means percentage. You must specify both trailingType and trailingAmt for TRAIL and TRAILLMT order  | [optional] 
**TrailingType** | Pointer to **string** | This is the trailing type for trailing amount. We only support two types here: amt or %. You must specify both trailingType and trailingAmt for TRAIL and TRAILLMT order  | [optional] 
**Referrer** | Pointer to **string** | Custom order reference  | [optional] 
**Quantity** | Pointer to **float32** | Usually integer, for some special cases such as fractional orders can specify as a float, e.g. quantity &#x3D; 0.001. In some special cases quantity is not specified, such as when using &#39;cashQty&#39; or &#39;fxQty&#39;.  | [optional] 
**CashQty** | Pointer to **float32** | Cash Quantity - used to specify the monetary value of an order instead of the number of shares. When using &#39;cashQty&#39; don&#39;t specify &#39;quantity&#39; Orders that express size using a monetary value, e.g. cash quantity can result in fractional shares and are provided on a non-guaranteed basis. The system simulates the order by canceling it once the specified amount is spent (for buy orders) or collected (for sell orders). In addition to the monetary value, the order uses a maximum size that is calculated using the Cash Quantity Estimated Factor, which can be modified in Order Presets.  | [optional] 
**FxQty** | Pointer to **float32** | double number, this is the cash quantity field which can only be used for Currency Conversion Orders. When using &#39;fxQty&#39; don&#39;t specify &#39;quantity&#39;.  | [optional] 
**UseAdaptive** | Pointer to **bool** | If true, the system will use the Price Management Algo to submit the order. https://www.interactivebrokers.com/en/index.php?f&#x3D;43423  | [optional] 
**IsCcyConv** | Pointer to **bool** | set to true if the order is a FX conversion order  | [optional] 
**AllocationMethod** | Pointer to **string** | Set the allocation method when placing an order using an FA account for a group Possible allocation methods are \&quot;NetLiquidity\&quot;, \&quot;AvailableEquity\&quot;, \&quot;EqualQuantity\&quot; and \&quot;PctChange\&quot;.  | [optional] 
**Strategy** | Pointer to **string** | Specify which IB Algo algorithm to use for this order.  | [optional] 
**StrategyParameters** | Pointer to **map[string]interface{}** | The IB Algo parameters for the specified algorithm.  | [optional] 

## Methods

### NewOrderRequest

`func NewOrderRequest() *OrderRequest`

NewOrderRequest instantiates a new OrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRequestWithDefaults

`func NewOrderRequestWithDefaults() *OrderRequest`

NewOrderRequestWithDefaults instantiates a new OrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctId

`func (o *OrderRequest) GetAcctId() string`

GetAcctId returns the AcctId field if non-nil, zero value otherwise.

### GetAcctIdOk

`func (o *OrderRequest) GetAcctIdOk() (*string, bool)`

GetAcctIdOk returns a tuple with the AcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctId

`func (o *OrderRequest) SetAcctId(v string)`

SetAcctId sets AcctId field to given value.

### HasAcctId

`func (o *OrderRequest) HasAcctId() bool`

HasAcctId returns a boolean if a field has been set.

### GetConid

`func (o *OrderRequest) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *OrderRequest) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *OrderRequest) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *OrderRequest) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetConidex

`func (o *OrderRequest) GetConidex() string`

GetConidex returns the Conidex field if non-nil, zero value otherwise.

### GetConidexOk

`func (o *OrderRequest) GetConidexOk() (*string, bool)`

GetConidexOk returns a tuple with the Conidex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConidex

`func (o *OrderRequest) SetConidex(v string)`

SetConidex sets Conidex field to given value.

### HasConidex

`func (o *OrderRequest) HasConidex() bool`

HasConidex returns a boolean if a field has been set.

### GetSecType

`func (o *OrderRequest) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *OrderRequest) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *OrderRequest) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *OrderRequest) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetCOID

`func (o *OrderRequest) GetCOID() string`

GetCOID returns the COID field if non-nil, zero value otherwise.

### GetCOIDOk

`func (o *OrderRequest) GetCOIDOk() (*string, bool)`

GetCOIDOk returns a tuple with the COID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCOID

`func (o *OrderRequest) SetCOID(v string)`

SetCOID sets COID field to given value.

### HasCOID

`func (o *OrderRequest) HasCOID() bool`

HasCOID returns a boolean if a field has been set.

### GetParentId

`func (o *OrderRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *OrderRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *OrderRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *OrderRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderRequest) GetOrderType() string`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderRequest) GetOrderTypeOk() (*string, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderRequest) SetOrderType(v string)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderRequest) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetListingExchange

`func (o *OrderRequest) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *OrderRequest) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *OrderRequest) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *OrderRequest) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetIsSingleGroup

`func (o *OrderRequest) GetIsSingleGroup() bool`

GetIsSingleGroup returns the IsSingleGroup field if non-nil, zero value otherwise.

### GetIsSingleGroupOk

`func (o *OrderRequest) GetIsSingleGroupOk() (*bool, bool)`

GetIsSingleGroupOk returns a tuple with the IsSingleGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleGroup

`func (o *OrderRequest) SetIsSingleGroup(v bool)`

SetIsSingleGroup sets IsSingleGroup field to given value.

### HasIsSingleGroup

`func (o *OrderRequest) HasIsSingleGroup() bool`

HasIsSingleGroup returns a boolean if a field has been set.

### GetOutsideRTH

`func (o *OrderRequest) GetOutsideRTH() bool`

GetOutsideRTH returns the OutsideRTH field if non-nil, zero value otherwise.

### GetOutsideRTHOk

`func (o *OrderRequest) GetOutsideRTHOk() (*bool, bool)`

GetOutsideRTHOk returns a tuple with the OutsideRTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRTH

`func (o *OrderRequest) SetOutsideRTH(v bool)`

SetOutsideRTH sets OutsideRTH field to given value.

### HasOutsideRTH

`func (o *OrderRequest) HasOutsideRTH() bool`

HasOutsideRTH returns a boolean if a field has been set.

### GetPrice

`func (o *OrderRequest) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderRequest) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderRequest) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetAuxPrice

`func (o *OrderRequest) GetAuxPrice() map[string]interface{}`

GetAuxPrice returns the AuxPrice field if non-nil, zero value otherwise.

### GetAuxPriceOk

`func (o *OrderRequest) GetAuxPriceOk() (*map[string]interface{}, bool)`

GetAuxPriceOk returns a tuple with the AuxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuxPrice

`func (o *OrderRequest) SetAuxPrice(v map[string]interface{})`

SetAuxPrice sets AuxPrice field to given value.

### HasAuxPrice

`func (o *OrderRequest) HasAuxPrice() bool`

HasAuxPrice returns a boolean if a field has been set.

### GetSide

`func (o *OrderRequest) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *OrderRequest) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *OrderRequest) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *OrderRequest) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetTicker

`func (o *OrderRequest) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *OrderRequest) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *OrderRequest) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *OrderRequest) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetTif

`func (o *OrderRequest) GetTif() string`

GetTif returns the Tif field if non-nil, zero value otherwise.

### GetTifOk

`func (o *OrderRequest) GetTifOk() (*string, bool)`

GetTifOk returns a tuple with the Tif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTif

`func (o *OrderRequest) SetTif(v string)`

SetTif sets Tif field to given value.

### HasTif

`func (o *OrderRequest) HasTif() bool`

HasTif returns a boolean if a field has been set.

### GetTrailingAmt

`func (o *OrderRequest) GetTrailingAmt() float32`

GetTrailingAmt returns the TrailingAmt field if non-nil, zero value otherwise.

### GetTrailingAmtOk

`func (o *OrderRequest) GetTrailingAmtOk() (*float32, bool)`

GetTrailingAmtOk returns a tuple with the TrailingAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingAmt

`func (o *OrderRequest) SetTrailingAmt(v float32)`

SetTrailingAmt sets TrailingAmt field to given value.

### HasTrailingAmt

`func (o *OrderRequest) HasTrailingAmt() bool`

HasTrailingAmt returns a boolean if a field has been set.

### GetTrailingType

`func (o *OrderRequest) GetTrailingType() string`

GetTrailingType returns the TrailingType field if non-nil, zero value otherwise.

### GetTrailingTypeOk

`func (o *OrderRequest) GetTrailingTypeOk() (*string, bool)`

GetTrailingTypeOk returns a tuple with the TrailingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingType

`func (o *OrderRequest) SetTrailingType(v string)`

SetTrailingType sets TrailingType field to given value.

### HasTrailingType

`func (o *OrderRequest) HasTrailingType() bool`

HasTrailingType returns a boolean if a field has been set.

### GetReferrer

`func (o *OrderRequest) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *OrderRequest) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *OrderRequest) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *OrderRequest) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderRequest) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderRequest) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderRequest) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCashQty

`func (o *OrderRequest) GetCashQty() float32`

GetCashQty returns the CashQty field if non-nil, zero value otherwise.

### GetCashQtyOk

`func (o *OrderRequest) GetCashQtyOk() (*float32, bool)`

GetCashQtyOk returns a tuple with the CashQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashQty

`func (o *OrderRequest) SetCashQty(v float32)`

SetCashQty sets CashQty field to given value.

### HasCashQty

`func (o *OrderRequest) HasCashQty() bool`

HasCashQty returns a boolean if a field has been set.

### GetFxQty

`func (o *OrderRequest) GetFxQty() float32`

GetFxQty returns the FxQty field if non-nil, zero value otherwise.

### GetFxQtyOk

`func (o *OrderRequest) GetFxQtyOk() (*float32, bool)`

GetFxQtyOk returns a tuple with the FxQty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxQty

`func (o *OrderRequest) SetFxQty(v float32)`

SetFxQty sets FxQty field to given value.

### HasFxQty

`func (o *OrderRequest) HasFxQty() bool`

HasFxQty returns a boolean if a field has been set.

### GetUseAdaptive

`func (o *OrderRequest) GetUseAdaptive() bool`

GetUseAdaptive returns the UseAdaptive field if non-nil, zero value otherwise.

### GetUseAdaptiveOk

`func (o *OrderRequest) GetUseAdaptiveOk() (*bool, bool)`

GetUseAdaptiveOk returns a tuple with the UseAdaptive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAdaptive

`func (o *OrderRequest) SetUseAdaptive(v bool)`

SetUseAdaptive sets UseAdaptive field to given value.

### HasUseAdaptive

`func (o *OrderRequest) HasUseAdaptive() bool`

HasUseAdaptive returns a boolean if a field has been set.

### GetIsCcyConv

`func (o *OrderRequest) GetIsCcyConv() bool`

GetIsCcyConv returns the IsCcyConv field if non-nil, zero value otherwise.

### GetIsCcyConvOk

`func (o *OrderRequest) GetIsCcyConvOk() (*bool, bool)`

GetIsCcyConvOk returns a tuple with the IsCcyConv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCcyConv

`func (o *OrderRequest) SetIsCcyConv(v bool)`

SetIsCcyConv sets IsCcyConv field to given value.

### HasIsCcyConv

`func (o *OrderRequest) HasIsCcyConv() bool`

HasIsCcyConv returns a boolean if a field has been set.

### GetAllocationMethod

`func (o *OrderRequest) GetAllocationMethod() string`

GetAllocationMethod returns the AllocationMethod field if non-nil, zero value otherwise.

### GetAllocationMethodOk

`func (o *OrderRequest) GetAllocationMethodOk() (*string, bool)`

GetAllocationMethodOk returns a tuple with the AllocationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationMethod

`func (o *OrderRequest) SetAllocationMethod(v string)`

SetAllocationMethod sets AllocationMethod field to given value.

### HasAllocationMethod

`func (o *OrderRequest) HasAllocationMethod() bool`

HasAllocationMethod returns a boolean if a field has been set.

### GetStrategy

`func (o *OrderRequest) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *OrderRequest) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *OrderRequest) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *OrderRequest) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetStrategyParameters

`func (o *OrderRequest) GetStrategyParameters() map[string]interface{}`

GetStrategyParameters returns the StrategyParameters field if non-nil, zero value otherwise.

### GetStrategyParametersOk

`func (o *OrderRequest) GetStrategyParametersOk() (*map[string]interface{}, bool)`

GetStrategyParametersOk returns a tuple with the StrategyParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyParameters

`func (o *OrderRequest) SetStrategyParameters(v map[string]interface{})`

SetStrategyParameters sets StrategyParameters field to given value.

### HasStrategyParameters

`func (o *OrderRequest) HasStrategyParameters() bool`

HasStrategyParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


