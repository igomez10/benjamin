# \AlertAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IserverAccountAccountIdAlertActivatePost**](AlertAPI.md#IserverAccountAccountIdAlertActivatePost) | **Post** /iserver/account/{accountId}/alert/activate | Activate or deactivate an alert
[**IserverAccountAccountIdAlertAlertIdDelete**](AlertAPI.md#IserverAccountAccountIdAlertAlertIdDelete) | **Delete** /iserver/account/{accountId}/alert/{alertId} | Delete an alert
[**IserverAccountAccountIdAlertPost**](AlertAPI.md#IserverAccountAccountIdAlertPost) | **Post** /iserver/account/{accountId}/alert | Create or modify alert
[**IserverAccountAccountIdAlertsGet**](AlertAPI.md#IserverAccountAccountIdAlertsGet) | **Get** /iserver/account/{accountId}/alerts | Get a list of available alerts
[**IserverAccountAlertIdGet**](AlertAPI.md#IserverAccountAlertIdGet) | **Get** /iserver/account/alert/{id} | Get details of an alert
[**IserverAccountMtaGet**](AlertAPI.md#IserverAccountMtaGet) | **Get** /iserver/account/mta | Get MTA alert



## IserverAccountAccountIdAlertActivatePost

> IserverAccountAccountIdAlertActivatePost200Response IserverAccountAccountIdAlertActivatePost(ctx, accountId).Body(body).Execute()

Activate or deactivate an alert



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
	accountId := "accountId_example" // string | account id
	body := *openapiclient.NewIserverAccountAccountIdAlertActivatePostRequest() // IserverAccountAccountIdAlertActivatePostRequest | order request info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.IserverAccountAccountIdAlertActivatePost(context.Background(), accountId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.IserverAccountAccountIdAlertActivatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAccountIdAlertActivatePost`: IserverAccountAccountIdAlertActivatePost200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.IserverAccountAccountIdAlertActivatePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAccountIdAlertActivatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IserverAccountAccountIdAlertActivatePostRequest**](IserverAccountAccountIdAlertActivatePostRequest.md) | order request info | 

### Return type

[**IserverAccountAccountIdAlertActivatePost200Response**](IserverAccountAccountIdAlertActivatePost200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountAccountIdAlertAlertIdDelete

> IserverAccountAccountIdAlertActivatePost200Response IserverAccountAccountIdAlertAlertIdDelete(ctx, accountId, alertId).Execute()

Delete an alert



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
	accountId := "accountId_example" // string | account id
	alertId := "alertId_example" // string | alert id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.IserverAccountAccountIdAlertAlertIdDelete(context.Background(), accountId, alertId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.IserverAccountAccountIdAlertAlertIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAccountIdAlertAlertIdDelete`: IserverAccountAccountIdAlertActivatePost200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.IserverAccountAccountIdAlertAlertIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 
**alertId** | **string** | alert id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAccountIdAlertAlertIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IserverAccountAccountIdAlertActivatePost200Response**](IserverAccountAccountIdAlertActivatePost200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountAccountIdAlertPost

> IserverAccountAccountIdAlertPost200Response IserverAccountAccountIdAlertPost(ctx, accountId).Body(body).Execute()

Create or modify alert



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
	accountId := "accountId_example" // string | account id
	body := *openapiclient.NewAlertRequest() // AlertRequest | alert info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.IserverAccountAccountIdAlertPost(context.Background(), accountId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.IserverAccountAccountIdAlertPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAccountIdAlertPost`: IserverAccountAccountIdAlertPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.IserverAccountAccountIdAlertPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAccountIdAlertPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AlertRequest**](AlertRequest.md) | alert info | 

### Return type

[**IserverAccountAccountIdAlertPost200Response**](IserverAccountAccountIdAlertPost200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountAccountIdAlertsGet

> []IserverAccountAccountIdAlertsGet200ResponseInner IserverAccountAccountIdAlertsGet(ctx, accountId).Execute()

Get a list of available alerts



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
	accountId := "accountId_example" // string | account id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.IserverAccountAccountIdAlertsGet(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.IserverAccountAccountIdAlertsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAccountIdAlertsGet`: []IserverAccountAccountIdAlertsGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.IserverAccountAccountIdAlertsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAccountIdAlertsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IserverAccountAccountIdAlertsGet200ResponseInner**](IserverAccountAccountIdAlertsGet200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountAlertIdGet

> AlertResponse IserverAccountAlertIdGet(ctx, id).Execute()

Get details of an alert



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
	id := "id_example" // string | alert id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.IserverAccountAlertIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.IserverAccountAlertIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAlertIdGet`: AlertResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.IserverAccountAlertIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | alert id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAlertIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertResponse**](AlertResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountMtaGet

> AlertResponse IserverAccountMtaGet(ctx).Execute()

Get MTA alert



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
	resp, r, err := apiClient.AlertAPI.IserverAccountMtaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.IserverAccountMtaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountMtaGet`: AlertResponse
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.IserverAccountMtaGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountMtaGetRequest struct via the builder pattern


### Return type

[**AlertResponse**](AlertResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

