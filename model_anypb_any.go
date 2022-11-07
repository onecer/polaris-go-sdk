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

// AnypbAny struct for AnypbAny
type AnypbAny struct {
	TypeUrl *string `json:"type_url,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewAnypbAny instantiates a new AnypbAny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnypbAny() *AnypbAny {
	this := AnypbAny{}
	return &this
}

// NewAnypbAnyWithDefaults instantiates a new AnypbAny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnypbAnyWithDefaults() *AnypbAny {
	this := AnypbAny{}
	return &this
}

// GetTypeUrl returns the TypeUrl field value if set, zero value otherwise.
func (o *AnypbAny) GetTypeUrl() string {
	if o == nil || o.TypeUrl == nil {
		var ret string
		return ret
	}
	return *o.TypeUrl
}

// GetTypeUrlOk returns a tuple with the TypeUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnypbAny) GetTypeUrlOk() (*string, bool) {
	if o == nil || o.TypeUrl == nil {
		return nil, false
	}
	return o.TypeUrl, true
}

// HasTypeUrl returns a boolean if a field has been set.
func (o *AnypbAny) HasTypeUrl() bool {
	if o != nil && o.TypeUrl != nil {
		return true
	}

	return false
}

// SetTypeUrl gets a reference to the given string and assigns it to the TypeUrl field.
func (o *AnypbAny) SetTypeUrl(v string) {
	o.TypeUrl = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AnypbAny) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnypbAny) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AnypbAny) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AnypbAny) SetValue(v string) {
	o.Value = &v
}

func (o AnypbAny) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TypeUrl != nil {
		toSerialize["type_url"] = o.TypeUrl
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableAnypbAny struct {
	value *AnypbAny
	isSet bool
}

func (v NullableAnypbAny) Get() *AnypbAny {
	return v.value
}

func (v *NullableAnypbAny) Set(val *AnypbAny) {
	v.value = val
	v.isSet = true
}

func (v NullableAnypbAny) IsSet() bool {
	return v.isSet
}

func (v *NullableAnypbAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnypbAny(val *AnypbAny) *NullableAnypbAny {
	return &NullableAnypbAny{value: val, isSet: true}
}

func (v NullableAnypbAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnypbAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


