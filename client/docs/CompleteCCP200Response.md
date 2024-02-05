# CompleteCCP200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passed** | Pointer to **bool** | If sso authentication completed | [optional] 
**Authenticated** | Pointer to **bool** | If connection is authenticated | [optional] 
**Connected** | Pointer to **bool** | Connected to CCP session | [optional] 
**Competing** | Pointer to **bool** | If user already has an existing brokerage session running. | [optional] 

## Methods

### NewCompleteCCP200Response

`func NewCompleteCCP200Response() *CompleteCCP200Response`

NewCompleteCCP200Response instantiates a new CompleteCCP200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteCCP200ResponseWithDefaults

`func NewCompleteCCP200ResponseWithDefaults() *CompleteCCP200Response`

NewCompleteCCP200ResponseWithDefaults instantiates a new CompleteCCP200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassed

`func (o *CompleteCCP200Response) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *CompleteCCP200Response) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *CompleteCCP200Response) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *CompleteCCP200Response) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### GetAuthenticated

`func (o *CompleteCCP200Response) GetAuthenticated() bool`

GetAuthenticated returns the Authenticated field if non-nil, zero value otherwise.

### GetAuthenticatedOk

`func (o *CompleteCCP200Response) GetAuthenticatedOk() (*bool, bool)`

GetAuthenticatedOk returns a tuple with the Authenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticated

`func (o *CompleteCCP200Response) SetAuthenticated(v bool)`

SetAuthenticated sets Authenticated field to given value.

### HasAuthenticated

`func (o *CompleteCCP200Response) HasAuthenticated() bool`

HasAuthenticated returns a boolean if a field has been set.

### GetConnected

`func (o *CompleteCCP200Response) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *CompleteCCP200Response) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *CompleteCCP200Response) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *CompleteCCP200Response) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetCompeting

`func (o *CompleteCCP200Response) GetCompeting() bool`

GetCompeting returns the Competing field if non-nil, zero value otherwise.

### GetCompetingOk

`func (o *CompleteCCP200Response) GetCompetingOk() (*bool, bool)`

GetCompetingOk returns a tuple with the Competing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompeting

`func (o *CompleteCCP200Response) SetCompeting(v bool)`

SetCompeting sets Competing field to given value.

### HasCompeting

`func (o *CompleteCCP200Response) HasCompeting() bool`

HasCompeting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


