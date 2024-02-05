# AlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int32** | orderId is required when modifying alert. You can get it from /iserver/account/{accountId}/alerts  | [optional] 
**AlertName** | Pointer to **string** | name of alert | [optional] 
**AlertMessage** | Pointer to **string** | The message you want to receive via email or text message | [optional] 
**AlertRepeatable** | Pointer to **int32** | whether alert is repeatable or not, so value can only be 0 or 1, this has to be 1 for MTA alert | [optional] 
**Email** | Pointer to **string** | email address to receive alert | [optional] 
**SendMessage** | Pointer to **int32** | whether allowing to send email or not, so value can only be 0 or 1,  | [optional] 
**Tif** | Pointer to **string** | time in force, can only be GTC or GTD | [optional] 
**ExpireTime** | Pointer to **string** | format, YYYYMMDD-HH:mm:ss, please NOTE this will only work when tif is GTD  | [optional] 
**OutsideRth** | Pointer to **int32** | value can only be 0 or 1, set to 1 if the alert can be triggered outside regular trading hours.  | [optional] 
**ITWSOrdersOnly** | Pointer to **int32** | value can only be 0 or 1, set to 1 to enable the alert only in IBKR mobile  | [optional] 
**ShowPopup** | Pointer to **int32** | value can only be 0 or 1, set to 1 to allow to show alert in pop-ups | [optional] 
**ToolId** | Pointer to **int32** | for MTA alert only, each user has a unique toolId and it will stay the same, do not send for normal alert  | [optional] 
**PlayAudio** | Pointer to **string** | audio message to play when alert is triggered | [optional] 
**Conditions** | Pointer to [**[]AlertRequestConditionsInner**](AlertRequestConditionsInner.md) |  | [optional] 

## Methods

### NewAlertRequest

`func NewAlertRequest() *AlertRequest`

NewAlertRequest instantiates a new AlertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRequestWithDefaults

`func NewAlertRequestWithDefaults() *AlertRequest`

NewAlertRequestWithDefaults instantiates a new AlertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *AlertRequest) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AlertRequest) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AlertRequest) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AlertRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAlertName

`func (o *AlertRequest) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *AlertRequest) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *AlertRequest) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *AlertRequest) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetAlertMessage

`func (o *AlertRequest) GetAlertMessage() string`

GetAlertMessage returns the AlertMessage field if non-nil, zero value otherwise.

### GetAlertMessageOk

`func (o *AlertRequest) GetAlertMessageOk() (*string, bool)`

GetAlertMessageOk returns a tuple with the AlertMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertMessage

`func (o *AlertRequest) SetAlertMessage(v string)`

SetAlertMessage sets AlertMessage field to given value.

### HasAlertMessage

`func (o *AlertRequest) HasAlertMessage() bool`

HasAlertMessage returns a boolean if a field has been set.

### GetAlertRepeatable

`func (o *AlertRequest) GetAlertRepeatable() int32`

GetAlertRepeatable returns the AlertRepeatable field if non-nil, zero value otherwise.

### GetAlertRepeatableOk

`func (o *AlertRequest) GetAlertRepeatableOk() (*int32, bool)`

GetAlertRepeatableOk returns a tuple with the AlertRepeatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRepeatable

`func (o *AlertRequest) SetAlertRepeatable(v int32)`

SetAlertRepeatable sets AlertRepeatable field to given value.

### HasAlertRepeatable

`func (o *AlertRequest) HasAlertRepeatable() bool`

HasAlertRepeatable returns a boolean if a field has been set.

### GetEmail

`func (o *AlertRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AlertRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AlertRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AlertRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSendMessage

`func (o *AlertRequest) GetSendMessage() int32`

GetSendMessage returns the SendMessage field if non-nil, zero value otherwise.

### GetSendMessageOk

`func (o *AlertRequest) GetSendMessageOk() (*int32, bool)`

GetSendMessageOk returns a tuple with the SendMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendMessage

`func (o *AlertRequest) SetSendMessage(v int32)`

SetSendMessage sets SendMessage field to given value.

### HasSendMessage

`func (o *AlertRequest) HasSendMessage() bool`

HasSendMessage returns a boolean if a field has been set.

### GetTif

`func (o *AlertRequest) GetTif() string`

GetTif returns the Tif field if non-nil, zero value otherwise.

### GetTifOk

`func (o *AlertRequest) GetTifOk() (*string, bool)`

GetTifOk returns a tuple with the Tif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTif

`func (o *AlertRequest) SetTif(v string)`

SetTif sets Tif field to given value.

### HasTif

`func (o *AlertRequest) HasTif() bool`

HasTif returns a boolean if a field has been set.

### GetExpireTime

`func (o *AlertRequest) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *AlertRequest) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *AlertRequest) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *AlertRequest) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetOutsideRth

`func (o *AlertRequest) GetOutsideRth() int32`

GetOutsideRth returns the OutsideRth field if non-nil, zero value otherwise.

### GetOutsideRthOk

`func (o *AlertRequest) GetOutsideRthOk() (*int32, bool)`

GetOutsideRthOk returns a tuple with the OutsideRth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRth

`func (o *AlertRequest) SetOutsideRth(v int32)`

SetOutsideRth sets OutsideRth field to given value.

### HasOutsideRth

`func (o *AlertRequest) HasOutsideRth() bool`

HasOutsideRth returns a boolean if a field has been set.

### GetITWSOrdersOnly

`func (o *AlertRequest) GetITWSOrdersOnly() int32`

GetITWSOrdersOnly returns the ITWSOrdersOnly field if non-nil, zero value otherwise.

### GetITWSOrdersOnlyOk

`func (o *AlertRequest) GetITWSOrdersOnlyOk() (*int32, bool)`

GetITWSOrdersOnlyOk returns a tuple with the ITWSOrdersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetITWSOrdersOnly

`func (o *AlertRequest) SetITWSOrdersOnly(v int32)`

SetITWSOrdersOnly sets ITWSOrdersOnly field to given value.

### HasITWSOrdersOnly

`func (o *AlertRequest) HasITWSOrdersOnly() bool`

HasITWSOrdersOnly returns a boolean if a field has been set.

### GetShowPopup

`func (o *AlertRequest) GetShowPopup() int32`

GetShowPopup returns the ShowPopup field if non-nil, zero value otherwise.

### GetShowPopupOk

`func (o *AlertRequest) GetShowPopupOk() (*int32, bool)`

GetShowPopupOk returns a tuple with the ShowPopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPopup

`func (o *AlertRequest) SetShowPopup(v int32)`

SetShowPopup sets ShowPopup field to given value.

### HasShowPopup

`func (o *AlertRequest) HasShowPopup() bool`

HasShowPopup returns a boolean if a field has been set.

### GetToolId

`func (o *AlertRequest) GetToolId() int32`

GetToolId returns the ToolId field if non-nil, zero value otherwise.

### GetToolIdOk

`func (o *AlertRequest) GetToolIdOk() (*int32, bool)`

GetToolIdOk returns a tuple with the ToolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolId

`func (o *AlertRequest) SetToolId(v int32)`

SetToolId sets ToolId field to given value.

### HasToolId

`func (o *AlertRequest) HasToolId() bool`

HasToolId returns a boolean if a field has been set.

### GetPlayAudio

`func (o *AlertRequest) GetPlayAudio() string`

GetPlayAudio returns the PlayAudio field if non-nil, zero value otherwise.

### GetPlayAudioOk

`func (o *AlertRequest) GetPlayAudioOk() (*string, bool)`

GetPlayAudioOk returns a tuple with the PlayAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayAudio

`func (o *AlertRequest) SetPlayAudio(v string)`

SetPlayAudio sets PlayAudio field to given value.

### HasPlayAudio

`func (o *AlertRequest) HasPlayAudio() bool`

HasPlayAudio returns a boolean if a field has been set.

### GetConditions

`func (o *AlertRequest) GetConditions() []AlertRequestConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AlertRequest) GetConditionsOk() (*[]AlertRequestConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AlertRequest) SetConditions(v []AlertRequestConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AlertRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


