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

// V1CbPolicy struct for V1CbPolicy
type V1CbPolicy struct {
	Consecutive *V1CbPolicyConsecutiveErrConfig `json:"consecutive,omitempty"`
	ErrorRate *V1CbPolicyErrRateConfig `json:"errorRate,omitempty"`
	JudgeDuration *DurationpbDuration `json:"judgeDuration,omitempty"`
	MaxEjectionPercent *int32 `json:"maxEjectionPercent,omitempty"`
	SlowRate *V1CbPolicySlowRateConfig `json:"slowRate,omitempty"`
}

// NewV1CbPolicy instantiates a new V1CbPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CbPolicy() *V1CbPolicy {
	this := V1CbPolicy{}
	return &this
}

// NewV1CbPolicyWithDefaults instantiates a new V1CbPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CbPolicyWithDefaults() *V1CbPolicy {
	this := V1CbPolicy{}
	return &this
}

// GetConsecutive returns the Consecutive field value if set, zero value otherwise.
func (o *V1CbPolicy) GetConsecutive() V1CbPolicyConsecutiveErrConfig {
	if o == nil || o.Consecutive == nil {
		var ret V1CbPolicyConsecutiveErrConfig
		return ret
	}
	return *o.Consecutive
}

// GetConsecutiveOk returns a tuple with the Consecutive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicy) GetConsecutiveOk() (*V1CbPolicyConsecutiveErrConfig, bool) {
	if o == nil || o.Consecutive == nil {
		return nil, false
	}
	return o.Consecutive, true
}

// HasConsecutive returns a boolean if a field has been set.
func (o *V1CbPolicy) HasConsecutive() bool {
	if o != nil && o.Consecutive != nil {
		return true
	}

	return false
}

// SetConsecutive gets a reference to the given V1CbPolicyConsecutiveErrConfig and assigns it to the Consecutive field.
func (o *V1CbPolicy) SetConsecutive(v V1CbPolicyConsecutiveErrConfig) {
	o.Consecutive = &v
}

// GetErrorRate returns the ErrorRate field value if set, zero value otherwise.
func (o *V1CbPolicy) GetErrorRate() V1CbPolicyErrRateConfig {
	if o == nil || o.ErrorRate == nil {
		var ret V1CbPolicyErrRateConfig
		return ret
	}
	return *o.ErrorRate
}

// GetErrorRateOk returns a tuple with the ErrorRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicy) GetErrorRateOk() (*V1CbPolicyErrRateConfig, bool) {
	if o == nil || o.ErrorRate == nil {
		return nil, false
	}
	return o.ErrorRate, true
}

// HasErrorRate returns a boolean if a field has been set.
func (o *V1CbPolicy) HasErrorRate() bool {
	if o != nil && o.ErrorRate != nil {
		return true
	}

	return false
}

// SetErrorRate gets a reference to the given V1CbPolicyErrRateConfig and assigns it to the ErrorRate field.
func (o *V1CbPolicy) SetErrorRate(v V1CbPolicyErrRateConfig) {
	o.ErrorRate = &v
}

// GetJudgeDuration returns the JudgeDuration field value if set, zero value otherwise.
func (o *V1CbPolicy) GetJudgeDuration() DurationpbDuration {
	if o == nil || o.JudgeDuration == nil {
		var ret DurationpbDuration
		return ret
	}
	return *o.JudgeDuration
}

// GetJudgeDurationOk returns a tuple with the JudgeDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicy) GetJudgeDurationOk() (*DurationpbDuration, bool) {
	if o == nil || o.JudgeDuration == nil {
		return nil, false
	}
	return o.JudgeDuration, true
}

// HasJudgeDuration returns a boolean if a field has been set.
func (o *V1CbPolicy) HasJudgeDuration() bool {
	if o != nil && o.JudgeDuration != nil {
		return true
	}

	return false
}

// SetJudgeDuration gets a reference to the given DurationpbDuration and assigns it to the JudgeDuration field.
func (o *V1CbPolicy) SetJudgeDuration(v DurationpbDuration) {
	o.JudgeDuration = &v
}

// GetMaxEjectionPercent returns the MaxEjectionPercent field value if set, zero value otherwise.
func (o *V1CbPolicy) GetMaxEjectionPercent() int32 {
	if o == nil || o.MaxEjectionPercent == nil {
		var ret int32
		return ret
	}
	return *o.MaxEjectionPercent
}

// GetMaxEjectionPercentOk returns a tuple with the MaxEjectionPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicy) GetMaxEjectionPercentOk() (*int32, bool) {
	if o == nil || o.MaxEjectionPercent == nil {
		return nil, false
	}
	return o.MaxEjectionPercent, true
}

// HasMaxEjectionPercent returns a boolean if a field has been set.
func (o *V1CbPolicy) HasMaxEjectionPercent() bool {
	if o != nil && o.MaxEjectionPercent != nil {
		return true
	}

	return false
}

// SetMaxEjectionPercent gets a reference to the given int32 and assigns it to the MaxEjectionPercent field.
func (o *V1CbPolicy) SetMaxEjectionPercent(v int32) {
	o.MaxEjectionPercent = &v
}

// GetSlowRate returns the SlowRate field value if set, zero value otherwise.
func (o *V1CbPolicy) GetSlowRate() V1CbPolicySlowRateConfig {
	if o == nil || o.SlowRate == nil {
		var ret V1CbPolicySlowRateConfig
		return ret
	}
	return *o.SlowRate
}

// GetSlowRateOk returns a tuple with the SlowRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicy) GetSlowRateOk() (*V1CbPolicySlowRateConfig, bool) {
	if o == nil || o.SlowRate == nil {
		return nil, false
	}
	return o.SlowRate, true
}

// HasSlowRate returns a boolean if a field has been set.
func (o *V1CbPolicy) HasSlowRate() bool {
	if o != nil && o.SlowRate != nil {
		return true
	}

	return false
}

// SetSlowRate gets a reference to the given V1CbPolicySlowRateConfig and assigns it to the SlowRate field.
func (o *V1CbPolicy) SetSlowRate(v V1CbPolicySlowRateConfig) {
	o.SlowRate = &v
}

func (o V1CbPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Consecutive != nil {
		toSerialize["consecutive"] = o.Consecutive
	}
	if o.ErrorRate != nil {
		toSerialize["errorRate"] = o.ErrorRate
	}
	if o.JudgeDuration != nil {
		toSerialize["judgeDuration"] = o.JudgeDuration
	}
	if o.MaxEjectionPercent != nil {
		toSerialize["maxEjectionPercent"] = o.MaxEjectionPercent
	}
	if o.SlowRate != nil {
		toSerialize["slowRate"] = o.SlowRate
	}
	return json.Marshal(toSerialize)
}

type NullableV1CbPolicy struct {
	value *V1CbPolicy
	isSet bool
}

func (v NullableV1CbPolicy) Get() *V1CbPolicy {
	return v.value
}

func (v *NullableV1CbPolicy) Set(val *V1CbPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CbPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CbPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CbPolicy(val *V1CbPolicy) *NullableV1CbPolicy {
	return &NullableV1CbPolicy{value: val, isSet: true}
}

func (v NullableV1CbPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CbPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

