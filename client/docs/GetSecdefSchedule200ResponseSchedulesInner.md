# GetSecdefSchedule200ResponseSchedulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClearingCycleEndTime** | Pointer to **int32** |  | [optional] 
**TradingScheduleDate** | Pointer to **int32** | 20000101 stands for any Sat, 20000102 stands for any Sun, ... 20000107 stands for any Fri. Any other date stands for itself. | [optional] 
**Sessions** | Pointer to [**GetSecdefSchedule200ResponseSchedulesInnerSessions**](GetSecdefSchedule200ResponseSchedulesInnerSessions.md) |  | [optional] 
**TradingTimes** | Pointer to [**GetSecdefSchedule200ResponseSchedulesInnerTradingTimes**](GetSecdefSchedule200ResponseSchedulesInnerTradingTimes.md) |  | [optional] 

## Methods

### NewGetSecdefSchedule200ResponseSchedulesInner

`func NewGetSecdefSchedule200ResponseSchedulesInner() *GetSecdefSchedule200ResponseSchedulesInner`

NewGetSecdefSchedule200ResponseSchedulesInner instantiates a new GetSecdefSchedule200ResponseSchedulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecdefSchedule200ResponseSchedulesInnerWithDefaults

`func NewGetSecdefSchedule200ResponseSchedulesInnerWithDefaults() *GetSecdefSchedule200ResponseSchedulesInner`

NewGetSecdefSchedule200ResponseSchedulesInnerWithDefaults instantiates a new GetSecdefSchedule200ResponseSchedulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClearingCycleEndTime

`func (o *GetSecdefSchedule200ResponseSchedulesInner) GetClearingCycleEndTime() int32`

GetClearingCycleEndTime returns the ClearingCycleEndTime field if non-nil, zero value otherwise.

### GetClearingCycleEndTimeOk

`func (o *GetSecdefSchedule200ResponseSchedulesInner) GetClearingCycleEndTimeOk() (*int32, bool)`

GetClearingCycleEndTimeOk returns a tuple with the ClearingCycleEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingCycleEndTime

`func (o *GetSecdefSchedule200ResponseSchedulesInner) SetClearingCycleEndTime(v int32)`

SetClearingCycleEndTime sets ClearingCycleEndTime field to given value.

### HasClearingCycleEndTime

`func (o *GetSecdefSchedule200ResponseSchedulesInner) HasClearingCycleEndTime() bool`

HasClearingCycleEndTime returns a boolean if a field has been set.

### GetTradingScheduleDate

`func (o *GetSecdefSchedule200ResponseSchedulesInner) GetTradingScheduleDate() int32`

GetTradingScheduleDate returns the TradingScheduleDate field if non-nil, zero value otherwise.

### GetTradingScheduleDateOk

`func (o *GetSecdefSchedule200ResponseSchedulesInner) GetTradingScheduleDateOk() (*int32, bool)`

GetTradingScheduleDateOk returns a tuple with the TradingScheduleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingScheduleDate

`func (o *GetSecdefSchedule200ResponseSchedulesInner) SetTradingScheduleDate(v int32)`

SetTradingScheduleDate sets TradingScheduleDate field to given value.

### HasTradingScheduleDate

`func (o *GetSecdefSchedule200ResponseSchedulesInner) HasTradingScheduleDate() bool`

HasTradingScheduleDate returns a boolean if a field has been set.

### GetSessions

`func (o *GetSecdefSchedule200ResponseSchedulesInner) GetSessions() GetSecdefSchedule200ResponseSchedulesInnerSessions`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *GetSecdefSchedule200ResponseSchedulesInner) GetSessionsOk() (*GetSecdefSchedule200ResponseSchedulesInnerSessions, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *GetSecdefSchedule200ResponseSchedulesInner) SetSessions(v GetSecdefSchedule200ResponseSchedulesInnerSessions)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *GetSecdefSchedule200ResponseSchedulesInner) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetTradingTimes

`func (o *GetSecdefSchedule200ResponseSchedulesInner) GetTradingTimes() GetSecdefSchedule200ResponseSchedulesInnerTradingTimes`

GetTradingTimes returns the TradingTimes field if non-nil, zero value otherwise.

### GetTradingTimesOk

`func (o *GetSecdefSchedule200ResponseSchedulesInner) GetTradingTimesOk() (*GetSecdefSchedule200ResponseSchedulesInnerTradingTimes, bool)`

GetTradingTimesOk returns a tuple with the TradingTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingTimes

`func (o *GetSecdefSchedule200ResponseSchedulesInner) SetTradingTimes(v GetSecdefSchedule200ResponseSchedulesInnerTradingTimes)`

SetTradingTimes sets TradingTimes field to given value.

### HasTradingTimes

`func (o *GetSecdefSchedule200ResponseSchedulesInner) HasTradingTimes() bool`

HasTradingTimes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


