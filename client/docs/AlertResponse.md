# AlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | account id | [optional] 
**OrderId** | Pointer to **int32** |  | [optional] 
**AlertName** | Pointer to **string** | name of alert | [optional] 
**AlertMessage** | Pointer to **string** | The message you want to receive via email or text message | [optional] 
**AlertActive** | Pointer to **int32** | whether alert is active or not, so value can only be 0 or 1 | [optional] 
**AlertRepeatable** | Pointer to **int32** | whether alert is repeatable or not, so value can only be 0 or 1 | [optional] 
**AlertEmail** | Pointer to **string** | email address to receive alert | [optional] 
**AlertSendMessage** | Pointer to **int32** | whether allowing to send email or not, so value can only be 0 or 1,  | [optional] 
**Tif** | Pointer to **string** | time in force, can only be GTC or GTD | [optional] 
**ExpireTime** | Pointer to **string** | format, YYYYMMDD-HH:mm:ss  | [optional] 
**OrderStatus** | Pointer to **string** | status of alert | [optional] 
**OutsideRth** | Pointer to **int32** | value can only be 0 or 1, set to 1 if the alert can be triggered outside regular trading hours.  | [optional] 
**ItwsOrdersOnly** | Pointer to **int32** | value can only be 0 or 1, set to 1 to enable the alert only in IBKR mobile  | [optional] 
**AlertShowPopup** | Pointer to **int32** | value can only be 0 or 1, set to 1 to allow to show alert in pop-ups | [optional] 
**AlertTriggered** | Pointer to **bool** | whether the alert has been triggered | [optional] 
**OrderNotEditable** | Pointer to **bool** | whether the alert can be edited | [optional] 
**ToolId** | Pointer to **int32** | for MTA alert only, each user has a unique toolId and it will stay the same, do not send for normal alert  | [optional] 
**AlertPlayAudio** | Pointer to **string** | audio message to play when alert is triggered | [optional] 
**AlertMtaCurrency** | Pointer to **string** | MTA alert only | [optional] 
**AlertMtaDefaults** | Pointer to **string** | MTA alert only | [optional] 
**TimeZone** | Pointer to **string** | MTA alert only | [optional] 
**AlertDefaultType** | Pointer to **string** | MTA alert only | [optional] 
**ConditionSize** | Pointer to **int32** | size of conditions array | [optional] 
**ConditionOutsideRth** | Pointer to **int32** | whether allowing the condition can be triggered outside of regular trading hours, 1 means allow | [optional] 
**Conditions** | Pointer to [**[]AlertResponseConditionsInner**](AlertResponseConditionsInner.md) |  | [optional] 

## Methods

### NewAlertResponse

`func NewAlertResponse() *AlertResponse`

NewAlertResponse instantiates a new AlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResponseWithDefaults

`func NewAlertResponseWithDefaults() *AlertResponse`

NewAlertResponseWithDefaults instantiates a new AlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AlertResponse) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AlertResponse) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AlertResponse) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AlertResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOrderId

`func (o *AlertResponse) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AlertResponse) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AlertResponse) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AlertResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetAlertName

`func (o *AlertResponse) GetAlertName() string`

GetAlertName returns the AlertName field if non-nil, zero value otherwise.

### GetAlertNameOk

`func (o *AlertResponse) GetAlertNameOk() (*string, bool)`

GetAlertNameOk returns a tuple with the AlertName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertName

`func (o *AlertResponse) SetAlertName(v string)`

SetAlertName sets AlertName field to given value.

### HasAlertName

`func (o *AlertResponse) HasAlertName() bool`

HasAlertName returns a boolean if a field has been set.

### GetAlertMessage

`func (o *AlertResponse) GetAlertMessage() string`

GetAlertMessage returns the AlertMessage field if non-nil, zero value otherwise.

### GetAlertMessageOk

`func (o *AlertResponse) GetAlertMessageOk() (*string, bool)`

GetAlertMessageOk returns a tuple with the AlertMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertMessage

`func (o *AlertResponse) SetAlertMessage(v string)`

SetAlertMessage sets AlertMessage field to given value.

### HasAlertMessage

`func (o *AlertResponse) HasAlertMessage() bool`

HasAlertMessage returns a boolean if a field has been set.

### GetAlertActive

`func (o *AlertResponse) GetAlertActive() int32`

GetAlertActive returns the AlertActive field if non-nil, zero value otherwise.

### GetAlertActiveOk

`func (o *AlertResponse) GetAlertActiveOk() (*int32, bool)`

GetAlertActiveOk returns a tuple with the AlertActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertActive

`func (o *AlertResponse) SetAlertActive(v int32)`

SetAlertActive sets AlertActive field to given value.

### HasAlertActive

`func (o *AlertResponse) HasAlertActive() bool`

HasAlertActive returns a boolean if a field has been set.

### GetAlertRepeatable

`func (o *AlertResponse) GetAlertRepeatable() int32`

GetAlertRepeatable returns the AlertRepeatable field if non-nil, zero value otherwise.

### GetAlertRepeatableOk

`func (o *AlertResponse) GetAlertRepeatableOk() (*int32, bool)`

GetAlertRepeatableOk returns a tuple with the AlertRepeatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRepeatable

`func (o *AlertResponse) SetAlertRepeatable(v int32)`

SetAlertRepeatable sets AlertRepeatable field to given value.

### HasAlertRepeatable

`func (o *AlertResponse) HasAlertRepeatable() bool`

HasAlertRepeatable returns a boolean if a field has been set.

### GetAlertEmail

`func (o *AlertResponse) GetAlertEmail() string`

GetAlertEmail returns the AlertEmail field if non-nil, zero value otherwise.

### GetAlertEmailOk

`func (o *AlertResponse) GetAlertEmailOk() (*string, bool)`

GetAlertEmailOk returns a tuple with the AlertEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertEmail

`func (o *AlertResponse) SetAlertEmail(v string)`

SetAlertEmail sets AlertEmail field to given value.

### HasAlertEmail

`func (o *AlertResponse) HasAlertEmail() bool`

HasAlertEmail returns a boolean if a field has been set.

### GetAlertSendMessage

`func (o *AlertResponse) GetAlertSendMessage() int32`

GetAlertSendMessage returns the AlertSendMessage field if non-nil, zero value otherwise.

### GetAlertSendMessageOk

`func (o *AlertResponse) GetAlertSendMessageOk() (*int32, bool)`

GetAlertSendMessageOk returns a tuple with the AlertSendMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSendMessage

`func (o *AlertResponse) SetAlertSendMessage(v int32)`

SetAlertSendMessage sets AlertSendMessage field to given value.

### HasAlertSendMessage

`func (o *AlertResponse) HasAlertSendMessage() bool`

HasAlertSendMessage returns a boolean if a field has been set.

### GetTif

`func (o *AlertResponse) GetTif() string`

GetTif returns the Tif field if non-nil, zero value otherwise.

### GetTifOk

`func (o *AlertResponse) GetTifOk() (*string, bool)`

GetTifOk returns a tuple with the Tif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTif

`func (o *AlertResponse) SetTif(v string)`

SetTif sets Tif field to given value.

### HasTif

`func (o *AlertResponse) HasTif() bool`

HasTif returns a boolean if a field has been set.

### GetExpireTime

`func (o *AlertResponse) GetExpireTime() string`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *AlertResponse) GetExpireTimeOk() (*string, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *AlertResponse) SetExpireTime(v string)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *AlertResponse) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetOrderStatus

`func (o *AlertResponse) GetOrderStatus() string`

GetOrderStatus returns the OrderStatus field if non-nil, zero value otherwise.

### GetOrderStatusOk

`func (o *AlertResponse) GetOrderStatusOk() (*string, bool)`

GetOrderStatusOk returns a tuple with the OrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStatus

`func (o *AlertResponse) SetOrderStatus(v string)`

SetOrderStatus sets OrderStatus field to given value.

### HasOrderStatus

`func (o *AlertResponse) HasOrderStatus() bool`

HasOrderStatus returns a boolean if a field has been set.

### GetOutsideRth

`func (o *AlertResponse) GetOutsideRth() int32`

GetOutsideRth returns the OutsideRth field if non-nil, zero value otherwise.

### GetOutsideRthOk

`func (o *AlertResponse) GetOutsideRthOk() (*int32, bool)`

GetOutsideRthOk returns a tuple with the OutsideRth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutsideRth

`func (o *AlertResponse) SetOutsideRth(v int32)`

SetOutsideRth sets OutsideRth field to given value.

### HasOutsideRth

`func (o *AlertResponse) HasOutsideRth() bool`

HasOutsideRth returns a boolean if a field has been set.

### GetItwsOrdersOnly

`func (o *AlertResponse) GetItwsOrdersOnly() int32`

GetItwsOrdersOnly returns the ItwsOrdersOnly field if non-nil, zero value otherwise.

### GetItwsOrdersOnlyOk

`func (o *AlertResponse) GetItwsOrdersOnlyOk() (*int32, bool)`

GetItwsOrdersOnlyOk returns a tuple with the ItwsOrdersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItwsOrdersOnly

`func (o *AlertResponse) SetItwsOrdersOnly(v int32)`

SetItwsOrdersOnly sets ItwsOrdersOnly field to given value.

### HasItwsOrdersOnly

`func (o *AlertResponse) HasItwsOrdersOnly() bool`

HasItwsOrdersOnly returns a boolean if a field has been set.

### GetAlertShowPopup

`func (o *AlertResponse) GetAlertShowPopup() int32`

GetAlertShowPopup returns the AlertShowPopup field if non-nil, zero value otherwise.

### GetAlertShowPopupOk

`func (o *AlertResponse) GetAlertShowPopupOk() (*int32, bool)`

GetAlertShowPopupOk returns a tuple with the AlertShowPopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertShowPopup

`func (o *AlertResponse) SetAlertShowPopup(v int32)`

SetAlertShowPopup sets AlertShowPopup field to given value.

### HasAlertShowPopup

`func (o *AlertResponse) HasAlertShowPopup() bool`

HasAlertShowPopup returns a boolean if a field has been set.

### GetAlertTriggered

`func (o *AlertResponse) GetAlertTriggered() bool`

GetAlertTriggered returns the AlertTriggered field if non-nil, zero value otherwise.

### GetAlertTriggeredOk

`func (o *AlertResponse) GetAlertTriggeredOk() (*bool, bool)`

GetAlertTriggeredOk returns a tuple with the AlertTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertTriggered

`func (o *AlertResponse) SetAlertTriggered(v bool)`

SetAlertTriggered sets AlertTriggered field to given value.

### HasAlertTriggered

`func (o *AlertResponse) HasAlertTriggered() bool`

HasAlertTriggered returns a boolean if a field has been set.

### GetOrderNotEditable

`func (o *AlertResponse) GetOrderNotEditable() bool`

GetOrderNotEditable returns the OrderNotEditable field if non-nil, zero value otherwise.

### GetOrderNotEditableOk

`func (o *AlertResponse) GetOrderNotEditableOk() (*bool, bool)`

GetOrderNotEditableOk returns a tuple with the OrderNotEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNotEditable

`func (o *AlertResponse) SetOrderNotEditable(v bool)`

SetOrderNotEditable sets OrderNotEditable field to given value.

### HasOrderNotEditable

`func (o *AlertResponse) HasOrderNotEditable() bool`

HasOrderNotEditable returns a boolean if a field has been set.

### GetToolId

`func (o *AlertResponse) GetToolId() int32`

GetToolId returns the ToolId field if non-nil, zero value otherwise.

### GetToolIdOk

`func (o *AlertResponse) GetToolIdOk() (*int32, bool)`

GetToolIdOk returns a tuple with the ToolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolId

`func (o *AlertResponse) SetToolId(v int32)`

SetToolId sets ToolId field to given value.

### HasToolId

`func (o *AlertResponse) HasToolId() bool`

HasToolId returns a boolean if a field has been set.

### GetAlertPlayAudio

`func (o *AlertResponse) GetAlertPlayAudio() string`

GetAlertPlayAudio returns the AlertPlayAudio field if non-nil, zero value otherwise.

### GetAlertPlayAudioOk

`func (o *AlertResponse) GetAlertPlayAudioOk() (*string, bool)`

GetAlertPlayAudioOk returns a tuple with the AlertPlayAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertPlayAudio

`func (o *AlertResponse) SetAlertPlayAudio(v string)`

SetAlertPlayAudio sets AlertPlayAudio field to given value.

### HasAlertPlayAudio

`func (o *AlertResponse) HasAlertPlayAudio() bool`

HasAlertPlayAudio returns a boolean if a field has been set.

### GetAlertMtaCurrency

`func (o *AlertResponse) GetAlertMtaCurrency() string`

GetAlertMtaCurrency returns the AlertMtaCurrency field if non-nil, zero value otherwise.

### GetAlertMtaCurrencyOk

`func (o *AlertResponse) GetAlertMtaCurrencyOk() (*string, bool)`

GetAlertMtaCurrencyOk returns a tuple with the AlertMtaCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertMtaCurrency

`func (o *AlertResponse) SetAlertMtaCurrency(v string)`

SetAlertMtaCurrency sets AlertMtaCurrency field to given value.

### HasAlertMtaCurrency

`func (o *AlertResponse) HasAlertMtaCurrency() bool`

HasAlertMtaCurrency returns a boolean if a field has been set.

### GetAlertMtaDefaults

`func (o *AlertResponse) GetAlertMtaDefaults() string`

GetAlertMtaDefaults returns the AlertMtaDefaults field if non-nil, zero value otherwise.

### GetAlertMtaDefaultsOk

`func (o *AlertResponse) GetAlertMtaDefaultsOk() (*string, bool)`

GetAlertMtaDefaultsOk returns a tuple with the AlertMtaDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertMtaDefaults

`func (o *AlertResponse) SetAlertMtaDefaults(v string)`

SetAlertMtaDefaults sets AlertMtaDefaults field to given value.

### HasAlertMtaDefaults

`func (o *AlertResponse) HasAlertMtaDefaults() bool`

HasAlertMtaDefaults returns a boolean if a field has been set.

### GetTimeZone

`func (o *AlertResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AlertResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AlertResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AlertResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetAlertDefaultType

`func (o *AlertResponse) GetAlertDefaultType() string`

GetAlertDefaultType returns the AlertDefaultType field if non-nil, zero value otherwise.

### GetAlertDefaultTypeOk

`func (o *AlertResponse) GetAlertDefaultTypeOk() (*string, bool)`

GetAlertDefaultTypeOk returns a tuple with the AlertDefaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertDefaultType

`func (o *AlertResponse) SetAlertDefaultType(v string)`

SetAlertDefaultType sets AlertDefaultType field to given value.

### HasAlertDefaultType

`func (o *AlertResponse) HasAlertDefaultType() bool`

HasAlertDefaultType returns a boolean if a field has been set.

### GetConditionSize

`func (o *AlertResponse) GetConditionSize() int32`

GetConditionSize returns the ConditionSize field if non-nil, zero value otherwise.

### GetConditionSizeOk

`func (o *AlertResponse) GetConditionSizeOk() (*int32, bool)`

GetConditionSizeOk returns a tuple with the ConditionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionSize

`func (o *AlertResponse) SetConditionSize(v int32)`

SetConditionSize sets ConditionSize field to given value.

### HasConditionSize

`func (o *AlertResponse) HasConditionSize() bool`

HasConditionSize returns a boolean if a field has been set.

### GetConditionOutsideRth

`func (o *AlertResponse) GetConditionOutsideRth() int32`

GetConditionOutsideRth returns the ConditionOutsideRth field if non-nil, zero value otherwise.

### GetConditionOutsideRthOk

`func (o *AlertResponse) GetConditionOutsideRthOk() (*int32, bool)`

GetConditionOutsideRthOk returns a tuple with the ConditionOutsideRth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionOutsideRth

`func (o *AlertResponse) SetConditionOutsideRth(v int32)`

SetConditionOutsideRth sets ConditionOutsideRth field to given value.

### HasConditionOutsideRth

`func (o *AlertResponse) HasConditionOutsideRth() bool`

HasConditionOutsideRth returns a boolean if a field has been set.

### GetConditions

`func (o *AlertResponse) GetConditions() []AlertResponseConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AlertResponse) GetConditionsOk() (*[]AlertResponseConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AlertResponse) SetConditions(v []AlertResponseConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AlertResponse) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


