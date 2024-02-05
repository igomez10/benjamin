# alert-response
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **account** | **String** | account id | [optional] [default to null] |
| **order\_id** | **Integer** |  | [optional] [default to null] |
| **alert\_name** | **String** | name of alert | [optional] [default to null] |
| **alert\_message** | **String** | The message you want to receive via email or text message | [optional] [default to null] |
| **alert\_active** | **Integer** | whether alert is active or not, so value can only be 0 or 1 | [optional] [default to null] |
| **alert\_repeatable** | **Integer** | whether alert is repeatable or not, so value can only be 0 or 1 | [optional] [default to null] |
| **alert\_email** | **String** | email address to receive alert | [optional] [default to null] |
| **alert\_send\_message** | **Integer** | whether allowing to send email or not, so value can only be 0 or 1,  | [optional] [default to null] |
| **tif** | **String** | time in force, can only be GTC or GTD | [optional] [default to null] |
| **expire\_time** | **String** | format, YYYYMMDD-HH:mm:ss  | [optional] [default to null] |
| **order\_status** | **String** | status of alert | [optional] [default to null] |
| **outsideRth** | **Integer** | value can only be 0 or 1, set to 1 if the alert can be triggered outside regular trading hours.  | [optional] [default to null] |
| **itws\_orders\_only** | **Integer** | value can only be 0 or 1, set to 1 to enable the alert only in IBKR mobile  | [optional] [default to null] |
| **alert\_show\_popup** | **Integer** | value can only be 0 or 1, set to 1 to allow to show alert in pop-ups | [optional] [default to null] |
| **alert\_triggered** | **Boolean** | whether the alert has been triggered | [optional] [default to null] |
| **order\_not\_editable** | **Boolean** | whether the alert can be edited | [optional] [default to null] |
| **tool\_id** | **Integer** | for MTA alert only, each user has a unique toolId and it will stay the same, do not send for normal alert  | [optional] [default to null] |
| **alert\_play\_audio** | **String** | audio message to play when alert is triggered | [optional] [default to null] |
| **alert\_mta\_currency** | **String** | MTA alert only | [optional] [default to null] |
| **alert\_mta\_defaults** | **String** | MTA alert only | [optional] [default to null] |
| **time\_zone** | **String** | MTA alert only | [optional] [default to null] |
| **alert\_default\_type** | **String** | MTA alert only | [optional] [default to null] |
| **condition\_size** | **Integer** | size of conditions array | [optional] [default to null] |
| **condition\_outside\_rth** | **Integer** | whether allowing the condition can be triggered outside of regular trading hours, 1 means allow | [optional] [default to null] |
| **conditions** | [**List**](alert_response_conditions_inner.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

