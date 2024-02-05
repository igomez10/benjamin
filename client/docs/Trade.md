# Trade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **string** | execution identifier for the order | [optional] 
**Symbol** | Pointer to **string** | Underlying Symbol | [optional] 
**Side** | Pointer to **string** | The side of the market of the order.   * B - Buy contract near posted ask price   * S - Sell contract near posted bid price   * X - Option expired  | [optional] 
**OrderDescription** | Pointer to **string** | Formatted description of the order \&quot;%side% %size% @ %price% on %exchange%\&quot;. | [optional] 
**TradeTime** | Pointer to **string** | Time of Status update in format \&quot;YYYYMMDD-hh:mm:ss\&quot;. | [optional] 
**TradeTimeR** | Pointer to **float32** | Time of status update in format unix time. | [optional] 
**Size** | Pointer to **string** | Quantity of the order | [optional] 
**Price** | Pointer to **string** | Average Price | [optional] 
**OrderRef** | Pointer to **string** | User defined string used to identify the order. Value is set using \&quot;cOID\&quot; field while placing an order. | [optional] 
**Submitter** | Pointer to **string** | User that submitted order | [optional] 
**Exchange** | Pointer to **string** | Exchange or venue of order | [optional] 
**Commission** | Pointer to **float32** | Commission of the order | [optional] 
**NetAmount** | Pointer to **float32** | Net cost of the order, including contract multiplier and quantity. | [optional] 
**Account** | Pointer to **string** | accountCode | [optional] 
**AcountCode** | Pointer to **string** | Account Number | [optional] 
**CompanyName** | Pointer to **string** | Contracts company name | [optional] 
**ContractDescription1** | Pointer to **string** | Format contract name | [optional] 
**SecType** | Pointer to **string** | Asset class | [optional] 
**Conid** | Pointer to **string** | IBKR&#39;s contract identifier | [optional] 
**Conidex** | Pointer to **string** | conid and exchange. Format supports conid or conid@exchange | [optional] 
**Position** | Pointer to **string** | Total quantity owned for this contract | [optional] 
**ClearingId** | Pointer to **string** | Firm which will settle the trade. For IBExecution customers only. | [optional] 
**ClearingName** | Pointer to **string** | Specifies the true beneficiary of the order. For IBExecution customers only. | [optional] 
**LiquidationTrade** | Pointer to **float32** | If order adds liquidity to the market. | [optional] 

## Methods

### NewTrade

`func NewTrade() *Trade`

NewTrade instantiates a new Trade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeWithDefaults

`func NewTradeWithDefaults() *Trade`

NewTradeWithDefaults instantiates a new Trade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *Trade) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *Trade) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *Trade) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *Trade) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### GetSymbol

`func (o *Trade) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Trade) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Trade) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Trade) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetSide

`func (o *Trade) GetSide() string`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Trade) GetSideOk() (*string, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Trade) SetSide(v string)`

SetSide sets Side field to given value.

### HasSide

`func (o *Trade) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetOrderDescription

`func (o *Trade) GetOrderDescription() string`

GetOrderDescription returns the OrderDescription field if non-nil, zero value otherwise.

### GetOrderDescriptionOk

`func (o *Trade) GetOrderDescriptionOk() (*string, bool)`

GetOrderDescriptionOk returns a tuple with the OrderDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDescription

`func (o *Trade) SetOrderDescription(v string)`

SetOrderDescription sets OrderDescription field to given value.

### HasOrderDescription

`func (o *Trade) HasOrderDescription() bool`

HasOrderDescription returns a boolean if a field has been set.

### GetTradeTime

`func (o *Trade) GetTradeTime() string`

GetTradeTime returns the TradeTime field if non-nil, zero value otherwise.

### GetTradeTimeOk

`func (o *Trade) GetTradeTimeOk() (*string, bool)`

GetTradeTimeOk returns a tuple with the TradeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTime

`func (o *Trade) SetTradeTime(v string)`

SetTradeTime sets TradeTime field to given value.

### HasTradeTime

`func (o *Trade) HasTradeTime() bool`

HasTradeTime returns a boolean if a field has been set.

### GetTradeTimeR

`func (o *Trade) GetTradeTimeR() float32`

GetTradeTimeR returns the TradeTimeR field if non-nil, zero value otherwise.

### GetTradeTimeROk

`func (o *Trade) GetTradeTimeROk() (*float32, bool)`

GetTradeTimeROk returns a tuple with the TradeTimeR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeTimeR

`func (o *Trade) SetTradeTimeR(v float32)`

SetTradeTimeR sets TradeTimeR field to given value.

### HasTradeTimeR

`func (o *Trade) HasTradeTimeR() bool`

HasTradeTimeR returns a boolean if a field has been set.

### GetSize

`func (o *Trade) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Trade) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Trade) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *Trade) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPrice

`func (o *Trade) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Trade) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Trade) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Trade) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetOrderRef

`func (o *Trade) GetOrderRef() string`

GetOrderRef returns the OrderRef field if non-nil, zero value otherwise.

### GetOrderRefOk

`func (o *Trade) GetOrderRefOk() (*string, bool)`

GetOrderRefOk returns a tuple with the OrderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRef

`func (o *Trade) SetOrderRef(v string)`

SetOrderRef sets OrderRef field to given value.

### HasOrderRef

`func (o *Trade) HasOrderRef() bool`

HasOrderRef returns a boolean if a field has been set.

### GetSubmitter

`func (o *Trade) GetSubmitter() string`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *Trade) GetSubmitterOk() (*string, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *Trade) SetSubmitter(v string)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *Trade) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetExchange

`func (o *Trade) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *Trade) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *Trade) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *Trade) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetCommission

`func (o *Trade) GetCommission() float32`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *Trade) GetCommissionOk() (*float32, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *Trade) SetCommission(v float32)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *Trade) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetNetAmount

`func (o *Trade) GetNetAmount() float32`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *Trade) GetNetAmountOk() (*float32, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *Trade) SetNetAmount(v float32)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *Trade) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetAccount

`func (o *Trade) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Trade) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Trade) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Trade) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAcountCode

`func (o *Trade) GetAcountCode() string`

GetAcountCode returns the AcountCode field if non-nil, zero value otherwise.

### GetAcountCodeOk

`func (o *Trade) GetAcountCodeOk() (*string, bool)`

GetAcountCodeOk returns a tuple with the AcountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcountCode

`func (o *Trade) SetAcountCode(v string)`

SetAcountCode sets AcountCode field to given value.

### HasAcountCode

`func (o *Trade) HasAcountCode() bool`

HasAcountCode returns a boolean if a field has been set.

### GetCompanyName

`func (o *Trade) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Trade) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Trade) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Trade) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetContractDescription1

`func (o *Trade) GetContractDescription1() string`

GetContractDescription1 returns the ContractDescription1 field if non-nil, zero value otherwise.

### GetContractDescription1Ok

`func (o *Trade) GetContractDescription1Ok() (*string, bool)`

GetContractDescription1Ok returns a tuple with the ContractDescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDescription1

`func (o *Trade) SetContractDescription1(v string)`

SetContractDescription1 sets ContractDescription1 field to given value.

### HasContractDescription1

`func (o *Trade) HasContractDescription1() bool`

HasContractDescription1 returns a boolean if a field has been set.

### GetSecType

`func (o *Trade) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *Trade) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *Trade) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *Trade) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetConid

`func (o *Trade) GetConid() string`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *Trade) GetConidOk() (*string, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *Trade) SetConid(v string)`

SetConid sets Conid field to given value.

### HasConid

`func (o *Trade) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetConidex

`func (o *Trade) GetConidex() string`

GetConidex returns the Conidex field if non-nil, zero value otherwise.

### GetConidexOk

`func (o *Trade) GetConidexOk() (*string, bool)`

GetConidexOk returns a tuple with the Conidex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConidex

`func (o *Trade) SetConidex(v string)`

SetConidex sets Conidex field to given value.

### HasConidex

`func (o *Trade) HasConidex() bool`

HasConidex returns a boolean if a field has been set.

### GetPosition

`func (o *Trade) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Trade) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Trade) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Trade) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetClearingId

`func (o *Trade) GetClearingId() string`

GetClearingId returns the ClearingId field if non-nil, zero value otherwise.

### GetClearingIdOk

`func (o *Trade) GetClearingIdOk() (*string, bool)`

GetClearingIdOk returns a tuple with the ClearingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingId

`func (o *Trade) SetClearingId(v string)`

SetClearingId sets ClearingId field to given value.

### HasClearingId

`func (o *Trade) HasClearingId() bool`

HasClearingId returns a boolean if a field has been set.

### GetClearingName

`func (o *Trade) GetClearingName() string`

GetClearingName returns the ClearingName field if non-nil, zero value otherwise.

### GetClearingNameOk

`func (o *Trade) GetClearingNameOk() (*string, bool)`

GetClearingNameOk returns a tuple with the ClearingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingName

`func (o *Trade) SetClearingName(v string)`

SetClearingName sets ClearingName field to given value.

### HasClearingName

`func (o *Trade) HasClearingName() bool`

HasClearingName returns a boolean if a field has been set.

### GetLiquidationTrade

`func (o *Trade) GetLiquidationTrade() float32`

GetLiquidationTrade returns the LiquidationTrade field if non-nil, zero value otherwise.

### GetLiquidationTradeOk

`func (o *Trade) GetLiquidationTradeOk() (*float32, bool)`

GetLiquidationTradeOk returns a tuple with the LiquidationTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationTrade

`func (o *Trade) SetLiquidationTrade(v float32)`

SetLiquidationTrade sets LiquidationTrade field to given value.

### HasLiquidationTrade

`func (o *Trade) HasLiquidationTrade() bool`

HasLiquidationTrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

