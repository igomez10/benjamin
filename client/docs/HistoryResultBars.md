# HistoryResultBars

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Open** | Pointer to **float32** | First price returned for bar value. | [optional] 
**StartTime** | Pointer to **string** | Start Time in the format YYYYMMDD. | [optional] 
**StartTimeVal** | Pointer to **int32** | Start Time Value - Formatted in unix time in ms. | [optional] 
**EndTime** | Pointer to **string** | End Time in the format YYYYMMDD. | [optional] 
**EndTimeVal** | Pointer to **int32** | End Time Value - Formatted in unix time in ms. | [optional] 
**Points** | Pointer to **int32** | total number of data points. | [optional] 
**Data** | Pointer to [**[]HistoryResultBarsDataInner**](HistoryResultBarsDataInner.md) |  | [optional] 
**MktDataDelay** | Pointer to **int32** | If 0 then data is returned in real time. Otherwise will return the number of seconds history data is delayed. | [optional] 

## Methods

### NewHistoryResultBars

`func NewHistoryResultBars() *HistoryResultBars`

NewHistoryResultBars instantiates a new HistoryResultBars object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryResultBarsWithDefaults

`func NewHistoryResultBarsWithDefaults() *HistoryResultBars`

NewHistoryResultBarsWithDefaults instantiates a new HistoryResultBars object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpen

`func (o *HistoryResultBars) GetOpen() float32`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *HistoryResultBars) GetOpenOk() (*float32, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *HistoryResultBars) SetOpen(v float32)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *HistoryResultBars) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetStartTime

`func (o *HistoryResultBars) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HistoryResultBars) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HistoryResultBars) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HistoryResultBars) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStartTimeVal

`func (o *HistoryResultBars) GetStartTimeVal() int32`

GetStartTimeVal returns the StartTimeVal field if non-nil, zero value otherwise.

### GetStartTimeValOk

`func (o *HistoryResultBars) GetStartTimeValOk() (*int32, bool)`

GetStartTimeValOk returns a tuple with the StartTimeVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeVal

`func (o *HistoryResultBars) SetStartTimeVal(v int32)`

SetStartTimeVal sets StartTimeVal field to given value.

### HasStartTimeVal

`func (o *HistoryResultBars) HasStartTimeVal() bool`

HasStartTimeVal returns a boolean if a field has been set.

### GetEndTime

`func (o *HistoryResultBars) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HistoryResultBars) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HistoryResultBars) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HistoryResultBars) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEndTimeVal

`func (o *HistoryResultBars) GetEndTimeVal() int32`

GetEndTimeVal returns the EndTimeVal field if non-nil, zero value otherwise.

### GetEndTimeValOk

`func (o *HistoryResultBars) GetEndTimeValOk() (*int32, bool)`

GetEndTimeValOk returns a tuple with the EndTimeVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeVal

`func (o *HistoryResultBars) SetEndTimeVal(v int32)`

SetEndTimeVal sets EndTimeVal field to given value.

### HasEndTimeVal

`func (o *HistoryResultBars) HasEndTimeVal() bool`

HasEndTimeVal returns a boolean if a field has been set.

### GetPoints

`func (o *HistoryResultBars) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *HistoryResultBars) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *HistoryResultBars) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *HistoryResultBars) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetData

`func (o *HistoryResultBars) GetData() []HistoryResultBarsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *HistoryResultBars) GetDataOk() (*[]HistoryResultBarsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *HistoryResultBars) SetData(v []HistoryResultBarsDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *HistoryResultBars) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMktDataDelay

`func (o *HistoryResultBars) GetMktDataDelay() int32`

GetMktDataDelay returns the MktDataDelay field if non-nil, zero value otherwise.

### GetMktDataDelayOk

`func (o *HistoryResultBars) GetMktDataDelayOk() (*int32, bool)`

GetMktDataDelayOk returns a tuple with the MktDataDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktDataDelay

`func (o *HistoryResultBars) SetMktDataDelay(v int32)`

SetMktDataDelay sets MktDataDelay field to given value.

### HasMktDataDelay

`func (o *HistoryResultBars) HasMktDataDelay() bool`

HasMktDataDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


