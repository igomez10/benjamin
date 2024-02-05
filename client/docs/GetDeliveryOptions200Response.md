# GetDeliveryOptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**M** | Pointer to **int32** | Email option is enabled or not 0-off, 1-on. | [optional] 
**E** | Pointer to [**[]GetDeliveryOptions200ResponseEInner**](GetDeliveryOptions200ResponseEInner.md) |  | [optional] 

## Methods

### NewGetDeliveryOptions200Response

`func NewGetDeliveryOptions200Response() *GetDeliveryOptions200Response`

NewGetDeliveryOptions200Response instantiates a new GetDeliveryOptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeliveryOptions200ResponseWithDefaults

`func NewGetDeliveryOptions200ResponseWithDefaults() *GetDeliveryOptions200Response`

NewGetDeliveryOptions200ResponseWithDefaults instantiates a new GetDeliveryOptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetM

`func (o *GetDeliveryOptions200Response) GetM() int32`

GetM returns the M field if non-nil, zero value otherwise.

### GetMOk

`func (o *GetDeliveryOptions200Response) GetMOk() (*int32, bool)`

GetMOk returns a tuple with the M field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM

`func (o *GetDeliveryOptions200Response) SetM(v int32)`

SetM sets M field to given value.

### HasM

`func (o *GetDeliveryOptions200Response) HasM() bool`

HasM returns a boolean if a field has been set.

### GetE

`func (o *GetDeliveryOptions200Response) GetE() []GetDeliveryOptions200ResponseEInner`

GetE returns the E field if non-nil, zero value otherwise.

### GetEOk

`func (o *GetDeliveryOptions200Response) GetEOk() (*[]GetDeliveryOptions200ResponseEInner, bool)`

GetEOk returns a tuple with the E field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetE

`func (o *GetDeliveryOptions200Response) SetE(v []GetDeliveryOptions200ResponseEInner)`

SetE sets E field to given value.

### HasE

`func (o *GetDeliveryOptions200Response) HasE() bool`

HasE returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


