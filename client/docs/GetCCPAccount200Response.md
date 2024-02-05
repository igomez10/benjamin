# GetCCPAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MainAcct** | Pointer to **string** | The primary or parent account. | [optional] 
**AcctList** | Pointer to [**[]GetCCPAccount200ResponseAcctListInner**](GetCCPAccount200ResponseAcctListInner.md) | List of tradeable or Sub Accounts | [optional] 

## Methods

### NewGetCCPAccount200Response

`func NewGetCCPAccount200Response() *GetCCPAccount200Response`

NewGetCCPAccount200Response instantiates a new GetCCPAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCCPAccount200ResponseWithDefaults

`func NewGetCCPAccount200ResponseWithDefaults() *GetCCPAccount200Response`

NewGetCCPAccount200ResponseWithDefaults instantiates a new GetCCPAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMainAcct

`func (o *GetCCPAccount200Response) GetMainAcct() string`

GetMainAcct returns the MainAcct field if non-nil, zero value otherwise.

### GetMainAcctOk

`func (o *GetCCPAccount200Response) GetMainAcctOk() (*string, bool)`

GetMainAcctOk returns a tuple with the MainAcct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainAcct

`func (o *GetCCPAccount200Response) SetMainAcct(v string)`

SetMainAcct sets MainAcct field to given value.

### HasMainAcct

`func (o *GetCCPAccount200Response) HasMainAcct() bool`

HasMainAcct returns a boolean if a field has been set.

### GetAcctList

`func (o *GetCCPAccount200Response) GetAcctList() []GetCCPAccount200ResponseAcctListInner`

GetAcctList returns the AcctList field if non-nil, zero value otherwise.

### GetAcctListOk

`func (o *GetCCPAccount200Response) GetAcctListOk() (*[]GetCCPAccount200ResponseAcctListInner, bool)`

GetAcctListOk returns a tuple with the AcctList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctList

`func (o *GetCCPAccount200Response) SetAcctList(v []GetCCPAccount200ResponseAcctListInner)`

SetAcctList sets AcctList field to given value.

### HasAcctList

`func (o *GetCCPAccount200Response) HasAcctList() bool`

HasAcctList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


