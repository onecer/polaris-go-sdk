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

// V1Instance struct for V1Instance
type V1Instance struct {
	Ctime *string `json:"ctime,omitempty"`
	EnableHealthCheck *bool `json:"enable_health_check,omitempty"`
	HealthCheck *V1HealthCheck `json:"health_check,omitempty"`
	Healthy *bool `json:"healthy,omitempty"`
	Host *string `json:"host,omitempty"`
	Id *string `json:"id,omitempty"`
	Isolate *bool `json:"isolate,omitempty"`
	Location *V1Location `json:"location,omitempty"`
	LogicSet *string `json:"logic_set,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Mtime *string `json:"mtime,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Port *int32 `json:"port,omitempty"`
	Priority *int32 `json:"priority,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Revision *string `json:"revision,omitempty"`
	Service *string `json:"service,omitempty"`
	ServiceToken *string `json:"service_token,omitempty"`
	Version *string `json:"version,omitempty"`
	VpcId *string `json:"vpc_id,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
}

// NewV1Instance instantiates a new V1Instance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Instance() *V1Instance {
	this := V1Instance{}
	return &this
}

// NewV1InstanceWithDefaults instantiates a new V1Instance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1InstanceWithDefaults() *V1Instance {
	this := V1Instance{}
	return &this
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *V1Instance) GetCtime() string {
	if o == nil || o.Ctime == nil {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetCtimeOk() (*string, bool) {
	if o == nil || o.Ctime == nil {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *V1Instance) HasCtime() bool {
	if o != nil && o.Ctime != nil {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *V1Instance) SetCtime(v string) {
	o.Ctime = &v
}

// GetEnableHealthCheck returns the EnableHealthCheck field value if set, zero value otherwise.
func (o *V1Instance) GetEnableHealthCheck() bool {
	if o == nil || o.EnableHealthCheck == nil {
		var ret bool
		return ret
	}
	return *o.EnableHealthCheck
}

// GetEnableHealthCheckOk returns a tuple with the EnableHealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetEnableHealthCheckOk() (*bool, bool) {
	if o == nil || o.EnableHealthCheck == nil {
		return nil, false
	}
	return o.EnableHealthCheck, true
}

// HasEnableHealthCheck returns a boolean if a field has been set.
func (o *V1Instance) HasEnableHealthCheck() bool {
	if o != nil && o.EnableHealthCheck != nil {
		return true
	}

	return false
}

// SetEnableHealthCheck gets a reference to the given bool and assigns it to the EnableHealthCheck field.
func (o *V1Instance) SetEnableHealthCheck(v bool) {
	o.EnableHealthCheck = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *V1Instance) GetHealthCheck() V1HealthCheck {
	if o == nil || o.HealthCheck == nil {
		var ret V1HealthCheck
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetHealthCheckOk() (*V1HealthCheck, bool) {
	if o == nil || o.HealthCheck == nil {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *V1Instance) HasHealthCheck() bool {
	if o != nil && o.HealthCheck != nil {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given V1HealthCheck and assigns it to the HealthCheck field.
func (o *V1Instance) SetHealthCheck(v V1HealthCheck) {
	o.HealthCheck = &v
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *V1Instance) GetHealthy() bool {
	if o == nil || o.Healthy == nil {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetHealthyOk() (*bool, bool) {
	if o == nil || o.Healthy == nil {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *V1Instance) HasHealthy() bool {
	if o != nil && o.Healthy != nil {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *V1Instance) SetHealthy(v bool) {
	o.Healthy = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *V1Instance) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *V1Instance) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *V1Instance) SetHost(v string) {
	o.Host = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1Instance) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1Instance) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V1Instance) SetId(v string) {
	o.Id = &v
}

// GetIsolate returns the Isolate field value if set, zero value otherwise.
func (o *V1Instance) GetIsolate() bool {
	if o == nil || o.Isolate == nil {
		var ret bool
		return ret
	}
	return *o.Isolate
}

// GetIsolateOk returns a tuple with the Isolate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetIsolateOk() (*bool, bool) {
	if o == nil || o.Isolate == nil {
		return nil, false
	}
	return o.Isolate, true
}

// HasIsolate returns a boolean if a field has been set.
func (o *V1Instance) HasIsolate() bool {
	if o != nil && o.Isolate != nil {
		return true
	}

	return false
}

// SetIsolate gets a reference to the given bool and assigns it to the Isolate field.
func (o *V1Instance) SetIsolate(v bool) {
	o.Isolate = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *V1Instance) GetLocation() V1Location {
	if o == nil || o.Location == nil {
		var ret V1Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetLocationOk() (*V1Location, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *V1Instance) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given V1Location and assigns it to the Location field.
func (o *V1Instance) SetLocation(v V1Location) {
	o.Location = &v
}

// GetLogicSet returns the LogicSet field value if set, zero value otherwise.
func (o *V1Instance) GetLogicSet() string {
	if o == nil || o.LogicSet == nil {
		var ret string
		return ret
	}
	return *o.LogicSet
}

// GetLogicSetOk returns a tuple with the LogicSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetLogicSetOk() (*string, bool) {
	if o == nil || o.LogicSet == nil {
		return nil, false
	}
	return o.LogicSet, true
}

// HasLogicSet returns a boolean if a field has been set.
func (o *V1Instance) HasLogicSet() bool {
	if o != nil && o.LogicSet != nil {
		return true
	}

	return false
}

// SetLogicSet gets a reference to the given string and assigns it to the LogicSet field.
func (o *V1Instance) SetLogicSet(v string) {
	o.LogicSet = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *V1Instance) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *V1Instance) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *V1Instance) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMtime returns the Mtime field value if set, zero value otherwise.
func (o *V1Instance) GetMtime() string {
	if o == nil || o.Mtime == nil {
		var ret string
		return ret
	}
	return *o.Mtime
}

// GetMtimeOk returns a tuple with the Mtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetMtimeOk() (*string, bool) {
	if o == nil || o.Mtime == nil {
		return nil, false
	}
	return o.Mtime, true
}

// HasMtime returns a boolean if a field has been set.
func (o *V1Instance) HasMtime() bool {
	if o != nil && o.Mtime != nil {
		return true
	}

	return false
}

// SetMtime gets a reference to the given string and assigns it to the Mtime field.
func (o *V1Instance) SetMtime(v string) {
	o.Mtime = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1Instance) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1Instance) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1Instance) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V1Instance) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V1Instance) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *V1Instance) SetPort(v int32) {
	o.Port = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *V1Instance) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *V1Instance) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *V1Instance) SetPriority(v int32) {
	o.Priority = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *V1Instance) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *V1Instance) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *V1Instance) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *V1Instance) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *V1Instance) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *V1Instance) SetRevision(v string) {
	o.Revision = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *V1Instance) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *V1Instance) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *V1Instance) SetService(v string) {
	o.Service = &v
}

// GetServiceToken returns the ServiceToken field value if set, zero value otherwise.
func (o *V1Instance) GetServiceToken() string {
	if o == nil || o.ServiceToken == nil {
		var ret string
		return ret
	}
	return *o.ServiceToken
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetServiceTokenOk() (*string, bool) {
	if o == nil || o.ServiceToken == nil {
		return nil, false
	}
	return o.ServiceToken, true
}

// HasServiceToken returns a boolean if a field has been set.
func (o *V1Instance) HasServiceToken() bool {
	if o != nil && o.ServiceToken != nil {
		return true
	}

	return false
}

// SetServiceToken gets a reference to the given string and assigns it to the ServiceToken field.
func (o *V1Instance) SetServiceToken(v string) {
	o.ServiceToken = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V1Instance) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V1Instance) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *V1Instance) SetVersion(v string) {
	o.Version = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *V1Instance) GetVpcId() string {
	if o == nil || o.VpcId == nil {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetVpcIdOk() (*string, bool) {
	if o == nil || o.VpcId == nil {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *V1Instance) HasVpcId() bool {
	if o != nil && o.VpcId != nil {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *V1Instance) SetVpcId(v string) {
	o.VpcId = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *V1Instance) GetWeight() int32 {
	if o == nil || o.Weight == nil {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Instance) GetWeightOk() (*int32, bool) {
	if o == nil || o.Weight == nil {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *V1Instance) HasWeight() bool {
	if o != nil && o.Weight != nil {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *V1Instance) SetWeight(v int32) {
	o.Weight = &v
}

func (o V1Instance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ctime != nil {
		toSerialize["ctime"] = o.Ctime
	}
	if o.EnableHealthCheck != nil {
		toSerialize["enable_health_check"] = o.EnableHealthCheck
	}
	if o.HealthCheck != nil {
		toSerialize["health_check"] = o.HealthCheck
	}
	if o.Healthy != nil {
		toSerialize["healthy"] = o.Healthy
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Isolate != nil {
		toSerialize["isolate"] = o.Isolate
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.LogicSet != nil {
		toSerialize["logic_set"] = o.LogicSet
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Mtime != nil {
		toSerialize["mtime"] = o.Mtime
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ServiceToken != nil {
		toSerialize["service_token"] = o.ServiceToken
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.VpcId != nil {
		toSerialize["vpc_id"] = o.VpcId
	}
	if o.Weight != nil {
		toSerialize["weight"] = o.Weight
	}
	return json.Marshal(toSerialize)
}

type NullableV1Instance struct {
	value *V1Instance
	isSet bool
}

func (v NullableV1Instance) Get() *V1Instance {
	return v.value
}

func (v *NullableV1Instance) Set(val *V1Instance) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Instance) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Instance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Instance(val *V1Instance) *NullableV1Instance {
	return &NullableV1Instance{value: val, isSet: true}
}

func (v NullableV1Instance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Instance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


