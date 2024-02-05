# _iserver_account_orders_get_200_response_orders_inner
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **acct** | **String** | Account number | [optional] [default to null] |
| **conidex** | **String** | conid and exchange. Format supports conid or conid@exchange | [optional] [default to null] |
| **conid** | **BigDecimal** | Contract identifier | [optional] [default to null] |
| **orderId** | **String** | Order identifier | [optional] [default to null] |
| **cashCcy** | **String** | Cash currency | [optional] [default to null] |
| **sizeAndFills** | **String** | Quantity outstanding and total quantity concatenated with forward slash separator | [optional] [default to null] |
| **orderDesc** | **String** | Order description | [optional] [default to null] |
| **description1** | **String** | Formatted ticker description | [optional] [default to null] |
| **ticker** | **String** | Underlying symbol | [optional] [default to null] |
| **secType** | **String** | Asset class | [optional] [default to null] |
| **listingExchange** | **String** | Listing Exchange | [optional] [default to null] |
| **remainingQuantity** | **BigDecimal** | Quantity remaining | [optional] [default to null] |
| **filledQuantity** | **BigDecimal** | Quantity filled | [optional] [default to null] |
| **companyName** | **String** | Company Name | [optional] [default to null] |
| **status** | **String** | Status of the order | [optional] [default to null] |
| **origOrderType** | **String** | Original order type | [optional] [default to null] |
| **supportsTaxOpt** | **BigDecimal** | Supports Tax Optimization with 0 for no and 1 for yes | [optional] [default to null] |
| **lastExecutionTime** | **BigDecimal** | Last status update in format YYMMDDhhmms based in GMT | [optional] [default to null] |
| **lastExecutionTime\_r** | **BigDecimal** | Last status update unix time in ms | [optional] [default to null] |
| **orderType** | **String** | Order type | [optional] [default to null] |
| **order\_ref** | **String** | Order reference | [optional] [default to null] |
| **side** | **String** | The side of the market of the order.  * BUY: Buy contract near posted ask price  * SELL: Sell contract near posted bid price  * ASSN: Option Assignment, if BUYSELL&#x3D;BUY and OptionType&#x3D;PUT or BUYSELL&#x3D;SELL and OptionType&#x3D;CALL  * EXER: Option Exercise, if BUYSELL&#x3D;SELL and OptionType&#x3D;PUT or BUYSELL&#x3D;BUY and OptionType&#x3D;CALL  | [optional] [default to null] |
| **timeInForce** | **String** | Time in force | [optional] [default to null] |
| **price** | **BigDecimal** | Price of order | [optional] [default to null] |
| **bgColor** | **String** | Background color in hex format | [optional] [default to null] |
| **fgColor** | **String** | Foreground color in hex format | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

