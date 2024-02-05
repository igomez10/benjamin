# RunScannerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instrument** | Pointer to **string** |  | [optional] 
**Locations** | Pointer to **string** |  | [optional] 
**ScanCode** | Pointer to **string** |  | [optional] 
**SecType** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]RunScannerRequestFiltersInner**](RunScannerRequestFiltersInner.md) |  | [optional] 

## Methods

### NewRunScannerRequest

`func NewRunScannerRequest() *RunScannerRequest`

NewRunScannerRequest instantiates a new RunScannerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunScannerRequestWithDefaults

`func NewRunScannerRequestWithDefaults() *RunScannerRequest`

NewRunScannerRequestWithDefaults instantiates a new RunScannerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrument

`func (o *RunScannerRequest) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *RunScannerRequest) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *RunScannerRequest) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *RunScannerRequest) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetLocations

`func (o *RunScannerRequest) GetLocations() string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *RunScannerRequest) GetLocationsOk() (*string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *RunScannerRequest) SetLocations(v string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *RunScannerRequest) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetScanCode

`func (o *RunScannerRequest) GetScanCode() string`

GetScanCode returns the ScanCode field if non-nil, zero value otherwise.

### GetScanCodeOk

`func (o *RunScannerRequest) GetScanCodeOk() (*string, bool)`

GetScanCodeOk returns a tuple with the ScanCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanCode

`func (o *RunScannerRequest) SetScanCode(v string)`

SetScanCode sets ScanCode field to given value.

### HasScanCode

`func (o *RunScannerRequest) HasScanCode() bool`

HasScanCode returns a boolean if a field has been set.

### GetSecType

`func (o *RunScannerRequest) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *RunScannerRequest) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *RunScannerRequest) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *RunScannerRequest) HasSecType() bool`

HasSecType returns a boolean if a field has been set.

### GetFilters

`func (o *RunScannerRequest) GetFilters() []RunScannerRequestFiltersInner`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RunScannerRequest) GetFiltersOk() (*[]RunScannerRequestFiltersInner, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RunScannerRequest) SetFilters(v []RunScannerRequestFiltersInner)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RunScannerRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


