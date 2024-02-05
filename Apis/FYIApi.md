# FYIApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteDevice**](FYIApi.md#deleteDevice) | **DELETE** /fyi/deliveryoptions/{deviceId} | Delete a device |
| [**enableDisableDevice**](FYIApi.md#enableDisableDevice) | **POST** /fyi/deliveryoptions/device | Enable/Disable device option |
| [**enableDisableEmail**](FYIApi.md#enableDisableEmail) | **PUT** /fyi/deliveryoptions/email | Enable/Disable email option |
| [**enableDisableSubscription**](FYIApi.md#enableDisableSubscription) | **POST** /fyi/settings/{typecode} | Enable/Disable certain subscription |
| [**getDeliveryOptions**](FYIApi.md#getDeliveryOptions) | **GET** /fyi/deliveryoptions | Get delivery options |
| [**getDisclaimer**](FYIApi.md#getDisclaimer) | **GET** /fyi/disclaimer/{typecode} | Get disclaimer for a certain kind of fyi |
| [**getMoreNotifications**](FYIApi.md#getMoreNotifications) | **GET** /fyi/notifications/more | Get more notifications based on a certain one |
| [**getNotifications**](FYIApi.md#getNotifications) | **GET** /fyi/notifications | Get a list of notifications |
| [**getSettings**](FYIApi.md#getSettings) | **GET** /fyi/settings | Get a list of subscriptions |
| [**getUnreadNumber**](FYIApi.md#getUnreadNumber) | **GET** /fyi/unreadnumber | Get unread number of fyis. The HTTP method POST is also supported. |
| [**markDisclaimerRead**](FYIApi.md#markDisclaimerRead) | **PUT** /fyi/disclaimer/{typecode} | Mark disclaimer read |
| [**markNotificationRead**](FYIApi.md#markNotificationRead) | **PUT** /fyi/notifications/{notificationId} | Get a list of notifications |


<a name="deleteDevice"></a>
# **deleteDevice**
> Object deleteDevice(deviceId)

Delete a device

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **deviceId** | **String**| device ID | [default to null] |

### Return type

**Object**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="enableDisableDevice"></a>
# **enableDisableDevice**
> markDisclaimerRead_200_response enableDisableDevice(body)

Enable/Disable device option

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**enableDisableDevice_request**](../Models/enableDisableDevice_request.md)| device info | |

### Return type

[**markDisclaimerRead_200_response**](../Models/markDisclaimerRead_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="enableDisableEmail"></a>
# **enableDisableEmail**
> markDisclaimerRead_200_response enableDisableEmail(enabled)

Enable/Disable email option

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **enabled** | **String**| true/false | [default to null] |

### Return type

[**markDisclaimerRead_200_response**](../Models/markDisclaimerRead_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="enableDisableSubscription"></a>
# **enableDisableSubscription**
> Object enableDisableSubscription(typecode, body)

Enable/Disable certain subscription

    Configure which typecode you would like to enable/disable. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **typecode** | **String**| fyi code | [default to null] |
| **body** | [**enableDisableSubscription_request**](../Models/enableDisableSubscription_request.md)|  | |

### Return type

**Object**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getDeliveryOptions"></a>
# **getDeliveryOptions**
> getDeliveryOptions_200_response getDeliveryOptions()

Get delivery options

    options for sending fyis to email and other devices 

### Parameters
This endpoint does not need any parameter.

### Return type

[**getDeliveryOptions_200_response**](../Models/getDeliveryOptions_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getDisclaimer"></a>
# **getDisclaimer**
> getDisclaimer_200_response getDisclaimer(typecode)

Get disclaimer for a certain kind of fyi

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **typecode** | **String**| fyi code, for example --M8, EA | [default to null] |

### Return type

[**getDisclaimer_200_response**](../Models/getDisclaimer_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getMoreNotifications"></a>
# **getMoreNotifications**
> List getMoreNotifications(id)

Get more notifications based on a certain one

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **id** | **String**| id of last notification in the list | [default to null] |

### Return type

[**List**](../Models/notifications_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNotifications"></a>
# **getNotifications**
> List getNotifications(max, exclude, include)

Get a list of notifications

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **max** | **String**| max number of fyis in response | [default to null] |
| **exclude** | **String**| if set, don&#39;t set include | [optional] [default to null] |
| **include** | **String**| if set, don&#39;t set exclude | [optional] [default to null] |

### Return type

[**List**](../Models/notifications_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSettings"></a>
# **getSettings**
> List getSettings()

Get a list of subscriptions

    Return the current choices of subscriptions, we can toggle the option 

### Parameters
This endpoint does not need any parameter.

### Return type

[**List**](../Models/getSettings_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getUnreadNumber"></a>
# **getUnreadNumber**
> getUnreadNumber_200_response getUnreadNumber()

Get unread number of fyis. The HTTP method POST is also supported.

    Returns the total number of unread fyis 

### Parameters
This endpoint does not need any parameter.

### Return type

[**getUnreadNumber_200_response**](../Models/getUnreadNumber_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="markDisclaimerRead"></a>
# **markDisclaimerRead**
> markDisclaimerRead_200_response markDisclaimerRead(typecode)

Mark disclaimer read

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **typecode** | **String**| fyi code, for example --M8, EA | [default to null] |

### Return type

[**markDisclaimerRead_200_response**](../Models/markDisclaimerRead_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="markNotificationRead"></a>
# **markNotificationRead**
> Object markNotificationRead(notificationId)

Get a list of notifications

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **notificationId** | **String**| mark a notification read | [default to null] |

### Return type

**Object**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

