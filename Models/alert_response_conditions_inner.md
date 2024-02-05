# alert_response_conditions_inner
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **condition\_type** | **Integer** | Types: 1-Price, 3-Time, 4-Margin, 5-Trade, 6-Volume, 7: MTA market 8: MTA Position, 9: MTA Acc. Daily PN&amp;  | [optional] [default to null] |
| **conidex** | **String** | conid and exchange. Format supports conid or conid@exchange | [optional] [default to null] |
| **contract\_description\_1** | **String** | Format contract name | [optional] [default to null] |
| **condition\_operator** | **String** | optional, operator for the current condition   * &gt;&#x3D; Greater than or equal to   * &lt;&#x3D; Less than or equal to  | [optional] [default to null] |
| **condition\_trigger\_method** | **String** | optional, only some type of conditions have triggerMethod | [optional] [default to null] |
| **condition\_value** | **String** | can not be empty, can pass default value \&quot;*\&quot; | [optional] [default to null] |
| **condition\_logic\_bind** | **String** | Condition array should end with \&quot;n\&quot;   * a - AND   * o - OR   * n - END  | [optional] [default to null] |
| **condition\_time\_zone** | **String** | only needed for some MTA alert condition | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

