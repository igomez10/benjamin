# SearchBySymbolOrName200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conid** | Pointer to **int32** | Contract Identifier | [optional] 
**CompanyHeader** | Pointer to **string** | Company Name - Exchange | [optional] 
**CompanyName** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** | Underlying symbol | [optional] 
**Description** | Pointer to **string** | Exchange | [optional] 
**Restricted** | Pointer to **string** |  | [optional] 
**Fop** | Pointer to **string** | List of Future Option expirations in YYYMMDD format separated by semicolon | [optional] 
**Opt** | Pointer to **string** | List of Option expirations in YYYYMMDD format separated by semicolon | [optional] 
**War** | Pointer to **string** | List of Warrant expirations in YYYYMMDD format separated by semicolon | [optional] 
**Sections** | Pointer to [**[]SearchBySymbolOrName200ResponseInnerSectionsInner**](SearchBySymbolOrName200ResponseInnerSectionsInner.md) |  | [optional] 

## Methods

### NewSearchBySymbolOrName200ResponseInner

`func NewSearchBySymbolOrName200ResponseInner() *SearchBySymbolOrName200ResponseInner`

NewSearchBySymbolOrName200ResponseInner instantiates a new SearchBySymbolOrName200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchBySymbolOrName200ResponseInnerWithDefaults

`func NewSearchBySymbolOrName200ResponseInnerWithDefaults() *SearchBySymbolOrName200ResponseInner`

NewSearchBySymbolOrName200ResponseInnerWithDefaults instantiates a new SearchBySymbolOrName200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConid

`func (o *SearchBySymbolOrName200ResponseInner) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *SearchBySymbolOrName200ResponseInner) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *SearchBySymbolOrName200ResponseInner) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *SearchBySymbolOrName200ResponseInner) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetCompanyHeader

`func (o *SearchBySymbolOrName200ResponseInner) GetCompanyHeader() string`

GetCompanyHeader returns the CompanyHeader field if non-nil, zero value otherwise.

### GetCompanyHeaderOk

`func (o *SearchBySymbolOrName200ResponseInner) GetCompanyHeaderOk() (*string, bool)`

GetCompanyHeaderOk returns a tuple with the CompanyHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyHeader

`func (o *SearchBySymbolOrName200ResponseInner) SetCompanyHeader(v string)`

SetCompanyHeader sets CompanyHeader field to given value.

### HasCompanyHeader

`func (o *SearchBySymbolOrName200ResponseInner) HasCompanyHeader() bool`

HasCompanyHeader returns a boolean if a field has been set.

### GetCompanyName

`func (o *SearchBySymbolOrName200ResponseInner) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *SearchBySymbolOrName200ResponseInner) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *SearchBySymbolOrName200ResponseInner) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *SearchBySymbolOrName200ResponseInner) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetSymbol

`func (o *SearchBySymbolOrName200ResponseInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SearchBySymbolOrName200ResponseInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SearchBySymbolOrName200ResponseInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SearchBySymbolOrName200ResponseInner) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *SearchBySymbolOrName200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchBySymbolOrName200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchBySymbolOrName200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SearchBySymbolOrName200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRestricted

`func (o *SearchBySymbolOrName200ResponseInner) GetRestricted() string`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *SearchBySymbolOrName200ResponseInner) GetRestrictedOk() (*string, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *SearchBySymbolOrName200ResponseInner) SetRestricted(v string)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *SearchBySymbolOrName200ResponseInner) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetFop

`func (o *SearchBySymbolOrName200ResponseInner) GetFop() string`

GetFop returns the Fop field if non-nil, zero value otherwise.

### GetFopOk

`func (o *SearchBySymbolOrName200ResponseInner) GetFopOk() (*string, bool)`

GetFopOk returns a tuple with the Fop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFop

`func (o *SearchBySymbolOrName200ResponseInner) SetFop(v string)`

SetFop sets Fop field to given value.

### HasFop

`func (o *SearchBySymbolOrName200ResponseInner) HasFop() bool`

HasFop returns a boolean if a field has been set.

### GetOpt

`func (o *SearchBySymbolOrName200ResponseInner) GetOpt() string`

GetOpt returns the Opt field if non-nil, zero value otherwise.

### GetOptOk

`func (o *SearchBySymbolOrName200ResponseInner) GetOptOk() (*string, bool)`

GetOptOk returns a tuple with the Opt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpt

`func (o *SearchBySymbolOrName200ResponseInner) SetOpt(v string)`

SetOpt sets Opt field to given value.

### HasOpt

`func (o *SearchBySymbolOrName200ResponseInner) HasOpt() bool`

HasOpt returns a boolean if a field has been set.

### GetWar

`func (o *SearchBySymbolOrName200ResponseInner) GetWar() string`

GetWar returns the War field if non-nil, zero value otherwise.

### GetWarOk

`func (o *SearchBySymbolOrName200ResponseInner) GetWarOk() (*string, bool)`

GetWarOk returns a tuple with the War field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWar

`func (o *SearchBySymbolOrName200ResponseInner) SetWar(v string)`

SetWar sets War field to given value.

### HasWar

`func (o *SearchBySymbolOrName200ResponseInner) HasWar() bool`

HasWar returns a boolean if a field has been set.

### GetSections

`func (o *SearchBySymbolOrName200ResponseInner) GetSections() []SearchBySymbolOrName200ResponseInnerSectionsInner`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *SearchBySymbolOrName200ResponseInner) GetSectionsOk() (*[]SearchBySymbolOrName200ResponseInnerSectionsInner, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *SearchBySymbolOrName200ResponseInner) SetSections(v []SearchBySymbolOrName200ResponseInnerSectionsInner)`

SetSections sets Sections field to given value.

### HasSections

`func (o *SearchBySymbolOrName200ResponseInner) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


