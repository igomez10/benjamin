/*
Client Portal Web API

Client Poral Web API

API version: 1.0.0
Contact: e@e.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the Trade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Trade{}

// Trade struct for Trade
type Trade struct {
	// execution identifier for the order
	ExecutionId *string `json:"execution_id,omitempty"`
	// Underlying Symbol
	Symbol *string `json:"symbol,omitempty"`
	// The side of the market of the order.   * B - Buy contract near posted ask price   * S - Sell contract near posted bid price   * X - Option expired
	Side *string `json:"side,omitempty"`
	// Formatted description of the order \"%side% %size% @ %price% on %exchange%\".
	OrderDescription *string `json:"order_description,omitempty"`
	// Time of Status update in format \"YYYYMMDD-hh:mm:ss\".
	TradeTime *string `json:"trade_time,omitempty"`
	// Time of status update in format unix time.
	TradeTimeR *float32 `json:"trade_time_r,omitempty"`
	// Quantity of the order
	Size *string `json:"size,omitempty"`
	// Average Price
	Price *string `json:"price,omitempty"`
	// User defined string used to identify the order. Value is set using \"cOID\" field while placing an order.
	OrderRef *string `json:"order_ref,omitempty"`
	// User that submitted order
	Submitter *string `json:"submitter,omitempty"`
	// Exchange or venue of order
	Exchange *string `json:"exchange,omitempty"`
	// Commission of the order
	Commission *float32 `json:"commission,omitempty"`
	// Net cost of the order, including contract multiplier and quantity.
	NetAmount *float32 `json:"net_amount,omitempty"`
	// accountCode
	Account *string `json:"account,omitempty"`
	// Account Number
	AcountCode *string `json:"acountCode,omitempty"`
	// Contracts company name
	CompanyName *string `json:"company_name,omitempty"`
	// Format contract name
	ContractDescription1 *string `json:"contract_description_1,omitempty"`
	// Asset class
	SecType *string `json:"sec_type,omitempty"`
	// IBKR's contract identifier
	Conid *string `json:"conid,omitempty"`
	// conid and exchange. Format supports conid or conid@exchange
	Conidex *string `json:"conidex,omitempty"`
	// Total quantity owned for this contract
	Position *string `json:"position,omitempty"`
	// Firm which will settle the trade. For IBExecution customers only.
	ClearingId *string `json:"clearing_id,omitempty"`
	// Specifies the true beneficiary of the order. For IBExecution customers only.
	ClearingName *string `json:"clearing_name,omitempty"`
	// If order adds liquidity to the market.
	LiquidationTrade *float32 `json:"liquidation_trade,omitempty"`
}

// NewTrade instantiates a new Trade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrade() *Trade {
	this := Trade{}
	return &this
}

// NewTradeWithDefaults instantiates a new Trade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeWithDefaults() *Trade {
	this := Trade{}
	return &this
}

// GetExecutionId returns the ExecutionId field value if set, zero value otherwise.
func (o *Trade) GetExecutionId() string {
	if o == nil || IsNil(o.ExecutionId) {
		var ret string
		return ret
	}
	return *o.ExecutionId
}

// GetExecutionIdOk returns a tuple with the ExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetExecutionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionId) {
		return nil, false
	}
	return o.ExecutionId, true
}

// HasExecutionId returns a boolean if a field has been set.
func (o *Trade) HasExecutionId() bool {
	if o != nil && !IsNil(o.ExecutionId) {
		return true
	}

	return false
}

// SetExecutionId gets a reference to the given string and assigns it to the ExecutionId field.
func (o *Trade) SetExecutionId(v string) {
	o.ExecutionId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *Trade) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *Trade) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *Trade) SetSymbol(v string) {
	o.Symbol = &v
}

// GetSide returns the Side field value if set, zero value otherwise.
func (o *Trade) GetSide() string {
	if o == nil || IsNil(o.Side) {
		var ret string
		return ret
	}
	return *o.Side
}

// GetSideOk returns a tuple with the Side field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetSideOk() (*string, bool) {
	if o == nil || IsNil(o.Side) {
		return nil, false
	}
	return o.Side, true
}

// HasSide returns a boolean if a field has been set.
func (o *Trade) HasSide() bool {
	if o != nil && !IsNil(o.Side) {
		return true
	}

	return false
}

// SetSide gets a reference to the given string and assigns it to the Side field.
func (o *Trade) SetSide(v string) {
	o.Side = &v
}

// GetOrderDescription returns the OrderDescription field value if set, zero value otherwise.
func (o *Trade) GetOrderDescription() string {
	if o == nil || IsNil(o.OrderDescription) {
		var ret string
		return ret
	}
	return *o.OrderDescription
}

// GetOrderDescriptionOk returns a tuple with the OrderDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetOrderDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.OrderDescription) {
		return nil, false
	}
	return o.OrderDescription, true
}

// HasOrderDescription returns a boolean if a field has been set.
func (o *Trade) HasOrderDescription() bool {
	if o != nil && !IsNil(o.OrderDescription) {
		return true
	}

	return false
}

// SetOrderDescription gets a reference to the given string and assigns it to the OrderDescription field.
func (o *Trade) SetOrderDescription(v string) {
	o.OrderDescription = &v
}

// GetTradeTime returns the TradeTime field value if set, zero value otherwise.
func (o *Trade) GetTradeTime() string {
	if o == nil || IsNil(o.TradeTime) {
		var ret string
		return ret
	}
	return *o.TradeTime
}

// GetTradeTimeOk returns a tuple with the TradeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetTradeTimeOk() (*string, bool) {
	if o == nil || IsNil(o.TradeTime) {
		return nil, false
	}
	return o.TradeTime, true
}

// HasTradeTime returns a boolean if a field has been set.
func (o *Trade) HasTradeTime() bool {
	if o != nil && !IsNil(o.TradeTime) {
		return true
	}

	return false
}

// SetTradeTime gets a reference to the given string and assigns it to the TradeTime field.
func (o *Trade) SetTradeTime(v string) {
	o.TradeTime = &v
}

// GetTradeTimeR returns the TradeTimeR field value if set, zero value otherwise.
func (o *Trade) GetTradeTimeR() float32 {
	if o == nil || IsNil(o.TradeTimeR) {
		var ret float32
		return ret
	}
	return *o.TradeTimeR
}

// GetTradeTimeROk returns a tuple with the TradeTimeR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetTradeTimeROk() (*float32, bool) {
	if o == nil || IsNil(o.TradeTimeR) {
		return nil, false
	}
	return o.TradeTimeR, true
}

// HasTradeTimeR returns a boolean if a field has been set.
func (o *Trade) HasTradeTimeR() bool {
	if o != nil && !IsNil(o.TradeTimeR) {
		return true
	}

	return false
}

// SetTradeTimeR gets a reference to the given float32 and assigns it to the TradeTimeR field.
func (o *Trade) SetTradeTimeR(v float32) {
	o.TradeTimeR = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *Trade) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *Trade) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *Trade) SetSize(v string) {
	o.Size = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Trade) GetPrice() string {
	if o == nil || IsNil(o.Price) {
		var ret string
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetPriceOk() (*string, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *Trade) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given string and assigns it to the Price field.
func (o *Trade) SetPrice(v string) {
	o.Price = &v
}

// GetOrderRef returns the OrderRef field value if set, zero value otherwise.
func (o *Trade) GetOrderRef() string {
	if o == nil || IsNil(o.OrderRef) {
		var ret string
		return ret
	}
	return *o.OrderRef
}

// GetOrderRefOk returns a tuple with the OrderRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetOrderRefOk() (*string, bool) {
	if o == nil || IsNil(o.OrderRef) {
		return nil, false
	}
	return o.OrderRef, true
}

// HasOrderRef returns a boolean if a field has been set.
func (o *Trade) HasOrderRef() bool {
	if o != nil && !IsNil(o.OrderRef) {
		return true
	}

	return false
}

// SetOrderRef gets a reference to the given string and assigns it to the OrderRef field.
func (o *Trade) SetOrderRef(v string) {
	o.OrderRef = &v
}

// GetSubmitter returns the Submitter field value if set, zero value otherwise.
func (o *Trade) GetSubmitter() string {
	if o == nil || IsNil(o.Submitter) {
		var ret string
		return ret
	}
	return *o.Submitter
}

// GetSubmitterOk returns a tuple with the Submitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetSubmitterOk() (*string, bool) {
	if o == nil || IsNil(o.Submitter) {
		return nil, false
	}
	return o.Submitter, true
}

// HasSubmitter returns a boolean if a field has been set.
func (o *Trade) HasSubmitter() bool {
	if o != nil && !IsNil(o.Submitter) {
		return true
	}

	return false
}

// SetSubmitter gets a reference to the given string and assigns it to the Submitter field.
func (o *Trade) SetSubmitter(v string) {
	o.Submitter = &v
}

// GetExchange returns the Exchange field value if set, zero value otherwise.
func (o *Trade) GetExchange() string {
	if o == nil || IsNil(o.Exchange) {
		var ret string
		return ret
	}
	return *o.Exchange
}

// GetExchangeOk returns a tuple with the Exchange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetExchangeOk() (*string, bool) {
	if o == nil || IsNil(o.Exchange) {
		return nil, false
	}
	return o.Exchange, true
}

// HasExchange returns a boolean if a field has been set.
func (o *Trade) HasExchange() bool {
	if o != nil && !IsNil(o.Exchange) {
		return true
	}

	return false
}

// SetExchange gets a reference to the given string and assigns it to the Exchange field.
func (o *Trade) SetExchange(v string) {
	o.Exchange = &v
}

// GetCommission returns the Commission field value if set, zero value otherwise.
func (o *Trade) GetCommission() float32 {
	if o == nil || IsNil(o.Commission) {
		var ret float32
		return ret
	}
	return *o.Commission
}

// GetCommissionOk returns a tuple with the Commission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetCommissionOk() (*float32, bool) {
	if o == nil || IsNil(o.Commission) {
		return nil, false
	}
	return o.Commission, true
}

// HasCommission returns a boolean if a field has been set.
func (o *Trade) HasCommission() bool {
	if o != nil && !IsNil(o.Commission) {
		return true
	}

	return false
}

// SetCommission gets a reference to the given float32 and assigns it to the Commission field.
func (o *Trade) SetCommission(v float32) {
	o.Commission = &v
}

// GetNetAmount returns the NetAmount field value if set, zero value otherwise.
func (o *Trade) GetNetAmount() float32 {
	if o == nil || IsNil(o.NetAmount) {
		var ret float32
		return ret
	}
	return *o.NetAmount
}

// GetNetAmountOk returns a tuple with the NetAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetNetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.NetAmount) {
		return nil, false
	}
	return o.NetAmount, true
}

// HasNetAmount returns a boolean if a field has been set.
func (o *Trade) HasNetAmount() bool {
	if o != nil && !IsNil(o.NetAmount) {
		return true
	}

	return false
}

// SetNetAmount gets a reference to the given float32 and assigns it to the NetAmount field.
func (o *Trade) SetNetAmount(v float32) {
	o.NetAmount = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Trade) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Trade) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Trade) SetAccount(v string) {
	o.Account = &v
}

// GetAcountCode returns the AcountCode field value if set, zero value otherwise.
func (o *Trade) GetAcountCode() string {
	if o == nil || IsNil(o.AcountCode) {
		var ret string
		return ret
	}
	return *o.AcountCode
}

// GetAcountCodeOk returns a tuple with the AcountCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetAcountCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AcountCode) {
		return nil, false
	}
	return o.AcountCode, true
}

// HasAcountCode returns a boolean if a field has been set.
func (o *Trade) HasAcountCode() bool {
	if o != nil && !IsNil(o.AcountCode) {
		return true
	}

	return false
}

// SetAcountCode gets a reference to the given string and assigns it to the AcountCode field.
func (o *Trade) SetAcountCode(v string) {
	o.AcountCode = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *Trade) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *Trade) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *Trade) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetContractDescription1 returns the ContractDescription1 field value if set, zero value otherwise.
func (o *Trade) GetContractDescription1() string {
	if o == nil || IsNil(o.ContractDescription1) {
		var ret string
		return ret
	}
	return *o.ContractDescription1
}

// GetContractDescription1Ok returns a tuple with the ContractDescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetContractDescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.ContractDescription1) {
		return nil, false
	}
	return o.ContractDescription1, true
}

// HasContractDescription1 returns a boolean if a field has been set.
func (o *Trade) HasContractDescription1() bool {
	if o != nil && !IsNil(o.ContractDescription1) {
		return true
	}

	return false
}

// SetContractDescription1 gets a reference to the given string and assigns it to the ContractDescription1 field.
func (o *Trade) SetContractDescription1(v string) {
	o.ContractDescription1 = &v
}

// GetSecType returns the SecType field value if set, zero value otherwise.
func (o *Trade) GetSecType() string {
	if o == nil || IsNil(o.SecType) {
		var ret string
		return ret
	}
	return *o.SecType
}

// GetSecTypeOk returns a tuple with the SecType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetSecTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SecType) {
		return nil, false
	}
	return o.SecType, true
}

// HasSecType returns a boolean if a field has been set.
func (o *Trade) HasSecType() bool {
	if o != nil && !IsNil(o.SecType) {
		return true
	}

	return false
}

// SetSecType gets a reference to the given string and assigns it to the SecType field.
func (o *Trade) SetSecType(v string) {
	o.SecType = &v
}

// GetConid returns the Conid field value if set, zero value otherwise.
func (o *Trade) GetConid() string {
	if o == nil || IsNil(o.Conid) {
		var ret string
		return ret
	}
	return *o.Conid
}

// GetConidOk returns a tuple with the Conid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetConidOk() (*string, bool) {
	if o == nil || IsNil(o.Conid) {
		return nil, false
	}
	return o.Conid, true
}

// HasConid returns a boolean if a field has been set.
func (o *Trade) HasConid() bool {
	if o != nil && !IsNil(o.Conid) {
		return true
	}

	return false
}

// SetConid gets a reference to the given string and assigns it to the Conid field.
func (o *Trade) SetConid(v string) {
	o.Conid = &v
}

// GetConidex returns the Conidex field value if set, zero value otherwise.
func (o *Trade) GetConidex() string {
	if o == nil || IsNil(o.Conidex) {
		var ret string
		return ret
	}
	return *o.Conidex
}

// GetConidexOk returns a tuple with the Conidex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetConidexOk() (*string, bool) {
	if o == nil || IsNil(o.Conidex) {
		return nil, false
	}
	return o.Conidex, true
}

// HasConidex returns a boolean if a field has been set.
func (o *Trade) HasConidex() bool {
	if o != nil && !IsNil(o.Conidex) {
		return true
	}

	return false
}

// SetConidex gets a reference to the given string and assigns it to the Conidex field.
func (o *Trade) SetConidex(v string) {
	o.Conidex = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *Trade) GetPosition() string {
	if o == nil || IsNil(o.Position) {
		var ret string
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetPositionOk() (*string, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *Trade) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given string and assigns it to the Position field.
func (o *Trade) SetPosition(v string) {
	o.Position = &v
}

// GetClearingId returns the ClearingId field value if set, zero value otherwise.
func (o *Trade) GetClearingId() string {
	if o == nil || IsNil(o.ClearingId) {
		var ret string
		return ret
	}
	return *o.ClearingId
}

// GetClearingIdOk returns a tuple with the ClearingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetClearingIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClearingId) {
		return nil, false
	}
	return o.ClearingId, true
}

// HasClearingId returns a boolean if a field has been set.
func (o *Trade) HasClearingId() bool {
	if o != nil && !IsNil(o.ClearingId) {
		return true
	}

	return false
}

// SetClearingId gets a reference to the given string and assigns it to the ClearingId field.
func (o *Trade) SetClearingId(v string) {
	o.ClearingId = &v
}

// GetClearingName returns the ClearingName field value if set, zero value otherwise.
func (o *Trade) GetClearingName() string {
	if o == nil || IsNil(o.ClearingName) {
		var ret string
		return ret
	}
	return *o.ClearingName
}

// GetClearingNameOk returns a tuple with the ClearingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetClearingNameOk() (*string, bool) {
	if o == nil || IsNil(o.ClearingName) {
		return nil, false
	}
	return o.ClearingName, true
}

// HasClearingName returns a boolean if a field has been set.
func (o *Trade) HasClearingName() bool {
	if o != nil && !IsNil(o.ClearingName) {
		return true
	}

	return false
}

// SetClearingName gets a reference to the given string and assigns it to the ClearingName field.
func (o *Trade) SetClearingName(v string) {
	o.ClearingName = &v
}

// GetLiquidationTrade returns the LiquidationTrade field value if set, zero value otherwise.
func (o *Trade) GetLiquidationTrade() float32 {
	if o == nil || IsNil(o.LiquidationTrade) {
		var ret float32
		return ret
	}
	return *o.LiquidationTrade
}

// GetLiquidationTradeOk returns a tuple with the LiquidationTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trade) GetLiquidationTradeOk() (*float32, bool) {
	if o == nil || IsNil(o.LiquidationTrade) {
		return nil, false
	}
	return o.LiquidationTrade, true
}

// HasLiquidationTrade returns a boolean if a field has been set.
func (o *Trade) HasLiquidationTrade() bool {
	if o != nil && !IsNil(o.LiquidationTrade) {
		return true
	}

	return false
}

// SetLiquidationTrade gets a reference to the given float32 and assigns it to the LiquidationTrade field.
func (o *Trade) SetLiquidationTrade(v float32) {
	o.LiquidationTrade = &v
}

func (o Trade) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Trade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecutionId) {
		toSerialize["execution_id"] = o.ExecutionId
	}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.Side) {
		toSerialize["side"] = o.Side
	}
	if !IsNil(o.OrderDescription) {
		toSerialize["order_description"] = o.OrderDescription
	}
	if !IsNil(o.TradeTime) {
		toSerialize["trade_time"] = o.TradeTime
	}
	if !IsNil(o.TradeTimeR) {
		toSerialize["trade_time_r"] = o.TradeTimeR
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.OrderRef) {
		toSerialize["order_ref"] = o.OrderRef
	}
	if !IsNil(o.Submitter) {
		toSerialize["submitter"] = o.Submitter
	}
	if !IsNil(o.Exchange) {
		toSerialize["exchange"] = o.Exchange
	}
	if !IsNil(o.Commission) {
		toSerialize["commission"] = o.Commission
	}
	if !IsNil(o.NetAmount) {
		toSerialize["net_amount"] = o.NetAmount
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.AcountCode) {
		toSerialize["acountCode"] = o.AcountCode
	}
	if !IsNil(o.CompanyName) {
		toSerialize["company_name"] = o.CompanyName
	}
	if !IsNil(o.ContractDescription1) {
		toSerialize["contract_description_1"] = o.ContractDescription1
	}
	if !IsNil(o.SecType) {
		toSerialize["sec_type"] = o.SecType
	}
	if !IsNil(o.Conid) {
		toSerialize["conid"] = o.Conid
	}
	if !IsNil(o.Conidex) {
		toSerialize["conidex"] = o.Conidex
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.ClearingId) {
		toSerialize["clearing_id"] = o.ClearingId
	}
	if !IsNil(o.ClearingName) {
		toSerialize["clearing_name"] = o.ClearingName
	}
	if !IsNil(o.LiquidationTrade) {
		toSerialize["liquidation_trade"] = o.LiquidationTrade
	}
	return toSerialize, nil
}

type NullableTrade struct {
	value *Trade
	isSet bool
}

func (v NullableTrade) Get() *Trade {
	return v.value
}

func (v *NullableTrade) Set(val *Trade) {
	v.value = val
	v.isSet = true
}

func (v NullableTrade) IsSet() bool {
	return v.isSet
}

func (v *NullableTrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrade(val *Trade) *NullableTrade {
	return &NullableTrade{value: val, isSet: true}
}

func (v NullableTrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}