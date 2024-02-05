# HistoryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Underlying symbol | [optional] 
**Text** | Pointer to **string** | companyName | [optional] 
**PriceFactor** | Pointer to **int32** | priceFactor is price increment obtained from display rule | [optional] 
**StartTime** | Pointer to **string** | start date time in the format YYYYMMDD-HH:mm:ss | [optional] 
**High** | Pointer to **string** | High value during this time series with format %h/%v/%t. %h is the high price (scaled by priceFactor), %v is volume (volume factor will always be 100 (reported volume &#x3D; actual volume/100)) and %t is minutes from start time of the chart  | [optional] 
**Low** | Pointer to **string** | Low value during this time series with format %l/%v/%t. %l is the low price (scaled by priceFactor), %v is volume (volume factor will always be 100 (reported volume &#x3D; actual volume/100)) and %t is minutes from start time of the chart  | [optional] 
**TimePeriod** | Pointer to **string** | The duration for the historical data request | [optional] 
**BarLength** | Pointer to **int32** | The number of seconds in a bar | [optional] 
**MdAvailability** | Pointer to **string** | Market Data Availability. The field may contain two chars. The first char is the primary code: S &#x3D; Streaming, R &#x3D; Realtime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed. The second char is the secondary code: P &#x3D; Snapshot Available, p &#x3D; Consolidated.  | [optional] 
**MktDataDelay** | Pointer to **int32** | The time it takes, in milliseconds, to process the historical data request | [optional] 
**OutsideRth** | Pointer to **bool** | The historical data returned includes outside of regular trading hours  | [optional] 
**TradingDayDuration** | Pointer to **int32** | The number of seconds in the trading day | [optional] 
**VolumeFactor** | Pointer to **int32** |  | [optional] 
**PriceDisplayRule** | Pointer to **int32** |  | [optional] 
**PriceDisplayValue** | Pointer to **string** |  | [optional] 
**NegativeCapable** | Pointer to **bool** |  | [optional] 
**MessageVersion** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]HistoryDataDataInner**](HistoryDataDataInner.md) |  | [optional] 
**Points** | Pointer to **int32** | total number of points | [optional] 
**TravelTime** | Pointer to **int32** |  | [optional] 

## Methods

### NewHistoryData

`func NewHistoryData() *HistoryData`

NewHistoryData instantiates a new HistoryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryDataWithDefaults

`func NewHistoryDataWithDefaults() *HistoryData`

NewHistoryDataWithDefaults instantiates a new HistoryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *HistoryData) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *HistoryData) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *HistoryData) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *HistoryData) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetText

`func (o *HistoryData) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *HistoryData) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *HistoryData) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *HistoryData) HasText() bool`

HasText returns a boolean if a field has been set.

### GetPriceFactor

`func (o *HistoryData) GetPriceFactor() int32`

GetPriceFactor returns the PriceFactor field if non-nil, zero value otherwise.

### GetPriceFactorOk

`func (o *HistoryData) GetPriceFactorOk() (*int32, bool)`

GetPriceFactorOk returns a tuple with the PriceFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFactor

`func (o *HistoryData) SetPriceFactor(v int32)`

SetPriceFactor sets PriceFactor field to given value.

### HasPriceFactor

`func (o *HistoryData) HasPriceFactor() bool`

HasPriceFactor returns a boolean if a field has been set.

### GetStartTime

`func (o *HistoryData) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HistoryData) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HistoryData) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HistoryData) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetHigh

`func (o *HistoryData) GetHigh() string`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *HistoryData) GetHighOk() (*string, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *HistoryData) SetHigh(v string)`

SetHigh sets High field to given value.

### HasHigh

`func (o *HistoryData) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetLow

`func (o *HistoryData) GetLow() string`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *HistoryData) GetLowOk() (*string, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *HistoryData) SetLow(v string)`

SetLow sets Low field to given value.

### HasLow

`func (o *HistoryData) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetTimePeriod

`func (o *HistoryData) GetTimePeriod() string`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *HistoryData) GetTimePeriodOk() (*string, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *HistoryData) SetTimePeriod(v string)`

SetTimePeriod sets TimePeriod field to given value.

### HasTimePeriod

`func (o *HistoryData) HasTimePeriod() bool`

HasTimePeriod returns a boolean if a field has been set.

### GetBarLength

`func (o *HistoryData) GetBarLength() int32`

GetBarLength returns the BarLength field if non-nil, zero value otherwise.

### GetBarLengthOk

`func (o *HistoryData) GetBarLengthOk() (*int32, bool)`

GetBarLengthOk returns a tuple with the BarLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBarLength

`func (o *HistoryData) SetBarLength(v int32)`

SetBarLength sets BarLength field to given value.

### HasBarLength

`func (o *HistoryData) HasBarLength() bool`

HasBarLength returns a boolean if a field has been set.

### GetMdAvailability

`func (o *HistoryData) GetMdAvailability() string`

GetMdAvailability returns the MdAvailability field if non-nil, zero value otherwise.

### GetMdAvailabilityOk

`func (o *HistoryData) GetMdAvailabilityOk() (*string, bool)`

GetMdAvailabilityOk returns a tuple with the MdAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMdAvailability

`func (o *HistoryData) SetMdAvailability(v string)`

SetMdAvailability sets MdAvailability field to given value.

### HasMdAvailability

`func (o *HistoryData) HasMdAvailability() bool`

HasMdAvailability returns a boolean if a field has been set.

### GetMktDataDelay

`func (o *HistoryData) GetMktDataDelay() int32`

GetMktDataDelay returns the MktDataDelay field if non-nil, zero value otherwise.

### GetMktDataDelayOk

`func (o *HistoryData) GetMktDataDelayOk() (*int32, bool)`

GetMktDataDelayOk returns a tuple with the MktDataDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktDataDelay

`func (o *HistoryData) SetMktDataDelay(v int32)`

SetMktDataDelay sets MktDataDelay field to given value.

### HasMktDataDelay

`func (o *HistoryData) HasMktDataDelay() bool`

HasMktDataDelay returns a boolean if a field has been set.

### GetOutsideRth

`func (o *HistoryData) GetOutsideRth() bool`

GetOutsideRth returns the OutsideRth field if non-nil, zero value otherwise.

### GetOutsideRthOk

`func (o *HistoryData) GetOutsideRthOk() (*bool, bool)`

GetOutsideRthOk returns a tuple with the OutsideRth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRth

`func (o *HistoryData) SetOutsideRth(v bool)`

SetOutsideRth sets OutsideRth field to given value.

### HasOutsideRth

`func (o *HistoryData) HasOutsideRth() bool`

HasOutsideRth returns a boolean if a field has been set.

### GetTradingDayDuration

`func (o *HistoryData) GetTradingDayDuration() int32`

GetTradingDayDuration returns the TradingDayDuration field if non-nil, zero value otherwise.

### GetTradingDayDurationOk

`func (o *HistoryData) GetTradingDayDurationOk() (*int32, bool)`

GetTradingDayDurationOk returns a tuple with the TradingDayDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingDayDuration

`func (o *HistoryData) SetTradingDayDuration(v int32)`

SetTradingDayDuration sets TradingDayDuration field to given value.

### HasTradingDayDuration

`func (o *HistoryData) HasTradingDayDuration() bool`

HasTradingDayDuration returns a boolean if a field has been set.

### GetVolumeFactor

`func (o *HistoryData) GetVolumeFactor() int32`

GetVolumeFactor returns the VolumeFactor field if non-nil, zero value otherwise.

### GetVolumeFactorOk

`func (o *HistoryData) GetVolumeFactorOk() (*int32, bool)`

GetVolumeFactorOk returns a tuple with the VolumeFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeFactor

`func (o *HistoryData) SetVolumeFactor(v int32)`

SetVolumeFactor sets VolumeFactor field to given value.

### HasVolumeFactor

`func (o *HistoryData) HasVolumeFactor() bool`

HasVolumeFactor returns a boolean if a field has been set.

### GetPriceDisplayRule

`func (o *HistoryData) GetPriceDisplayRule() int32`

GetPriceDisplayRule returns the PriceDisplayRule field if non-nil, zero value otherwise.

### GetPriceDisplayRuleOk

`func (o *HistoryData) GetPriceDisplayRuleOk() (*int32, bool)`

GetPriceDisplayRuleOk returns a tuple with the PriceDisplayRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplayRule

`func (o *HistoryData) SetPriceDisplayRule(v int32)`

SetPriceDisplayRule sets PriceDisplayRule field to given value.

### HasPriceDisplayRule

`func (o *HistoryData) HasPriceDisplayRule() bool`

HasPriceDisplayRule returns a boolean if a field has been set.

### GetPriceDisplayValue

`func (o *HistoryData) GetPriceDisplayValue() string`

GetPriceDisplayValue returns the PriceDisplayValue field if non-nil, zero value otherwise.

### GetPriceDisplayValueOk

`func (o *HistoryData) GetPriceDisplayValueOk() (*string, bool)`

GetPriceDisplayValueOk returns a tuple with the PriceDisplayValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDisplayValue

`func (o *HistoryData) SetPriceDisplayValue(v string)`

SetPriceDisplayValue sets PriceDisplayValue field to given value.

### HasPriceDisplayValue

`func (o *HistoryData) HasPriceDisplayValue() bool`

HasPriceDisplayValue returns a boolean if a field has been set.

### GetNegativeCapable

`func (o *HistoryData) GetNegativeCapable() bool`

GetNegativeCapable returns the NegativeCapable field if non-nil, zero value otherwise.

### GetNegativeCapableOk

`func (o *HistoryData) GetNegativeCapableOk() (*bool, bool)`

GetNegativeCapableOk returns a tuple with the NegativeCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeCapable

`func (o *HistoryData) SetNegativeCapable(v bool)`

SetNegativeCapable sets NegativeCapable field to given value.

### HasNegativeCapable

`func (o *HistoryData) HasNegativeCapable() bool`

HasNegativeCapable returns a boolean if a field has been set.

### GetMessageVersion

`func (o *HistoryData) GetMessageVersion() int32`

GetMessageVersion returns the MessageVersion field if non-nil, zero value otherwise.

### GetMessageVersionOk

`func (o *HistoryData) GetMessageVersionOk() (*int32, bool)`

GetMessageVersionOk returns a tuple with the MessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVersion

`func (o *HistoryData) SetMessageVersion(v int32)`

SetMessageVersion sets MessageVersion field to given value.

### HasMessageVersion

`func (o *HistoryData) HasMessageVersion() bool`

HasMessageVersion returns a boolean if a field has been set.

### GetData

`func (o *HistoryData) GetData() []HistoryDataDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryData) GetDataOk() (*[]HistoryDataDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoryData) SetData(v []HistoryDataDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *HistoryData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPoints

`func (o *HistoryData) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *HistoryData) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *HistoryData) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *HistoryData) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetTravelTime

`func (o *HistoryData) GetTravelTime() int32`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *HistoryData) GetTravelTimeOk() (*int32, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *HistoryData) SetTravelTime(v int32)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *HistoryData) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


