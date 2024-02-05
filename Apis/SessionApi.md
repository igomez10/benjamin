# SessionApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**iserverAuthStatusPost**](SessionApi.md#iserverAuthStatusPost) | **POST** /iserver/auth/status | Authentication Status |
| [**iserverReauthenticatePost**](SessionApi.md#iserverReauthenticatePost) | **POST** /iserver/reauthenticate | Tries to re-authenticate to Brokerage |
| [**logout**](SessionApi.md#logout) | **POST** /logout | Ends the current session |
| [**tickle**](SessionApi.md#tickle) | **POST** /tickle | Ping the server to keep the session open |
| [**validateSSO**](SessionApi.md#validateSSO) | **GET** /sso/validate | Validate SSO |


<a name="iserverAuthStatusPost"></a>
# **iserverAuthStatusPost**
> authStatus iserverAuthStatusPost()

Authentication Status

    Current Authentication status to the Brokerage system. Market Data and Trading is not possible if not authenticated, e.g. authenticated shows false

### Parameters
This endpoint does not need any parameter.

### Return type

[**authStatus**](../Models/authStatus.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverReauthenticatePost"></a>
# **iserverReauthenticatePost**
> authStatus iserverReauthenticatePost()

Tries to re-authenticate to Brokerage

    When using the CP Gateway, this endpoint provides a way to reauthenticate to the Brokerage system as long as there is a valid SSO session, see /sso/validate. 

### Parameters
This endpoint does not need any parameter.

### Return type

[**authStatus**](../Models/authStatus.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="logout"></a>
# **logout**
> logout_200_response logout()

Ends the current session

    Logs the user out of the gateway session. Any further activity requires re-authentication.

### Parameters
This endpoint does not need any parameter.

### Return type

[**logout_200_response**](../Models/logout_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="tickle"></a>
# **tickle**
> tickle()

Ping the server to keep the session open

    If the gateway has not received any requests for several minutes an open session will automatically timeout. The tickle endpoint pings the server to prevent the session from ending.

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

<a name="validateSSO"></a>
# **validateSSO**
> validateSSO_200_response validateSSO()

Validate SSO

    Validates the current session for the SSO user

### Parameters
This endpoint does not need any parameter.

### Return type

[**validateSSO_200_response**](../Models/validateSSO_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

