# Transactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | will always be getTransactions | [optional] 
**Currency** | Pointer to **string** | same as request | [optional] 
**IncludesRealTime** | Pointer to **bool** | Indicates whether current day and realtime data is included in the result | [optional] 
**From** | Pointer to **float32** | Period start date. Epoch time, GMT | [optional] 
**To** | Pointer to **float32** | Period end date. Epoch time, GMT | [optional] 
**Transactions** | Pointer to [**[]TransactionsTransactionsInner**](TransactionsTransactionsInner.md) | Sorted by date descending | [optional] 

## Methods

### NewTransactions

`func NewTransactions() *Transactions`

NewTransactions instantiates a new Transactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsWithDefaults

`func NewTransactionsWithDefaults() *Transactions`

NewTransactionsWithDefaults instantiates a new Transactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transactions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transactions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transactions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Transactions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCurrency

`func (o *Transactions) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transactions) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transactions) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Transactions) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIncludesRealTime

`func (o *Transactions) GetIncludesRealTime() bool`

GetIncludesRealTime returns the IncludesRealTime field if non-nil, zero value otherwise.

### GetIncludesRealTimeOk

`func (o *Transactions) GetIncludesRealTimeOk() (*bool, bool)`

GetIncludesRealTimeOk returns a tuple with the IncludesRealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludesRealTime

`func (o *Transactions) SetIncludesRealTime(v bool)`

SetIncludesRealTime sets IncludesRealTime field to given value.

### HasIncludesRealTime

`func (o *Transactions) HasIncludesRealTime() bool`

HasIncludesRealTime returns a boolean if a field has been set.

### GetFrom

`func (o *Transactions) GetFrom() float32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Transactions) GetFromOk() (*float32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Transactions) SetFrom(v float32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *Transactions) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *Transactions) GetTo() float32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Transactions) GetToOk() (*float32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Transactions) SetTo(v float32)`

SetTo sets To field to given value.

### HasTo

`func (o *Transactions) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTransactions

`func (o *Transactions) GetTransactions() []TransactionsTransactionsInner`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Transactions) GetTransactionsOk() (*[]TransactionsTransactionsInner, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Transactions) SetTransactions(v []TransactionsTransactionsInner)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Transactions) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


