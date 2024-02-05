# SearchBySymbolOrNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | symbol or name to be searched | 
**Name** | Pointer to **bool** | should be true if the search is to be performed by name. false by default. | [optional] 
**SecType** | Pointer to **string** | If search is done by name, only the assets provided in this field will be returned. Currently, only STK is supported. | [optional] 

## Methods

### NewSearchBySymbolOrNameRequest

`func NewSearchBySymbolOrNameRequest(symbol string, ) *SearchBySymbolOrNameRequest`

NewSearchBySymbolOrNameRequest instantiates a new SearchBySymbolOrNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBySymbolOrNameRequestWithDefaults

`func NewSearchBySymbolOrNameRequestWithDefaults() *SearchBySymbolOrNameRequest`

NewSearchBySymbolOrNameRequestWithDefaults instantiates a new SearchBySymbolOrNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SearchBySymbolOrNameRequest) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SearchBySymbolOrNameRequest) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SearchBySymbolOrNameRequest) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *SearchBySymbolOrNameRequest) GetName() bool`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchBySymbolOrNameRequest) GetNameOk() (*bool, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchBySymbolOrNameRequest) SetName(v bool)`

SetName sets Name field to given value.

### HasName

`func (o *SearchBySymbolOrNameRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecType

`func (o *SearchBySymbolOrNameRequest) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *SearchBySymbolOrNameRequest) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *SearchBySymbolOrNameRequest) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *SearchBySymbolOrNameRequest) HasSecType() bool`

HasSecType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


