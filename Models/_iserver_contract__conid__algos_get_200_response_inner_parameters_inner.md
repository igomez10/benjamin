# _iserver_contract__conid__algos_get_200_response_inner_parameters_inner
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The algo parameter | [default to null] |
| **required** | **Boolean** | If true a value must be entered. | [optional] [default to null] |
| **name** | **String** | Descriptive name of the parameter. | [optional] [default to null] |
| **valueClassName** | **String** | Format of the parameter. | [default to null] |
| **minValue** | **BigDecimal** | Smallest value, only applies to parameters with valueClassName&#x3D;Double. | [optional] [default to null] |
| **maxValue** | **BigDecimal** | Largest value, only applies to parameters with valueClassName&#x3D;Double. | [optional] [default to null] |
| **defaultValue** | **Boolean** | User configured preset for this parameter. | [optional] [default to null] |
| **legalStrings** | **String** | The list of choices | [optional] [default to null] |
| **description** | **String** | Detailed description of the parameter. | [optional] [default to null] |
| **guiRank** | **BigDecimal** | The order in UI, used when building dynamic UI so that more important parameters are presented first. | [optional] [default to null] |
| **priceMarketRule** | **Boolean** | If true, must specify parameter using market rule format. Only applies to parameters with valueClassName&#x3D;Double. | [optional] [default to null] |
| **enabledConditions** | **String** | The rules that UI should apply to algo parameters depending on chosen order type:  * MKT:speedUp:&#x3D;:no - hide SpeedUp param when MKT is chosen for order type.  * LMT:strategyType:&lt;&gt;:empty - strategyType param cannot be empty when LMT is chosen for order type.  * MKT:strategyType:&#x3D;:Marketable - set strategyType param to Marketable and disable (no other choice) when MKT is chosen for order type.  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

