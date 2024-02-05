# PortfolioSubaccounts2Get200ResponseSubaccountsInner

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
**ClearingStatus** | Pointer to **string** | Status of the Account   * O &#x3D; Open   * P or N &#x3D; Pending   * A &#x3D; Abandoned   * R &#x3D; Rejected   * C &#x3D; Closed   covestor:     type: boolean     description: Is a Covestor Account   parent:     type: object     properties:       mmc:         type: array         items:           type: string           description: Money Manager Client (MMC) Account       accountId:         type: string         description: Account Number for Money Manager Client       isMParent:         type: boolean         description: Is MM a Parent Account       isMChild:         type: boolean         description: Is MM a Child Account       isMultiplex:         type: boolean         description: Is a Multiplex Account. These are account models with individual account being parent and managed account being child.   desc:     type: string     description: Formatted \&quot;accountId - accountAlias\&quot;  | [optional] 

## Methods

### NewPortfolioSubaccounts2Get200ResponseSubaccountsInner

`func NewPortfolioSubaccounts2Get200ResponseSubaccountsInner() *PortfolioSubaccounts2Get200ResponseSubaccountsInner`

NewPortfolioSubaccounts2Get200ResponseSubaccountsInner instantiates a new PortfolioSubaccounts2Get200ResponseSubaccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortfolioSubaccounts2Get200ResponseSubaccountsInnerWithDefaults

`func NewPortfolioSubaccounts2Get200ResponseSubaccountsInnerWithDefaults() *PortfolioSubaccounts2Get200ResponseSubaccountsInner`

NewPortfolioSubaccounts2Get200ResponseSubaccountsInnerWithDefaults instantiates a new PortfolioSubaccounts2Get200ResponseSubaccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountVan

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountVan() string`

GetAccountVan returns the AccountVan field if non-nil, zero value otherwise.

### GetAccountVanOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountVanOk() (*string, bool)`

GetAccountVanOk returns a tuple with the AccountVan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountVan

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountVan(v string)`

SetAccountVan sets AccountVan field to given value.

### HasAccountVan

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountVan() bool`

HasAccountVan returns a boolean if a field has been set.

### GetAccountTitle

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountTitle() string`

GetAccountTitle returns the AccountTitle field if non-nil, zero value otherwise.

### GetAccountTitleOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountTitleOk() (*string, bool)`

GetAccountTitleOk returns a tuple with the AccountTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTitle

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountTitle(v string)`

SetAccountTitle sets AccountTitle field to given value.

### HasAccountTitle

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountTitle() bool`

HasAccountTitle returns a boolean if a field has been set.

### GetDisplayName

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetAccountAlias

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountAlias() string`

GetAccountAlias returns the AccountAlias field if non-nil, zero value otherwise.

### GetAccountAliasOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountAliasOk() (*string, bool)`

GetAccountAliasOk returns a tuple with the AccountAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAlias

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountAlias(v string)`

SetAccountAlias sets AccountAlias field to given value.

### HasAccountAlias

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountAlias() bool`

HasAccountAlias returns a boolean if a field has been set.

### GetAccountStatus

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountStatus() float32`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetAccountStatusOk() (*float32, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetAccountStatus(v float32)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetType

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTradingType

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetTradingType() string`

GetTradingType returns the TradingType field if non-nil, zero value otherwise.

### GetTradingTypeOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetTradingTypeOk() (*string, bool)`

GetTradingTypeOk returns a tuple with the TradingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingType

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetTradingType(v string)`

SetTradingType sets TradingType field to given value.

### HasTradingType

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasTradingType() bool`

HasTradingType returns a boolean if a field has been set.

### GetFaclient

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetFaclient() bool`

GetFaclient returns the Faclient field if non-nil, zero value otherwise.

### GetFaclientOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetFaclientOk() (*bool, bool)`

GetFaclientOk returns a tuple with the Faclient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaclient

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetFaclient(v bool)`

SetFaclient sets Faclient field to given value.

### HasFaclient

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasFaclient() bool`

HasFaclient returns a boolean if a field has been set.

### GetClearingStatus

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetClearingStatus() string`

GetClearingStatus returns the ClearingStatus field if non-nil, zero value otherwise.

### GetClearingStatusOk

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) GetClearingStatusOk() (*string, bool)`

GetClearingStatusOk returns a tuple with the ClearingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingStatus

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) SetClearingStatus(v string)`

SetClearingStatus sets ClearingStatus field to given value.

### HasClearingStatus

`func (o *PortfolioSubaccounts2Get200ResponseSubaccountsInner) HasClearingStatus() bool`

HasClearingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


