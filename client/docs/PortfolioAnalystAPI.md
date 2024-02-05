# \PortfolioAnalystAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPerformance**](PortfolioAnalystAPI.md#GetPerformance) | **Post** /pa/performance | Account Performance
[**PaSummaryPost**](PortfolioAnalystAPI.md#PaSummaryPost) | **Post** /pa/summary | Account Balance&#39;s Summary (Deprecated)
[**PaTransactionsPost**](PortfolioAnalystAPI.md#PaTransactionsPost) | **Post** /pa/transactions | Position&#39;s Transaction History



## GetPerformance

> Performance GetPerformance(ctx).Body(body).Execute()

Account Performance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/igomez10/benjamin/client"
)

func main() {
	body := *openapiclient.NewGetPerformanceRequest() // GetPerformanceRequest | an array of account ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioAnalystAPI.GetPerformance(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAnalystAPI.GetPerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPerformance`: Performance
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAnalystAPI.GetPerformance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetPerformanceRequest**](GetPerformanceRequest.md) | an array of account ids | 

### Return type

[**Performance**](Performance.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaSummaryPost

> Summary PaSummaryPost(ctx).Body(body).Execute()

Account Balance's Summary (Deprecated)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/igomez10/benjamin/client"
)

func main() {
	body := *openapiclient.NewPaSummaryPostRequest() // PaSummaryPostRequest | an array of account ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioAnalystAPI.PaSummaryPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAnalystAPI.PaSummaryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaSummaryPost`: Summary
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAnalystAPI.PaSummaryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaSummaryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaSummaryPostRequest**](PaSummaryPostRequest.md) | an array of account ids | 

### Return type

[**Summary**](Summary.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaTransactionsPost

> Transactions PaTransactionsPost(ctx).Body(body).Execute()

Position's Transaction History



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/igomez10/benjamin/client"
)

func main() {
	body := *openapiclient.NewPaTransactionsPostRequest() // PaTransactionsPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PortfolioAnalystAPI.PaTransactionsPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PortfolioAnalystAPI.PaTransactionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaTransactionsPost`: Transactions
	fmt.Fprintf(os.Stdout, "Response from `PortfolioAnalystAPI.PaTransactionsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaTransactionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PaTransactionsPostRequest**](PaTransactionsPostRequest.md) |  | 

### Return type

[**Transactions**](Transactions.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

