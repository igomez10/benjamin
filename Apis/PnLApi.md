# PnLApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**iserverAccountPnlPartitionedGet**](PnLApi.md#iserverAccountPnlPartitionedGet) | **GET** /iserver/account/pnl/partitioned | PnL for the selected account |


<a name="iserverAccountPnlPartitionedGet"></a>
# **iserverAccountPnlPartitionedGet**
> _iserver_account_pnl_partitioned_get_200_response iserverAccountPnlPartitionedGet()

PnL for the selected account

    Returns an object containing PnL for the selected account and its models (if any). To receive streaming PnL the endpoint /ws can be used. Refer to [Streaming WebSocket Data](https://interactivebrokers.github.io/cpwebapi/RealtimeSubscription.html) for details. 

### Parameters
This endpoint does not need any parameter.

### Return type

[**_iserver_account_pnl_partitioned_get_200_response**](../Models/_iserver_account_pnl_partitioned_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

