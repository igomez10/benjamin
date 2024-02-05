# SecdefInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conid** | Pointer to **int32** | IBKR contract identifier. | [optional] 
**Currency** | Pointer to **string** | Currency contract trades in. | [optional] 
**CrossCurrency** | Pointer to **bool** | Defines if a derivative contract has a different currency. | [optional] 
**Time** | Pointer to **int32** |  | [optional] 
**ChineseName** | Pointer to **string** | HTML encoded company description in Chinese. | [optional] 
**AllExchanges** | Pointer to **string** | List of exchanges and venues contract trades. | [optional] 
**ListingExchange** | Pointer to **string** | Main trading venue. | [optional] 
**Name** | Pointer to **string** | Company Name. | [optional] 
**AssetClass** | Pointer to **string** | Group of financial instruments which have similar financial characteristics and behave similar in the marketplace. | [optional] 
**Expiry** | Pointer to **string** | Specific data contract expires. | [optional] 
**LastTradingDay** | Pointer to **string** | Final day derivative contract can be traded before delivery of the underlying asset or cash settlement. | [optional] 
**Group** | Pointer to **string** | Potential characteristic of each product. | [optional] 
**PutOrCall** | Pointer to **string** | Defines the right to buy or sell of the underlying security. | [optional] 
**Sector** | Pointer to **string** | The category of the economy. | [optional] 
**SectorGroup** | Pointer to **string** | Stock Group contract belongs too. | [optional] 
**Strike** | Pointer to **float32** | Set price at which a derivative contract can be bought or sold. | [optional] 
**Ticker** | Pointer to **string** | Contract symbol. | [optional] 
**UndConid** | Pointer to **int32** | Underlying contract identifier. | [optional] 
**Multiplier** | Pointer to **int32** | Multiplier for total premium paid or received for derivative contract. | [optional] 
**Type** | Pointer to **string** | Stock type. | [optional] 
**UndComp** | Pointer to **string** | Company name for underlying contract. | [optional] 
**UndSym** | Pointer to **string** | IBKR Symbol for underlying contract. | [optional] 
**HasOptions** | Pointer to **bool** | If contract has an option. | [optional] 
**FullName** | Pointer to **string** | Formatted company name with underlying symbol, expiration, strike, right. | [optional] 
**IsUS** | Pointer to **bool** | If contract is a US contract. Currently supported for stocks, options and warrants. | [optional] 
**IncrementRules** | Pointer to [**SecdefInnerIncrementRules**](SecdefInnerIncrementRules.md) |  | [optional] 

## Methods

### NewSecdefInner

`func NewSecdefInner() *SecdefInner`

NewSecdefInner instantiates a new SecdefInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecdefInnerWithDefaults

`func NewSecdefInnerWithDefaults() *SecdefInner`

NewSecdefInnerWithDefaults instantiates a new SecdefInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConid

`func (o *SecdefInner) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *SecdefInner) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *SecdefInner) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *SecdefInner) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetCurrency

`func (o *SecdefInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SecdefInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SecdefInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SecdefInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCrossCurrency

`func (o *SecdefInner) GetCrossCurrency() bool`

GetCrossCurrency returns the CrossCurrency field if non-nil, zero value otherwise.

### GetCrossCurrencyOk

`func (o *SecdefInner) GetCrossCurrencyOk() (*bool, bool)`

GetCrossCurrencyOk returns a tuple with the CrossCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCurrency

`func (o *SecdefInner) SetCrossCurrency(v bool)`

SetCrossCurrency sets CrossCurrency field to given value.

### HasCrossCurrency

`func (o *SecdefInner) HasCrossCurrency() bool`

HasCrossCurrency returns a boolean if a field has been set.

### GetTime

`func (o *SecdefInner) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SecdefInner) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SecdefInner) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *SecdefInner) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetChineseName

`func (o *SecdefInner) GetChineseName() string`

GetChineseName returns the ChineseName field if non-nil, zero value otherwise.

### GetChineseNameOk

`func (o *SecdefInner) GetChineseNameOk() (*string, bool)`

GetChineseNameOk returns a tuple with the ChineseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChineseName

`func (o *SecdefInner) SetChineseName(v string)`

SetChineseName sets ChineseName field to given value.

### HasChineseName

`func (o *SecdefInner) HasChineseName() bool`

HasChineseName returns a boolean if a field has been set.

### GetAllExchanges

`func (o *SecdefInner) GetAllExchanges() string`

GetAllExchanges returns the AllExchanges field if non-nil, zero value otherwise.

### GetAllExchangesOk

`func (o *SecdefInner) GetAllExchangesOk() (*string, bool)`

GetAllExchangesOk returns a tuple with the AllExchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllExchanges

`func (o *SecdefInner) SetAllExchanges(v string)`

SetAllExchanges sets AllExchanges field to given value.

### HasAllExchanges

`func (o *SecdefInner) HasAllExchanges() bool`

HasAllExchanges returns a boolean if a field has been set.

### GetListingExchange

`func (o *SecdefInner) GetListingExchange() string`

GetListingExchange returns the ListingExchange field if non-nil, zero value otherwise.

### GetListingExchangeOk

`func (o *SecdefInner) GetListingExchangeOk() (*string, bool)`

GetListingExchangeOk returns a tuple with the ListingExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingExchange

`func (o *SecdefInner) SetListingExchange(v string)`

SetListingExchange sets ListingExchange field to given value.

### HasListingExchange

`func (o *SecdefInner) HasListingExchange() bool`

HasListingExchange returns a boolean if a field has been set.

### GetName

`func (o *SecdefInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecdefInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecdefInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecdefInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAssetClass

`func (o *SecdefInner) GetAssetClass() string`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *SecdefInner) GetAssetClassOk() (*string, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *SecdefInner) SetAssetClass(v string)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *SecdefInner) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetExpiry

`func (o *SecdefInner) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *SecdefInner) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *SecdefInner) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *SecdefInner) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLastTradingDay

`func (o *SecdefInner) GetLastTradingDay() string`

GetLastTradingDay returns the LastTradingDay field if non-nil, zero value otherwise.

### GetLastTradingDayOk

`func (o *SecdefInner) GetLastTradingDayOk() (*string, bool)`

GetLastTradingDayOk returns a tuple with the LastTradingDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradingDay

`func (o *SecdefInner) SetLastTradingDay(v string)`

SetLastTradingDay sets LastTradingDay field to given value.

### HasLastTradingDay

`func (o *SecdefInner) HasLastTradingDay() bool`

HasLastTradingDay returns a boolean if a field has been set.

### GetGroup

`func (o *SecdefInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SecdefInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SecdefInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SecdefInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetPutOrCall

`func (o *SecdefInner) GetPutOrCall() string`

GetPutOrCall returns the PutOrCall field if non-nil, zero value otherwise.

### GetPutOrCallOk

`func (o *SecdefInner) GetPutOrCallOk() (*string, bool)`

GetPutOrCallOk returns a tuple with the PutOrCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutOrCall

`func (o *SecdefInner) SetPutOrCall(v string)`

SetPutOrCall sets PutOrCall field to given value.

### HasPutOrCall

`func (o *SecdefInner) HasPutOrCall() bool`

HasPutOrCall returns a boolean if a field has been set.

### GetSector

`func (o *SecdefInner) GetSector() string`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *SecdefInner) GetSectorOk() (*string, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *SecdefInner) SetSector(v string)`

SetSector sets Sector field to given value.

### HasSector

`func (o *SecdefInner) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetSectorGroup

`func (o *SecdefInner) GetSectorGroup() string`

GetSectorGroup returns the SectorGroup field if non-nil, zero value otherwise.

### GetSectorGroupOk

`func (o *SecdefInner) GetSectorGroupOk() (*string, bool)`

GetSectorGroupOk returns a tuple with the SectorGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorGroup

`func (o *SecdefInner) SetSectorGroup(v string)`

SetSectorGroup sets SectorGroup field to given value.

### HasSectorGroup

`func (o *SecdefInner) HasSectorGroup() bool`

HasSectorGroup returns a boolean if a field has been set.

### GetStrike

`func (o *SecdefInner) GetStrike() float32`

GetStrike returns the Strike field if non-nil, zero value otherwise.

### GetStrikeOk

`func (o *SecdefInner) GetStrikeOk() (*float32, bool)`

GetStrikeOk returns a tuple with the Strike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrike

`func (o *SecdefInner) SetStrike(v float32)`

SetStrike sets Strike field to given value.

### HasStrike

`func (o *SecdefInner) HasStrike() bool`

HasStrike returns a boolean if a field has been set.

### GetTicker

`func (o *SecdefInner) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *SecdefInner) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *SecdefInner) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *SecdefInner) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetUndConid

`func (o *SecdefInner) GetUndConid() int32`

GetUndConid returns the UndConid field if non-nil, zero value otherwise.

### GetUndConidOk

`func (o *SecdefInner) GetUndConidOk() (*int32, bool)`

GetUndConidOk returns a tuple with the UndConid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndConid

`func (o *SecdefInner) SetUndConid(v int32)`

SetUndConid sets UndConid field to given value.

### HasUndConid

`func (o *SecdefInner) HasUndConid() bool`

HasUndConid returns a boolean if a field has been set.

### GetMultiplier

`func (o *SecdefInner) GetMultiplier() int32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *SecdefInner) GetMultiplierOk() (*int32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *SecdefInner) SetMultiplier(v int32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *SecdefInner) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetType

`func (o *SecdefInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecdefInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecdefInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecdefInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUndComp

`func (o *SecdefInner) GetUndComp() string`

GetUndComp returns the UndComp field if non-nil, zero value otherwise.

### GetUndCompOk

`func (o *SecdefInner) GetUndCompOk() (*string, bool)`

GetUndCompOk returns a tuple with the UndComp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndComp

`func (o *SecdefInner) SetUndComp(v string)`

SetUndComp sets UndComp field to given value.

### HasUndComp

`func (o *SecdefInner) HasUndComp() bool`

HasUndComp returns a boolean if a field has been set.

### GetUndSym

`func (o *SecdefInner) GetUndSym() string`

GetUndSym returns the UndSym field if non-nil, zero value otherwise.

### GetUndSymOk

`func (o *SecdefInner) GetUndSymOk() (*string, bool)`

GetUndSymOk returns a tuple with the UndSym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndSym

`func (o *SecdefInner) SetUndSym(v string)`

SetUndSym sets UndSym field to given value.

### HasUndSym

`func (o *SecdefInner) HasUndSym() bool`

HasUndSym returns a boolean if a field has been set.

### GetHasOptions

`func (o *SecdefInner) GetHasOptions() bool`

GetHasOptions returns the HasOptions field if non-nil, zero value otherwise.

### GetHasOptionsOk

`func (o *SecdefInner) GetHasOptionsOk() (*bool, bool)`

GetHasOptionsOk returns a tuple with the HasOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOptions

`func (o *SecdefInner) SetHasOptions(v bool)`

SetHasOptions sets HasOptions field to given value.

### HasHasOptions

`func (o *SecdefInner) HasHasOptions() bool`

HasHasOptions returns a boolean if a field has been set.

### GetFullName

`func (o *SecdefInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SecdefInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SecdefInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *SecdefInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetIsUS

`func (o *SecdefInner) GetIsUS() bool`

GetIsUS returns the IsUS field if non-nil, zero value otherwise.

### GetIsUSOk

`func (o *SecdefInner) GetIsUSOk() (*bool, bool)`

GetIsUSOk returns a tuple with the IsUS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUS

`func (o *SecdefInner) SetIsUS(v bool)`

SetIsUS sets IsUS field to given value.

### HasIsUS

`func (o *SecdefInner) HasIsUS() bool`

HasIsUS returns a boolean if a field has been set.

### GetIncrementRules

`func (o *SecdefInner) GetIncrementRules() SecdefInnerIncrementRules`

GetIncrementRules returns the IncrementRules field if non-nil, zero value otherwise.

### GetIncrementRulesOk

`func (o *SecdefInner) GetIncrementRulesOk() (*SecdefInnerIncrementRules, bool)`

GetIncrementRulesOk returns a tuple with the IncrementRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementRules

`func (o *SecdefInner) SetIncrementRules(v SecdefInnerIncrementRules)`

SetIncrementRules sets IncrementRules field to given value.

### HasIncrementRules

`func (o *SecdefInner) HasIncrementRules() bool`

HasIncrementRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


