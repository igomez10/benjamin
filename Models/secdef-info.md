# secdef-info
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **conid** | **BigDecimal** | IBKR contract identifier | [optional] [default to null] |
| **symbol** | **String** | Underlying symbol | [optional] [default to null] |
| **secType** | **String** | Security type | [optional] [default to null] |
| **exchange** | **String** | Primary Exchange, Routing or Trading Venue | [optional] [default to null] |
| **listingExchange** | **String** | Main Trading Venue | [optional] [default to null] |
| **right** | **String** | Put or Call of the option. C &#x3D; Call Option, P &#x3D; Put Option | [optional] [default to null] |
| **strike** | **BigDecimal** | Set price at which a derivative contract can be bought or sold. The strike price also known as exercise price. | [optional] [default to null] |
| **currency** | **String** | Currency the contract trades in | [optional] [default to null] |
| **cusip** | **String** | Committee on Uniform Securities Identification Procedures number | [optional] [default to null] |
| **coupon** | **String** | Annual interest rate paid on a bond | [optional] [default to null] |
| **desc1** | **String** | Currency pairs for Forex e.g. EUR.AUD, EUR.CAD, EUR.CHF etc. | [optional] [default to null] |
| **desc2** | **String** | Formatted expiration, strike and right | [optional] [default to null] |
| **maturityDate** | **BigDecimal** | Format YYYYMMDD, the date on which the underlying transaction settles if the option is exercised | [optional] [default to null] |
| **multiplier** | **String** | Multiplier for total premium paid or received for derivative contract. | [optional] [default to null] |
| **tradingClass** | **String** | Designation of the contract. | [optional] [default to null] |
| **validExchanges** | **String** | Comma separated list of exchanges or trading venues. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

