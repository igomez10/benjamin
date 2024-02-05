# IserverAccountsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** | Unique account id | [optional] 
**Aliases** | Pointer to **map[string]interface{}** | Account Id and its alias | [optional] 
**SelectedAccount** | Pointer to **string** |  | [optional] 

## Methods

### NewIserverAccountsGet200Response

`func NewIserverAccountsGet200Response() *IserverAccountsGet200Response`

NewIserverAccountsGet200Response instantiates a new IserverAccountsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIserverAccountsGet200ResponseWithDefaults

`func NewIserverAccountsGet200ResponseWithDefaults() *IserverAccountsGet200Response`

NewIserverAccountsGet200ResponseWithDefaults instantiates a new IserverAccountsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *IserverAccountsGet200Response) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *IserverAccountsGet200Response) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *IserverAccountsGet200Response) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *IserverAccountsGet200Response) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetAliases

`func (o *IserverAccountsGet200Response) GetAliases() map[string]interface{}`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *IserverAccountsGet200Response) GetAliasesOk() (*map[string]interface{}, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *IserverAccountsGet200Response) SetAliases(v map[string]interface{})`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *IserverAccountsGet200Response) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSelectedAccount

`func (o *IserverAccountsGet200Response) GetSelectedAccount() string`

GetSelectedAccount returns the SelectedAccount field if non-nil, zero value otherwise.

### GetSelectedAccountOk

`func (o *IserverAccountsGet200Response) GetSelectedAccountOk() (*string, bool)`

GetSelectedAccountOk returns a tuple with the SelectedAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedAccount

`func (o *IserverAccountsGet200Response) SetSelectedAccount(v string)`

SetSelectedAccount sets SelectedAccount field to given value.

### HasSelectedAccount

`func (o *IserverAccountsGet200Response) HasSelectedAccount() bool`

HasSelectedAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


