# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The account identification value | [optional] 
**AccountId** | Pointer to **string** | The account number | [optional] 
**AccountVan** | Pointer to **string** | The accountAlias | [optional] 
**AccountTitle** | Pointer to **string** | Title of the account | [optional] 
**DisplayName** | Pointer to **string** | Whichever value is not null in this priority | [optional] 
**AccountAlias** | Pointer to **string** | User customizable account alias. Refer to [Configure Account Alias](https://guides.interactivebrokers.com/cp/cp.htm#am/settings/accountalias.htm) for details. | [optional] 
**AccountStatus** | Pointer to **float32** | When the account was opened in unix time. | [optional] 
**Currency** | Pointer to **string** | Base currency of the account. | [optional] 
**Type** | Pointer to **string** | Account Type | [optional] 
**TradingType** | Pointer to **string** | UNI - Deprecated property | [optional] 
**Faclient** | Pointer to **bool** | If an account is a sub-account to a Financial Advisor. | [optional] 
**ClearingStatus** | Pointer to **string** | Status of the Account   * O &#x3D; Open   * P or N &#x3D; Pending   * A &#x3D; Abandoned   * R &#x3D; Rejected   * C &#x3D; Closed  | [optional] 
**Covestor** | Pointer to **bool** | Is a Covestor Account | [optional] 
**Parent** | Pointer to [**AccountParent**](AccountParent.md) |  | [optional] 
**Desc** | Pointer to **string** | Formatted \&quot;accountId - accountAlias\&quot; | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *Account) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Account) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Account) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Account) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountVan

`func (o *Account) GetAccountVan() string`

GetAccountVan returns the AccountVan field if non-nil, zero value otherwise.

### GetAccountVanOk

`func (o *Account) GetAccountVanOk() (*string, bool)`

GetAccountVanOk returns a tuple with the AccountVan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountVan

`func (o *Account) SetAccountVan(v string)`

SetAccountVan sets AccountVan field to given value.

### HasAccountVan

`func (o *Account) HasAccountVan() bool`

HasAccountVan returns a boolean if a field has been set.

### GetAccountTitle

`func (o *Account) GetAccountTitle() string`

GetAccountTitle returns the AccountTitle field if non-nil, zero value otherwise.

### GetAccountTitleOk

`func (o *Account) GetAccountTitleOk() (*string, bool)`

GetAccountTitleOk returns a tuple with the AccountTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTitle

`func (o *Account) SetAccountTitle(v string)`

SetAccountTitle sets AccountTitle field to given value.

### HasAccountTitle

`func (o *Account) HasAccountTitle() bool`

HasAccountTitle returns a boolean if a field has been set.

### GetDisplayName

`func (o *Account) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Account) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Account) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Account) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAccountAlias

`func (o *Account) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *Account) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *Account) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *Account) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetAccountStatus

`func (o *Account) GetAccountStatus() float32`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *Account) GetAccountStatusOk() (*float32, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *Account) SetAccountStatus(v float32)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *Account) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *Account) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Account) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Account) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Account) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Account) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTradingType

`func (o *Account) GetTradingType() string`

GetTradingType returns the TradingType field if non-nil, zero value otherwise.

### GetTradingTypeOk

`func (o *Account) GetTradingTypeOk() (*string, bool)`

GetTradingTypeOk returns a tuple with the TradingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingType

`func (o *Account) SetTradingType(v string)`

SetTradingType sets TradingType field to given value.

### HasTradingType

`func (o *Account) HasTradingType() bool`

HasTradingType returns a boolean if a field has been set.

### GetFaclient

`func (o *Account) GetFaclient() bool`

GetFaclient returns the Faclient field if non-nil, zero value otherwise.

### GetFaclientOk

`func (o *Account) GetFaclientOk() (*bool, bool)`

GetFaclientOk returns a tuple with the Faclient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaclient

`func (o *Account) SetFaclient(v bool)`

SetFaclient sets Faclient field to given value.

### HasFaclient

`func (o *Account) HasFaclient() bool`

HasFaclient returns a boolean if a field has been set.

### GetClearingStatus

`func (o *Account) GetClearingStatus() string`

GetClearingStatus returns the ClearingStatus field if non-nil, zero value otherwise.

### GetClearingStatusOk

`func (o *Account) GetClearingStatusOk() (*string, bool)`

GetClearingStatusOk returns a tuple with the ClearingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingStatus

`func (o *Account) SetClearingStatus(v string)`

SetClearingStatus sets ClearingStatus field to given value.

### HasClearingStatus

`func (o *Account) HasClearingStatus() bool`

HasClearingStatus returns a boolean if a field has been set.

### GetCovestor

`func (o *Account) GetCovestor() bool`

GetCovestor returns the Covestor field if non-nil, zero value otherwise.

### GetCovestorOk

`func (o *Account) GetCovestorOk() (*bool, bool)`

GetCovestorOk returns a tuple with the Covestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCovestor

`func (o *Account) SetCovestor(v bool)`

SetCovestor sets Covestor field to given value.

### HasCovestor

`func (o *Account) HasCovestor() bool`

HasCovestor returns a boolean if a field has been set.

### GetParent

`func (o *Account) GetParent() AccountParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Account) GetParentOk() (*AccountParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Account) SetParent(v AccountParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Account) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetDesc

`func (o *Account) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *Account) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *Account) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *Account) HasDesc() bool`

HasDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


