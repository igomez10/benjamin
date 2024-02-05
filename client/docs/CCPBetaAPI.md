# \CCPBetaAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteCCP**](CCPBetaAPI.md#CompleteCCP) | **Post** /ccp/auth/response | Complete CCP Session
[**DeleteOrder**](CCPBetaAPI.md#DeleteOrder) | **Delete** /ccp/order | Delete Order
[**GetCCPAccount**](CCPBetaAPI.md#GetCCPAccount) | **Get** /ccp/account | Brokerage Accounts
[**GetCCPOrders**](CCPBetaAPI.md#GetCCPOrders) | **Get** /ccp/orders | Order Status
[**GetCCPPositions**](CCPBetaAPI.md#GetCCPPositions) | **Get** /ccp/positions | Positions
[**GetCCPStatus**](CCPBetaAPI.md#GetCCPStatus) | **Get** /ccp/status | CCP Status
[**GetCCPTrades**](CCPBetaAPI.md#GetCCPTrades) | **Get** /ccp/trades | Trades
[**InitCCP**](CCPBetaAPI.md#InitCCP) | **Post** /ccp/auth/init | Start CCP Session
[**SubmitOrder**](CCPBetaAPI.md#SubmitOrder) | **Post** /ccp/order | Submit Order
[**UpdateOrder**](CCPBetaAPI.md#UpdateOrder) | **Put** /ccp/order | Update Order



## CompleteCCP

> CompleteCCP200Response CompleteCCP(ctx).Auth(auth).Execute()

Complete CCP Session



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
	auth := *openapiclient.NewCompleteCCPRequest() // CompleteCCPRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCPBetaAPI.CompleteCCP(context.Background()).Auth(auth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.CompleteCCP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteCCP`: CompleteCCP200Response
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.CompleteCCP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteCCPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **auth** | [**CompleteCCPRequest**](CompleteCCPRequest.md) |  | 

### Return type

[**CompleteCCP200Response**](CompleteCCP200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrder

> OrderData DeleteOrder(ctx).Acct(acct).Id(id).Execute()

Delete Order



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
	acct := "acct_example" // string | Account Number
	id := float32(8.14) // float32 | Order Identifier of original submit order

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCPBetaAPI.DeleteOrder(context.Background()).Acct(acct).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.DeleteOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOrder`: OrderData
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.DeleteOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acct** | **string** | Account Number | 
 **id** | **float32** | Order Identifier of original submit order | 

### Return type

[**OrderData**](OrderData.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCCPAccount

> GetCCPAccount200Response GetCCPAccount(ctx).Execute()

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
	resp, r, err := apiClient.CCPBetaAPI.GetCCPAccount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.GetCCPAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCCPAccount`: GetCCPAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.GetCCPAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCCPAccountRequest struct via the builder pattern


### Return type

[**GetCCPAccount200Response**](GetCCPAccount200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCCPOrders

> GetCCPOrders200Response GetCCPOrders(ctx).Acct(acct).Cancelled(cancelled).Execute()

Order Status



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
	acct := "acct_example" // string | User Account
	cancelled := true // bool | Return only Rejected or Cancelled orders since today midnight (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCPBetaAPI.GetCCPOrders(context.Background()).Acct(acct).Cancelled(cancelled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.GetCCPOrders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCCPOrders`: GetCCPOrders200Response
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.GetCCPOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCCPOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acct** | **string** | User Account | 
 **cancelled** | **bool** | Return only Rejected or Cancelled orders since today midnight | 

### Return type

[**GetCCPOrders200Response**](GetCCPOrders200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCCPPositions

> PositionData GetCCPPositions(ctx).Execute()

Positions



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
	resp, r, err := apiClient.CCPBetaAPI.GetCCPPositions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.GetCCPPositions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCCPPositions`: PositionData
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.GetCCPPositions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCCPPositionsRequest struct via the builder pattern


### Return type

[**PositionData**](PositionData.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCCPStatus

> GetCCPStatus200Response GetCCPStatus(ctx).Execute()

CCP Status



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
	resp, r, err := apiClient.CCPBetaAPI.GetCCPStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.GetCCPStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCCPStatus`: GetCCPStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.GetCCPStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCCPStatusRequest struct via the builder pattern


### Return type

[**GetCCPStatus200Response**](GetCCPStatus200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCCPTrades

> GetCCPOrders200Response GetCCPTrades(ctx).From(from).To(to).Execute()

Trades



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
	from := "from_example" // string | From Date (YYYYMMDD-HH:mm:ss) or offset (-1,-2,-3..) (optional)
	to := "to_example" // string | To Date (YYYYMMDD-HH:mm:ss) or offset (-1,-2,-3..). To value should be bigger than from value.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCPBetaAPI.GetCCPTrades(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.GetCCPTrades``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCCPTrades`: GetCCPOrders200Response
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.GetCCPTrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCCPTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | From Date (YYYYMMDD-HH:mm:ss) or offset (-1,-2,-3..) | 
 **to** | **string** | To Date (YYYYMMDD-HH:mm:ss) or offset (-1,-2,-3..). To value should be bigger than from value.  | 

### Return type

[**GetCCPOrders200Response**](GetCCPOrders200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InitCCP

> InitCCP200Response InitCCP(ctx).Compete(compete).Locale(locale).Mac(mac).MachineId(machineId).Username(username).Execute()

Start CCP Session



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
	compete := true // bool | Allow competing CCP session to run (optional)
	locale := "locale_example" // string | Concatenate value for language and region, set to \\\"en_US\\\" (optional)
	mac := "mac_example" // string | Local MAC Address (optional)
	machineId := "machineId_example" // string | Local machine ID (optional)
	username := "username_example" // string | Login user, set to dash \\\"-\\\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCPBetaAPI.InitCCP(context.Background()).Compete(compete).Locale(locale).Mac(mac).MachineId(machineId).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.InitCCP``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InitCCP`: InitCCP200Response
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.InitCCP`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitCCPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compete** | **bool** | Allow competing CCP session to run | 
 **locale** | **string** | Concatenate value for language and region, set to \\\&quot;en_US\\\&quot; | 
 **mac** | **string** | Local MAC Address | 
 **machineId** | **string** | Local machine ID | 
 **username** | **string** | Login user, set to dash \\\&quot;-\\\&quot; | 

### Return type

[**InitCCP200Response**](InitCCP200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitOrder

> OrderData SubmitOrder(ctx).Acct(acct).Conid(conid).Ccy(ccy).Exchange(exchange).Qty(qty).Type_(type_).Side(side).Price(price).Tif(tif).Execute()

Submit Order



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
	acct := "acct_example" // string | User Account
	conid := float32(8.14) // float32 | Contract identifier from IBKR's database.
	ccy := "ccy_example" // string | Contract Currency
	exchange := "exchange_example" // string | Exchange
	qty := float32(8.14) // float32 | Order Quantity
	type_ := "type__example" // string | Order Price; required if order type is limit (optional)
	side := "side_example" // string | Side (optional)
	price := float32(8.14) // float32 | Order Price; required if order type is limit (optional)
	tif := "tif_example" // string | Time in Force (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCPBetaAPI.SubmitOrder(context.Background()).Acct(acct).Conid(conid).Ccy(ccy).Exchange(exchange).Qty(qty).Type_(type_).Side(side).Price(price).Tif(tif).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.SubmitOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitOrder`: OrderData
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.SubmitOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acct** | **string** | User Account | 
 **conid** | **float32** | Contract identifier from IBKR&#39;s database. | 
 **ccy** | **string** | Contract Currency | 
 **exchange** | **string** | Exchange | 
 **qty** | **float32** | Order Quantity | 
 **type_** | **string** | Order Price; required if order type is limit | 
 **side** | **string** | Side | 
 **price** | **float32** | Order Price; required if order type is limit | 
 **tif** | **string** | Time in Force | 

### Return type

[**OrderData**](OrderData.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrder

> OrderData UpdateOrder(ctx).Acct(acct).Id(id).Execute()

Update Order



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
	acct := "acct_example" // string | User Account
	id := float32(8.14) // float32 | Order ID to be modified

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CCPBetaAPI.UpdateOrder(context.Background()).Acct(acct).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CCPBetaAPI.UpdateOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrder`: OrderData
	fmt.Fprintf(os.Stdout, "Response from `CCPBetaAPI.UpdateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **acct** | **string** | User Account | 
 **id** | **float32** | Order ID to be modified | 

### Return type

[**OrderData**](OrderData.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

