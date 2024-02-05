# \ScannerAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IserverScannerParamsGet**](ScannerAPI.md#IserverScannerParamsGet) | **Get** /iserver/scanner/params | Scanner Parameters
[**IserverScannerRunPost**](ScannerAPI.md#IserverScannerRunPost) | **Post** /iserver/scanner/run | Scanner Run
[**RunScanner**](ScannerAPI.md#RunScanner) | **Post** /hmds/scanner | Run Scanner (Beta)



## IserverScannerParamsGet

> IserverScannerParamsGet200Response IserverScannerParamsGet(ctx).Execute()

Scanner Parameters



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScannerAPI.IserverScannerParamsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScannerAPI.IserverScannerParamsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverScannerParamsGet`: IserverScannerParamsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ScannerAPI.IserverScannerParamsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIserverScannerParamsGetRequest struct via the builder pattern


### Return type

[**IserverScannerParamsGet200Response**](IserverScannerParamsGet200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverScannerRunPost

> []IserverScannerRunPost200ResponseInner IserverScannerRunPost(ctx).Body(body).Execute()

Scanner Run



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
	body := *openapiclient.NewScannerParams() // ScannerParams | scanner-params request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScannerAPI.IserverScannerRunPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScannerAPI.IserverScannerRunPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverScannerRunPost`: []IserverScannerRunPost200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ScannerAPI.IserverScannerRunPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIserverScannerRunPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScannerParams**](ScannerParams.md) | scanner-params request | 

### Return type

[**[]IserverScannerRunPost200ResponseInner**](IserverScannerRunPost200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunScanner

> ScannerResult RunScanner(ctx).Body(body).Execute()

Run Scanner (Beta)



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
	body := *openapiclient.NewRunScannerRequest() // RunScannerRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScannerAPI.RunScanner(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScannerAPI.RunScanner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunScanner`: ScannerResult
	fmt.Fprintf(os.Stdout, "Response from `ScannerAPI.RunScanner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunScannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RunScannerRequest**](RunScannerRequest.md) | request body | 

### Return type

[**ScannerResult**](ScannerResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

