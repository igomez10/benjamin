# PortfolioAnalystApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getPerformance**](PortfolioAnalystApi.md#getPerformance) | **POST** /pa/performance | Account Performance |
| [**paSummaryPost**](PortfolioAnalystApi.md#paSummaryPost) | **POST** /pa/summary | Account Balance&#39;s Summary (Deprecated) |
| [**paTransactionsPost**](PortfolioAnalystApi.md#paTransactionsPost) | **POST** /pa/transactions | Position&#39;s Transaction History |


<a name="getPerformance"></a>
# **getPerformance**
> performance getPerformance(body)

Account Performance

    Returns the performance (MTM) for the given accounts, if more than one account is passed, the result is consolidated.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**getPerformance_request**](../Models/getPerformance_request.md)| an array of account ids | |

### Return type

[**performance**](../Models/performance.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="paSummaryPost"></a>
# **paSummaryPost**
> summary paSummaryPost(body)

Account Balance&#39;s Summary (Deprecated)

    This endpoint is going to be deprecated. Please use /pa/performance instead. Returns a summary of all account balances for the given accounts, if more than one account is passe, the result is consolidated.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**_pa_summary_post_request**](../Models/_pa_summary_post_request.md)| an array of account ids | |

### Return type

[**summary**](../Models/summary.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="paTransactionsPost"></a>
# **paTransactionsPost**
> transactions paTransactionsPost(body)

Position&#39;s Transaction History

    transaction history for a given number of conids and accounts. Types of transactions include dividend payments, buy and sell transactions, transfers. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**_pa_transactions_post_request**](../Models/_pa_transactions_post_request.md)|  | |

### Return type

[**transactions**](../Models/transactions.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

