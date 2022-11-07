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

// V1CbPolicyErrRateConfig struct for V1CbPolicyErrRateConfig
type V1CbPolicyErrRateConfig struct {
	Enable *bool `json:"enable,omitempty"`
	ErrorRateToOpen *int32 `json:"errorRateToOpen,omitempty"`
	ErrorRateToPreserved *int32 `json:"errorRateToPreserved,omitempty"`
	RequestVolumeThreshold *int32 `json:"requestVolumeThreshold,omitempty"`
	Specials *[]V1CbPolicyErrRateConfigSpecialConfig `json:"specials,omitempty"`
}

// NewV1CbPolicyErrRateConfig instantiates a new V1CbPolicyErrRateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CbPolicyErrRateConfig() *V1CbPolicyErrRateConfig {
	this := V1CbPolicyErrRateConfig{}
	return &this
}

// NewV1CbPolicyErrRateConfigWithDefaults instantiates a new V1CbPolicyErrRateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CbPolicyErrRateConfigWithDefaults() *V1CbPolicyErrRateConfig {
	this := V1CbPolicyErrRateConfig{}
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *V1CbPolicyErrRateConfig) GetEnable() bool {
	if o == nil || o.Enable == nil {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicyErrRateConfig) GetEnableOk() (*bool, bool) {
	if o == nil || o.Enable == nil {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *V1CbPolicyErrRateConfig) HasEnable() bool {
	if o != nil && o.Enable != nil {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *V1CbPolicyErrRateConfig) SetEnable(v bool) {
	o.Enable = &v
}

// GetErrorRateToOpen returns the ErrorRateToOpen field value if set, zero value otherwise.
func (o *V1CbPolicyErrRateConfig) GetErrorRateToOpen() int32 {
	if o == nil || o.ErrorRateToOpen == nil {
		var ret int32
		return ret
	}
	return *o.ErrorRateToOpen
}

// GetErrorRateToOpenOk returns a tuple with the ErrorRateToOpen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicyErrRateConfig) GetErrorRateToOpenOk() (*int32, bool) {
	if o == nil || o.ErrorRateToOpen == nil {
		return nil, false
	}
	return o.ErrorRateToOpen, true
}

// HasErrorRateToOpen returns a boolean if a field has been set.
func (o *V1CbPolicyErrRateConfig) HasErrorRateToOpen() bool {
	if o != nil && o.ErrorRateToOpen != nil {
		return true
	}

	return false
}

// SetErrorRateToOpen gets a reference to the given int32 and assigns it to the ErrorRateToOpen field.
func (o *V1CbPolicyErrRateConfig) SetErrorRateToOpen(v int32) {
	o.ErrorRateToOpen = &v
}

// GetErrorRateToPreserved returns the ErrorRateToPreserved field value if set, zero value otherwise.
func (o *V1CbPolicyErrRateConfig) GetErrorRateToPreserved() int32 {
	if o == nil || o.ErrorRateToPreserved == nil {
		var ret int32
		return ret
	}
	return *o.ErrorRateToPreserved
}

// GetErrorRateToPreservedOk returns a tuple with the ErrorRateToPreserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicyErrRateConfig) GetErrorRateToPreservedOk() (*int32, bool) {
	if o == nil || o.ErrorRateToPreserved == nil {
		return nil, false
	}
	return o.ErrorRateToPreserved, true
}

// HasErrorRateToPreserved returns a boolean if a field has been set.
func (o *V1CbPolicyErrRateConfig) HasErrorRateToPreserved() bool {
	if o != nil && o.ErrorRateToPreserved != nil {
		return true
	}

	return false
}

// SetErrorRateToPreserved gets a reference to the given int32 and assigns it to the ErrorRateToPreserved field.
func (o *V1CbPolicyErrRateConfig) SetErrorRateToPreserved(v int32) {
	o.ErrorRateToPreserved = &v
}

// GetRequestVolumeThreshold returns the RequestVolumeThreshold field value if set, zero value otherwise.
func (o *V1CbPolicyErrRateConfig) GetRequestVolumeThreshold() int32 {
	if o == nil || o.RequestVolumeThreshold == nil {
		var ret int32
		return ret
	}
	return *o.RequestVolumeThreshold
}

// GetRequestVolumeThresholdOk returns a tuple with the RequestVolumeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicyErrRateConfig) GetRequestVolumeThresholdOk() (*int32, bool) {
	if o == nil || o.RequestVolumeThreshold == nil {
		return nil, false
	}
	return o.RequestVolumeThreshold, true
}

// HasRequestVolumeThreshold returns a boolean if a field has been set.
func (o *V1CbPolicyErrRateConfig) HasRequestVolumeThreshold() bool {
	if o != nil && o.RequestVolumeThreshold != nil {
		return true
	}

	return false
}

// SetRequestVolumeThreshold gets a reference to the given int32 and assigns it to the RequestVolumeThreshold field.
func (o *V1CbPolicyErrRateConfig) SetRequestVolumeThreshold(v int32) {
	o.RequestVolumeThreshold = &v
}

// GetSpecials returns the Specials field value if set, zero value otherwise.
func (o *V1CbPolicyErrRateConfig) GetSpecials() []V1CbPolicyErrRateConfigSpecialConfig {
	if o == nil || o.Specials == nil {
		var ret []V1CbPolicyErrRateConfigSpecialConfig
		return ret
	}
	return *o.Specials
}

// GetSpecialsOk returns a tuple with the Specials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CbPolicyErrRateConfig) GetSpecialsOk() (*[]V1CbPolicyErrRateConfigSpecialConfig, bool) {
	if o == nil || o.Specials == nil {
		return nil, false
	}
	return o.Specials, true
}

// HasSpecials returns a boolean if a field has been set.
func (o *V1CbPolicyErrRateConfig) HasSpecials() bool {
	if o != nil && o.Specials != nil {
		return true
	}

	return false
}

// SetSpecials gets a reference to the given []V1CbPolicyErrRateConfigSpecialConfig and assigns it to the Specials field.
func (o *V1CbPolicyErrRateConfig) SetSpecials(v []V1CbPolicyErrRateConfigSpecialConfig) {
	o.Specials = &v
}

func (o V1CbPolicyErrRateConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enable != nil {
		toSerialize["enable"] = o.Enable
	}
	if o.ErrorRateToOpen != nil {
		toSerialize["errorRateToOpen"] = o.ErrorRateToOpen
	}
	if o.ErrorRateToPreserved != nil {
		toSerialize["errorRateToPreserved"] = o.ErrorRateToPreserved
	}
	if o.RequestVolumeThreshold != nil {
		toSerialize["requestVolumeThreshold"] = o.RequestVolumeThreshold
	}
	if o.Specials != nil {
		toSerialize["specials"] = o.Specials
	}
	return json.Marshal(toSerialize)
}

type NullableV1CbPolicyErrRateConfig struct {
	value *V1CbPolicyErrRateConfig
	isSet bool
}

func (v NullableV1CbPolicyErrRateConfig) Get() *V1CbPolicyErrRateConfig {
	return v.value
}

func (v *NullableV1CbPolicyErrRateConfig) Set(val *V1CbPolicyErrRateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CbPolicyErrRateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CbPolicyErrRateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CbPolicyErrRateConfig(val *V1CbPolicyErrRateConfig) *NullableV1CbPolicyErrRateConfig {
	return &NullableV1CbPolicyErrRateConfig{value: val, isSet: true}
}

func (v NullableV1CbPolicyErrRateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CbPolicyErrRateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

