# MarketDataApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getHistory**](MarketDataApi.md#getHistory) | **GET** /hmds/history | Market Data History (Beta) |
| [**getSnapshot**](MarketDataApi.md#getSnapshot) | **GET** /md/snapshot | Market Data Snapshot (Beta) |
| [**iserverMarketdataConidUnsubscribeGet**](MarketDataApi.md#iserverMarketdataConidUnsubscribeGet) | **GET** /iserver/marketdata/{conid}/unsubscribe | Market Data Cancel (Single) |
| [**iserverMarketdataHistoryGet**](MarketDataApi.md#iserverMarketdataHistoryGet) | **GET** /iserver/marketdata/history | Market Data History |
| [**iserverMarketdataSnapshotGet**](MarketDataApi.md#iserverMarketdataSnapshotGet) | **GET** /iserver/marketdata/snapshot | Market Data |
| [**iserverMarketdataUnsubscribeallGet**](MarketDataApi.md#iserverMarketdataUnsubscribeallGet) | **GET** /iserver/marketdata/unsubscribeall | Market Data Cancel (All) |


<a name="getHistory"></a>
# **getHistory**
> history-result getHistory(conid, period, bar, outsideRth)

Market Data History (Beta)

    Using a direct connection to the market data farm, will provide a list of historical market data for given conid.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **Integer**| contract id | [default to null] |
| **period** | **String**| Time period for history request.    * min: Minutes   * h: Hours   * d: Days   * w: Weeks   * m: Months   * y: Years  | [default to null] [enum: min, h, d, w, m, y] |
| **bar** | **String**| Duration of time for each candlestick bar.   * min: Minutes   * h: Hours   * d: Days   * w: Weeks   * m: Months  | [optional] [default to null] [enum: min, h, d, w, m] |
| **outsideRth** | **Boolean**| For contracts that support it, will determine if history data includes outside of regular trading hours. | [optional] [default to null] |

### Return type

[**history-result**](../Models/history-result.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSnapshot"></a>
# **getSnapshot**
> market-data getSnapshot(conids, fields)

Market Data Snapshot (Beta)

    Get a snapshot of Market Data for the given conid(s).See response for a list of available fields that can be requested from the fields argument. Must be connected to a brokerage session before can query snapshot data. First /snapshot endpoint call for given conid(s) will initiate the market data request, make an additional request to receive field values back. To receive all available fields the /snapshot endpoint will need to be called several times. To receive streaming market data the endpoint /ws can be used. Refer to [Streaming WebSocket Data](https://interactivebrokers.github.io/cpwebapi/RealtimeSubscription.html) for details. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conids** | **String**| List of conids comma separated. Optional exchange and instrument type can be specified.   * conid: IBKR Contract Identifier   * exchange: Exchange or venue   * instrType: Instrument Type supported values: CS (Stocks), OPT (Options), FUT (Futures), FOP (Future Options), WAR (Warrants), BOND (Bonds), FUND (Mutual Funds), CASH (Forex), CFD (Contract for difference), IND (Index)  | [default to null] [enum: conid@exchange:instrType] |
| **fields** | [**List**](../Models/BigDecimal.md)| list of fields separated by comma | [optional] [default to null] |

### Return type

[**market-data**](../Models/market-data.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverMarketdataConidUnsubscribeGet"></a>
# **iserverMarketdataConidUnsubscribeGet**
> _iserver_marketdata__conid__unsubscribe_get_200_response iserverMarketdataConidUnsubscribeGet(conid)

Market Data Cancel (Single)

    Cancel market data for given conid. To cancel all market data request(s), see /iserver/marketdata/unsubscribeall. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **String**| contract id | [default to null] |

### Return type

[**_iserver_marketdata__conid__unsubscribe_get_200_response**](../Models/_iserver_marketdata__conid__unsubscribe_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverMarketdataHistoryGet"></a>
# **iserverMarketdataHistoryGet**
> history-data iserverMarketdataHistoryGet(conid, period, exchange, bar, outsideRth)

Market Data History

    Get historical market Data for given conid, length of data is controlled by &#39;period&#39; and &#39;bar&#39;. Formatted as: min&#x3D;minute, h&#x3D;hour, d&#x3D;day, w&#x3D;week, m&#x3D;month, y&#x3D;year e.g. period &#x3D;1y with bar &#x3D;1w returns 52 data points (Max of 1000 data points supported). **Note**: There&#39;s a limit of 5 concurrent requests. Excessive requests will return a &#39;Too many requests&#39; status 429 response. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **String**| contract id | [default to null] |
| **period** | **String**| available time period-- {1-30}min, {1-8}h, {1-1000}d, {1-792}w, {1-182}m, {1-15}y | [default to null] |
| **exchange** | **String**| Exchange of the conid (e.g. ISLAND, NYSE, etc.). Default value is empty which corresponds to primary exchange of the conid. | [optional] [default to null] |
| **bar** | **String**| possible value-- 1min, 2min, 3min, 5min, 10min, 15min, 30min, 1h, 2h, 3h, 4h, 8h, 1d, 1w, 1m | [optional] [default to null] |
| **outsideRth** | **Boolean**| For contracts that support it, will determine if historical data includes outside of regular trading hours. | [optional] [default to null] |

### Return type

[**history-data**](../Models/history-data.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverMarketdataSnapshotGet"></a>
# **iserverMarketdataSnapshotGet**
> List iserverMarketdataSnapshotGet(conids, since, fields)

Market Data

    Get Market Data for the given conid(s). The endpoint will return by default bid, ask, last, change, change pct, close, listing exchange. See response fields for a list of available fields that can be request via fields argument. The endpoint /iserver/accounts must be called prior to /iserver/marketdata/snapshot. For derivative contracts the endpoint /iserver/secdef/search must be called first. First /snapshot endpoint call for given conid will initiate the market data request.  To receive all available fields the /snapshot endpoint will need to be called several times. To receive streaming market data the endpoint /ws can be used. Refer to [Streaming WebSocket Data](https://interactivebrokers.github.io/cpwebapi/RealtimeSubscription.html) for details. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conids** | **String**| list of conids separated by comma | [default to null] |
| **since** | **Integer**| time period since which updates are required. uses epoch time with milliseconds. | [optional] [default to null] |
| **fields** | **String**| list of fields separated by comma | [optional] [default to null] |

### Return type

[**List**](../Models/_iserver_marketdata_snapshot_get_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverMarketdataUnsubscribeallGet"></a>
# **iserverMarketdataUnsubscribeallGet**
> _iserver_marketdata_unsubscribeall_get_200_response iserverMarketdataUnsubscribeallGet()

Market Data Cancel (All)

    Cancel all market data request(s). To cancel market data for given conid, see /iserver/marketdata/{conid}/unsubscribe. 

### Parameters
This endpoint does not need any parameter.

### Return type

[**_iserver_marketdata_unsubscribeall_get_200_response**](../Models/_iserver_marketdata_unsubscribeall_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

