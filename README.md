# Documentation for Client Portal Web API

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to *http://localhost:5000/v1/api*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AccountApi* | [**getLedger**](Apis/AccountApi.md#getledger) | **GET** /portfolio/{accountId}/ledger | Account Ledger |
*AccountApi* | [**iserverAccountPnlPartitionedGet**](Apis/AccountApi.md#iserveraccountpnlpartitionedget) | **GET** /iserver/account/pnl/partitioned | PnL for the selected account |
*AccountApi* | [**iserverAccountPost**](Apis/AccountApi.md#iserveraccountpost) | **POST** /iserver/account | Switch Account |
*AccountApi* | [**iserverAccountsGet**](Apis/AccountApi.md#iserveraccountsget) | **GET** /iserver/accounts | Brokerage Accounts |
*AccountApi* | [**portfolioAccountIdMetaGet**](Apis/AccountApi.md#portfolioaccountidmetaget) | **GET** /portfolio/{accountId}/meta | Account Information |
*AccountApi* | [**portfolioAccountIdSummaryGet**](Apis/AccountApi.md#portfolioaccountidsummaryget) | **GET** /portfolio/{accountId}/summary | Account Summary |
*AccountApi* | [**portfolioAccountsGet**](Apis/AccountApi.md#portfolioaccountsget) | **GET** /portfolio/accounts | Portfolio Accounts |
*AccountApi* | [**portfolioSubaccounts2Get**](Apis/AccountApi.md#portfoliosubaccounts2get) | **GET** /portfolio/subaccounts2 | List of Sub-Accounts (Large Accounts) |
*AccountApi* | [**portfolioSubaccountsGet**](Apis/AccountApi.md#portfoliosubaccountsget) | **GET** /portfolio/subaccounts | List of Sub-Accounts |
| *AlertApi* | [**iserverAccountAccountIdAlertActivatePost**](Apis/AlertApi.md#iserveraccountaccountidalertactivatepost) | **POST** /iserver/account/{accountId}/alert/activate | Activate or deactivate an alert |
*AlertApi* | [**iserverAccountAccountIdAlertAlertIdDelete**](Apis/AlertApi.md#iserveraccountaccountidalertalertiddelete) | **DELETE** /iserver/account/{accountId}/alert/{alertId} | Delete an alert |
*AlertApi* | [**iserverAccountAccountIdAlertPost**](Apis/AlertApi.md#iserveraccountaccountidalertpost) | **POST** /iserver/account/{accountId}/alert | Create or modify alert |
*AlertApi* | [**iserverAccountAccountIdAlertsGet**](Apis/AlertApi.md#iserveraccountaccountidalertsget) | **GET** /iserver/account/{accountId}/alerts | Get a list of available alerts |
*AlertApi* | [**iserverAccountAlertIdGet**](Apis/AlertApi.md#iserveraccountalertidget) | **GET** /iserver/account/alert/{id} | Get details of an alert |
*AlertApi* | [**iserverAccountMtaGet**](Apis/AlertApi.md#iserveraccountmtaget) | **GET** /iserver/account/mta | Get MTA alert |
| *CCPBetaApi* | [**completeCCP**](Apis/CCPBetaApi.md#completeccp) | **POST** /ccp/auth/response | Complete CCP Session |
*CCPBetaApi* | [**deleteOrder**](Apis/CCPBetaApi.md#deleteorder) | **DELETE** /ccp/order | Delete Order |
*CCPBetaApi* | [**getCCPAccount**](Apis/CCPBetaApi.md#getccpaccount) | **GET** /ccp/account | Brokerage Accounts |
*CCPBetaApi* | [**getCCPOrders**](Apis/CCPBetaApi.md#getccporders) | **GET** /ccp/orders | Order Status |
*CCPBetaApi* | [**getCCPPositions**](Apis/CCPBetaApi.md#getccppositions) | **GET** /ccp/positions | Positions |
*CCPBetaApi* | [**getCCPStatus**](Apis/CCPBetaApi.md#getccpstatus) | **GET** /ccp/status | CCP Status |
*CCPBetaApi* | [**getCCPTrades**](Apis/CCPBetaApi.md#getccptrades) | **GET** /ccp/trades | Trades |
*CCPBetaApi* | [**initCCP**](Apis/CCPBetaApi.md#initccp) | **POST** /ccp/auth/init | Start CCP Session |
*CCPBetaApi* | [**submitOrder**](Apis/CCPBetaApi.md#submitorder) | **POST** /ccp/order | Submit Order |
*CCPBetaApi* | [**updateOrder**](Apis/CCPBetaApi.md#updateorder) | **PUT** /ccp/order | Update Order |
| *ContractApi* | [**getFutures**](Apis/ContractApi.md#getfutures) | **GET** /trsrv/futures | Security Futures by Symbol |
*ContractApi* | [**getSecdef**](Apis/ContractApi.md#getsecdef) | **POST** /trsrv/secdef | Secdef by Conid |
*ContractApi* | [**getSecdefSchedule**](Apis/ContractApi.md#getsecdefschedule) | **GET** /trsrv/secdef/schedule | Get trading schedule for symbol |
*ContractApi* | [**getStocks**](Apis/ContractApi.md#getstocks) | **GET** /trsrv/stocks | Security Stocks by Symbol |
*ContractApi* | [**iserverContractConidAlgosGet**](Apis/ContractApi.md#iservercontractconidalgosget) | **GET** /iserver/contract/{conid}/algos | IB Algo Params |
*ContractApi* | [**iserverContractConidInfoAndRulesGet**](Apis/ContractApi.md#iservercontractconidinfoandrulesget) | **GET** /iserver/contract/{conid}/info-and-rules | Info and Rules |
*ContractApi* | [**iserverContractConidInfoGet**](Apis/ContractApi.md#iservercontractconidinfoget) | **GET** /iserver/contract/{conid}/info | Contract Details |
*ContractApi* | [**iserverContractRulesPost**](Apis/ContractApi.md#iservercontractrulespost) | **POST** /iserver/contract/rules | Contract Rules |
*ContractApi* | [**iserverSecdefInfoGet**](Apis/ContractApi.md#iserversecdefinfoget) | **GET** /iserver/secdef/info | Secdef Info |
*ContractApi* | [**iserverSecdefStrikesGet**](Apis/ContractApi.md#iserversecdefstrikesget) | **GET** /iserver/secdef/strikes | Search Strikes |
*ContractApi* | [**searchBySymbolOrName**](Apis/ContractApi.md#searchbysymbolorname) | **POST** /iserver/secdef/search | Search by Symbol or Name |
| *FYIApi* | [**deleteDevice**](Apis/FYIApi.md#deletedevice) | **DELETE** /fyi/deliveryoptions/{deviceId} | Delete a device |
*FYIApi* | [**enableDisableDevice**](Apis/FYIApi.md#enabledisabledevice) | **POST** /fyi/deliveryoptions/device | Enable/Disable device option |
*FYIApi* | [**enableDisableEmail**](Apis/FYIApi.md#enabledisableemail) | **PUT** /fyi/deliveryoptions/email | Enable/Disable email option |
*FYIApi* | [**enableDisableSubscription**](Apis/FYIApi.md#enabledisablesubscription) | **POST** /fyi/settings/{typecode} | Enable/Disable certain subscription |
*FYIApi* | [**getDeliveryOptions**](Apis/FYIApi.md#getdeliveryoptions) | **GET** /fyi/deliveryoptions | Get delivery options |
*FYIApi* | [**getDisclaimer**](Apis/FYIApi.md#getdisclaimer) | **GET** /fyi/disclaimer/{typecode} | Get disclaimer for a certain kind of fyi |
*FYIApi* | [**getMoreNotifications**](Apis/FYIApi.md#getmorenotifications) | **GET** /fyi/notifications/more | Get more notifications based on a certain one |
*FYIApi* | [**getNotifications**](Apis/FYIApi.md#getnotifications) | **GET** /fyi/notifications | Get a list of notifications |
*FYIApi* | [**getSettings**](Apis/FYIApi.md#getsettings) | **GET** /fyi/settings | Get a list of subscriptions |
*FYIApi* | [**getUnreadNumber**](Apis/FYIApi.md#getunreadnumber) | **GET** /fyi/unreadnumber | Get unread number of fyis. The HTTP method POST is also supported. |
*FYIApi* | [**markDisclaimerRead**](Apis/FYIApi.md#markdisclaimerread) | **PUT** /fyi/disclaimer/{typecode} | Mark disclaimer read |
*FYIApi* | [**markNotificationRead**](Apis/FYIApi.md#marknotificationread) | **PUT** /fyi/notifications/{notificationId} | Get a list of notifications |
| *MarketDataApi* | [**getHistory**](Apis/MarketDataApi.md#gethistory) | **GET** /hmds/history | Market Data History (Beta) |
*MarketDataApi* | [**getSnapshot**](Apis/MarketDataApi.md#getsnapshot) | **GET** /md/snapshot | Market Data Snapshot (Beta) |
*MarketDataApi* | [**iserverMarketdataConidUnsubscribeGet**](Apis/MarketDataApi.md#iservermarketdataconidunsubscribeget) | **GET** /iserver/marketdata/{conid}/unsubscribe | Market Data Cancel (Single) |
*MarketDataApi* | [**iserverMarketdataHistoryGet**](Apis/MarketDataApi.md#iservermarketdatahistoryget) | **GET** /iserver/marketdata/history | Market Data History |
*MarketDataApi* | [**iserverMarketdataSnapshotGet**](Apis/MarketDataApi.md#iservermarketdatasnapshotget) | **GET** /iserver/marketdata/snapshot | Market Data |
*MarketDataApi* | [**iserverMarketdataUnsubscribeallGet**](Apis/MarketDataApi.md#iservermarketdataunsubscribeallget) | **GET** /iserver/marketdata/unsubscribeall | Market Data Cancel (All) |
| *OrderApi* | [**getLiveOrders**](Apis/OrderApi.md#getliveorders) | **GET** /iserver/account/orders | Live Orders |
*OrderApi* | [**iserverAccountAccountIdOrderOrderIdDelete**](Apis/OrderApi.md#iserveraccountaccountidorderorderiddelete) | **DELETE** /iserver/account/{accountId}/order/{orderId} | Cancel Order |
*OrderApi* | [**iserverAccountAccountIdOrderOrderIdPost**](Apis/OrderApi.md#iserveraccountaccountidorderorderidpost) | **POST** /iserver/account/{accountId}/order/{orderId} | Modify Order |
*OrderApi* | [**iserverAccountAccountIdOrderWhatifPost**](Apis/OrderApi.md#iserveraccountaccountidorderwhatifpost) | **POST** /iserver/account/{accountId}/order/whatif | Preview Order (Deprecated) |
*OrderApi* | [**iserverAccountAccountIdOrdersWhatifPost**](Apis/OrderApi.md#iserveraccountaccountidorderswhatifpost) | **POST** /iserver/account/{accountId}/orders/whatif | Preview Orders |
*OrderApi* | [**iserverAccountOrderStatusOrderIdGet**](Apis/OrderApi.md#iserveraccountorderstatusorderidget) | **GET** /iserver/account/order/status/{orderId} | Order Status |
*OrderApi* | [**iserverAccountOrdersFaGroupPost**](Apis/OrderApi.md#iserveraccountordersfagrouppost) | **POST** /iserver/account/orders/{faGroup} | Place Orders for FA |
*OrderApi* | [**iserverReplyReplyidPost**](Apis/OrderApi.md#iserverreplyreplyidpost) | **POST** /iserver/reply/{replyid} | Place Order Reply |
*OrderApi* | [**placeOrder**](Apis/OrderApi.md#placeorder) | **POST** /iserver/account/{accountId}/orders | Place Orders |
*OrderApi* | [**placeOrderDeprecated**](Apis/OrderApi.md#placeorderdeprecated) | **POST** /iserver/account/{accountId}/order | Place Order (Deprecated) |
| *PnLApi* | [**iserverAccountPnlPartitionedGet**](Apis/PnLApi.md#iserveraccountpnlpartitionedget) | **GET** /iserver/account/pnl/partitioned | PnL for the selected account |
| *PortfolioApi* | [**getLedger**](Apis/PortfolioApi.md#getledger) | **GET** /portfolio/{accountId}/ledger | Account Ledger |
*PortfolioApi* | [**portfolioAccountIdAllocationGet**](Apis/PortfolioApi.md#portfolioaccountidallocationget) | **GET** /portfolio/{accountId}/allocation | Account Allocation |
*PortfolioApi* | [**portfolioAccountIdMetaGet**](Apis/PortfolioApi.md#portfolioaccountidmetaget) | **GET** /portfolio/{accountId}/meta | Account Information |
*PortfolioApi* | [**portfolioAccountIdPositionConidGet**](Apis/PortfolioApi.md#portfolioaccountidpositionconidget) | **GET** /portfolio/{accountId}/position/{conid} | Position by Conid |
*PortfolioApi* | [**portfolioAccountIdPositionsInvalidatePost**](Apis/PortfolioApi.md#portfolioaccountidpositionsinvalidatepost) | **POST** /portfolio/{accountId}/positions/invalidate | Invalidates the backend cache of the Portfolio |
*PortfolioApi* | [**portfolioAccountIdPositionsPageIdGet**](Apis/PortfolioApi.md#portfolioaccountidpositionspageidget) | **GET** /portfolio/{accountId}/positions/{pageId} | Portfolio Positions |
*PortfolioApi* | [**portfolioAccountIdSummaryGet**](Apis/PortfolioApi.md#portfolioaccountidsummaryget) | **GET** /portfolio/{accountId}/summary | Account Summary |
*PortfolioApi* | [**portfolioAccountsGet**](Apis/PortfolioApi.md#portfolioaccountsget) | **GET** /portfolio/accounts | Portfolio Accounts |
*PortfolioApi* | [**portfolioAllocationPost**](Apis/PortfolioApi.md#portfolioallocationpost) | **POST** /portfolio/allocation | Account Alloction (All Accounts) |
*PortfolioApi* | [**portfolioPositionsConidGet**](Apis/PortfolioApi.md#portfoliopositionsconidget) | **GET** /portfolio/positions/{conid} | Positions by Conid |
*PortfolioApi* | [**portfolioSubaccounts2Get**](Apis/PortfolioApi.md#portfoliosubaccounts2get) | **GET** /portfolio/subaccounts2 | List of Sub-Accounts (Large Accounts) |
*PortfolioApi* | [**portfolioSubaccountsGet**](Apis/PortfolioApi.md#portfoliosubaccountsget) | **GET** /portfolio/subaccounts | List of Sub-Accounts |
| *PortfolioAnalystApi* | [**getPerformance**](Apis/PortfolioAnalystApi.md#getperformance) | **POST** /pa/performance | Account Performance |
*PortfolioAnalystApi* | [**paSummaryPost**](Apis/PortfolioAnalystApi.md#pasummarypost) | **POST** /pa/summary | Account Balance's Summary (Deprecated) |
*PortfolioAnalystApi* | [**paTransactionsPost**](Apis/PortfolioAnalystApi.md#patransactionspost) | **POST** /pa/transactions | Position's Transaction History |
| *ScannerApi* | [**iserverScannerParamsGet**](Apis/ScannerApi.md#iserverscannerparamsget) | **GET** /iserver/scanner/params | Scanner Parameters |
*ScannerApi* | [**iserverScannerRunPost**](Apis/ScannerApi.md#iserverscannerrunpost) | **POST** /iserver/scanner/run | Scanner Run |
*ScannerApi* | [**runScanner**](Apis/ScannerApi.md#runscanner) | **POST** /hmds/scanner | Run Scanner (Beta) |
| *SessionApi* | [**iserverAuthStatusPost**](Apis/SessionApi.md#iserverauthstatuspost) | **POST** /iserver/auth/status | Authentication Status |
*SessionApi* | [**iserverReauthenticatePost**](Apis/SessionApi.md#iserverreauthenticatepost) | **POST** /iserver/reauthenticate | Tries to re-authenticate to Brokerage |
*SessionApi* | [**logout**](Apis/SessionApi.md#logout) | **POST** /logout | Ends the current session |
*SessionApi* | [**tickle**](Apis/SessionApi.md#tickle) | **POST** /tickle | Ping the server to keep the session open |
*SessionApi* | [**validateSSO**](Apis/SessionApi.md#validatesso) | **GET** /sso/validate | Validate SSO |
| *StreamingApi* | [**websocket**](Apis/StreamingApi.md#websocket) | **POST** /ws | Websocket Endpoint |
| *TradesApi* | [**iserverAccountTradesGet**](Apis/TradesApi.md#iserveraccounttradesget) | **GET** /iserver/account/trades | List of Trades for the selected account |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [PlaceOrderDeprecated_200_response_inner](./Models/PlaceOrderDeprecated_200_response_inner.md)
 - [PlaceOrder_request](./Models/PlaceOrder_request.md)
 - [SearchBySymbolOrName_200_response_inner](./Models/SearchBySymbolOrName_200_response_inner.md)
 - [SearchBySymbolOrName_200_response_inner_sections_inner](./Models/SearchBySymbolOrName_200_response_inner_sections_inner.md)
 - [SearchBySymbolOrName_request](./Models/SearchBySymbolOrName_request.md)
 - [_iserver_account__accountId__alert_activate_post_200_response](./Models/_iserver_account__accountId__alert_activate_post_200_response.md)
 - [_iserver_account__accountId__alert_activate_post_request](./Models/_iserver_account__accountId__alert_activate_post_request.md)
 - [_iserver_account__accountId__alert_post_200_response](./Models/_iserver_account__accountId__alert_post_200_response.md)
 - [_iserver_account__accountId__alerts_get_200_response_inner](./Models/_iserver_account__accountId__alerts_get_200_response_inner.md)
 - [_iserver_account__accountId__order__orderId__delete_200_response](./Models/_iserver_account__accountId__order__orderId__delete_200_response.md)
 - [_iserver_account__accountId__order__orderId__post_200_response_inner](./Models/_iserver_account__accountId__order__orderId__post_200_response_inner.md)
 - [_iserver_account__accountId__order_whatif_post_200_response](./Models/_iserver_account__accountId__order_whatif_post_200_response.md)
 - [_iserver_account__accountId__order_whatif_post_200_response_amount](./Models/_iserver_account__accountId__order_whatif_post_200_response_amount.md)
 - [_iserver_account__accountId__order_whatif_post_200_response_equity](./Models/_iserver_account__accountId__order_whatif_post_200_response_equity.md)
 - [_iserver_account_pnl_partitioned_get_200_response](./Models/_iserver_account_pnl_partitioned_get_200_response.md)
 - [_iserver_account_post_200_response](./Models/_iserver_account_post_200_response.md)
 - [_iserver_accounts_get_200_response](./Models/_iserver_accounts_get_200_response.md)
 - [_iserver_contract__conid__algos_get_200_response_inner](./Models/_iserver_contract__conid__algos_get_200_response_inner.md)
 - [_iserver_contract__conid__algos_get_200_response_inner_parameters_inner](./Models/_iserver_contract__conid__algos_get_200_response_inner_parameters_inner.md)
 - [_iserver_contract__conid__info_and_rules_get_200_response](./Models/_iserver_contract__conid__info_and_rules_get_200_response.md)
 - [_iserver_contract__conid__info_and_rules_get_200_response_rules_inner](./Models/_iserver_contract__conid__info_and_rules_get_200_response_rules_inner.md)
 - [_iserver_contract_rules_post_200_response](./Models/_iserver_contract_rules_post_200_response.md)
 - [_iserver_contract_rules_post_200_response_rules_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_canTradeAcctIds_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_canTradeAcctIds_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_cqtTypes_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_cqtTypes_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_fraqTypes_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_fraqTypes_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_ibalgoTypes_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_ibalgoTypes_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_orderDefaults_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_orderDefaults_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_orderDefaults_inner_string_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_orderDefaults_inner_string_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_orderTypesOutside_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_orderTypesOutside_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_orderTypes_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_orderTypes_inner.md)
 - [_iserver_contract_rules_post_200_response_rules_inner_tifTypes_inner](./Models/_iserver_contract_rules_post_200_response_rules_inner_tifTypes_inner.md)
 - [_iserver_contract_rules_post_request](./Models/_iserver_contract_rules_post_request.md)
 - [_iserver_marketdata__conid__unsubscribe_get_200_response](./Models/_iserver_marketdata__conid__unsubscribe_get_200_response.md)
 - [_iserver_marketdata_snapshot_get_200_response_inner](./Models/_iserver_marketdata_snapshot_get_200_response_inner.md)
 - [_iserver_marketdata_snapshot_get_400_response](./Models/_iserver_marketdata_snapshot_get_400_response.md)
 - [_iserver_marketdata_unsubscribeall_get_200_response](./Models/_iserver_marketdata_unsubscribeall_get_200_response.md)
 - [_iserver_reply__replyid__post_200_response_inner](./Models/_iserver_reply__replyid__post_200_response_inner.md)
 - [_iserver_reply__replyid__post_400_response](./Models/_iserver_reply__replyid__post_400_response.md)
 - [_iserver_reply__replyid__post_request](./Models/_iserver_reply__replyid__post_request.md)
 - [_iserver_scanner_params_get_200_response](./Models/_iserver_scanner_params_get_200_response.md)
 - [_iserver_scanner_params_get_200_response_filter_list_inner](./Models/_iserver_scanner_params_get_200_response_filter_list_inner.md)
 - [_iserver_scanner_params_get_200_response_instrument_list_inner](./Models/_iserver_scanner_params_get_200_response_instrument_list_inner.md)
 - [_iserver_scanner_params_get_200_response_location_tree_inner](./Models/_iserver_scanner_params_get_200_response_location_tree_inner.md)
 - [_iserver_scanner_params_get_200_response_location_tree_inner_locations_inner](./Models/_iserver_scanner_params_get_200_response_location_tree_inner_locations_inner.md)
 - [_iserver_scanner_params_get_200_response_scan_type_list_inner](./Models/_iserver_scanner_params_get_200_response_scan_type_list_inner.md)
 - [_iserver_scanner_run_post_200_response_inner](./Models/_iserver_scanner_run_post_200_response_inner.md)
 - [_iserver_secdef_strikes_get_200_response](./Models/_iserver_secdef_strikes_get_200_response.md)
 - [_pa_summary_post_request](./Models/_pa_summary_post_request.md)
 - [_pa_transactions_post_request](./Models/_pa_transactions_post_request.md)
 - [_portfolio__accountId__summary_get_200_response](./Models/_portfolio__accountId__summary_get_200_response.md)
 - [_portfolio_positions__conid__get_200_response](./Models/_portfolio_positions__conid__get_200_response.md)
 - [_portfolio_subaccounts2_get_200_response](./Models/_portfolio_subaccounts2_get_200_response.md)
 - [_portfolio_subaccounts2_get_200_response_metadata](./Models/_portfolio_subaccounts2_get_200_response_metadata.md)
 - [_portfolio_subaccounts2_get_200_response_subaccounts_inner](./Models/_portfolio_subaccounts2_get_200_response_subaccounts_inner.md)
 - [account](./Models/account.md)
 - [account_parent](./Models/account_parent.md)
 - [alert-request](./Models/alert-request.md)
 - [alert-response](./Models/alert-response.md)
 - [alert_request_conditions_inner](./Models/alert_request_conditions_inner.md)
 - [alert_response_conditions_inner](./Models/alert_response_conditions_inner.md)
 - [allocation_inner](./Models/allocation_inner.md)
 - [allocation_inner_assetClass](./Models/allocation_inner_assetClass.md)
 - [allocation_inner_assetClass_long](./Models/allocation_inner_assetClass_long.md)
 - [allocation_inner_assetClass_short](./Models/allocation_inner_assetClass_short.md)
 - [allocation_inner_group](./Models/allocation_inner_group.md)
 - [allocation_inner_group_long](./Models/allocation_inner_group_long.md)
 - [allocation_inner_group_short](./Models/allocation_inner_group_short.md)
 - [allocation_inner_sector](./Models/allocation_inner_sector.md)
 - [allocation_inner_sector_long](./Models/allocation_inner_sector_long.md)
 - [allocation_inner_sector_short](./Models/allocation_inner_sector_short.md)
 - [authStatus](./Models/authStatus.md)
 - [calendar_request](./Models/calendar_request.md)
 - [calendar_request_date](./Models/calendar_request_date.md)
 - [calendar_request_filters](./Models/calendar_request_filters.md)
 - [completeCCP_200_response](./Models/completeCCP_200_response.md)
 - [completeCCP_request](./Models/completeCCP_request.md)
 - [contract](./Models/contract.md)
 - [contract_rules](./Models/contract_rules.md)
 - [enableDisableDevice_request](./Models/enableDisableDevice_request.md)
 - [enableDisableSubscription_request](./Models/enableDisableSubscription_request.md)
 - [events_inner](./Models/events_inner.md)
 - [futures_inner](./Models/futures_inner.md)
 - [getCCPAccount_200_response](./Models/getCCPAccount_200_response.md)
 - [getCCPAccount_200_response_acctList_inner](./Models/getCCPAccount_200_response_acctList_inner.md)
 - [getCCPOrders_200_response](./Models/getCCPOrders_200_response.md)
 - [getCCPStatus_200_response](./Models/getCCPStatus_200_response.md)
 - [getDeliveryOptions_200_response](./Models/getDeliveryOptions_200_response.md)
 - [getDeliveryOptions_200_response_E_inner](./Models/getDeliveryOptions_200_response_E_inner.md)
 - [getDisclaimer_200_response](./Models/getDisclaimer_200_response.md)
 - [getFutures_200_response](./Models/getFutures_200_response.md)
 - [getFutures_500_response](./Models/getFutures_500_response.md)
 - [getLedger_200_response](./Models/getLedger_200_response.md)
 - [getLiveOrders_200_response](./Models/getLiveOrders_200_response.md)
 - [getLiveOrders_200_response_orders_inner](./Models/getLiveOrders_200_response_orders_inner.md)
 - [getPerformance_request](./Models/getPerformance_request.md)
 - [getSecdefSchedule_200_response](./Models/getSecdefSchedule_200_response.md)
 - [getSecdefSchedule_200_response_schedules_inner](./Models/getSecdefSchedule_200_response_schedules_inner.md)
 - [getSecdefSchedule_200_response_schedules_inner_sessions](./Models/getSecdefSchedule_200_response_schedules_inner_sessions.md)
 - [getSecdefSchedule_200_response_schedules_inner_tradingTimes](./Models/getSecdefSchedule_200_response_schedules_inner_tradingTimes.md)
 - [getSecdef_request](./Models/getSecdef_request.md)
 - [getSettings_200_response_inner](./Models/getSettings_200_response_inner.md)
 - [getStocks_200_response](./Models/getStocks_200_response.md)
 - [getUnreadNumber_200_response](./Models/getUnreadNumber_200_response.md)
 - [history-data](./Models/history-data.md)
 - [history-result](./Models/history-result.md)
 - [history_data_data_inner](./Models/history_data_data_inner.md)
 - [history_result_bars](./Models/history_result_bars.md)
 - [history_result_bars_data_inner](./Models/history_result_bars_data_inner.md)
 - [inds_inner](./Models/inds_inner.md)
 - [initCCP_200_response](./Models/initCCP_200_response.md)
 - [ledger](./Models/ledger.md)
 - [logout_200_response](./Models/logout_200_response.md)
 - [markDisclaimerRead_200_response](./Models/markDisclaimerRead_200_response.md)
 - [market-data](./Models/market-data.md)
 - [modify-order](./Models/modify-order.md)
 - [notifications_inner](./Models/notifications_inner.md)
 - [order](./Models/order.md)
 - [order-data](./Models/order-data.md)
 - [order-request](./Models/order-request.md)
 - [order-status](./Models/order-status.md)
 - [order_data_warnings](./Models/order_data_warnings.md)
 - [performance](./Models/performance.md)
 - [performance_cps](./Models/performance_cps.md)
 - [performance_cps_data_inner](./Models/performance_cps_data_inner.md)
 - [performance_nav](./Models/performance_nav.md)
 - [performance_tpps](./Models/performance_tpps.md)
 - [position-data](./Models/position-data.md)
 - [position_inner](./Models/position_inner.md)
 - [runScanner_request](./Models/runScanner_request.md)
 - [runScanner_request_filters_inner](./Models/runScanner_request_filters_inner.md)
 - [scanner-params](./Models/scanner-params.md)
 - [scanner-result](./Models/scanner-result.md)
 - [scanner_params_filter_inner](./Models/scanner_params_filter_inner.md)
 - [scanner_result_Contracts](./Models/scanner_result_Contracts.md)
 - [scanner_result_Contracts_Contract_inner](./Models/scanner_result_Contracts_Contract_inner.md)
 - [secdef-info](./Models/secdef-info.md)
 - [secdef_inner](./Models/secdef_inner.md)
 - [secdef_inner_incrementRules](./Models/secdef_inner_incrementRules.md)
 - [set-account](./Models/set-account.md)
 - [stats-data](./Models/stats-data.md)
 - [stocks_inner](./Models/stocks_inner.md)
 - [stocks_inner_contracts_inner](./Models/stocks_inner_contracts_inner.md)
 - [summary](./Models/summary.md)
 - [system-error](./Models/system-error.md)
 - [trade](./Models/trade.md)
 - [transactions](./Models/transactions.md)
 - [transactions_transactions_inner](./Models/transactions_transactions_inner.md)
 - [validateSSO_200_response](./Models/validateSSO_200_response.md)
 - [wagers_inner](./Models/wagers_inner.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

<a name="api_key"></a>
### api_key

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

