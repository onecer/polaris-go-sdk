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

// V1Route struct for V1Route
type V1Route struct {
	Destinations *[]V1Destination `json:"destinations,omitempty"`
	ExtendInfo *map[string]string `json:"extendInfo,omitempty"`
	Sources *[]V1Source `json:"sources,omitempty"`
}

// NewV1Route instantiates a new V1Route object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Route() *V1Route {
	this := V1Route{}
	return &this
}

// NewV1RouteWithDefaults instantiates a new V1Route object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1RouteWithDefaults() *V1Route {
	this := V1Route{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *V1Route) GetDestinations() []V1Destination {
	if o == nil || o.Destinations == nil {
		var ret []V1Destination
		return ret
	}
	return *o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Route) GetDestinationsOk() (*[]V1Destination, bool) {
	if o == nil || o.Destinations == nil {
		return nil, false
	}
	return o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *V1Route) HasDestinations() bool {
	if o != nil && o.Destinations != nil {
		return true
	}

	return false
}

// SetDestinations gets a reference to the given []V1Destination and assigns it to the Destinations field.
func (o *V1Route) SetDestinations(v []V1Destination) {
	o.Destinations = &v
}

// GetExtendInfo returns the ExtendInfo field value if set, zero value otherwise.
func (o *V1Route) GetExtendInfo() map[string]string {
	if o == nil || o.ExtendInfo == nil {
		var ret map[string]string
		return ret
	}
	return *o.ExtendInfo
}

// GetExtendInfoOk returns a tuple with the ExtendInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Route) GetExtendInfoOk() (*map[string]string, bool) {
	if o == nil || o.ExtendInfo == nil {
		return nil, false
	}
	return o.ExtendInfo, true
}

// HasExtendInfo returns a boolean if a field has been set.
func (o *V1Route) HasExtendInfo() bool {
	if o != nil && o.ExtendInfo != nil {
		return true
	}

	return false
}

// SetExtendInfo gets a reference to the given map[string]string and assigns it to the ExtendInfo field.
func (o *V1Route) SetExtendInfo(v map[string]string) {
	o.ExtendInfo = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *V1Route) GetSources() []V1Source {
	if o == nil || o.Sources == nil {
		var ret []V1Source
		return ret
	}
	return *o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Route) GetSourcesOk() (*[]V1Source, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *V1Route) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []V1Source and assigns it to the Sources field.
func (o *V1Route) SetSources(v []V1Source) {
	o.Sources = &v
}

func (o V1Route) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Destinations != nil {
		toSerialize["destinations"] = o.Destinations
	}
	if o.ExtendInfo != nil {
		toSerialize["extendInfo"] = o.ExtendInfo
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	return json.Marshal(toSerialize)
}

type NullableV1Route struct {
	value *V1Route
	isSet bool
}

func (v NullableV1Route) Get() *V1Route {
	return v.value
}

func (v *NullableV1Route) Set(val *V1Route) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Route) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Route) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Route(val *V1Route) *NullableV1Route {
	return &NullableV1Route{value: val, isSet: true}
}

func (v NullableV1Route) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Route) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


