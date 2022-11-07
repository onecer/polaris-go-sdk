# V1UserGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthToken** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Ctime** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Relation** | Pointer to [**V1UserGroupRelation**](V1UserGroupRelation.md) |  | [optional] 
**TokenEnable** | Pointer to **bool** |  | [optional] 
**UserCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1UserGroup

`func NewV1UserGroup() *V1UserGroup`

NewV1UserGroup instantiates a new V1UserGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UserGroupWithDefaults

`func NewV1UserGroupWithDefaults() *V1UserGroup`

NewV1UserGroupWithDefaults instantiates a new V1UserGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthToken

`func (o *V1UserGroup) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *V1UserGroup) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *V1UserGroup) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *V1UserGroup) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetComment

`func (o *V1UserGroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1UserGroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1UserGroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1UserGroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCtime

`func (o *V1UserGroup) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1UserGroup) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1UserGroup) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1UserGroup) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetId

`func (o *V1UserGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1UserGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1UserGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1UserGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtime

`func (o *V1UserGroup) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1UserGroup) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1UserGroup) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1UserGroup) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetName

`func (o *V1UserGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1UserGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1UserGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1UserGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *V1UserGroup) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1UserGroup) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1UserGroup) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1UserGroup) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRelation

`func (o *V1UserGroup) GetRelation() V1UserGroupRelation`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *V1UserGroup) GetRelationOk() (*V1UserGroupRelation, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *V1UserGroup) SetRelation(v V1UserGroupRelation)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *V1UserGroup) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetTokenEnable

`func (o *V1UserGroup) GetTokenEnable() bool`

GetTokenEnable returns the TokenEnable field if non-nil, zero value otherwise.

### GetTokenEnableOk

`func (o *V1UserGroup) GetTokenEnableOk() (*bool, bool)`

GetTokenEnableOk returns a tuple with the TokenEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEnable

`func (o *V1UserGroup) SetTokenEnable(v bool)`

SetTokenEnable sets TokenEnable field to given value.

### HasTokenEnable

`func (o *V1UserGroup) HasTokenEnable() bool`

HasTokenEnable returns a boolean if a field has been set.

### GetUserCount

`func (o *V1UserGroup) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *V1UserGroup) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *V1UserGroup) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *V1UserGroup) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


