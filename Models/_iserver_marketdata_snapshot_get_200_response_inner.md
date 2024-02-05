# _iserver_marketdata_snapshot_get_200_response_inner
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **31** | **String** | Last Price - The last price at which the contract traded. May contain one of the following prefixes:   * C - Previous day&#39;s closing price.   * H - Trading has halted.  | [optional] [default to null] |
| **55** | **String** | Symbol | [optional] [default to null] |
| **58** | **String** | Text | [optional] [default to null] |
| **70** | **String** | High - Current day high price | [optional] [default to null] |
| **71** | **String** | Low - Current day low price | [optional] [default to null] |
| **73** | **String** | Market Value - The current market value of  your position in the security. Market Value is calculated with real time market data (even when not subscribed to market data). | [optional] [default to null] |
| **74** | **String** | Avg Price - The average price of the position. | [optional] [default to null] |
| **75** | **String** | Unrealized PnL - Unrealized profit or loss. Unrealized PnL is calculated with real time market data (even when not subscribed to market data). | [optional] [default to null] |
| **76** | **String** | Formatted position | [optional] [default to null] |
| **77** | **String** | Formatted Unrealized PnL | [optional] [default to null] |
| **78** | **String** | Daily PnL - Your profit or loss of the day since prior close. Daily PnL is calculated with real time market data (even when not subscribed to market data). | [optional] [default to null] |
| **79** | **String** | Realized PnL - Realized profit or loss. Realized PnL is calculated with real time market data (even when not subscribed to market data). | [optional] [default to null] |
| **80** | **String** | Unrealized PnL % - Unrealized profit or loss expressed in percentage. | [optional] [default to null] |
| **82** | **String** | Change - The difference between the last price and the close on the previous trading day | [optional] [default to null] |
| **83** | **String** | Change % - The difference between the last price and the close on the previous trading day in percentage. | [optional] [default to null] |
| **84** | **String** | Bid Price - The highest-priced bid on the contract. | [optional] [default to null] |
| **85** | **String** | Ask Size - The number of contracts or shares offered at the ask price. For US stocks, the number displayed is divided by 100. | [optional] [default to null] |
| **86** | **String** | Ask Price - The lowest-priced offer on the contract. | [optional] [default to null] |
| **87** | **String** | Volume - Volume for the day, formatted with &#39;K&#39; for thousands or &#39;M&#39; for millions. For higher precision volume refer to field 7762. | [optional] [default to null] |
| **88** | **String** | Bid Size - The number of contracts or shares bid for at the bid price. For US stocks, the number displayed is divided by 100. | [optional] [default to null] |
| **6004** | **String** | Exchange | [optional] [default to null] |
| **6008** | **Integer** | Conid - Contract identifier from IBKR&#39;s database. | [optional] [default to null] |
| **6070** | **String** | SecType - The asset class of the instrument. | [optional] [default to null] |
| **6072** | **String** | Months | [optional] [default to null] |
| **6073** | **String** | Regular Expiry | [optional] [default to null] |
| **6119** | **String** | Marker for market data delivery method (similar to request id) | [optional] [default to null] |
| **6457** | **Integer** | Underlying Conid. Use /trsrv/secdef to get more information about the security | [optional] [default to null] |
| **6508** | **String** | Service Params. | [optional] [default to null] |
| **6509** | **String** | Market Data Availability. The field may contain three chars. First char defines: R &#x3D; RealTime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed, N &#x3D; Not Subscribed. Second char defines: P &#x3D; Snapshot, p &#x3D; Consolidated. Third char defines: B &#x3D; Book   * RealTime - Data is relayed back in real time without delay, market data subscription(s) are required.   * Delayed - Data is relayed back 15-20 min delayed.    * Frozen - Last recorded data at market close, relayed back in real time.   * Frozen Delayed - Last recorded data at market close, relayed back delayed.   * Not Subscribed - User does not have the required market data subscription(s) to relay back either real time or delayed data.   * Snapshot - Snapshot request is available for contract.   * Consolidated - Market data is aggregated across multiple exchanges or venues.   * Book - Top of the book data is available for contract.  | [optional] [default to null] |
| **7051** | **String** | Company name | [optional] [default to null] |
| **7057** | **String** | Ask Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY  | [optional] [default to null] |
| **7058** | **String** | Last Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY  | [optional] [default to null] |
| **7059** | **String** | Last Size - The number of unites traded at the last price | [optional] [default to null] |
| **7068** | **String** | Bid Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY  | [optional] [default to null] |
| **7084** | **String** | Implied Vol./Hist. Vol % - The ratio of the implied volatility over the historical volatility, expressed as a percentage. | [optional] [default to null] |
| **7085** | **String** | Put/Call Interest - Put option open interest/call option open interest for the trading day. | [optional] [default to null] |
| **7086** | **String** | Put/Call Volume - Put option volume/call option volume for the trading day. | [optional] [default to null] |
| **7087** | **String** | Hist. Vol. % - 30-day real-time historical volatility. | [optional] [default to null] |
| **7088** | **String** | Hist. Vol. Close % - Shows the historical volatility based on previous close price. | [optional] [default to null] |
| **7089** | **String** | Opt. Volume - Option Volume | [optional] [default to null] |
| **7094** | **String** | Conid + Exchange | [optional] [default to null] |
| **7184** | **String** | canBeTraded - If contract is a trade-able instrument. Returns 1(true) or 0(false). | [optional] [default to null] |
| **7219** | **String** | Contract Description | [optional] [default to null] |
| **7220** | **String** | Contract Description | [optional] [default to null] |
| **7221** | **String** | Listing Exchange | [optional] [default to null] |
| **7280** | **String** | Industry - Displays the type of industry under which the underlying company can be categorized. | [optional] [default to null] |
| **7281** | **String** | Category - Displays a more detailed level of description within the industry under which the underlying company can be categorized. | [optional] [default to null] |
| **7282** | **String** | Average Volume - The average daily trading volume over 90 days. | [optional] [default to null] |
| **7283** | **String** | Option Implied Vol. % - A prediction of how volatile an underlying will be in the future. At the market volatility estimated for a maturity thirty calendar days forward of the current trading day, and based on option prices from two consecutive expiration months. To query the Implied Vol. % of a specific strike refer to field 7633.  | [optional] [default to null] |
| **7284** | **String** | Historic Volume (30d) | [optional] [default to null] |
| **7285** | **String** | Put/Call Ratio | [optional] [default to null] |
| **7286** | **String** | Dividend Amount - Displays the amount of the next dividend. | [optional] [default to null] |
| **7287** | **String** | Dividend Yield % - This value is the toal of the expected dividend payments over the next twelve months per share divided by the Current Price and is expressed as a percentage. For derivatives, this displays the total of the expected dividend payments over the expiry date.  | [optional] [default to null] |
| **7288** | **String** | Ex-date of the dividend | [optional] [default to null] |
| **7289** | **String** | Market Cap | [optional] [default to null] |
| **7290** | **String** | P/E | [optional] [default to null] |
| **7291** | **String** | EPS | [optional] [default to null] |
| **7292** | **String** | Cost Basis - Your current position in this security multiplied by the average price and multiplier. | [optional] [default to null] |
| **7293** | **String** | 52 Week High - The highest price for the past 52 weeks. | [optional] [default to null] |
| **7294** | **String** | 52 Week Low - The lowest price for the past 52 weeks. | [optional] [default to null] |
| **7295** | **String** | Open - Today&#39;s opening price. | [optional] [default to null] |
| **7296** | **String** | Close - Today&#39;s closing price. | [optional] [default to null] |
| **7308** | **String** | Delta - The ratio of the change in the price of the option to the corresponding change in the price of the underlying. | [optional] [default to null] |
| **7309** | **String** | Gamma - The rate of change for the delta with respect to the underlying asset&#39;s price. | [optional] [default to null] |
| **7310** | **String** | Theta - A measure of the rate of decline the value of an option due to the passage of time. | [optional] [default to null] |
| **7311** | **String** | Vega - The amount that the price of an option changes compared to a 1% change in the volatility. | [optional] [default to null] |
| **7607** | **String** | Opt. Volume Change % - Today&#39;s option volume as a percentage of the average option volume. | [optional] [default to null] |
| **7633** | **String** | Implied Vol. % - The implied volatility for the specific strike of the option in percentage. To query the Option Implied Vol. % from the underlying refer to field 7283.  | [optional] [default to null] |
| **7635** | **String** | Mark - The mark price is, the ask price if ask is less than last price, the bid price if bid is more than the last price, otherwise it&#39;s equal to last price. | [optional] [default to null] |
| **7636** | **String** | Shortable Shares - Number of shares available for shorting. | [optional] [default to null] |
| **7637** | **String** | Fee Rate - Interest rate charged on borrowed shares. | [optional] [default to null] |
| **7638** | **String** | Option Open Interest | [optional] [default to null] |
| **7639** | **String** | % of Mark Value - Displays the market value of the contract as a percentage of the total market value of the account. Mark Value is calculated with real time market data (even when not subscribed to market data).  | [optional] [default to null] |
| **7644** | **String** | Shortable - Describes the level of difficulty with which the security can be sold short. | [optional] [default to null] |
| **7655** | **String** | Morningstar Rating - Displays Morningstar Rating provided value. Requires [Morningstar](https://www.interactivebrokers.com/en/index.php?f&#x3D;14262) subscription. | [optional] [default to null] |
| **7671** | **String** | Dividends - This value is the total of the expected dividend payments over the next twelve months per share. | [optional] [default to null] |
| **7672** | **String** | Dividends TTM - This value is the total of the expected dividend payments over the last twelve months per share. | [optional] [default to null] |
| **7674** | **String** | EMA(200) - Exponential moving average (N&#x3D;200). | [optional] [default to null] |
| **7675** | **String** | EMA(100) - Exponential moving average (N&#x3D;100). | [optional] [default to null] |
| **7676** | **String** | EMA(50) - Exponential moving average (N&#x3D;50). | [optional] [default to null] |
| **7677** | **String** | EMA(20) - Exponential moving average (N&#x3D;20). | [optional] [default to null] |
| **7678** | **String** | Price/EMA(200) - Price to Exponential moving average (N&#x3D;200) ratio -1, displayed in percents. | [optional] [default to null] |
| **7679** | **String** | Price/EMA(100) - Price to Exponential moving average (N&#x3D;100) ratio -1, displayed in percents. | [optional] [default to null] |
| **7680** | **String** | Price/EMA(50) - Price to Exponential moving average (N&#x3D;50) ratio -1, displayed in percents. | [optional] [default to null] |
| **7681** | **String** | Price/EMA(20) - Price to Exponential moving average (N&#x3D;20) ratio -1, displayed in percents. | [optional] [default to null] |
| **7682** | **String** | Change Since Open - The difference between the last price and the open price. | [optional] [default to null] |
| **7683** | **String** | Upcoming Event - Shows the next major company event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] [default to null] |
| **7684** | **String** | Upcoming Event Date - The date of the next major company event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] [default to null] |
| **7685** | **String** | Upcoming Analyst Meeting - The date and time of the next scheduled analyst meeting. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] [default to null] |
| **7686** | **String** | Upcoming Earnings - The date and time of the next scheduled earnings/earnings call event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] [default to null] |
| **7687** | **String** | Upcoming Misc Event - The date and time of the next shareholder meeting, presentation or other event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] [default to null] |
| **7688** | **String** | Recent Analyst Meeting - The date and time of the most recent analyst meeting. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] [default to null] |
| **7689** | **String** | Recent Earnings - The date and time of the most recent earnings/earning call event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] [default to null] |
| **7690** | **String** | Recent Misc Event - The date and time of the most recent shareholder meeting, presentation or other event. Requires [Wall Street Horizon](https://www.interactivebrokers.com/en/index.php?f&#x3D;24674) subscription. | [optional] [default to null] |
| **7694** | **String** | Probability of Max Return - Customer implied probability of maximum potential gain. | [optional] [default to null] |
| **7695** | **String** | Break Even - Break even points | [optional] [default to null] |
| **7696** | **String** | SPX Delta - Beta Weighted Delta is calculated using the formula; Delta x dollar adjusted beta, where adjusted beta is adjusted by the ratio of the close price. | [optional] [default to null] |
| **7697** | **String** | Futures Open Interest - Total number of outstanding futures contracts | [optional] [default to null] |
| **7698** | **String** | Last Yield - Implied yield of the bond if it is purchased at the current last price. Last yield is calculated using the Last price on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] [default to null] |
| **7699** | **String** | Bid Yield - Implied yield of the bond if it is purchased at the current bid price. Bid yield is calculated using the Ask on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] [default to null] |
| **7700** | **String** | Probability of Max Return - Customer implied probability of maximum potential gain. | [optional] [default to null] |
| **7702** | **String** | Probability of Max Loss - Customer implied probability of maximum potential loss. | [optional] [default to null] |
| **7703** | **String** | Profit Probability - Customer implied probability of any gain. | [optional] [default to null] |
| **7704** | **String** | Organization Type | [optional] [default to null] |
| **7705** | **String** | Debt Class | [optional] [default to null] |
| **7706** | **String** | Ratings - Ratings issued for bond contract. | [optional] [default to null] |
| **7707** | **String** | Bond State Code | [optional] [default to null] |
| **7708** | **String** | Bond Type | [optional] [default to null] |
| **7714** | **String** | Last Trading Date | [optional] [default to null] |
| **7715** | **String** | Issue Date | [optional] [default to null] |
| **7718** | **String** | Beta - Beta is against standard index. | [optional] [default to null] |
| **7720** | **String** | Ask Yield - Implied yield of the bond if it is purchased at the current offer. Ask yield is calculated using the Bid on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] [default to null] |
| **7741** | **String** | Prior Close - Yesterday&#39;s closing price | [optional] [default to null] |
| **7762** | **String** | Volume Long - High precision volume for the day. For formatted volume refer to field 87. | [optional] [default to null] |
| **7768** | **String** | hasTradingPermissions - if user has trading permissions for specified contract. Returns 1(true) or 0(false). | [optional] [default to null] |
| **server\_id** | **String** |  | [optional] [default to null] |
| **conid** | **Integer** |  | [optional] [default to null] |
| **\_updated** | **Integer** |  | [optional] [default to null] |
| **87\_raw (deprecated)** | **String** | Raw Volume - Volume for the day, provided in long form without formatted with K/M. This field value is deprecated, for high precision volume refer to field 7762. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

