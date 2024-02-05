# PortfolioApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getLedger**](PortfolioApi.md#getLedger) | **GET** /portfolio/{accountId}/ledger | Account Ledger |
| [**portfolioAccountIdAllocationGet**](PortfolioApi.md#portfolioAccountIdAllocationGet) | **GET** /portfolio/{accountId}/allocation | Account Allocation |
| [**portfolioAccountIdMetaGet**](PortfolioApi.md#portfolioAccountIdMetaGet) | **GET** /portfolio/{accountId}/meta | Account Information |
| [**portfolioAccountIdPositionConidGet**](PortfolioApi.md#portfolioAccountIdPositionConidGet) | **GET** /portfolio/{accountId}/position/{conid} | Position by Conid |
| [**portfolioAccountIdPositionsInvalidatePost**](PortfolioApi.md#portfolioAccountIdPositionsInvalidatePost) | **POST** /portfolio/{accountId}/positions/invalidate | Invalidates the backend cache of the Portfolio |
| [**portfolioAccountIdPositionsPageIdGet**](PortfolioApi.md#portfolioAccountIdPositionsPageIdGet) | **GET** /portfolio/{accountId}/positions/{pageId} | Portfolio Positions |
| [**portfolioAccountIdSummaryGet**](PortfolioApi.md#portfolioAccountIdSummaryGet) | **GET** /portfolio/{accountId}/summary | Account Summary |
| [**portfolioAccountsGet**](PortfolioApi.md#portfolioAccountsGet) | **GET** /portfolio/accounts | Portfolio Accounts |
| [**portfolioAllocationPost**](PortfolioApi.md#portfolioAllocationPost) | **POST** /portfolio/allocation | Account Alloction (All Accounts) |
| [**portfolioPositionsConidGet**](PortfolioApi.md#portfolioPositionsConidGet) | **GET** /portfolio/positions/{conid} | Positions by Conid |
| [**portfolioSubaccounts2Get**](PortfolioApi.md#portfolioSubaccounts2Get) | **GET** /portfolio/subaccounts2 | List of Sub-Accounts (Large Accounts) |
| [**portfolioSubaccountsGet**](PortfolioApi.md#portfolioSubaccountsGet) | **GET** /portfolio/subaccounts | List of Sub-Accounts |


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

<a name="portfolioAccountIdAllocationGet"></a>
# **portfolioAccountIdAllocationGet**
> List portfolioAccountIdAllocationGet(accountId)

Account Allocation

    Information about the account&#39;s portfolio allocation by Asset Class, Industry and Category.  /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |

### Return type

[**List**](../Models/allocation_inner.md)

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

<a name="portfolioAccountIdPositionConidGet"></a>
# **portfolioAccountIdPositionConidGet**
> List portfolioAccountIdPositionConidGet(accountId, conid)

Position by Conid

    Returns a list of all positions matching the conid. For portfolio models the conid could be in more than one model, returning an array with the name of the model it belongs to.  /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **conid** | **Integer**| contract id | [default to null] |

### Return type

[**List**](../Models/position_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="portfolioAccountIdPositionsInvalidatePost"></a>
# **portfolioAccountIdPositionsInvalidatePost**
> Object portfolioAccountIdPositionsInvalidatePost(accountId)

Invalidates the backend cache of the Portfolio

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |

### Return type

**Object**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="portfolioAccountIdPositionsPageIdGet"></a>
# **portfolioAccountIdPositionsPageIdGet**
> List portfolioAccountIdPositionsPageIdGet(accountId, pageId, model, sort, direction, period)

Portfolio Positions

    Returns a list of positions for the given account. The endpoint supports paging, page&#39;s default size is 30 positions. /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **pageId** | **String**| page id | [default to 0] |
| **model** | **String**| optional | [optional] [default to null] |
| **sort** | **String**| declare the table to be sorted by which column | [optional] [default to null] |
| **direction** | **String**| in which order, a means ascending - d means descending | [optional] [default to null] |
| **period** | **String**| period for pnl column, can be 1D, 7D, 1M... | [optional] [default to null] |

### Return type

[**List**](../Models/position_inner.md)

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

<a name="portfolioAllocationPost"></a>
# **portfolioAllocationPost**
> List portfolioAllocationPost(body)

Account Alloction (All Accounts)

    Similar to /portfolio/{accountId}/allocation but returns a consolidated view of of all the accounts returned by /portfolio/accounts. /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**_pa_summary_post_request**](../Models/_pa_summary_post_request.md)| accounts info | |

### Return type

[**List**](../Models/allocation_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="portfolioPositionsConidGet"></a>
# **portfolioPositionsConidGet**
> _portfolio_positions__conid__get_200_response portfolioPositionsConidGet(conid)

Positions by Conid

    Returns an object of all positions matching the conid for all the selected accounts. For portfolio models the conid could be in more than one model, returning an array with the name of the model it belongs to. /portfolio/accounts or /portfolio/subaccounts must be called prior to this endpoint.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **Integer**| contract id | [default to null] |

### Return type

[**_portfolio_positions__conid__get_200_response**](../Models/_portfolio_positions__conid__get_200_response.md)

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

