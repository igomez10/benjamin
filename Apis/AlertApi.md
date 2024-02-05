# AlertApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**iserverAccountAccountIdAlertActivatePost**](AlertApi.md#iserverAccountAccountIdAlertActivatePost) | **POST** /iserver/account/{accountId}/alert/activate | Activate or deactivate an alert |
| [**iserverAccountAccountIdAlertAlertIdDelete**](AlertApi.md#iserverAccountAccountIdAlertAlertIdDelete) | **DELETE** /iserver/account/{accountId}/alert/{alertId} | Delete an alert |
| [**iserverAccountAccountIdAlertPost**](AlertApi.md#iserverAccountAccountIdAlertPost) | **POST** /iserver/account/{accountId}/alert | Create or modify alert |
| [**iserverAccountAccountIdAlertsGet**](AlertApi.md#iserverAccountAccountIdAlertsGet) | **GET** /iserver/account/{accountId}/alerts | Get a list of available alerts |
| [**iserverAccountAlertIdGet**](AlertApi.md#iserverAccountAlertIdGet) | **GET** /iserver/account/alert/{id} | Get details of an alert |
| [**iserverAccountMtaGet**](AlertApi.md#iserverAccountMtaGet) | **GET** /iserver/account/mta | Get MTA alert |


<a name="iserverAccountAccountIdAlertActivatePost"></a>
# **iserverAccountAccountIdAlertActivatePost**
> _iserver_account__accountId__alert_activate_post_200_response iserverAccountAccountIdAlertActivatePost(accountId, body)

Activate or deactivate an alert

    Please note, if alertId is 0, it will activate/deactivate all alerts

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **body** | [**_iserver_account__accountId__alert_activate_post_request**](../Models/_iserver_account__accountId__alert_activate_post_request.md)| order request info | |

### Return type

[**_iserver_account__accountId__alert_activate_post_200_response**](../Models/_iserver_account__accountId__alert_activate_post_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountAccountIdAlertAlertIdDelete"></a>
# **iserverAccountAccountIdAlertAlertIdDelete**
> _iserver_account__accountId__alert_activate_post_200_response iserverAccountAccountIdAlertAlertIdDelete(accountId, alertId)

Delete an alert

    Please be careful, if alertId is 0, it will delete all alerts

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **alertId** | **String**| alert id | [default to null] |

### Return type

[**_iserver_account__accountId__alert_activate_post_200_response**](../Models/_iserver_account__accountId__alert_activate_post_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountAccountIdAlertPost"></a>
# **iserverAccountAccountIdAlertPost**
> _iserver_account__accountId__alert_post_200_response iserverAccountAccountIdAlertPost(accountId, body)

Create or modify alert

    Please note here, DO NOT pass orderId when creating a new alert, toolId is only required for MTA alert 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |
| **body** | [**alert-request**](../Models/alert-request.md)| alert info | |

### Return type

[**_iserver_account__accountId__alert_post_200_response**](../Models/_iserver_account__accountId__alert_post_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountAccountIdAlertsGet"></a>
# **iserverAccountAccountIdAlertsGet**
> List iserverAccountAccountIdAlertsGet(accountId)

Get a list of available alerts

    The response will contain both active and inactive alerts, but it won&#39;t have MTA alert

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **accountId** | **String**| account id | [default to null] |

### Return type

[**List**](../Models/_iserver_account__accountId__alerts_get_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountAlertIdGet"></a>
# **iserverAccountAlertIdGet**
> alert-response iserverAccountAlertIdGet(id)

Get details of an alert

    Use the endpoint /iserver/account/{accountId}/alerts to receive the alert id. The order_id in the response is the alert id. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **String**| alert id | [default to null] |

### Return type

[**alert-response**](../Models/alert-response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverAccountMtaGet"></a>
# **iserverAccountMtaGet**
> alert-response iserverAccountMtaGet()

Get MTA alert

    Each login user only has one mobile trading assistant (MTA) alert with it&#39;s own unique tool id. The tool id cannot be changed. When modified a new order Id is generated. MTA alerts can not be created or deleted. If you call delete /iserver/account/{accountId}/alert/{alertId}, it will reset MTA to default. See [here](https://www.interactivebrokers.com/en/software/mobileiphonemobile/mobileiphone.htm#monitor/trading-assistant.htm) for more information on MTA alerts. 

### Parameters
This endpoint does not need any parameter.

### Return type

[**alert-response**](../Models/alert-response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

