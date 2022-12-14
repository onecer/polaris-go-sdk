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

// V1RateLimitCluster struct for V1RateLimitCluster
type V1RateLimitCluster struct {
	Namespace *string `json:"namespace,omitempty"`
	Service *string `json:"service,omitempty"`
}

// NewV1RateLimitCluster instantiates a new V1RateLimitCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RateLimitCluster() *V1RateLimitCluster {
	this := V1RateLimitCluster{}
	return &this
}

// NewV1RateLimitClusterWithDefaults instantiates a new V1RateLimitCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RateLimitClusterWithDefaults() *V1RateLimitCluster {
	this := V1RateLimitCluster{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1RateLimitCluster) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RateLimitCluster) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1RateLimitCluster) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1RateLimitCluster) SetNamespace(v string) {
	o.Namespace = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *V1RateLimitCluster) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RateLimitCluster) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *V1RateLimitCluster) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *V1RateLimitCluster) SetService(v string) {
	o.Service = &v
}

func (o V1RateLimitCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	return json.Marshal(toSerialize)
}

type NullableV1RateLimitCluster struct {
	value *V1RateLimitCluster
	isSet bool
}

func (v NullableV1RateLimitCluster) Get() *V1RateLimitCluster {
	return v.value
}

func (v *NullableV1RateLimitCluster) Set(val *V1RateLimitCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RateLimitCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RateLimitCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RateLimitCluster(val *V1RateLimitCluster) *NullableV1RateLimitCluster {
	return &NullableV1RateLimitCluster{value: val, isSet: true}
}

func (v NullableV1RateLimitCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RateLimitCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


