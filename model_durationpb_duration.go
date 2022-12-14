/*
Polaris Server

一个支持多语言、多框架的云原生服务发现和治理中心  提供高性能SDK和无侵入Sidecar两种接入方式  

API version: v0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package polaris

import (
	"encoding/json"
)

// DurationpbDuration struct for DurationpbDuration
type DurationpbDuration struct {
	Nanos *int32 `json:"nanos,omitempty"`
	Seconds *int64 `json:"seconds,omitempty"`
}

// NewDurationpbDuration instantiates a new DurationpbDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDurationpbDuration() *DurationpbDuration {
	this := DurationpbDuration{}
	return &this
}

// NewDurationpbDurationWithDefaults instantiates a new DurationpbDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDurationpbDurationWithDefaults() *DurationpbDuration {
	this := DurationpbDuration{}
	return &this
}

// GetNanos returns the Nanos field value if set, zero value otherwise.
func (o *DurationpbDuration) GetNanos() int32 {
	if o == nil || o.Nanos == nil {
		var ret int32
		return ret
	}
	return *o.Nanos
}

// GetNanosOk returns a tuple with the Nanos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DurationpbDuration) GetNanosOk() (*int32, bool) {
	if o == nil || o.Nanos == nil {
		return nil, false
	}
	return o.Nanos, true
}

// HasNanos returns a boolean if a field has been set.
func (o *DurationpbDuration) HasNanos() bool {
	if o != nil && o.Nanos != nil {
		return true
	}

	return false
}

// SetNanos gets a reference to the given int32 and assigns it to the Nanos field.
func (o *DurationpbDuration) SetNanos(v int32) {
	o.Nanos = &v
}

// GetSeconds returns the Seconds field value if set, zero value otherwise.
func (o *DurationpbDuration) GetSeconds() int64 {
	if o == nil || o.Seconds == nil {
		var ret int64
		return ret
	}
	return *o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DurationpbDuration) GetSecondsOk() (*int64, bool) {
	if o == nil || o.Seconds == nil {
		return nil, false
	}
	return o.Seconds, true
}

// HasSeconds returns a boolean if a field has been set.
func (o *DurationpbDuration) HasSeconds() bool {
	if o != nil && o.Seconds != nil {
		return true
	}

	return false
}

// SetSeconds gets a reference to the given int64 and assigns it to the Seconds field.
func (o *DurationpbDuration) SetSeconds(v int64) {
	o.Seconds = &v
}

func (o DurationpbDuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Nanos != nil {
		toSerialize["nanos"] = o.Nanos
	}
	if o.Seconds != nil {
		toSerialize["seconds"] = o.Seconds
	}
	return json.Marshal(toSerialize)
}

type NullableDurationpbDuration struct {
	value *DurationpbDuration
	isSet bool
}

func (v NullableDurationpbDuration) Get() *DurationpbDuration {
	return v.value
}

func (v *NullableDurationpbDuration) Set(val *DurationpbDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableDurationpbDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableDurationpbDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDurationpbDuration(val *DurationpbDuration) *NullableDurationpbDuration {
	return &NullableDurationpbDuration{value: val, isSet: true}
}

func (v NullableDurationpbDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDurationpbDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


