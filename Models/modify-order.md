# modify-order
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **acctId** | **String** |  | [optional] [default to null] |
| **conid** | **Integer** |  | [optional] [default to null] |
| **orderType** | **String** | for example LMT | [optional] [default to null] |
| **outsideRTH** | **Boolean** |  | [optional] [default to null] |
| **price** | **BigDecimal** |  | [optional] [default to null] |
| **auxPrice** | **BigDecimal** |  | [optional] [default to null] |
| **side** | **String** | SELL or BUY | [optional] [default to null] |
| **listingExchange** | **String** | optional, not required | [optional] [default to null] |
| **ticker** | **String** | The ticker symbol of the original place order | [optional] [default to null] |
| **tif** | **String** | Specify a time in force to change how long your order will continue to work in the market | [optional] [default to null] |
| **quantity** | **BigDecimal** | usually integer, for some special cases can be float numbers | [optional] [default to null] |
| **deactivated** | **Boolean** | Set to true if you want to pause a working order. For details refer to the [TWS Users&#39; Guide:](https://guides.interactivebrokers.com/tws/twsguide.html#usersguidebook/getstarted/pause_execution.htm)  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

