# PositionInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctId** | Pointer to **string** |  | [optional] 
**Conid** | Pointer to **int32** |  | [optional] 
**ContractDesc** | Pointer to **string** |  | [optional] 
**AssetClass** | Pointer to **string** |  | [optional] 
**Position** | Pointer to **float32** |  | [optional] 
**MktPrice** | Pointer to **float32** |  | [optional] 
**MktValue** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**AvgCost** | Pointer to **float32** |  | [optional] 
**AvgPrice** | Pointer to **float32** |  | [optional] 
**RealizedPnl** | Pointer to **float32** |  | [optional] 
**UnrealizedPnl** | Pointer to **float32** |  | [optional] 
**Exchs** | Pointer to **string** |  | [optional] 
**Expiry** | Pointer to **string** |  | [optional] 
**PutOrCall** | Pointer to **string** |  | [optional] 
**Multiplier** | Pointer to **float32** |  | [optional] 
**Strike** | Pointer to **float32** |  | [optional] 
**ExerciseStyle** | Pointer to **string** |  | [optional] 
**UndConid** | Pointer to **int32** |  | [optional] 
**ConExchMap** | Pointer to **[]string** |  | [optional] 
**BaseMktValue** | Pointer to **float32** |  | [optional] 
**BaseMktPrice** | Pointer to **float32** |  | [optional] 
**BaseAvgCost** | Pointer to **float32** |  | [optional] 
**BaseAvgPrice** | Pointer to **float32** |  | [optional] 
**BaseRealizedPnl** | Pointer to **float32** |  | [optional] 
**BaseUnrealizedPnl** | Pointer to **float32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**LastTradingDay** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Sector** | Pointer to **string** |  | [optional] 
**SectorGroup** | Pointer to **string** |  | [optional] 
**Ticker** | Pointer to **string** |  | [optional] 
**UndComp** | Pointer to **string** |  | [optional] 
**UndSym** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 

## Methods

### NewPositionInner

`func NewPositionInner() *PositionInner`

NewPositionInner instantiates a new PositionInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionInnerWithDefaults

`func NewPositionInnerWithDefaults() *PositionInner`

NewPositionInnerWithDefaults instantiates a new PositionInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctId

`func (o *PositionInner) GetAcctId() string`

GetAcctId returns the AcctId field if non-nil, zero value otherwise.

### GetAcctIdOk

`func (o *PositionInner) GetAcctIdOk() (*string, bool)`

GetAcctIdOk returns a tuple with the AcctId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctId

`func (o *PositionInner) SetAcctId(v string)`

SetAcctId sets AcctId field to given value.

### HasAcctId

`func (o *PositionInner) HasAcctId() bool`

HasAcctId returns a boolean if a field has been set.

### GetConid

`func (o *PositionInner) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *PositionInner) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *PositionInner) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *PositionInner) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetContractDesc

`func (o *PositionInner) GetContractDesc() string`

GetContractDesc returns the ContractDesc field if non-nil, zero value otherwise.

### GetContractDescOk

`func (o *PositionInner) GetContractDescOk() (*string, bool)`

GetContractDescOk returns a tuple with the ContractDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractDesc

`func (o *PositionInner) SetContractDesc(v string)`

SetContractDesc sets ContractDesc field to given value.

### HasContractDesc

`func (o *PositionInner) HasContractDesc() bool`

HasContractDesc returns a boolean if a field has been set.

### GetAssetClass

`func (o *PositionInner) GetAssetClass() string`

GetAssetClass returns the AssetClass field if non-nil, zero value otherwise.

### GetAssetClassOk

`func (o *PositionInner) GetAssetClassOk() (*string, bool)`

GetAssetClassOk returns a tuple with the AssetClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetClass

`func (o *PositionInner) SetAssetClass(v string)`

SetAssetClass sets AssetClass field to given value.

### HasAssetClass

`func (o *PositionInner) HasAssetClass() bool`

HasAssetClass returns a boolean if a field has been set.

### GetPosition

`func (o *PositionInner) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *PositionInner) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *PositionInner) SetPosition(v float32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *PositionInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetMktPrice

`func (o *PositionInner) GetMktPrice() float32`

GetMktPrice returns the MktPrice field if non-nil, zero value otherwise.

### GetMktPriceOk

`func (o *PositionInner) GetMktPriceOk() (*float32, bool)`

GetMktPriceOk returns a tuple with the MktPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktPrice

`func (o *PositionInner) SetMktPrice(v float32)`

SetMktPrice sets MktPrice field to given value.

### HasMktPrice

`func (o *PositionInner) HasMktPrice() bool`

HasMktPrice returns a boolean if a field has been set.

### GetMktValue

`func (o *PositionInner) GetMktValue() float32`

GetMktValue returns the MktValue field if non-nil, zero value otherwise.

### GetMktValueOk

`func (o *PositionInner) GetMktValueOk() (*float32, bool)`

GetMktValueOk returns a tuple with the MktValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktValue

`func (o *PositionInner) SetMktValue(v float32)`

SetMktValue sets MktValue field to given value.

### HasMktValue

`func (o *PositionInner) HasMktValue() bool`

HasMktValue returns a boolean if a field has been set.

### GetCurrency

`func (o *PositionInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PositionInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PositionInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *PositionInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAvgCost

`func (o *PositionInner) GetAvgCost() float32`

GetAvgCost returns the AvgCost field if non-nil, zero value otherwise.

### GetAvgCostOk

`func (o *PositionInner) GetAvgCostOk() (*float32, bool)`

GetAvgCostOk returns a tuple with the AvgCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgCost

`func (o *PositionInner) SetAvgCost(v float32)`

SetAvgCost sets AvgCost field to given value.

### HasAvgCost

`func (o *PositionInner) HasAvgCost() bool`

HasAvgCost returns a boolean if a field has been set.

### GetAvgPrice

`func (o *PositionInner) GetAvgPrice() float32`

GetAvgPrice returns the AvgPrice field if non-nil, zero value otherwise.

### GetAvgPriceOk

`func (o *PositionInner) GetAvgPriceOk() (*float32, bool)`

GetAvgPriceOk returns a tuple with the AvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPrice

`func (o *PositionInner) SetAvgPrice(v float32)`

SetAvgPrice sets AvgPrice field to given value.

### HasAvgPrice

`func (o *PositionInner) HasAvgPrice() bool`

HasAvgPrice returns a boolean if a field has been set.

### GetRealizedPnl

`func (o *PositionInner) GetRealizedPnl() float32`

GetRealizedPnl returns the RealizedPnl field if non-nil, zero value otherwise.

### GetRealizedPnlOk

`func (o *PositionInner) GetRealizedPnlOk() (*float32, bool)`

GetRealizedPnlOk returns a tuple with the RealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealizedPnl

`func (o *PositionInner) SetRealizedPnl(v float32)`

SetRealizedPnl sets RealizedPnl field to given value.

### HasRealizedPnl

`func (o *PositionInner) HasRealizedPnl() bool`

HasRealizedPnl returns a boolean if a field has been set.

### GetUnrealizedPnl

`func (o *PositionInner) GetUnrealizedPnl() float32`

GetUnrealizedPnl returns the UnrealizedPnl field if non-nil, zero value otherwise.

### GetUnrealizedPnlOk

`func (o *PositionInner) GetUnrealizedPnlOk() (*float32, bool)`

GetUnrealizedPnlOk returns a tuple with the UnrealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrealizedPnl

`func (o *PositionInner) SetUnrealizedPnl(v float32)`

SetUnrealizedPnl sets UnrealizedPnl field to given value.

### HasUnrealizedPnl

`func (o *PositionInner) HasUnrealizedPnl() bool`

HasUnrealizedPnl returns a boolean if a field has been set.

### GetExchs

`func (o *PositionInner) GetExchs() string`

GetExchs returns the Exchs field if non-nil, zero value otherwise.

### GetExchsOk

`func (o *PositionInner) GetExchsOk() (*string, bool)`

GetExchsOk returns a tuple with the Exchs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchs

`func (o *PositionInner) SetExchs(v string)`

SetExchs sets Exchs field to given value.

### HasExchs

`func (o *PositionInner) HasExchs() bool`

HasExchs returns a boolean if a field has been set.

### GetExpiry

`func (o *PositionInner) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PositionInner) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PositionInner) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *PositionInner) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetPutOrCall

`func (o *PositionInner) GetPutOrCall() string`

GetPutOrCall returns the PutOrCall field if non-nil, zero value otherwise.

### GetPutOrCallOk

`func (o *PositionInner) GetPutOrCallOk() (*string, bool)`

GetPutOrCallOk returns a tuple with the PutOrCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutOrCall

`func (o *PositionInner) SetPutOrCall(v string)`

SetPutOrCall sets PutOrCall field to given value.

### HasPutOrCall

`func (o *PositionInner) HasPutOrCall() bool`

HasPutOrCall returns a boolean if a field has been set.

### GetMultiplier

`func (o *PositionInner) GetMultiplier() float32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *PositionInner) GetMultiplierOk() (*float32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *PositionInner) SetMultiplier(v float32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *PositionInner) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetStrike

`func (o *PositionInner) GetStrike() float32`

GetStrike returns the Strike field if non-nil, zero value otherwise.

### GetStrikeOk

`func (o *PositionInner) GetStrikeOk() (*float32, bool)`

GetStrikeOk returns a tuple with the Strike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrike

`func (o *PositionInner) SetStrike(v float32)`

SetStrike sets Strike field to given value.

### HasStrike

`func (o *PositionInner) HasStrike() bool`

HasStrike returns a boolean if a field has been set.

### GetExerciseStyle

`func (o *PositionInner) GetExerciseStyle() string`

GetExerciseStyle returns the ExerciseStyle field if non-nil, zero value otherwise.

### GetExerciseStyleOk

`func (o *PositionInner) GetExerciseStyleOk() (*string, bool)`

GetExerciseStyleOk returns a tuple with the ExerciseStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExerciseStyle

`func (o *PositionInner) SetExerciseStyle(v string)`

SetExerciseStyle sets ExerciseStyle field to given value.

### HasExerciseStyle

`func (o *PositionInner) HasExerciseStyle() bool`

HasExerciseStyle returns a boolean if a field has been set.

### GetUndConid

`func (o *PositionInner) GetUndConid() int32`

GetUndConid returns the UndConid field if non-nil, zero value otherwise.

### GetUndConidOk

`func (o *PositionInner) GetUndConidOk() (*int32, bool)`

GetUndConidOk returns a tuple with the UndConid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndConid

`func (o *PositionInner) SetUndConid(v int32)`

SetUndConid sets UndConid field to given value.

### HasUndConid

`func (o *PositionInner) HasUndConid() bool`

HasUndConid returns a boolean if a field has been set.

### GetConExchMap

`func (o *PositionInner) GetConExchMap() []string`

GetConExchMap returns the ConExchMap field if non-nil, zero value otherwise.

### GetConExchMapOk

`func (o *PositionInner) GetConExchMapOk() (*[]string, bool)`

GetConExchMapOk returns a tuple with the ConExchMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConExchMap

`func (o *PositionInner) SetConExchMap(v []string)`

SetConExchMap sets ConExchMap field to given value.

### HasConExchMap

`func (o *PositionInner) HasConExchMap() bool`

HasConExchMap returns a boolean if a field has been set.

### GetBaseMktValue

`func (o *PositionInner) GetBaseMktValue() float32`

GetBaseMktValue returns the BaseMktValue field if non-nil, zero value otherwise.

### GetBaseMktValueOk

`func (o *PositionInner) GetBaseMktValueOk() (*float32, bool)`

GetBaseMktValueOk returns a tuple with the BaseMktValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMktValue

`func (o *PositionInner) SetBaseMktValue(v float32)`

SetBaseMktValue sets BaseMktValue field to given value.

### HasBaseMktValue

`func (o *PositionInner) HasBaseMktValue() bool`

HasBaseMktValue returns a boolean if a field has been set.

### GetBaseMktPrice

`func (o *PositionInner) GetBaseMktPrice() float32`

GetBaseMktPrice returns the BaseMktPrice field if non-nil, zero value otherwise.

### GetBaseMktPriceOk

`func (o *PositionInner) GetBaseMktPriceOk() (*float32, bool)`

GetBaseMktPriceOk returns a tuple with the BaseMktPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseMktPrice

`func (o *PositionInner) SetBaseMktPrice(v float32)`

SetBaseMktPrice sets BaseMktPrice field to given value.

### HasBaseMktPrice

`func (o *PositionInner) HasBaseMktPrice() bool`

HasBaseMktPrice returns a boolean if a field has been set.

### GetBaseAvgCost

`func (o *PositionInner) GetBaseAvgCost() float32`

GetBaseAvgCost returns the BaseAvgCost field if non-nil, zero value otherwise.

### GetBaseAvgCostOk

`func (o *PositionInner) GetBaseAvgCostOk() (*float32, bool)`

GetBaseAvgCostOk returns a tuple with the BaseAvgCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAvgCost

`func (o *PositionInner) SetBaseAvgCost(v float32)`

SetBaseAvgCost sets BaseAvgCost field to given value.

### HasBaseAvgCost

`func (o *PositionInner) HasBaseAvgCost() bool`

HasBaseAvgCost returns a boolean if a field has been set.

### GetBaseAvgPrice

`func (o *PositionInner) GetBaseAvgPrice() float32`

GetBaseAvgPrice returns the BaseAvgPrice field if non-nil, zero value otherwise.

### GetBaseAvgPriceOk

`func (o *PositionInner) GetBaseAvgPriceOk() (*float32, bool)`

GetBaseAvgPriceOk returns a tuple with the BaseAvgPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseAvgPrice

`func (o *PositionInner) SetBaseAvgPrice(v float32)`

SetBaseAvgPrice sets BaseAvgPrice field to given value.

### HasBaseAvgPrice

`func (o *PositionInner) HasBaseAvgPrice() bool`

HasBaseAvgPrice returns a boolean if a field has been set.

### GetBaseRealizedPnl

`func (o *PositionInner) GetBaseRealizedPnl() float32`

GetBaseRealizedPnl returns the BaseRealizedPnl field if non-nil, zero value otherwise.

### GetBaseRealizedPnlOk

`func (o *PositionInner) GetBaseRealizedPnlOk() (*float32, bool)`

GetBaseRealizedPnlOk returns a tuple with the BaseRealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseRealizedPnl

`func (o *PositionInner) SetBaseRealizedPnl(v float32)`

SetBaseRealizedPnl sets BaseRealizedPnl field to given value.

### HasBaseRealizedPnl

`func (o *PositionInner) HasBaseRealizedPnl() bool`

HasBaseRealizedPnl returns a boolean if a field has been set.

### GetBaseUnrealizedPnl

`func (o *PositionInner) GetBaseUnrealizedPnl() float32`

GetBaseUnrealizedPnl returns the BaseUnrealizedPnl field if non-nil, zero value otherwise.

### GetBaseUnrealizedPnlOk

`func (o *PositionInner) GetBaseUnrealizedPnlOk() (*float32, bool)`

GetBaseUnrealizedPnlOk returns a tuple with the BaseUnrealizedPnl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnrealizedPnl

`func (o *PositionInner) SetBaseUnrealizedPnl(v float32)`

SetBaseUnrealizedPnl sets BaseUnrealizedPnl field to given value.

### HasBaseUnrealizedPnl

`func (o *PositionInner) HasBaseUnrealizedPnl() bool`

HasBaseUnrealizedPnl returns a boolean if a field has been set.

### GetName

`func (o *PositionInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PositionInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PositionInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PositionInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastTradingDay

`func (o *PositionInner) GetLastTradingDay() string`

GetLastTradingDay returns the LastTradingDay field if non-nil, zero value otherwise.

### GetLastTradingDayOk

`func (o *PositionInner) GetLastTradingDayOk() (*string, bool)`

GetLastTradingDayOk returns a tuple with the LastTradingDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTradingDay

`func (o *PositionInner) SetLastTradingDay(v string)`

SetLastTradingDay sets LastTradingDay field to given value.

### HasLastTradingDay

`func (o *PositionInner) HasLastTradingDay() bool`

HasLastTradingDay returns a boolean if a field has been set.

### GetGroup

`func (o *PositionInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PositionInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PositionInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PositionInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSector

`func (o *PositionInner) GetSector() string`

GetSector returns the Sector field if non-nil, zero value otherwise.

### GetSectorOk

`func (o *PositionInner) GetSectorOk() (*string, bool)`

GetSectorOk returns a tuple with the Sector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSector

`func (o *PositionInner) SetSector(v string)`

SetSector sets Sector field to given value.

### HasSector

`func (o *PositionInner) HasSector() bool`

HasSector returns a boolean if a field has been set.

### GetSectorGroup

`func (o *PositionInner) GetSectorGroup() string`

GetSectorGroup returns the SectorGroup field if non-nil, zero value otherwise.

### GetSectorGroupOk

`func (o *PositionInner) GetSectorGroupOk() (*string, bool)`

GetSectorGroupOk returns a tuple with the SectorGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectorGroup

`func (o *PositionInner) SetSectorGroup(v string)`

SetSectorGroup sets SectorGroup field to given value.

### HasSectorGroup

`func (o *PositionInner) HasSectorGroup() bool`

HasSectorGroup returns a boolean if a field has been set.

### GetTicker

`func (o *PositionInner) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *PositionInner) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *PositionInner) SetTicker(v string)`

SetTicker sets Ticker field to given value.

### HasTicker

`func (o *PositionInner) HasTicker() bool`

HasTicker returns a boolean if a field has been set.

### GetUndComp

`func (o *PositionInner) GetUndComp() string`

GetUndComp returns the UndComp field if non-nil, zero value otherwise.

### GetUndCompOk

`func (o *PositionInner) GetUndCompOk() (*string, bool)`

GetUndCompOk returns a tuple with the UndComp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndComp

`func (o *PositionInner) SetUndComp(v string)`

SetUndComp sets UndComp field to given value.

### HasUndComp

`func (o *PositionInner) HasUndComp() bool`

HasUndComp returns a boolean if a field has been set.

### GetUndSym

`func (o *PositionInner) GetUndSym() string`

GetUndSym returns the UndSym field if non-nil, zero value otherwise.

### GetUndSymOk

`func (o *PositionInner) GetUndSymOk() (*string, bool)`

GetUndSymOk returns a tuple with the UndSym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndSym

`func (o *PositionInner) SetUndSym(v string)`

SetUndSym sets UndSym field to given value.

### HasUndSym

`func (o *PositionInner) HasUndSym() bool`

HasUndSym returns a boolean if a field has been set.

### GetFullName

`func (o *PositionInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PositionInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PositionInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *PositionInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetPageSize

`func (o *PositionInner) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PositionInner) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PositionInner) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PositionInner) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetModel

`func (o *PositionInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PositionInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PositionInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PositionInner) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


