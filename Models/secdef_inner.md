# secdef_inner
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **conid** | **Integer** | IBKR contract identifier. | [optional] [default to null] |
| **currency** | **String** | Currency contract trades in. | [optional] [default to null] |
| **crossCurrency** | **Boolean** | Defines if a derivative contract has a different currency. | [optional] [default to null] |
| **time** | **Integer** |  | [optional] [default to null] |
| **chineseName** | **String** | HTML encoded company description in Chinese. | [optional] [default to null] |
| **allExchanges** | **String** | List of exchanges and venues contract trades. | [optional] [default to null] |
| **listingExchange** | **String** | Main trading venue. | [optional] [default to null] |
| **name** | **String** | Company Name. | [optional] [default to null] |
| **assetClass** | **String** | Group of financial instruments which have similar financial characteristics and behave similar in the marketplace. | [optional] [default to null] |
| **expiry** | **String** | Specific data contract expires. | [optional] [default to null] |
| **lastTradingDay** | **String** | Final day derivative contract can be traded before delivery of the underlying asset or cash settlement. | [optional] [default to null] |
| **group** | **String** | Potential characteristic of each product. | [optional] [default to null] |
| **putOrCall** | **String** | Defines the right to buy or sell of the underlying security. | [optional] [default to null] |
| **sector** | **String** | The category of the economy. | [optional] [default to null] |
| **sectorGroup** | **String** | Stock Group contract belongs too. | [optional] [default to null] |
| **strike** | **BigDecimal** | Set price at which a derivative contract can be bought or sold. | [optional] [default to null] |
| **ticker** | **String** | Contract symbol. | [optional] [default to null] |
| **undConid** | **Integer** | Underlying contract identifier. | [optional] [default to null] |
| **multiplier** | **Integer** | Multiplier for total premium paid or received for derivative contract. | [optional] [default to null] |
| **type** | **String** | Stock type. | [optional] [default to null] |
| **undComp** | **String** | Company name for underlying contract. | [optional] [default to null] |
| **undSym** | **String** | IBKR Symbol for underlying contract. | [optional] [default to null] |
| **hasOptions** | **Boolean** | If contract has an option. | [optional] [default to null] |
| **fullName** | **String** | Formatted company name with underlying symbol, expiration, strike, right. | [optional] [default to null] |
| **isUS** | **Boolean** | If contract is a US contract. Currently supported for stocks, options and warrants. | [optional] [default to null] |
| **incrementRules** | [**secdef_inner_incrementRules**](secdef_inner_incrementRules.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

