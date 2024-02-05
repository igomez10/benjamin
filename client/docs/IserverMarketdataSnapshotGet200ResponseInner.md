# IserverMarketdataSnapshotGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var31** | Pointer to **string** | Last Price - The last price at which the contract traded. May contain one of the following prefixes:   * C - Previous day&#39;s closing price.   * H - Trading has halted.  | [optional] 
**Var55** | Pointer to **string** | Symbol | [optional] 
**Var58** | Pointer to **string** | Text | [optional] 
**Var70** | Pointer to **string** | High - Current day high price | [optional] 
**Var71** | Pointer to **string** | Low - Current day low price | [optional] 
**Var73** | Pointer to **string** | Market Value - The current market value of  your position in the security. Market Value is calculated with real time market data (even when not subscribed to market data). | [optional] 
**Var74** | Pointer to **string** | Avg Price - The average price of the position. | [optional] 
**Var75** | Pointer to **string** | Unrealized PnL - Unrealized profit or loss. Unrealized PnL is calculated with real time market data (even when not subscribed to market data). | [optional] 
**Var76** | Pointer to **string** | Formatted position | [optional] 
**Var77** | Pointer to **string** | Formatted Unrealized PnL | [optional] 
**Var78** | Pointer to **string** | Daily PnL - Your profit or loss of the day since prior close. Daily PnL is calculated with real time market data (even when not subscribed to market data). | [optional] 
**Var79** | Pointer to **string** | Realized PnL - Realized profit or loss. Realized PnL is calculated with real time market data (even when not subscribed to market data). | [optional] 
**Var80** | Pointer to **string** | Unrealized PnL % - Unrealized profit or loss expressed in percentage. | [optional] 
**Var82** | Pointer to **string** | Change - The difference between the last price and the close on the previous trading day | [optional] 
**Var83** | Pointer to **string** | Change % - The difference between the last price and the close on the previous trading day in percentage. | [optional] 
**Var84** | Pointer to **string** | Bid Price - The highest-priced bid on the contract. | [optional] 
**Var85** | Pointer to **string** | Ask Size - The number of contracts or shares offered at the ask price. For US stocks, the number displayed is divided by 100. | [optional] 
**Var86** | Pointer to **string** | Ask Price - The lowest-priced offer on the contract. | [optional] 
**Var87** | Pointer to **string** | Volume - Volume for the day, formatted with &#39;K&#39; for thousands or &#39;M&#39; for millions. For higher precision volume refer to field 7762. | [optional] 
**Var88** | Pointer to **string** | Bid Size - The number of contracts or shares bid for at the bid price. For US stocks, the number displayed is divided by 100. | [optional] 
**Var6004** | Pointer to **string** | Exchange | [optional] 
**Var6008** | Pointer to **int32** | Conid - Contract identifier from IBKR&#39;s database. | [optional] 
**Var6070** | Pointer to **string** | SecType - The asset class of the instrument. | [optional] 
**Var6072** | Pointer to **string** | Months | [optional] 
**Var6073** | Pointer to **string** | Regular Expiry | [optional] 
**Var6119** | Pointer to **string** | Marker for market data delivery method (similar to request id) | [optional] 
**Var6457** | Pointer to **int32** | Underlying Conid. Use /trsrv/secdef to get more information about the security | [optional] 
**Var6508** | Pointer to **string** | Service Params. | [optional] 
**Var6509** | Pointer to **string** | Market Data Availability. The field may contain three chars. First char defines: R &#x3D; RealTime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed, N &#x3D; Not Subscribed. Second char defines: P &#x3D; Snapshot, p &#x3D; Consolidated. Third char defines: B &#x3D; Book   * RealTime - Data is relayed back in real time without delay, market data subscription(s) are required.   * Delayed - Data is relayed back 15-20 min delayed.    * Frozen - Last recorded data at market close, relayed back in real time.   * Frozen Delayed - Last recorded data at market close, relayed back delayed.   * Not Subscribed - User does not have the required market data subscription(s) to relay back either real time or delayed data.   * Snapshot - Snapshot request is available for contract.   * Consolidated - Market data is aggregated across multiple exchanges or venues.   * Book - Top of the book data is available for contract.  | [optional] 
**Var7051** | Pointer to **string** | Company name | [optional] 
**Var7057** | Pointer to **string** | Ask Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY  | [optional] 
**Var7058** | Pointer to **string** | Last Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY  | [optional] 
**Var7059** | Pointer to **string** | Last Size - The number of unites traded at the last price | [optional] 
**Var7068** | Pointer to **string** | Bid Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY  | [optional] 
**Var7084** | Pointer to **string** | Implied Vol./Hist. Vol % - The ratio of the implied volatility over the historical volatility, expressed as a percentage. | [optional] 
**Var7085** | Pointer to **string** | Put/Call Interest - Put option open interest/call option open interest for the trading day. | [optional] 
**Var7086** | Pointer to **string** | Put/Call Volume - Put option volume/call option volume for the trading day. | [optional] 
**Var7087** | Pointer to **string** | Hist. Vol. % - 30-day real-time historical volatility. | [optional] 
**Var7088** | Pointer to **string** | Hist. Vol. Close % - Shows the historical volatility based on previous close price. | [optional] 
**Var7089** | Pointer to **string** | Opt. Volume - Option Volume | [optional] 
**Var7094** | Pointer to **string** | Conid + Exchange | [optional] 
**Var7184** | Pointer to **string** | canBeTraded - If contract is a trade-able instrument. Returns 1(true) or 0(false). | [optional] 
**Var7219** | Pointer to **string** | Contract Description | [optional] 
**Var7220** | Pointer to **string** | Contract Description | [optional] 
**Var7221** | Pointer to **string** | Listing Exchange | [optional] 
**Var7280** | Pointer to **string** | Industry - Displays the type of industry under which the underlying company can be categorized. | [optional] 
**Var7281** | Pointer to **string** | Category - Displays a more detailed level of description within the industry under which the underlying company can be categorized. | [optional] 
**Var7282** | Pointer to **string** | Average Volume - The average daily trading volume over 90 days. | [optional] 
**Var7283** | Pointer to **string** | Option Implied Vol. % - A prediction of how volatile an underlying will be in the future. At the market volatility estimated for a maturity thirty calendar days forward of the current trading day, and based on option prices from two consecutive expiration months. To query the Implied Vol. % of a specific strike refer to field 7633.  | [optional] 
**Var7284** | Pointer to **string** | Historic Volume (30d) | [optional] 
**Var7285** | Pointer to **string** | Put/Call Ratio | [optional] 
**Var7286** | Pointer to **string** | Dividend Amount - Displays the amount of the next dividend. | [optional] 
**Var7287** | Pointer to **string** | Dividend Yield % - This value is the toal of the expected dividend payments over the next twelve months per share divided by the Current Price and is expressed as a percentage. For derivatives, this displays the total of the expected dividend payments over the expiry date.  | [optional] 
**Var7288** | Pointer to **string** | Ex-date of the dividend | [optional] 
**Var7289** | Pointer to **string** | Market Cap | [optional] 
**Var7290** | Pointer to **string** | P/E | [optional] 
**Var7291** | Pointer to **string** | EPS | [optional] 
**Var7292** | Pointer to **string** | Cost Basis - Your current position in this security multiplied by the average price and multiplier. | [optional] 
**Var7293** | Pointer to **string** | 52 Week High - The highest price for the past 52 weeks. | [optional] 
**Var7294** | Pointer to **string** | 52 Week Low - The lowest price for the past 52 weeks. | [optional] 
**Var7295** | Pointer to **string** | Open - Today&#39;s opening price. | [optional] 
**Var7296** | Pointer to **string** | Close - Today&#39;s closing price. | [optional] 
**Var7308** | Pointer to **string** | Delta - The ratio of the change in the price of the option to the corresponding change in the price of the underlying. | [optional] 
**Var7309** | Pointer to **string** | Gamma - The rate of change for the delta with respect to the underlying asset&#39;s price. | [optional] 
**Var7310** | Pointer to **string** | Theta - A measure of the rate of decline the value of an option due to the passage of time. | [optional] 
**Var7311** | Pointer to **string** | Vega - The amount that the price of an option changes compared to a 1% change in the volatility. | [optional] 
**Var7607** | Pointer to **string** | Opt. Volume Change % - Today&#39;s option volume as a percentage of the average option volume. | [optional] 
**Var7633** | Pointer to **string** | Implied Vol. % - The implied volatility for the specific strike of the option in percentage. To query the Option Implied Vol. % from the underlying refer to field 7283.  | [optional] 
**Var7635** | Pointer to **string** | Mark - The mark price is, the ask price if ask is less than last price, the bid price if bid is more than the last price, otherwise it&#39;s equal to last price. | [optional] 
**Var7636** | Pointer to **string** | Shortable Shares - Number of shares available for shorting. | [optional] 
**Var7637** | Pointer to **string** | Fee Rate - Interest rate charged on borrowed shares. | [optional] 
**Var7638** | Pointer to **string** | Option Open Interest | [optional] 
**Var7639** | Pointer to **string** | % of Mark Value - Displays the market value of the contract as a percentage of the total market value of the account. Mark Value is calculated with real time market data (even when not subscribed to market data).  | [optional] 
**Var7644** | Pointer to **string** | Shortable - Describes the level of difficulty with which the security can be sold short. | [optional] 
**Var7655** | Pointer to **string** | Morningstar Rating - Displays Morningstar Rating provided value. Requires [Morningstar](https://www.interactivebrokers.com/en/index.php?f&#x3D;14262) subscription. | [optional] 
**Var7671** | Pointer to **string** | Dividends - This value is the total of the expected dividend payments over the next twelve months per share. | [optional] 
**Var7672** | Pointer to **string** | Dividends TTM - This value is the total of the expected dividend payments over the last twelve months per share. | [optional] 
**Var7674** | Pointer to **string** | EMA(200) - Exponential moving average (N&#x3D;200). | [optional] 
**Var7675** | Pointer to **string** | EMA(100) - Exponential moving average (N&#x3D;100). | [optional] 
**Var7676** | Pointer to **string** | EMA(50) - Exponential moving average (N&#x3D;50). | [optional] 
**Var7677** | Pointer to **string** | EMA(20) - Exponential moving average (N&#x3D;20). | [optional] 
**Var7678** | Pointer to **string** | Price/EMA(200) - Price to Exponential moving average (N&#x3D;200) ratio -1, displayed in percents. | [optional] 
**Var7679** | Pointer to **string** | Price/EMA(100) - Price to Exponential moving average (N&#x3D;100) ratio -1, displayed in percents. | [optional] 
**Var7680** | Pointer to **string** | Price/EMA(50) - Price to Exponential moving average (N&#x3D;50) ratio -1, displayed in percents. | [optional] 
**Var7681** | Pointer to **string** | Price/EMA(20) - Price to Exponential moving average (N&#x3D;20) ratio -1, displayed in percents. | [optional] 
**Var7682** | Pointer to **string** | Change Since Open - The difference between the last price and the open price. | [optional] 
**Var7683** | Pointer to **string** | Upcoming Event - Shows the next major company event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] 
**Var7684** | Pointer to **string** | Upcoming Event Date - The date of the next major company event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] 
**Var7685** | Pointer to **string** | Upcoming Analyst Meeting - The date and time of the next scheduled analyst meeting. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] 
**Var7686** | Pointer to **string** | Upcoming Earnings - The date and time of the next scheduled earnings/earnings call event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] 
**Var7687** | Pointer to **string** | Upcoming Misc Event - The date and time of the next shareholder meeting, presentation or other event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] 
**Var7688** | Pointer to **string** | Recent Analyst Meeting - The date and time of the most recent analyst meeting. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] 
**Var7689** | Pointer to **string** | Recent Earnings - The date and time of the most recent earnings/earning call event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] 
**Var7690** | Pointer to **string** | Recent Misc Event - The date and time of the most recent shareholder meeting, presentation or other event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] 
**Var7694** | Pointer to **string** | Probability of Max Return - Customer implied probability of maximum potential gain. | [optional] 
**Var7695** | Pointer to **string** | Break Even - Break even points | [optional] 
**Var7696** | Pointer to **string** | SPX Delta - Beta Weighted Delta is calculated using the formula; Delta x dollar adjusted beta, where adjusted beta is adjusted by the ratio of the close price. | [optional] 
**Var7697** | Pointer to **string** | Futures Open Interest - Total number of outstanding futures contracts | [optional] 
**Var7698** | Pointer to **string** | Last Yield - Implied yield of the bond if it is purchased at the current last price. Last yield is calculated using the Last price on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] 
**Var7699** | Pointer to **string** | Bid Yield - Implied yield of the bond if it is purchased at the current bid price. Bid yield is calculated using the Ask on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] 
**Var7700** | Pointer to **string** | Probability of Max Return - Customer implied probability of maximum potential gain. | [optional] 
**Var7702** | Pointer to **string** | Probability of Max Loss - Customer implied probability of maximum potential loss. | [optional] 
**Var7703** | Pointer to **string** | Profit Probability - Customer implied probability of any gain. | [optional] 
**Var7704** | Pointer to **string** | Organization Type | [optional] 
**Var7705** | Pointer to **string** | Debt Class | [optional] 
**Var7706** | Pointer to **string** | Ratings - Ratings issued for bond contract. | [optional] 
**Var7707** | Pointer to **string** | Bond State Code | [optional] 
**Var7708** | Pointer to **string** | Bond Type | [optional] 
**Var7714** | Pointer to **string** | Last Trading Date | [optional] 
**Var7715** | Pointer to **string** | Issue Date | [optional] 
**Var7718** | Pointer to **string** | Beta - Beta is against standard index. | [optional] 
**Var7720** | Pointer to **string** | Ask Yield - Implied yield of the bond if it is purchased at the current offer. Ask yield is calculated using the Bid on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] 
**Var7741** | Pointer to **string** | Prior Close - Yesterday&#39;s closing price | [optional] 
**Var7762** | Pointer to **string** | Volume Long - High precision volume for the day. For formatted volume refer to field 87. | [optional] 
**Var7768** | Pointer to **string** | hasTradingPermissions - if user has trading permissions for specified contract. Returns 1(true) or 0(false). | [optional] 
**ServerId** | Pointer to **string** |  | [optional] 
**Conid** | Pointer to **int32** |  | [optional] 
**Updated** | Pointer to **int32** |  | [optional] 
**Var87RawDeprecated** | Pointer to **string** | Raw Volume - Volume for the day, provided in long form without formatted with K/M. This field value is deprecated, for high precision volume refer to field 7762. | [optional] 

## Methods

### NewIserverMarketdataSnapshotGet200ResponseInner

`func NewIserverMarketdataSnapshotGet200ResponseInner() *IserverMarketdataSnapshotGet200ResponseInner`

NewIserverMarketdataSnapshotGet200ResponseInner instantiates a new IserverMarketdataSnapshotGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIserverMarketdataSnapshotGet200ResponseInnerWithDefaults

`func NewIserverMarketdataSnapshotGet200ResponseInnerWithDefaults() *IserverMarketdataSnapshotGet200ResponseInner`

NewIserverMarketdataSnapshotGet200ResponseInnerWithDefaults instantiates a new IserverMarketdataSnapshotGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar31

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar31() string`

GetVar31 returns the Var31 field if non-nil, zero value otherwise.

### GetVar31Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar31Ok() (*string, bool)`

GetVar31Ok returns a tuple with the Var31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar31

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar31(v string)`

SetVar31 sets Var31 field to given value.

### HasVar31

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar31() bool`

HasVar31 returns a boolean if a field has been set.

### GetVar55

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar55() string`

GetVar55 returns the Var55 field if non-nil, zero value otherwise.

### GetVar55Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar55Ok() (*string, bool)`

GetVar55Ok returns a tuple with the Var55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar55

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar55(v string)`

SetVar55 sets Var55 field to given value.

### HasVar55

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar55() bool`

HasVar55 returns a boolean if a field has been set.

### GetVar58

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar58() string`

GetVar58 returns the Var58 field if non-nil, zero value otherwise.

### GetVar58Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar58Ok() (*string, bool)`

GetVar58Ok returns a tuple with the Var58 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar58

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar58(v string)`

SetVar58 sets Var58 field to given value.

### HasVar58

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar58() bool`

HasVar58 returns a boolean if a field has been set.

### GetVar70

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar70() string`

GetVar70 returns the Var70 field if non-nil, zero value otherwise.

### GetVar70Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar70Ok() (*string, bool)`

GetVar70Ok returns a tuple with the Var70 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar70

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar70(v string)`

SetVar70 sets Var70 field to given value.

### HasVar70

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar70() bool`

HasVar70 returns a boolean if a field has been set.

### GetVar71

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar71() string`

GetVar71 returns the Var71 field if non-nil, zero value otherwise.

### GetVar71Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar71Ok() (*string, bool)`

GetVar71Ok returns a tuple with the Var71 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar71

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar71(v string)`

SetVar71 sets Var71 field to given value.

### HasVar71

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar71() bool`

HasVar71 returns a boolean if a field has been set.

### GetVar73

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar73() string`

GetVar73 returns the Var73 field if non-nil, zero value otherwise.

### GetVar73Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar73Ok() (*string, bool)`

GetVar73Ok returns a tuple with the Var73 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar73

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar73(v string)`

SetVar73 sets Var73 field to given value.

### HasVar73

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar73() bool`

HasVar73 returns a boolean if a field has been set.

### GetVar74

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar74() string`

GetVar74 returns the Var74 field if non-nil, zero value otherwise.

### GetVar74Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar74Ok() (*string, bool)`

GetVar74Ok returns a tuple with the Var74 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar74

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar74(v string)`

SetVar74 sets Var74 field to given value.

### HasVar74

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar74() bool`

HasVar74 returns a boolean if a field has been set.

### GetVar75

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar75() string`

GetVar75 returns the Var75 field if non-nil, zero value otherwise.

### GetVar75Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar75Ok() (*string, bool)`

GetVar75Ok returns a tuple with the Var75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar75

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar75(v string)`

SetVar75 sets Var75 field to given value.

### HasVar75

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar75() bool`

HasVar75 returns a boolean if a field has been set.

### GetVar76

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar76() string`

GetVar76 returns the Var76 field if non-nil, zero value otherwise.

### GetVar76Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar76Ok() (*string, bool)`

GetVar76Ok returns a tuple with the Var76 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar76

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar76(v string)`

SetVar76 sets Var76 field to given value.

### HasVar76

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar76() bool`

HasVar76 returns a boolean if a field has been set.

### GetVar77

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar77() string`

GetVar77 returns the Var77 field if non-nil, zero value otherwise.

### GetVar77Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar77Ok() (*string, bool)`

GetVar77Ok returns a tuple with the Var77 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar77

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar77(v string)`

SetVar77 sets Var77 field to given value.

### HasVar77

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar77() bool`

HasVar77 returns a boolean if a field has been set.

### GetVar78

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar78() string`

GetVar78 returns the Var78 field if non-nil, zero value otherwise.

### GetVar78Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar78Ok() (*string, bool)`

GetVar78Ok returns a tuple with the Var78 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar78

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar78(v string)`

SetVar78 sets Var78 field to given value.

### HasVar78

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar78() bool`

HasVar78 returns a boolean if a field has been set.

### GetVar79

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar79() string`

GetVar79 returns the Var79 field if non-nil, zero value otherwise.

### GetVar79Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar79Ok() (*string, bool)`

GetVar79Ok returns a tuple with the Var79 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar79

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar79(v string)`

SetVar79 sets Var79 field to given value.

### HasVar79

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar79() bool`

HasVar79 returns a boolean if a field has been set.

### GetVar80

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar80() string`

GetVar80 returns the Var80 field if non-nil, zero value otherwise.

### GetVar80Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar80Ok() (*string, bool)`

GetVar80Ok returns a tuple with the Var80 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar80

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar80(v string)`

SetVar80 sets Var80 field to given value.

### HasVar80

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar80() bool`

HasVar80 returns a boolean if a field has been set.

### GetVar82

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar82() string`

GetVar82 returns the Var82 field if non-nil, zero value otherwise.

### GetVar82Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar82Ok() (*string, bool)`

GetVar82Ok returns a tuple with the Var82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar82

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar82(v string)`

SetVar82 sets Var82 field to given value.

### HasVar82

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar82() bool`

HasVar82 returns a boolean if a field has been set.

### GetVar83

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar83() string`

GetVar83 returns the Var83 field if non-nil, zero value otherwise.

### GetVar83Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar83Ok() (*string, bool)`

GetVar83Ok returns a tuple with the Var83 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar83

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar83(v string)`

SetVar83 sets Var83 field to given value.

### HasVar83

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar83() bool`

HasVar83 returns a boolean if a field has been set.

### GetVar84

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar84() string`

GetVar84 returns the Var84 field if non-nil, zero value otherwise.

### GetVar84Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar84Ok() (*string, bool)`

GetVar84Ok returns a tuple with the Var84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar84

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar84(v string)`

SetVar84 sets Var84 field to given value.

### HasVar84

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar84() bool`

HasVar84 returns a boolean if a field has been set.

### GetVar85

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar85() string`

GetVar85 returns the Var85 field if non-nil, zero value otherwise.

### GetVar85Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar85Ok() (*string, bool)`

GetVar85Ok returns a tuple with the Var85 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar85

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar85(v string)`

SetVar85 sets Var85 field to given value.

### HasVar85

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar85() bool`

HasVar85 returns a boolean if a field has been set.

### GetVar86

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar86() string`

GetVar86 returns the Var86 field if non-nil, zero value otherwise.

### GetVar86Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar86Ok() (*string, bool)`

GetVar86Ok returns a tuple with the Var86 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar86

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar86(v string)`

SetVar86 sets Var86 field to given value.

### HasVar86

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar86() bool`

HasVar86 returns a boolean if a field has been set.

### GetVar87

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar87() string`

GetVar87 returns the Var87 field if non-nil, zero value otherwise.

### GetVar87Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar87Ok() (*string, bool)`

GetVar87Ok returns a tuple with the Var87 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar87

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar87(v string)`

SetVar87 sets Var87 field to given value.

### HasVar87

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar87() bool`

HasVar87 returns a boolean if a field has been set.

### GetVar88

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar88() string`

GetVar88 returns the Var88 field if non-nil, zero value otherwise.

### GetVar88Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar88Ok() (*string, bool)`

GetVar88Ok returns a tuple with the Var88 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar88

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar88(v string)`

SetVar88 sets Var88 field to given value.

### HasVar88

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar88() bool`

HasVar88 returns a boolean if a field has been set.

### GetVar6004

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6004() string`

GetVar6004 returns the Var6004 field if non-nil, zero value otherwise.

### GetVar6004Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6004Ok() (*string, bool)`

GetVar6004Ok returns a tuple with the Var6004 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6004

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6004(v string)`

SetVar6004 sets Var6004 field to given value.

### HasVar6004

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6004() bool`

HasVar6004 returns a boolean if a field has been set.

### GetVar6008

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6008() int32`

GetVar6008 returns the Var6008 field if non-nil, zero value otherwise.

### GetVar6008Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6008Ok() (*int32, bool)`

GetVar6008Ok returns a tuple with the Var6008 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6008

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6008(v int32)`

SetVar6008 sets Var6008 field to given value.

### HasVar6008

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6008() bool`

HasVar6008 returns a boolean if a field has been set.

### GetVar6070

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6070() string`

GetVar6070 returns the Var6070 field if non-nil, zero value otherwise.

### GetVar6070Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6070Ok() (*string, bool)`

GetVar6070Ok returns a tuple with the Var6070 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6070

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6070(v string)`

SetVar6070 sets Var6070 field to given value.

### HasVar6070

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6070() bool`

HasVar6070 returns a boolean if a field has been set.

### GetVar6072

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6072() string`

GetVar6072 returns the Var6072 field if non-nil, zero value otherwise.

### GetVar6072Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6072Ok() (*string, bool)`

GetVar6072Ok returns a tuple with the Var6072 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6072

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6072(v string)`

SetVar6072 sets Var6072 field to given value.

### HasVar6072

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6072() bool`

HasVar6072 returns a boolean if a field has been set.

### GetVar6073

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6073() string`

GetVar6073 returns the Var6073 field if non-nil, zero value otherwise.

### GetVar6073Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6073Ok() (*string, bool)`

GetVar6073Ok returns a tuple with the Var6073 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6073

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6073(v string)`

SetVar6073 sets Var6073 field to given value.

### HasVar6073

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6073() bool`

HasVar6073 returns a boolean if a field has been set.

### GetVar6119

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6119() string`

GetVar6119 returns the Var6119 field if non-nil, zero value otherwise.

### GetVar6119Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6119Ok() (*string, bool)`

GetVar6119Ok returns a tuple with the Var6119 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6119

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6119(v string)`

SetVar6119 sets Var6119 field to given value.

### HasVar6119

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6119() bool`

HasVar6119 returns a boolean if a field has been set.

### GetVar6457

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6457() int32`

GetVar6457 returns the Var6457 field if non-nil, zero value otherwise.

### GetVar6457Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6457Ok() (*int32, bool)`

GetVar6457Ok returns a tuple with the Var6457 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6457

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6457(v int32)`

SetVar6457 sets Var6457 field to given value.

### HasVar6457

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6457() bool`

HasVar6457 returns a boolean if a field has been set.

### GetVar6508

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6508() string`

GetVar6508 returns the Var6508 field if non-nil, zero value otherwise.

### GetVar6508Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6508Ok() (*string, bool)`

GetVar6508Ok returns a tuple with the Var6508 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6508

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6508(v string)`

SetVar6508 sets Var6508 field to given value.

### HasVar6508

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6508() bool`

HasVar6508 returns a boolean if a field has been set.

### GetVar6509

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6509() string`

GetVar6509 returns the Var6509 field if non-nil, zero value otherwise.

### GetVar6509Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar6509Ok() (*string, bool)`

GetVar6509Ok returns a tuple with the Var6509 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar6509

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar6509(v string)`

SetVar6509 sets Var6509 field to given value.

### HasVar6509

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar6509() bool`

HasVar6509 returns a boolean if a field has been set.

### GetVar7051

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7051() string`

GetVar7051 returns the Var7051 field if non-nil, zero value otherwise.

### GetVar7051Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7051Ok() (*string, bool)`

GetVar7051Ok returns a tuple with the Var7051 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7051

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7051(v string)`

SetVar7051 sets Var7051 field to given value.

### HasVar7051

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7051() bool`

HasVar7051 returns a boolean if a field has been set.

### GetVar7057

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7057() string`

GetVar7057 returns the Var7057 field if non-nil, zero value otherwise.

### GetVar7057Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7057Ok() (*string, bool)`

GetVar7057Ok returns a tuple with the Var7057 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7057

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7057(v string)`

SetVar7057 sets Var7057 field to given value.

### HasVar7057

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7057() bool`

HasVar7057 returns a boolean if a field has been set.

### GetVar7058

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7058() string`

GetVar7058 returns the Var7058 field if non-nil, zero value otherwise.

### GetVar7058Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7058Ok() (*string, bool)`

GetVar7058Ok returns a tuple with the Var7058 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7058

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7058(v string)`

SetVar7058 sets Var7058 field to given value.

### HasVar7058

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7058() bool`

HasVar7058 returns a boolean if a field has been set.

### GetVar7059

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7059() string`

GetVar7059 returns the Var7059 field if non-nil, zero value otherwise.

### GetVar7059Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7059Ok() (*string, bool)`

GetVar7059Ok returns a tuple with the Var7059 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7059

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7059(v string)`

SetVar7059 sets Var7059 field to given value.

### HasVar7059

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7059() bool`

HasVar7059 returns a boolean if a field has been set.

### GetVar7068

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7068() string`

GetVar7068 returns the Var7068 field if non-nil, zero value otherwise.

### GetVar7068Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7068Ok() (*string, bool)`

GetVar7068Ok returns a tuple with the Var7068 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7068

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7068(v string)`

SetVar7068 sets Var7068 field to given value.

### HasVar7068

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7068() bool`

HasVar7068 returns a boolean if a field has been set.

### GetVar7084

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7084() string`

GetVar7084 returns the Var7084 field if non-nil, zero value otherwise.

### GetVar7084Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7084Ok() (*string, bool)`

GetVar7084Ok returns a tuple with the Var7084 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7084

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7084(v string)`

SetVar7084 sets Var7084 field to given value.

### HasVar7084

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7084() bool`

HasVar7084 returns a boolean if a field has been set.

### GetVar7085

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7085() string`

GetVar7085 returns the Var7085 field if non-nil, zero value otherwise.

### GetVar7085Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7085Ok() (*string, bool)`

GetVar7085Ok returns a tuple with the Var7085 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7085

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7085(v string)`

SetVar7085 sets Var7085 field to given value.

### HasVar7085

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7085() bool`

HasVar7085 returns a boolean if a field has been set.

### GetVar7086

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7086() string`

GetVar7086 returns the Var7086 field if non-nil, zero value otherwise.

### GetVar7086Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7086Ok() (*string, bool)`

GetVar7086Ok returns a tuple with the Var7086 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7086

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7086(v string)`

SetVar7086 sets Var7086 field to given value.

### HasVar7086

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7086() bool`

HasVar7086 returns a boolean if a field has been set.

### GetVar7087

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7087() string`

GetVar7087 returns the Var7087 field if non-nil, zero value otherwise.

### GetVar7087Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7087Ok() (*string, bool)`

GetVar7087Ok returns a tuple with the Var7087 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7087

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7087(v string)`

SetVar7087 sets Var7087 field to given value.

### HasVar7087

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7087() bool`

HasVar7087 returns a boolean if a field has been set.

### GetVar7088

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7088() string`

GetVar7088 returns the Var7088 field if non-nil, zero value otherwise.

### GetVar7088Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7088Ok() (*string, bool)`

GetVar7088Ok returns a tuple with the Var7088 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7088

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7088(v string)`

SetVar7088 sets Var7088 field to given value.

### HasVar7088

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7088() bool`

HasVar7088 returns a boolean if a field has been set.

### GetVar7089

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7089() string`

GetVar7089 returns the Var7089 field if non-nil, zero value otherwise.

### GetVar7089Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7089Ok() (*string, bool)`

GetVar7089Ok returns a tuple with the Var7089 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7089

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7089(v string)`

SetVar7089 sets Var7089 field to given value.

### HasVar7089

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7089() bool`

HasVar7089 returns a boolean if a field has been set.

### GetVar7094

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7094() string`

GetVar7094 returns the Var7094 field if non-nil, zero value otherwise.

### GetVar7094Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7094Ok() (*string, bool)`

GetVar7094Ok returns a tuple with the Var7094 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7094

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7094(v string)`

SetVar7094 sets Var7094 field to given value.

### HasVar7094

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7094() bool`

HasVar7094 returns a boolean if a field has been set.

### GetVar7184

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7184() string`

GetVar7184 returns the Var7184 field if non-nil, zero value otherwise.

### GetVar7184Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7184Ok() (*string, bool)`

GetVar7184Ok returns a tuple with the Var7184 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7184

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7184(v string)`

SetVar7184 sets Var7184 field to given value.

### HasVar7184

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7184() bool`

HasVar7184 returns a boolean if a field has been set.

### GetVar7219

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7219() string`

GetVar7219 returns the Var7219 field if non-nil, zero value otherwise.

### GetVar7219Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7219Ok() (*string, bool)`

GetVar7219Ok returns a tuple with the Var7219 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7219

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7219(v string)`

SetVar7219 sets Var7219 field to given value.

### HasVar7219

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7219() bool`

HasVar7219 returns a boolean if a field has been set.

### GetVar7220

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7220() string`

GetVar7220 returns the Var7220 field if non-nil, zero value otherwise.

### GetVar7220Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7220Ok() (*string, bool)`

GetVar7220Ok returns a tuple with the Var7220 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7220

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7220(v string)`

SetVar7220 sets Var7220 field to given value.

### HasVar7220

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7220() bool`

HasVar7220 returns a boolean if a field has been set.

### GetVar7221

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7221() string`

GetVar7221 returns the Var7221 field if non-nil, zero value otherwise.

### GetVar7221Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7221Ok() (*string, bool)`

GetVar7221Ok returns a tuple with the Var7221 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7221

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7221(v string)`

SetVar7221 sets Var7221 field to given value.

### HasVar7221

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7221() bool`

HasVar7221 returns a boolean if a field has been set.

### GetVar7280

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7280() string`

GetVar7280 returns the Var7280 field if non-nil, zero value otherwise.

### GetVar7280Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7280Ok() (*string, bool)`

GetVar7280Ok returns a tuple with the Var7280 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7280

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7280(v string)`

SetVar7280 sets Var7280 field to given value.

### HasVar7280

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7280() bool`

HasVar7280 returns a boolean if a field has been set.

### GetVar7281

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7281() string`

GetVar7281 returns the Var7281 field if non-nil, zero value otherwise.

### GetVar7281Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7281Ok() (*string, bool)`

GetVar7281Ok returns a tuple with the Var7281 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7281

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7281(v string)`

SetVar7281 sets Var7281 field to given value.

### HasVar7281

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7281() bool`

HasVar7281 returns a boolean if a field has been set.

### GetVar7282

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7282() string`

GetVar7282 returns the Var7282 field if non-nil, zero value otherwise.

### GetVar7282Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7282Ok() (*string, bool)`

GetVar7282Ok returns a tuple with the Var7282 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7282

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7282(v string)`

SetVar7282 sets Var7282 field to given value.

### HasVar7282

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7282() bool`

HasVar7282 returns a boolean if a field has been set.

### GetVar7283

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7283() string`

GetVar7283 returns the Var7283 field if non-nil, zero value otherwise.

### GetVar7283Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7283Ok() (*string, bool)`

GetVar7283Ok returns a tuple with the Var7283 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7283

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7283(v string)`

SetVar7283 sets Var7283 field to given value.

### HasVar7283

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7283() bool`

HasVar7283 returns a boolean if a field has been set.

### GetVar7284

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7284() string`

GetVar7284 returns the Var7284 field if non-nil, zero value otherwise.

### GetVar7284Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7284Ok() (*string, bool)`

GetVar7284Ok returns a tuple with the Var7284 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7284

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7284(v string)`

SetVar7284 sets Var7284 field to given value.

### HasVar7284

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7284() bool`

HasVar7284 returns a boolean if a field has been set.

### GetVar7285

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7285() string`

GetVar7285 returns the Var7285 field if non-nil, zero value otherwise.

### GetVar7285Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7285Ok() (*string, bool)`

GetVar7285Ok returns a tuple with the Var7285 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7285

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7285(v string)`

SetVar7285 sets Var7285 field to given value.

### HasVar7285

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7285() bool`

HasVar7285 returns a boolean if a field has been set.

### GetVar7286

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7286() string`

GetVar7286 returns the Var7286 field if non-nil, zero value otherwise.

### GetVar7286Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7286Ok() (*string, bool)`

GetVar7286Ok returns a tuple with the Var7286 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7286

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7286(v string)`

SetVar7286 sets Var7286 field to given value.

### HasVar7286

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7286() bool`

HasVar7286 returns a boolean if a field has been set.

### GetVar7287

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7287() string`

GetVar7287 returns the Var7287 field if non-nil, zero value otherwise.

### GetVar7287Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7287Ok() (*string, bool)`

GetVar7287Ok returns a tuple with the Var7287 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7287

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7287(v string)`

SetVar7287 sets Var7287 field to given value.

### HasVar7287

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7287() bool`

HasVar7287 returns a boolean if a field has been set.

### GetVar7288

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7288() string`

GetVar7288 returns the Var7288 field if non-nil, zero value otherwise.

### GetVar7288Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7288Ok() (*string, bool)`

GetVar7288Ok returns a tuple with the Var7288 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7288

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7288(v string)`

SetVar7288 sets Var7288 field to given value.

### HasVar7288

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7288() bool`

HasVar7288 returns a boolean if a field has been set.

### GetVar7289

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7289() string`

GetVar7289 returns the Var7289 field if non-nil, zero value otherwise.

### GetVar7289Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7289Ok() (*string, bool)`

GetVar7289Ok returns a tuple with the Var7289 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7289

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7289(v string)`

SetVar7289 sets Var7289 field to given value.

### HasVar7289

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7289() bool`

HasVar7289 returns a boolean if a field has been set.

### GetVar7290

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7290() string`

GetVar7290 returns the Var7290 field if non-nil, zero value otherwise.

### GetVar7290Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7290Ok() (*string, bool)`

GetVar7290Ok returns a tuple with the Var7290 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7290

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7290(v string)`

SetVar7290 sets Var7290 field to given value.

### HasVar7290

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7290() bool`

HasVar7290 returns a boolean if a field has been set.

### GetVar7291

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7291() string`

GetVar7291 returns the Var7291 field if non-nil, zero value otherwise.

### GetVar7291Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7291Ok() (*string, bool)`

GetVar7291Ok returns a tuple with the Var7291 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7291

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7291(v string)`

SetVar7291 sets Var7291 field to given value.

### HasVar7291

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7291() bool`

HasVar7291 returns a boolean if a field has been set.

### GetVar7292

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7292() string`

GetVar7292 returns the Var7292 field if non-nil, zero value otherwise.

### GetVar7292Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7292Ok() (*string, bool)`

GetVar7292Ok returns a tuple with the Var7292 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7292

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7292(v string)`

SetVar7292 sets Var7292 field to given value.

### HasVar7292

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7292() bool`

HasVar7292 returns a boolean if a field has been set.

### GetVar7293

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7293() string`

GetVar7293 returns the Var7293 field if non-nil, zero value otherwise.

### GetVar7293Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7293Ok() (*string, bool)`

GetVar7293Ok returns a tuple with the Var7293 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7293

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7293(v string)`

SetVar7293 sets Var7293 field to given value.

### HasVar7293

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7293() bool`

HasVar7293 returns a boolean if a field has been set.

### GetVar7294

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7294() string`

GetVar7294 returns the Var7294 field if non-nil, zero value otherwise.

### GetVar7294Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7294Ok() (*string, bool)`

GetVar7294Ok returns a tuple with the Var7294 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7294

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7294(v string)`

SetVar7294 sets Var7294 field to given value.

### HasVar7294

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7294() bool`

HasVar7294 returns a boolean if a field has been set.

### GetVar7295

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7295() string`

GetVar7295 returns the Var7295 field if non-nil, zero value otherwise.

### GetVar7295Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7295Ok() (*string, bool)`

GetVar7295Ok returns a tuple with the Var7295 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7295

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7295(v string)`

SetVar7295 sets Var7295 field to given value.

### HasVar7295

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7295() bool`

HasVar7295 returns a boolean if a field has been set.

### GetVar7296

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7296() string`

GetVar7296 returns the Var7296 field if non-nil, zero value otherwise.

### GetVar7296Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7296Ok() (*string, bool)`

GetVar7296Ok returns a tuple with the Var7296 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7296

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7296(v string)`

SetVar7296 sets Var7296 field to given value.

### HasVar7296

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7296() bool`

HasVar7296 returns a boolean if a field has been set.

### GetVar7308

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7308() string`

GetVar7308 returns the Var7308 field if non-nil, zero value otherwise.

### GetVar7308Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7308Ok() (*string, bool)`

GetVar7308Ok returns a tuple with the Var7308 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7308

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7308(v string)`

SetVar7308 sets Var7308 field to given value.

### HasVar7308

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7308() bool`

HasVar7308 returns a boolean if a field has been set.

### GetVar7309

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7309() string`

GetVar7309 returns the Var7309 field if non-nil, zero value otherwise.

### GetVar7309Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7309Ok() (*string, bool)`

GetVar7309Ok returns a tuple with the Var7309 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7309

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7309(v string)`

SetVar7309 sets Var7309 field to given value.

### HasVar7309

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7309() bool`

HasVar7309 returns a boolean if a field has been set.

### GetVar7310

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7310() string`

GetVar7310 returns the Var7310 field if non-nil, zero value otherwise.

### GetVar7310Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7310Ok() (*string, bool)`

GetVar7310Ok returns a tuple with the Var7310 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7310

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7310(v string)`

SetVar7310 sets Var7310 field to given value.

### HasVar7310

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7310() bool`

HasVar7310 returns a boolean if a field has been set.

### GetVar7311

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7311() string`

GetVar7311 returns the Var7311 field if non-nil, zero value otherwise.

### GetVar7311Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7311Ok() (*string, bool)`

GetVar7311Ok returns a tuple with the Var7311 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7311

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7311(v string)`

SetVar7311 sets Var7311 field to given value.

### HasVar7311

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7311() bool`

HasVar7311 returns a boolean if a field has been set.

### GetVar7607

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7607() string`

GetVar7607 returns the Var7607 field if non-nil, zero value otherwise.

### GetVar7607Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7607Ok() (*string, bool)`

GetVar7607Ok returns a tuple with the Var7607 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7607

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7607(v string)`

SetVar7607 sets Var7607 field to given value.

### HasVar7607

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7607() bool`

HasVar7607 returns a boolean if a field has been set.

### GetVar7633

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7633() string`

GetVar7633 returns the Var7633 field if non-nil, zero value otherwise.

### GetVar7633Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7633Ok() (*string, bool)`

GetVar7633Ok returns a tuple with the Var7633 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7633

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7633(v string)`

SetVar7633 sets Var7633 field to given value.

### HasVar7633

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7633() bool`

HasVar7633 returns a boolean if a field has been set.

### GetVar7635

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7635() string`

GetVar7635 returns the Var7635 field if non-nil, zero value otherwise.

### GetVar7635Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7635Ok() (*string, bool)`

GetVar7635Ok returns a tuple with the Var7635 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7635

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7635(v string)`

SetVar7635 sets Var7635 field to given value.

### HasVar7635

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7635() bool`

HasVar7635 returns a boolean if a field has been set.

### GetVar7636

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7636() string`

GetVar7636 returns the Var7636 field if non-nil, zero value otherwise.

### GetVar7636Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7636Ok() (*string, bool)`

GetVar7636Ok returns a tuple with the Var7636 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7636

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7636(v string)`

SetVar7636 sets Var7636 field to given value.

### HasVar7636

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7636() bool`

HasVar7636 returns a boolean if a field has been set.

### GetVar7637

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7637() string`

GetVar7637 returns the Var7637 field if non-nil, zero value otherwise.

### GetVar7637Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7637Ok() (*string, bool)`

GetVar7637Ok returns a tuple with the Var7637 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7637

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7637(v string)`

SetVar7637 sets Var7637 field to given value.

### HasVar7637

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7637() bool`

HasVar7637 returns a boolean if a field has been set.

### GetVar7638

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7638() string`

GetVar7638 returns the Var7638 field if non-nil, zero value otherwise.

### GetVar7638Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7638Ok() (*string, bool)`

GetVar7638Ok returns a tuple with the Var7638 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7638

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7638(v string)`

SetVar7638 sets Var7638 field to given value.

### HasVar7638

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7638() bool`

HasVar7638 returns a boolean if a field has been set.

### GetVar7639

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7639() string`

GetVar7639 returns the Var7639 field if non-nil, zero value otherwise.

### GetVar7639Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7639Ok() (*string, bool)`

GetVar7639Ok returns a tuple with the Var7639 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7639

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7639(v string)`

SetVar7639 sets Var7639 field to given value.

### HasVar7639

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7639() bool`

HasVar7639 returns a boolean if a field has been set.

### GetVar7644

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7644() string`

GetVar7644 returns the Var7644 field if non-nil, zero value otherwise.

### GetVar7644Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7644Ok() (*string, bool)`

GetVar7644Ok returns a tuple with the Var7644 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7644

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7644(v string)`

SetVar7644 sets Var7644 field to given value.

### HasVar7644

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7644() bool`

HasVar7644 returns a boolean if a field has been set.

### GetVar7655

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7655() string`

GetVar7655 returns the Var7655 field if non-nil, zero value otherwise.

### GetVar7655Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7655Ok() (*string, bool)`

GetVar7655Ok returns a tuple with the Var7655 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7655

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7655(v string)`

SetVar7655 sets Var7655 field to given value.

### HasVar7655

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7655() bool`

HasVar7655 returns a boolean if a field has been set.

### GetVar7671

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7671() string`

GetVar7671 returns the Var7671 field if non-nil, zero value otherwise.

### GetVar7671Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7671Ok() (*string, bool)`

GetVar7671Ok returns a tuple with the Var7671 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7671

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7671(v string)`

SetVar7671 sets Var7671 field to given value.

### HasVar7671

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7671() bool`

HasVar7671 returns a boolean if a field has been set.

### GetVar7672

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7672() string`

GetVar7672 returns the Var7672 field if non-nil, zero value otherwise.

### GetVar7672Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7672Ok() (*string, bool)`

GetVar7672Ok returns a tuple with the Var7672 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7672

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7672(v string)`

SetVar7672 sets Var7672 field to given value.

### HasVar7672

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7672() bool`

HasVar7672 returns a boolean if a field has been set.

### GetVar7674

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7674() string`

GetVar7674 returns the Var7674 field if non-nil, zero value otherwise.

### GetVar7674Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7674Ok() (*string, bool)`

GetVar7674Ok returns a tuple with the Var7674 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7674

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7674(v string)`

SetVar7674 sets Var7674 field to given value.

### HasVar7674

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7674() bool`

HasVar7674 returns a boolean if a field has been set.

### GetVar7675

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7675() string`

GetVar7675 returns the Var7675 field if non-nil, zero value otherwise.

### GetVar7675Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7675Ok() (*string, bool)`

GetVar7675Ok returns a tuple with the Var7675 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7675

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7675(v string)`

SetVar7675 sets Var7675 field to given value.

### HasVar7675

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7675() bool`

HasVar7675 returns a boolean if a field has been set.

### GetVar7676

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7676() string`

GetVar7676 returns the Var7676 field if non-nil, zero value otherwise.

### GetVar7676Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7676Ok() (*string, bool)`

GetVar7676Ok returns a tuple with the Var7676 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7676

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7676(v string)`

SetVar7676 sets Var7676 field to given value.

### HasVar7676

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7676() bool`

HasVar7676 returns a boolean if a field has been set.

### GetVar7677

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7677() string`

GetVar7677 returns the Var7677 field if non-nil, zero value otherwise.

### GetVar7677Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7677Ok() (*string, bool)`

GetVar7677Ok returns a tuple with the Var7677 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7677

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7677(v string)`

SetVar7677 sets Var7677 field to given value.

### HasVar7677

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7677() bool`

HasVar7677 returns a boolean if a field has been set.

### GetVar7678

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7678() string`

GetVar7678 returns the Var7678 field if non-nil, zero value otherwise.

### GetVar7678Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7678Ok() (*string, bool)`

GetVar7678Ok returns a tuple with the Var7678 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7678

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7678(v string)`

SetVar7678 sets Var7678 field to given value.

### HasVar7678

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7678() bool`

HasVar7678 returns a boolean if a field has been set.

### GetVar7679

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7679() string`

GetVar7679 returns the Var7679 field if non-nil, zero value otherwise.

### GetVar7679Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7679Ok() (*string, bool)`

GetVar7679Ok returns a tuple with the Var7679 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7679

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7679(v string)`

SetVar7679 sets Var7679 field to given value.

### HasVar7679

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7679() bool`

HasVar7679 returns a boolean if a field has been set.

### GetVar7680

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7680() string`

GetVar7680 returns the Var7680 field if non-nil, zero value otherwise.

### GetVar7680Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7680Ok() (*string, bool)`

GetVar7680Ok returns a tuple with the Var7680 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7680

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7680(v string)`

SetVar7680 sets Var7680 field to given value.

### HasVar7680

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7680() bool`

HasVar7680 returns a boolean if a field has been set.

### GetVar7681

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7681() string`

GetVar7681 returns the Var7681 field if non-nil, zero value otherwise.

### GetVar7681Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7681Ok() (*string, bool)`

GetVar7681Ok returns a tuple with the Var7681 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7681

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7681(v string)`

SetVar7681 sets Var7681 field to given value.

### HasVar7681

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7681() bool`

HasVar7681 returns a boolean if a field has been set.

### GetVar7682

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7682() string`

GetVar7682 returns the Var7682 field if non-nil, zero value otherwise.

### GetVar7682Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7682Ok() (*string, bool)`

GetVar7682Ok returns a tuple with the Var7682 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7682

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7682(v string)`

SetVar7682 sets Var7682 field to given value.

### HasVar7682

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7682() bool`

HasVar7682 returns a boolean if a field has been set.

### GetVar7683

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7683() string`

GetVar7683 returns the Var7683 field if non-nil, zero value otherwise.

### GetVar7683Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7683Ok() (*string, bool)`

GetVar7683Ok returns a tuple with the Var7683 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7683

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7683(v string)`

SetVar7683 sets Var7683 field to given value.

### HasVar7683

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7683() bool`

HasVar7683 returns a boolean if a field has been set.

### GetVar7684

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7684() string`

GetVar7684 returns the Var7684 field if non-nil, zero value otherwise.

### GetVar7684Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7684Ok() (*string, bool)`

GetVar7684Ok returns a tuple with the Var7684 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7684

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7684(v string)`

SetVar7684 sets Var7684 field to given value.

### HasVar7684

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7684() bool`

HasVar7684 returns a boolean if a field has been set.

### GetVar7685

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7685() string`

GetVar7685 returns the Var7685 field if non-nil, zero value otherwise.

### GetVar7685Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7685Ok() (*string, bool)`

GetVar7685Ok returns a tuple with the Var7685 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7685

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7685(v string)`

SetVar7685 sets Var7685 field to given value.

### HasVar7685

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7685() bool`

HasVar7685 returns a boolean if a field has been set.

### GetVar7686

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7686() string`

GetVar7686 returns the Var7686 field if non-nil, zero value otherwise.

### GetVar7686Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7686Ok() (*string, bool)`

GetVar7686Ok returns a tuple with the Var7686 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7686

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7686(v string)`

SetVar7686 sets Var7686 field to given value.

### HasVar7686

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7686() bool`

HasVar7686 returns a boolean if a field has been set.

### GetVar7687

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7687() string`

GetVar7687 returns the Var7687 field if non-nil, zero value otherwise.

### GetVar7687Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7687Ok() (*string, bool)`

GetVar7687Ok returns a tuple with the Var7687 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7687

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7687(v string)`

SetVar7687 sets Var7687 field to given value.

### HasVar7687

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7687() bool`

HasVar7687 returns a boolean if a field has been set.

### GetVar7688

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7688() string`

GetVar7688 returns the Var7688 field if non-nil, zero value otherwise.

### GetVar7688Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7688Ok() (*string, bool)`

GetVar7688Ok returns a tuple with the Var7688 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7688

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7688(v string)`

SetVar7688 sets Var7688 field to given value.

### HasVar7688

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7688() bool`

HasVar7688 returns a boolean if a field has been set.

### GetVar7689

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7689() string`

GetVar7689 returns the Var7689 field if non-nil, zero value otherwise.

### GetVar7689Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7689Ok() (*string, bool)`

GetVar7689Ok returns a tuple with the Var7689 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7689

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7689(v string)`

SetVar7689 sets Var7689 field to given value.

### HasVar7689

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7689() bool`

HasVar7689 returns a boolean if a field has been set.

### GetVar7690

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7690() string`

GetVar7690 returns the Var7690 field if non-nil, zero value otherwise.

### GetVar7690Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7690Ok() (*string, bool)`

GetVar7690Ok returns a tuple with the Var7690 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7690

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7690(v string)`

SetVar7690 sets Var7690 field to given value.

### HasVar7690

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7690() bool`

HasVar7690 returns a boolean if a field has been set.

### GetVar7694

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7694() string`

GetVar7694 returns the Var7694 field if non-nil, zero value otherwise.

### GetVar7694Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7694Ok() (*string, bool)`

GetVar7694Ok returns a tuple with the Var7694 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7694

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7694(v string)`

SetVar7694 sets Var7694 field to given value.

### HasVar7694

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7694() bool`

HasVar7694 returns a boolean if a field has been set.

### GetVar7695

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7695() string`

GetVar7695 returns the Var7695 field if non-nil, zero value otherwise.

### GetVar7695Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7695Ok() (*string, bool)`

GetVar7695Ok returns a tuple with the Var7695 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7695

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7695(v string)`

SetVar7695 sets Var7695 field to given value.

### HasVar7695

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7695() bool`

HasVar7695 returns a boolean if a field has been set.

### GetVar7696

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7696() string`

GetVar7696 returns the Var7696 field if non-nil, zero value otherwise.

### GetVar7696Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7696Ok() (*string, bool)`

GetVar7696Ok returns a tuple with the Var7696 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7696

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7696(v string)`

SetVar7696 sets Var7696 field to given value.

### HasVar7696

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7696() bool`

HasVar7696 returns a boolean if a field has been set.

### GetVar7697

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7697() string`

GetVar7697 returns the Var7697 field if non-nil, zero value otherwise.

### GetVar7697Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7697Ok() (*string, bool)`

GetVar7697Ok returns a tuple with the Var7697 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7697

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7697(v string)`

SetVar7697 sets Var7697 field to given value.

### HasVar7697

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7697() bool`

HasVar7697 returns a boolean if a field has been set.

### GetVar7698

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7698() string`

GetVar7698 returns the Var7698 field if non-nil, zero value otherwise.

### GetVar7698Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7698Ok() (*string, bool)`

GetVar7698Ok returns a tuple with the Var7698 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7698

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7698(v string)`

SetVar7698 sets Var7698 field to given value.

### HasVar7698

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7698() bool`

HasVar7698 returns a boolean if a field has been set.

### GetVar7699

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7699() string`

GetVar7699 returns the Var7699 field if non-nil, zero value otherwise.

### GetVar7699Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7699Ok() (*string, bool)`

GetVar7699Ok returns a tuple with the Var7699 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7699

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7699(v string)`

SetVar7699 sets Var7699 field to given value.

### HasVar7699

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7699() bool`

HasVar7699 returns a boolean if a field has been set.

### GetVar7700

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7700() string`

GetVar7700 returns the Var7700 field if non-nil, zero value otherwise.

### GetVar7700Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7700Ok() (*string, bool)`

GetVar7700Ok returns a tuple with the Var7700 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7700

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7700(v string)`

SetVar7700 sets Var7700 field to given value.

### HasVar7700

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7700() bool`

HasVar7700 returns a boolean if a field has been set.

### GetVar7702

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7702() string`

GetVar7702 returns the Var7702 field if non-nil, zero value otherwise.

### GetVar7702Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7702Ok() (*string, bool)`

GetVar7702Ok returns a tuple with the Var7702 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7702

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7702(v string)`

SetVar7702 sets Var7702 field to given value.

### HasVar7702

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7702() bool`

HasVar7702 returns a boolean if a field has been set.

### GetVar7703

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7703() string`

GetVar7703 returns the Var7703 field if non-nil, zero value otherwise.

### GetVar7703Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7703Ok() (*string, bool)`

GetVar7703Ok returns a tuple with the Var7703 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7703

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7703(v string)`

SetVar7703 sets Var7703 field to given value.

### HasVar7703

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7703() bool`

HasVar7703 returns a boolean if a field has been set.

### GetVar7704

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7704() string`

GetVar7704 returns the Var7704 field if non-nil, zero value otherwise.

### GetVar7704Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7704Ok() (*string, bool)`

GetVar7704Ok returns a tuple with the Var7704 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7704

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7704(v string)`

SetVar7704 sets Var7704 field to given value.

### HasVar7704

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7704() bool`

HasVar7704 returns a boolean if a field has been set.

### GetVar7705

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7705() string`

GetVar7705 returns the Var7705 field if non-nil, zero value otherwise.

### GetVar7705Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7705Ok() (*string, bool)`

GetVar7705Ok returns a tuple with the Var7705 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7705

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7705(v string)`

SetVar7705 sets Var7705 field to given value.

### HasVar7705

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7705() bool`

HasVar7705 returns a boolean if a field has been set.

### GetVar7706

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7706() string`

GetVar7706 returns the Var7706 field if non-nil, zero value otherwise.

### GetVar7706Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7706Ok() (*string, bool)`

GetVar7706Ok returns a tuple with the Var7706 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7706

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7706(v string)`

SetVar7706 sets Var7706 field to given value.

### HasVar7706

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7706() bool`

HasVar7706 returns a boolean if a field has been set.

### GetVar7707

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7707() string`

GetVar7707 returns the Var7707 field if non-nil, zero value otherwise.

### GetVar7707Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7707Ok() (*string, bool)`

GetVar7707Ok returns a tuple with the Var7707 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7707

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7707(v string)`

SetVar7707 sets Var7707 field to given value.

### HasVar7707

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7707() bool`

HasVar7707 returns a boolean if a field has been set.

### GetVar7708

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7708() string`

GetVar7708 returns the Var7708 field if non-nil, zero value otherwise.

### GetVar7708Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7708Ok() (*string, bool)`

GetVar7708Ok returns a tuple with the Var7708 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7708

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7708(v string)`

SetVar7708 sets Var7708 field to given value.

### HasVar7708

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7708() bool`

HasVar7708 returns a boolean if a field has been set.

### GetVar7714

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7714() string`

GetVar7714 returns the Var7714 field if non-nil, zero value otherwise.

### GetVar7714Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7714Ok() (*string, bool)`

GetVar7714Ok returns a tuple with the Var7714 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7714

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7714(v string)`

SetVar7714 sets Var7714 field to given value.

### HasVar7714

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7714() bool`

HasVar7714 returns a boolean if a field has been set.

### GetVar7715

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7715() string`

GetVar7715 returns the Var7715 field if non-nil, zero value otherwise.

### GetVar7715Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7715Ok() (*string, bool)`

GetVar7715Ok returns a tuple with the Var7715 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7715

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7715(v string)`

SetVar7715 sets Var7715 field to given value.

### HasVar7715

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7715() bool`

HasVar7715 returns a boolean if a field has been set.

### GetVar7718

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7718() string`

GetVar7718 returns the Var7718 field if non-nil, zero value otherwise.

### GetVar7718Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7718Ok() (*string, bool)`

GetVar7718Ok returns a tuple with the Var7718 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7718

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7718(v string)`

SetVar7718 sets Var7718 field to given value.

### HasVar7718

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7718() bool`

HasVar7718 returns a boolean if a field has been set.

### GetVar7720

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7720() string`

GetVar7720 returns the Var7720 field if non-nil, zero value otherwise.

### GetVar7720Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7720Ok() (*string, bool)`

GetVar7720Ok returns a tuple with the Var7720 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7720

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7720(v string)`

SetVar7720 sets Var7720 field to given value.

### HasVar7720

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7720() bool`

HasVar7720 returns a boolean if a field has been set.

### GetVar7741

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7741() string`

GetVar7741 returns the Var7741 field if non-nil, zero value otherwise.

### GetVar7741Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7741Ok() (*string, bool)`

GetVar7741Ok returns a tuple with the Var7741 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7741

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7741(v string)`

SetVar7741 sets Var7741 field to given value.

### HasVar7741

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7741() bool`

HasVar7741 returns a boolean if a field has been set.

### GetVar7762

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7762() string`

GetVar7762 returns the Var7762 field if non-nil, zero value otherwise.

### GetVar7762Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7762Ok() (*string, bool)`

GetVar7762Ok returns a tuple with the Var7762 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7762

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7762(v string)`

SetVar7762 sets Var7762 field to given value.

### HasVar7762

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7762() bool`

HasVar7762 returns a boolean if a field has been set.

### GetVar7768

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7768() string`

GetVar7768 returns the Var7768 field if non-nil, zero value otherwise.

### GetVar7768Ok

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar7768Ok() (*string, bool)`

GetVar7768Ok returns a tuple with the Var7768 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar7768

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar7768(v string)`

SetVar7768 sets Var7768 field to given value.

### HasVar7768

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar7768() bool`

HasVar7768 returns a boolean if a field has been set.

### GetServerId

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetConid

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetConid() int32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetConidOk() (*int32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetConid(v int32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetUpdated

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVar87RawDeprecated

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar87RawDeprecated() string`

GetVar87RawDeprecated returns the Var87RawDeprecated field if non-nil, zero value otherwise.

### GetVar87RawDeprecatedOk

`func (o *IserverMarketdataSnapshotGet200ResponseInner) GetVar87RawDeprecatedOk() (*string, bool)`

GetVar87RawDeprecatedOk returns a tuple with the Var87RawDeprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar87RawDeprecated

`func (o *IserverMarketdataSnapshotGet200ResponseInner) SetVar87RawDeprecated(v string)`

SetVar87RawDeprecated sets Var87RawDeprecated field to given value.

### HasVar87RawDeprecated

`func (o *IserverMarketdataSnapshotGet200ResponseInner) HasVar87RawDeprecated() bool`

HasVar87RawDeprecated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


