# AlertResponseConditionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | Pointer to **int32** | Types: 1-Price, 3-Time, 4-Margin, 5-Trade, 6-Volume, 7: MTA market 8: MTA Position, 9: MTA Acc. Daily PN&amp;  | [optional] 
**Conidex** | Pointer to **string** | conid and exchange. Format supports conid or conid@exchange | [optional] 
**ContractDescription1** | Pointer to **string** | Format contract name | [optional] 
**ConditionOperator** | Pointer to **string** | optional, operator for the current condition   * &gt;&#x3D; Greater than or equal to   * &lt;&#x3D; Less than or equal to  | [optional] 
**ConditionTriggerMethod** | Pointer to **string** | optional, only some type of conditions have triggerMethod | [optional] 
**ConditionValue** | Pointer to **string** | can not be empty, can pass default value \&quot;*\&quot; | [optional] 
**ConditionLogicBind** | Pointer to **string** | Condition array should end with \&quot;n\&quot;   * a - AND   * o - OR   * n - END  | [optional] 
**ConditionTimeZone** | Pointer to **string** | only needed for some MTA alert condition | [optional] 

## Methods

### NewAlertResponseConditionsInner

`func NewAlertResponseConditionsInner() *AlertResponseConditionsInner`

NewAlertResponseConditionsInner instantiates a new AlertResponseConditionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertResponseConditionsInnerWithDefaults

`func NewAlertResponseConditionsInnerWithDefaults() *AlertResponseConditionsInner`

NewAlertResponseConditionsInnerWithDefaults instantiates a new AlertResponseConditionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *AlertResponseConditionsInner) GetConditionType() int32`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *AlertResponseConditionsInner) GetConditionTypeOk() (*int32, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *AlertResponseConditionsInner) SetConditionType(v int32)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *AlertResponseConditionsInner) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### GetConidex

`func (o *AlertResponseConditionsInner) GetConidex() string`

GetConidex returns the Conidex field if non-nil, zero value otherwise.

### GetConidexOk

`func (o *AlertResponseConditionsInner) GetConidexOk() (*string, bool)`

GetConidexOk returns a tuple with the Conidex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConidex

`func (o *AlertResponseConditionsInner) SetConidex(v string)`

SetConidex sets Conidex field to given value.

### HasConidex

`func (o *AlertResponseConditionsInner) HasConidex() bool`

HasConidex returns a boolean if a field has been set.

### GetContractDescription1

`func (o *AlertResponseConditionsInner) GetContractDescription1() string`

GetContractDescription1 returns the ContractDescription1 field if non-nil, zero value otherwise.

### GetContractDescription1Ok

`func (o *AlertResponseConditionsInner) GetContractDescription1Ok() (*string, bool)`

GetContractDescription1Ok returns a tuple with the ContractDescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDescription1

`func (o *AlertResponseConditionsInner) SetContractDescription1(v string)`

SetContractDescription1 sets ContractDescription1 field to given value.

### HasContractDescription1

`func (o *AlertResponseConditionsInner) HasContractDescription1() bool`

HasContractDescription1 returns a boolean if a field has been set.

### GetConditionOperator

`func (o *AlertResponseConditionsInner) GetConditionOperator() string`

GetConditionOperator returns the ConditionOperator field if non-nil, zero value otherwise.

### GetConditionOperatorOk

`func (o *AlertResponseConditionsInner) GetConditionOperatorOk() (*string, bool)`

GetConditionOperatorOk returns a tuple with the ConditionOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionOperator

`func (o *AlertResponseConditionsInner) SetConditionOperator(v string)`

SetConditionOperator sets ConditionOperator field to given value.

### HasConditionOperator

`func (o *AlertResponseConditionsInner) HasConditionOperator() bool`

HasConditionOperator returns a boolean if a field has been set.

### GetConditionTriggerMethod

`func (o *AlertResponseConditionsInner) GetConditionTriggerMethod() string`

GetConditionTriggerMethod returns the ConditionTriggerMethod field if non-nil, zero value otherwise.

### GetConditionTriggerMethodOk

`func (o *AlertResponseConditionsInner) GetConditionTriggerMethodOk() (*string, bool)`

GetConditionTriggerMethodOk returns a tuple with the ConditionTriggerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTriggerMethod

`func (o *AlertResponseConditionsInner) SetConditionTriggerMethod(v string)`

SetConditionTriggerMethod sets ConditionTriggerMethod field to given value.

### HasConditionTriggerMethod

`func (o *AlertResponseConditionsInner) HasConditionTriggerMethod() bool`

HasConditionTriggerMethod returns a boolean if a field has been set.

### GetConditionValue

`func (o *AlertResponseConditionsInner) GetConditionValue() string`

GetConditionValue returns the ConditionValue field if non-nil, zero value otherwise.

### GetConditionValueOk

`func (o *AlertResponseConditionsInner) GetConditionValueOk() (*string, bool)`

GetConditionValueOk returns a tuple with the ConditionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionValue

`func (o *AlertResponseConditionsInner) SetConditionValue(v string)`

SetConditionValue sets ConditionValue field to given value.

### HasConditionValue

`func (o *AlertResponseConditionsInner) HasConditionValue() bool`

HasConditionValue returns a boolean if a field has been set.

### GetConditionLogicBind

`func (o *AlertResponseConditionsInner) GetConditionLogicBind() string`

GetConditionLogicBind returns the ConditionLogicBind field if non-nil, zero value otherwise.

### GetConditionLogicBindOk

`func (o *AlertResponseConditionsInner) GetConditionLogicBindOk() (*string, bool)`

GetConditionLogicBindOk returns a tuple with the ConditionLogicBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionLogicBind

`func (o *AlertResponseConditionsInner) SetConditionLogicBind(v string)`

SetConditionLogicBind sets ConditionLogicBind field to given value.

### HasConditionLogicBind

`func (o *AlertResponseConditionsInner) HasConditionLogicBind() bool`

HasConditionLogicBind returns a boolean if a field has been set.

### GetConditionTimeZone

`func (o *AlertResponseConditionsInner) GetConditionTimeZone() string`

GetConditionTimeZone returns the ConditionTimeZone field if non-nil, zero value otherwise.

### GetConditionTimeZoneOk

`func (o *AlertResponseConditionsInner) GetConditionTimeZoneOk() (*string, bool)`

GetConditionTimeZoneOk returns a tuple with the ConditionTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTimeZone

`func (o *AlertResponseConditionsInner) SetConditionTimeZone(v string)`

SetConditionTimeZone sets ConditionTimeZone field to given value.

### HasConditionTimeZone

`func (o *AlertResponseConditionsInner) HasConditionTimeZone() bool`

HasConditionTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


