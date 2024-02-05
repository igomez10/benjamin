# AccountApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getLedger**](AccountApi.md#getLedger) | **GET** /portfolio/{accountId}/ledger | Account Ledger |
| [**iserverAccountPnlPartitionedGet**](AccountApi.md#iserverAccountPnlPartitionedGet) | **GET** /iserver/account/pnl/partitioned | PnL for the selected account |
| [**iserverAccountPost**](AccountApi.md#iserverAccountPost) | **POST** /iserver/account | Switch Account |
| [**iserverAccountsGet**](AccountApi.md#iserverAccountsGet) | **GET** /iserver/accounts | Brokerage Accounts |
| [**portfolioAccountIdMetaGet**](AccountApi.md#portfolioAccountIdMetaGet) | **GET** /portfolio/{accountId}/meta | Account Information |
| [**portfolioAccountIdSummaryGet**](AccountApi.md#portfolioAccountIdSummaryGet) | **GET** /portfolio/{accountId}/summary | Account Summary |
| [**portfolioAccountsGet**](AccountApi.md#portfolioAccountsGet) | **GET** /portfolio/accounts | Portfolio Accounts |
| [**portfolioSubaccounts2Get**](AccountApi.md#portfolioSubaccounts2Get) | **GET** /portfolio/subaccounts2 | List of Sub-Accounts (Large Accounts) |
| [**portfolioSubaccountsGet**](AccountApi.md#portfolioSubaccountsGet) | **GET** /portfolio/subaccounts | List of Sub-Accounts |


<a name="getLedger"></a>
# **getLedger**
> getLedger_200_response getLedger(accountId)

Account Ledger

    Information regarding settled cash, cash balances, etc. in the account&#39;s base currency and any other cash balances hold in other currencies.  /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint. The list of supported currencies is available at https://www.interactivebrokers.com/en/index.php?f&#x3D;3185.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |

### Return type

[**getLedger_200_response**](../Models/getLedger_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

<a name="iserverAccountPost"></a>
# **iserverAccountPost**
> _iserver_account_post_200_response iserverAccountPost(body)

Switch Account

    If an user has multiple accounts, and user wants to get orders, trades, etc. of an account other than currently selected account, then user can update the currently selected account using this API and then can fetch required information for the newly updated account.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**set-account**](../Models/set-account.md)| account id | |

### Return type

[**_iserver_account_post_200_response**](../Models/_iserver_account_post_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountsGet"></a>
# **iserverAccountsGet**
> _iserver_accounts_get_200_response iserverAccountsGet()

Brokerage Accounts

    Returns a list of accounts the user has trading access to, their respective aliases and the currently selected account. Note this endpoint must be called before modifying an order or querying open orders.

### Parameters
This endpoint does not need any parameter.

### Return type

[**_iserver_accounts_get_200_response**](../Models/_iserver_accounts_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="portfolioAccountIdMetaGet"></a>
# **portfolioAccountIdMetaGet**
> List portfolioAccountIdMetaGet(accountId)

Account Information

    Account information related to account Id /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |

### Return type

[**List**](../Models/account.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="portfolioAccountIdSummaryGet"></a>
# **portfolioAccountIdSummaryGet**
> _portfolio__accountId__summary_get_200_response portfolioAccountIdSummaryGet(accountId)

Account Summary

    Returns information about margin, cash balances and other information related to specified account. See also /portfolio/{accountId}/ledger. /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |

### Return type

[**_portfolio__accountId__summary_get_200_response**](../Models/_portfolio__accountId__summary_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="portfolioAccountsGet"></a>
# **portfolioAccountsGet**
> List portfolioAccountsGet()

Portfolio Accounts

    In non-tiered account structures, returns a list of accounts for which the user can view position and account information. This endpoint must be called prior to calling other /portfolio endpoints for those accounts. For querying a list of accounts which the user can trade, see /iserver/accounts. For a list of subaccounts in tiered account structures (e.g. financial advisor or ibroker accounts) see /portfolio/subaccounts.

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/account.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="portfolioSubaccounts2Get"></a>
# **portfolioSubaccounts2Get**
> _portfolio_subaccounts2_get_200_response portfolioSubaccounts2Get(page)

List of Sub-Accounts (Large Accounts)

    Used in tiered account structures (such as Financial Advisor and IBroker Accounts) to return a list of sub-accounts, paginated up to 20 accounts per page, for which the user can view position and account-related information. This endpoint must be called prior to calling other /portfolio endpoints for those sub-accounts. If you have less than 100 sub-accounts use /portfolio/subaccounts. To query a list of accounts the user can trade, see /iserver/accounts.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **page** | **String**|  | [default to 0] |

### Return type

[**_portfolio_subaccounts2_get_200_response**](../Models/_portfolio_subaccounts2_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="portfolioSubaccountsGet"></a>
# **portfolioSubaccountsGet**
> account portfolioSubaccountsGet()

List of Sub-Accounts

    Used in tiered account structures (such as Financial Advisor and IBroker Accounts) to return a list of up to 100 sub-accounts for which the user can view position and account-related information. This endpoint must be called prior to calling other /portfolio endpoints for those sub-accounts. If you have more than 100 sub-accounts use /portfolio/subaccounts2. To query a list of accounts the user can trade, see /iserver/accounts.

### Parameters
This endpoint does not need any parameter.

### Return type

[**account**](../Models/account.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

