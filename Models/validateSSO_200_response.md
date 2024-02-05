# validateSSO_200_response
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **LOGIN\_TYPE** | **BigDecimal** | 1 for Live, 2 for Paper | [optional] [default to null] |
| **USER\_NAME** | **String** | Username | [optional] [default to null] |
| **USER\_ID** | **BigDecimal** | User ID | [optional] [default to null] |
| **expire** | **BigDecimal** | Time in milliseconds until session expires. Caller needs to call the again to re-validate session | [optional] [default to null] |
| **RESULT** | **Boolean** | true if session was validated; false if not. | [optional] [default to null] |
| **AUTH\_TIME** | **BigDecimal** | Time of session validation | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

