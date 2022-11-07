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

// V1ConfigFileRelease struct for V1ConfigFileRelease
type V1ConfigFileRelease struct {
	Comment *string `json:"comment,omitempty"`
	Content *string `json:"content,omitempty"`
	CreateBy *string `json:"create_by,omitempty"`
	CreateTime *string `json:"create_time,omitempty"`
	FileName *string `json:"file_name,omitempty"`
	Group *string `json:"group,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Md5 *string `json:"md5,omitempty"`
	ModifyBy *string `json:"modify_by,omitempty"`
	ModifyTime *string `json:"modify_time,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewV1ConfigFileRelease instantiates a new V1ConfigFileRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ConfigFileRelease() *V1ConfigFileRelease {
	this := V1ConfigFileRelease{}
	return &this
}

// NewV1ConfigFileReleaseWithDefaults instantiates a new V1ConfigFileRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ConfigFileReleaseWithDefaults() *V1ConfigFileRelease {
	this := V1ConfigFileRelease{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *V1ConfigFileRelease) SetComment(v string) {
	o.Comment = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *V1ConfigFileRelease) SetContent(v string) {
	o.Content = &v
}

// GetCreateBy returns the CreateBy field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetCreateBy() string {
	if o == nil || o.CreateBy == nil {
		var ret string
		return ret
	}
	return *o.CreateBy
}

// GetCreateByOk returns a tuple with the CreateBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetCreateByOk() (*string, bool) {
	if o == nil || o.CreateBy == nil {
		return nil, false
	}
	return o.CreateBy, true
}

// HasCreateBy returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasCreateBy() bool {
	if o != nil && o.CreateBy != nil {
		return true
	}

	return false
}

// SetCreateBy gets a reference to the given string and assigns it to the CreateBy field.
func (o *V1ConfigFileRelease) SetCreateBy(v string) {
	o.CreateBy = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetCreateTime() string {
	if o == nil || o.CreateTime == nil {
		var ret string
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetCreateTimeOk() (*string, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given string and assigns it to the CreateTime field.
func (o *V1ConfigFileRelease) SetCreateTime(v string) {
	o.CreateTime = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *V1ConfigFileRelease) SetFileName(v string) {
	o.FileName = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *V1ConfigFileRelease) SetGroup(v string) {
	o.Group = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *V1ConfigFileRelease) SetId(v int32) {
	o.Id = &v
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetMd5() string {
	if o == nil || o.Md5 == nil {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetMd5Ok() (*string, bool) {
	if o == nil || o.Md5 == nil {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasMd5() bool {
	if o != nil && o.Md5 != nil {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *V1ConfigFileRelease) SetMd5(v string) {
	o.Md5 = &v
}

// GetModifyBy returns the ModifyBy field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetModifyBy() string {
	if o == nil || o.ModifyBy == nil {
		var ret string
		return ret
	}
	return *o.ModifyBy
}

// GetModifyByOk returns a tuple with the ModifyBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetModifyByOk() (*string, bool) {
	if o == nil || o.ModifyBy == nil {
		return nil, false
	}
	return o.ModifyBy, true
}

// HasModifyBy returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasModifyBy() bool {
	if o != nil && o.ModifyBy != nil {
		return true
	}

	return false
}

// SetModifyBy gets a reference to the given string and assigns it to the ModifyBy field.
func (o *V1ConfigFileRelease) SetModifyBy(v string) {
	o.ModifyBy = &v
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetModifyTime() string {
	if o == nil || o.ModifyTime == nil {
		var ret string
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetModifyTimeOk() (*string, bool) {
	if o == nil || o.ModifyTime == nil {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasModifyTime() bool {
	if o != nil && o.ModifyTime != nil {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given string and assigns it to the ModifyTime field.
func (o *V1ConfigFileRelease) SetModifyTime(v string) {
	o.ModifyTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1ConfigFileRelease) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *V1ConfigFileRelease) SetNamespace(v string) {
	o.Namespace = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *V1ConfigFileRelease) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1ConfigFileRelease) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *V1ConfigFileRelease) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *V1ConfigFileRelease) SetVersion(v int32) {
	o.Version = &v
}

func (o V1ConfigFileRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.CreateBy != nil {
		toSerialize["create_by"] = o.CreateBy
	}
	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Md5 != nil {
		toSerialize["md5"] = o.Md5
	}
	if o.ModifyBy != nil {
		toSerialize["modify_by"] = o.ModifyBy
	}
	if o.ModifyTime != nil {
		toSerialize["modify_time"] = o.ModifyTime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableV1ConfigFileRelease struct {
	value *V1ConfigFileRelease
	isSet bool
}

func (v NullableV1ConfigFileRelease) Get() *V1ConfigFileRelease {
	return v.value
}

func (v *NullableV1ConfigFileRelease) Set(val *V1ConfigFileRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ConfigFileRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ConfigFileRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ConfigFileRelease(val *V1ConfigFileRelease) *NullableV1ConfigFileRelease {
	return &NullableV1ConfigFileRelease{value: val, isSet: true}
}

func (v NullableV1ConfigFileRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ConfigFileRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

