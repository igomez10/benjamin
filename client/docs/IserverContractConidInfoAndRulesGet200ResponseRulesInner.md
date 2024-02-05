# IserverContractConidInfoAndRulesGet200ResponseRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgoEligible** | Pointer to **bool** | Contract supports algo orders | [optional] 
**CanTradeAcctIds** | Pointer to [**[]IserverContractRulesPost200ResponseRulesInnerCanTradeAcctIdsInner**](IserverContractRulesPost200ResponseRulesInnerCanTradeAcctIdsInner.md) |  | [optional] 
**Error** | Pointer to **string** | Returns a description on any errors with order presets | [optional] 
**OrderTypes** | Pointer to [**[]IserverContractRulesPost200ResponseRulesInnerOrderTypesInner**](IserverContractRulesPost200ResponseRulesInnerOrderTypesInner.md) |  | [optional] 
**IbalgoTypes** | Pointer to [**[]IserverContractRulesPost200ResponseRulesInnerIbalgoTypesInner**](IserverContractRulesPost200ResponseRulesInnerIbalgoTypesInner.md) |  | [optional] 
**FraqTypes** | Pointer to [**[]IserverContractRulesPost200ResponseRulesInnerFraqTypesInner**](IserverContractRulesPost200ResponseRulesInnerFraqTypesInner.md) |  | [optional] 
**CqtTypes** | Pointer to [**[]IserverContractRulesPost200ResponseRulesInnerCqtTypesInner**](IserverContractRulesPost200ResponseRulesInnerCqtTypesInner.md) |  | [optional] 
**OrderDefaults** | Pointer to [**[]IserverContractRulesPost200ResponseRulesInnerOrderDefaultsInner**](IserverContractRulesPost200ResponseRulesInnerOrderDefaultsInner.md) | If object returned will provide the defaults based on user settings | [optional] 
**OrderTypesOutside** | Pointer to [**[]IserverContractRulesPost200ResponseRulesInnerOrderTypesOutsideInner**](IserverContractRulesPost200ResponseRulesInnerOrderTypesOutsideInner.md) |  | [optional] 
**DefaultSize** | Pointer to **int32** | Default quantity | [optional] 
**CashSize** | Pointer to **int32** | cash value | [optional] 
**SizeIncrement** | Pointer to **int32** | increment quantity value | [optional] 
**TifTypes** | Pointer to [**[]IserverContractRulesPost200ResponseRulesInnerTifTypesInner**](IserverContractRulesPost200ResponseRulesInnerTifTypesInner.md) |  | [optional] 
**DefaultTIF** | Pointer to **string** | Default time in force value | [optional] 
**LimitPrice** | Pointer to **float32** | Limit price | [optional] 
**Stopprice** | Pointer to **float32** | Stop price | [optional] 
**OrderOrigination** | Pointer to **float32** | Order origin designation for US securities options and Options Clearing Corporation | [optional] 
**Preview** | Pointer to **bool** | order preview required | [optional] 
**DisplaySize** | Pointer to **float32** |  | [optional] 
**FraqInt** | Pointer to **float32** | decimal places for fractional order size | [optional] 
**CashCcy** | Pointer to **string** | Cash currency for the contract | [optional] 
**CashQtyIncr** | Pointer to **float32** | Increment value for cash quantity | [optional] 
**PriceMagnifier** | Pointer to **float32** | Price Magnifier | [optional] 
**NegativeCapable** | Pointer to **bool** | trading negative price support | [optional] 
**Increment** | Pointer to **float32** | Price increment value | [optional] 
**IncrementDigits** | Pointer to **int32** | Number of digits for price increment | [optional] 

## Methods

### NewIserverContractConidInfoAndRulesGet200ResponseRulesInner

`func NewIserverContractConidInfoAndRulesGet200ResponseRulesInner() *IserverContractConidInfoAndRulesGet200ResponseRulesInner`

NewIserverContractConidInfoAndRulesGet200ResponseRulesInner instantiates a new IserverContractConidInfoAndRulesGet200ResponseRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIserverContractConidInfoAndRulesGet200ResponseRulesInnerWithDefaults

`func NewIserverContractConidInfoAndRulesGet200ResponseRulesInnerWithDefaults() *IserverContractConidInfoAndRulesGet200ResponseRulesInner`

NewIserverContractConidInfoAndRulesGet200ResponseRulesInnerWithDefaults instantiates a new IserverContractConidInfoAndRulesGet200ResponseRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgoEligible

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetAlgoEligible() bool`

GetAlgoEligible returns the AlgoEligible field if non-nil, zero value otherwise.

### GetAlgoEligibleOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetAlgoEligibleOk() (*bool, bool)`

GetAlgoEligibleOk returns a tuple with the AlgoEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgoEligible

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetAlgoEligible(v bool)`

SetAlgoEligible sets AlgoEligible field to given value.

### HasAlgoEligible

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasAlgoEligible() bool`

HasAlgoEligible returns a boolean if a field has been set.

### GetCanTradeAcctIds

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCanTradeAcctIds() []IserverContractRulesPost200ResponseRulesInnerCanTradeAcctIdsInner`

GetCanTradeAcctIds returns the CanTradeAcctIds field if non-nil, zero value otherwise.

### GetCanTradeAcctIdsOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCanTradeAcctIdsOk() (*[]IserverContractRulesPost200ResponseRulesInnerCanTradeAcctIdsInner, bool)`

GetCanTradeAcctIdsOk returns a tuple with the CanTradeAcctIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanTradeAcctIds

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetCanTradeAcctIds(v []IserverContractRulesPost200ResponseRulesInnerCanTradeAcctIdsInner)`

SetCanTradeAcctIds sets CanTradeAcctIds field to given value.

### HasCanTradeAcctIds

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasCanTradeAcctIds() bool`

HasCanTradeAcctIds returns a boolean if a field has been set.

### GetError

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOrderTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetOrderTypes() []IserverContractRulesPost200ResponseRulesInnerOrderTypesInner`

GetOrderTypes returns the OrderTypes field if non-nil, zero value otherwise.

### GetOrderTypesOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetOrderTypesOk() (*[]IserverContractRulesPost200ResponseRulesInnerOrderTypesInner, bool)`

GetOrderTypesOk returns a tuple with the OrderTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetOrderTypes(v []IserverContractRulesPost200ResponseRulesInnerOrderTypesInner)`

SetOrderTypes sets OrderTypes field to given value.

### HasOrderTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasOrderTypes() bool`

HasOrderTypes returns a boolean if a field has been set.

### GetIbalgoTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetIbalgoTypes() []IserverContractRulesPost200ResponseRulesInnerIbalgoTypesInner`

GetIbalgoTypes returns the IbalgoTypes field if non-nil, zero value otherwise.

### GetIbalgoTypesOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetIbalgoTypesOk() (*[]IserverContractRulesPost200ResponseRulesInnerIbalgoTypesInner, bool)`

GetIbalgoTypesOk returns a tuple with the IbalgoTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbalgoTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetIbalgoTypes(v []IserverContractRulesPost200ResponseRulesInnerIbalgoTypesInner)`

SetIbalgoTypes sets IbalgoTypes field to given value.

### HasIbalgoTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasIbalgoTypes() bool`

HasIbalgoTypes returns a boolean if a field has been set.

### GetFraqTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetFraqTypes() []IserverContractRulesPost200ResponseRulesInnerFraqTypesInner`

GetFraqTypes returns the FraqTypes field if non-nil, zero value otherwise.

### GetFraqTypesOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetFraqTypesOk() (*[]IserverContractRulesPost200ResponseRulesInnerFraqTypesInner, bool)`

GetFraqTypesOk returns a tuple with the FraqTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraqTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetFraqTypes(v []IserverContractRulesPost200ResponseRulesInnerFraqTypesInner)`

SetFraqTypes sets FraqTypes field to given value.

### HasFraqTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasFraqTypes() bool`

HasFraqTypes returns a boolean if a field has been set.

### GetCqtTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCqtTypes() []IserverContractRulesPost200ResponseRulesInnerCqtTypesInner`

GetCqtTypes returns the CqtTypes field if non-nil, zero value otherwise.

### GetCqtTypesOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCqtTypesOk() (*[]IserverContractRulesPost200ResponseRulesInnerCqtTypesInner, bool)`

GetCqtTypesOk returns a tuple with the CqtTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCqtTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetCqtTypes(v []IserverContractRulesPost200ResponseRulesInnerCqtTypesInner)`

SetCqtTypes sets CqtTypes field to given value.

### HasCqtTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasCqtTypes() bool`

HasCqtTypes returns a boolean if a field has been set.

### GetOrderDefaults

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetOrderDefaults() []IserverContractRulesPost200ResponseRulesInnerOrderDefaultsInner`

GetOrderDefaults returns the OrderDefaults field if non-nil, zero value otherwise.

### GetOrderDefaultsOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetOrderDefaultsOk() (*[]IserverContractRulesPost200ResponseRulesInnerOrderDefaultsInner, bool)`

GetOrderDefaultsOk returns a tuple with the OrderDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderDefaults

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetOrderDefaults(v []IserverContractRulesPost200ResponseRulesInnerOrderDefaultsInner)`

SetOrderDefaults sets OrderDefaults field to given value.

### HasOrderDefaults

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasOrderDefaults() bool`

HasOrderDefaults returns a boolean if a field has been set.

### GetOrderTypesOutside

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetOrderTypesOutside() []IserverContractRulesPost200ResponseRulesInnerOrderTypesOutsideInner`

GetOrderTypesOutside returns the OrderTypesOutside field if non-nil, zero value otherwise.

### GetOrderTypesOutsideOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetOrderTypesOutsideOk() (*[]IserverContractRulesPost200ResponseRulesInnerOrderTypesOutsideInner, bool)`

GetOrderTypesOutsideOk returns a tuple with the OrderTypesOutside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTypesOutside

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetOrderTypesOutside(v []IserverContractRulesPost200ResponseRulesInnerOrderTypesOutsideInner)`

SetOrderTypesOutside sets OrderTypesOutside field to given value.

### HasOrderTypesOutside

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasOrderTypesOutside() bool`

HasOrderTypesOutside returns a boolean if a field has been set.

### GetDefaultSize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetDefaultSize() int32`

GetDefaultSize returns the DefaultSize field if non-nil, zero value otherwise.

### GetDefaultSizeOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetDefaultSizeOk() (*int32, bool)`

GetDefaultSizeOk returns a tuple with the DefaultSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetDefaultSize(v int32)`

SetDefaultSize sets DefaultSize field to given value.

### HasDefaultSize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasDefaultSize() bool`

HasDefaultSize returns a boolean if a field has been set.

### GetCashSize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCashSize() int32`

GetCashSize returns the CashSize field if non-nil, zero value otherwise.

### GetCashSizeOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCashSizeOk() (*int32, bool)`

GetCashSizeOk returns a tuple with the CashSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashSize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetCashSize(v int32)`

SetCashSize sets CashSize field to given value.

### HasCashSize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasCashSize() bool`

HasCashSize returns a boolean if a field has been set.

### GetSizeIncrement

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetSizeIncrement() int32`

GetSizeIncrement returns the SizeIncrement field if non-nil, zero value otherwise.

### GetSizeIncrementOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetSizeIncrementOk() (*int32, bool)`

GetSizeIncrementOk returns a tuple with the SizeIncrement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeIncrement

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetSizeIncrement(v int32)`

SetSizeIncrement sets SizeIncrement field to given value.

### HasSizeIncrement

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasSizeIncrement() bool`

HasSizeIncrement returns a boolean if a field has been set.

### GetTifTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetTifTypes() []IserverContractRulesPost200ResponseRulesInnerTifTypesInner`

GetTifTypes returns the TifTypes field if non-nil, zero value otherwise.

### GetTifTypesOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetTifTypesOk() (*[]IserverContractRulesPost200ResponseRulesInnerTifTypesInner, bool)`

GetTifTypesOk returns a tuple with the TifTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTifTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetTifTypes(v []IserverContractRulesPost200ResponseRulesInnerTifTypesInner)`

SetTifTypes sets TifTypes field to given value.

### HasTifTypes

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasTifTypes() bool`

HasTifTypes returns a boolean if a field has been set.

### GetDefaultTIF

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetDefaultTIF() string`

GetDefaultTIF returns the DefaultTIF field if non-nil, zero value otherwise.

### GetDefaultTIFOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetDefaultTIFOk() (*string, bool)`

GetDefaultTIFOk returns a tuple with the DefaultTIF field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTIF

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetDefaultTIF(v string)`

SetDefaultTIF sets DefaultTIF field to given value.

### HasDefaultTIF

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasDefaultTIF() bool`

HasDefaultTIF returns a boolean if a field has been set.

### GetLimitPrice

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetLimitPrice() float32`

GetLimitPrice returns the LimitPrice field if non-nil, zero value otherwise.

### GetLimitPriceOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetLimitPriceOk() (*float32, bool)`

GetLimitPriceOk returns a tuple with the LimitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitPrice

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetLimitPrice(v float32)`

SetLimitPrice sets LimitPrice field to given value.

### HasLimitPrice

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasLimitPrice() bool`

HasLimitPrice returns a boolean if a field has been set.

### GetStopprice

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetStopprice() float32`

GetStopprice returns the Stopprice field if non-nil, zero value otherwise.

### GetStoppriceOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetStoppriceOk() (*float32, bool)`

GetStoppriceOk returns a tuple with the Stopprice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopprice

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetStopprice(v float32)`

SetStopprice sets Stopprice field to given value.

### HasStopprice

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasStopprice() bool`

HasStopprice returns a boolean if a field has been set.

### GetOrderOrigination

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetOrderOrigination() float32`

GetOrderOrigination returns the OrderOrigination field if non-nil, zero value otherwise.

### GetOrderOriginationOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetOrderOriginationOk() (*float32, bool)`

GetOrderOriginationOk returns a tuple with the OrderOrigination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderOrigination

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetOrderOrigination(v float32)`

SetOrderOrigination sets OrderOrigination field to given value.

### HasOrderOrigination

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasOrderOrigination() bool`

HasOrderOrigination returns a boolean if a field has been set.

### GetPreview

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetPreview() bool`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetPreviewOk() (*bool, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetPreview(v bool)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetDisplaySize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetDisplaySize() float32`

GetDisplaySize returns the DisplaySize field if non-nil, zero value otherwise.

### GetDisplaySizeOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetDisplaySizeOk() (*float32, bool)`

GetDisplaySizeOk returns a tuple with the DisplaySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetDisplaySize(v float32)`

SetDisplaySize sets DisplaySize field to given value.

### HasDisplaySize

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasDisplaySize() bool`

HasDisplaySize returns a boolean if a field has been set.

### GetFraqInt

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetFraqInt() float32`

GetFraqInt returns the FraqInt field if non-nil, zero value otherwise.

### GetFraqIntOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetFraqIntOk() (*float32, bool)`

GetFraqIntOk returns a tuple with the FraqInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFraqInt

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetFraqInt(v float32)`

SetFraqInt sets FraqInt field to given value.

### HasFraqInt

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasFraqInt() bool`

HasFraqInt returns a boolean if a field has been set.

### GetCashCcy

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCashCcy() string`

GetCashCcy returns the CashCcy field if non-nil, zero value otherwise.

### GetCashCcyOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCashCcyOk() (*string, bool)`

GetCashCcyOk returns a tuple with the CashCcy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashCcy

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetCashCcy(v string)`

SetCashCcy sets CashCcy field to given value.

### HasCashCcy

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasCashCcy() bool`

HasCashCcy returns a boolean if a field has been set.

### GetCashQtyIncr

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCashQtyIncr() float32`

GetCashQtyIncr returns the CashQtyIncr field if non-nil, zero value otherwise.

### GetCashQtyIncrOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetCashQtyIncrOk() (*float32, bool)`

GetCashQtyIncrOk returns a tuple with the CashQtyIncr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashQtyIncr

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetCashQtyIncr(v float32)`

SetCashQtyIncr sets CashQtyIncr field to given value.

### HasCashQtyIncr

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasCashQtyIncr() bool`

HasCashQtyIncr returns a boolean if a field has been set.

### GetPriceMagnifier

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetPriceMagnifier() float32`

GetPriceMagnifier returns the PriceMagnifier field if non-nil, zero value otherwise.

### GetPriceMagnifierOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetPriceMagnifierOk() (*float32, bool)`

GetPriceMagnifierOk returns a tuple with the PriceMagnifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMagnifier

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetPriceMagnifier(v float32)`

SetPriceMagnifier sets PriceMagnifier field to given value.

### HasPriceMagnifier

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasPriceMagnifier() bool`

HasPriceMagnifier returns a boolean if a field has been set.

### GetNegativeCapable

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetNegativeCapable() bool`

GetNegativeCapable returns the NegativeCapable field if non-nil, zero value otherwise.

### GetNegativeCapableOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetNegativeCapableOk() (*bool, bool)`

GetNegativeCapableOk returns a tuple with the NegativeCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCapable

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetNegativeCapable(v bool)`

SetNegativeCapable sets NegativeCapable field to given value.

### HasNegativeCapable

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasNegativeCapable() bool`

HasNegativeCapable returns a boolean if a field has been set.

### GetIncrement

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetIncrement() float32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetIncrementOk() (*float32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetIncrement(v float32)`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.

### GetIncrementDigits

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetIncrementDigits() int32`

GetIncrementDigits returns the IncrementDigits field if non-nil, zero value otherwise.

### GetIncrementDigitsOk

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) GetIncrementDigitsOk() (*int32, bool)`

GetIncrementDigitsOk returns a tuple with the IncrementDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementDigits

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) SetIncrementDigits(v int32)`

SetIncrementDigits sets IncrementDigits field to given value.

### HasIncrementDigits

`func (o *IserverContractConidInfoAndRulesGet200ResponseRulesInner) HasIncrementDigits() bool`

HasIncrementDigits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


