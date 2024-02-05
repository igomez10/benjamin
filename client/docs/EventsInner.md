# EventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndexDateType** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | will be different for different event types | [optional] 
**Conids** | Pointer to **[]string** |  | [optional] 
**IndexDate** | Pointer to **string** | for exmple 20180817T040000+0000 | [optional] 
**Source** | Pointer to **string** | for example RSE | [optional] 
**EventKey** | Pointer to **string** | for example 11662135 | [optional] 
**Tooltips** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewEventsInner

`func NewEventsInner() *EventsInner`

NewEventsInner instantiates a new EventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsInnerWithDefaults

`func NewEventsInnerWithDefaults() *EventsInner`

NewEventsInnerWithDefaults instantiates a new EventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexDateType

`func (o *EventsInner) GetIndexDateType() string`

GetIndexDateType returns the IndexDateType field if non-nil, zero value otherwise.

### GetIndexDateTypeOk

`func (o *EventsInner) GetIndexDateTypeOk() (*string, bool)`

GetIndexDateTypeOk returns a tuple with the IndexDateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDateType

`func (o *EventsInner) SetIndexDateType(v string)`

SetIndexDateType sets IndexDateType field to given value.

### HasIndexDateType

`func (o *EventsInner) HasIndexDateType() bool`

HasIndexDateType returns a boolean if a field has been set.

### GetEventType

`func (o *EventsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *EventsInner) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetData

`func (o *EventsInner) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventsInner) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventsInner) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *EventsInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetConids

`func (o *EventsInner) GetConids() []string`

GetConids returns the Conids field if non-nil, zero value otherwise.

### GetConidsOk

`func (o *EventsInner) GetConidsOk() (*[]string, bool)`

GetConidsOk returns a tuple with the Conids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConids

`func (o *EventsInner) SetConids(v []string)`

SetConids sets Conids field to given value.

### HasConids

`func (o *EventsInner) HasConids() bool`

HasConids returns a boolean if a field has been set.

### GetIndexDate

`func (o *EventsInner) GetIndexDate() string`

GetIndexDate returns the IndexDate field if non-nil, zero value otherwise.

### GetIndexDateOk

`func (o *EventsInner) GetIndexDateOk() (*string, bool)`

GetIndexDateOk returns a tuple with the IndexDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDate

`func (o *EventsInner) SetIndexDate(v string)`

SetIndexDate sets IndexDate field to given value.

### HasIndexDate

`func (o *EventsInner) HasIndexDate() bool`

HasIndexDate returns a boolean if a field has been set.

### GetSource

`func (o *EventsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EventsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EventsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *EventsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetEventKey

`func (o *EventsInner) GetEventKey() string`

GetEventKey returns the EventKey field if non-nil, zero value otherwise.

### GetEventKeyOk

`func (o *EventsInner) GetEventKeyOk() (*string, bool)`

GetEventKeyOk returns a tuple with the EventKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventKey

`func (o *EventsInner) SetEventKey(v string)`

SetEventKey sets EventKey field to given value.

### HasEventKey

`func (o *EventsInner) HasEventKey() bool`

HasEventKey returns a boolean if a field has been set.

### GetTooltips

`func (o *EventsInner) GetTooltips() map[string]interface{}`

GetTooltips returns the Tooltips field if non-nil, zero value otherwise.

### GetTooltipsOk

`func (o *EventsInner) GetTooltipsOk() (*map[string]interface{}, bool)`

GetTooltipsOk returns a tuple with the Tooltips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTooltips

`func (o *EventsInner) SetTooltips(v map[string]interface{})`

SetTooltips sets Tooltips field to given value.

### HasTooltips

`func (o *EventsInner) HasTooltips() bool`

HasTooltips returns a boolean if a field has been set.

### GetStatus

`func (o *EventsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


