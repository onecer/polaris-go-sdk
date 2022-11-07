# V1ConfigFileGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**CreateBy** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**FileCount** | Pointer to **int32** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ModifyBy** | Pointer to **string** |  | [optional] 
**ModifyTime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**RemoveGroupIds** | Pointer to **[]string** |  | [optional] 
**RemoveUserIds** | Pointer to **[]string** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1ConfigFileGroup

`func NewV1ConfigFileGroup() *V1ConfigFileGroup`

NewV1ConfigFileGroup instantiates a new V1ConfigFileGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigFileGroupWithDefaults

`func NewV1ConfigFileGroupWithDefaults() *V1ConfigFileGroup`

NewV1ConfigFileGroupWithDefaults instantiates a new V1ConfigFileGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V1ConfigFileGroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1ConfigFileGroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1ConfigFileGroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1ConfigFileGroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreateBy

`func (o *V1ConfigFileGroup) GetCreateBy() string`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *V1ConfigFileGroup) GetCreateByOk() (*string, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *V1ConfigFileGroup) SetCreateBy(v string)`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *V1ConfigFileGroup) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *V1ConfigFileGroup) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *V1ConfigFileGroup) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *V1ConfigFileGroup) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *V1ConfigFileGroup) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetEditable

`func (o *V1ConfigFileGroup) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *V1ConfigFileGroup) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *V1ConfigFileGroup) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *V1ConfigFileGroup) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetFileCount

`func (o *V1ConfigFileGroup) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *V1ConfigFileGroup) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *V1ConfigFileGroup) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *V1ConfigFileGroup) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetGroupIds

`func (o *V1ConfigFileGroup) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *V1ConfigFileGroup) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *V1ConfigFileGroup) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *V1ConfigFileGroup) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetId

`func (o *V1ConfigFileGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ConfigFileGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ConfigFileGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V1ConfigFileGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifyBy

`func (o *V1ConfigFileGroup) GetModifyBy() string`

GetModifyBy returns the ModifyBy field if non-nil, zero value otherwise.

### GetModifyByOk

`func (o *V1ConfigFileGroup) GetModifyByOk() (*string, bool)`

GetModifyByOk returns a tuple with the ModifyBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyBy

`func (o *V1ConfigFileGroup) SetModifyBy(v string)`

SetModifyBy sets ModifyBy field to given value.

### HasModifyBy

`func (o *V1ConfigFileGroup) HasModifyBy() bool`

HasModifyBy returns a boolean if a field has been set.

### GetModifyTime

`func (o *V1ConfigFileGroup) GetModifyTime() string`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *V1ConfigFileGroup) GetModifyTimeOk() (*string, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *V1ConfigFileGroup) SetModifyTime(v string)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *V1ConfigFileGroup) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *V1ConfigFileGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ConfigFileGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ConfigFileGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ConfigFileGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1ConfigFileGroup) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1ConfigFileGroup) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1ConfigFileGroup) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1ConfigFileGroup) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOwner

`func (o *V1ConfigFileGroup) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1ConfigFileGroup) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1ConfigFileGroup) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1ConfigFileGroup) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRemoveGroupIds

`func (o *V1ConfigFileGroup) GetRemoveGroupIds() []string`

GetRemoveGroupIds returns the RemoveGroupIds field if non-nil, zero value otherwise.

### GetRemoveGroupIdsOk

`func (o *V1ConfigFileGroup) GetRemoveGroupIdsOk() (*[]string, bool)`

GetRemoveGroupIdsOk returns a tuple with the RemoveGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveGroupIds

`func (o *V1ConfigFileGroup) SetRemoveGroupIds(v []string)`

SetRemoveGroupIds sets RemoveGroupIds field to given value.

### HasRemoveGroupIds

`func (o *V1ConfigFileGroup) HasRemoveGroupIds() bool`

HasRemoveGroupIds returns a boolean if a field has been set.

### GetRemoveUserIds

`func (o *V1ConfigFileGroup) GetRemoveUserIds() []string`

GetRemoveUserIds returns the RemoveUserIds field if non-nil, zero value otherwise.

### GetRemoveUserIdsOk

`func (o *V1ConfigFileGroup) GetRemoveUserIdsOk() (*[]string, bool)`

GetRemoveUserIdsOk returns a tuple with the RemoveUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUserIds

`func (o *V1ConfigFileGroup) SetRemoveUserIds(v []string)`

SetRemoveUserIds sets RemoveUserIds field to given value.

### HasRemoveUserIds

`func (o *V1ConfigFileGroup) HasRemoveUserIds() bool`

HasRemoveUserIds returns a boolean if a field has been set.

### GetUserIds

`func (o *V1ConfigFileGroup) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *V1ConfigFileGroup) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *V1ConfigFileGroup) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *V1ConfigFileGroup) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


