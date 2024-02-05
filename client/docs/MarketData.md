# MarketData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var31** | Pointer to **string** | Last Price - The last price at which the contract traded. May contain one of the following prefixes:   * C - Previous day&#39;s closing price.   * H - Trading has halted.  | [optional] 
**Var70** | Pointer to **float32** | High - Current day high price | [optional] 
**Var71** | Pointer to **float32** | Low - Current day low price | [optional] 
**Var82** | Pointer to **string** | Change - The difference between the last price and the close on the previous trading day | [optional] 
**Var83** | Pointer to **float32** | Change % - The difference between the last price and the close on the previous trading day in percentage. | [optional] 
**Var84** | Pointer to **string** | Bid Price - The highest-priced bid on the contract. | [optional] 
**Var85** | Pointer to **string** | Ask Size - The number of contracts or shares offered at the ask price. For US stocks, the number displayed is divided by 100. | [optional] 
**Var86** | Pointer to **string** | Ask Price - The lowest-priced offer on the contract. | [optional] 
**Var87** | Pointer to **string** | Volume - Volume for the day, formatted with &#39;K&#39; for thousands or &#39;M&#39; for millions. For higher precision volume refer to field 7762. | [optional] 
**Var88** | Pointer to **string** | Bid Size - The number of contracts or shares bid for at the bid price. For US stocks, the number displayed is divided by 100. | [optional] 
**Var6509** | Pointer to **string** | Market Data Availability. The field may contain three chars. First char defines: R &#x3D; RealTime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed, N &#x3D; Not Subscribed. Second char defines: P &#x3D; Snapshot, p &#x3D; Consolidated. Third char defines: B &#x3D; Book   * RealTime - Data is relayed back in real time without delay, market data subscription(s) are required.   * Delayed - Data is relayed back 15-20 min delayed.   * Frozen - Last recorded data at market close, relayed back in real time.   * Frozen Delayed - Last recorded data at market close, relayed back delayed.   * Not Subscribed - User does not have the required market data subscription(s) to relay back either real time or delayed data.   * Snapshot - Snapshot request is available for contract.   * Consolidated - Market data is aggregated across multiple exchanges or venues.   * Book - Top of the book data is available for contract.  | [optional] 
**Var7057** | Pointer to **string** | Ask Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY | [optional] 
**Var7058** | Pointer to **string** | Last Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY | [optional] 
**Var7059** | Pointer to **float32** | Last Size - The number of unites traded at the last price | [optional] 
**Var7068** | Pointer to **string** | Bid Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY | [optional] 
**Var7195** | Pointer to **string** | IV Rank | [optional] 
**Var7196** | Pointer to **string** | IV Rank | [optional] 
**Var7197** | Pointer to **string** | IV Rank | [optional] 
**Var7198** | Pointer to **string** | IV Percentile | [optional] 
**Var7199** | Pointer to **string** | IV Percentile | [optional] 
**Var7200** | Pointer to **string** | IV Percentile | [optional] 
**Var7201** | Pointer to **string** | IV High Low | [optional] 
**Var7202** | Pointer to **string** | IV High Low | [optional] 
**Var7203** | Pointer to **string** | IV High Low | [optional] 
**Var7204** | Pointer to **string** | IV High Low | [optional] 
**Var7205** | Pointer to **string** | IV High Low | [optional] 
**Var7206** | Pointer to **string** | IV High Low | [optional] 
**Var7207** | Pointer to **string** | HV Rank | [optional] 
**Var7208** | Pointer to **string** | HV Rank | [optional] 
**Var7209** | Pointer to **string** | HV Rank | [optional] 
**Var7210** | Pointer to **string** | HV Percentile | [optional] 
**Var7211** | Pointer to **string** | HV Percentile | [optional] 
**Var7212** | Pointer to **string** | HV Percentile | [optional] 
**Var7245** | Pointer to **string** | HV High Low | [optional] 
**Var7246** | Pointer to **string** | HV High Low | [optional] 
**Var7247** | Pointer to **string** | HV High Low | [optional] 
**Var7248** | Pointer to **string** | HV High Low | [optional] 
**Var7249** | Pointer to **string** | HV High Low | [optional] 
**Var7263** | Pointer to **string** | HV High Low | [optional] 
**Var7264** | Pointer to **string** | ESG | [optional] 
**Var7265** | Pointer to **string** | ESG | [optional] 
**Var7266** | Pointer to **string** | ESG | [optional] 
**Var7267** | Pointer to **string** | ESG | [optional] 
**Var7268** | Pointer to **string** | ESG | [optional] 
**Var7269** | Pointer to **string** | ESG | [optional] 
**Var7271** | Pointer to **string** | ESG | [optional] 
**Var7272** | Pointer to **string** | ESG | [optional] 
**Var7273** | Pointer to **string** | ESG | [optional] 
**Var7274** | Pointer to **string** | ESG | [optional] 
**Var7275** | Pointer to **string** | ESG | [optional] 
**Var7276** | Pointer to **string** | ESG | [optional] 
**Var7277** | Pointer to **string** | ESG | [optional] 
**Var7282** | Pointer to **string** | Average Volume - The average daily trading volume over 90 days. | [optional] 
**Var7283** | Pointer to **string** | Option Implied Vol. % - A prediction of how volatile an underlying will be in the future. At the market volatility estimated for a maturity thirty calendar days forward of the current trading day, and based on option prices from two consecutive expiration months.  | [optional] 
**Var7284** | Pointer to **string** | Historic Volume (30d) | [optional] 
**Var7286** | Pointer to **float32** | Dividend Amount - Displays the amount of the next dividend. | [optional] 
**Var7287** | Pointer to **string** | Dividend Yield % - This value is the toal of the expected dividend payments over the next twelve months per share divided by the Current Price and is expressed as a percentage. For derivatives, this displays the total of the expected dividend payments over the expiry date.  | [optional] 
**Var7288** | Pointer to **string** | Ex-date of the dividend | [optional] 
**Var7289** | Pointer to **string** | Market Cap | [optional] 
**Var7290** | Pointer to **string** | P/E | [optional] 
**Var7293** | Pointer to **string** | 52 Week High - The highest price for the past 52 weeks. | [optional] 
**Var7294** | Pointer to **string** | 52 Week Low - The lowest price for the past 52 weeks. | [optional] 
**Var7295** | Pointer to **float32** | Open - Today&#39;s opening price. | [optional] 
**Var7296** | Pointer to **float32** | Close - Today&#39;s closing price. | [optional] 
**Var7331** | Pointer to **string** | Reuters Fundamentals | [optional] 
**Var7370** | Pointer to **string** | ESG | [optional] 
**Var7371** | Pointer to **string** | ESG | [optional] 
**Var7372** | Pointer to **string** | ESG | [optional] 
**Var7635** | Pointer to **string** | Mark - The mark price is, the ask price if ask is less than last price, the bid price if bid is more than the last price, otherwise it&#39;s equal to last price | [optional] 
**Var7636** | Pointer to **float32** | shortable invetory | [optional] 
**Var7637** | Pointer to **string** | Fee rebate rate | [optional] 
**Var7644** | Pointer to **string** | Shortable - Describes the level of difficulty with which the security can be sold short. | [optional] 
**Var7674** | Pointer to **string** | EMA(200) - Exponential moving average (N&#x3D;200). | [optional] 
**Var7675** | Pointer to **string** | EMA(100) - Exponential moving average (N&#x3D;100). | [optional] 
**Var7676** | Pointer to **string** | EMA(50) - Exponential moving average (N&#x3D;50). | [optional] 
**Var7677** | Pointer to **string** | EMA(20) - Exponential moving average (N&#x3D;20). | [optional] 
**Var7681** | Pointer to **string** | Price/EMA(20) - Price to Exponential moving average (N&#x3D;20) ratio -1, displayed in percents. | [optional] 
**Var7698** | Pointer to **string** | Last Yield - Implied yield of the bond if it is purchased at the current last price. Last yield is calculated using the Last price on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] 
**Var7699** | Pointer to **string** | Bid Yield - Implied yield of the bond if it is purchased at the current bid price. Bid yield is calculated using the Ask on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] 
**Var7718** | Pointer to **string** | Beta - Beta is against standard index. | [optional] 
**Var7720** | Pointer to **string** | Ask Yield - Implied yield of the bond if it is purchased at the current offer. Ask yield is calculated using the Bid on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] 
**Var7743** | Pointer to **string** | Reuters Fundamentals | [optional] 
**Var7761** | Pointer to **string** | ESG | [optional] 
**Var7992** | Pointer to **string** | 26 Week High - The highest price for the past 26 weeks. | [optional] 
**Var7993** | Pointer to **string** | 26 Week Low - The lowest price for the past 26 weeks. | [optional] 
**Var7994** | Pointer to **string** | 13 Week High - The highest price for the past 13 weeks. | [optional] 
**Var7995** | Pointer to **string** | 13 Week Low - The lowest price for the past 13 weeks. | [optional] 
**Conid** | Pointer to **int32** | IBKR Contract identifier | [optional] 
**MinTick** | Pointer to **float32** | minimum price increment | [optional] 
**BboExchange** | Pointer to **string** | Color for Best Bid/Offer Exchange in hex code | [optional] 
**HasDelayed** | Pointer to **bool** | If market data field values return delayed | [optional] 
**SizeMinTick** | Pointer to **int32** | minimum size increment | [optional] 
**BestEligible** | Pointer to **int32** |  | [optional] 
**BestBidExch** | Pointer to **int32** |  | [optional] 
**BestAskExch** | Pointer to **int32** |  | [optional] 
**PreOpenBid** | Pointer to **int32** |  | [optional] 
**LastAttribs** | Pointer to **int32** |  | [optional] 
**TimestampBase** | Pointer to **int32** | Base time stamp for last update in format YYYYMMDD | [optional] 
**TimestampDelta** | Pointer to **int32** |  | [optional] 
**LastExch** | Pointer to **int32** |  | [optional] 
**CloseAttribs** | Pointer to **int32** |  | [optional] 

## Methods

### NewMarketData

`func NewMarketData() *MarketData`

NewMarketData instantiates a new MarketData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketDataWithDefaults

`func NewMarketDataWithDefaults() *MarketData`

NewMarketDataWithDefaults instantiates a new MarketData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar31

`func (o *MarketData) GetVar31() string`

GetVar31 returns the Var31 field if non-nil, zero value otherwise.

### GetVar31Ok

`func (o *MarketData) GetVar31Ok() (*string, bool)`

GetVar31Ok returns a tuple with the Var31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar31

`func (o *MarketData) SetVar31(v string)`

SetVar31 sets Var31 field to given value.

### HasVar31

`func (o *MarketData) HasVar31() bool`

HasVar31 returns a boolean if a field has been set.

### GetVar70

`func (o *MarketData) GetVar70() float32`

GetVar70 returns the Var70 field if non-nil, zero value otherwise.

### GetVar70Ok

`func (o *MarketData) GetVar70Ok() (*float32, bool)`

GetVar70Ok returns a tuple with the Var70 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar70

`func (o *MarketData) SetVar70(v float32)`

SetVar70 sets Var70 field to given value.

### HasVar70

`func (o *MarketData) HasVar70() bool`

HasVar70 returns a boolean if a field has been set.

### GetVar71

`func (o *MarketData) GetVar71() float32`

GetVar71 returns the Var71 field if non-nil, zero value otherwise.

### GetVar71Ok

`func (o *MarketData) GetVar71Ok() (*float32, bool)`

GetVar71Ok returns a tuple with the Var71 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar71

`func (o *MarketData) SetVar71(v float32)`

SetVar71 sets Var71 field to given value.

### HasVar71

`func (o *MarketData) HasVar71() bool`

HasVar71 returns a boolean if a field has been set.

### GetVar82

`func (o *MarketData) GetVar82() string`

GetVar82 returns the Var82 field if non-nil, zero value otherwise.

### GetVar82Ok

`func (o *MarketData) GetVar82Ok() (*string, bool)`

GetVar82Ok returns a tuple with the Var82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar82

`func (o *MarketData) SetVar82(v string)`

SetVar82 sets Var82 field to given value.

### HasVar82

`func (o *MarketData) HasVar82() bool`

HasVar82 returns a boolean if a field has been set.

### GetVar83

`func (o *MarketData) GetVar83() float32`

GetVar83 returns the Var83 field if non-nil, zero value otherwise.

### GetVar83Ok

`func (o *MarketData) GetVar83Ok() (*float32, bool)`

GetVar83Ok returns a tuple with the Var83 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar83

`func (o *MarketData) SetVar83(v float32)`

SetVar83 sets Var83 field to given value.

### HasVar83

`func (o *MarketData) HasVar83() bool`

HasVar83 returns a boolean if a field has been set.

### GetVar84

`func (o *MarketData) GetVar84() string`

GetVar84 returns the Var84 field if non-nil, zero value otherwise.

### GetVar84Ok

`func (o *MarketData) GetVar84Ok() (*string, bool)`

GetVar84Ok returns a tuple with the Var84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar84

`func (o *MarketData) SetVar84(v string)`

SetVar84 sets Var84 field to given value.

### HasVar84

`func (o *MarketData) HasVar84() bool`

HasVar84 returns a boolean if a field has been set.

### GetVar85

`func (o *MarketData) GetVar85() string`

GetVar85 returns the Var85 field if non-nil, zero value otherwise.

### GetVar85Ok

`func (o *MarketData) GetVar85Ok() (*string, bool)`

GetVar85Ok returns a tuple with the Var85 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar85

`func (o *MarketData) SetVar85(v string)`

SetVar85 sets Var85 field to given value.

### HasVar85

`func (o *MarketData) HasVar85() bool`

HasVar85 returns a boolean if a field has been set.

### GetVar86

`func (o *MarketData) GetVar86() string`

GetVar86 returns the Var86 field if non-nil, zero value otherwise.

### GetVar86Ok

`func (o *MarketData) GetVar86Ok() (*string, bool)`

GetVar86Ok returns a tuple with the Var86 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar86

`func (o *MarketData) SetVar86(v string)`

SetVar86 sets Var86 field to given value.

### HasVar86

`func (o *MarketData) HasVar86() bool`

HasVar86 returns a boolean if a field has been set.

### GetVar87

`func (o *MarketData) GetVar87() string`

GetVar87 returns the Var87 field if non-nil, zero value otherwise.

### GetVar87Ok

`func (o *MarketData) GetVar87Ok() (*string, bool)`

GetVar87Ok returns a tuple with the Var87 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar87

`func (o *MarketData) SetVar87(v string)`

SetVar87 sets Var87 field to given value.

### HasVar87

`func (o *MarketData) HasVar87() bool`

HasVar87 returns a boolean if a field has been set.

### GetVar88

`func (o *MarketData) GetVar88() string`

GetVar88 returns the Var88 field if non-nil, zero value otherwise.

### GetVar88Ok

`func (o *MarketData) GetVar88Ok() (*string, bool)`

GetVar88Ok returns a tuple with the Var88 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar88

`func (o *MarketData) SetVar88(v string)`

SetVar88 sets Var88 field to given value.

### HasVar88

`func (o *MarketData) HasVar88() bool`

HasVar88 returns a boolean if a field has been set.

### GetVar6509

`func (o *MarketData) GetVar6509() string`

GetVar6509 returns the Var6509 field if non-nil, zero value otherwise.

### GetVar6509Ok

`func (o *MarketData) GetVar6509Ok() (*string, bool)`

GetVar6509Ok returns a tuple with the Var6509 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6509

`func (o *MarketData) SetVar6509(v string)`

SetVar6509 sets Var6509 field to given value.

### HasVar6509

`func (o *MarketData) HasVar6509() bool`

HasVar6509 returns a boolean if a field has been set.

### GetVar7057

`func (o *MarketData) GetVar7057() string`

GetVar7057 returns the Var7057 field if non-nil, zero value otherwise.

### GetVar7057Ok

`func (o *MarketData) GetVar7057Ok() (*string, bool)`

GetVar7057Ok returns a tuple with the Var7057 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7057

`func (o *MarketData) SetVar7057(v string)`

SetVar7057 sets Var7057 field to given value.

### HasVar7057

`func (o *MarketData) HasVar7057() bool`

HasVar7057 returns a boolean if a field has been set.

### GetVar7058

`func (o *MarketData) GetVar7058() string`

GetVar7058 returns the Var7058 field if non-nil, zero value otherwise.

### GetVar7058Ok

`func (o *MarketData) GetVar7058Ok() (*string, bool)`

GetVar7058Ok returns a tuple with the Var7058 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7058

`func (o *MarketData) SetVar7058(v string)`

SetVar7058 sets Var7058 field to given value.

### HasVar7058

`func (o *MarketData) HasVar7058() bool`

HasVar7058 returns a boolean if a field has been set.

### GetVar7059

`func (o *MarketData) GetVar7059() float32`

GetVar7059 returns the Var7059 field if non-nil, zero value otherwise.

### GetVar7059Ok

`func (o *MarketData) GetVar7059Ok() (*float32, bool)`

GetVar7059Ok returns a tuple with the Var7059 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7059

`func (o *MarketData) SetVar7059(v float32)`

SetVar7059 sets Var7059 field to given value.

### HasVar7059

`func (o *MarketData) HasVar7059() bool`

HasVar7059 returns a boolean if a field has been set.

### GetVar7068

`func (o *MarketData) GetVar7068() string`

GetVar7068 returns the Var7068 field if non-nil, zero value otherwise.

### GetVar7068Ok

`func (o *MarketData) GetVar7068Ok() (*string, bool)`

GetVar7068Ok returns a tuple with the Var7068 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7068

`func (o *MarketData) SetVar7068(v string)`

SetVar7068 sets Var7068 field to given value.

### HasVar7068

`func (o *MarketData) HasVar7068() bool`

HasVar7068 returns a boolean if a field has been set.

### GetVar7195

`func (o *MarketData) GetVar7195() string`

GetVar7195 returns the Var7195 field if non-nil, zero value otherwise.

### GetVar7195Ok

`func (o *MarketData) GetVar7195Ok() (*string, bool)`

GetVar7195Ok returns a tuple with the Var7195 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7195

`func (o *MarketData) SetVar7195(v string)`

SetVar7195 sets Var7195 field to given value.

### HasVar7195

`func (o *MarketData) HasVar7195() bool`

HasVar7195 returns a boolean if a field has been set.

### GetVar7196

`func (o *MarketData) GetVar7196() string`

GetVar7196 returns the Var7196 field if non-nil, zero value otherwise.

### GetVar7196Ok

`func (o *MarketData) GetVar7196Ok() (*string, bool)`

GetVar7196Ok returns a tuple with the Var7196 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7196

`func (o *MarketData) SetVar7196(v string)`

SetVar7196 sets Var7196 field to given value.

### HasVar7196

`func (o *MarketData) HasVar7196() bool`

HasVar7196 returns a boolean if a field has been set.

### GetVar7197

`func (o *MarketData) GetVar7197() string`

GetVar7197 returns the Var7197 field if non-nil, zero value otherwise.

### GetVar7197Ok

`func (o *MarketData) GetVar7197Ok() (*string, bool)`

GetVar7197Ok returns a tuple with the Var7197 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7197

`func (o *MarketData) SetVar7197(v string)`

SetVar7197 sets Var7197 field to given value.

### HasVar7197

`func (o *MarketData) HasVar7197() bool`

HasVar7197 returns a boolean if a field has been set.

### GetVar7198

`func (o *MarketData) GetVar7198() string`

GetVar7198 returns the Var7198 field if non-nil, zero value otherwise.

### GetVar7198Ok

`func (o *MarketData) GetVar7198Ok() (*string, bool)`

GetVar7198Ok returns a tuple with the Var7198 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7198

`func (o *MarketData) SetVar7198(v string)`

SetVar7198 sets Var7198 field to given value.

### HasVar7198

`func (o *MarketData) HasVar7198() bool`

HasVar7198 returns a boolean if a field has been set.

### GetVar7199

`func (o *MarketData) GetVar7199() string`

GetVar7199 returns the Var7199 field if non-nil, zero value otherwise.

### GetVar7199Ok

`func (o *MarketData) GetVar7199Ok() (*string, bool)`

GetVar7199Ok returns a tuple with the Var7199 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7199

`func (o *MarketData) SetVar7199(v string)`

SetVar7199 sets Var7199 field to given value.

### HasVar7199

`func (o *MarketData) HasVar7199() bool`

HasVar7199 returns a boolean if a field has been set.

### GetVar7200

`func (o *MarketData) GetVar7200() string`

GetVar7200 returns the Var7200 field if non-nil, zero value otherwise.

### GetVar7200Ok

`func (o *MarketData) GetVar7200Ok() (*string, bool)`

GetVar7200Ok returns a tuple with the Var7200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7200

`func (o *MarketData) SetVar7200(v string)`

SetVar7200 sets Var7200 field to given value.

### HasVar7200

`func (o *MarketData) HasVar7200() bool`

HasVar7200 returns a boolean if a field has been set.

### GetVar7201

`func (o *MarketData) GetVar7201() string`

GetVar7201 returns the Var7201 field if non-nil, zero value otherwise.

### GetVar7201Ok

`func (o *MarketData) GetVar7201Ok() (*string, bool)`

GetVar7201Ok returns a tuple with the Var7201 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7201

`func (o *MarketData) SetVar7201(v string)`

SetVar7201 sets Var7201 field to given value.

### HasVar7201

`func (o *MarketData) HasVar7201() bool`

HasVar7201 returns a boolean if a field has been set.

### GetVar7202

`func (o *MarketData) GetVar7202() string`

GetVar7202 returns the Var7202 field if non-nil, zero value otherwise.

### GetVar7202Ok

`func (o *MarketData) GetVar7202Ok() (*string, bool)`

GetVar7202Ok returns a tuple with the Var7202 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7202

`func (o *MarketData) SetVar7202(v string)`

SetVar7202 sets Var7202 field to given value.

### HasVar7202

`func (o *MarketData) HasVar7202() bool`

HasVar7202 returns a boolean if a field has been set.

### GetVar7203

`func (o *MarketData) GetVar7203() string`

GetVar7203 returns the Var7203 field if non-nil, zero value otherwise.

### GetVar7203Ok

`func (o *MarketData) GetVar7203Ok() (*string, bool)`

GetVar7203Ok returns a tuple with the Var7203 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7203

`func (o *MarketData) SetVar7203(v string)`

SetVar7203 sets Var7203 field to given value.

### HasVar7203

`func (o *MarketData) HasVar7203() bool`

HasVar7203 returns a boolean if a field has been set.

### GetVar7204

`func (o *MarketData) GetVar7204() string`

GetVar7204 returns the Var7204 field if non-nil, zero value otherwise.

### GetVar7204Ok

`func (o *MarketData) GetVar7204Ok() (*string, bool)`

GetVar7204Ok returns a tuple with the Var7204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7204

`func (o *MarketData) SetVar7204(v string)`

SetVar7204 sets Var7204 field to given value.

### HasVar7204

`func (o *MarketData) HasVar7204() bool`

HasVar7204 returns a boolean if a field has been set.

### GetVar7205

`func (o *MarketData) GetVar7205() string`

GetVar7205 returns the Var7205 field if non-nil, zero value otherwise.

### GetVar7205Ok

`func (o *MarketData) GetVar7205Ok() (*string, bool)`

GetVar7205Ok returns a tuple with the Var7205 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7205

`func (o *MarketData) SetVar7205(v string)`

SetVar7205 sets Var7205 field to given value.

### HasVar7205

`func (o *MarketData) HasVar7205() bool`

HasVar7205 returns a boolean if a field has been set.

### GetVar7206

`func (o *MarketData) GetVar7206() string`

GetVar7206 returns the Var7206 field if non-nil, zero value otherwise.

### GetVar7206Ok

`func (o *MarketData) GetVar7206Ok() (*string, bool)`

GetVar7206Ok returns a tuple with the Var7206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7206

`func (o *MarketData) SetVar7206(v string)`

SetVar7206 sets Var7206 field to given value.

### HasVar7206

`func (o *MarketData) HasVar7206() bool`

HasVar7206 returns a boolean if a field has been set.

### GetVar7207

`func (o *MarketData) GetVar7207() string`

GetVar7207 returns the Var7207 field if non-nil, zero value otherwise.

### GetVar7207Ok

`func (o *MarketData) GetVar7207Ok() (*string, bool)`

GetVar7207Ok returns a tuple with the Var7207 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7207

`func (o *MarketData) SetVar7207(v string)`

SetVar7207 sets Var7207 field to given value.

### HasVar7207

`func (o *MarketData) HasVar7207() bool`

HasVar7207 returns a boolean if a field has been set.

### GetVar7208

`func (o *MarketData) GetVar7208() string`

GetVar7208 returns the Var7208 field if non-nil, zero value otherwise.

### GetVar7208Ok

`func (o *MarketData) GetVar7208Ok() (*string, bool)`

GetVar7208Ok returns a tuple with the Var7208 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7208

`func (o *MarketData) SetVar7208(v string)`

SetVar7208 sets Var7208 field to given value.

### HasVar7208

`func (o *MarketData) HasVar7208() bool`

HasVar7208 returns a boolean if a field has been set.

### GetVar7209

`func (o *MarketData) GetVar7209() string`

GetVar7209 returns the Var7209 field if non-nil, zero value otherwise.

### GetVar7209Ok

`func (o *MarketData) GetVar7209Ok() (*string, bool)`

GetVar7209Ok returns a tuple with the Var7209 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7209

`func (o *MarketData) SetVar7209(v string)`

SetVar7209 sets Var7209 field to given value.

### HasVar7209

`func (o *MarketData) HasVar7209() bool`

HasVar7209 returns a boolean if a field has been set.

### GetVar7210

`func (o *MarketData) GetVar7210() string`

GetVar7210 returns the Var7210 field if non-nil, zero value otherwise.

### GetVar7210Ok

`func (o *MarketData) GetVar7210Ok() (*string, bool)`

GetVar7210Ok returns a tuple with the Var7210 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7210

`func (o *MarketData) SetVar7210(v string)`

SetVar7210 sets Var7210 field to given value.

### HasVar7210

`func (o *MarketData) HasVar7210() bool`

HasVar7210 returns a boolean if a field has been set.

### GetVar7211

`func (o *MarketData) GetVar7211() string`

GetVar7211 returns the Var7211 field if non-nil, zero value otherwise.

### GetVar7211Ok

`func (o *MarketData) GetVar7211Ok() (*string, bool)`

GetVar7211Ok returns a tuple with the Var7211 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7211

`func (o *MarketData) SetVar7211(v string)`

SetVar7211 sets Var7211 field to given value.

### HasVar7211

`func (o *MarketData) HasVar7211() bool`

HasVar7211 returns a boolean if a field has been set.

### GetVar7212

`func (o *MarketData) GetVar7212() string`

GetVar7212 returns the Var7212 field if non-nil, zero value otherwise.

### GetVar7212Ok

`func (o *MarketData) GetVar7212Ok() (*string, bool)`

GetVar7212Ok returns a tuple with the Var7212 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7212

`func (o *MarketData) SetVar7212(v string)`

SetVar7212 sets Var7212 field to given value.

### HasVar7212

`func (o *MarketData) HasVar7212() bool`

HasVar7212 returns a boolean if a field has been set.

### GetVar7245

`func (o *MarketData) GetVar7245() string`

GetVar7245 returns the Var7245 field if non-nil, zero value otherwise.

### GetVar7245Ok

`func (o *MarketData) GetVar7245Ok() (*string, bool)`

GetVar7245Ok returns a tuple with the Var7245 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7245

`func (o *MarketData) SetVar7245(v string)`

SetVar7245 sets Var7245 field to given value.

### HasVar7245

`func (o *MarketData) HasVar7245() bool`

HasVar7245 returns a boolean if a field has been set.

### GetVar7246

`func (o *MarketData) GetVar7246() string`

GetVar7246 returns the Var7246 field if non-nil, zero value otherwise.

### GetVar7246Ok

`func (o *MarketData) GetVar7246Ok() (*string, bool)`

GetVar7246Ok returns a tuple with the Var7246 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7246

`func (o *MarketData) SetVar7246(v string)`

SetVar7246 sets Var7246 field to given value.

### HasVar7246

`func (o *MarketData) HasVar7246() bool`

HasVar7246 returns a boolean if a field has been set.

### GetVar7247

`func (o *MarketData) GetVar7247() string`

GetVar7247 returns the Var7247 field if non-nil, zero value otherwise.

### GetVar7247Ok

`func (o *MarketData) GetVar7247Ok() (*string, bool)`

GetVar7247Ok returns a tuple with the Var7247 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7247

`func (o *MarketData) SetVar7247(v string)`

SetVar7247 sets Var7247 field to given value.

### HasVar7247

`func (o *MarketData) HasVar7247() bool`

HasVar7247 returns a boolean if a field has been set.

### GetVar7248

`func (o *MarketData) GetVar7248() string`

GetVar7248 returns the Var7248 field if non-nil, zero value otherwise.

### GetVar7248Ok

`func (o *MarketData) GetVar7248Ok() (*string, bool)`

GetVar7248Ok returns a tuple with the Var7248 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7248

`func (o *MarketData) SetVar7248(v string)`

SetVar7248 sets Var7248 field to given value.

### HasVar7248

`func (o *MarketData) HasVar7248() bool`

HasVar7248 returns a boolean if a field has been set.

### GetVar7249

`func (o *MarketData) GetVar7249() string`

GetVar7249 returns the Var7249 field if non-nil, zero value otherwise.

### GetVar7249Ok

`func (o *MarketData) GetVar7249Ok() (*string, bool)`

GetVar7249Ok returns a tuple with the Var7249 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7249

`func (o *MarketData) SetVar7249(v string)`

SetVar7249 sets Var7249 field to given value.

### HasVar7249

`func (o *MarketData) HasVar7249() bool`

HasVar7249 returns a boolean if a field has been set.

### GetVar7263

`func (o *MarketData) GetVar7263() string`

GetVar7263 returns the Var7263 field if non-nil, zero value otherwise.

### GetVar7263Ok

`func (o *MarketData) GetVar7263Ok() (*string, bool)`

GetVar7263Ok returns a tuple with the Var7263 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7263

`func (o *MarketData) SetVar7263(v string)`

SetVar7263 sets Var7263 field to given value.

### HasVar7263

`func (o *MarketData) HasVar7263() bool`

HasVar7263 returns a boolean if a field has been set.

### GetVar7264

`func (o *MarketData) GetVar7264() string`

GetVar7264 returns the Var7264 field if non-nil, zero value otherwise.

### GetVar7264Ok

`func (o *MarketData) GetVar7264Ok() (*string, bool)`

GetVar7264Ok returns a tuple with the Var7264 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7264

`func (o *MarketData) SetVar7264(v string)`

SetVar7264 sets Var7264 field to given value.

### HasVar7264

`func (o *MarketData) HasVar7264() bool`

HasVar7264 returns a boolean if a field has been set.

### GetVar7265

`func (o *MarketData) GetVar7265() string`

GetVar7265 returns the Var7265 field if non-nil, zero value otherwise.

### GetVar7265Ok

`func (o *MarketData) GetVar7265Ok() (*string, bool)`

GetVar7265Ok returns a tuple with the Var7265 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7265

`func (o *MarketData) SetVar7265(v string)`

SetVar7265 sets Var7265 field to given value.

### HasVar7265

`func (o *MarketData) HasVar7265() bool`

HasVar7265 returns a boolean if a field has been set.

### GetVar7266

`func (o *MarketData) GetVar7266() string`

GetVar7266 returns the Var7266 field if non-nil, zero value otherwise.

### GetVar7266Ok

`func (o *MarketData) GetVar7266Ok() (*string, bool)`

GetVar7266Ok returns a tuple with the Var7266 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7266

`func (o *MarketData) SetVar7266(v string)`

SetVar7266 sets Var7266 field to given value.

### HasVar7266

`func (o *MarketData) HasVar7266() bool`

HasVar7266 returns a boolean if a field has been set.

### GetVar7267

`func (o *MarketData) GetVar7267() string`

GetVar7267 returns the Var7267 field if non-nil, zero value otherwise.

### GetVar7267Ok

`func (o *MarketData) GetVar7267Ok() (*string, bool)`

GetVar7267Ok returns a tuple with the Var7267 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7267

`func (o *MarketData) SetVar7267(v string)`

SetVar7267 sets Var7267 field to given value.

### HasVar7267

`func (o *MarketData) HasVar7267() bool`

HasVar7267 returns a boolean if a field has been set.

### GetVar7268

`func (o *MarketData) GetVar7268() string`

GetVar7268 returns the Var7268 field if non-nil, zero value otherwise.

### GetVar7268Ok

`func (o *MarketData) GetVar7268Ok() (*string, bool)`

GetVar7268Ok returns a tuple with the Var7268 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7268

`func (o *MarketData) SetVar7268(v string)`

SetVar7268 sets Var7268 field to given value.

### HasVar7268

`func (o *MarketData) HasVar7268() bool`

HasVar7268 returns a boolean if a field has been set.

### GetVar7269

`func (o *MarketData) GetVar7269() string`

GetVar7269 returns the Var7269 field if non-nil, zero value otherwise.

### GetVar7269Ok

`func (o *MarketData) GetVar7269Ok() (*string, bool)`

GetVar7269Ok returns a tuple with the Var7269 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7269

`func (o *MarketData) SetVar7269(v string)`

SetVar7269 sets Var7269 field to given value.

### HasVar7269

`func (o *MarketData) HasVar7269() bool`

HasVar7269 returns a boolean if a field has been set.

### GetVar7271

`func (o *MarketData) GetVar7271() string`

GetVar7271 returns the Var7271 field if non-nil, zero value otherwise.

### GetVar7271Ok

`func (o *MarketData) GetVar7271Ok() (*string, bool)`

GetVar7271Ok returns a tuple with the Var7271 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7271

`func (o *MarketData) SetVar7271(v string)`

SetVar7271 sets Var7271 field to given value.

### HasVar7271

`func (o *MarketData) HasVar7271() bool`

HasVar7271 returns a boolean if a field has been set.

### GetVar7272

`func (o *MarketData) GetVar7272() string`

GetVar7272 returns the Var7272 field if non-nil, zero value otherwise.

### GetVar7272Ok

`func (o *MarketData) GetVar7272Ok() (*string, bool)`

GetVar7272Ok returns a tuple with the Var7272 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7272

`func (o *MarketData) SetVar7272(v string)`

SetVar7272 sets Var7272 field to given value.

### HasVar7272

`func (o *MarketData) HasVar7272() bool`

HasVar7272 returns a boolean if a field has been set.

### GetVar7273

`func (o *MarketData) GetVar7273() string`

GetVar7273 returns the Var7273 field if non-nil, zero value otherwise.

### GetVar7273Ok

`func (o *MarketData) GetVar7273Ok() (*string, bool)`

GetVar7273Ok returns a tuple with the Var7273 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7273

`func (o *MarketData) SetVar7273(v string)`

SetVar7273 sets Var7273 field to given value.

### HasVar7273

`func (o *MarketData) HasVar7273() bool`

HasVar7273 returns a boolean if a field has been set.

### GetVar7274

`func (o *MarketData) GetVar7274() string`

GetVar7274 returns the Var7274 field if non-nil, zero value otherwise.

### GetVar7274Ok

`func (o *MarketData) GetVar7274Ok() (*string, bool)`

GetVar7274Ok returns a tuple with the Var7274 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7274

`func (o *MarketData) SetVar7274(v string)`

SetVar7274 sets Var7274 field to given value.

### HasVar7274

`func (o *MarketData) HasVar7274() bool`

HasVar7274 returns a boolean if a field has been set.

### GetVar7275

`func (o *MarketData) GetVar7275() string`

GetVar7275 returns the Var7275 field if non-nil, zero value otherwise.

### GetVar7275Ok

`func (o *MarketData) GetVar7275Ok() (*string, bool)`

GetVar7275Ok returns a tuple with the Var7275 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7275

`func (o *MarketData) SetVar7275(v string)`

SetVar7275 sets Var7275 field to given value.

### HasVar7275

`func (o *MarketData) HasVar7275() bool`

HasVar7275 returns a boolean if a field has been set.

### GetVar7276

`func (o *MarketData) GetVar7276() string`

GetVar7276 returns the Var7276 field if non-nil, zero value otherwise.

### GetVar7276Ok

`func (o *MarketData) GetVar7276Ok() (*string, bool)`

GetVar7276Ok returns a tuple with the Var7276 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7276

`func (o *MarketData) SetVar7276(v string)`

SetVar7276 sets Var7276 field to given value.

### HasVar7276

`func (o *MarketData) HasVar7276() bool`

HasVar7276 returns a boolean if a field has been set.

### GetVar7277

`func (o *MarketData) GetVar7277() string`

GetVar7277 returns the Var7277 field if non-nil, zero value otherwise.

### GetVar7277Ok

`func (o *MarketData) GetVar7277Ok() (*string, bool)`

GetVar7277Ok returns a tuple with the Var7277 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7277

`func (o *MarketData) SetVar7277(v string)`

SetVar7277 sets Var7277 field to given value.

### HasVar7277

`func (o *MarketData) HasVar7277() bool`

HasVar7277 returns a boolean if a field has been set.

### GetVar7282

`func (o *MarketData) GetVar7282() string`

GetVar7282 returns the Var7282 field if non-nil, zero value otherwise.

### GetVar7282Ok

`func (o *MarketData) GetVar7282Ok() (*string, bool)`

GetVar7282Ok returns a tuple with the Var7282 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7282

`func (o *MarketData) SetVar7282(v string)`

SetVar7282 sets Var7282 field to given value.

### HasVar7282

`func (o *MarketData) HasVar7282() bool`

HasVar7282 returns a boolean if a field has been set.

### GetVar7283

`func (o *MarketData) GetVar7283() string`

GetVar7283 returns the Var7283 field if non-nil, zero value otherwise.

### GetVar7283Ok

`func (o *MarketData) GetVar7283Ok() (*string, bool)`

GetVar7283Ok returns a tuple with the Var7283 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7283

`func (o *MarketData) SetVar7283(v string)`

SetVar7283 sets Var7283 field to given value.

### HasVar7283

`func (o *MarketData) HasVar7283() bool`

HasVar7283 returns a boolean if a field has been set.

### GetVar7284

`func (o *MarketData) GetVar7284() string`

GetVar7284 returns the Var7284 field if non-nil, zero value otherwise.

### GetVar7284Ok

`func (o *MarketData) GetVar7284Ok() (*string, bool)`

GetVar7284Ok returns a tuple with the Var7284 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7284

`func (o *MarketData) SetVar7284(v string)`

SetVar7284 sets Var7284 field to given value.

### HasVar7284

`func (o *MarketData) HasVar7284() bool`

HasVar7284 returns a boolean if a field has been set.

### GetVar7286

`func (o *MarketData) GetVar7286() float32`

GetVar7286 returns the Var7286 field if non-nil, zero value otherwise.

### GetVar7286Ok

`func (o *MarketData) GetVar7286Ok() (*float32, bool)`

GetVar7286Ok returns a tuple with the Var7286 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7286

`func (o *MarketData) SetVar7286(v float32)`

SetVar7286 sets Var7286 field to given value.

### HasVar7286

`func (o *MarketData) HasVar7286() bool`

HasVar7286 returns a boolean if a field has been set.

### GetVar7287

`func (o *MarketData) GetVar7287() string`

GetVar7287 returns the Var7287 field if non-nil, zero value otherwise.

### GetVar7287Ok

`func (o *MarketData) GetVar7287Ok() (*string, bool)`

GetVar7287Ok returns a tuple with the Var7287 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7287

`func (o *MarketData) SetVar7287(v string)`

SetVar7287 sets Var7287 field to given value.

### HasVar7287

`func (o *MarketData) HasVar7287() bool`

HasVar7287 returns a boolean if a field has been set.

### GetVar7288

`func (o *MarketData) GetVar7288() string`

GetVar7288 returns the Var7288 field if non-nil, zero value otherwise.

### GetVar7288Ok

`func (o *MarketData) GetVar7288Ok() (*string, bool)`

GetVar7288Ok returns a tuple with the Var7288 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7288

`func (o *MarketData) SetVar7288(v string)`

SetVar7288 sets Var7288 field to given value.

### HasVar7288

`func (o *MarketData) HasVar7288() bool`

HasVar7288 returns a boolean if a field has been set.

### GetVar7289

`func (o *MarketData) GetVar7289() string`

GetVar7289 returns the Var7289 field if non-nil, zero value otherwise.

### GetVar7289Ok

`func (o *MarketData) GetVar7289Ok() (*string, bool)`

GetVar7289Ok returns a tuple with the Var7289 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7289

`func (o *MarketData) SetVar7289(v string)`

SetVar7289 sets Var7289 field to given value.

### HasVar7289

`func (o *MarketData) HasVar7289() bool`

HasVar7289 returns a boolean if a field has been set.

### GetVar7290

`func (o *MarketData) GetVar7290() string`

GetVar7290 returns the Var7290 field if non-nil, zero value otherwise.

### GetVar7290Ok

`func (o *MarketData) GetVar7290Ok() (*string, bool)`

GetVar7290Ok returns a tuple with the Var7290 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7290

`func (o *MarketData) SetVar7290(v string)`

SetVar7290 sets Var7290 field to given value.

### HasVar7290

`func (o *MarketData) HasVar7290() bool`

HasVar7290 returns a boolean if a field has been set.

### GetVar7293

`func (o *MarketData) GetVar7293() string`

GetVar7293 returns the Var7293 field if non-nil, zero value otherwise.

### GetVar7293Ok

`func (o *MarketData) GetVar7293Ok() (*string, bool)`

GetVar7293Ok returns a tuple with the Var7293 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7293

`func (o *MarketData) SetVar7293(v string)`

SetVar7293 sets Var7293 field to given value.

### HasVar7293

`func (o *MarketData) HasVar7293() bool`

HasVar7293 returns a boolean if a field has been set.

### GetVar7294

`func (o *MarketData) GetVar7294() string`

GetVar7294 returns the Var7294 field if non-nil, zero value otherwise.

### GetVar7294Ok

`func (o *MarketData) GetVar7294Ok() (*string, bool)`

GetVar7294Ok returns a tuple with the Var7294 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7294

`func (o *MarketData) SetVar7294(v string)`

SetVar7294 sets Var7294 field to given value.

### HasVar7294

`func (o *MarketData) HasVar7294() bool`

HasVar7294 returns a boolean if a field has been set.

### GetVar7295

`func (o *MarketData) GetVar7295() float32`

GetVar7295 returns the Var7295 field if non-nil, zero value otherwise.

### GetVar7295Ok

`func (o *MarketData) GetVar7295Ok() (*float32, bool)`

GetVar7295Ok returns a tuple with the Var7295 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7295

`func (o *MarketData) SetVar7295(v float32)`

SetVar7295 sets Var7295 field to given value.

### HasVar7295

`func (o *MarketData) HasVar7295() bool`

HasVar7295 returns a boolean if a field has been set.

### GetVar7296

`func (o *MarketData) GetVar7296() float32`

GetVar7296 returns the Var7296 field if non-nil, zero value otherwise.

### GetVar7296Ok

`func (o *MarketData) GetVar7296Ok() (*float32, bool)`

GetVar7296Ok returns a tuple with the Var7296 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7296

`func (o *MarketData) SetVar7296(v float32)`

SetVar7296 sets Var7296 field to given value.

### HasVar7296

`func (o *MarketData) HasVar7296() bool`

HasVar7296 returns a boolean if a field has been set.

### GetVar7331

`func (o *MarketData) GetVar7331() string`

GetVar7331 returns the Var7331 field if non-nil, zero value otherwise.

### GetVar7331Ok

`func (o *MarketData) GetVar7331Ok() (*string, bool)`

GetVar7331Ok returns a tuple with the Var7331 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7331

`func (o *MarketData) SetVar7331(v string)`

SetVar7331 sets Var7331 field to given value.

### HasVar7331

`func (o *MarketData) HasVar7331() bool`

HasVar7331 returns a boolean if a field has been set.

### GetVar7370

`func (o *MarketData) GetVar7370() string`

GetVar7370 returns the Var7370 field if non-nil, zero value otherwise.

### GetVar7370Ok

`func (o *MarketData) GetVar7370Ok() (*string, bool)`

GetVar7370Ok returns a tuple with the Var7370 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7370

`func (o *MarketData) SetVar7370(v string)`

SetVar7370 sets Var7370 field to given value.

### HasVar7370

`func (o *MarketData) HasVar7370() bool`

HasVar7370 returns a boolean if a field has been set.

### GetVar7371

`func (o *MarketData) GetVar7371() string`

GetVar7371 returns the Var7371 field if non-nil, zero value otherwise.

### GetVar7371Ok

`func (o *MarketData) GetVar7371Ok() (*string, bool)`

GetVar7371Ok returns a tuple with the Var7371 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7371

`func (o *MarketData) SetVar7371(v string)`

SetVar7371 sets Var7371 field to given value.

### HasVar7371

`func (o *MarketData) HasVar7371() bool`

HasVar7371 returns a boolean if a field has been set.

### GetVar7372

`func (o *MarketData) GetVar7372() string`

GetVar7372 returns the Var7372 field if non-nil, zero value otherwise.

### GetVar7372Ok

`func (o *MarketData) GetVar7372Ok() (*string, bool)`

GetVar7372Ok returns a tuple with the Var7372 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7372

`func (o *MarketData) SetVar7372(v string)`

SetVar7372 sets Var7372 field to given value.

### HasVar7372

`func (o *MarketData) HasVar7372() bool`

HasVar7372 returns a boolean if a field has been set.

### GetVar7635

`func (o *MarketData) GetVar7635() string`

GetVar7635 returns the Var7635 field if non-nil, zero value otherwise.

### GetVar7635Ok

`func (o *MarketData) GetVar7635Ok() (*string, bool)`

GetVar7635Ok returns a tuple with the Var7635 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7635

`func (o *MarketData) SetVar7635(v string)`

SetVar7635 sets Var7635 field to given value.

### HasVar7635

`func (o *MarketData) HasVar7635() bool`

HasVar7635 returns a boolean if a field has been set.

### GetVar7636

`func (o *MarketData) GetVar7636() float32`

GetVar7636 returns the Var7636 field if non-nil, zero value otherwise.

### GetVar7636Ok

`func (o *MarketData) GetVar7636Ok() (*float32, bool)`

GetVar7636Ok returns a tuple with the Var7636 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7636

`func (o *MarketData) SetVar7636(v float32)`

SetVar7636 sets Var7636 field to given value.

### HasVar7636

`func (o *MarketData) HasVar7636() bool`

HasVar7636 returns a boolean if a field has been set.

### GetVar7637

`func (o *MarketData) GetVar7637() string`

GetVar7637 returns the Var7637 field if non-nil, zero value otherwise.

### GetVar7637Ok

`func (o *MarketData) GetVar7637Ok() (*string, bool)`

GetVar7637Ok returns a tuple with the Var7637 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7637

`func (o *MarketData) SetVar7637(v string)`

SetVar7637 sets Var7637 field to given value.

### HasVar7637

`func (o *MarketData) HasVar7637() bool`

HasVar7637 returns a boolean if a field has been set.

### GetVar7644

`func (o *MarketData) GetVar7644() string`

GetVar7644 returns the Var7644 field if non-nil, zero value otherwise.

### GetVar7644Ok

`func (o *MarketData) GetVar7644Ok() (*string, bool)`

GetVar7644Ok returns a tuple with the Var7644 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7644

`func (o *MarketData) SetVar7644(v string)`

SetVar7644 sets Var7644 field to given value.

### HasVar7644

`func (o *MarketData) HasVar7644() bool`

HasVar7644 returns a boolean if a field has been set.

### GetVar7674

`func (o *MarketData) GetVar7674() string`

GetVar7674 returns the Var7674 field if non-nil, zero value otherwise.

### GetVar7674Ok

`func (o *MarketData) GetVar7674Ok() (*string, bool)`

GetVar7674Ok returns a tuple with the Var7674 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7674

`func (o *MarketData) SetVar7674(v string)`

SetVar7674 sets Var7674 field to given value.

### HasVar7674

`func (o *MarketData) HasVar7674() bool`

HasVar7674 returns a boolean if a field has been set.

### GetVar7675

`func (o *MarketData) GetVar7675() string`

GetVar7675 returns the Var7675 field if non-nil, zero value otherwise.

### GetVar7675Ok

`func (o *MarketData) GetVar7675Ok() (*string, bool)`

GetVar7675Ok returns a tuple with the Var7675 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7675

`func (o *MarketData) SetVar7675(v string)`

SetVar7675 sets Var7675 field to given value.

### HasVar7675

`func (o *MarketData) HasVar7675() bool`

HasVar7675 returns a boolean if a field has been set.

### GetVar7676

`func (o *MarketData) GetVar7676() string`

GetVar7676 returns the Var7676 field if non-nil, zero value otherwise.

### GetVar7676Ok

`func (o *MarketData) GetVar7676Ok() (*string, bool)`

GetVar7676Ok returns a tuple with the Var7676 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7676

`func (o *MarketData) SetVar7676(v string)`

SetVar7676 sets Var7676 field to given value.

### HasVar7676

`func (o *MarketData) HasVar7676() bool`

HasVar7676 returns a boolean if a field has been set.

### GetVar7677

`func (o *MarketData) GetVar7677() string`

GetVar7677 returns the Var7677 field if non-nil, zero value otherwise.

### GetVar7677Ok

`func (o *MarketData) GetVar7677Ok() (*string, bool)`

GetVar7677Ok returns a tuple with the Var7677 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7677

`func (o *MarketData) SetVar7677(v string)`

SetVar7677 sets Var7677 field to given value.

### HasVar7677

`func (o *MarketData) HasVar7677() bool`

HasVar7677 returns a boolean if a field has been set.

### GetVar7681

`func (o *MarketData) GetVar7681() string`

GetVar7681 returns the Var7681 field if non-nil, zero value otherwise.

### GetVar7681Ok

`func (o *MarketData) GetVar7681Ok() (*string, bool)`

GetVar7681Ok returns a tuple with the Var7681 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7681

`func (o *MarketData) SetVar7681(v string)`

SetVar7681 sets Var7681 field to given value.

### HasVar7681

`func (o *MarketData) HasVar7681() bool`

HasVar7681 returns a boolean if a field has been set.

### GetVar7698

`func (o *MarketData) GetVar7698() string`

GetVar7698 returns the Var7698 field if non-nil, zero value otherwise.

### GetVar7698Ok

`func (o *MarketData) GetVar7698Ok() (*string, bool)`

GetVar7698Ok returns a tuple with the Var7698 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7698

`func (o *MarketData) SetVar7698(v string)`

SetVar7698 sets Var7698 field to given value.

### HasVar7698

`func (o *MarketData) HasVar7698() bool`

HasVar7698 returns a boolean if a field has been set.

### GetVar7699

`func (o *MarketData) GetVar7699() string`

GetVar7699 returns the Var7699 field if non-nil, zero value otherwise.

### GetVar7699Ok

`func (o *MarketData) GetVar7699Ok() (*string, bool)`

GetVar7699Ok returns a tuple with the Var7699 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7699

`func (o *MarketData) SetVar7699(v string)`

SetVar7699 sets Var7699 field to given value.

### HasVar7699

`func (o *MarketData) HasVar7699() bool`

HasVar7699 returns a boolean if a field has been set.

### GetVar7718

`func (o *MarketData) GetVar7718() string`

GetVar7718 returns the Var7718 field if non-nil, zero value otherwise.

### GetVar7718Ok

`func (o *MarketData) GetVar7718Ok() (*string, bool)`

GetVar7718Ok returns a tuple with the Var7718 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7718

`func (o *MarketData) SetVar7718(v string)`

SetVar7718 sets Var7718 field to given value.

### HasVar7718

`func (o *MarketData) HasVar7718() bool`

HasVar7718 returns a boolean if a field has been set.

### GetVar7720

`func (o *MarketData) GetVar7720() string`

GetVar7720 returns the Var7720 field if non-nil, zero value otherwise.

### GetVar7720Ok

`func (o *MarketData) GetVar7720Ok() (*string, bool)`

GetVar7720Ok returns a tuple with the Var7720 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7720

`func (o *MarketData) SetVar7720(v string)`

SetVar7720 sets Var7720 field to given value.

### HasVar7720

`func (o *MarketData) HasVar7720() bool`

HasVar7720 returns a boolean if a field has been set.

### GetVar7743

`func (o *MarketData) GetVar7743() string`

GetVar7743 returns the Var7743 field if non-nil, zero value otherwise.

### GetVar7743Ok

`func (o *MarketData) GetVar7743Ok() (*string, bool)`

GetVar7743Ok returns a tuple with the Var7743 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7743

`func (o *MarketData) SetVar7743(v string)`

SetVar7743 sets Var7743 field to given value.

### HasVar7743

`func (o *MarketData) HasVar7743() bool`

HasVar7743 returns a boolean if a field has been set.

### GetVar7761

`func (o *MarketData) GetVar7761() string`

GetVar7761 returns the Var7761 field if non-nil, zero value otherwise.

### GetVar7761Ok

`func (o *MarketData) GetVar7761Ok() (*string, bool)`

GetVar7761Ok returns a tuple with the Var7761 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7761

`func (o *MarketData) SetVar7761(v string)`

SetVar7761 sets Var7761 field to given value.

### HasVar7761

`func (o *MarketData) HasVar7761() bool`

HasVar7761 returns a boolean if a field has been set.

### GetVar7992

`func (o *MarketData) GetVar7992() string`

GetVar7992 returns the Var7992 field if non-nil, zero value otherwise.

### GetVar7992Ok

`func (o *MarketData) GetVar7992Ok() (*string, bool)`

GetVar7992Ok returns a tuple with the Var7992 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7992

`func (o *MarketData) SetVar7992(v string)`

SetVar7992 sets Var7992 field to given value.

### HasVar7992

`func (o *MarketData) HasVar7992() bool`

HasVar7992 returns a boolean if a field has been set.

### GetVar7993

`func (o *MarketData) GetVar7993() string`

GetVar7993 returns the Var7993 field if non-nil, zero value otherwise.

### GetVar7993Ok

`func (o *MarketData) GetVar7993Ok() (*string, bool)`

GetVar7993Ok returns a tuple with the Var7993 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7993

`func (o *MarketData) SetVar7993(v string)`

SetVar7993 sets Var7993 field to given value.

### HasVar7993

`func (o *MarketData) HasVar7993() bool`

HasVar7993 returns a boolean if a field has been set.

### GetVar7994

`func (o *MarketData) GetVar7994() string`

GetVar7994 returns the Var7994 field if non-nil, zero value otherwise.

### GetVar7994Ok

`func (o *MarketData) GetVar7994Ok() (*string, bool)`

GetVar7994Ok returns a tuple with the Var7994 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7994

`func (o *MarketData) SetVar7994(v string)`

SetVar7994 sets Var7994 field to given value.

### HasVar7994

`func (o *MarketData) HasVar7994() bool`

HasVar7994 returns a boolean if a field has been set.

### GetVar7995

`func (o *MarketData) GetVar7995() string`

GetVar7995 returns the Var7995 field if non-nil, zero value otherwise.

### GetVar7995Ok

`func (o *MarketData) GetVar7995Ok() (*string, bool)`

GetVar7995Ok returns a tuple with the Var7995 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7995

`func (o *MarketData) SetVar7995(v string)`

SetVar7995 sets Var7995 field to given value.

### HasVar7995

`func (o *MarketData) HasVar7995() bool`

HasVar7995 returns a boolean if a field has been set.

### GetConid

`func (o *MarketData) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *MarketData) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *MarketData) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *MarketData) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetMinTick

`func (o *MarketData) GetMinTick() float32`

GetMinTick returns the MinTick field if non-nil, zero value otherwise.

### GetMinTickOk

`func (o *MarketData) GetMinTickOk() (*float32, bool)`

GetMinTickOk returns a tuple with the MinTick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTick

`func (o *MarketData) SetMinTick(v float32)`

SetMinTick sets MinTick field to given value.

### HasMinTick

`func (o *MarketData) HasMinTick() bool`

HasMinTick returns a boolean if a field has been set.

### GetBboExchange

`func (o *MarketData) GetBboExchange() string`

GetBboExchange returns the BboExchange field if non-nil, zero value otherwise.

### GetBboExchangeOk

`func (o *MarketData) GetBboExchangeOk() (*string, bool)`

GetBboExchangeOk returns a tuple with the BboExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBboExchange

`func (o *MarketData) SetBboExchange(v string)`

SetBboExchange sets BboExchange field to given value.

### HasBboExchange

`func (o *MarketData) HasBboExchange() bool`

HasBboExchange returns a boolean if a field has been set.

### GetHasDelayed

`func (o *MarketData) GetHasDelayed() bool`

GetHasDelayed returns the HasDelayed field if non-nil, zero value otherwise.

### GetHasDelayedOk

`func (o *MarketData) GetHasDelayedOk() (*bool, bool)`

GetHasDelayedOk returns a tuple with the HasDelayed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDelayed

`func (o *MarketData) SetHasDelayed(v bool)`

SetHasDelayed sets HasDelayed field to given value.

### HasHasDelayed

`func (o *MarketData) HasHasDelayed() bool`

HasHasDelayed returns a boolean if a field has been set.

### GetSizeMinTick

`func (o *MarketData) GetSizeMinTick() int32`

GetSizeMinTick returns the SizeMinTick field if non-nil, zero value otherwise.

### GetSizeMinTickOk

`func (o *MarketData) GetSizeMinTickOk() (*int32, bool)`

GetSizeMinTickOk returns a tuple with the SizeMinTick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMinTick

`func (o *MarketData) SetSizeMinTick(v int32)`

SetSizeMinTick sets SizeMinTick field to given value.

### HasSizeMinTick

`func (o *MarketData) HasSizeMinTick() bool`

HasSizeMinTick returns a boolean if a field has been set.

### GetBestEligible

`func (o *MarketData) GetBestEligible() int32`

GetBestEligible returns the BestEligible field if non-nil, zero value otherwise.

### GetBestEligibleOk

`func (o *MarketData) GetBestEligibleOk() (*int32, bool)`

GetBestEligibleOk returns a tuple with the BestEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestEligible

`func (o *MarketData) SetBestEligible(v int32)`

SetBestEligible sets BestEligible field to given value.

### HasBestEligible

`func (o *MarketData) HasBestEligible() bool`

HasBestEligible returns a boolean if a field has been set.

### GetBestBidExch

`func (o *MarketData) GetBestBidExch() int32`

GetBestBidExch returns the BestBidExch field if non-nil, zero value otherwise.

### GetBestBidExchOk

`func (o *MarketData) GetBestBidExchOk() (*int32, bool)`

GetBestBidExchOk returns a tuple with the BestBidExch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestBidExch

`func (o *MarketData) SetBestBidExch(v int32)`

SetBestBidExch sets BestBidExch field to given value.

### HasBestBidExch

`func (o *MarketData) HasBestBidExch() bool`

HasBestBidExch returns a boolean if a field has been set.

### GetBestAskExch

`func (o *MarketData) GetBestAskExch() int32`

GetBestAskExch returns the BestAskExch field if non-nil, zero value otherwise.

### GetBestAskExchOk

`func (o *MarketData) GetBestAskExchOk() (*int32, bool)`

GetBestAskExchOk returns a tuple with the BestAskExch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestAskExch

`func (o *MarketData) SetBestAskExch(v int32)`

SetBestAskExch sets BestAskExch field to given value.

### HasBestAskExch

`func (o *MarketData) HasBestAskExch() bool`

HasBestAskExch returns a boolean if a field has been set.

### GetPreOpenBid

`func (o *MarketData) GetPreOpenBid() int32`

GetPreOpenBid returns the PreOpenBid field if non-nil, zero value otherwise.

### GetPreOpenBidOk

`func (o *MarketData) GetPreOpenBidOk() (*int32, bool)`

GetPreOpenBidOk returns a tuple with the PreOpenBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOpenBid

`func (o *MarketData) SetPreOpenBid(v int32)`

SetPreOpenBid sets PreOpenBid field to given value.

### HasPreOpenBid

`func (o *MarketData) HasPreOpenBid() bool`

HasPreOpenBid returns a boolean if a field has been set.

### GetLastAttribs

`func (o *MarketData) GetLastAttribs() int32`

GetLastAttribs returns the LastAttribs field if non-nil, zero value otherwise.

### GetLastAttribsOk

`func (o *MarketData) GetLastAttribsOk() (*int32, bool)`

GetLastAttribsOk returns a tuple with the LastAttribs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAttribs

`func (o *MarketData) SetLastAttribs(v int32)`

SetLastAttribs sets LastAttribs field to given value.

### HasLastAttribs

`func (o *MarketData) HasLastAttribs() bool`

HasLastAttribs returns a boolean if a field has been set.

### GetTimestampBase

`func (o *MarketData) GetTimestampBase() int32`

GetTimestampBase returns the TimestampBase field if non-nil, zero value otherwise.

### GetTimestampBaseOk

`func (o *MarketData) GetTimestampBaseOk() (*int32, bool)`

GetTimestampBaseOk returns a tuple with the TimestampBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampBase

`func (o *MarketData) SetTimestampBase(v int32)`

SetTimestampBase sets TimestampBase field to given value.

### HasTimestampBase

`func (o *MarketData) HasTimestampBase() bool`

HasTimestampBase returns a boolean if a field has been set.

### GetTimestampDelta

`func (o *MarketData) GetTimestampDelta() int32`

GetTimestampDelta returns the TimestampDelta field if non-nil, zero value otherwise.

### GetTimestampDeltaOk

`func (o *MarketData) GetTimestampDeltaOk() (*int32, bool)`

GetTimestampDeltaOk returns a tuple with the TimestampDelta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampDelta

`func (o *MarketData) SetTimestampDelta(v int32)`

SetTimestampDelta sets TimestampDelta field to given value.

### HasTimestampDelta

`func (o *MarketData) HasTimestampDelta() bool`

HasTimestampDelta returns a boolean if a field has been set.

### GetLastExch

`func (o *MarketData) GetLastExch() int32`

GetLastExch returns the LastExch field if non-nil, zero value otherwise.

### GetLastExchOk

`func (o *MarketData) GetLastExchOk() (*int32, bool)`

GetLastExchOk returns a tuple with the LastExch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastExch

`func (o *MarketData) SetLastExch(v int32)`

SetLastExch sets LastExch field to given value.

### HasLastExch

`func (o *MarketData) HasLastExch() bool`

HasLastExch returns a boolean if a field has been set.

### GetCloseAttribs

`func (o *MarketData) GetCloseAttribs() int32`

GetCloseAttribs returns the CloseAttribs field if non-nil, zero value otherwise.

### GetCloseAttribsOk

`func (o *MarketData) GetCloseAttribsOk() (*int32, bool)`

GetCloseAttribsOk returns a tuple with the CloseAttribs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseAttribs

`func (o *MarketData) SetCloseAttribs(v int32)`

SetCloseAttribs sets CloseAttribs field to given value.

### HasCloseAttribs

`func (o *MarketData) HasCloseAttribs() bool`

HasCloseAttribs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


