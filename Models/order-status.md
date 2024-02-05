# order-status
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **sub\_type** | **String** | order sub-type | [optional] [default to null] |
| **request\_id** | **String** | order request id | [optional] [default to null] |
| **order\_id** | **Integer** | system generated order id, unique per account | [optional] [default to null] |
| **conidex** | **String** | conid and exchange. Format supports conid or conid@exchange | [optional] [default to null] |
| **symbol** | **String** | Underlying symbol | [optional] [default to null] |
| **side** | **String** | The side of the market of the order.   * B - Buy contract near posted ask price   * S - Sell contract near posted bid price   * X - Option expired  | [optional] [default to null] |
| **contract\_description\_1** | **String** | Format contract name | [optional] [default to null] |
| **listing\_exchange** | **String** | Trading Exchange or Venue | [optional] [default to null] |
| **option\_acct** | **String** |  | [optional] [default to null] |
| **company\_name** | **String** | Contracts company name | [optional] [default to null] |
| **size** | **String** | Quantity updated | [optional] [default to null] |
| **total\_size** | **String** | Total quantity | [optional] [default to null] |
| **currency** | **String** | Contract traded currency | [optional] [default to null] |
| **account** | **String** | account id | [optional] [default to null] |
| **order\_type** | **String** | Types of orders | [optional] [default to null] |
| **limit\_price** | **String** | Limit price | [optional] [default to null] |
| **stop\_price** | **String** | Stop price | [optional] [default to null] |
| **cum\_fill** | **String** | Cumulative fill | [optional] [default to null] |
| **order\_status** | **String** | *  PendingSubmit - Indicates the order was sent, but confirmation has not been received that it has been received by the destination.                    Occurs most commonly if an exchange is closed. *  PendingCancel - Indicates that a request has been sent to cancel an order but confirmation has not been received of its cancellation. *  PreSubmitted - Indicates that a simulated order type has been accepted by the IBKR system and that this order has yet to be elected.                   The order is held in the IBKR system until the election criteria are met. At that time the order is transmitted to the order destination as specified. *  Submitted - Indicates that the order has been accepted at the order destination and is working. *  Cancelled - Indicates that the balance of the order has been confirmed cancelled by the IB system.                This could occur unexpectedly when IB or the destination has rejected the order. *  Filled - Indicates that the order has been completely filled. *  Inactive - Indicates the order is not working, for instance if the order was invalid and triggered an error message,               or if the order was to short a security and shares have not yet been located.  | [optional] [default to null] |
| **order\_status\_description** | **String** | Description of the order status | [optional] [default to null] |
| **tif** | **String** | Time-in-Force - length of time order will continue working before it is canceled. | [optional] [default to null] |
| **fg\_color** | **String** | Foreground color in hex format | [optional] [default to null] |
| **bg\_color** | **String** | Background color in hex format | [optional] [default to null] |
| **order\_not\_editable** | **Boolean** | If true not allowed to modify order | [optional] [default to null] |
| **editable\_fields** | **String** | Fields that can be edited in escaped unicode characters | [optional] [default to null] |
| **cannot\_cancel\_order** | **Boolean** | If true not allowed to cancel order | [optional] [default to null] |
| **outside\_rth** | **Boolean** | If true order trades outside regular trading hours | [optional] [default to null] |
| **deactivate\_order** | **Boolean** | If true order is de-activated | [optional] [default to null] |
| **use\_price\_mgmt\_algo** | **Boolean** | If true price management algo is enabled, refer to https://www.interactivebrokers.com/en/index.php?f&#x3D;43423 | [optional] [default to null] |
| **sec\_type** | **String** | Asset class | [optional] [default to null] |
| **available\_chart\_periods** | **String** | List of available chart periods | [optional] [default to null] |
| **order\_description** | **String** | Format description of order | [optional] [default to null] |
| **order\_description\_with\_contract** | **String** | order_description with the symbol | [optional] [default to null] |
| **alert\_active** | **Integer** |  | [optional] [default to null] |
| **child\_order\_type** | **String** | type of the child order | [optional] [default to null] |
| **size\_and\_fills** | **String** | Format fillQuantity\\totalQuantity | [optional] [default to null] |
| **exit\_strategy\_display\_price** | **String** | Position display price | [optional] [default to null] |
| **exit\_strategy\_chart\_description** | **String** | Position description to display on chart | [optional] [default to null] |
| **exit\_strategy\_tool\_availability** | **String** | * 1: If your account has position or order for contract * 0: If your account has no position or order for contract  | [optional] [default to null] |
| **allowed\_duplicate\_opposite** | **Boolean** | Returns true if contract supports duplicate/opposite side order. | [optional] [default to null] |
| **order\_time** | **String** | Time of status update in unix time | [optional] [default to null] |
| **oca\_group\_id** | **String** | only exists for oca orders, oca orders in same group will have same id | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

