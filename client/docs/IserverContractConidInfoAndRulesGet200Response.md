# IserverContractConidInfoAndRulesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfiCode** | Pointer to **string** | Classification of Financial Instrument codes | [optional] 
**Symbol** | Pointer to **string** | Underlying symbol | [optional] 
**Cusip** | Pointer to **string** |  | [optional] 
**ExpiryFull** | Pointer to **float32** | Expiration Date in the format YYYYMMDD | [optional] 
**ConId** | Pointer to **float32** | IBKRs contract identifier | [optional] 
**MaturityDate** | Pointer to **float32** | Date on which the underlying transaction settles if the option is exercised | [optional] 
**Industry** | Pointer to **string** | Specific group of companies or businesses. | [optional] 
**InstrumentType** | Pointer to **string** | Asset Class of the contract | [optional] 
**TradingClass** | Pointer to **string** | Designation of the contract | [optional] 
**ValidExchanges** | Pointer to **string** | Comma separated list of exchanges or trading venues | [optional] 
**AllowSellLong** | Pointer to **bool** | Allowed to sell shares that you own | [optional] 
**IsZeroCommissionSecurity** | Pointer to **bool** | Supports zero commission trades | [optional] 
**LocalSymbol** | Pointer to **string** | Contracts symbol from primary exchange. For options it is the OCC symbol. | [optional] 
**Classifier** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** | Currency contract trades in | [optional] 
**Text** | Pointer to **string** | Formatted contract parameters | [optional] 
**UnderlyingConId** | Pointer to **float32** | IBKRs contract identifier for the underlying instrument | [optional] 
**RTH** | Pointer to **bool** | Provides trading outside of Regular Trading Hours | [optional] 
**Multiplier** | Pointer to **string** | numerical value of each point of price movement | [optional] 
**Strike** | Pointer to **string** | fixed price at which the owner of the option buys or sells the underlying | [optional] 
**Right** | Pointer to **string** | Put or Call of the option | [optional] 
**UnderlyingIssuer** | Pointer to **string** | Legal entity for underlying contract | [optional] 
**ContractMonth** | Pointer to **string** | Month the contract must be satisfied by making or accepting delivery | [optional] 
**CompanyName** | Pointer to **string** | Contracts company name | [optional] 
**SmartAvailable** | Pointer to **bool** | Support IBKRs SMART routing | [optional] 
**Exchange** | Pointer to **string** | Primary Exchange, Routing or Trading Venue | [optional] 
**Rules** | Pointer to [**[]IserverContractConidInfoAndRulesGet200ResponseRulesInner**](IserverContractConidInfoAndRulesGet200ResponseRulesInner.md) |  | [optional] 

## Methods

### NewIserverContractConidInfoAndRulesGet200Response

`func NewIserverContractConidInfoAndRulesGet200Response() *IserverContractConidInfoAndRulesGet200Response`

NewIserverContractConidInfoAndRulesGet200Response instantiates a new IserverContractConidInfoAndRulesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIserverContractConidInfoAndRulesGet200ResponseWithDefaults

`func NewIserverContractConidInfoAndRulesGet200ResponseWithDefaults() *IserverContractConidInfoAndRulesGet200Response`

NewIserverContractConidInfoAndRulesGet200ResponseWithDefaults instantiates a new IserverContractConidInfoAndRulesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfiCode

`func (o *IserverContractConidInfoAndRulesGet200Response) GetCfiCode() string`

GetCfiCode returns the CfiCode field if non-nil, zero value otherwise.

### GetCfiCodeOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetCfiCodeOk() (*string, bool)`

GetCfiCodeOk returns a tuple with the CfiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfiCode

`func (o *IserverContractConidInfoAndRulesGet200Response) SetCfiCode(v string)`

SetCfiCode sets CfiCode field to given value.

### HasCfiCode

`func (o *IserverContractConidInfoAndRulesGet200Response) HasCfiCode() bool`

HasCfiCode returns a boolean if a field has been set.

### GetSymbol

`func (o *IserverContractConidInfoAndRulesGet200Response) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *IserverContractConidInfoAndRulesGet200Response) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *IserverContractConidInfoAndRulesGet200Response) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetCusip

`func (o *IserverContractConidInfoAndRulesGet200Response) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *IserverContractConidInfoAndRulesGet200Response) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *IserverContractConidInfoAndRulesGet200Response) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetExpiryFull

`func (o *IserverContractConidInfoAndRulesGet200Response) GetExpiryFull() float32`

GetExpiryFull returns the ExpiryFull field if non-nil, zero value otherwise.

### GetExpiryFullOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetExpiryFullOk() (*float32, bool)`

GetExpiryFullOk returns a tuple with the ExpiryFull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryFull

`func (o *IserverContractConidInfoAndRulesGet200Response) SetExpiryFull(v float32)`

SetExpiryFull sets ExpiryFull field to given value.

### HasExpiryFull

`func (o *IserverContractConidInfoAndRulesGet200Response) HasExpiryFull() bool`

HasExpiryFull returns a boolean if a field has been set.

### GetConId

`func (o *IserverContractConidInfoAndRulesGet200Response) GetConId() float32`

GetConId returns the ConId field if non-nil, zero value otherwise.

### GetConIdOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetConIdOk() (*float32, bool)`

GetConIdOk returns a tuple with the ConId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConId

`func (o *IserverContractConidInfoAndRulesGet200Response) SetConId(v float32)`

SetConId sets ConId field to given value.

### HasConId

`func (o *IserverContractConidInfoAndRulesGet200Response) HasConId() bool`

HasConId returns a boolean if a field has been set.

### GetMaturityDate

`func (o *IserverContractConidInfoAndRulesGet200Response) GetMaturityDate() float32`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetMaturityDateOk() (*float32, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *IserverContractConidInfoAndRulesGet200Response) SetMaturityDate(v float32)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *IserverContractConidInfoAndRulesGet200Response) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetIndustry

`func (o *IserverContractConidInfoAndRulesGet200Response) GetIndustry() string`

GetIndustry returns the Industry field if non-nil, zero value otherwise.

### GetIndustryOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetIndustryOk() (*string, bool)`

GetIndustryOk returns a tuple with the Industry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustry

`func (o *IserverContractConidInfoAndRulesGet200Response) SetIndustry(v string)`

SetIndustry sets Industry field to given value.

### HasIndustry

`func (o *IserverContractConidInfoAndRulesGet200Response) HasIndustry() bool`

HasIndustry returns a boolean if a field has been set.

### GetInstrumentType

`func (o *IserverContractConidInfoAndRulesGet200Response) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *IserverContractConidInfoAndRulesGet200Response) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *IserverContractConidInfoAndRulesGet200Response) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetTradingClass

`func (o *IserverContractConidInfoAndRulesGet200Response) GetTradingClass() string`

GetTradingClass returns the TradingClass field if non-nil, zero value otherwise.

### GetTradingClassOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetTradingClassOk() (*string, bool)`

GetTradingClassOk returns a tuple with the TradingClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingClass

`func (o *IserverContractConidInfoAndRulesGet200Response) SetTradingClass(v string)`

SetTradingClass sets TradingClass field to given value.

### HasTradingClass

`func (o *IserverContractConidInfoAndRulesGet200Response) HasTradingClass() bool`

HasTradingClass returns a boolean if a field has been set.

### GetValidExchanges

`func (o *IserverContractConidInfoAndRulesGet200Response) GetValidExchanges() string`

GetValidExchanges returns the ValidExchanges field if non-nil, zero value otherwise.

### GetValidExchangesOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetValidExchangesOk() (*string, bool)`

GetValidExchangesOk returns a tuple with the ValidExchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidExchanges

`func (o *IserverContractConidInfoAndRulesGet200Response) SetValidExchanges(v string)`

SetValidExchanges sets ValidExchanges field to given value.

### HasValidExchanges

`func (o *IserverContractConidInfoAndRulesGet200Response) HasValidExchanges() bool`

HasValidExchanges returns a boolean if a field has been set.

### GetAllowSellLong

`func (o *IserverContractConidInfoAndRulesGet200Response) GetAllowSellLong() bool`

GetAllowSellLong returns the AllowSellLong field if non-nil, zero value otherwise.

### GetAllowSellLongOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetAllowSellLongOk() (*bool, bool)`

GetAllowSellLongOk returns a tuple with the AllowSellLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSellLong

`func (o *IserverContractConidInfoAndRulesGet200Response) SetAllowSellLong(v bool)`

SetAllowSellLong sets AllowSellLong field to given value.

### HasAllowSellLong

`func (o *IserverContractConidInfoAndRulesGet200Response) HasAllowSellLong() bool`

HasAllowSellLong returns a boolean if a field has been set.

### GetIsZeroCommissionSecurity

`func (o *IserverContractConidInfoAndRulesGet200Response) GetIsZeroCommissionSecurity() bool`

GetIsZeroCommissionSecurity returns the IsZeroCommissionSecurity field if non-nil, zero value otherwise.

### GetIsZeroCommissionSecurityOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetIsZeroCommissionSecurityOk() (*bool, bool)`

GetIsZeroCommissionSecurityOk returns a tuple with the IsZeroCommissionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsZeroCommissionSecurity

`func (o *IserverContractConidInfoAndRulesGet200Response) SetIsZeroCommissionSecurity(v bool)`

SetIsZeroCommissionSecurity sets IsZeroCommissionSecurity field to given value.

### HasIsZeroCommissionSecurity

`func (o *IserverContractConidInfoAndRulesGet200Response) HasIsZeroCommissionSecurity() bool`

HasIsZeroCommissionSecurity returns a boolean if a field has been set.

### GetLocalSymbol

`func (o *IserverContractConidInfoAndRulesGet200Response) GetLocalSymbol() string`

GetLocalSymbol returns the LocalSymbol field if non-nil, zero value otherwise.

### GetLocalSymbolOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetLocalSymbolOk() (*string, bool)`

GetLocalSymbolOk returns a tuple with the LocalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSymbol

`func (o *IserverContractConidInfoAndRulesGet200Response) SetLocalSymbol(v string)`

SetLocalSymbol sets LocalSymbol field to given value.

### HasLocalSymbol

`func (o *IserverContractConidInfoAndRulesGet200Response) HasLocalSymbol() bool`

HasLocalSymbol returns a boolean if a field has been set.

### GetClassifier

`func (o *IserverContractConidInfoAndRulesGet200Response) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *IserverContractConidInfoAndRulesGet200Response) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *IserverContractConidInfoAndRulesGet200Response) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetCurrency

`func (o *IserverContractConidInfoAndRulesGet200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *IserverContractConidInfoAndRulesGet200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *IserverContractConidInfoAndRulesGet200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetText

`func (o *IserverContractConidInfoAndRulesGet200Response) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *IserverContractConidInfoAndRulesGet200Response) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *IserverContractConidInfoAndRulesGet200Response) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUnderlyingConId

`func (o *IserverContractConidInfoAndRulesGet200Response) GetUnderlyingConId() float32`

GetUnderlyingConId returns the UnderlyingConId field if non-nil, zero value otherwise.

### GetUnderlyingConIdOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetUnderlyingConIdOk() (*float32, bool)`

GetUnderlyingConIdOk returns a tuple with the UnderlyingConId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingConId

`func (o *IserverContractConidInfoAndRulesGet200Response) SetUnderlyingConId(v float32)`

SetUnderlyingConId sets UnderlyingConId field to given value.

### HasUnderlyingConId

`func (o *IserverContractConidInfoAndRulesGet200Response) HasUnderlyingConId() bool`

HasUnderlyingConId returns a boolean if a field has been set.

### GetRTH

`func (o *IserverContractConidInfoAndRulesGet200Response) GetRTH() bool`

GetRTH returns the RTH field if non-nil, zero value otherwise.

### GetRTHOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetRTHOk() (*bool, bool)`

GetRTHOk returns a tuple with the RTH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRTH

`func (o *IserverContractConidInfoAndRulesGet200Response) SetRTH(v bool)`

SetRTH sets RTH field to given value.

### HasRTH

`func (o *IserverContractConidInfoAndRulesGet200Response) HasRTH() bool`

HasRTH returns a boolean if a field has been set.

### GetMultiplier

`func (o *IserverContractConidInfoAndRulesGet200Response) GetMultiplier() string`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetMultiplierOk() (*string, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *IserverContractConidInfoAndRulesGet200Response) SetMultiplier(v string)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *IserverContractConidInfoAndRulesGet200Response) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetStrike

`func (o *IserverContractConidInfoAndRulesGet200Response) GetStrike() string`

GetStrike returns the Strike field if non-nil, zero value otherwise.

### GetStrikeOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetStrikeOk() (*string, bool)`

GetStrikeOk returns a tuple with the Strike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrike

`func (o *IserverContractConidInfoAndRulesGet200Response) SetStrike(v string)`

SetStrike sets Strike field to given value.

### HasStrike

`func (o *IserverContractConidInfoAndRulesGet200Response) HasStrike() bool`

HasStrike returns a boolean if a field has been set.

### GetRight

`func (o *IserverContractConidInfoAndRulesGet200Response) GetRight() string`

GetRight returns the Right field if non-nil, zero value otherwise.

### GetRightOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetRightOk() (*string, bool)`

GetRightOk returns a tuple with the Right field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRight

`func (o *IserverContractConidInfoAndRulesGet200Response) SetRight(v string)`

SetRight sets Right field to given value.

### HasRight

`func (o *IserverContractConidInfoAndRulesGet200Response) HasRight() bool`

HasRight returns a boolean if a field has been set.

### GetUnderlyingIssuer

`func (o *IserverContractConidInfoAndRulesGet200Response) GetUnderlyingIssuer() string`

GetUnderlyingIssuer returns the UnderlyingIssuer field if non-nil, zero value otherwise.

### GetUnderlyingIssuerOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetUnderlyingIssuerOk() (*string, bool)`

GetUnderlyingIssuerOk returns a tuple with the UnderlyingIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingIssuer

`func (o *IserverContractConidInfoAndRulesGet200Response) SetUnderlyingIssuer(v string)`

SetUnderlyingIssuer sets UnderlyingIssuer field to given value.

### HasUnderlyingIssuer

`func (o *IserverContractConidInfoAndRulesGet200Response) HasUnderlyingIssuer() bool`

HasUnderlyingIssuer returns a boolean if a field has been set.

### GetContractMonth

`func (o *IserverContractConidInfoAndRulesGet200Response) GetContractMonth() string`

GetContractMonth returns the ContractMonth field if non-nil, zero value otherwise.

### GetContractMonthOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetContractMonthOk() (*string, bool)`

GetContractMonthOk returns a tuple with the ContractMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractMonth

`func (o *IserverContractConidInfoAndRulesGet200Response) SetContractMonth(v string)`

SetContractMonth sets ContractMonth field to given value.

### HasContractMonth

`func (o *IserverContractConidInfoAndRulesGet200Response) HasContractMonth() bool`

HasContractMonth returns a boolean if a field has been set.

### GetCompanyName

`func (o *IserverContractConidInfoAndRulesGet200Response) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *IserverContractConidInfoAndRulesGet200Response) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *IserverContractConidInfoAndRulesGet200Response) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetSmartAvailable

`func (o *IserverContractConidInfoAndRulesGet200Response) GetSmartAvailable() bool`

GetSmartAvailable returns the SmartAvailable field if non-nil, zero value otherwise.

### GetSmartAvailableOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetSmartAvailableOk() (*bool, bool)`

GetSmartAvailableOk returns a tuple with the SmartAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartAvailable

`func (o *IserverContractConidInfoAndRulesGet200Response) SetSmartAvailable(v bool)`

SetSmartAvailable sets SmartAvailable field to given value.

### HasSmartAvailable

`func (o *IserverContractConidInfoAndRulesGet200Response) HasSmartAvailable() bool`

HasSmartAvailable returns a boolean if a field has been set.

### GetExchange

`func (o *IserverContractConidInfoAndRulesGet200Response) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *IserverContractConidInfoAndRulesGet200Response) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *IserverContractConidInfoAndRulesGet200Response) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetRules

`func (o *IserverContractConidInfoAndRulesGet200Response) GetRules() []IserverContractConidInfoAndRulesGet200ResponseRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *IserverContractConidInfoAndRulesGet200Response) GetRulesOk() (*[]IserverContractConidInfoAndRulesGet200ResponseRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *IserverContractConidInfoAndRulesGet200Response) SetRules(v []IserverContractConidInfoAndRulesGet200ResponseRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *IserverContractConidInfoAndRulesGet200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


