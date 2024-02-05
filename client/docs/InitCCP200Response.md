# InitCCP200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenge** | Pointer to **map[string]interface{}** | Challenge in hex format | [optional] 

## Methods

### NewInitCCP200Response

`func NewInitCCP200Response() *InitCCP200Response`

NewInitCCP200Response instantiates a new InitCCP200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitCCP200ResponseWithDefaults

`func NewInitCCP200ResponseWithDefaults() *InitCCP200Response`

NewInitCCP200ResponseWithDefaults instantiates a new InitCCP200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *InitCCP200Response) GetChallenge() map[string]interface{}`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *InitCCP200Response) GetChallengeOk() (*map[string]interface{}, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *InitCCP200Response) SetChallenge(v map[string]interface{})`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *InitCCP200Response) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


