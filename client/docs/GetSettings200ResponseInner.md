# GetSettings200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | Pointer to **int32** | optional, if A doesn&#39;t exist, it means user can&#39;t toggle this option. 0-off, 1-on. | [optional] 
**FC** | Pointer to **string** | fyi code | [optional] 
**H** | Pointer to **int32** | disclaimer read, 1 &#x3D; yes, &#x3D; 0 no. | [optional] 
**FD** | Pointer to **string** | detailed description | [optional] 
**FN** | Pointer to **string** | title | [optional] 

## Methods

### NewGetSettings200ResponseInner

`func NewGetSettings200ResponseInner() *GetSettings200ResponseInner`

NewGetSettings200ResponseInner instantiates a new GetSettings200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSettings200ResponseInnerWithDefaults

`func NewGetSettings200ResponseInnerWithDefaults() *GetSettings200ResponseInner`

NewGetSettings200ResponseInnerWithDefaults instantiates a new GetSettings200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *GetSettings200ResponseInner) GetA() int32`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *GetSettings200ResponseInner) GetAOk() (*int32, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *GetSettings200ResponseInner) SetA(v int32)`

SetA sets A field to given value.

### HasA

`func (o *GetSettings200ResponseInner) HasA() bool`

HasA returns a boolean if a field has been set.

### GetFC

`func (o *GetSettings200ResponseInner) GetFC() string`

GetFC returns the FC field if non-nil, zero value otherwise.

### GetFCOk

`func (o *GetSettings200ResponseInner) GetFCOk() (*string, bool)`

GetFCOk returns a tuple with the FC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFC

`func (o *GetSettings200ResponseInner) SetFC(v string)`

SetFC sets FC field to given value.

### HasFC

`func (o *GetSettings200ResponseInner) HasFC() bool`

HasFC returns a boolean if a field has been set.

### GetH

`func (o *GetSettings200ResponseInner) GetH() int32`

GetH returns the H field if non-nil, zero value otherwise.

### GetHOk

`func (o *GetSettings200ResponseInner) GetHOk() (*int32, bool)`

GetHOk returns a tuple with the H field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH

`func (o *GetSettings200ResponseInner) SetH(v int32)`

SetH sets H field to given value.

### HasH

`func (o *GetSettings200ResponseInner) HasH() bool`

HasH returns a boolean if a field has been set.

### GetFD

`func (o *GetSettings200ResponseInner) GetFD() string`

GetFD returns the FD field if non-nil, zero value otherwise.

### GetFDOk

`func (o *GetSettings200ResponseInner) GetFDOk() (*string, bool)`

GetFDOk returns a tuple with the FD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFD

`func (o *GetSettings200ResponseInner) SetFD(v string)`

SetFD sets FD field to given value.

### HasFD

`func (o *GetSettings200ResponseInner) HasFD() bool`

HasFD returns a boolean if a field has been set.

### GetFN

`func (o *GetSettings200ResponseInner) GetFN() string`

GetFN returns the FN field if non-nil, zero value otherwise.

### GetFNOk

`func (o *GetSettings200ResponseInner) GetFNOk() (*string, bool)`

GetFNOk returns a tuple with the FN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFN

`func (o *GetSettings200ResponseInner) SetFN(v string)`

SetFN sets FN field to given value.

### HasFN

`func (o *GetSettings200ResponseInner) HasFN() bool`

HasFN returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


