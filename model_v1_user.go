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

// V1User struct for V1User
type V1User struct {
	AuthToken *string `json:"auth_token,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Ctime *string `json:"ctime,omitempty"`
	Email *string `json:"email,omitempty"`
	Id *string `json:"id,omitempty"`
	Mobile *string `json:"mobile,omitempty"`
	Mtime *string `json:"mtime,omitempty"`
	Name *string `json:"name,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Password *string `json:"password,omitempty"`
	Source *string `json:"source,omitempty"`
	TokenEnable *bool `json:"token_enable,omitempty"`
	UserType *string `json:"user_type,omitempty"`
}

// NewV1User instantiates a new V1User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1User() *V1User {
	this := V1User{}
	return &this
}

// NewV1UserWithDefaults instantiates a new V1User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1UserWithDefaults() *V1User {
	this := V1User{}
	return &this
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *V1User) GetAuthToken() string {
	if o == nil || o.AuthToken == nil {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetAuthTokenOk() (*string, bool) {
	if o == nil || o.AuthToken == nil {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *V1User) HasAuthToken() bool {
	if o != nil && o.AuthToken != nil {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *V1User) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *V1User) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *V1User) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *V1User) SetComment(v string) {
	o.Comment = &v
}

// GetCtime returns the Ctime field value if set, zero value otherwise.
func (o *V1User) GetCtime() string {
	if o == nil || o.Ctime == nil {
		var ret string
		return ret
	}
	return *o.Ctime
}

// GetCtimeOk returns a tuple with the Ctime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetCtimeOk() (*string, bool) {
	if o == nil || o.Ctime == nil {
		return nil, false
	}
	return o.Ctime, true
}

// HasCtime returns a boolean if a field has been set.
func (o *V1User) HasCtime() bool {
	if o != nil && o.Ctime != nil {
		return true
	}

	return false
}

// SetCtime gets a reference to the given string and assigns it to the Ctime field.
func (o *V1User) SetCtime(v string) {
	o.Ctime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *V1User) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *V1User) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *V1User) SetEmail(v string) {
	o.Email = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V1User) SetId(v string) {
	o.Id = &v
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *V1User) GetMobile() string {
	if o == nil || o.Mobile == nil {
		var ret string
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetMobileOk() (*string, bool) {
	if o == nil || o.Mobile == nil {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *V1User) HasMobile() bool {
	if o != nil && o.Mobile != nil {
		return true
	}

	return false
}

// SetMobile gets a reference to the given string and assigns it to the Mobile field.
func (o *V1User) SetMobile(v string) {
	o.Mobile = &v
}

// GetMtime returns the Mtime field value if set, zero value otherwise.
func (o *V1User) GetMtime() string {
	if o == nil || o.Mtime == nil {
		var ret string
		return ret
	}
	return *o.Mtime
}

// GetMtimeOk returns a tuple with the Mtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetMtimeOk() (*string, bool) {
	if o == nil || o.Mtime == nil {
		return nil, false
	}
	return o.Mtime, true
}

// HasMtime returns a boolean if a field has been set.
func (o *V1User) HasMtime() bool {
	if o != nil && o.Mtime != nil {
		return true
	}

	return false
}

// SetMtime gets a reference to the given string and assigns it to the Mtime field.
func (o *V1User) SetMtime(v string) {
	o.Mtime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1User) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *V1User) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *V1User) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *V1User) SetOwner(v string) {
	o.Owner = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *V1User) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *V1User) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *V1User) SetPassword(v string) {
	o.Password = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1User) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1User) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *V1User) SetSource(v string) {
	o.Source = &v
}

// GetTokenEnable returns the TokenEnable field value if set, zero value otherwise.
func (o *V1User) GetTokenEnable() bool {
	if o == nil || o.TokenEnable == nil {
		var ret bool
		return ret
	}
	return *o.TokenEnable
}

// GetTokenEnableOk returns a tuple with the TokenEnable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetTokenEnableOk() (*bool, bool) {
	if o == nil || o.TokenEnable == nil {
		return nil, false
	}
	return o.TokenEnable, true
}

// HasTokenEnable returns a boolean if a field has been set.
func (o *V1User) HasTokenEnable() bool {
	if o != nil && o.TokenEnable != nil {
		return true
	}

	return false
}

// SetTokenEnable gets a reference to the given bool and assigns it to the TokenEnable field.
func (o *V1User) SetTokenEnable(v bool) {
	o.TokenEnable = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *V1User) GetUserType() string {
	if o == nil || o.UserType == nil {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1User) GetUserTypeOk() (*string, bool) {
	if o == nil || o.UserType == nil {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *V1User) HasUserType() bool {
	if o != nil && o.UserType != nil {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *V1User) SetUserType(v string) {
	o.UserType = &v
}

func (o V1User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthToken != nil {
		toSerialize["auth_token"] = o.AuthToken
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Ctime != nil {
		toSerialize["ctime"] = o.Ctime
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Mobile != nil {
		toSerialize["mobile"] = o.Mobile
	}
	if o.Mtime != nil {
		toSerialize["mtime"] = o.Mtime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.TokenEnable != nil {
		toSerialize["token_enable"] = o.TokenEnable
	}
	if o.UserType != nil {
		toSerialize["user_type"] = o.UserType
	}
	return json.Marshal(toSerialize)
}

type NullableV1User struct {
	value *V1User
	isSet bool
}

func (v NullableV1User) Get() *V1User {
	return v.value
}

func (v *NullableV1User) Set(val *V1User) {
	v.value = val
	v.isSet = true
}

func (v NullableV1User) IsSet() bool {
	return v.isSet
}

func (v *NullableV1User) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1User(val *V1User) *NullableV1User {
	return &NullableV1User{value: val, isSet: true}
}

func (v NullableV1User) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1User) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


