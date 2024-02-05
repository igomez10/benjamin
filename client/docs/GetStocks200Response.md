# GetStocks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to [**[]StocksInner**](StocksInner.md) | This is an array of object(s), there could be multiple results under same symbol  | [optional] 

## Methods

### NewGetStocks200Response

`func NewGetStocks200Response() *GetStocks200Response`

NewGetStocks200Response instantiates a new GetStocks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStocks200ResponseWithDefaults

`func NewGetStocks200ResponseWithDefaults() *GetStocks200Response`

NewGetStocks200ResponseWithDefaults instantiates a new GetStocks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetStocks200Response) GetSymbol() []StocksInner`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetStocks200Response) GetSymbolOk() (*[]StocksInner, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetStocks200Response) SetSymbol(v []StocksInner)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *GetStocks200Response) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


