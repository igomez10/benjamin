# PaTransactionsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctIds** | Pointer to **[]string** |  | [optional] 
**Conids** | Pointer to **[]float32** |  | [optional] 
**Currency** | Pointer to **string** | optional defaults to USD. | [optional] 
**Days** | Pointer to **float32** | optional, default value is 90 | [optional] 

## Methods

### NewPaTransactionsPostRequest

`func NewPaTransactionsPostRequest() *PaTransactionsPostRequest`

NewPaTransactionsPostRequest instantiates a new PaTransactionsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaTransactionsPostRequestWithDefaults

`func NewPaTransactionsPostRequestWithDefaults() *PaTransactionsPostRequest`

NewPaTransactionsPostRequestWithDefaults instantiates a new PaTransactionsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctIds

`func (o *PaTransactionsPostRequest) GetAcctIds() []string`

GetAcctIds returns the AcctIds field if non-nil, zero value otherwise.

### GetAcctIdsOk

`func (o *PaTransactionsPostRequest) GetAcctIdsOk() (*[]string, bool)`

GetAcctIdsOk returns a tuple with the AcctIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctIds

`func (o *PaTransactionsPostRequest) SetAcctIds(v []string)`

SetAcctIds sets AcctIds field to given value.

### HasAcctIds

`func (o *PaTransactionsPostRequest) HasAcctIds() bool`

HasAcctIds returns a boolean if a field has been set.

### GetConids

`func (o *PaTransactionsPostRequest) GetConids() []float32`

GetConids returns the Conids field if non-nil, zero value otherwise.

### GetConidsOk

`func (o *PaTransactionsPostRequest) GetConidsOk() (*[]float32, bool)`

GetConidsOk returns a tuple with the Conids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConids

`func (o *PaTransactionsPostRequest) SetConids(v []float32)`

SetConids sets Conids field to given value.

### HasConids

`func (o *PaTransactionsPostRequest) HasConids() bool`

HasConids returns a boolean if a field has been set.

### GetCurrency

`func (o *PaTransactionsPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaTransactionsPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaTransactionsPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PaTransactionsPostRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDays

`func (o *PaTransactionsPostRequest) GetDays() float32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *PaTransactionsPostRequest) GetDaysOk() (*float32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *PaTransactionsPostRequest) SetDays(v float32)`

SetDays sets Days field to given value.

### HasDays

`func (o *PaTransactionsPostRequest) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


