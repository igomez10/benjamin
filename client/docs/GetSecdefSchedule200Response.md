# GetSecdefSchedule200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Exchange parameter id | [optional] 
**TradeVenueId** | Pointer to **string** | Reference on a trade venue of given exchange parameter | [optional] 
**Schedules** | Pointer to [**[]GetSecdefSchedule200ResponseSchedulesInner**](GetSecdefSchedule200ResponseSchedulesInner.md) | Always contains at least one &#39;tradingTime&#39;  and zero or more &#39;sessionTime&#39; tags | [optional] 

## Methods

### NewGetSecdefSchedule200Response

`func NewGetSecdefSchedule200Response() *GetSecdefSchedule200Response`

NewGetSecdefSchedule200Response instantiates a new GetSecdefSchedule200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecdefSchedule200ResponseWithDefaults

`func NewGetSecdefSchedule200ResponseWithDefaults() *GetSecdefSchedule200Response`

NewGetSecdefSchedule200ResponseWithDefaults instantiates a new GetSecdefSchedule200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSecdefSchedule200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSecdefSchedule200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSecdefSchedule200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetSecdefSchedule200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTradeVenueId

`func (o *GetSecdefSchedule200Response) GetTradeVenueId() string`

GetTradeVenueId returns the TradeVenueId field if non-nil, zero value otherwise.

### GetTradeVenueIdOk

`func (o *GetSecdefSchedule200Response) GetTradeVenueIdOk() (*string, bool)`

GetTradeVenueIdOk returns a tuple with the TradeVenueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeVenueId

`func (o *GetSecdefSchedule200Response) SetTradeVenueId(v string)`

SetTradeVenueId sets TradeVenueId field to given value.

### HasTradeVenueId

`func (o *GetSecdefSchedule200Response) HasTradeVenueId() bool`

HasTradeVenueId returns a boolean if a field has been set.

### GetSchedules

`func (o *GetSecdefSchedule200Response) GetSchedules() []GetSecdefSchedule200ResponseSchedulesInner`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *GetSecdefSchedule200Response) GetSchedulesOk() (*[]GetSecdefSchedule200ResponseSchedulesInner, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *GetSecdefSchedule200Response) SetSchedules(v []GetSecdefSchedule200ResponseSchedulesInner)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *GetSecdefSchedule200Response) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


