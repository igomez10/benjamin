# GetLiveOrders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to **[]string** |  | [optional] 
**Orders** | Pointer to [**[]GetLiveOrders200ResponseOrdersInner**](GetLiveOrders200ResponseOrdersInner.md) |  | [optional] 
**Snapshot** | Pointer to **bool** | If live order update is a snapshot | [optional] 

## Methods

### NewGetLiveOrders200Response

`func NewGetLiveOrders200Response() *GetLiveOrders200Response`

NewGetLiveOrders200Response instantiates a new GetLiveOrders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLiveOrders200ResponseWithDefaults

`func NewGetLiveOrders200ResponseWithDefaults() *GetLiveOrders200Response`

NewGetLiveOrders200ResponseWithDefaults instantiates a new GetLiveOrders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *GetLiveOrders200Response) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetLiveOrders200Response) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetLiveOrders200Response) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetLiveOrders200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetOrders

`func (o *GetLiveOrders200Response) GetOrders() []GetLiveOrders200ResponseOrdersInner`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GetLiveOrders200Response) GetOrdersOk() (*[]GetLiveOrders200ResponseOrdersInner, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GetLiveOrders200Response) SetOrders(v []GetLiveOrders200ResponseOrdersInner)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *GetLiveOrders200Response) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetSnapshot

`func (o *GetLiveOrders200Response) GetSnapshot() bool`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *GetLiveOrders200Response) GetSnapshotOk() (*bool, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *GetLiveOrders200Response) SetSnapshot(v bool)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *GetLiveOrders200Response) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


