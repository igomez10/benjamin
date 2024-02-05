# history_result_bars
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **open** | **BigDecimal** | First price returned for bar value. | [optional] [default to null] |
| **startTime** | **String** | Start Time in the format YYYYMMDD. | [optional] [default to null] |
| **startTimeVal** | **Integer** | Start Time Value - Formatted in unix time in ms. | [optional] [default to null] |
| **endTime** | **String** | End Time in the format YYYYMMDD. | [optional] [default to null] |
| **endTimeVal** | **Integer** | End Time Value - Formatted in unix time in ms. | [optional] [default to null] |
| **points** | **Integer** | total number of data points. | [optional] [default to null] |
| **data** | [**List**](history_result_bars_data_inner.md) |  | [optional] [default to null] |
| **mktDataDelay** | **Integer** | If 0 then data is returned in real time. Otherwise will return the number of seconds history data is delayed. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

