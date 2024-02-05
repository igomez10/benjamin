# PlaceOrderDeprecated200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **[]string** | Please note here, if the message is a question, you have to reply to question in order to submit the order successfully. See more in the \&quot;/iserver/reply/{replyid}\&quot; endpoint.  | [optional] 

## Methods

### NewPlaceOrderDeprecated200ResponseInner

`func NewPlaceOrderDeprecated200ResponseInner() *PlaceOrderDeprecated200ResponseInner`

NewPlaceOrderDeprecated200ResponseInner instantiates a new PlaceOrderDeprecated200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaceOrderDeprecated200ResponseInnerWithDefaults

`func NewPlaceOrderDeprecated200ResponseInnerWithDefaults() *PlaceOrderDeprecated200ResponseInner`

NewPlaceOrderDeprecated200ResponseInnerWithDefaults instantiates a new PlaceOrderDeprecated200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlaceOrderDeprecated200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlaceOrderDeprecated200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlaceOrderDeprecated200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PlaceOrderDeprecated200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *PlaceOrderDeprecated200ResponseInner) GetMessage() []string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PlaceOrderDeprecated200ResponseInner) GetMessageOk() (*[]string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PlaceOrderDeprecated200ResponseInner) SetMessage(v []string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PlaceOrderDeprecated200ResponseInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


