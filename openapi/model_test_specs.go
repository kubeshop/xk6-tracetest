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

// checks if the TestSpecs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSpecs{}

// TestSpecs struct for TestSpecs
type TestSpecs struct {
	Specs []TestSpec `json:"specs,omitempty"`
}

// NewTestSpecs instantiates a new TestSpecs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSpecs() *TestSpecs {
	this := TestSpecs{}
	return &this
}

// NewTestSpecsWithDefaults instantiates a new TestSpecs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSpecsWithDefaults() *TestSpecs {
	this := TestSpecs{}
	return &this
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *TestSpecs) GetSpecs() []TestSpec {
	if o == nil || isNil(o.Specs) {
		var ret []TestSpec
		return ret
	}
	return o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSpecs) GetSpecsOk() ([]TestSpec, bool) {
	if o == nil || isNil(o.Specs) {
		return nil, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *TestSpecs) HasSpecs() bool {
	if o != nil && !isNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given []TestSpec and assigns it to the Specs field.
func (o *TestSpecs) SetSpecs(v []TestSpec) {
	o.Specs = v
}

func (o TestSpecs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSpecs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Specs) {
		toSerialize["specs"] = o.Specs
	}
	return toSerialize, nil
}

type NullableTestSpecs struct {
	value *TestSpecs
	isSet bool
}

func (v NullableTestSpecs) Get() *TestSpecs {
	return v.value
}

func (v *NullableTestSpecs) Set(val *TestSpecs) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSpecs) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSpecs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSpecs(val *TestSpecs) *NullableTestSpecs {
	return &NullableTestSpecs{value: val, isSet: true}
}

func (v NullableTestSpecs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSpecs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
