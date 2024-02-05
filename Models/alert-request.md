# alert-request
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **orderId** | **Integer** | orderId is required when modifying alert. You can get it from /iserver/account/{accountId}/alerts  | [optional] [default to null] |
| **alertName** | **String** | name of alert | [optional] [default to null] |
| **alertMessage** | **String** | The message you want to receive via email or text message | [optional] [default to null] |
| **alertRepeatable** | **Integer** | whether alert is repeatable or not, so value can only be 0 or 1, this has to be 1 for MTA alert | [optional] [default to null] |
| **email** | **String** | email address to receive alert | [optional] [default to null] |
| **sendMessage** | **Integer** | whether allowing to send email or not, so value can only be 0 or 1,  | [optional] [default to null] |
| **tif** | **String** | time in force, can only be GTC or GTD | [optional] [default to null] |
| **expireTime** | **String** | format, YYYYMMDD-HH:mm:ss, please NOTE this will only work when tif is GTD  | [optional] [default to null] |
| **outsideRth** | **Integer** | value can only be 0 or 1, set to 1 if the alert can be triggered outside regular trading hours.  | [optional] [default to null] |
| **iTWSOrdersOnly** | **Integer** | value can only be 0 or 1, set to 1 to enable the alert only in IBKR mobile  | [optional] [default to null] |
| **showPopup** | **Integer** | value can only be 0 or 1, set to 1 to allow to show alert in pop-ups | [optional] [default to null] |
| **toolId** | **Integer** | for MTA alert only, each user has a unique toolId and it will stay the same, do not send for normal alert  | [optional] [default to null] |
| **playAudio** | **String** | audio message to play when alert is triggered | [optional] [default to null] |
| **conditions** | [**List**](alert_request_conditions_inner.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

