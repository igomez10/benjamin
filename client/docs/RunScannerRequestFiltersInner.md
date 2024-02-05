# RunScannerRequestFiltersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** | Value can be either an integer, double, boolean or a string depending upon the type of filter specified in the code section | [optional] 

## Methods

### NewRunScannerRequestFiltersInner

`func NewRunScannerRequestFiltersInner() *RunScannerRequestFiltersInner`

NewRunScannerRequestFiltersInner instantiates a new RunScannerRequestFiltersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunScannerRequestFiltersInnerWithDefaults

`func NewRunScannerRequestFiltersInnerWithDefaults() *RunScannerRequestFiltersInner`

NewRunScannerRequestFiltersInnerWithDefaults instantiates a new RunScannerRequestFiltersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *RunScannerRequestFiltersInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *RunScannerRequestFiltersInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *RunScannerRequestFiltersInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *RunScannerRequestFiltersInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetValue

`func (o *RunScannerRequestFiltersInner) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RunScannerRequestFiltersInner) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RunScannerRequestFiltersInner) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *RunScannerRequestFiltersInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


