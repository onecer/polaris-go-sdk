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

// V1ClimbConfigTriggerPolicy struct for V1ClimbConfigTriggerPolicy
type V1ClimbConfigTriggerPolicy struct {
	ErrorRate *V1ClimbConfigTriggerPolicyErrorRate `json:"errorRate,omitempty"`
	SlowRate *V1ClimbConfigTriggerPolicySlowRate `json:"slowRate,omitempty"`
}

// NewV1ClimbConfigTriggerPolicy instantiates a new V1ClimbConfigTriggerPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ClimbConfigTriggerPolicy() *V1ClimbConfigTriggerPolicy {
	this := V1ClimbConfigTriggerPolicy{}
	return &this
}

// NewV1ClimbConfigTriggerPolicyWithDefaults instantiates a new V1ClimbConfigTriggerPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ClimbConfigTriggerPolicyWithDefaults() *V1ClimbConfigTriggerPolicy {
	this := V1ClimbConfigTriggerPolicy{}
	return &this
}

// GetErrorRate returns the ErrorRate field value if set, zero value otherwise.
func (o *V1ClimbConfigTriggerPolicy) GetErrorRate() V1ClimbConfigTriggerPolicyErrorRate {
	if o == nil || o.ErrorRate == nil {
		var ret V1ClimbConfigTriggerPolicyErrorRate
		return ret
	}
	return *o.ErrorRate
}

// GetErrorRateOk returns a tuple with the ErrorRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClimbConfigTriggerPolicy) GetErrorRateOk() (*V1ClimbConfigTriggerPolicyErrorRate, bool) {
	if o == nil || o.ErrorRate == nil {
		return nil, false
	}
	return o.ErrorRate, true
}

// HasErrorRate returns a boolean if a field has been set.
func (o *V1ClimbConfigTriggerPolicy) HasErrorRate() bool {
	if o != nil && o.ErrorRate != nil {
		return true
	}

	return false
}

// SetErrorRate gets a reference to the given V1ClimbConfigTriggerPolicyErrorRate and assigns it to the ErrorRate field.
func (o *V1ClimbConfigTriggerPolicy) SetErrorRate(v V1ClimbConfigTriggerPolicyErrorRate) {
	o.ErrorRate = &v
}

// GetSlowRate returns the SlowRate field value if set, zero value otherwise.
func (o *V1ClimbConfigTriggerPolicy) GetSlowRate() V1ClimbConfigTriggerPolicySlowRate {
	if o == nil || o.SlowRate == nil {
		var ret V1ClimbConfigTriggerPolicySlowRate
		return ret
	}
	return *o.SlowRate
}

// GetSlowRateOk returns a tuple with the SlowRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClimbConfigTriggerPolicy) GetSlowRateOk() (*V1ClimbConfigTriggerPolicySlowRate, bool) {
	if o == nil || o.SlowRate == nil {
		return nil, false
	}
	return o.SlowRate, true
}

// HasSlowRate returns a boolean if a field has been set.
func (o *V1ClimbConfigTriggerPolicy) HasSlowRate() bool {
	if o != nil && o.SlowRate != nil {
		return true
	}

	return false
}

// SetSlowRate gets a reference to the given V1ClimbConfigTriggerPolicySlowRate and assigns it to the SlowRate field.
func (o *V1ClimbConfigTriggerPolicy) SetSlowRate(v V1ClimbConfigTriggerPolicySlowRate) {
	o.SlowRate = &v
}

func (o V1ClimbConfigTriggerPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorRate != nil {
		toSerialize["errorRate"] = o.ErrorRate
	}
	if o.SlowRate != nil {
		toSerialize["slowRate"] = o.SlowRate
	}
	return json.Marshal(toSerialize)
}

type NullableV1ClimbConfigTriggerPolicy struct {
	value *V1ClimbConfigTriggerPolicy
	isSet bool
}

func (v NullableV1ClimbConfigTriggerPolicy) Get() *V1ClimbConfigTriggerPolicy {
	return v.value
}

func (v *NullableV1ClimbConfigTriggerPolicy) Set(val *V1ClimbConfigTriggerPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ClimbConfigTriggerPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ClimbConfigTriggerPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ClimbConfigTriggerPolicy(val *V1ClimbConfigTriggerPolicy) *NullableV1ClimbConfigTriggerPolicy {
	return &NullableV1ClimbConfigTriggerPolicy{value: val, isSet: true}
}

func (v NullableV1ClimbConfigTriggerPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ClimbConfigTriggerPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

