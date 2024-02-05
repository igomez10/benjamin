# \PortfolioAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLedger**](PortfolioAPI.md#GetLedger) | **Get** /portfolio/{accountId}/ledger | Account Ledger
[**PortfolioAccountIdAllocationGet**](PortfolioAPI.md#PortfolioAccountIdAllocationGet) | **Get** /portfolio/{accountId}/allocation | Account Allocation
[**PortfolioAccountIdMetaGet**](PortfolioAPI.md#PortfolioAccountIdMetaGet) | **Get** /portfolio/{accountId}/meta | Account Information
[**PortfolioAccountIdPositionConidGet**](PortfolioAPI.md#PortfolioAccountIdPositionConidGet) | **Get** /portfolio/{accountId}/position/{conid} | Position by Conid
[**PortfolioAccountIdPositionsInvalidatePost**](PortfolioAPI.md#PortfolioAccountIdPositionsInvalidatePost) | **Post** /portfolio/{accountId}/positions/invalidate | Invalidates the backend cache of the Portfolio
[**PortfolioAccountIdPositionsPageIdGet**](PortfolioAPI.md#PortfolioAccountIdPositionsPageIdGet) | **Get** /portfolio/{accountId}/positions/{pageId} | Portfolio Positions
[**PortfolioAccountIdSummaryGet**](PortfolioAPI.md#PortfolioAccountIdSummaryGet) | **Get** /portfolio/{accountId}/summary | Account Summary
[**PortfolioAccountsGet**](PortfolioAPI.md#PortfolioAccountsGet) | **Get** /portfolio/accounts | Portfolio Accounts
[**PortfolioAllocationPost**](PortfolioAPI.md#PortfolioAllocationPost) | **Post** /portfolio/allocation | Account Alloction (All Accounts)
[**PortfolioPositionsConidGet**](PortfolioAPI.md#PortfolioPositionsConidGet) | **Get** /portfolio/positions/{conid} | Positions by Conid
[**PortfolioSubaccounts2Get**](PortfolioAPI.md#PortfolioSubaccounts2Get) | **Get** /portfolio/subaccounts2 | List of Sub-Accounts (Large Accounts)
[**PortfolioSubaccountsGet**](PortfolioAPI.md#PortfolioSubaccountsGet) | **Get** /portfolio/subaccounts | List of Sub-Accounts



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
	resp, r, err := apiClient.PortfolioAPI.GetLedger(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.GetLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLedger`: GetLedger200Response
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.GetLedger`: %v\n", resp)
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


## PortfolioAccountIdAllocationGet

> []AllocationInner PortfolioAccountIdAllocationGet(ctx, accountId).Execute()

Account Allocation



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
	resp, r, err := apiClient.PortfolioAPI.PortfolioAccountIdAllocationGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioAccountIdAllocationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountIdAllocationGet`: []AllocationInner
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioAccountIdAllocationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioAccountIdAllocationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AllocationInner**](AllocationInner.md)

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
	resp, r, err := apiClient.PortfolioAPI.PortfolioAccountIdMetaGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioAccountIdMetaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountIdMetaGet`: []Account
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioAccountIdMetaGet`: %v\n", resp)
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


## PortfolioAccountIdPositionConidGet

> []PositionInner PortfolioAccountIdPositionConidGet(ctx, accountId, conid).Execute()

Position by Conid



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
	conid := int32(56) // int32 | contract id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioAPI.PortfolioAccountIdPositionConidGet(context.Background(), accountId, conid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioAccountIdPositionConidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountIdPositionConidGet`: []PositionInner
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioAccountIdPositionConidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 
**conid** | **int32** | contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioAccountIdPositionConidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PositionInner**](PositionInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioAccountIdPositionsInvalidatePost

> map[string]interface{} PortfolioAccountIdPositionsInvalidatePost(ctx, accountId).Execute()

Invalidates the backend cache of the Portfolio

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
	resp, r, err := apiClient.PortfolioAPI.PortfolioAccountIdPositionsInvalidatePost(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioAccountIdPositionsInvalidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountIdPositionsInvalidatePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioAccountIdPositionsInvalidatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioAccountIdPositionsInvalidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioAccountIdPositionsPageIdGet

> []PositionInner PortfolioAccountIdPositionsPageIdGet(ctx, accountId, pageId).Model(model).Sort(sort).Direction(direction).Period(period).Execute()

Portfolio Positions



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
	pageId := "pageId_example" // string | page id (default to "0")
	model := "model_example" // string | optional (optional)
	sort := "sort_example" // string | declare the table to be sorted by which column (optional)
	direction := "direction_example" // string | in which order, a means ascending - d means descending (optional)
	period := "period_example" // string | period for pnl column, can be 1D, 7D, 1M... (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioAPI.PortfolioAccountIdPositionsPageIdGet(context.Background(), accountId, pageId).Model(model).Sort(sort).Direction(direction).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioAccountIdPositionsPageIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountIdPositionsPageIdGet`: []PositionInner
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioAccountIdPositionsPageIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 
**pageId** | **string** | page id | [default to &quot;0&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioAccountIdPositionsPageIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **model** | **string** | optional | 
 **sort** | **string** | declare the table to be sorted by which column | 
 **direction** | **string** | in which order, a means ascending - d means descending | 
 **period** | **string** | period for pnl column, can be 1D, 7D, 1M... | 

### Return type

[**[]PositionInner**](PositionInner.md)

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
	resp, r, err := apiClient.PortfolioAPI.PortfolioAccountIdSummaryGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioAccountIdSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountIdSummaryGet`: PortfolioAccountIdSummaryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioAccountIdSummaryGet`: %v\n", resp)
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
	resp, r, err := apiClient.PortfolioAPI.PortfolioAccountsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioAccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAccountsGet`: []Account
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioAccountsGet`: %v\n", resp)
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


## PortfolioAllocationPost

> []AllocationInner PortfolioAllocationPost(ctx).Body(body).Execute()

Account Alloction (All Accounts)



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
	body := *openapiclient.NewPaSummaryPostRequest() // PaSummaryPostRequest | accounts info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioAPI.PortfolioAllocationPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioAllocationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioAllocationPost`: []AllocationInner
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioAllocationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioAllocationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaSummaryPostRequest**](PaSummaryPostRequest.md) | accounts info | 

### Return type

[**[]AllocationInner**](AllocationInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PortfolioPositionsConidGet

> PortfolioPositionsConidGet200Response PortfolioPositionsConidGet(ctx, conid).Execute()

Positions by Conid



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
	conid := int32(56) // int32 | contract id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioAPI.PortfolioPositionsConidGet(context.Background(), conid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioPositionsConidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioPositionsConidGet`: PortfolioPositionsConidGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioPositionsConidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conid** | **int32** | contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPortfolioPositionsConidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortfolioPositionsConidGet200Response**](PortfolioPositionsConidGet200Response.md)

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
	resp, r, err := apiClient.PortfolioAPI.PortfolioSubaccounts2Get(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioSubaccounts2Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioSubaccounts2Get`: PortfolioSubaccounts2Get200Response
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioSubaccounts2Get`: %v\n", resp)
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
	resp, r, err := apiClient.PortfolioAPI.PortfolioSubaccountsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAPI.PortfolioSubaccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PortfolioSubaccountsGet`: Account
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAPI.PortfolioSubaccountsGet`: %v\n", resp)
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

