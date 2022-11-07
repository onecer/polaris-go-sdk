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

// V1RateLimit struct for V1RateLimit
type V1RateLimit struct {
	Revision *string `json:"revision,omitempty"`
	Rules *[]V1Rule `json:"rules,omitempty"`
}

// NewV1RateLimit instantiates a new V1RateLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1RateLimit() *V1RateLimit {
	this := V1RateLimit{}
	return &this
}

// NewV1RateLimitWithDefaults instantiates a new V1RateLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RateLimitWithDefaults() *V1RateLimit {
	this := V1RateLimit{}
	return &this
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *V1RateLimit) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RateLimit) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *V1RateLimit) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *V1RateLimit) SetRevision(v string) {
	o.Revision = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *V1RateLimit) GetRules() []V1Rule {
	if o == nil || o.Rules == nil {
		var ret []V1Rule
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1RateLimit) GetRulesOk() (*[]V1Rule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *V1RateLimit) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []V1Rule and assigns it to the Rules field.
func (o *V1RateLimit) SetRules(v []V1Rule) {
	o.Rules = &v
}

func (o V1RateLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableV1RateLimit struct {
	value *V1RateLimit
	isSet bool
}

func (v NullableV1RateLimit) Get() *V1RateLimit {
	return v.value
}

func (v *NullableV1RateLimit) Set(val *V1RateLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableV1RateLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableV1RateLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1RateLimit(val *V1RateLimit) *NullableV1RateLimit {
	return &NullableV1RateLimit{value: val, isSet: true}
}

func (v NullableV1RateLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1RateLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

