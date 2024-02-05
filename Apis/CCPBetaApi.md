# CCPBetaApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**completeCCP**](CCPBetaApi.md#completeCCP) | **POST** /ccp/auth/response | Complete CCP Session |
| [**deleteOrder**](CCPBetaApi.md#deleteOrder) | **DELETE** /ccp/order | Delete Order |
| [**getCCPAccount**](CCPBetaApi.md#getCCPAccount) | **GET** /ccp/account | Brokerage Accounts |
| [**getCCPOrders**](CCPBetaApi.md#getCCPOrders) | **GET** /ccp/orders | Order Status |
| [**getCCPPositions**](CCPBetaApi.md#getCCPPositions) | **GET** /ccp/positions | Positions |
| [**getCCPStatus**](CCPBetaApi.md#getCCPStatus) | **GET** /ccp/status | CCP Status |
| [**getCCPTrades**](CCPBetaApi.md#getCCPTrades) | **GET** /ccp/trades | Trades |
| [**initCCP**](CCPBetaApi.md#initCCP) | **POST** /ccp/auth/init | Start CCP Session |
| [**submitOrder**](CCPBetaApi.md#submitOrder) | **POST** /ccp/order | Submit Order |
| [**updateOrder**](CCPBetaApi.md#updateOrder) | **PUT** /ccp/order | Update Order |


<a name="completeCCP"></a>
# **completeCCP**
> completeCCP_200_response completeCCP(auth)

Complete CCP Session

    Session Token Authentication

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **auth** | [**completeCCP_request**](../Models/completeCCP_request.md)|  | [optional] |

### Return type

[**completeCCP_200_response**](../Models/completeCCP_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteOrder"></a>
# **deleteOrder**
> order-data deleteOrder(acct, id)

Delete Order

    Sends an Order cancellation request. The status of the order can be queried through /ccp/order. Passing arguments as GET is also supported (requires passing action&#x3D;delete) (GET is meant for development only) 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **acct** | **String**| Account Number | [default to null] |
| **id** | **BigDecimal**| Order Identifier of original submit order | [default to null] |

### Return type

[**order-data**](../Models/order-data.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCCPAccount"></a>
# **getCCPAccount**
> getCCPAccount_200_response getCCPAccount()

Brokerage Accounts

    Provides the list of tradeable accounts

### Parameters
This endpoint does not need any parameter.

### Return type

[**getCCPAccount_200_response**](../Models/getCCPAccount_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCCPOrders"></a>
# **getCCPOrders**
> getCCPOrders_200_response getCCPOrders(acct, cancelled)

Order Status

    Get status for all orders

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **acct** | **String**| User Account | [default to null] |
| **cancelled** | **Boolean**| Return only Rejected or Cancelled orders since today midnight | [optional] [default to null] |

### Return type

[**getCCPOrders_200_response**](../Models/getCCPOrders_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCCPPositions"></a>
# **getCCPPositions**
> position-data getCCPPositions()

Positions

    List of positions

### Parameters
This endpoint does not need any parameter.

### Return type

[**position-data**](../Models/position-data.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCCPStatus"></a>
# **getCCPStatus**
> getCCPStatus_200_response getCCPStatus()

CCP Status

    Provide the current CCP session status. When using the Gateway this endpoint will also initiate a brokerage session to CCP by sending /auth/init and response.

### Parameters
This endpoint does not need any parameter.

### Return type

[**getCCPStatus_200_response**](../Models/getCCPStatus_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getCCPTrades"></a>
# **getCCPTrades**
> getCCPOrders_200_response getCCPTrades(from, to)

Trades

    Get a list of Trades, by default, the list is from today midnight to Date.now(). 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **from** | **String**| From Date (YYYYMMDD-HH:mm:ss) or offset (-1,-2,-3..) | [optional] [default to null] |
| **to** | **String**| To Date (YYYYMMDD-HH:mm:ss) or offset (-1,-2,-3..). To value should be bigger than from value.  | [optional] [default to null] |

### Return type

[**getCCPOrders_200_response**](../Models/getCCPOrders_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="initCCP"></a>
# **initCCP**
> initCCP_200_response initCCP(compete, locale, mac, machineId, username)

Start CCP Session

    Initiate a brokerage session to CCP. Only one brokerage session type can run at a time. If an existing brokerage session to iServer is running then call the endpoint /logout first. Note at this time only order management is possible from CCP session, market data and scanner endpoints can&#39;t be used since they are only available from iServer session. Work is in progress to provide new CCP endpoints for market data and scanners.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **compete** | **Boolean**| Allow competing CCP session to run | [optional] [default to null] |
| **locale** | **String**| Concatenate value for language and region, set to \\\&quot;en_US\\\&quot; | [optional] [default to null] |
| **mac** | **String**| Local MAC Address | [optional] [default to null] |
| **machineId** | **String**| Local machine ID | [optional] [default to null] |
| **username** | **String**| Login user, set to dash \\\&quot;-\\\&quot; | [optional] [default to null] |

### Return type

[**initCCP_200_response**](../Models/initCCP_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

<a name="submitOrder"></a>
# **submitOrder**
> order-data submitOrder(acct, conid, ccy, exchange, qty, type, side, price, tif)

Submit Order

    Submits an Order. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **acct** | **String**| User Account | [default to null] |
| **conid** | **BigDecimal**| Contract identifier from IBKR&#39;s database. | [default to null] |
| **ccy** | **String**| Contract Currency | [default to null] [enum: USD, GBP, EUR] |
| **exchange** | **String**| Exchange | [default to null] [enum: NYSE, CBOE, NYMEX] |
| **qty** | **BigDecimal**| Order Quantity | [default to null] |
| **type** | **String**| Order Price; required if order type is limit | [optional] [default to null] [enum: limit, market] |
| **side** | **String**| Side | [optional] [default to null] [enum: sell, buy] |
| **price** | **BigDecimal**| Order Price; required if order type is limit | [optional] [default to null] |
| **tif** | **String**| Time in Force | [optional] [default to null] [enum: IOC, GTC] |

### Return type

[**order-data**](../Models/order-data.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateOrder"></a>
# **updateOrder**
> order-data updateOrder(acct, id)

Update Order

    Updates an Order. Updating an order requires the same arguments as placing an order besides the conid. Note: The status of the order can be queried through GET /ccp/order. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **acct** | **String**| User Account | [default to null] |
| **id** | **BigDecimal**| Order ID to be modified | [default to null] |

### Return type

[**order-data**](../Models/order-data.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

