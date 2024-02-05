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

// checks if the AlertResponseConditionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertResponseConditionsInner{}

// AlertResponseConditionsInner struct for AlertResponseConditionsInner
type AlertResponseConditionsInner struct {
	// Types: 1-Price, 3-Time, 4-Margin, 5-Trade, 6-Volume, 7: MTA market 8: MTA Position, 9: MTA Acc. Daily PN&
	ConditionType *int32 `json:"condition_type,omitempty"`
	// conid and exchange. Format supports conid or conid@exchange
	Conidex *string `json:"conidex,omitempty"`
	// Format contract name
	ContractDescription1 *string `json:"contract_description_1,omitempty"`
	// optional, operator for the current condition   * >= Greater than or equal to   * <= Less than or equal to
	ConditionOperator *string `json:"condition_operator,omitempty"`
	// optional, only some type of conditions have triggerMethod
	ConditionTriggerMethod *string `json:"condition_trigger_method,omitempty"`
	// can not be empty, can pass default value \"*\"
	ConditionValue *string `json:"condition_value,omitempty"`
	// Condition array should end with \"n\"   * a - AND   * o - OR   * n - END
	ConditionLogicBind *string `json:"condition_logic_bind,omitempty"`
	// only needed for some MTA alert condition
	ConditionTimeZone *string `json:"condition_time_zone,omitempty"`
}

// NewAlertResponseConditionsInner instantiates a new AlertResponseConditionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertResponseConditionsInner() *AlertResponseConditionsInner {
	this := AlertResponseConditionsInner{}
	return &this
}

// NewAlertResponseConditionsInnerWithDefaults instantiates a new AlertResponseConditionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertResponseConditionsInnerWithDefaults() *AlertResponseConditionsInner {
	this := AlertResponseConditionsInner{}
	return &this
}

// GetConditionType returns the ConditionType field value if set, zero value otherwise.
func (o *AlertResponseConditionsInner) GetConditionType() int32 {
	if o == nil || IsNil(o.ConditionType) {
		var ret int32
		return ret
	}
	return *o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponseConditionsInner) GetConditionTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.ConditionType) {
		return nil, false
	}
	return o.ConditionType, true
}

// HasConditionType returns a boolean if a field has been set.
func (o *AlertResponseConditionsInner) HasConditionType() bool {
	if o != nil && !IsNil(o.ConditionType) {
		return true
	}

	return false
}

// SetConditionType gets a reference to the given int32 and assigns it to the ConditionType field.
func (o *AlertResponseConditionsInner) SetConditionType(v int32) {
	o.ConditionType = &v
}

// GetConidex returns the Conidex field value if set, zero value otherwise.
func (o *AlertResponseConditionsInner) GetConidex() string {
	if o == nil || IsNil(o.Conidex) {
		var ret string
		return ret
	}
	return *o.Conidex
}

// GetConidexOk returns a tuple with the Conidex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponseConditionsInner) GetConidexOk() (*string, bool) {
	if o == nil || IsNil(o.Conidex) {
		return nil, false
	}
	return o.Conidex, true
}

// HasConidex returns a boolean if a field has been set.
func (o *AlertResponseConditionsInner) HasConidex() bool {
	if o != nil && !IsNil(o.Conidex) {
		return true
	}

	return false
}

// SetConidex gets a reference to the given string and assigns it to the Conidex field.
func (o *AlertResponseConditionsInner) SetConidex(v string) {
	o.Conidex = &v
}

// GetContractDescription1 returns the ContractDescription1 field value if set, zero value otherwise.
func (o *AlertResponseConditionsInner) GetContractDescription1() string {
	if o == nil || IsNil(o.ContractDescription1) {
		var ret string
		return ret
	}
	return *o.ContractDescription1
}

// GetContractDescription1Ok returns a tuple with the ContractDescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponseConditionsInner) GetContractDescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.ContractDescription1) {
		return nil, false
	}
	return o.ContractDescription1, true
}

// HasContractDescription1 returns a boolean if a field has been set.
func (o *AlertResponseConditionsInner) HasContractDescription1() bool {
	if o != nil && !IsNil(o.ContractDescription1) {
		return true
	}

	return false
}

// SetContractDescription1 gets a reference to the given string and assigns it to the ContractDescription1 field.
func (o *AlertResponseConditionsInner) SetContractDescription1(v string) {
	o.ContractDescription1 = &v
}

// GetConditionOperator returns the ConditionOperator field value if set, zero value otherwise.
func (o *AlertResponseConditionsInner) GetConditionOperator() string {
	if o == nil || IsNil(o.ConditionOperator) {
		var ret string
		return ret
	}
	return *o.ConditionOperator
}

// GetConditionOperatorOk returns a tuple with the ConditionOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponseConditionsInner) GetConditionOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionOperator) {
		return nil, false
	}
	return o.ConditionOperator, true
}

// HasConditionOperator returns a boolean if a field has been set.
func (o *AlertResponseConditionsInner) HasConditionOperator() bool {
	if o != nil && !IsNil(o.ConditionOperator) {
		return true
	}

	return false
}

// SetConditionOperator gets a reference to the given string and assigns it to the ConditionOperator field.
func (o *AlertResponseConditionsInner) SetConditionOperator(v string) {
	o.ConditionOperator = &v
}

// GetConditionTriggerMethod returns the ConditionTriggerMethod field value if set, zero value otherwise.
func (o *AlertResponseConditionsInner) GetConditionTriggerMethod() string {
	if o == nil || IsNil(o.ConditionTriggerMethod) {
		var ret string
		return ret
	}
	return *o.ConditionTriggerMethod
}

// GetConditionTriggerMethodOk returns a tuple with the ConditionTriggerMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponseConditionsInner) GetConditionTriggerMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionTriggerMethod) {
		return nil, false
	}
	return o.ConditionTriggerMethod, true
}

// HasConditionTriggerMethod returns a boolean if a field has been set.
func (o *AlertResponseConditionsInner) HasConditionTriggerMethod() bool {
	if o != nil && !IsNil(o.ConditionTriggerMethod) {
		return true
	}

	return false
}

// SetConditionTriggerMethod gets a reference to the given string and assigns it to the ConditionTriggerMethod field.
func (o *AlertResponseConditionsInner) SetConditionTriggerMethod(v string) {
	o.ConditionTriggerMethod = &v
}

// GetConditionValue returns the ConditionValue field value if set, zero value otherwise.
func (o *AlertResponseConditionsInner) GetConditionValue() string {
	if o == nil || IsNil(o.ConditionValue) {
		var ret string
		return ret
	}
	return *o.ConditionValue
}

// GetConditionValueOk returns a tuple with the ConditionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponseConditionsInner) GetConditionValueOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionValue) {
		return nil, false
	}
	return o.ConditionValue, true
}

// HasConditionValue returns a boolean if a field has been set.
func (o *AlertResponseConditionsInner) HasConditionValue() bool {
	if o != nil && !IsNil(o.ConditionValue) {
		return true
	}

	return false
}

// SetConditionValue gets a reference to the given string and assigns it to the ConditionValue field.
func (o *AlertResponseConditionsInner) SetConditionValue(v string) {
	o.ConditionValue = &v
}

// GetConditionLogicBind returns the ConditionLogicBind field value if set, zero value otherwise.
func (o *AlertResponseConditionsInner) GetConditionLogicBind() string {
	if o == nil || IsNil(o.ConditionLogicBind) {
		var ret string
		return ret
	}
	return *o.ConditionLogicBind
}

// GetConditionLogicBindOk returns a tuple with the ConditionLogicBind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponseConditionsInner) GetConditionLogicBindOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionLogicBind) {
		return nil, false
	}
	return o.ConditionLogicBind, true
}

// HasConditionLogicBind returns a boolean if a field has been set.
func (o *AlertResponseConditionsInner) HasConditionLogicBind() bool {
	if o != nil && !IsNil(o.ConditionLogicBind) {
		return true
	}

	return false
}

// SetConditionLogicBind gets a reference to the given string and assigns it to the ConditionLogicBind field.
func (o *AlertResponseConditionsInner) SetConditionLogicBind(v string) {
	o.ConditionLogicBind = &v
}

// GetConditionTimeZone returns the ConditionTimeZone field value if set, zero value otherwise.
func (o *AlertResponseConditionsInner) GetConditionTimeZone() string {
	if o == nil || IsNil(o.ConditionTimeZone) {
		var ret string
		return ret
	}
	return *o.ConditionTimeZone
}

// GetConditionTimeZoneOk returns a tuple with the ConditionTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponseConditionsInner) GetConditionTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionTimeZone) {
		return nil, false
	}
	return o.ConditionTimeZone, true
}

// HasConditionTimeZone returns a boolean if a field has been set.
func (o *AlertResponseConditionsInner) HasConditionTimeZone() bool {
	if o != nil && !IsNil(o.ConditionTimeZone) {
		return true
	}

	return false
}

// SetConditionTimeZone gets a reference to the given string and assigns it to the ConditionTimeZone field.
func (o *AlertResponseConditionsInner) SetConditionTimeZone(v string) {
	o.ConditionTimeZone = &v
}

func (o AlertResponseConditionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertResponseConditionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConditionType) {
		toSerialize["condition_type"] = o.ConditionType
	}
	if !IsNil(o.Conidex) {
		toSerialize["conidex"] = o.Conidex
	}
	if !IsNil(o.ContractDescription1) {
		toSerialize["contract_description_1"] = o.ContractDescription1
	}
	if !IsNil(o.ConditionOperator) {
		toSerialize["condition_operator"] = o.ConditionOperator
	}
	if !IsNil(o.ConditionTriggerMethod) {
		toSerialize["condition_trigger_method"] = o.ConditionTriggerMethod
	}
	if !IsNil(o.ConditionValue) {
		toSerialize["condition_value"] = o.ConditionValue
	}
	if !IsNil(o.ConditionLogicBind) {
		toSerialize["condition_logic_bind"] = o.ConditionLogicBind
	}
	if !IsNil(o.ConditionTimeZone) {
		toSerialize["condition_time_zone"] = o.ConditionTimeZone
	}
	return toSerialize, nil
}

type NullableAlertResponseConditionsInner struct {
	value *AlertResponseConditionsInner
	isSet bool
}

func (v NullableAlertResponseConditionsInner) Get() *AlertResponseConditionsInner {
	return v.value
}

func (v *NullableAlertResponseConditionsInner) Set(val *AlertResponseConditionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertResponseConditionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertResponseConditionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertResponseConditionsInner(val *AlertResponseConditionsInner) *NullableAlertResponseConditionsInner {
	return &NullableAlertResponseConditionsInner{value: val, isSet: true}
}

func (v NullableAlertResponseConditionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertResponseConditionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}