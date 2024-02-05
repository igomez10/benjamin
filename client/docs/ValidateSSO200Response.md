# ValidateSSO200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LOGIN_TYPE** | Pointer to **float32** | 1 for Live, 2 for Paper | [optional] 
**USER_NAME** | Pointer to **string** | Username | [optional] 
**USER_ID** | Pointer to **float32** | User ID | [optional] 
**Expire** | Pointer to **float32** | Time in milliseconds until session expires. Caller needs to call the again to re-validate session | [optional] 
**RESULT** | Pointer to **bool** | true if session was validated; false if not. | [optional] 
**AUTH_TIME** | Pointer to **float32** | Time of session validation | [optional] 

## Methods

### NewValidateSSO200Response

`func NewValidateSSO200Response() *ValidateSSO200Response`

NewValidateSSO200Response instantiates a new ValidateSSO200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateSSO200ResponseWithDefaults

`func NewValidateSSO200ResponseWithDefaults() *ValidateSSO200Response`

NewValidateSSO200ResponseWithDefaults instantiates a new ValidateSSO200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLOGIN_TYPE

`func (o *ValidateSSO200Response) GetLOGIN_TYPE() float32`

GetLOGIN_TYPE returns the LOGIN_TYPE field if non-nil, zero value otherwise.

### GetLOGIN_TYPEOk

`func (o *ValidateSSO200Response) GetLOGIN_TYPEOk() (*float32, bool)`

GetLOGIN_TYPEOk returns a tuple with the LOGIN_TYPE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLOGIN_TYPE

`func (o *ValidateSSO200Response) SetLOGIN_TYPE(v float32)`

SetLOGIN_TYPE sets LOGIN_TYPE field to given value.

### HasLOGIN_TYPE

`func (o *ValidateSSO200Response) HasLOGIN_TYPE() bool`

HasLOGIN_TYPE returns a boolean if a field has been set.

### GetUSER_NAME

`func (o *ValidateSSO200Response) GetUSER_NAME() string`

GetUSER_NAME returns the USER_NAME field if non-nil, zero value otherwise.

### GetUSER_NAMEOk

`func (o *ValidateSSO200Response) GetUSER_NAMEOk() (*string, bool)`

GetUSER_NAMEOk returns a tuple with the USER_NAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_NAME

`func (o *ValidateSSO200Response) SetUSER_NAME(v string)`

SetUSER_NAME sets USER_NAME field to given value.

### HasUSER_NAME

`func (o *ValidateSSO200Response) HasUSER_NAME() bool`

HasUSER_NAME returns a boolean if a field has been set.

### GetUSER_ID

`func (o *ValidateSSO200Response) GetUSER_ID() float32`

GetUSER_ID returns the USER_ID field if non-nil, zero value otherwise.

### GetUSER_IDOk

`func (o *ValidateSSO200Response) GetUSER_IDOk() (*float32, bool)`

GetUSER_IDOk returns a tuple with the USER_ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUSER_ID

`func (o *ValidateSSO200Response) SetUSER_ID(v float32)`

SetUSER_ID sets USER_ID field to given value.

### HasUSER_ID

`func (o *ValidateSSO200Response) HasUSER_ID() bool`

HasUSER_ID returns a boolean if a field has been set.

### GetExpire

`func (o *ValidateSSO200Response) GetExpire() float32`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *ValidateSSO200Response) GetExpireOk() (*float32, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *ValidateSSO200Response) SetExpire(v float32)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *ValidateSSO200Response) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetRESULT

`func (o *ValidateSSO200Response) GetRESULT() bool`

GetRESULT returns the RESULT field if non-nil, zero value otherwise.

### GetRESULTOk

`func (o *ValidateSSO200Response) GetRESULTOk() (*bool, bool)`

GetRESULTOk returns a tuple with the RESULT field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRESULT

`func (o *ValidateSSO200Response) SetRESULT(v bool)`

SetRESULT sets RESULT field to given value.

### HasRESULT

`func (o *ValidateSSO200Response) HasRESULT() bool`

HasRESULT returns a boolean if a field has been set.

### GetAUTH_TIME

`func (o *ValidateSSO200Response) GetAUTH_TIME() float32`

GetAUTH_TIME returns the AUTH_TIME field if non-nil, zero value otherwise.

### GetAUTH_TIMEOk

`func (o *ValidateSSO200Response) GetAUTH_TIMEOk() (*float32, bool)`

GetAUTH_TIMEOk returns a tuple with the AUTH_TIME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAUTH_TIME

`func (o *ValidateSSO200Response) SetAUTH_TIME(v float32)`

SetAUTH_TIME sets AUTH_TIME field to given value.

### HasAUTH_TIME

`func (o *ValidateSSO200Response) HasAUTH_TIME() bool`

HasAUTH_TIME returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


