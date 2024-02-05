# PlaceOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orders** | Pointer to [**[]OrderRequest**](OrderRequest.md) | Notes for bracket orders: 1. Children orders will not have its own \&quot;cOID\&quot;, so please donot pass \&quot;cOID\&quot; parameter in child order.Instead, they will have a \&quot;parentId\&quot; which must be equal to \&quot;cOID\&quot; of parent. 2. When you cancel a parent order, it will cancel all bracket orders, when you cancel one child order, it will also cancel its sibling order.  | [optional] 

## Methods

### NewPlaceOrderRequest

`func NewPlaceOrderRequest() *PlaceOrderRequest`

NewPlaceOrderRequest instantiates a new PlaceOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceOrderRequestWithDefaults

`func NewPlaceOrderRequestWithDefaults() *PlaceOrderRequest`

NewPlaceOrderRequestWithDefaults instantiates a new PlaceOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrders

`func (o *PlaceOrderRequest) GetOrders() []OrderRequest`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *PlaceOrderRequest) GetOrdersOk() (*[]OrderRequest, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *PlaceOrderRequest) SetOrders(v []OrderRequest)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *PlaceOrderRequest) HasOrders() bool`

HasOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

