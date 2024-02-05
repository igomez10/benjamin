# IserverScannerRunPost200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** |  | [optional] 
**ColumnName** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** | Underlying symbol | [optional] 
**Conidex** | Pointer to **string** | conid and exchange. Format supports conid or conid@exchange | [optional] 
**ConId** | Pointer to **float32** |  | [optional] 
**AvailableChartPeriods** | Pointer to **string** | List of available chart periods | [optional] 
**CompanyName** | Pointer to **string** | Contracts company name | [optional] 
**ContractDescription1** | Pointer to **string** | Format contract name | [optional] 
**ListingExchange** | Pointer to **string** |  | [optional] 
**SecType** | Pointer to **string** |  | [optional] 

## Methods

### NewIserverScannerRunPost200ResponseInner

`func NewIserverScannerRunPost200ResponseInner() *IserverScannerRunPost200ResponseInner`

NewIserverScannerRunPost200ResponseInner instantiates a new IserverScannerRunPost200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIserverScannerRunPost200ResponseInnerWithDefaults

`func NewIserverScannerRunPost200ResponseInnerWithDefaults() *IserverScannerRunPost200ResponseInner`

NewIserverScannerRunPost200ResponseInnerWithDefaults instantiates a new IserverScannerRunPost200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *IserverScannerRunPost200ResponseInner) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *IserverScannerRunPost200ResponseInner) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *IserverScannerRunPost200ResponseInner) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *IserverScannerRunPost200ResponseInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetColumnName

`func (o *IserverScannerRunPost200ResponseInner) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *IserverScannerRunPost200ResponseInner) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *IserverScannerRunPost200ResponseInner) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *IserverScannerRunPost200ResponseInner) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetSymbol

`func (o *IserverScannerRunPost200ResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *IserverScannerRunPost200ResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *IserverScannerRunPost200ResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *IserverScannerRunPost200ResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetConidex

`func (o *IserverScannerRunPost200ResponseInner) GetConidex() string`

GetConidex returns the Conidex field if non-nil, zero value otherwise.

### GetConidexOk

`func (o *IserverScannerRunPost200ResponseInner) GetConidexOk() (*string, bool)`

GetConidexOk returns a tuple with the Conidex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConidex

`func (o *IserverScannerRunPost200ResponseInner) SetConidex(v string)`

SetConidex sets Conidex field to given value.

### HasConidex

`func (o *IserverScannerRunPost200ResponseInner) HasConidex() bool`

HasConidex returns a boolean if a field has been set.

### GetConId

`func (o *IserverScannerRunPost200ResponseInner) GetConId() float32`

GetConId returns the ConId field if non-nil, zero value otherwise.

### GetConIdOk

`func (o *IserverScannerRunPost200ResponseInner) GetConIdOk() (*float32, bool)`

GetConIdOk returns a tuple with the ConId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConId

`func (o *IserverScannerRunPost200ResponseInner) SetConId(v float32)`

SetConId sets ConId field to given value.

### HasConId

`func (o *IserverScannerRunPost200ResponseInner) HasConId() bool`

HasConId returns a boolean if a field has been set.

### GetAvailableChartPeriods

`func (o *IserverScannerRunPost200ResponseInner) GetAvailableChartPeriods() string`

GetAvailableChartPeriods returns the AvailableChartPeriods field if non-nil, zero value otherwise.

### GetAvailableChartPeriodsOk

`func (o *IserverScannerRunPost200ResponseInner) GetAvailableChartPeriodsOk() (*string, bool)`

GetAvailableChartPeriodsOk returns a tuple with the AvailableChartPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableChartPeriods

`func (o *IserverScannerRunPost200ResponseInner) SetAvailableChartPeriods(v string)`

SetAvailableChartPeriods sets AvailableChartPeriods field to given value.

### HasAvailableChartPeriods

`func (o *IserverScannerRunPost200ResponseInner) HasAvailableChartPeriods() bool`

HasAvailableChartPeriods returns a boolean if a field has been set.

### GetCompanyName

`func (o *IserverScannerRunPost200ResponseInner) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *IserverScannerRunPost200ResponseInner) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *IserverScannerRunPost200ResponseInner) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *IserverScannerRunPost200ResponseInner) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetContractDescription1

`func (o *IserverScannerRunPost200ResponseInner) GetContractDescription1() string`

GetContractDescription1 returns the ContractDescription1 field if non-nil, zero value otherwise.

### GetContractDescription1Ok

`func (o *IserverScannerRunPost200ResponseInner) GetContractDescription1Ok() (*string, bool)`

GetContractDescription1Ok returns a tuple with the ContractDescription1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDescription1

`func (o *IserverScannerRunPost200ResponseInner) SetContractDescription1(v string)`

SetContractDescription1 sets ContractDescription1 field to given value.

### HasContractDescription1

`func (o *IserverScannerRunPost200ResponseInner) HasContractDescription1() bool`

HasContractDescription1 returns a boolean if a field has been set.

### GetListingExchange

`func (o *IserverScannerRunPost200ResponseInner) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *IserverScannerRunPost200ResponseInner) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *IserverScannerRunPost200ResponseInner) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *IserverScannerRunPost200ResponseInner) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetSecType

`func (o *IserverScannerRunPost200ResponseInner) GetSecType() string`

GetSecType returns the SecType field if non-nil, zero value otherwise.

### GetSecTypeOk

`func (o *IserverScannerRunPost200ResponseInner) GetSecTypeOk() (*string, bool)`

GetSecTypeOk returns a tuple with the SecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecType

`func (o *IserverScannerRunPost200ResponseInner) SetSecType(v string)`

SetSecType sets SecType field to given value.

### HasSecType

`func (o *IserverScannerRunPost200ResponseInner) HasSecType() bool`

HasSecType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


