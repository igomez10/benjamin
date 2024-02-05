# AllocationInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetClass** | Pointer to [**AllocationInnerAssetClass**](AllocationInnerAssetClass.md) |  | [optional] 
**Sector** | Pointer to [**AllocationInnerSector**](AllocationInnerSector.md) |  | [optional] 
**Group** | Pointer to [**AllocationInnerGroup**](AllocationInnerGroup.md) |  | [optional] 

## Methods

### NewAllocationInner

`func NewAllocationInner() *AllocationInner`

NewAllocationInner instantiates a new AllocationInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationInnerWithDefaults

`func NewAllocationInnerWithDefaults() *AllocationInner`

NewAllocationInnerWithDefaults instantiates a new AllocationInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetClass

`func (o *AllocationInner) GetAssetClass() AllocationInnerAssetClass`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *AllocationInner) GetAssetClassOk() (*AllocationInnerAssetClass, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *AllocationInner) SetAssetClass(v AllocationInnerAssetClass)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *AllocationInner) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetSector

`func (o *AllocationInner) GetSector() AllocationInnerSector`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *AllocationInner) GetSectorOk() (*AllocationInnerSector, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *AllocationInner) SetSector(v AllocationInnerSector)`

SetSector sets Sector field to given value.

### HasSector

`func (o *AllocationInner) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetGroup

`func (o *AllocationInner) GetGroup() AllocationInnerGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AllocationInner) GetGroupOk() (*AllocationInnerGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AllocationInner) SetGroup(v AllocationInnerGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AllocationInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


