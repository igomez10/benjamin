# _iserver_contract__conid__info_and_rules_get_200_response
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **cfi\_code** | **String** | Classification of Financial Instrument codes | [optional] [default to null] |
| **symbol** | **String** | Underlying symbol | [optional] [default to null] |
| **cusip** | **String** |  | [optional] [default to null] |
| **expiry\_full** | **BigDecimal** | Expiration Date in the format YYYYMMDD | [optional] [default to null] |
| **con\_id** | **BigDecimal** | IBKRs contract identifier | [optional] [default to null] |
| **maturity\_date** | **BigDecimal** | Date on which the underlying transaction settles if the option is exercised | [optional] [default to null] |
| **industry** | **String** | Specific group of companies or businesses. | [optional] [default to null] |
| **instrument\_type** | **String** | Asset Class of the contract | [optional] [default to null] |
| **trading\_class** | **String** | Designation of the contract | [optional] [default to null] |
| **valid\_exchanges** | **String** | Comma separated list of exchanges or trading venues | [optional] [default to null] |
| **allow\_sell\_long** | **Boolean** | Allowed to sell shares that you own | [optional] [default to null] |
| **is\_zero\_commission\_security** | **Boolean** | Supports zero commission trades | [optional] [default to null] |
| **local\_symbol** | **String** | Contracts symbol from primary exchange. For options it is the OCC symbol. | [optional] [default to null] |
| **classifier** | **String** |  | [optional] [default to null] |
| **currency** | **String** | Currency contract trades in | [optional] [default to null] |
| **text** | **String** | Formatted contract parameters | [optional] [default to null] |
| **underlying\_con\_id** | **BigDecimal** | IBKRs contract identifier for the underlying instrument | [optional] [default to null] |
| **r\_t\_h** | **Boolean** | Provides trading outside of Regular Trading Hours | [optional] [default to null] |
| **multiplier** | **String** | numerical value of each point of price movement | [optional] [default to null] |
| **strike** | **String** | fixed price at which the owner of the option buys or sells the underlying | [optional] [default to null] |
| **right** | **String** | Put or Call of the option | [optional] [default to null] |
| **underlying\_issuer** | **String** | Legal entity for underlying contract | [optional] [default to null] |
| **contract\_month** | **String** | Month the contract must be satisfied by making or accepting delivery | [optional] [default to null] |
| **company\_name** | **String** | Contracts company name | [optional] [default to null] |
| **smart\_available** | **Boolean** | Support IBKRs SMART routing | [optional] [default to null] |
| **exchange** | **String** | Primary Exchange, Routing or Trading Venue | [optional] [default to null] |
| **rules** | [**List**](_iserver_contract__conid__info_and_rules_get_200_response_rules_inner.md) |  | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

