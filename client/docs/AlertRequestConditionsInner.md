# AlertRequestConditionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **int32** | Types: 1-Price, 3-Time, 4-Margin, 5-Trade, 6-Volume, 7: MTA market 8: MTA Position, 9: MTA Acc. Daily PN&amp;  | [optional] 
**Conidex** | Pointer to **string** | conid and exchange. Format supports conid or conid@exchange | [optional] 
**Operator** | Pointer to **string** | optional, operator for the current condition, can be &gt;&#x3D; or &lt;&#x3D; | [optional] 
**TriggerMethod** | Pointer to **string** | optional, only some type of conditions have triggerMethod | [optional] 
**Value** | Pointer to **string** | can not be empty, can pass default value \&quot;*\&quot; | [optional] 
**LogicBind** | Pointer to **string** | \&quot;a\&quot; means \&quot;AND\&quot;, \&quot;o\&quot; means \&quot;OR\&quot;, \&quot;n\&quot; means \&quot;END\&quot;, the last one condition in the condition array should \&quot;n\&quot;  | [optional] 
**TimeZone** | Pointer to **string** | only needed for some MTA alert condition | [optional] 

## Methods

### NewAlertRequestConditionsInner

`func NewAlertRequestConditionsInner() *AlertRequestConditionsInner`

NewAlertRequestConditionsInner instantiates a new AlertRequestConditionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertRequestConditionsInnerWithDefaults

`func NewAlertRequestConditionsInnerWithDefaults() *AlertRequestConditionsInner`

NewAlertRequestConditionsInnerWithDefaults instantiates a new AlertRequestConditionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AlertRequestConditionsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertRequestConditionsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertRequestConditionsInner) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *AlertRequestConditionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConidex

`func (o *AlertRequestConditionsInner) GetConidex() string`

GetConidex returns the Conidex field if non-nil, zero value otherwise.

### GetConidexOk

`func (o *AlertRequestConditionsInner) GetConidexOk() (*string, bool)`

GetConidexOk returns a tuple with the Conidex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConidex

`func (o *AlertRequestConditionsInner) SetConidex(v string)`

SetConidex sets Conidex field to given value.

### HasConidex

`func (o *AlertRequestConditionsInner) HasConidex() bool`

HasConidex returns a boolean if a field has been set.

### GetOperator

`func (o *AlertRequestConditionsInner) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *AlertRequestConditionsInner) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *AlertRequestConditionsInner) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *AlertRequestConditionsInner) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetTriggerMethod

`func (o *AlertRequestConditionsInner) GetTriggerMethod() string`

GetTriggerMethod returns the TriggerMethod field if non-nil, zero value otherwise.

### GetTriggerMethodOk

`func (o *AlertRequestConditionsInner) GetTriggerMethodOk() (*string, bool)`

GetTriggerMethodOk returns a tuple with the TriggerMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerMethod

`func (o *AlertRequestConditionsInner) SetTriggerMethod(v string)`

SetTriggerMethod sets TriggerMethod field to given value.

### HasTriggerMethod

`func (o *AlertRequestConditionsInner) HasTriggerMethod() bool`

HasTriggerMethod returns a boolean if a field has been set.

### GetValue

`func (o *AlertRequestConditionsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AlertRequestConditionsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AlertRequestConditionsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AlertRequestConditionsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLogicBind

`func (o *AlertRequestConditionsInner) GetLogicBind() string`

GetLogicBind returns the LogicBind field if non-nil, zero value otherwise.

### GetLogicBindOk

`func (o *AlertRequestConditionsInner) GetLogicBindOk() (*string, bool)`

GetLogicBindOk returns a tuple with the LogicBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicBind

`func (o *AlertRequestConditionsInner) SetLogicBind(v string)`

SetLogicBind sets LogicBind field to given value.

### HasLogicBind

`func (o *AlertRequestConditionsInner) HasLogicBind() bool`

HasLogicBind returns a boolean if a field has been set.

### GetTimeZone

`func (o *AlertRequestConditionsInner) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AlertRequestConditionsInner) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AlertRequestConditionsInner) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *AlertRequestConditionsInner) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


