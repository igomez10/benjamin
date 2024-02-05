# \TradesAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IserverAccountTradesGet**](TradesAPI.md#IserverAccountTradesGet) | **Get** /iserver/account/trades | List of Trades for the selected account



## IserverAccountTradesGet

> []Trade IserverAccountTradesGet(ctx).Execute()

List of Trades for the selected account



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
	resp, r, err := apiClient.TradesAPI.IserverAccountTradesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TradesAPI.IserverAccountTradesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountTradesGet`: []Trade
	fmt.Fprintf(os.Stdout, "Response from `TradesAPI.IserverAccountTradesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountTradesGetRequest struct via the builder pattern


### Return type

[**[]Trade**](Trade.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

