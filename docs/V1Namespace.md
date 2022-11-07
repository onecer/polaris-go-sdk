# V1Namespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Ctime** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owners** | Pointer to **string** |  | [optional] 
**RemoveGroupIds** | Pointer to **[]string** |  | [optional] 
**RemoveUserIds** | Pointer to **[]string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TotalHealthInstanceCount** | Pointer to **int32** |  | [optional] 
**TotalInstanceCount** | Pointer to **int32** |  | [optional] 
**TotalServiceCount** | Pointer to **int32** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1Namespace

`func NewV1Namespace() *V1Namespace`

NewV1Namespace instantiates a new V1Namespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NamespaceWithDefaults

`func NewV1NamespaceWithDefaults() *V1Namespace`

NewV1NamespaceWithDefaults instantiates a new V1Namespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *V1Namespace) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1Namespace) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1Namespace) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1Namespace) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCtime

`func (o *V1Namespace) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1Namespace) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1Namespace) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1Namespace) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetEditable

`func (o *V1Namespace) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *V1Namespace) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *V1Namespace) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *V1Namespace) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetGroupIds

`func (o *V1Namespace) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *V1Namespace) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *V1Namespace) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *V1Namespace) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetId

`func (o *V1Namespace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Namespace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Namespace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Namespace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtime

`func (o *V1Namespace) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1Namespace) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1Namespace) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1Namespace) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetName

`func (o *V1Namespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Namespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Namespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1Namespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwners

`func (o *V1Namespace) GetOwners() string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *V1Namespace) GetOwnersOk() (*string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *V1Namespace) SetOwners(v string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *V1Namespace) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetRemoveGroupIds

`func (o *V1Namespace) GetRemoveGroupIds() []string`

GetRemoveGroupIds returns the RemoveGroupIds field if non-nil, zero value otherwise.

### GetRemoveGroupIdsOk

`func (o *V1Namespace) GetRemoveGroupIdsOk() (*[]string, bool)`

GetRemoveGroupIdsOk returns a tuple with the RemoveGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveGroupIds

`func (o *V1Namespace) SetRemoveGroupIds(v []string)`

SetRemoveGroupIds sets RemoveGroupIds field to given value.

### HasRemoveGroupIds

`func (o *V1Namespace) HasRemoveGroupIds() bool`

HasRemoveGroupIds returns a boolean if a field has been set.

### GetRemoveUserIds

`func (o *V1Namespace) GetRemoveUserIds() []string`

GetRemoveUserIds returns the RemoveUserIds field if non-nil, zero value otherwise.

### GetRemoveUserIdsOk

`func (o *V1Namespace) GetRemoveUserIdsOk() (*[]string, bool)`

GetRemoveUserIdsOk returns a tuple with the RemoveUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUserIds

`func (o *V1Namespace) SetRemoveUserIds(v []string)`

SetRemoveUserIds sets RemoveUserIds field to given value.

### HasRemoveUserIds

`func (o *V1Namespace) HasRemoveUserIds() bool`

HasRemoveUserIds returns a boolean if a field has been set.

### GetToken

`func (o *V1Namespace) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1Namespace) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1Namespace) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1Namespace) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTotalHealthInstanceCount

`func (o *V1Namespace) GetTotalHealthInstanceCount() int32`

GetTotalHealthInstanceCount returns the TotalHealthInstanceCount field if non-nil, zero value otherwise.

### GetTotalHealthInstanceCountOk

`func (o *V1Namespace) GetTotalHealthInstanceCountOk() (*int32, bool)`

GetTotalHealthInstanceCountOk returns a tuple with the TotalHealthInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHealthInstanceCount

`func (o *V1Namespace) SetTotalHealthInstanceCount(v int32)`

SetTotalHealthInstanceCount sets TotalHealthInstanceCount field to given value.

### HasTotalHealthInstanceCount

`func (o *V1Namespace) HasTotalHealthInstanceCount() bool`

HasTotalHealthInstanceCount returns a boolean if a field has been set.

### GetTotalInstanceCount

`func (o *V1Namespace) GetTotalInstanceCount() int32`

GetTotalInstanceCount returns the TotalInstanceCount field if non-nil, zero value otherwise.

### GetTotalInstanceCountOk

`func (o *V1Namespace) GetTotalInstanceCountOk() (*int32, bool)`

GetTotalInstanceCountOk returns a tuple with the TotalInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstanceCount

`func (o *V1Namespace) SetTotalInstanceCount(v int32)`

SetTotalInstanceCount sets TotalInstanceCount field to given value.

### HasTotalInstanceCount

`func (o *V1Namespace) HasTotalInstanceCount() bool`

HasTotalInstanceCount returns a boolean if a field has been set.

### GetTotalServiceCount

`func (o *V1Namespace) GetTotalServiceCount() int32`

GetTotalServiceCount returns the TotalServiceCount field if non-nil, zero value otherwise.

### GetTotalServiceCountOk

`func (o *V1Namespace) GetTotalServiceCountOk() (*int32, bool)`

GetTotalServiceCountOk returns a tuple with the TotalServiceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServiceCount

`func (o *V1Namespace) SetTotalServiceCount(v int32)`

SetTotalServiceCount sets TotalServiceCount field to given value.

### HasTotalServiceCount

`func (o *V1Namespace) HasTotalServiceCount() bool`

HasTotalServiceCount returns a boolean if a field has been set.

### GetUserIds

`func (o *V1Namespace) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *V1Namespace) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *V1Namespace) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *V1Namespace) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


