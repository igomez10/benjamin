# order
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **acct** | **String** | account id | [optional] [default to null] |
| **conid** | **Integer** |  | [optional] [default to null] |
| **orderDesc** | **String** |  | [optional] [default to null] |
| **description1** | **String** |  | [optional] [default to null] |
| **ticker** | **String** | for example FB | [optional] [default to null] |
| **secType** | **String** | for example STK | [optional] [default to null] |
| **listingExchange** | **String** | for example NASDAQ.NMS | [optional] [default to null] |
| **remainingQuantity** | **String** |  | [optional] [default to null] |
| **filledQuantity** | **String** |  | [optional] [default to null] |
| **companyName** | **String** |  | [optional] [default to null] |
| **status** | **String** | * PendingSubmit - Indicates the order was sent, but confirmation has not been received that it has been received by the destination.                   Occurs most commonly if an exchange is closed. * PendingCancel - Indicates that a request has been sent to cancel an order but confirmation has not been received of its cancellation. * PreSubmitted - Indicates that a simulated order type has been accepted by the IBKR system and that this order has yet to be elected.                  The order is held in the IBKR system until the election criteria are met. At that time the order is transmitted to the order destination as specified. * Submitted - Indicates that the order has been accepted at the order destination and is working. * Cancelled - Indicates that the balance of the order has been confirmed cancelled by the IB system.               This could occur unexpectedly when IB or the destination has rejected the order. * Filled - Indicates that the order has been completely filled. * Inactive - Indicates the order is not working, for instance if the order was invalid and triggered an error message,              or if the order was to short a security and shares have not yet been located.  | [optional] [default to null] |
| **origOrderType** | **String** | for example Limit | [optional] [default to null] |
| **side** | **String** | BUY or SELL | [optional] [default to null] |
| **price** | **BigDecimal** |  | [optional] [default to null] |
| **bgColor** | **String** | back-ground color | [optional] [default to null] |
| **fgColor** | **String** |  | [optional] [default to null] |
| **orderId** | **Integer** |  | [optional] [default to null] |
| **parentId** | **Integer** | Only exists in child order of bracket | [optional] [default to null] |
| **order\_ref** | **String** | User defined string used to identify the order. Value is set using \&quot;cOID\&quot; field while placing an order. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

