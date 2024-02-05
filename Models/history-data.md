# history-data
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **symbol** | **String** | Underlying symbol | [optional] [default to null] |
| **text** | **String** | companyName | [optional] [default to null] |
| **priceFactor** | **Integer** | priceFactor is price increment obtained from display rule | [optional] [default to null] |
| **startTime** | **String** | start date time in the format YYYYMMDD-HH:mm:ss | [optional] [default to null] |
| **high** | **String** | High value during this time series with format %h/%v/%t. %h is the high price (scaled by priceFactor), %v is volume (volume factor will always be 100 (reported volume &#x3D; actual volume/100)) and %t is minutes from start time of the chart  | [optional] [default to null] |
| **low** | **String** | Low value during this time series with format %l/%v/%t. %l is the low price (scaled by priceFactor), %v is volume (volume factor will always be 100 (reported volume &#x3D; actual volume/100)) and %t is minutes from start time of the chart  | [optional] [default to null] |
| **timePeriod** | **String** | The duration for the historical data request | [optional] [default to null] |
| **barLength** | **Integer** | The number of seconds in a bar | [optional] [default to null] |
| **mdAvailability** | **String** | Market Data Availability. The field may contain two chars. The first char is the primary code: S &#x3D; Streaming, R &#x3D; Realtime, D &#x3D; Delayed, Z &#x3D; Frozen, Y &#x3D; Frozen Delayed. The second char is the secondary code: P &#x3D; Snapshot Available, p &#x3D; Consolidated.  | [optional] [default to null] |
| **mktDataDelay** | **Integer** | The time it takes, in milliseconds, to process the historical data request | [optional] [default to null] |
| **outsideRth** | **Boolean** | The historical data returned includes outside of regular trading hours  | [optional] [default to null] |
| **tradingDayDuration** | **Integer** | The number of seconds in the trading day | [optional] [default to null] |
| **volumeFactor** | **Integer** |  | [optional] [default to null] |
| **priceDisplayRule** | **Integer** |  | [optional] [default to null] |
| **priceDisplayValue** | **String** |  | [optional] [default to null] |
| **negativeCapable** | **Boolean** |  | [optional] [default to null] |
| **messageVersion** | **Integer** |  | [optional] [default to null] |
| **data** | [**List**](history_data_data_inner.md) |  | [optional] [default to null] |
| **points** | **Integer** | total number of points | [optional] [default to null] |
| **travelTime** | **Integer** |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

