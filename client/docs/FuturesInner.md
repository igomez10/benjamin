# FuturesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Conid** | Pointer to **int32** | conid of the future contract | [optional] 
**UnderlyingConid** | Pointer to **int32** |  | [optional] 
**ExpirationDate** | Pointer to **string** |  | [optional] 
**Ltd** | Pointer to **string** | last trading day | [optional] 

## Methods

### NewFuturesInner

`func NewFuturesInner() *FuturesInner`

NewFuturesInner instantiates a new FuturesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuturesInnerWithDefaults

`func NewFuturesInnerWithDefaults() *FuturesInner`

NewFuturesInnerWithDefaults instantiates a new FuturesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *FuturesInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *FuturesInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *FuturesInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *FuturesInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetConid

`func (o *FuturesInner) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *FuturesInner) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *FuturesInner) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *FuturesInner) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetUnderlyingConid

`func (o *FuturesInner) GetUnderlyingConid() int32`

GetUnderlyingConid returns the UnderlyingConid field if non-nil, zero value otherwise.

### GetUnderlyingConidOk

`func (o *FuturesInner) GetUnderlyingConidOk() (*int32, bool)`

GetUnderlyingConidOk returns a tuple with the UnderlyingConid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingConid

`func (o *FuturesInner) SetUnderlyingConid(v int32)`

SetUnderlyingConid sets UnderlyingConid field to given value.

### HasUnderlyingConid

`func (o *FuturesInner) HasUnderlyingConid() bool`

HasUnderlyingConid returns a boolean if a field has been set.

### GetExpirationDate

`func (o *FuturesInner) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *FuturesInner) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *FuturesInner) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *FuturesInner) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLtd

`func (o *FuturesInner) GetLtd() string`

GetLtd returns the Ltd field if non-nil, zero value otherwise.

### GetLtdOk

`func (o *FuturesInner) GetLtdOk() (*string, bool)`

GetLtdOk returns a tuple with the Ltd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLtd

`func (o *FuturesInner) SetLtd(v string)`

SetLtd sets Ltd field to given value.

### HasLtd

`func (o *FuturesInner) HasLtd() bool`

HasLtd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


