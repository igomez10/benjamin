# ScannerApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**iserverScannerParamsGet**](ScannerApi.md#iserverScannerParamsGet) | **GET** /iserver/scanner/params | Scanner Parameters |
| [**iserverScannerRunPost**](ScannerApi.md#iserverScannerRunPost) | **POST** /iserver/scanner/run | Scanner Run |
| [**runScanner**](ScannerApi.md#runScanner) | **POST** /hmds/scanner | Run Scanner (Beta) |


<a name="iserverScannerParamsGet"></a>
# **iserverScannerParamsGet**
> _iserver_scanner_params_get_200_response iserverScannerParamsGet()

Scanner Parameters

    Returns an object contains four lists contain all parameters for scanners

### Parameters
This endpoint does not need any parameter.

### Return type

[**_iserver_scanner_params_get_200_response**](../Models/_iserver_scanner_params_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverScannerRunPost"></a>
# **iserverScannerRunPost**
> List iserverScannerRunPost(body)

Scanner Run

    Searches for contracts according to the filters specified in scanner/params endpoint

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**scanner-params**](../Models/scanner-params.md)| scanner-params request | |

### Return type

[**List**](../Models/_iserver_scanner_run_post_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="runScanner"></a>
# **runScanner**
> scanner-result runScanner(body)

Run Scanner (Beta)

    Using a direct connection to the market data farm, will provide results to the requested scanner.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**runScanner_request**](../Models/runScanner_request.md)| request body | |

### Return type

[**scanner-result**](../Models/scanner-result.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

