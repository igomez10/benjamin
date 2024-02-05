# \FYIAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevice**](FYIAPI.md#DeleteDevice) | **Delete** /fyi/deliveryoptions/{deviceId} | Delete a device
[**EnableDisableDevice**](FYIAPI.md#EnableDisableDevice) | **Post** /fyi/deliveryoptions/device | Enable/Disable device option
[**EnableDisableEmail**](FYIAPI.md#EnableDisableEmail) | **Put** /fyi/deliveryoptions/email | Enable/Disable email option
[**EnableDisableSubscription**](FYIAPI.md#EnableDisableSubscription) | **Post** /fyi/settings/{typecode} | Enable/Disable certain subscription
[**GetDeliveryOptions**](FYIAPI.md#GetDeliveryOptions) | **Get** /fyi/deliveryoptions | Get delivery options
[**GetDisclaimer**](FYIAPI.md#GetDisclaimer) | **Get** /fyi/disclaimer/{typecode} | Get disclaimer for a certain kind of fyi
[**GetMoreNotifications**](FYIAPI.md#GetMoreNotifications) | **Get** /fyi/notifications/more | Get more notifications based on a certain one
[**GetNotifications**](FYIAPI.md#GetNotifications) | **Get** /fyi/notifications | Get a list of notifications
[**GetSettings**](FYIAPI.md#GetSettings) | **Get** /fyi/settings | Get a list of subscriptions
[**GetUnreadNumber**](FYIAPI.md#GetUnreadNumber) | **Get** /fyi/unreadnumber | Get unread number of fyis. The HTTP method POST is also supported.
[**MarkDisclaimerRead**](FYIAPI.md#MarkDisclaimerRead) | **Put** /fyi/disclaimer/{typecode} | Mark disclaimer read
[**MarkNotificationRead**](FYIAPI.md#MarkNotificationRead) | **Put** /fyi/notifications/{notificationId} | Get a list of notifications



## DeleteDevice

> map[string]interface{} DeleteDevice(ctx, deviceId).Execute()

Delete a device

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
	deviceId := "deviceId_example" // string | device ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.DeleteDevice(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.DeleteDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.DeleteDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## EnableDisableDevice

> MarkDisclaimerRead200Response EnableDisableDevice(ctx).Body(body).Execute()

Enable/Disable device option

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
	body := *openapiclient.NewEnableDisableDeviceRequest() // EnableDisableDeviceRequest | device info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.EnableDisableDevice(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.EnableDisableDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableDisableDevice`: MarkDisclaimerRead200Response
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.EnableDisableDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableDisableDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EnableDisableDeviceRequest**](EnableDisableDeviceRequest.md) | device info | 

### Return type

[**MarkDisclaimerRead200Response**](MarkDisclaimerRead200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableDisableEmail

> MarkDisclaimerRead200Response EnableDisableEmail(ctx).Enabled(enabled).Execute()

Enable/Disable email option

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
	enabled := "enabled_example" // string | true/false

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.EnableDisableEmail(context.Background()).Enabled(enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.EnableDisableEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableDisableEmail`: MarkDisclaimerRead200Response
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.EnableDisableEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableDisableEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enabled** | **string** | true/false | 

### Return type

[**MarkDisclaimerRead200Response**](MarkDisclaimerRead200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableDisableSubscription

> map[string]interface{} EnableDisableSubscription(ctx, typecode).Body(body).Execute()

Enable/Disable certain subscription



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
	typecode := "typecode_example" // string | fyi code
	body := *openapiclient.NewEnableDisableSubscriptionRequest() // EnableDisableSubscriptionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.EnableDisableSubscription(context.Background(), typecode).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.EnableDisableSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnableDisableSubscription`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.EnableDisableSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typecode** | **string** | fyi code | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableDisableSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EnableDisableSubscriptionRequest**](EnableDisableSubscriptionRequest.md) |  | 

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


## GetDeliveryOptions

> GetDeliveryOptions200Response GetDeliveryOptions(ctx).Execute()

Get delivery options



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
	resp, r, err := apiClient.FYIAPI.GetDeliveryOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.GetDeliveryOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeliveryOptions`: GetDeliveryOptions200Response
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.GetDeliveryOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeliveryOptionsRequest struct via the builder pattern


### Return type

[**GetDeliveryOptions200Response**](GetDeliveryOptions200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDisclaimer

> GetDisclaimer200Response GetDisclaimer(ctx, typecode).Execute()

Get disclaimer for a certain kind of fyi

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
	typecode := "typecode_example" // string | fyi code, for example --M8, EA

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.GetDisclaimer(context.Background(), typecode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.GetDisclaimer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDisclaimer`: GetDisclaimer200Response
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.GetDisclaimer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typecode** | **string** | fyi code, for example --M8, EA | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDisclaimerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDisclaimer200Response**](GetDisclaimer200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMoreNotifications

> []NotificationsInner GetMoreNotifications(ctx).Id(id).Execute()

Get more notifications based on a certain one

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
	id := "id_example" // string | id of last notification in the list

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.GetMoreNotifications(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.GetMoreNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMoreNotifications`: []NotificationsInner
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.GetMoreNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMoreNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | id of last notification in the list | 

### Return type

[**[]NotificationsInner**](NotificationsInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotifications

> []NotificationsInner GetNotifications(ctx).Max(max).Exclude(exclude).Include(include).Execute()

Get a list of notifications

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
	max := "max_example" // string | max number of fyis in response
	exclude := "exclude_example" // string | if set, don't set include (optional)
	include := "include_example" // string | if set, don't set exclude (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.GetNotifications(context.Background()).Max(max).Exclude(exclude).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.GetNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotifications`: []NotificationsInner
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.GetNotifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **max** | **string** | max number of fyis in response | 
 **exclude** | **string** | if set, don&#39;t set include | 
 **include** | **string** | if set, don&#39;t set exclude | 

### Return type

[**[]NotificationsInner**](NotificationsInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettings

> []GetSettings200ResponseInner GetSettings(ctx).Execute()

Get a list of subscriptions



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
	resp, r, err := apiClient.FYIAPI.GetSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.GetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings`: []GetSettings200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.GetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


### Return type

[**[]GetSettings200ResponseInner**](GetSettings200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnreadNumber

> GetUnreadNumber200Response GetUnreadNumber(ctx).Execute()

Get unread number of fyis. The HTTP method POST is also supported.



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
	resp, r, err := apiClient.FYIAPI.GetUnreadNumber(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.GetUnreadNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnreadNumber`: GetUnreadNumber200Response
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.GetUnreadNumber`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnreadNumberRequest struct via the builder pattern


### Return type

[**GetUnreadNumber200Response**](GetUnreadNumber200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkDisclaimerRead

> MarkDisclaimerRead200Response MarkDisclaimerRead(ctx, typecode).Execute()

Mark disclaimer read

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
	typecode := "typecode_example" // string | fyi code, for example --M8, EA

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.MarkDisclaimerRead(context.Background(), typecode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.MarkDisclaimerRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkDisclaimerRead`: MarkDisclaimerRead200Response
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.MarkDisclaimerRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**typecode** | **string** | fyi code, for example --M8, EA | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkDisclaimerReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MarkDisclaimerRead200Response**](MarkDisclaimerRead200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkNotificationRead

> map[string]interface{} MarkNotificationRead(ctx, notificationId).Execute()

Get a list of notifications

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
	notificationId := "notificationId_example" // string | mark a notification read

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FYIAPI.MarkNotificationRead(context.Background(), notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FYIAPI.MarkNotificationRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkNotificationRead`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FYIAPI.MarkNotificationRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationId** | **string** | mark a notification read | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkNotificationReadRequest struct via the builder pattern


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

