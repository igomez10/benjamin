# ContractApi

All URIs are relative to *http://localhost:5000/v1/api*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getFutures**](ContractApi.md#getFutures) | **GET** /trsrv/futures | Security Futures by Symbol |
| [**getSecdef**](ContractApi.md#getSecdef) | **POST** /trsrv/secdef | Secdef by Conid |
| [**getSecdefSchedule**](ContractApi.md#getSecdefSchedule) | **GET** /trsrv/secdef/schedule | Get trading schedule for symbol |
| [**getStocks**](ContractApi.md#getStocks) | **GET** /trsrv/stocks | Security Stocks by Symbol |
| [**iserverContractConidAlgosGet**](ContractApi.md#iserverContractConidAlgosGet) | **GET** /iserver/contract/{conid}/algos | IB Algo Params |
| [**iserverContractConidInfoAndRulesGet**](ContractApi.md#iserverContractConidInfoAndRulesGet) | **GET** /iserver/contract/{conid}/info-and-rules | Info and Rules |
| [**iserverContractConidInfoGet**](ContractApi.md#iserverContractConidInfoGet) | **GET** /iserver/contract/{conid}/info | Contract Details |
| [**iserverContractRulesPost**](ContractApi.md#iserverContractRulesPost) | **POST** /iserver/contract/rules | Contract Rules |
| [**iserverSecdefInfoGet**](ContractApi.md#iserverSecdefInfoGet) | **GET** /iserver/secdef/info | Secdef Info |
| [**iserverSecdefStrikesGet**](ContractApi.md#iserverSecdefStrikesGet) | **GET** /iserver/secdef/strikes | Search Strikes |
| [**searchBySymbolOrName**](ContractApi.md#searchBySymbolOrName) | **POST** /iserver/secdef/search | Search by Symbol or Name |


<a name="getFutures"></a>
# **getFutures**
> getFutures_200_response getFutures(symbols)

Security Futures by Symbol

    Returns a list of non-expired future contracts for given symbol(s)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **symbols** | **String**| list of case-sensitive symbols separated by comma | [default to null] |

### Return type

[**getFutures_200_response**](../Models/getFutures_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSecdef"></a>
# **getSecdef**
> List getSecdef(body)

Secdef by Conid

    Returns a list of security definitions for the given conids

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **body** | [**getSecdef_request**](../Models/getSecdef_request.md)| request body | |

### Return type

[**List**](../Models/secdef_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSecdefSchedule"></a>
# **getSecdefSchedule**
> getSecdefSchedule_200_response getSecdefSchedule(assetClass, symbol, exchange, exchangeFilter)

Get trading schedule for symbol

    Returns the trading schedule up to a month for the requested contract

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **assetClass** | **String**| specify the asset class of the contract. Available values-- Stock: STK, Option: OPT, Future: FUT, Contract For Difference: CFD, Warrant: WAR, Forex: SWP, Mutual Fund: FND, Bond: BND, Inter-Commodity Spreads: ICS  | [default to null] |
| **symbol** | **String**| Underlying Symbol for specified contract, for example &#39;AAPL&#39; for US Stock - Apple Inc. | [default to null] |
| **exchange** | **String**| Native exchange for contract, for example &#39;NASDAQ&#39; for US Stock - Apple Inc. | [optional] [default to null] |
| **exchangeFilter** | **String**| Response only returns trading schedule for specified exchange | [optional] [default to null] |

### Return type

[**getSecdefSchedule_200_response**](../Models/getSecdefSchedule_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getStocks"></a>
# **getStocks**
> getStocks_200_response getStocks(symbols)

Security Stocks by Symbol

    Returns an object contains all stock contracts for given symbol(s)

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **symbols** | **String**| list of upper-sensitive symbols separated by comma | [default to null] |

### Return type

[**getStocks_200_response**](../Models/getStocks_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverContractConidAlgosGet"></a>
# **iserverContractConidAlgosGet**
> List iserverContractConidAlgosGet(conid, algos, addDescription, addParams)

IB Algo Params

    Returns supported IB Algos for contract. Must be called a second time to query the list of available parameters.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **String**| IBKR contract identifier | [default to null] |
| **algos** | **String**| List of algo ids delimited by \&quot;;\&quot; to filter by. Max of 8 algos ids can be specified. | [optional] [default to null] |
| **addDescription** | **String**| Whether or not to add algo descriptions to response. Set to 1 for yes, 0 for no. | [optional] [default to null] |
| **addParams** | **String**| Whether or not to show algo parameters.  Set to 1 for yes, 0 for no. | [optional] [default to null] |

### Return type

[**List**](../Models/_iserver_contract__conid__algos_get_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverContractConidInfoAndRulesGet"></a>
# **iserverContractConidInfoAndRulesGet**
> _iserver_contract__conid__info_and_rules_get_200_response iserverContractConidInfoAndRulesGet(conid, isBuy)

Info and Rules

    Returns both contract info and rules from a single endpoint. For only contract rules, use the endpoint /iserver/contract/rules. For only contract info, use the endpoint /iserver/contract/{conid}/info. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **String**| IBKR contract identifier | [default to null] |
| **isBuy** | **Boolean**| Side of the market rules apply too. Set to **true** for Buy Orders, set to **false** for Sell Orders | [default to null] |

### Return type

[**_iserver_contract__conid__info_and_rules_get_200_response**](../Models/_iserver_contract__conid__info_and_rules_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverContractConidInfoGet"></a>
# **iserverContractConidInfoGet**
> contract iserverContractConidInfoGet(conid)

Contract Details

    Using the Contract Identifier get contract info. You can use this to prefill your order before you submit an order

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **String**| contract id | [default to null] |

### Return type

[**contract**](../Models/contract.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverContractRulesPost"></a>
# **iserverContractRulesPost**
> _iserver_contract_rules_post_200_response iserverContractRulesPost(conid)

Contract Rules

    Returns trading related rules for a specific contract and side. For both contract info and rules use the endpoint /iserver/contract/{conid}/info-and-rules.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | [**_iserver_contract_rules_post_request**](../Models/_iserver_contract_rules_post_request.md)|  | |

### Return type

[**_iserver_contract_rules_post_200_response**](../Models/_iserver_contract_rules_post_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverSecdefInfoGet"></a>
# **iserverSecdefInfoGet**
> List iserverSecdefInfoGet(conid, sectype, month, exchange, strike, right)

Secdef Info

    Provides Contract Details of Futures, Options, Warrants, Cash and CFDs based on conid. To get the strike price for Options/Warrants use \&quot;/iserver/secdef/strikes\&quot; endpoint. Must call /secdef/search for the underlying contract first.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **String**| underlying contract id | [default to null] |
| **sectype** | **String**| FUT/OPT/WAR/CASH/CFD | [default to null] |
| **month** | **String**| contract month, only required for FUT/OPT/WAR in the format MMMYY, example JAN00 | [optional] [default to null] |
| **exchange** | **String**| optional, default is SMART | [optional] [default to null] |
| **strike** | **BigDecimal**| optional, only required for OPT/WAR | [optional] [default to null] |
| **right** | **String**| C for call, P for put | [optional] [default to null] |

### Return type

[**List**](../Models/secdef-info.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="iserverSecdefStrikesGet"></a>
# **iserverSecdefStrikesGet**
> _iserver_secdef_strikes_get_200_response iserverSecdefStrikesGet(conid, sectype, month, exchange)

Search Strikes

    Query strikes for Options/Warrants. For the conid of the underlying contract, available contract months and exchanges use \&quot;/iserver/secdef/search\&quot;

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **conid** | **String**| contract id of the underlying contract | [default to null] |
| **sectype** | **String**| OPT/WAR | [default to null] |
| **month** | **String**| contract month | [default to null] |
| **exchange** | **String**| optional, default is SMART | [optional] [default to null] |

### Return type

[**_iserver_secdef_strikes_get_200_response**](../Models/_iserver_secdef_strikes_get_200_response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="searchBySymbolOrName"></a>
# **searchBySymbolOrName**
> List searchBySymbolOrName(symbol)

Search by Symbol or Name

    Search by underlying symbol or company name. Relays back what derivative contract(s) it has. This endpoint must be called before using /secdef/info. If company name is specified will only receive limited response: conid, companyName, companyHeader and symbol. 

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **symbol** | [**SearchBySymbolOrName_request**](../Models/SearchBySymbolOrName_request.md)| Symbol or Company Name to be searched | |

### Return type

[**List**](../Models/SearchBySymbolOrName_200_response_inner.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

