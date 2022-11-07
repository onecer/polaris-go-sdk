# V1ConfigFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**CreateBy** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**ModifyBy** | Pointer to **string** |  | [optional] 
**ModifyTime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**ReleaseBy** | Pointer to **string** |  | [optional] 
**ReleaseTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]V1ConfigFileTag**](V1ConfigFileTag.md) |  | [optional] 

## Methods

### NewV1ConfigFile

`func NewV1ConfigFile() *V1ConfigFile`

NewV1ConfigFile instantiates a new V1ConfigFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigFileWithDefaults

`func NewV1ConfigFileWithDefaults() *V1ConfigFile`

NewV1ConfigFileWithDefaults instantiates a new V1ConfigFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V1ConfigFile) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1ConfigFile) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1ConfigFile) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1ConfigFile) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContent

`func (o *V1ConfigFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *V1ConfigFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *V1ConfigFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *V1ConfigFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreateBy

`func (o *V1ConfigFile) GetCreateBy() string`

GetCreateBy returns the CreateBy field if non-nil, zero value otherwise.

### GetCreateByOk

`func (o *V1ConfigFile) GetCreateByOk() (*string, bool)`

GetCreateByOk returns a tuple with the CreateBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBy

`func (o *V1ConfigFile) SetCreateBy(v string)`

SetCreateBy sets CreateBy field to given value.

### HasCreateBy

`func (o *V1ConfigFile) HasCreateBy() bool`

HasCreateBy returns a boolean if a field has been set.

### GetCreateTime

`func (o *V1ConfigFile) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *V1ConfigFile) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *V1ConfigFile) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *V1ConfigFile) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetFormat

`func (o *V1ConfigFile) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V1ConfigFile) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V1ConfigFile) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *V1ConfigFile) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetGroup

`func (o *V1ConfigFile) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *V1ConfigFile) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *V1ConfigFile) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *V1ConfigFile) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *V1ConfigFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ConfigFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ConfigFile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *V1ConfigFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModifyBy

`func (o *V1ConfigFile) GetModifyBy() string`

GetModifyBy returns the ModifyBy field if non-nil, zero value otherwise.

### GetModifyByOk

`func (o *V1ConfigFile) GetModifyByOk() (*string, bool)`

GetModifyByOk returns a tuple with the ModifyBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyBy

`func (o *V1ConfigFile) SetModifyBy(v string)`

SetModifyBy sets ModifyBy field to given value.

### HasModifyBy

`func (o *V1ConfigFile) HasModifyBy() bool`

HasModifyBy returns a boolean if a field has been set.

### GetModifyTime

`func (o *V1ConfigFile) GetModifyTime() string`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *V1ConfigFile) GetModifyTimeOk() (*string, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *V1ConfigFile) SetModifyTime(v string)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *V1ConfigFile) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *V1ConfigFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ConfigFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ConfigFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1ConfigFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1ConfigFile) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1ConfigFile) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1ConfigFile) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1ConfigFile) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetReleaseBy

`func (o *V1ConfigFile) GetReleaseBy() string`

GetReleaseBy returns the ReleaseBy field if non-nil, zero value otherwise.

### GetReleaseByOk

`func (o *V1ConfigFile) GetReleaseByOk() (*string, bool)`

GetReleaseByOk returns a tuple with the ReleaseBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseBy

`func (o *V1ConfigFile) SetReleaseBy(v string)`

SetReleaseBy sets ReleaseBy field to given value.

### HasReleaseBy

`func (o *V1ConfigFile) HasReleaseBy() bool`

HasReleaseBy returns a boolean if a field has been set.

### GetReleaseTime

`func (o *V1ConfigFile) GetReleaseTime() string`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *V1ConfigFile) GetReleaseTimeOk() (*string, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *V1ConfigFile) SetReleaseTime(v string)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *V1ConfigFile) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetStatus

`func (o *V1ConfigFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1ConfigFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1ConfigFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1ConfigFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *V1ConfigFile) GetTags() []V1ConfigFileTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V1ConfigFile) GetTagsOk() (*[]V1ConfigFileTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V1ConfigFile) SetTags(v []V1ConfigFileTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V1ConfigFile) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


