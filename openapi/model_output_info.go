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

// checks if the OutputInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputInfo{}

// OutputInfo struct for OutputInfo
type OutputInfo struct {
	LogLevel   *string `json:"logLevel,omitempty"`
	Message    *string `json:"message,omitempty"`
	OutputName *string `json:"outputName,omitempty"`
}

// NewOutputInfo instantiates a new OutputInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputInfo() *OutputInfo {
	this := OutputInfo{}
	return &this
}

// NewOutputInfoWithDefaults instantiates a new OutputInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputInfoWithDefaults() *OutputInfo {
	this := OutputInfo{}
	return &this
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *OutputInfo) GetLogLevel() string {
	if o == nil || isNil(o.LogLevel) {
		var ret string
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfo) GetLogLevelOk() (*string, bool) {
	if o == nil || isNil(o.LogLevel) {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *OutputInfo) HasLogLevel() bool {
	if o != nil && !isNil(o.LogLevel) {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given string and assigns it to the LogLevel field.
func (o *OutputInfo) SetLogLevel(v string) {
	o.LogLevel = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *OutputInfo) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfo) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *OutputInfo) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *OutputInfo) SetMessage(v string) {
	o.Message = &v
}

// GetOutputName returns the OutputName field value if set, zero value otherwise.
func (o *OutputInfo) GetOutputName() string {
	if o == nil || isNil(o.OutputName) {
		var ret string
		return ret
	}
	return *o.OutputName
}

// GetOutputNameOk returns a tuple with the OutputName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutputInfo) GetOutputNameOk() (*string, bool) {
	if o == nil || isNil(o.OutputName) {
		return nil, false
	}
	return o.OutputName, true
}

// HasOutputName returns a boolean if a field has been set.
func (o *OutputInfo) HasOutputName() bool {
	if o != nil && !isNil(o.OutputName) {
		return true
	}

	return false
}

// SetOutputName gets a reference to the given string and assigns it to the OutputName field.
func (o *OutputInfo) SetOutputName(v string) {
	o.OutputName = &v
}

func (o OutputInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LogLevel) {
		toSerialize["logLevel"] = o.LogLevel
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.OutputName) {
		toSerialize["outputName"] = o.OutputName
	}
	return toSerialize, nil
}

type NullableOutputInfo struct {
	value *OutputInfo
	isSet bool
}

func (v NullableOutputInfo) Get() *OutputInfo {
	return v.value
}

func (v *NullableOutputInfo) Set(val *OutputInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputInfo(val *OutputInfo) *NullableOutputInfo {
	return &NullableOutputInfo{value: val, isSet: true}
}

func (v NullableOutputInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
