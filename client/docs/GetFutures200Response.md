# GetFutures200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to [**[]FuturesInner**](FuturesInner.md) |  | [optional] 

## Methods

### NewGetFutures200Response

`func NewGetFutures200Response() *GetFutures200Response`

NewGetFutures200Response instantiates a new GetFutures200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFutures200ResponseWithDefaults

`func NewGetFutures200ResponseWithDefaults() *GetFutures200Response`

NewGetFutures200ResponseWithDefaults instantiates a new GetFutures200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetFutures200Response) GetSymbol() []FuturesInner`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetFutures200Response) GetSymbolOk() (*[]FuturesInner, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetFutures200Response) SetSymbol(v []FuturesInner)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetFutures200Response) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


