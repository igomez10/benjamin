# market-data
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **31** | **String** | Last Price - The last price at which the contract traded. May contain one of the following prefixes:   * C - Previous day&#39;s closing price.   * H - Trading has halted.  | [optional] [default to null] |
| **70** | **BigDecimal** | High - Current day high price | [optional] [default to null] |
| **71** | **BigDecimal** | Low - Current day low price | [optional] [default to null] |
| **82** | **String** | Change - The difference between the last price and the close on the previous trading day | [optional] [default to null] |
| **83** | **BigDecimal** | Change % - The difference between the last price and the close on the previous trading day in percentage. | [optional] [default to null] |
| **84** | **String** | Bid Price - The highest-priced bid on the contract. | [optional] [default to null] |
| **85** | **String** | Ask Size - The number of contracts or shares offered at the ask price. For US stocks, the number displayed is divided by 100. | [optional] [default to null] |
| **86** | **String** | Ask Price - The lowest-priced offer on the contract. | [optional] [default to null] |
| **87** | **String** | Volume - Volume for the day, formatted with &#39;K&#39; for thousands or &#39;M&#39; for millions. For higher precision volume refer to field 7762. | [optional] [default to null] |
| **88** | **String** | Bid Size - The number of contracts or shares bid for at the bid price. For US stocks, the number displayed is divided by 100. | [optional] [default to null] |
| **6509** | **String** | Market Data Availability. The field may contain three chars. First char defines: R &#x3D; RealTime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed, N &#x3D; Not Subscribed. Second char defines: P &#x3D; Snapshot, p &#x3D; Consolidated. Third char defines: B &#x3D; Book   * RealTime - Data is relayed back in real time without delay, market data subscription(s) are required.   * Delayed - Data is relayed back 15-20 min delayed.   * Frozen - Last recorded data at market close, relayed back in real time.   * Frozen Delayed - Last recorded data at market close, relayed back delayed.   * Not Subscribed - User does not have the required market data subscription(s) to relay back either real time or delayed data.   * Snapshot - Snapshot request is available for contract.   * Consolidated - Market data is aggregated across multiple exchanges or venues.   * Book - Top of the book data is available for contract.  | [optional] [default to null] |
| **7057** | **String** | Ask Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY | [optional] [default to null] |
| **7058** | **String** | Last Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY | [optional] [default to null] |
| **7059** | **BigDecimal** | Last Size - The number of unites traded at the last price | [optional] [default to null] |
| **7068** | **String** | Bid Exch - Displays the exchange(s) offering the SMART price. A&#x3D;AMEX, C&#x3D;CBOE, I&#x3D;ISE, X&#x3D;PHLX, N&#x3D;PSE, B&#x3D;BOX, Q&#x3D;NASDAQOM, Z&#x3D;BATS, W&#x3D;CBOE2, T&#x3D;NASDAQBX, M&#x3D;MIAX, H&#x3D;GEMINI, E&#x3D;EDGX, J&#x3D;MERCURY | [optional] [default to null] |
| **7195** | **String** | IV Rank | [optional] [default to null] |
| **7196** | **String** | IV Rank | [optional] [default to null] |
| **7197** | **String** | IV Rank | [optional] [default to null] |
| **7198** | **String** | IV Percentile | [optional] [default to null] |
| **7199** | **String** | IV Percentile | [optional] [default to null] |
| **7200** | **String** | IV Percentile | [optional] [default to null] |
| **7201** | **String** | IV High Low | [optional] [default to null] |
| **7202** | **String** | IV High Low | [optional] [default to null] |
| **7203** | **String** | IV High Low | [optional] [default to null] |
| **7204** | **String** | IV High Low | [optional] [default to null] |
| **7205** | **String** | IV High Low | [optional] [default to null] |
| **7206** | **String** | IV High Low | [optional] [default to null] |
| **7207** | **String** | HV Rank | [optional] [default to null] |
| **7208** | **String** | HV Rank | [optional] [default to null] |
| **7209** | **String** | HV Rank | [optional] [default to null] |
| **7210** | **String** | HV Percentile | [optional] [default to null] |
| **7211** | **String** | HV Percentile | [optional] [default to null] |
| **7212** | **String** | HV Percentile | [optional] [default to null] |
| **7245** | **String** | HV High Low | [optional] [default to null] |
| **7246** | **String** | HV High Low | [optional] [default to null] |
| **7247** | **String** | HV High Low | [optional] [default to null] |
| **7248** | **String** | HV High Low | [optional] [default to null] |
| **7249** | **String** | HV High Low | [optional] [default to null] |
| **7263** | **String** | HV High Low | [optional] [default to null] |
| **7264** | **String** | ESG | [optional] [default to null] |
| **7265** | **String** | ESG | [optional] [default to null] |
| **7266** | **String** | ESG | [optional] [default to null] |
| **7267** | **String** | ESG | [optional] [default to null] |
| **7268** | **String** | ESG | [optional] [default to null] |
| **7269** | **String** | ESG | [optional] [default to null] |
| **7271** | **String** | ESG | [optional] [default to null] |
| **7272** | **String** | ESG | [optional] [default to null] |
| **7273** | **String** | ESG | [optional] [default to null] |
| **7274** | **String** | ESG | [optional] [default to null] |
| **7275** | **String** | ESG | [optional] [default to null] |
| **7276** | **String** | ESG | [optional] [default to null] |
| **7277** | **String** | ESG | [optional] [default to null] |
| **7282** | **String** | Average Volume - The average daily trading volume over 90 days. | [optional] [default to null] |
| **7283** | **String** | Option Implied Vol. % - A prediction of how volatile an underlying will be in the future. At the market volatility estimated for a maturity thirty calendar days forward of the current trading day, and based on option prices from two consecutive expiration months.  | [optional] [default to null] |
| **7284** | **String** | Historic Volume (30d) | [optional] [default to null] |
| **7286** | **BigDecimal** | Dividend Amount - Displays the amount of the next dividend. | [optional] [default to null] |
| **7287** | **String** | Dividend Yield % - This value is the toal of the expected dividend payments over the next twelve months per share divided by the Current Price and is expressed as a percentage. For derivatives, this displays the total of the expected dividend payments over the expiry date.  | [optional] [default to null] |
| **7288** | **String** | Ex-date of the dividend | [optional] [default to null] |
| **7289** | **String** | Market Cap | [optional] [default to null] |
| **7290** | **String** | P/E | [optional] [default to null] |
| **7293** | **String** | 52 Week High - The highest price for the past 52 weeks. | [optional] [default to null] |
| **7294** | **String** | 52 Week Low - The lowest price for the past 52 weeks. | [optional] [default to null] |
| **7295** | **BigDecimal** | Open - Today&#39;s opening price. | [optional] [default to null] |
| **7296** | **BigDecimal** | Close - Today&#39;s closing price. | [optional] [default to null] |
| **7331** | **String** | Reuters Fundamentals | [optional] [default to null] |
| **7370** | **String** | ESG | [optional] [default to null] |
| **7371** | **String** | ESG | [optional] [default to null] |
| **7372** | **String** | ESG | [optional] [default to null] |
| **7635** | **String** | Mark - The mark price is, the ask price if ask is less than last price, the bid price if bid is more than the last price, otherwise it&#39;s equal to last price | [optional] [default to null] |
| **7636** | **BigDecimal** | shortable invetory | [optional] [default to null] |
| **7637** | **String** | Fee rebate rate | [optional] [default to null] |
| **7644** | **String** | Shortable - Describes the level of difficulty with which the security can be sold short. | [optional] [default to null] |
| **7674** | **String** | EMA(200) - Exponential moving average (N&#x3D;200). | [optional] [default to null] |
| **7675** | **String** | EMA(100) - Exponential moving average (N&#x3D;100). | [optional] [default to null] |
| **7676** | **String** | EMA(50) - Exponential moving average (N&#x3D;50). | [optional] [default to null] |
| **7677** | **String** | EMA(20) - Exponential moving average (N&#x3D;20). | [optional] [default to null] |
| **7681** | **String** | Price/EMA(20) - Price to Exponential moving average (N&#x3D;20) ratio -1, displayed in percents. | [optional] [default to null] |
| **7698** | **String** | Last Yield - Implied yield of the bond if it is purchased at the current last price. Last yield is calculated using the Last price on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] [default to null] |
| **7699** | **String** | Bid Yield - Implied yield of the bond if it is purchased at the current bid price. Bid yield is calculated using the Ask on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] [default to null] |
| **7718** | **String** | Beta - Beta is against standard index. | [optional] [default to null] |
| **7720** | **String** | Ask Yield - Implied yield of the bond if it is purchased at the current offer. Ask yield is calculated using the Bid on all possible call dates. It is assumed that prepayment occurs if the bond has call or put provisions and the issuer can offer a lower coupon rate based on current market rates. The yield to worst will be the lowest of the yield to maturity or yield to call (if the bond has prepayment provisions). Yield to worse may be the same as yield to maturity but never higher.  | [optional] [default to null] |
| **7743** | **String** | Reuters Fundamentals | [optional] [default to null] |
| **7761** | **String** | ESG | [optional] [default to null] |
| **7992** | **String** | 26 Week High - The highest price for the past 26 weeks. | [optional] [default to null] |
| **7993** | **String** | 26 Week Low - The lowest price for the past 26 weeks. | [optional] [default to null] |
| **7994** | **String** | 13 Week High - The highest price for the past 13 weeks. | [optional] [default to null] |
| **7995** | **String** | 13 Week Low - The lowest price for the past 13 weeks. | [optional] [default to null] |
| **conid** | **Integer** | IBKR Contract identifier | [optional] [default to null] |
| **minTick** | **BigDecimal** | minimum price increment | [optional] [default to null] |
| **BboExchange** | **String** | Color for Best Bid/Offer Exchange in hex code | [optional] [default to null] |
| **HasDelayed** | **Boolean** | If market data field values return delayed | [optional] [default to null] |
| **sizeMinTick** | **Integer** | minimum size increment | [optional] [default to null] |
| **BestEligible** | **Integer** |  | [optional] [default to null] |
| **BestBidExch** | **Integer** |  | [optional] [default to null] |
| **BestAskExch** | **Integer** |  | [optional] [default to null] |
| **PreOpenBid** | **Integer** |  | [optional] [default to null] |
| **LastAttribs** | **Integer** |  | [optional] [default to null] |
| **TimestampBase** | **Integer** | Base time stamp for last update in format YYYYMMDD | [optional] [default to null] |
| **TimestampDelta** | **Integer** |  | [optional] [default to null] |
| **LastExch** | **Integer** |  | [optional] [default to null] |
| **CloseAttribs** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

