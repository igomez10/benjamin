# TransactionsTransactionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acctid** | Pointer to **string** |  | [optional] 
**Conid** | Pointer to **float32** |  | [optional] 
**Cur** | Pointer to **string** | currency code | [optional] 
**FxRate** | Pointer to **float32** | Conversion rate from asset currency to response currency | [optional] 
**Desc** | Pointer to **string** | Transaction description | [optional] 
**Date** | Pointer to **string** | Date of transaction.  Epoch time, GMT | [optional] 
**Type** | Pointer to **string** | Transaction Type Name: Examples: \&quot;Sell\&quot;, \&quot;Buy\&quot;, \&quot;Corporate Action\&quot;, \&quot;Dividend Payment\&quot;, \&quot;Transfer\&quot;, \&quot;Payment in Lieu\&quot; Dividends and Transfers do not have price and quantity in response  | [optional] 
**Qty** | Pointer to **float32** | Not applicable for all transaction types | [optional] 
**Pr** | Pointer to **float32** | In asset currency. Not be applicable for all transaction types. | [optional] 
**Amt** | Pointer to **float32** | Raw value, no formatting. Net transaction amount (may include commission, tax). In asset currency | [optional] 

## Methods

### NewTransactionsTransactionsInner

`func NewTransactionsTransactionsInner() *TransactionsTransactionsInner`

NewTransactionsTransactionsInner instantiates a new TransactionsTransactionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsTransactionsInnerWithDefaults

`func NewTransactionsTransactionsInnerWithDefaults() *TransactionsTransactionsInner`

NewTransactionsTransactionsInnerWithDefaults instantiates a new TransactionsTransactionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctid

`func (o *TransactionsTransactionsInner) GetAcctid() string`

GetAcctid returns the Acctid field if non-nil, zero value otherwise.

### GetAcctidOk

`func (o *TransactionsTransactionsInner) GetAcctidOk() (*string, bool)`

GetAcctidOk returns a tuple with the Acctid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctid

`func (o *TransactionsTransactionsInner) SetAcctid(v string)`

SetAcctid sets Acctid field to given value.

### HasAcctid

`func (o *TransactionsTransactionsInner) HasAcctid() bool`

HasAcctid returns a boolean if a field has been set.

### GetConid

`func (o *TransactionsTransactionsInner) GetConid() float32`

GetConid returns the Conid field if non-nil, zero value otherwise.

### GetConidOk

`func (o *TransactionsTransactionsInner) GetConidOk() (*float32, bool)`

GetConidOk returns a tuple with the Conid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConid

`func (o *TransactionsTransactionsInner) SetConid(v float32)`

SetConid sets Conid field to given value.

### HasConid

`func (o *TransactionsTransactionsInner) HasConid() bool`

HasConid returns a boolean if a field has been set.

### GetCur

`func (o *TransactionsTransactionsInner) GetCur() string`

GetCur returns the Cur field if non-nil, zero value otherwise.

### GetCurOk

`func (o *TransactionsTransactionsInner) GetCurOk() (*string, bool)`

GetCurOk returns a tuple with the Cur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCur

`func (o *TransactionsTransactionsInner) SetCur(v string)`

SetCur sets Cur field to given value.

### HasCur

`func (o *TransactionsTransactionsInner) HasCur() bool`

HasCur returns a boolean if a field has been set.

### GetFxRate

`func (o *TransactionsTransactionsInner) GetFxRate() float32`

GetFxRate returns the FxRate field if non-nil, zero value otherwise.

### GetFxRateOk

`func (o *TransactionsTransactionsInner) GetFxRateOk() (*float32, bool)`

GetFxRateOk returns a tuple with the FxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFxRate

`func (o *TransactionsTransactionsInner) SetFxRate(v float32)`

SetFxRate sets FxRate field to given value.

### HasFxRate

`func (o *TransactionsTransactionsInner) HasFxRate() bool`

HasFxRate returns a boolean if a field has been set.

### GetDesc

`func (o *TransactionsTransactionsInner) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *TransactionsTransactionsInner) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *TransactionsTransactionsInner) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *TransactionsTransactionsInner) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetDate

`func (o *TransactionsTransactionsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionsTransactionsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionsTransactionsInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TransactionsTransactionsInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetType

`func (o *TransactionsTransactionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionsTransactionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionsTransactionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionsTransactionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQty

`func (o *TransactionsTransactionsInner) GetQty() float32`

GetQty returns the Qty field if non-nil, zero value otherwise.

### GetQtyOk

`func (o *TransactionsTransactionsInner) GetQtyOk() (*float32, bool)`

GetQtyOk returns a tuple with the Qty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQty

`func (o *TransactionsTransactionsInner) SetQty(v float32)`

SetQty sets Qty field to given value.

### HasQty

`func (o *TransactionsTransactionsInner) HasQty() bool`

HasQty returns a boolean if a field has been set.

### GetPr

`func (o *TransactionsTransactionsInner) GetPr() float32`

GetPr returns the Pr field if non-nil, zero value otherwise.

### GetPrOk

`func (o *TransactionsTransactionsInner) GetPrOk() (*float32, bool)`

GetPrOk returns a tuple with the Pr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPr

`func (o *TransactionsTransactionsInner) SetPr(v float32)`

SetPr sets Pr field to given value.

### HasPr

`func (o *TransactionsTransactionsInner) HasPr() bool`

HasPr returns a boolean if a field has been set.

### GetAmt

`func (o *TransactionsTransactionsInner) GetAmt() float32`

GetAmt returns the Amt field if non-nil, zero value otherwise.

### GetAmtOk

`func (o *TransactionsTransactionsInner) GetAmtOk() (*float32, bool)`

GetAmtOk returns a tuple with the Amt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmt

`func (o *TransactionsTransactionsInner) SetAmt(v float32)`

SetAmt sets Amt field to given value.

### HasAmt

`func (o *TransactionsTransactionsInner) HasAmt() bool`

HasAmt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


