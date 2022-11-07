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

// V1LoginRequest struct for V1LoginRequest
type V1LoginRequest struct {
	Name *string `json:"name,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewV1LoginRequest instantiates a new V1LoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1LoginRequest() *V1LoginRequest {
	this := V1LoginRequest{}
	return &this
}

// NewV1LoginRequestWithDefaults instantiates a new V1LoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1LoginRequestWithDefaults() *V1LoginRequest {
	this := V1LoginRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1LoginRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LoginRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1LoginRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1LoginRequest) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *V1LoginRequest) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LoginRequest) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *V1LoginRequest) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *V1LoginRequest) SetOwner(v string) {
	o.Owner = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *V1LoginRequest) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1LoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *V1LoginRequest) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *V1LoginRequest) SetPassword(v string) {
	o.Password = &v
}

func (o V1LoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableV1LoginRequest struct {
	value *V1LoginRequest
	isSet bool
}

func (v NullableV1LoginRequest) Get() *V1LoginRequest {
	return v.value
}

func (v *NullableV1LoginRequest) Set(val *V1LoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1LoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1LoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1LoginRequest(val *V1LoginRequest) *NullableV1LoginRequest {
	return &NullableV1LoginRequest{value: val, isSet: true}
}

func (v NullableV1LoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1LoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

