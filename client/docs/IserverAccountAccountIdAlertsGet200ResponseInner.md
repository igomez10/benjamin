# IserverAccountAccountIdAlertsGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to **string** | account id | [optional] 
**AlertName** | Pointer to **string** |  | [optional] 
**AlertActive** | Pointer to **int32** | Value can only be 0 or 1, 1 means active | [optional] 
**OrderTime** | Pointer to **string** | format, YYYYMMDD-HH:mm:ss, the time when you created the alert  | [optional] 
**AlertTriggered** | Pointer to **bool** | whether the alert has been triggered or not | [optional] 
**AlertRepeatable** | Pointer to **int32** | whether the alert can be repeatable or not, value can be 1 or 0. 1 means true | [optional] 

## Methods

### NewIserverAccountAccountIdAlertsGet200ResponseInner

`func NewIserverAccountAccountIdAlertsGet200ResponseInner() *IserverAccountAccountIdAlertsGet200ResponseInner`

NewIserverAccountAccountIdAlertsGet200ResponseInner instantiates a new IserverAccountAccountIdAlertsGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIserverAccountAccountIdAlertsGet200ResponseInnerWithDefaults

`func NewIserverAccountAccountIdAlertsGet200ResponseInnerWithDefaults() *IserverAccountAccountIdAlertsGet200ResponseInner`

NewIserverAccountAccountIdAlertsGet200ResponseInnerWithDefaults instantiates a new IserverAccountAccountIdAlertsGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAccount

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAlertName

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetAlertActive

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAlertActive() int32`

GetAlertActive returns the AlertActive field if non-nil, zero value otherwise.

### GetAlertActiveOk

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAlertActiveOk() (*int32, bool)`

GetAlertActiveOk returns a tuple with the AlertActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertActive

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) SetAlertActive(v int32)`

SetAlertActive sets AlertActive field to given value.

### HasAlertActive

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) HasAlertActive() bool`

HasAlertActive returns a boolean if a field has been set.

### GetOrderTime

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetOrderTime() string`

GetOrderTime returns the OrderTime field if non-nil, zero value otherwise.

### GetOrderTimeOk

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetOrderTimeOk() (*string, bool)`

GetOrderTimeOk returns a tuple with the OrderTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderTime

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) SetOrderTime(v string)`

SetOrderTime sets OrderTime field to given value.

### HasOrderTime

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) HasOrderTime() bool`

HasOrderTime returns a boolean if a field has been set.

### GetAlertTriggered

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAlertTriggered() bool`

GetAlertTriggered returns the AlertTriggered field if non-nil, zero value otherwise.

### GetAlertTriggeredOk

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAlertTriggeredOk() (*bool, bool)`

GetAlertTriggeredOk returns a tuple with the AlertTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTriggered

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) SetAlertTriggered(v bool)`

SetAlertTriggered sets AlertTriggered field to given value.

### HasAlertTriggered

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) HasAlertTriggered() bool`

HasAlertTriggered returns a boolean if a field has been set.

### GetAlertRepeatable

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAlertRepeatable() int32`

GetAlertRepeatable returns the AlertRepeatable field if non-nil, zero value otherwise.

### GetAlertRepeatableOk

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) GetAlertRepeatableOk() (*int32, bool)`

GetAlertRepeatableOk returns a tuple with the AlertRepeatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRepeatable

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) SetAlertRepeatable(v int32)`

SetAlertRepeatable sets AlertRepeatable field to given value.

### HasAlertRepeatable

`func (o *IserverAccountAccountIdAlertsGet200ResponseInner) HasAlertRepeatable() bool`

HasAlertRepeatable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


