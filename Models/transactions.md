# transactions
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | will always be getTransactions | [optional] [default to null] |
| **currency** | **String** | same as request | [optional] [default to null] |
| **includesRealTime** | **Boolean** | Indicates whether current day and realtime data is included in the result | [optional] [default to null] |
| **from** | **BigDecimal** | Period start date. Epoch time, GMT | [optional] [default to null] |
| **to** | **BigDecimal** | Period end date. Epoch time, GMT | [optional] [default to null] |
| **transactions** | [**List**](transactions_transactions_inner.md) | Sorted by date descending | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

