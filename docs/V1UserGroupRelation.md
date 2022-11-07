# V1UserGroupRelation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]V1User**](V1User.md) |  | [optional] 

## Methods

### NewV1UserGroupRelation

`func NewV1UserGroupRelation() *V1UserGroupRelation`

NewV1UserGroupRelation instantiates a new V1UserGroupRelation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UserGroupRelationWithDefaults

`func NewV1UserGroupRelationWithDefaults() *V1UserGroupRelation`

NewV1UserGroupRelationWithDefaults instantiates a new V1UserGroupRelation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *V1UserGroupRelation) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V1UserGroupRelation) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V1UserGroupRelation) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V1UserGroupRelation) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUsers

`func (o *V1UserGroupRelation) GetUsers() []V1User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V1UserGroupRelation) GetUsersOk() (*[]V1User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V1UserGroupRelation) SetUsers(v []V1User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V1UserGroupRelation) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


