# GetCCPStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticated** | Pointer to **bool** | Login session is authenticated to the CCP. | [optional] 
**Connected** | Pointer to **bool** | Login session is connected | [optional] 
**Name** | Pointer to **string** | server name | [optional] 

## Methods

### NewGetCCPStatus200Response

`func NewGetCCPStatus200Response() *GetCCPStatus200Response`

NewGetCCPStatus200Response instantiates a new GetCCPStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCCPStatus200ResponseWithDefaults

`func NewGetCCPStatus200ResponseWithDefaults() *GetCCPStatus200Response`

NewGetCCPStatus200ResponseWithDefaults instantiates a new GetCCPStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticated

`func (o *GetCCPStatus200Response) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *GetCCPStatus200Response) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *GetCCPStatus200Response) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *GetCCPStatus200Response) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetConnected

`func (o *GetCCPStatus200Response) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *GetCCPStatus200Response) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *GetCCPStatus200Response) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *GetCCPStatus200Response) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetName

`func (o *GetCCPStatus200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCCPStatus200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCCPStatus200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCCPStatus200Response) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


