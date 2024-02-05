# NotificationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**D** | Pointer to **string** | notification date | [optional] 
**ID** | Pointer to **string** | unique way to reference this notification | [optional] 
**FC** | Pointer to **string** | FYI code, we can use it to find whether the disclaimer is accepted or not in settings | [optional] 
**MD** | Pointer to **string** | content of notification | [optional] 
**MS** | Pointer to **string** | title of notification | [optional] 
**R** | Pointer to **string** | 0-unread, 1-read | [optional] 

## Methods

### NewNotificationsInner

`func NewNotificationsInner() *NotificationsInner`

NewNotificationsInner instantiates a new NotificationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsInnerWithDefaults

`func NewNotificationsInnerWithDefaults() *NotificationsInner`

NewNotificationsInnerWithDefaults instantiates a new NotificationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetD

`func (o *NotificationsInner) GetD() string`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *NotificationsInner) GetDOk() (*string, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *NotificationsInner) SetD(v string)`

SetD sets D field to given value.

### HasD

`func (o *NotificationsInner) HasD() bool`

HasD returns a boolean if a field has been set.

### GetID

`func (o *NotificationsInner) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *NotificationsInner) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *NotificationsInner) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *NotificationsInner) HasID() bool`

HasID returns a boolean if a field has been set.

### GetFC

`func (o *NotificationsInner) GetFC() string`

GetFC returns the FC field if non-nil, zero value otherwise.

### GetFCOk

`func (o *NotificationsInner) GetFCOk() (*string, bool)`

GetFCOk returns a tuple with the FC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFC

`func (o *NotificationsInner) SetFC(v string)`

SetFC sets FC field to given value.

### HasFC

`func (o *NotificationsInner) HasFC() bool`

HasFC returns a boolean if a field has been set.

### GetMD

`func (o *NotificationsInner) GetMD() string`

GetMD returns the MD field if non-nil, zero value otherwise.

### GetMDOk

`func (o *NotificationsInner) GetMDOk() (*string, bool)`

GetMDOk returns a tuple with the MD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMD

`func (o *NotificationsInner) SetMD(v string)`

SetMD sets MD field to given value.

### HasMD

`func (o *NotificationsInner) HasMD() bool`

HasMD returns a boolean if a field has been set.

### GetMS

`func (o *NotificationsInner) GetMS() string`

GetMS returns the MS field if non-nil, zero value otherwise.

### GetMSOk

`func (o *NotificationsInner) GetMSOk() (*string, bool)`

GetMSOk returns a tuple with the MS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMS

`func (o *NotificationsInner) SetMS(v string)`

SetMS sets MS field to given value.

### HasMS

`func (o *NotificationsInner) HasMS() bool`

HasMS returns a boolean if a field has been set.

### GetR

`func (o *NotificationsInner) GetR() string`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *NotificationsInner) GetROk() (*string, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *NotificationsInner) SetR(v string)`

SetR sets R field to given value.

### HasR

`func (o *NotificationsInner) HasR() bool`

HasR returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


