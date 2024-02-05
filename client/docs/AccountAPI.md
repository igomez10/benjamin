# \AccountAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLedger**](AccountAPI.md#GetLedger) | **Get** /portfolio/{accountId}/ledger | Account Ledger
[**IserverAccountPnlPartitionedGet**](AccountAPI.md#IserverAccountPnlPartitionedGet) | **Get** /iserver/account/pnl/partitioned | PnL for the selected account
[**IserverAccountPost**](AccountAPI.md#IserverAccountPost) | **Post** /iserver/account | Switch Account
[**IserverAccountsGet**](AccountAPI.md#IserverAccountsGet) | **Get** /iserver/accounts | Brokerage Accounts
[**PortfolioAccountIdMetaGet**](AccountAPI.md#PortfolioAccountIdMetaGet) | **Get** /portfolio/{accountId}/meta | Account Information
[**PortfolioAccountIdSummaryGet**](AccountAPI.md#PortfolioAccountIdSummaryGet) | **Get** /portfolio/{accountId}/summary | Account Summary
[**PortfolioAccountsGet**](AccountAPI.md#PortfolioAccountsGet) | **Get** /portfolio/accounts | Portfolio Accounts
[**PortfolioSubaccounts2Get**](AccountAPI.md#PortfolioSubaccounts2Get) | **Get** /portfolio/subaccounts2 | List of Sub-Accounts (Large Accounts)
[**PortfolioSubaccountsGet**](AccountAPI.md#PortfolioSubaccountsGet) | **Get** /portfolio/subaccounts | List of Sub-Accounts



## GetLedger

> GetLedger200Response GetLedger(ctx, accountId).Execute()

Account Ledger



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
	accountId := "accountId_example" // string | account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.GetLedger(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.GetLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedger`: GetLedger200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.GetLedger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLedger200Response**](GetLedger200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountPnlPartitionedGet

> IserverAccountPnlPartitionedGet200Response IserverAccountPnlPartitionedGet(ctx).Execute()

PnL for the selected account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.IserverAccountPnlPartitionedGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.IserverAccountPnlPartitionedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountPnlPartitionedGet`: IserverAccountPnlPartitionedGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.IserverAccountPnlPartitionedGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountPnlPartitionedGetRequest struct via the builder pattern


### Return type

[**IserverAccountPnlPartitionedGet200Response**](IserverAccountPnlPartitionedGet200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountPost

> IserverAccountPost200Response IserverAccountPost(ctx).Body(body).Execute()

Switch Account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
	body := *openapiclient.NewSetAccount() // SetAccount | account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.IserverAccountPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.IserverAccountPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountPost`: IserverAccountPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.IserverAccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SetAccount**](SetAccount.md) | account id | 

### Return type

[**IserverAccountPost200Response**](IserverAccountPost200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountsGet

> IserverAccountsGet200Response IserverAccountsGet(ctx).Execute()

Brokerage Accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.IserverAccountsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.IserverAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountsGet`: IserverAccountsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.IserverAccountsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountsGetRequest struct via the builder pattern


### Return type

[**IserverAccountsGet200Response**](IserverAccountsGet200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioAccountIdMetaGet

> []Account PortfolioAccountIdMetaGet(ctx, accountId).Execute()

Account Information



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
	accountId := "accountId_example" // string | account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PortfolioAccountIdMetaGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PortfolioAccountIdMetaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountIdMetaGet`: []Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PortfolioAccountIdMetaGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioAccountIdMetaGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Account**](Account.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioAccountIdSummaryGet

> PortfolioAccountIdSummaryGet200Response PortfolioAccountIdSummaryGet(ctx, accountId).Execute()

Account Summary



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
	accountId := "accountId_example" // string | account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PortfolioAccountIdSummaryGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PortfolioAccountIdSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountIdSummaryGet`: PortfolioAccountIdSummaryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PortfolioAccountIdSummaryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioAccountIdSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortfolioAccountIdSummaryGet200Response**](PortfolioAccountIdSummaryGet200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioAccountsGet

> []Account PortfolioAccountsGet(ctx).Execute()

Portfolio Accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PortfolioAccountsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PortfolioAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountsGet`: []Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PortfolioAccountsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioAccountsGetRequest struct via the builder pattern


### Return type

[**[]Account**](Account.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioSubaccounts2Get

> PortfolioSubaccounts2Get200Response PortfolioSubaccounts2Get(ctx).Page(page).Execute()

List of Sub-Accounts (Large Accounts)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {
	page := "page_example" // string |  (default to "0")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PortfolioSubaccounts2Get(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PortfolioSubaccounts2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioSubaccounts2Get`: PortfolioSubaccounts2Get200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PortfolioSubaccounts2Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioSubaccounts2GetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** |  | [default to &quot;0&quot;]

### Return type

[**PortfolioSubaccounts2Get200Response**](PortfolioSubaccounts2Get200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioSubaccountsGet

> Account PortfolioSubaccountsGet(ctx).Execute()

List of Sub-Accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountAPI.PortfolioSubaccountsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountAPI.PortfolioSubaccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioSubaccountsGet`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountAPI.PortfolioSubaccountsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioSubaccountsGetRequest struct via the builder pattern


### Return type

[**Account**](Account.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

