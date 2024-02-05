# StocksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | company name | [optional] 
**ChineseName** | Pointer to **string** | company name in Chinese | [optional] 
**AssetClass** | Pointer to **string** |  | [optional] 
**Contracts** | Pointer to [**[]StocksInnerContractsInner**](StocksInnerContractsInner.md) | array of contracts from different exchanges | [optional] 

## Methods

### NewStocksInner

`func NewStocksInner() *StocksInner`

NewStocksInner instantiates a new StocksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStocksInnerWithDefaults

`func NewStocksInnerWithDefaults() *StocksInner`

NewStocksInnerWithDefaults instantiates a new StocksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StocksInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StocksInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StocksInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StocksInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetChineseName

`func (o *StocksInner) GetChineseName() string`

GetChineseName returns the ChineseName field if non-nil, zero value otherwise.

### GetChineseNameOk

`func (o *StocksInner) GetChineseNameOk() (*string, bool)`

GetChineseNameOk returns a tuple with the ChineseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChineseName

`func (o *StocksInner) SetChineseName(v string)`

SetChineseName sets ChineseName field to given value.

### HasChineseName

`func (o *StocksInner) HasChineseName() bool`

HasChineseName returns a boolean if a field has been set.

### GetAssetClass

`func (o *StocksInner) GetAssetClass() string`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *StocksInner) GetAssetClassOk() (*string, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *StocksInner) SetAssetClass(v string)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *StocksInner) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetContracts

`func (o *StocksInner) GetContracts() []StocksInnerContractsInner`

GetContracts returns the Contracts field if non-nil, zero value otherwise.

### GetContractsOk

`func (o *StocksInner) GetContractsOk() (*[]StocksInnerContractsInner, bool)`

GetContractsOk returns a tuple with the Contracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContracts

`func (o *StocksInner) SetContracts(v []StocksInnerContractsInner)`

SetContracts sets Contracts field to given value.

### HasContracts

`func (o *StocksInner) HasContracts() bool`

HasContracts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


