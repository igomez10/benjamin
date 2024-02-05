# \MarketDataAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistory**](MarketDataAPI.md#GetHistory) | **Get** /hmds/history | Market Data History (Beta)
[**GetSnapshot**](MarketDataAPI.md#GetSnapshot) | **Get** /md/snapshot | Market Data Snapshot (Beta)
[**IserverMarketdataConidUnsubscribeGet**](MarketDataAPI.md#IserverMarketdataConidUnsubscribeGet) | **Get** /iserver/marketdata/{conid}/unsubscribe | Market Data Cancel (Single)
[**IserverMarketdataHistoryGet**](MarketDataAPI.md#IserverMarketdataHistoryGet) | **Get** /iserver/marketdata/history | Market Data History
[**IserverMarketdataSnapshotGet**](MarketDataAPI.md#IserverMarketdataSnapshotGet) | **Get** /iserver/marketdata/snapshot | Market Data
[**IserverMarketdataUnsubscribeallGet**](MarketDataAPI.md#IserverMarketdataUnsubscribeallGet) | **Get** /iserver/marketdata/unsubscribeall | Market Data Cancel (All)



## GetHistory

> HistoryResult GetHistory(ctx).Conid(conid).Period(period).Bar(bar).OutsideRth(outsideRth).Execute()

Market Data History (Beta)



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
	period := "period_example" // string | Time period for history request.    * min: Minutes   * h: Hours   * d: Days   * w: Weeks   * m: Months   * y: Years 
	bar := "bar_example" // string | Duration of time for each candlestick bar.   * min: Minutes   * h: Hours   * d: Days   * w: Weeks   * m: Months  (optional)
	outsideRth := true // bool | For contracts that support it, will determine if history data includes outside of regular trading hours. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetHistory(context.Background()).Conid(conid).Period(period).Bar(bar).OutsideRth(outsideRth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHistory`: HistoryResult
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conid** | **int32** | contract id | 
 **period** | **string** | Time period for history request.    * min: Minutes   * h: Hours   * d: Days   * w: Weeks   * m: Months   * y: Years  | 
 **bar** | **string** | Duration of time for each candlestick bar.   * min: Minutes   * h: Hours   * d: Days   * w: Weeks   * m: Months  | 
 **outsideRth** | **bool** | For contracts that support it, will determine if history data includes outside of regular trading hours. | 

### Return type

[**HistoryResult**](HistoryResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> MarketData GetSnapshot(ctx).Conids(conids).Fields(fields).Execute()

Market Data Snapshot (Beta)



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
	conids := "conids_example" // string | List of conids comma separated. Optional exchange and instrument type can be specified.   * conid: IBKR Contract Identifier   * exchange: Exchange or venue   * instrType: Instrument Type supported values: CS (Stocks), OPT (Options), FUT (Futures), FOP (Future Options), WAR (Warrants), BOND (Bonds), FUND (Mutual Funds), CASH (Forex), CFD (Contract for difference), IND (Index) 
	fields := []float32{float32(123)} // []float32 | list of fields separated by comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.GetSnapshot(context.Background()).Conids(conids).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.GetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshot`: MarketData
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.GetSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conids** | **string** | List of conids comma separated. Optional exchange and instrument type can be specified.   * conid: IBKR Contract Identifier   * exchange: Exchange or venue   * instrType: Instrument Type supported values: CS (Stocks), OPT (Options), FUT (Futures), FOP (Future Options), WAR (Warrants), BOND (Bonds), FUND (Mutual Funds), CASH (Forex), CFD (Contract for difference), IND (Index)  | 
 **fields** | **[]float32** | list of fields separated by comma | 

### Return type

[**MarketData**](MarketData.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverMarketdataConidUnsubscribeGet

> IserverMarketdataConidUnsubscribeGet200Response IserverMarketdataConidUnsubscribeGet(ctx, conid).Execute()

Market Data Cancel (Single)



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
	conid := "conid_example" // string | contract id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.IserverMarketdataConidUnsubscribeGet(context.Background(), conid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.IserverMarketdataConidUnsubscribeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverMarketdataConidUnsubscribeGet`: IserverMarketdataConidUnsubscribeGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.IserverMarketdataConidUnsubscribeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conid** | **string** | contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverMarketdataConidUnsubscribeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IserverMarketdataConidUnsubscribeGet200Response**](IserverMarketdataConidUnsubscribeGet200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverMarketdataHistoryGet

> HistoryData IserverMarketdataHistoryGet(ctx).Conid(conid).Period(period).Exchange(exchange).Bar(bar).OutsideRth(outsideRth).Execute()

Market Data History



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
	conid := "conid_example" // string | contract id
	period := "period_example" // string | available time period-- {1-30}min, {1-8}h, {1-1000}d, {1-792}w, {1-182}m, {1-15}y
	exchange := "exchange_example" // string | Exchange of the conid (e.g. ISLAND, NYSE, etc.). Default value is empty which corresponds to primary exchange of the conid. (optional)
	bar := "bar_example" // string | possible value-- 1min, 2min, 3min, 5min, 10min, 15min, 30min, 1h, 2h, 3h, 4h, 8h, 1d, 1w, 1m (optional)
	outsideRth := true // bool | For contracts that support it, will determine if historical data includes outside of regular trading hours. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.IserverMarketdataHistoryGet(context.Background()).Conid(conid).Period(period).Exchange(exchange).Bar(bar).OutsideRth(outsideRth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.IserverMarketdataHistoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverMarketdataHistoryGet`: HistoryData
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.IserverMarketdataHistoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIserverMarketdataHistoryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conid** | **string** | contract id | 
 **period** | **string** | available time period-- {1-30}min, {1-8}h, {1-1000}d, {1-792}w, {1-182}m, {1-15}y | 
 **exchange** | **string** | Exchange of the conid (e.g. ISLAND, NYSE, etc.). Default value is empty which corresponds to primary exchange of the conid. | 
 **bar** | **string** | possible value-- 1min, 2min, 3min, 5min, 10min, 15min, 30min, 1h, 2h, 3h, 4h, 8h, 1d, 1w, 1m | 
 **outsideRth** | **bool** | For contracts that support it, will determine if historical data includes outside of regular trading hours. | 

### Return type

[**HistoryData**](HistoryData.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverMarketdataSnapshotGet

> []IserverMarketdataSnapshotGet200ResponseInner IserverMarketdataSnapshotGet(ctx).Conids(conids).Since(since).Fields(fields).Execute()

Market Data



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
	conids := "conids_example" // string | list of conids separated by comma
	since := int32(56) // int32 | time period since which updates are required. uses epoch time with milliseconds. (optional)
	fields := "fields_example" // string | list of fields separated by comma (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketDataAPI.IserverMarketdataSnapshotGet(context.Background()).Conids(conids).Since(since).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.IserverMarketdataSnapshotGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverMarketdataSnapshotGet`: []IserverMarketdataSnapshotGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.IserverMarketdataSnapshotGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIserverMarketdataSnapshotGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conids** | **string** | list of conids separated by comma | 
 **since** | **int32** | time period since which updates are required. uses epoch time with milliseconds. | 
 **fields** | **string** | list of fields separated by comma | 

### Return type

[**[]IserverMarketdataSnapshotGet200ResponseInner**](IserverMarketdataSnapshotGet200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverMarketdataUnsubscribeallGet

> IserverMarketdataUnsubscribeallGet200Response IserverMarketdataUnsubscribeallGet(ctx).Execute()

Market Data Cancel (All)



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
	resp, r, err := apiClient.MarketDataAPI.IserverMarketdataUnsubscribeallGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketDataAPI.IserverMarketdataUnsubscribeallGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverMarketdataUnsubscribeallGet`: IserverMarketdataUnsubscribeallGet200Response
	fmt.Fprintf(os.Stdout, "Response from `MarketDataAPI.IserverMarketdataUnsubscribeallGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiIserverMarketdataUnsubscribeallGetRequest struct via the builder pattern


### Return type

[**IserverMarketdataUnsubscribeallGet200Response**](IserverMarketdataUnsubscribeallGet200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

