/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MissingVariables struct for MissingVariables
type MissingVariables struct {
	Key          *string `json:"key,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
}

// NewMissingVariables instantiates a new MissingVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissingVariables() *MissingVariables {
	this := MissingVariables{}
	return &this
}

// NewMissingVariablesWithDefaults instantiates a new MissingVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissingVariablesWithDefaults() *MissingVariables {
	this := MissingVariables{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MissingVariables) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissingVariables) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MissingVariables) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *MissingVariables) SetKey(v string) {
	o.Key = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *MissingVariables) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MissingVariables) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *MissingVariables) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *MissingVariables) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

func (o MissingVariables) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	return json.Marshal(toSerialize)
}

type NullableMissingVariables struct {
	value *MissingVariables
	isSet bool
}

func (v NullableMissingVariables) Get() *MissingVariables {
	return v.value
}

func (v *NullableMissingVariables) Set(val *MissingVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableMissingVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableMissingVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissingVariables(val *MissingVariables) *NullableMissingVariables {
	return &NullableMissingVariables{value: val, isSet: true}
}

func (v NullableMissingVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissingVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
