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

// checks if the AuthStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthStatus{}

// AuthStatus struct for AuthStatus
type AuthStatus struct {
	// Brokerage session is authenticated
	Authenticated *bool `json:"authenticated,omitempty"`
	// Connected to backend
	Connected *bool `json:"connected,omitempty"`
	// Brokerage session is competing, e.g. user is logged in to IBKR Mobile, WebTrader, TWS or other trading platforms.
	Competing *bool `json:"competing,omitempty"`
	// Authentication failed, why.
	Fail *string `json:"fail,omitempty"`
	// System messages that may affect trading
	Message *string `json:"message,omitempty"`
	// Prompt messages that may affect trading or the account
	Prompts []string `json:"prompts,omitempty"`
}

// NewAuthStatus instantiates a new AuthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthStatus() *AuthStatus {
	this := AuthStatus{}
	return &this
}

// NewAuthStatusWithDefaults instantiates a new AuthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthStatusWithDefaults() *AuthStatus {
	this := AuthStatus{}
	return &this
}

// GetAuthenticated returns the Authenticated field value if set, zero value otherwise.
func (o *AuthStatus) GetAuthenticated() bool {
	if o == nil || IsNil(o.Authenticated) {
		var ret bool
		return ret
	}
	return *o.Authenticated
}

// GetAuthenticatedOk returns a tuple with the Authenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthStatus) GetAuthenticatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Authenticated) {
		return nil, false
	}
	return o.Authenticated, true
}

// HasAuthenticated returns a boolean if a field has been set.
func (o *AuthStatus) HasAuthenticated() bool {
	if o != nil && !IsNil(o.Authenticated) {
		return true
	}

	return false
}

// SetAuthenticated gets a reference to the given bool and assigns it to the Authenticated field.
func (o *AuthStatus) SetAuthenticated(v bool) {
	o.Authenticated = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *AuthStatus) GetConnected() bool {
	if o == nil || IsNil(o.Connected) {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthStatus) GetConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *AuthStatus) HasConnected() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *AuthStatus) SetConnected(v bool) {
	o.Connected = &v
}

// GetCompeting returns the Competing field value if set, zero value otherwise.
func (o *AuthStatus) GetCompeting() bool {
	if o == nil || IsNil(o.Competing) {
		var ret bool
		return ret
	}
	return *o.Competing
}

// GetCompetingOk returns a tuple with the Competing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthStatus) GetCompetingOk() (*bool, bool) {
	if o == nil || IsNil(o.Competing) {
		return nil, false
	}
	return o.Competing, true
}

// HasCompeting returns a boolean if a field has been set.
func (o *AuthStatus) HasCompeting() bool {
	if o != nil && !IsNil(o.Competing) {
		return true
	}

	return false
}

// SetCompeting gets a reference to the given bool and assigns it to the Competing field.
func (o *AuthStatus) SetCompeting(v bool) {
	o.Competing = &v
}

// GetFail returns the Fail field value if set, zero value otherwise.
func (o *AuthStatus) GetFail() string {
	if o == nil || IsNil(o.Fail) {
		var ret string
		return ret
	}
	return *o.Fail
}

// GetFailOk returns a tuple with the Fail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthStatus) GetFailOk() (*string, bool) {
	if o == nil || IsNil(o.Fail) {
		return nil, false
	}
	return o.Fail, true
}

// HasFail returns a boolean if a field has been set.
func (o *AuthStatus) HasFail() bool {
	if o != nil && !IsNil(o.Fail) {
		return true
	}

	return false
}

// SetFail gets a reference to the given string and assigns it to the Fail field.
func (o *AuthStatus) SetFail(v string) {
	o.Fail = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuthStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuthStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuthStatus) SetMessage(v string) {
	o.Message = &v
}

// GetPrompts returns the Prompts field value if set, zero value otherwise.
func (o *AuthStatus) GetPrompts() []string {
	if o == nil || IsNil(o.Prompts) {
		var ret []string
		return ret
	}
	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthStatus) GetPromptsOk() ([]string, bool) {
	if o == nil || IsNil(o.Prompts) {
		return nil, false
	}
	return o.Prompts, true
}

// HasPrompts returns a boolean if a field has been set.
func (o *AuthStatus) HasPrompts() bool {
	if o != nil && !IsNil(o.Prompts) {
		return true
	}

	return false
}

// SetPrompts gets a reference to the given []string and assigns it to the Prompts field.
func (o *AuthStatus) SetPrompts(v []string) {
	o.Prompts = v
}

func (o AuthStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authenticated) {
		toSerialize["authenticated"] = o.Authenticated
	}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	if !IsNil(o.Competing) {
		toSerialize["competing"] = o.Competing
	}
	if !IsNil(o.Fail) {
		toSerialize["fail"] = o.Fail
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Prompts) {
		toSerialize["prompts"] = o.Prompts
	}
	return toSerialize, nil
}

type NullableAuthStatus struct {
	value *AuthStatus
	isSet bool
}

func (v NullableAuthStatus) Get() *AuthStatus {
	return v.value
}

func (v *NullableAuthStatus) Set(val *AuthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthStatus(val *AuthStatus) *NullableAuthStatus {
	return &NullableAuthStatus{value: val, isSet: true}
}

func (v NullableAuthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
