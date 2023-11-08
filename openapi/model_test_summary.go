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

// checks if the TestSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestSummary{}

// TestSummary struct for TestSummary
type TestSummary struct {
	Runs    *int32              `json:"runs,omitempty"`
	LastRun *TestSummaryLastRun `json:"lastRun,omitempty"`
}

// NewTestSummary instantiates a new TestSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSummary() *TestSummary {
	this := TestSummary{}
	return &this
}

// NewTestSummaryWithDefaults instantiates a new TestSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSummaryWithDefaults() *TestSummary {
	this := TestSummary{}
	return &this
}

// GetRuns returns the Runs field value if set, zero value otherwise.
func (o *TestSummary) GetRuns() int32 {
	if o == nil || isNil(o.Runs) {
		var ret int32
		return ret
	}
	return *o.Runs
}

// GetRunsOk returns a tuple with the Runs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSummary) GetRunsOk() (*int32, bool) {
	if o == nil || isNil(o.Runs) {
		return nil, false
	}
	return o.Runs, true
}

// HasRuns returns a boolean if a field has been set.
func (o *TestSummary) HasRuns() bool {
	if o != nil && !isNil(o.Runs) {
		return true
	}

	return false
}

// SetRuns gets a reference to the given int32 and assigns it to the Runs field.
func (o *TestSummary) SetRuns(v int32) {
	o.Runs = &v
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *TestSummary) GetLastRun() TestSummaryLastRun {
	if o == nil || isNil(o.LastRun) {
		var ret TestSummaryLastRun
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSummary) GetLastRunOk() (*TestSummaryLastRun, bool) {
	if o == nil || isNil(o.LastRun) {
		return nil, false
	}
	return o.LastRun, true
}

// HasLastRun returns a boolean if a field has been set.
func (o *TestSummary) HasLastRun() bool {
	if o != nil && !isNil(o.LastRun) {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given TestSummaryLastRun and assigns it to the LastRun field.
func (o *TestSummary) SetLastRun(v TestSummaryLastRun) {
	o.LastRun = &v
}

func (o TestSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: runs is readOnly
	if !isNil(o.LastRun) {
		toSerialize["lastRun"] = o.LastRun
	}
	return toSerialize, nil
}

type NullableTestSummary struct {
	value *TestSummary
	isSet bool
}

func (v NullableTestSummary) Get() *TestSummary {
	return v.value
}

func (v *NullableTestSummary) Set(val *TestSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSummary(val *TestSummary) *NullableTestSummary {
	return &NullableTestSummary{value: val, isSet: true}
}

func (v NullableTestSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
