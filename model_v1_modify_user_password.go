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

// V1ModifyUserPassword struct for V1ModifyUserPassword
type V1ModifyUserPassword struct {
	Id *string `json:"id,omitempty"`
	NewPassword *string `json:"new_password,omitempty"`
	OldPassword *string `json:"old_password,omitempty"`
}

// NewV1ModifyUserPassword instantiates a new V1ModifyUserPassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ModifyUserPassword() *V1ModifyUserPassword {
	this := V1ModifyUserPassword{}
	return &this
}

// NewV1ModifyUserPasswordWithDefaults instantiates a new V1ModifyUserPassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ModifyUserPasswordWithDefaults() *V1ModifyUserPassword {
	this := V1ModifyUserPassword{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1ModifyUserPassword) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ModifyUserPassword) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1ModifyUserPassword) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V1ModifyUserPassword) SetId(v string) {
	o.Id = &v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *V1ModifyUserPassword) GetNewPassword() string {
	if o == nil || o.NewPassword == nil {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ModifyUserPassword) GetNewPasswordOk() (*string, bool) {
	if o == nil || o.NewPassword == nil {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *V1ModifyUserPassword) HasNewPassword() bool {
	if o != nil && o.NewPassword != nil {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *V1ModifyUserPassword) SetNewPassword(v string) {
	o.NewPassword = &v
}

// GetOldPassword returns the OldPassword field value if set, zero value otherwise.
func (o *V1ModifyUserPassword) GetOldPassword() string {
	if o == nil || o.OldPassword == nil {
		var ret string
		return ret
	}
	return *o.OldPassword
}

// GetOldPasswordOk returns a tuple with the OldPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ModifyUserPassword) GetOldPasswordOk() (*string, bool) {
	if o == nil || o.OldPassword == nil {
		return nil, false
	}
	return o.OldPassword, true
}

// HasOldPassword returns a boolean if a field has been set.
func (o *V1ModifyUserPassword) HasOldPassword() bool {
	if o != nil && o.OldPassword != nil {
		return true
	}

	return false
}

// SetOldPassword gets a reference to the given string and assigns it to the OldPassword field.
func (o *V1ModifyUserPassword) SetOldPassword(v string) {
	o.OldPassword = &v
}

func (o V1ModifyUserPassword) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NewPassword != nil {
		toSerialize["new_password"] = o.NewPassword
	}
	if o.OldPassword != nil {
		toSerialize["old_password"] = o.OldPassword
	}
	return json.Marshal(toSerialize)
}

type NullableV1ModifyUserPassword struct {
	value *V1ModifyUserPassword
	isSet bool
}

func (v NullableV1ModifyUserPassword) Get() *V1ModifyUserPassword {
	return v.value
}

func (v *NullableV1ModifyUserPassword) Set(val *V1ModifyUserPassword) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ModifyUserPassword) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ModifyUserPassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ModifyUserPassword(val *V1ModifyUserPassword) *NullableV1ModifyUserPassword {
	return &NullableV1ModifyUserPassword{value: val, isSet: true}
}

func (v NullableV1ModifyUserPassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ModifyUserPassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

