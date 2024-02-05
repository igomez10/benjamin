# GetCCPOrders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]OrderData**](OrderData.md) |  | [optional] 

## Methods

### NewGetCCPOrders200Response

`func NewGetCCPOrders200Response() *GetCCPOrders200Response`

NewGetCCPOrders200Response instantiates a new GetCCPOrders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCCPOrders200ResponseWithDefaults

`func NewGetCCPOrders200ResponseWithDefaults() *GetCCPOrders200Response`

NewGetCCPOrders200ResponseWithDefaults instantiates a new GetCCPOrders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *GetCCPOrders200Response) GetOrders() []OrderData`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GetCCPOrders200Response) GetOrdersOk() (*[]OrderData, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GetCCPOrders200Response) SetOrders(v []OrderData)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *GetCCPOrders200Response) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


