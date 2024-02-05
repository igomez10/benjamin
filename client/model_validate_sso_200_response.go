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

// checks if the ValidateSSO200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateSSO200Response{}

// ValidateSSO200Response struct for ValidateSSO200Response
type ValidateSSO200Response struct {
	// 1 for Live, 2 for Paper
	LOGIN_TYPE *float32 `json:"LOGIN_TYPE,omitempty"`
	// Username
	USER_NAME *string `json:"USER_NAME,omitempty"`
	// User ID
	USER_ID *float32 `json:"USER_ID,omitempty"`
	// Time in milliseconds until session expires. Caller needs to call the again to re-validate session
	Expire *float32 `json:"expire,omitempty"`
	// true if session was validated; false if not.
	RESULT *bool `json:"RESULT,omitempty"`
	// Time of session validation
	AUTH_TIME *float32 `json:"AUTH_TIME,omitempty"`
}

// NewValidateSSO200Response instantiates a new ValidateSSO200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateSSO200Response() *ValidateSSO200Response {
	this := ValidateSSO200Response{}
	return &this
}

// NewValidateSSO200ResponseWithDefaults instantiates a new ValidateSSO200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateSSO200ResponseWithDefaults() *ValidateSSO200Response {
	this := ValidateSSO200Response{}
	return &this
}

// GetLOGIN_TYPE returns the LOGIN_TYPE field value if set, zero value otherwise.
func (o *ValidateSSO200Response) GetLOGIN_TYPE() float32 {
	if o == nil || IsNil(o.LOGIN_TYPE) {
		var ret float32
		return ret
	}
	return *o.LOGIN_TYPE
}

// GetLOGIN_TYPEOk returns a tuple with the LOGIN_TYPE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateSSO200Response) GetLOGIN_TYPEOk() (*float32, bool) {
	if o == nil || IsNil(o.LOGIN_TYPE) {
		return nil, false
	}
	return o.LOGIN_TYPE, true
}

// HasLOGIN_TYPE returns a boolean if a field has been set.
func (o *ValidateSSO200Response) HasLOGIN_TYPE() bool {
	if o != nil && !IsNil(o.LOGIN_TYPE) {
		return true
	}

	return false
}

// SetLOGIN_TYPE gets a reference to the given float32 and assigns it to the LOGIN_TYPE field.
func (o *ValidateSSO200Response) SetLOGIN_TYPE(v float32) {
	o.LOGIN_TYPE = &v
}

// GetUSER_NAME returns the USER_NAME field value if set, zero value otherwise.
func (o *ValidateSSO200Response) GetUSER_NAME() string {
	if o == nil || IsNil(o.USER_NAME) {
		var ret string
		return ret
	}
	return *o.USER_NAME
}

// GetUSER_NAMEOk returns a tuple with the USER_NAME field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateSSO200Response) GetUSER_NAMEOk() (*string, bool) {
	if o == nil || IsNil(o.USER_NAME) {
		return nil, false
	}
	return o.USER_NAME, true
}

// HasUSER_NAME returns a boolean if a field has been set.
func (o *ValidateSSO200Response) HasUSER_NAME() bool {
	if o != nil && !IsNil(o.USER_NAME) {
		return true
	}

	return false
}

// SetUSER_NAME gets a reference to the given string and assigns it to the USER_NAME field.
func (o *ValidateSSO200Response) SetUSER_NAME(v string) {
	o.USER_NAME = &v
}

// GetUSER_ID returns the USER_ID field value if set, zero value otherwise.
func (o *ValidateSSO200Response) GetUSER_ID() float32 {
	if o == nil || IsNil(o.USER_ID) {
		var ret float32
		return ret
	}
	return *o.USER_ID
}

// GetUSER_IDOk returns a tuple with the USER_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateSSO200Response) GetUSER_IDOk() (*float32, bool) {
	if o == nil || IsNil(o.USER_ID) {
		return nil, false
	}
	return o.USER_ID, true
}

// HasUSER_ID returns a boolean if a field has been set.
func (o *ValidateSSO200Response) HasUSER_ID() bool {
	if o != nil && !IsNil(o.USER_ID) {
		return true
	}

	return false
}

// SetUSER_ID gets a reference to the given float32 and assigns it to the USER_ID field.
func (o *ValidateSSO200Response) SetUSER_ID(v float32) {
	o.USER_ID = &v
}

// GetExpire returns the Expire field value if set, zero value otherwise.
func (o *ValidateSSO200Response) GetExpire() float32 {
	if o == nil || IsNil(o.Expire) {
		var ret float32
		return ret
	}
	return *o.Expire
}

// GetExpireOk returns a tuple with the Expire field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateSSO200Response) GetExpireOk() (*float32, bool) {
	if o == nil || IsNil(o.Expire) {
		return nil, false
	}
	return o.Expire, true
}

// HasExpire returns a boolean if a field has been set.
func (o *ValidateSSO200Response) HasExpire() bool {
	if o != nil && !IsNil(o.Expire) {
		return true
	}

	return false
}

// SetExpire gets a reference to the given float32 and assigns it to the Expire field.
func (o *ValidateSSO200Response) SetExpire(v float32) {
	o.Expire = &v
}

// GetRESULT returns the RESULT field value if set, zero value otherwise.
func (o *ValidateSSO200Response) GetRESULT() bool {
	if o == nil || IsNil(o.RESULT) {
		var ret bool
		return ret
	}
	return *o.RESULT
}

// GetRESULTOk returns a tuple with the RESULT field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateSSO200Response) GetRESULTOk() (*bool, bool) {
	if o == nil || IsNil(o.RESULT) {
		return nil, false
	}
	return o.RESULT, true
}

// HasRESULT returns a boolean if a field has been set.
func (o *ValidateSSO200Response) HasRESULT() bool {
	if o != nil && !IsNil(o.RESULT) {
		return true
	}

	return false
}

// SetRESULT gets a reference to the given bool and assigns it to the RESULT field.
func (o *ValidateSSO200Response) SetRESULT(v bool) {
	o.RESULT = &v
}

// GetAUTH_TIME returns the AUTH_TIME field value if set, zero value otherwise.
func (o *ValidateSSO200Response) GetAUTH_TIME() float32 {
	if o == nil || IsNil(o.AUTH_TIME) {
		var ret float32
		return ret
	}
	return *o.AUTH_TIME
}

// GetAUTH_TIMEOk returns a tuple with the AUTH_TIME field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateSSO200Response) GetAUTH_TIMEOk() (*float32, bool) {
	if o == nil || IsNil(o.AUTH_TIME) {
		return nil, false
	}
	return o.AUTH_TIME, true
}

// HasAUTH_TIME returns a boolean if a field has been set.
func (o *ValidateSSO200Response) HasAUTH_TIME() bool {
	if o != nil && !IsNil(o.AUTH_TIME) {
		return true
	}

	return false
}

// SetAUTH_TIME gets a reference to the given float32 and assigns it to the AUTH_TIME field.
func (o *ValidateSSO200Response) SetAUTH_TIME(v float32) {
	o.AUTH_TIME = &v
}

func (o ValidateSSO200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateSSO200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LOGIN_TYPE) {
		toSerialize["LOGIN_TYPE"] = o.LOGIN_TYPE
	}
	if !IsNil(o.USER_NAME) {
		toSerialize["USER_NAME"] = o.USER_NAME
	}
	if !IsNil(o.USER_ID) {
		toSerialize["USER_ID"] = o.USER_ID
	}
	if !IsNil(o.Expire) {
		toSerialize["expire"] = o.Expire
	}
	if !IsNil(o.RESULT) {
		toSerialize["RESULT"] = o.RESULT
	}
	if !IsNil(o.AUTH_TIME) {
		toSerialize["AUTH_TIME"] = o.AUTH_TIME
	}
	return toSerialize, nil
}

type NullableValidateSSO200Response struct {
	value *ValidateSSO200Response
	isSet bool
}

func (v NullableValidateSSO200Response) Get() *ValidateSSO200Response {
	return v.value
}

func (v *NullableValidateSSO200Response) Set(val *ValidateSSO200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateSSO200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateSSO200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateSSO200Response(val *ValidateSSO200Response) *NullableValidateSSO200Response {
	return &NullableValidateSSO200Response{value: val, isSet: true}
}

func (v NullableValidateSSO200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateSSO200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}