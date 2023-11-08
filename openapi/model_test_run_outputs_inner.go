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

// checks if the TestRunOutputsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRunOutputsInner{}

// TestRunOutputsInner struct for TestRunOutputsInner
type TestRunOutputsInner struct {
	Name   *string `json:"name,omitempty"`
	SpanId *string `json:"spanId,omitempty"`
	Value  *string `json:"value,omitempty"`
	Error  *string `json:"error,omitempty"`
}

// NewTestRunOutputsInner instantiates a new TestRunOutputsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRunOutputsInner() *TestRunOutputsInner {
	this := TestRunOutputsInner{}
	return &this
}

// NewTestRunOutputsInnerWithDefaults instantiates a new TestRunOutputsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunOutputsInnerWithDefaults() *TestRunOutputsInner {
	this := TestRunOutputsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TestRunOutputsInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunOutputsInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TestRunOutputsInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TestRunOutputsInner) SetName(v string) {
	o.Name = &v
}

// GetSpanId returns the SpanId field value if set, zero value otherwise.
func (o *TestRunOutputsInner) GetSpanId() string {
	if o == nil || isNil(o.SpanId) {
		var ret string
		return ret
	}
	return *o.SpanId
}

// GetSpanIdOk returns a tuple with the SpanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunOutputsInner) GetSpanIdOk() (*string, bool) {
	if o == nil || isNil(o.SpanId) {
		return nil, false
	}
	return o.SpanId, true
}

// HasSpanId returns a boolean if a field has been set.
func (o *TestRunOutputsInner) HasSpanId() bool {
	if o != nil && !isNil(o.SpanId) {
		return true
	}

	return false
}

// SetSpanId gets a reference to the given string and assigns it to the SpanId field.
func (o *TestRunOutputsInner) SetSpanId(v string) {
	o.SpanId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TestRunOutputsInner) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunOutputsInner) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TestRunOutputsInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TestRunOutputsInner) SetValue(v string) {
	o.Value = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TestRunOutputsInner) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRunOutputsInner) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TestRunOutputsInner) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *TestRunOutputsInner) SetError(v string) {
	o.Error = &v
}

func (o TestRunOutputsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRunOutputsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SpanId) {
		toSerialize["spanId"] = o.SpanId
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableTestRunOutputsInner struct {
	value *TestRunOutputsInner
	isSet bool
}

func (v NullableTestRunOutputsInner) Get() *TestRunOutputsInner {
	return v.value
}

func (v *NullableTestRunOutputsInner) Set(val *TestRunOutputsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRunOutputsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRunOutputsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRunOutputsInner(val *TestRunOutputsInner) *NullableTestRunOutputsInner {
	return &NullableTestRunOutputsInner{value: val, isSet: true}
}

func (v NullableTestRunOutputsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRunOutputsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
