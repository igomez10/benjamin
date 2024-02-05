# TradesApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**iserverAccountTradesGet**](TradesApi.md#iserverAccountTradesGet) | **GET** /iserver/account/trades | List of Trades for the selected account |


<a name="iserverAccountTradesGet"></a>
# **iserverAccountTradesGet**
> List iserverAccountTradesGet()

List of Trades for the selected account

    Returns a list of trades for the currently selected account for current day and six previous days. It is advised to call this endpoint once per session. 

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/trade.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

