# OrderApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getLiveOrders**](OrderApi.md#getLiveOrders) | **GET** /iserver/account/orders | Live Orders |
| [**iserverAccountAccountIdOrderOrderIdDelete**](OrderApi.md#iserverAccountAccountIdOrderOrderIdDelete) | **DELETE** /iserver/account/{accountId}/order/{orderId} | Cancel Order |
| [**iserverAccountAccountIdOrderOrderIdPost**](OrderApi.md#iserverAccountAccountIdOrderOrderIdPost) | **POST** /iserver/account/{accountId}/order/{orderId} | Modify Order |
| [**iserverAccountAccountIdOrderWhatifPost**](OrderApi.md#iserverAccountAccountIdOrderWhatifPost) | **POST** /iserver/account/{accountId}/order/whatif | Preview Order (Deprecated) |
| [**iserverAccountAccountIdOrdersWhatifPost**](OrderApi.md#iserverAccountAccountIdOrdersWhatifPost) | **POST** /iserver/account/{accountId}/orders/whatif | Preview Orders |
| [**iserverAccountOrderStatusOrderIdGet**](OrderApi.md#iserverAccountOrderStatusOrderIdGet) | **GET** /iserver/account/order/status/{orderId} | Order Status |
| [**iserverAccountOrdersFaGroupPost**](OrderApi.md#iserverAccountOrdersFaGroupPost) | **POST** /iserver/account/orders/{faGroup} | Place Orders for FA |
| [**iserverReplyReplyidPost**](OrderApi.md#iserverReplyReplyidPost) | **POST** /iserver/reply/{replyid} | Place Order Reply |
| [**placeOrder**](OrderApi.md#placeOrder) | **POST** /iserver/account/{accountId}/orders | Place Orders |
| [**placeOrderDeprecated**](OrderApi.md#placeOrderDeprecated) | **POST** /iserver/account/{accountId}/order | Place Order (Deprecated) |


<a name="getLiveOrders"></a>
# **getLiveOrders**
> getLiveOrders_200_response getLiveOrders(Filters)

Live Orders

    The endpoint is meant to be used in polling mode, e.g. requesting every x seconds. The response will contain two objects, one is notification, the other is orders. Orders is the list of live orders (cancelled, filled, submitted). Notifications contains information about execute orders as they happen, see status field. To receive streaming live orders the endpoint /ws can be used. Refer to [Streaming WebSocket Data](https://interactivebrokers.github.io/cpwebapi/RealtimeSubscription.html) for details. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Filters** | **String**| list of filters separated by comma | [optional] [default to null] |

### Return type

[**getLiveOrders_200_response**](../Models/getLiveOrders_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountAccountIdOrderOrderIdDelete"></a>
# **iserverAccountAccountIdOrderOrderIdDelete**
> _iserver_account__accountId__order__orderId__delete_200_response iserverAccountAccountIdOrderOrderIdDelete(accountId, orderId)

Cancel Order

    Cancels an open order. Must call /iserver/accounts endpoint prior to cancelling an order. Use /iservers/account/orders endpoint to review open-order(s) and get latest order status.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id, or fa group if deleting a group order | [default to null] |
| **orderId** | **String**| Customer order id, use /iservers/account/orders endpoint to query orderId. | [default to null] |

### Return type

[**_iserver_account__accountId__order__orderId__delete_200_response**](../Models/_iserver_account__accountId__order__orderId__delete_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountAccountIdOrderOrderIdPost"></a>
# **iserverAccountAccountIdOrderOrderIdPost**
> List iserverAccountAccountIdOrderOrderIdPost(accountId, orderId, body)

Modify Order

    Modifies an open order. Must call /iserver/accounts endpoint prior to modifying an order. Use /iservers/account/orders endpoint to review open-order(s).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id, or fa group if modifying a group order | [default to null] |
| **orderId** | **String**| Customer order id, use /iservers/account/orders endpoint to query orderId. | [default to null] |
| **body** | [**modify-order**](../Models/modify-order.md)| modify-order request | |

### Return type

[**List**](../Models/_iserver_account__accountId__order__orderId__post_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountAccountIdOrderWhatifPost"></a>
# **iserverAccountAccountIdOrderWhatifPost**
> _iserver_account__accountId__order_whatif_post_200_response iserverAccountAccountIdOrderWhatifPost(accountId, body)

Preview Order (Deprecated)

    This end-point is going to be deprecated, you can use /iserver/account/{accountId}/orders/whatif, just pass one order in the array, the order structure will be same. This endpoint allows you to preview order without actually submitting the order and you can get commission information in the response. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **body** | [**order-request**](../Models/order-request.md)| order info | |

### Return type

[**_iserver_account__accountId__order_whatif_post_200_response**](../Models/_iserver_account__accountId__order_whatif_post_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountAccountIdOrdersWhatifPost"></a>
# **iserverAccountAccountIdOrdersWhatifPost**
> _iserver_account__accountId__order_whatif_post_200_response iserverAccountAccountIdOrdersWhatifPost(accountId, body)

Preview Orders

    This endpoint allows you to preview order without actually submitting the order and you can get commission information in the response. Also supports bracket orders. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **body** | [**PlaceOrder_request**](../Models/PlaceOrder_request.md)| order info | |

### Return type

[**_iserver_account__accountId__order_whatif_post_200_response**](../Models/_iserver_account__accountId__order_whatif_post_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountOrderStatusOrderIdGet"></a>
# **iserverAccountOrderStatusOrderIdGet**
> order-status iserverAccountOrderStatusOrderIdGet(orderId)

Order Status

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **orderId** | **String**| Customer order id, use /iservers/account/orders endpoint to query orderId. | [default to null] |

### Return type

[**order-status**](../Models/order-status.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountOrdersFaGroupPost"></a>
# **iserverAccountOrdersFaGroupPost**
> List iserverAccountOrdersFaGroupPost(faGroup, body)

Place Orders for FA

    Financial Advisors can use this endpoint to place an order for a specified group. To place an order for a specified account use the endpoint /iserver/account/{accountId}/order. More information about groups can be found in the [TWS Users&#39; Guide:](https://guides.interactivebrokers.com/twsguide/twsguide.htm#usersguidebook/financialadvisors/create_an_account_group_for_share_allocation.htm). 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **faGroup** | **String**| financial advisor group | [default to null] |
| **body** | [**order-request**](../Models/order-request.md)| order request info | |

### Return type

[**List**](../Models/PlaceOrderDeprecated_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverReplyReplyidPost"></a>
# **iserverReplyReplyidPost**
> List iserverReplyReplyidPost(replyid, body)

Place Order Reply

    Reply to questions when placing orders and submit orders

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **replyid** | **String**| Please use the \&quot;id\&quot; from the response of \&quot;Place Order\&quot; endpoint | [default to null] |
| **body** | [**_iserver_reply__replyid__post_request**](../Models/_iserver_reply__replyid__post_request.md)| Answer to question | |

### Return type

[**List**](../Models/_iserver_reply__replyid__post_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="placeOrder"></a>
# **placeOrder**
> List placeOrder(accountId, body)

Place Orders

    When connected to an IServer Brokerage Session, this endpoint will allow you to submit orders.  CP WEB API supports various advanced orderTypes, for additional details and examples refer to [IBKR Quant Blog](https://www.tradersinsight.news/category/ibkr-quant-news/programming_languages/rest-development/).   * Bracket - Attach additional opposite-side order(s) by using a single **cOID** sent with the parent and set the same value for **parentId** in each child order(s).   * Cash Quantity -  Send orders using monetary value by specifying **cashQty** instead of quantity, e.g. cashQty: 200. The endpoint /iserver/contract/rules returns list of valid orderTypes in cqtTypes.   * Currency Conversion - Convert cash from one currency to another by including **isCcyConv** &#x3D; **true**. To specify the cash quantity use **fxQTY** instead of quantity, e.g. fxQTY: 100.   * Fractional - Contracts that support fractional shares can be traded by specifying **quantity** as a float, e.g. quantity: 0.001. The endpoint /iserver/contract/rules returns a list of valid orderTypes in fraqTypes.   * IB Algos - Attached user-defined settings to your trades by using any of IBKR&#39;s Algo Orders. Use the endpoint /iserver/contract/{conid}/algos to identify the available strategies for a contract.   * One-Cancels-All (OCA) - Group multiple unrelated orders by passing order request info in an array and including **isSingleGroup &#x3D; true** for each order. All orders will be assigned the same oca_group_id. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **body** | [**PlaceOrder_request**](../Models/PlaceOrder_request.md)| order request info | |

### Return type

[**List**](../Models/PlaceOrderDeprecated_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="placeOrderDeprecated"></a>
# **placeOrderDeprecated**
> List placeOrderDeprecated(accountId, body)

Place Order (Deprecated)

    This endpoint is going to be deprecated, you can use /iserver/account/{accountId}/orders, just pass one order in the array, the order structure will be same. Please note here, sometimes this endpoint alone can&#39;t make sure you submit the order successfully, you could receive some questions in the response, you have to to answer them in order to submit the order successfully. You can use \&quot;/iserver/reply/{replyid}\&quot; endpoint to answer questions 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **body** | [**order-request**](../Models/order-request.md)| order request info | |

### Return type

[**List**](../Models/PlaceOrderDeprecated_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

