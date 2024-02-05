# PositionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conid** | Pointer to **float32** | Contract identifier from IBKR&#39;s database. | [optional] 
**Position** | Pointer to **float32** | Number of shares or quantity of the position. | [optional] 
**AvgCost** | Pointer to **float32** | Average cost of the position. | [optional] 

## Methods

### NewPositionData

`func NewPositionData() *PositionData`

NewPositionData instantiates a new PositionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionDataWithDefaults

`func NewPositionDataWithDefaults() *PositionData`

NewPositionDataWithDefaults instantiates a new PositionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConid

`func (o *PositionData) GetConid() float32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *PositionData) GetConidOk() (*float32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *PositionData) SetConid(v float32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *PositionData) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetPosition

`func (o *PositionData) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PositionData) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PositionData) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PositionData) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetAvgCost

`func (o *PositionData) GetAvgCost() float32`

GetAvgCost returns the AvgCost field if non-nil, zero value otherwise.

### GetAvgCostOk

`func (o *PositionData) GetAvgCostOk() (*float32, bool)`

GetAvgCostOk returns a tuple with the AvgCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgCost

`func (o *PositionData) SetAvgCost(v float32)`

SetAvgCost sets AvgCost field to given value.

### HasAvgCost

`func (o *PositionData) HasAvgCost() bool`

HasAvgCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


