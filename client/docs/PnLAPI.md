# \PnLAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IserverAccountPnlPartitionedGet**](PnLAPI.md#IserverAccountPnlPartitionedGet) | **Get** /iserver/account/pnl/partitioned | PnL for the selected account



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
	resp, r, err := apiClient.PnLAPI.IserverAccountPnlPartitionedGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PnLAPI.IserverAccountPnlPartitionedGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountPnlPartitionedGet`: IserverAccountPnlPartitionedGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PnLAPI.IserverAccountPnlPartitionedGet`: %v\n", resp)
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

