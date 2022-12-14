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

// V1StatInfo struct for V1StatInfo
type V1StatInfo struct {
	Path *string `json:"path,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Target *string `json:"target,omitempty"`
}

// NewV1StatInfo instantiates a new V1StatInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1StatInfo() *V1StatInfo {
	this := V1StatInfo{}
	return &this
}

// NewV1StatInfoWithDefaults instantiates a new V1StatInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1StatInfoWithDefaults() *V1StatInfo {
	this := V1StatInfo{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *V1StatInfo) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatInfo) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *V1StatInfo) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *V1StatInfo) SetPath(v string) {
	o.Path = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V1StatInfo) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatInfo) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V1StatInfo) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *V1StatInfo) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *V1StatInfo) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatInfo) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *V1StatInfo) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *V1StatInfo) SetProtocol(v string) {
	o.Protocol = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *V1StatInfo) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StatInfo) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *V1StatInfo) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *V1StatInfo) SetTarget(v string) {
	o.Target = &v
}

func (o V1StatInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableV1StatInfo struct {
	value *V1StatInfo
	isSet bool
}

func (v NullableV1StatInfo) Get() *V1StatInfo {
	return v.value
}

func (v *NullableV1StatInfo) Set(val *V1StatInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableV1StatInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableV1StatInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1StatInfo(val *V1StatInfo) *NullableV1StatInfo {
	return &NullableV1StatInfo{value: val, isSet: true}
}

func (v NullableV1StatInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1StatInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


