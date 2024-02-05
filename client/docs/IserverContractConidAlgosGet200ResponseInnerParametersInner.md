# IserverContractConidAlgosGet200ResponseInnerParametersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The algo parameter | 
**Required** | Pointer to **bool** | If true a value must be entered. | [optional] 
**Name** | Pointer to **string** | Descriptive name of the parameter. | [optional] 
**ValueClassName** | **string** | Format of the parameter. | 
**MinValue** | Pointer to **float32** | Smallest value, only applies to parameters with valueClassName&#x3D;Double. | [optional] 
**MaxValue** | Pointer to **float32** | Largest value, only applies to parameters with valueClassName&#x3D;Double. | [optional] 
**DefaultValue** | Pointer to **bool** | User configured preset for this parameter. | [optional] 
**LegalStrings** | Pointer to **string** | The list of choices | [optional] 
**Description** | Pointer to **string** | Detailed description of the parameter. | [optional] 
**GuiRank** | Pointer to **float32** | The order in UI, used when building dynamic UI so that more important parameters are presented first. | [optional] 
**PriceMarketRule** | Pointer to **bool** | If true, must specify parameter using market rule format. Only applies to parameters with valueClassName&#x3D;Double. | [optional] 
**EnabledConditions** | Pointer to **string** | The rules that UI should apply to algo parameters depending on chosen order type:  * MKT:speedUp:&#x3D;:no - hide SpeedUp param when MKT is chosen for order type.  * LMT:strategyType:&lt;&gt;:empty - strategyType param cannot be empty when LMT is chosen for order type.  * MKT:strategyType:&#x3D;:Marketable - set strategyType param to Marketable and disable (no other choice) when MKT is chosen for order type.  | [optional] 

## Methods

### NewIserverContractConidAlgosGet200ResponseInnerParametersInner

`func NewIserverContractConidAlgosGet200ResponseInnerParametersInner(id string, valueClassName string, ) *IserverContractConidAlgosGet200ResponseInnerParametersInner`

NewIserverContractConidAlgosGet200ResponseInnerParametersInner instantiates a new IserverContractConidAlgosGet200ResponseInnerParametersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIserverContractConidAlgosGet200ResponseInnerParametersInnerWithDefaults

`func NewIserverContractConidAlgosGet200ResponseInnerParametersInnerWithDefaults() *IserverContractConidAlgosGet200ResponseInnerParametersInner`

NewIserverContractConidAlgosGet200ResponseInnerParametersInnerWithDefaults instantiates a new IserverContractConidAlgosGet200ResponseInnerParametersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetId(v string)`

SetId sets Id field to given value.


### GetRequired

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetName

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValueClassName

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetValueClassName() string`

GetValueClassName returns the ValueClassName field if non-nil, zero value otherwise.

### GetValueClassNameOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetValueClassNameOk() (*string, bool)`

GetValueClassNameOk returns a tuple with the ValueClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueClassName

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetValueClassName(v string)`

SetValueClassName sets ValueClassName field to given value.


### GetMinValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### GetMaxValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### GetDefaultValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetDefaultValue() bool`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetDefaultValueOk() (*bool, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetDefaultValue(v bool)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetLegalStrings

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetLegalStrings() string`

GetLegalStrings returns the LegalStrings field if non-nil, zero value otherwise.

### GetLegalStringsOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetLegalStringsOk() (*string, bool)`

GetLegalStringsOk returns a tuple with the LegalStrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalStrings

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetLegalStrings(v string)`

SetLegalStrings sets LegalStrings field to given value.

### HasLegalStrings

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasLegalStrings() bool`

HasLegalStrings returns a boolean if a field has been set.

### GetDescription

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGuiRank

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetGuiRank() float32`

GetGuiRank returns the GuiRank field if non-nil, zero value otherwise.

### GetGuiRankOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetGuiRankOk() (*float32, bool)`

GetGuiRankOk returns a tuple with the GuiRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuiRank

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetGuiRank(v float32)`

SetGuiRank sets GuiRank field to given value.

### HasGuiRank

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasGuiRank() bool`

HasGuiRank returns a boolean if a field has been set.

### GetPriceMarketRule

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetPriceMarketRule() bool`

GetPriceMarketRule returns the PriceMarketRule field if non-nil, zero value otherwise.

### GetPriceMarketRuleOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetPriceMarketRuleOk() (*bool, bool)`

GetPriceMarketRuleOk returns a tuple with the PriceMarketRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMarketRule

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetPriceMarketRule(v bool)`

SetPriceMarketRule sets PriceMarketRule field to given value.

### HasPriceMarketRule

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasPriceMarketRule() bool`

HasPriceMarketRule returns a boolean if a field has been set.

### GetEnabledConditions

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetEnabledConditions() string`

GetEnabledConditions returns the EnabledConditions field if non-nil, zero value otherwise.

### GetEnabledConditionsOk

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) GetEnabledConditionsOk() (*string, bool)`

GetEnabledConditionsOk returns a tuple with the EnabledConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledConditions

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) SetEnabledConditions(v string)`

SetEnabledConditions sets EnabledConditions field to given value.

### HasEnabledConditions

`func (o *IserverContractConidAlgosGet200ResponseInnerParametersInner) HasEnabledConditions() bool`

HasEnabledConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


