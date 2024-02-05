# IserverContractRulesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conid** | **string** | IBKR contract identifier | 
**IsBuy** | **bool** | Side of the market rules apply too. Set to **true** for Buy Orders, set to **false** for Sell Orders | 

## Methods

### NewIserverContractRulesPostRequest

`func NewIserverContractRulesPostRequest(conid string, isBuy bool, ) *IserverContractRulesPostRequest`

NewIserverContractRulesPostRequest instantiates a new IserverContractRulesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIserverContractRulesPostRequestWithDefaults

`func NewIserverContractRulesPostRequestWithDefaults() *IserverContractRulesPostRequest`

NewIserverContractRulesPostRequestWithDefaults instantiates a new IserverContractRulesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConid

`func (o *IserverContractRulesPostRequest) GetConid() string`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *IserverContractRulesPostRequest) GetConidOk() (*string, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *IserverContractRulesPostRequest) SetConid(v string)`

SetConid sets Conid field to given value.


### GetIsBuy

`func (o *IserverContractRulesPostRequest) GetIsBuy() bool`

GetIsBuy returns the IsBuy field if non-nil, zero value otherwise.

### GetIsBuyOk

`func (o *IserverContractRulesPostRequest) GetIsBuyOk() (*bool, bool)`

GetIsBuyOk returns a tuple with the IsBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBuy

`func (o *IserverContractRulesPostRequest) SetIsBuy(v bool)`

SetIsBuy sets IsBuy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


