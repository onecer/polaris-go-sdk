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

// V1ClientWatchConfigFileRequest struct for V1ClientWatchConfigFileRequest
type V1ClientWatchConfigFileRequest struct {
	ClientIp *string `json:"client_ip,omitempty"`
	ServiceName *string `json:"service_name,omitempty"`
	WatchFiles *[]V1ClientConfigFileInfo `json:"watch_files,omitempty"`
}

// NewV1ClientWatchConfigFileRequest instantiates a new V1ClientWatchConfigFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ClientWatchConfigFileRequest() *V1ClientWatchConfigFileRequest {
	this := V1ClientWatchConfigFileRequest{}
	return &this
}

// NewV1ClientWatchConfigFileRequestWithDefaults instantiates a new V1ClientWatchConfigFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ClientWatchConfigFileRequestWithDefaults() *V1ClientWatchConfigFileRequest {
	this := V1ClientWatchConfigFileRequest{}
	return &this
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *V1ClientWatchConfigFileRequest) GetClientIp() string {
	if o == nil || o.ClientIp == nil {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClientWatchConfigFileRequest) GetClientIpOk() (*string, bool) {
	if o == nil || o.ClientIp == nil {
		return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *V1ClientWatchConfigFileRequest) HasClientIp() bool {
	if o != nil && o.ClientIp != nil {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *V1ClientWatchConfigFileRequest) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *V1ClientWatchConfigFileRequest) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClientWatchConfigFileRequest) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *V1ClientWatchConfigFileRequest) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *V1ClientWatchConfigFileRequest) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetWatchFiles returns the WatchFiles field value if set, zero value otherwise.
func (o *V1ClientWatchConfigFileRequest) GetWatchFiles() []V1ClientConfigFileInfo {
	if o == nil || o.WatchFiles == nil {
		var ret []V1ClientConfigFileInfo
		return ret
	}
	return *o.WatchFiles
}

// GetWatchFilesOk returns a tuple with the WatchFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ClientWatchConfigFileRequest) GetWatchFilesOk() (*[]V1ClientConfigFileInfo, bool) {
	if o == nil || o.WatchFiles == nil {
		return nil, false
	}
	return o.WatchFiles, true
}

// HasWatchFiles returns a boolean if a field has been set.
func (o *V1ClientWatchConfigFileRequest) HasWatchFiles() bool {
	if o != nil && o.WatchFiles != nil {
		return true
	}

	return false
}

// SetWatchFiles gets a reference to the given []V1ClientConfigFileInfo and assigns it to the WatchFiles field.
func (o *V1ClientWatchConfigFileRequest) SetWatchFiles(v []V1ClientConfigFileInfo) {
	o.WatchFiles = &v
}

func (o V1ClientWatchConfigFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientIp != nil {
		toSerialize["client_ip"] = o.ClientIp
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.WatchFiles != nil {
		toSerialize["watch_files"] = o.WatchFiles
	}
	return json.Marshal(toSerialize)
}

type NullableV1ClientWatchConfigFileRequest struct {
	value *V1ClientWatchConfigFileRequest
	isSet bool
}

func (v NullableV1ClientWatchConfigFileRequest) Get() *V1ClientWatchConfigFileRequest {
	return v.value
}

func (v *NullableV1ClientWatchConfigFileRequest) Set(val *V1ClientWatchConfigFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ClientWatchConfigFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ClientWatchConfigFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ClientWatchConfigFileRequest(val *V1ClientWatchConfigFileRequest) *NullableV1ClientWatchConfigFileRequest {
	return &NullableV1ClientWatchConfigFileRequest{value: val, isSet: true}
}

func (v NullableV1ClientWatchConfigFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ClientWatchConfigFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

