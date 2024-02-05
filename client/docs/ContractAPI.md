# \ContractAPI

All URIs are relative to *http://localhost:5000/v1/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFutures**](ContractAPI.md#GetFutures) | **Get** /trsrv/futures | Security Futures by Symbol
[**GetSecdef**](ContractAPI.md#GetSecdef) | **Post** /trsrv/secdef | Secdef by Conid
[**GetSecdefSchedule**](ContractAPI.md#GetSecdefSchedule) | **Get** /trsrv/secdef/schedule | Get trading schedule for symbol
[**GetStocks**](ContractAPI.md#GetStocks) | **Get** /trsrv/stocks | Security Stocks by Symbol
[**IserverContractConidAlgosGet**](ContractAPI.md#IserverContractConidAlgosGet) | **Get** /iserver/contract/{conid}/algos | IB Algo Params
[**IserverContractConidInfoAndRulesGet**](ContractAPI.md#IserverContractConidInfoAndRulesGet) | **Get** /iserver/contract/{conid}/info-and-rules | Info and Rules
[**IserverContractConidInfoGet**](ContractAPI.md#IserverContractConidInfoGet) | **Get** /iserver/contract/{conid}/info | Contract Details
[**IserverContractRulesPost**](ContractAPI.md#IserverContractRulesPost) | **Post** /iserver/contract/rules | Contract Rules
[**IserverSecdefInfoGet**](ContractAPI.md#IserverSecdefInfoGet) | **Get** /iserver/secdef/info | Secdef Info
[**IserverSecdefStrikesGet**](ContractAPI.md#IserverSecdefStrikesGet) | **Get** /iserver/secdef/strikes | Search Strikes
[**SearchBySymbolOrName**](ContractAPI.md#SearchBySymbolOrName) | **Post** /iserver/secdef/search | Search by Symbol or Name



## GetFutures

> GetFutures200Response GetFutures(ctx).Symbols(symbols).Execute()

Security Futures by Symbol



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
	symbols := "symbols_example" // string | list of case-sensitive symbols separated by comma

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetFutures(context.Background()).Symbols(symbols).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetFutures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFutures`: GetFutures200Response
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetFutures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFuturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **string** | list of case-sensitive symbols separated by comma | 

### Return type

[**GetFutures200Response**](GetFutures200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecdef

> []SecdefInner GetSecdef(ctx).Body(body).Execute()

Secdef by Conid



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
	body := *openapiclient.NewGetSecdefRequest() // GetSecdefRequest | request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetSecdef(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetSecdef``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecdef`: []SecdefInner
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetSecdef`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecdefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GetSecdefRequest**](GetSecdefRequest.md) | request body | 

### Return type

[**[]SecdefInner**](SecdefInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecdefSchedule

> GetSecdefSchedule200Response GetSecdefSchedule(ctx).AssetClass(assetClass).Symbol(symbol).Exchange(exchange).ExchangeFilter(exchangeFilter).Execute()

Get trading schedule for symbol



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
	assetClass := "assetClass_example" // string | specify the asset class of the contract. Available values-- Stock: STK, Option: OPT, Future: FUT, Contract For Difference: CFD, Warrant: WAR, Forex: SWP, Mutual Fund: FND, Bond: BND, Inter-Commodity Spreads: ICS 
	symbol := "symbol_example" // string | Underlying Symbol for specified contract, for example 'AAPL' for US Stock - Apple Inc.
	exchange := "exchange_example" // string | Native exchange for contract, for example 'NASDAQ' for US Stock - Apple Inc. (optional)
	exchangeFilter := "exchangeFilter_example" // string | Response only returns trading schedule for specified exchange (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetSecdefSchedule(context.Background()).AssetClass(assetClass).Symbol(symbol).Exchange(exchange).ExchangeFilter(exchangeFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetSecdefSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecdefSchedule`: GetSecdefSchedule200Response
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetSecdefSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecdefScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetClass** | **string** | specify the asset class of the contract. Available values-- Stock: STK, Option: OPT, Future: FUT, Contract For Difference: CFD, Warrant: WAR, Forex: SWP, Mutual Fund: FND, Bond: BND, Inter-Commodity Spreads: ICS  | 
 **symbol** | **string** | Underlying Symbol for specified contract, for example &#39;AAPL&#39; for US Stock - Apple Inc. | 
 **exchange** | **string** | Native exchange for contract, for example &#39;NASDAQ&#39; for US Stock - Apple Inc. | 
 **exchangeFilter** | **string** | Response only returns trading schedule for specified exchange | 

### Return type

[**GetSecdefSchedule200Response**](GetSecdefSchedule200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStocks

> GetStocks200Response GetStocks(ctx).Symbols(symbols).Execute()

Security Stocks by Symbol



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
	symbols := "symbols_example" // string | list of upper-sensitive symbols separated by comma

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.GetStocks(context.Background()).Symbols(symbols).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.GetStocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStocks`: GetStocks200Response
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.GetStocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **string** | list of upper-sensitive symbols separated by comma | 

### Return type

[**GetStocks200Response**](GetStocks200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverContractConidAlgosGet

> []IserverContractConidAlgosGet200ResponseInner IserverContractConidAlgosGet(ctx, conid).Algos(algos).AddDescription(addDescription).AddParams(addParams).Execute()

IB Algo Params



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
	conid := "conid_example" // string | IBKR contract identifier
	algos := "algos_example" // string | List of algo ids delimited by \";\" to filter by. Max of 8 algos ids can be specified. (optional)
	addDescription := "addDescription_example" // string | Whether or not to add algo descriptions to response. Set to 1 for yes, 0 for no. (optional)
	addParams := "addParams_example" // string | Whether or not to show algo parameters.  Set to 1 for yes, 0 for no. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.IserverContractConidAlgosGet(context.Background(), conid).Algos(algos).AddDescription(addDescription).AddParams(addParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.IserverContractConidAlgosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverContractConidAlgosGet`: []IserverContractConidAlgosGet200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.IserverContractConidAlgosGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conid** | **string** | IBKR contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverContractConidAlgosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **algos** | **string** | List of algo ids delimited by \&quot;;\&quot; to filter by. Max of 8 algos ids can be specified. | 
 **addDescription** | **string** | Whether or not to add algo descriptions to response. Set to 1 for yes, 0 for no. | 
 **addParams** | **string** | Whether or not to show algo parameters.  Set to 1 for yes, 0 for no. | 

### Return type

[**[]IserverContractConidAlgosGet200ResponseInner**](IserverContractConidAlgosGet200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverContractConidInfoAndRulesGet

> IserverContractConidInfoAndRulesGet200Response IserverContractConidInfoAndRulesGet(ctx, conid).IsBuy(isBuy).Execute()

Info and Rules



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
	conid := "conid_example" // string | IBKR contract identifier
	isBuy := true // bool | Side of the market rules apply too. Set to **true** for Buy Orders, set to **false** for Sell Orders

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.IserverContractConidInfoAndRulesGet(context.Background(), conid).IsBuy(isBuy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.IserverContractConidInfoAndRulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverContractConidInfoAndRulesGet`: IserverContractConidInfoAndRulesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.IserverContractConidInfoAndRulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conid** | **string** | IBKR contract identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverContractConidInfoAndRulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isBuy** | **bool** | Side of the market rules apply too. Set to **true** for Buy Orders, set to **false** for Sell Orders | 

### Return type

[**IserverContractConidInfoAndRulesGet200Response**](IserverContractConidInfoAndRulesGet200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverContractConidInfoGet

> Contract IserverContractConidInfoGet(ctx, conid).Execute()

Contract Details



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
	resp, r, err := apiClient.ContractAPI.IserverContractConidInfoGet(context.Background(), conid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.IserverContractConidInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverContractConidInfoGet`: Contract
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.IserverContractConidInfoGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conid** | **string** | contract id | 

### Other Parameters

Other parameters are passed through a pointer to a apiIserverContractConidInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contract**](Contract.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverContractRulesPost

> IserverContractRulesPost200Response IserverContractRulesPost(ctx).Conid(conid).Execute()

Contract Rules



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
	conid := *openapiclient.NewIserverContractRulesPostRequest("Conid_example", false) // IserverContractRulesPostRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.IserverContractRulesPost(context.Background()).Conid(conid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.IserverContractRulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverContractRulesPost`: IserverContractRulesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.IserverContractRulesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIserverContractRulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conid** | [**IserverContractRulesPostRequest**](IserverContractRulesPostRequest.md) |  | 

### Return type

[**IserverContractRulesPost200Response**](IserverContractRulesPost200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverSecdefInfoGet

> []SecdefInfo IserverSecdefInfoGet(ctx).Conid(conid).Sectype(sectype).Month(month).Exchange(exchange).Strike(strike).Right(right).Execute()

Secdef Info



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
	conid := "conid_example" // string | underlying contract id
	sectype := "sectype_example" // string | FUT/OPT/WAR/CASH/CFD
	month := "month_example" // string | contract month, only required for FUT/OPT/WAR in the format MMMYY, example JAN00 (optional)
	exchange := "exchange_example" // string | optional, default is SMART (optional)
	strike := float32(8.14) // float32 | optional, only required for OPT/WAR (optional)
	right := "right_example" // string | C for call, P for put (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.IserverSecdefInfoGet(context.Background()).Conid(conid).Sectype(sectype).Month(month).Exchange(exchange).Strike(strike).Right(right).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.IserverSecdefInfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverSecdefInfoGet`: []SecdefInfo
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.IserverSecdefInfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIserverSecdefInfoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conid** | **string** | underlying contract id | 
 **sectype** | **string** | FUT/OPT/WAR/CASH/CFD | 
 **month** | **string** | contract month, only required for FUT/OPT/WAR in the format MMMYY, example JAN00 | 
 **exchange** | **string** | optional, default is SMART | 
 **strike** | **float32** | optional, only required for OPT/WAR | 
 **right** | **string** | C for call, P for put | 

### Return type

[**[]SecdefInfo**](SecdefInfo.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IserverSecdefStrikesGet

> IserverSecdefStrikesGet200Response IserverSecdefStrikesGet(ctx).Conid(conid).Sectype(sectype).Month(month).Exchange(exchange).Execute()

Search Strikes



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
	conid := "conid_example" // string | contract id of the underlying contract
	sectype := "sectype_example" // string | OPT/WAR
	month := "month_example" // string | contract month
	exchange := "exchange_example" // string | optional, default is SMART (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.IserverSecdefStrikesGet(context.Background()).Conid(conid).Sectype(sectype).Month(month).Exchange(exchange).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.IserverSecdefStrikesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IserverSecdefStrikesGet`: IserverSecdefStrikesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.IserverSecdefStrikesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIserverSecdefStrikesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conid** | **string** | contract id of the underlying contract | 
 **sectype** | **string** | OPT/WAR | 
 **month** | **string** | contract month | 
 **exchange** | **string** | optional, default is SMART | 

### Return type

[**IserverSecdefStrikesGet200Response**](IserverSecdefStrikesGet200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBySymbolOrName

> []SearchBySymbolOrName200ResponseInner SearchBySymbolOrName(ctx).Symbol(symbol).Execute()

Search by Symbol or Name



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
	symbol := *openapiclient.NewSearchBySymbolOrNameRequest("Symbol_example") // SearchBySymbolOrNameRequest | Symbol or Company Name to be searched

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContractAPI.SearchBySymbolOrName(context.Background()).Symbol(symbol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContractAPI.SearchBySymbolOrName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchBySymbolOrName`: []SearchBySymbolOrName200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ContractAPI.SearchBySymbolOrName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBySymbolOrNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | [**SearchBySymbolOrNameRequest**](SearchBySymbolOrNameRequest.md) | Symbol or Company Name to be searched | 

### Return type

[**[]SearchBySymbolOrName200ResponseInner**](SearchBySymbolOrName200ResponseInner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

