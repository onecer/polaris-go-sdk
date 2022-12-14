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

// V1StrategyResources struct for V1StrategyResources
type V1StrategyResources struct {
	ConfigGroups *[]V1StrategyResourceEntry `json:"config_groups,omitempty"`
	Namespaces *[]V1StrategyResourceEntry `json:"namespaces,omitempty"`
	Services *[]V1StrategyResourceEntry `json:"services,omitempty"`
	StrategyId *string `json:"strategy_id,omitempty"`
}

// NewV1StrategyResources instantiates a new V1StrategyResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1StrategyResources() *V1StrategyResources {
	this := V1StrategyResources{}
	return &this
}

// NewV1StrategyResourcesWithDefaults instantiates a new V1StrategyResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1StrategyResourcesWithDefaults() *V1StrategyResources {
	this := V1StrategyResources{}
	return &this
}

// GetConfigGroups returns the ConfigGroups field value if set, zero value otherwise.
func (o *V1StrategyResources) GetConfigGroups() []V1StrategyResourceEntry {
	if o == nil || o.ConfigGroups == nil {
		var ret []V1StrategyResourceEntry
		return ret
	}
	return *o.ConfigGroups
}

// GetConfigGroupsOk returns a tuple with the ConfigGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StrategyResources) GetConfigGroupsOk() (*[]V1StrategyResourceEntry, bool) {
	if o == nil || o.ConfigGroups == nil {
		return nil, false
	}
	return o.ConfigGroups, true
}

// HasConfigGroups returns a boolean if a field has been set.
func (o *V1StrategyResources) HasConfigGroups() bool {
	if o != nil && o.ConfigGroups != nil {
		return true
	}

	return false
}

// SetConfigGroups gets a reference to the given []V1StrategyResourceEntry and assigns it to the ConfigGroups field.
func (o *V1StrategyResources) SetConfigGroups(v []V1StrategyResourceEntry) {
	o.ConfigGroups = &v
}

// GetNamespaces returns the Namespaces field value if set, zero value otherwise.
func (o *V1StrategyResources) GetNamespaces() []V1StrategyResourceEntry {
	if o == nil || o.Namespaces == nil {
		var ret []V1StrategyResourceEntry
		return ret
	}
	return *o.Namespaces
}

// GetNamespacesOk returns a tuple with the Namespaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StrategyResources) GetNamespacesOk() (*[]V1StrategyResourceEntry, bool) {
	if o == nil || o.Namespaces == nil {
		return nil, false
	}
	return o.Namespaces, true
}

// HasNamespaces returns a boolean if a field has been set.
func (o *V1StrategyResources) HasNamespaces() bool {
	if o != nil && o.Namespaces != nil {
		return true
	}

	return false
}

// SetNamespaces gets a reference to the given []V1StrategyResourceEntry and assigns it to the Namespaces field.
func (o *V1StrategyResources) SetNamespaces(v []V1StrategyResourceEntry) {
	o.Namespaces = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *V1StrategyResources) GetServices() []V1StrategyResourceEntry {
	if o == nil || o.Services == nil {
		var ret []V1StrategyResourceEntry
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StrategyResources) GetServicesOk() (*[]V1StrategyResourceEntry, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *V1StrategyResources) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []V1StrategyResourceEntry and assigns it to the Services field.
func (o *V1StrategyResources) SetServices(v []V1StrategyResourceEntry) {
	o.Services = &v
}

// GetStrategyId returns the StrategyId field value if set, zero value otherwise.
func (o *V1StrategyResources) GetStrategyId() string {
	if o == nil || o.StrategyId == nil {
		var ret string
		return ret
	}
	return *o.StrategyId
}

// GetStrategyIdOk returns a tuple with the StrategyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1StrategyResources) GetStrategyIdOk() (*string, bool) {
	if o == nil || o.StrategyId == nil {
		return nil, false
	}
	return o.StrategyId, true
}

// HasStrategyId returns a boolean if a field has been set.
func (o *V1StrategyResources) HasStrategyId() bool {
	if o != nil && o.StrategyId != nil {
		return true
	}

	return false
}

// SetStrategyId gets a reference to the given string and assigns it to the StrategyId field.
func (o *V1StrategyResources) SetStrategyId(v string) {
	o.StrategyId = &v
}

func (o V1StrategyResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConfigGroups != nil {
		toSerialize["config_groups"] = o.ConfigGroups
	}
	if o.Namespaces != nil {
		toSerialize["namespaces"] = o.Namespaces
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.StrategyId != nil {
		toSerialize["strategy_id"] = o.StrategyId
	}
	return json.Marshal(toSerialize)
}

type NullableV1StrategyResources struct {
	value *V1StrategyResources
	isSet bool
}

func (v NullableV1StrategyResources) Get() *V1StrategyResources {
	return v.value
}

func (v *NullableV1StrategyResources) Set(val *V1StrategyResources) {
	v.value = val
	v.isSet = true
}

func (v NullableV1StrategyResources) IsSet() bool {
	return v.isSet
}

func (v *NullableV1StrategyResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1StrategyResources(val *V1StrategyResources) *NullableV1StrategyResources {
	return &NullableV1StrategyResources{value: val, isSet: true}
}

func (v NullableV1StrategyResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1StrategyResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


