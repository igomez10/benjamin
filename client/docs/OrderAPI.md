# \OrderAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLiveOrders**](OrderAPI.md#GetLiveOrders) | **Get** /iserver/account/orders | Live Orders
[**IserverAccountAccountIdOrderOrderIdDelete**](OrderAPI.md#IserverAccountAccountIdOrderOrderIdDelete) | **Delete** /iserver/account/{accountId}/order/{orderId} | Cancel Order
[**IserverAccountAccountIdOrderOrderIdPost**](OrderAPI.md#IserverAccountAccountIdOrderOrderIdPost) | **Post** /iserver/account/{accountId}/order/{orderId} | Modify Order
[**IserverAccountAccountIdOrderWhatifPost**](OrderAPI.md#IserverAccountAccountIdOrderWhatifPost) | **Post** /iserver/account/{accountId}/order/whatif | Preview Order (Deprecated)
[**IserverAccountAccountIdOrdersWhatifPost**](OrderAPI.md#IserverAccountAccountIdOrdersWhatifPost) | **Post** /iserver/account/{accountId}/orders/whatif | Preview Orders
[**IserverAccountOrderStatusOrderIdGet**](OrderAPI.md#IserverAccountOrderStatusOrderIdGet) | **Get** /iserver/account/order/status/{orderId} | Order Status
[**IserverAccountOrdersFaGroupPost**](OrderAPI.md#IserverAccountOrdersFaGroupPost) | **Post** /iserver/account/orders/{faGroup} | Place Orders for FA
[**IserverReplyReplyidPost**](OrderAPI.md#IserverReplyReplyidPost) | **Post** /iserver/reply/{replyid} | Place Order Reply
[**PlaceOrder**](OrderAPI.md#PlaceOrder) | **Post** /iserver/account/{accountId}/orders | Place Orders
[**PlaceOrderDeprecated**](OrderAPI.md#PlaceOrderDeprecated) | **Post** /iserver/account/{accountId}/order | Place Order (Deprecated)



## GetLiveOrders

> GetLiveOrders200Response GetLiveOrders(ctx).Filters(filters).Execute()

Live Orders



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
	filters := "filters_example" // string | list of filters separated by comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.GetLiveOrders(context.Background()).Filters(filters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.GetLiveOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLiveOrders`: GetLiveOrders200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.GetLiveOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **string** | list of filters separated by comma | 

### Return type

[**GetLiveOrders200Response**](GetLiveOrders200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountAccountIdOrderOrderIdDelete

> IserverAccountAccountIdOrderOrderIdDelete200Response IserverAccountAccountIdOrderOrderIdDelete(ctx, accountId, orderId).Execute()

Cancel Order



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
	accountId := "accountId_example" // string | account id, or fa group if deleting a group order
	orderId := "orderId_example" // string | Customer order id, use /iservers/account/orders endpoint to query orderId.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.IserverAccountAccountIdOrderOrderIdDelete(context.Background(), accountId, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.IserverAccountAccountIdOrderOrderIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAccountIdOrderOrderIdDelete`: IserverAccountAccountIdOrderOrderIdDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.IserverAccountAccountIdOrderOrderIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id, or fa group if deleting a group order | 
**orderId** | **string** | Customer order id, use /iservers/account/orders endpoint to query orderId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAccountIdOrderOrderIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IserverAccountAccountIdOrderOrderIdDelete200Response**](IserverAccountAccountIdOrderOrderIdDelete200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountAccountIdOrderOrderIdPost

> []IserverAccountAccountIdOrderOrderIdPost200ResponseInner IserverAccountAccountIdOrderOrderIdPost(ctx, accountId, orderId).Body(body).Execute()

Modify Order



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
	accountId := "accountId_example" // string | account id, or fa group if modifying a group order
	orderId := "orderId_example" // string | Customer order id, use /iservers/account/orders endpoint to query orderId.
	body := *openapiclient.NewModifyOrder() // ModifyOrder | modify-order request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.IserverAccountAccountIdOrderOrderIdPost(context.Background(), accountId, orderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.IserverAccountAccountIdOrderOrderIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAccountIdOrderOrderIdPost`: []IserverAccountAccountIdOrderOrderIdPost200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.IserverAccountAccountIdOrderOrderIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id, or fa group if modifying a group order | 
**orderId** | **string** | Customer order id, use /iservers/account/orders endpoint to query orderId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAccountIdOrderOrderIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ModifyOrder**](ModifyOrder.md) | modify-order request | 

### Return type

[**[]IserverAccountAccountIdOrderOrderIdPost200ResponseInner**](IserverAccountAccountIdOrderOrderIdPost200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountAccountIdOrderWhatifPost

> IserverAccountAccountIdOrderWhatifPost200Response IserverAccountAccountIdOrderWhatifPost(ctx, accountId).Body(body).Execute()

Preview Order (Deprecated)



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
	body := *openapiclient.NewOrderRequest() // OrderRequest | order info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.IserverAccountAccountIdOrderWhatifPost(context.Background(), accountId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.IserverAccountAccountIdOrderWhatifPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAccountIdOrderWhatifPost`: IserverAccountAccountIdOrderWhatifPost200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.IserverAccountAccountIdOrderWhatifPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAccountIdOrderWhatifPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OrderRequest**](OrderRequest.md) | order info | 

### Return type

[**IserverAccountAccountIdOrderWhatifPost200Response**](IserverAccountAccountIdOrderWhatifPost200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountAccountIdOrdersWhatifPost

> IserverAccountAccountIdOrderWhatifPost200Response IserverAccountAccountIdOrdersWhatifPost(ctx, accountId).Body(body).Execute()

Preview Orders



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
	body := *openapiclient.NewPlaceOrderRequest() // PlaceOrderRequest | order info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.IserverAccountAccountIdOrdersWhatifPost(context.Background(), accountId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.IserverAccountAccountIdOrdersWhatifPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountAccountIdOrdersWhatifPost`: IserverAccountAccountIdOrderWhatifPost200Response
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.IserverAccountAccountIdOrdersWhatifPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountAccountIdOrdersWhatifPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PlaceOrderRequest**](PlaceOrderRequest.md) | order info | 

### Return type

[**IserverAccountAccountIdOrderWhatifPost200Response**](IserverAccountAccountIdOrderWhatifPost200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountOrderStatusOrderIdGet

> OrderStatus IserverAccountOrderStatusOrderIdGet(ctx, orderId).Execute()

Order Status

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
	orderId := "orderId_example" // string | Customer order id, use /iservers/account/orders endpoint to query orderId.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.IserverAccountOrderStatusOrderIdGet(context.Background(), orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.IserverAccountOrderStatusOrderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountOrderStatusOrderIdGet`: OrderStatus
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.IserverAccountOrderStatusOrderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | Customer order id, use /iservers/account/orders endpoint to query orderId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountOrderStatusOrderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrderStatus**](OrderStatus.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverAccountOrdersFaGroupPost

> []PlaceOrderDeprecated200ResponseInner IserverAccountOrdersFaGroupPost(ctx, faGroup).Body(body).Execute()

Place Orders for FA



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
	faGroup := "faGroup_example" // string | financial advisor group
	body := *openapiclient.NewOrderRequest() // OrderRequest | order request info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.IserverAccountOrdersFaGroupPost(context.Background(), faGroup).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.IserverAccountOrdersFaGroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverAccountOrdersFaGroupPost`: []PlaceOrderDeprecated200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.IserverAccountOrdersFaGroupPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**faGroup** | **string** | financial advisor group | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverAccountOrdersFaGroupPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OrderRequest**](OrderRequest.md) | order request info | 

### Return type

[**[]PlaceOrderDeprecated200ResponseInner**](PlaceOrderDeprecated200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverReplyReplyidPost

> []IserverReplyReplyidPost200ResponseInner IserverReplyReplyidPost(ctx, replyid).Body(body).Execute()

Place Order Reply



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
	replyid := "replyid_example" // string | Please use the \"id\" from the response of \"Place Order\" endpoint
	body := *openapiclient.NewIserverReplyReplyidPostRequest() // IserverReplyReplyidPostRequest | Answer to question

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.IserverReplyReplyidPost(context.Background(), replyid).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.IserverReplyReplyidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverReplyReplyidPost`: []IserverReplyReplyidPost200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.IserverReplyReplyidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**replyid** | **string** | Please use the \&quot;id\&quot; from the response of \&quot;Place Order\&quot; endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverReplyReplyidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IserverReplyReplyidPostRequest**](IserverReplyReplyidPostRequest.md) | Answer to question | 

### Return type

[**[]IserverReplyReplyidPost200ResponseInner**](IserverReplyReplyidPost200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOrder

> []PlaceOrderDeprecated200ResponseInner PlaceOrder(ctx, accountId).Body(body).Execute()

Place Orders



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
	body := *openapiclient.NewPlaceOrderRequest() // PlaceOrderRequest | order request info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.PlaceOrder(context.Background(), accountId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.PlaceOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaceOrder`: []PlaceOrderDeprecated200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.PlaceOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PlaceOrderRequest**](PlaceOrderRequest.md) | order request info | 

### Return type

[**[]PlaceOrderDeprecated200ResponseInner**](PlaceOrderDeprecated200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOrderDeprecated

> []PlaceOrderDeprecated200ResponseInner PlaceOrderDeprecated(ctx, accountId).Body(body).Execute()

Place Order (Deprecated)



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
	body := *openapiclient.NewOrderRequest() // OrderRequest | order request info

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderAPI.PlaceOrderDeprecated(context.Background(), accountId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderAPI.PlaceOrderDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaceOrderDeprecated`: []PlaceOrderDeprecated200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `OrderAPI.PlaceOrderDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOrderDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OrderRequest**](OrderRequest.md) | order request info | 

### Return type

[**[]PlaceOrderDeprecated200ResponseInner**](PlaceOrderDeprecated200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

