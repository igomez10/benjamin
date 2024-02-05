# AccountParent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mmc** | Pointer to **[]string** |  | [optional] 
**AccountId** | Pointer to **string** | Account Number for Money Manager Client | [optional] 
**IsMParent** | Pointer to **bool** | Is MM a Parent Account | [optional] 
**IsMChild** | Pointer to **bool** | Is MM a Child Account | [optional] 
**IsMultiplex** | Pointer to **bool** | Is a Multiplex Account. These are account models with individual account being parent and managed account being child. | [optional] 

## Methods

### NewAccountParent

`func NewAccountParent() *AccountParent`

NewAccountParent instantiates a new AccountParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountParentWithDefaults

`func NewAccountParentWithDefaults() *AccountParent`

NewAccountParentWithDefaults instantiates a new AccountParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMmc

`func (o *AccountParent) GetMmc() []string`

GetMmc returns the Mmc field if non-nil, zero value otherwise.

### GetMmcOk

`func (o *AccountParent) GetMmcOk() (*[]string, bool)`

GetMmcOk returns a tuple with the Mmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmc

`func (o *AccountParent) SetMmc(v []string)`

SetMmc sets Mmc field to given value.

### HasMmc

`func (o *AccountParent) HasMmc() bool`

HasMmc returns a boolean if a field has been set.

### GetAccountId

`func (o *AccountParent) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountParent) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountParent) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AccountParent) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetIsMParent

`func (o *AccountParent) GetIsMParent() bool`

GetIsMParent returns the IsMParent field if non-nil, zero value otherwise.

### GetIsMParentOk

`func (o *AccountParent) GetIsMParentOk() (*bool, bool)`

GetIsMParentOk returns a tuple with the IsMParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMParent

`func (o *AccountParent) SetIsMParent(v bool)`

SetIsMParent sets IsMParent field to given value.

### HasIsMParent

`func (o *AccountParent) HasIsMParent() bool`

HasIsMParent returns a boolean if a field has been set.

### GetIsMChild

`func (o *AccountParent) GetIsMChild() bool`

GetIsMChild returns the IsMChild field if non-nil, zero value otherwise.

### GetIsMChildOk

`func (o *AccountParent) GetIsMChildOk() (*bool, bool)`

GetIsMChildOk returns a tuple with the IsMChild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMChild

`func (o *AccountParent) SetIsMChild(v bool)`

SetIsMChild sets IsMChild field to given value.

### HasIsMChild

`func (o *AccountParent) HasIsMChild() bool`

HasIsMChild returns a boolean if a field has been set.

### GetIsMultiplex

`func (o *AccountParent) GetIsMultiplex() bool`

GetIsMultiplex returns the IsMultiplex field if non-nil, zero value otherwise.

### GetIsMultiplexOk

`func (o *AccountParent) GetIsMultiplexOk() (*bool, bool)`

GetIsMultiplexOk returns a tuple with the IsMultiplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiplex

`func (o *AccountParent) SetIsMultiplex(v bool)`

SetIsMultiplex sets IsMultiplex field to given value.

### HasIsMultiplex

`func (o *AccountParent) HasIsMultiplex() bool`

HasIsMultiplex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


