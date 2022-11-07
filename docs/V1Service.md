# V1Service

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to **string** |  | [optional] 
**CmdbMod1** | Pointer to **string** |  | [optional] 
**CmdbMod2** | Pointer to **string** |  | [optional] 
**CmdbMod3** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Ctime** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**GroupIds** | Pointer to **[]string** |  | [optional] 
**HealthyInstanceCount** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Owners** | Pointer to **string** |  | [optional] 
**PlatformId** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to **string** |  | [optional] 
**RemoveGroupIds** | Pointer to **[]string** |  | [optional] 
**RemoveUserIds** | Pointer to **[]string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TotalInstanceCount** | Pointer to **int32** |  | [optional] 
**UserIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1Service

`func NewV1Service() *V1Service`

NewV1Service instantiates a new V1Service object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ServiceWithDefaults

`func NewV1ServiceWithDefaults() *V1Service`

NewV1ServiceWithDefaults instantiates a new V1Service object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *V1Service) GetBusiness() string`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *V1Service) GetBusinessOk() (*string, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *V1Service) SetBusiness(v string)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *V1Service) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetCmdbMod1

`func (o *V1Service) GetCmdbMod1() string`

GetCmdbMod1 returns the CmdbMod1 field if non-nil, zero value otherwise.

### GetCmdbMod1Ok

`func (o *V1Service) GetCmdbMod1Ok() (*string, bool)`

GetCmdbMod1Ok returns a tuple with the CmdbMod1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbMod1

`func (o *V1Service) SetCmdbMod1(v string)`

SetCmdbMod1 sets CmdbMod1 field to given value.

### HasCmdbMod1

`func (o *V1Service) HasCmdbMod1() bool`

HasCmdbMod1 returns a boolean if a field has been set.

### GetCmdbMod2

`func (o *V1Service) GetCmdbMod2() string`

GetCmdbMod2 returns the CmdbMod2 field if non-nil, zero value otherwise.

### GetCmdbMod2Ok

`func (o *V1Service) GetCmdbMod2Ok() (*string, bool)`

GetCmdbMod2Ok returns a tuple with the CmdbMod2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbMod2

`func (o *V1Service) SetCmdbMod2(v string)`

SetCmdbMod2 sets CmdbMod2 field to given value.

### HasCmdbMod2

`func (o *V1Service) HasCmdbMod2() bool`

HasCmdbMod2 returns a boolean if a field has been set.

### GetCmdbMod3

`func (o *V1Service) GetCmdbMod3() string`

GetCmdbMod3 returns the CmdbMod3 field if non-nil, zero value otherwise.

### GetCmdbMod3Ok

`func (o *V1Service) GetCmdbMod3Ok() (*string, bool)`

GetCmdbMod3Ok returns a tuple with the CmdbMod3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmdbMod3

`func (o *V1Service) SetCmdbMod3(v string)`

SetCmdbMod3 sets CmdbMod3 field to given value.

### HasCmdbMod3

`func (o *V1Service) HasCmdbMod3() bool`

HasCmdbMod3 returns a boolean if a field has been set.

### GetComment

`func (o *V1Service) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1Service) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1Service) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1Service) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCtime

`func (o *V1Service) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1Service) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1Service) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1Service) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetDepartment

`func (o *V1Service) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *V1Service) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *V1Service) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *V1Service) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetEditable

`func (o *V1Service) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *V1Service) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *V1Service) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *V1Service) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetGroupIds

`func (o *V1Service) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *V1Service) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *V1Service) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *V1Service) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetHealthyInstanceCount

`func (o *V1Service) GetHealthyInstanceCount() int32`

GetHealthyInstanceCount returns the HealthyInstanceCount field if non-nil, zero value otherwise.

### GetHealthyInstanceCountOk

`func (o *V1Service) GetHealthyInstanceCountOk() (*int32, bool)`

GetHealthyInstanceCountOk returns a tuple with the HealthyInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyInstanceCount

`func (o *V1Service) SetHealthyInstanceCount(v int32)`

SetHealthyInstanceCount sets HealthyInstanceCount field to given value.

### HasHealthyInstanceCount

`func (o *V1Service) HasHealthyInstanceCount() bool`

HasHealthyInstanceCount returns a boolean if a field has been set.

### GetId

`func (o *V1Service) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Service) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Service) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Service) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Service) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Service) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Service) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Service) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMtime

`func (o *V1Service) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1Service) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1Service) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1Service) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetName

`func (o *V1Service) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Service) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Service) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1Service) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1Service) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1Service) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1Service) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1Service) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOwners

`func (o *V1Service) GetOwners() string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *V1Service) GetOwnersOk() (*string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *V1Service) SetOwners(v string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *V1Service) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetPlatformId

`func (o *V1Service) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *V1Service) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *V1Service) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.

### HasPlatformId

`func (o *V1Service) HasPlatformId() bool`

HasPlatformId returns a boolean if a field has been set.

### GetPorts

`func (o *V1Service) GetPorts() string`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *V1Service) GetPortsOk() (*string, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *V1Service) SetPorts(v string)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *V1Service) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetRemoveGroupIds

`func (o *V1Service) GetRemoveGroupIds() []string`

GetRemoveGroupIds returns the RemoveGroupIds field if non-nil, zero value otherwise.

### GetRemoveGroupIdsOk

`func (o *V1Service) GetRemoveGroupIdsOk() (*[]string, bool)`

GetRemoveGroupIdsOk returns a tuple with the RemoveGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveGroupIds

`func (o *V1Service) SetRemoveGroupIds(v []string)`

SetRemoveGroupIds sets RemoveGroupIds field to given value.

### HasRemoveGroupIds

`func (o *V1Service) HasRemoveGroupIds() bool`

HasRemoveGroupIds returns a boolean if a field has been set.

### GetRemoveUserIds

`func (o *V1Service) GetRemoveUserIds() []string`

GetRemoveUserIds returns the RemoveUserIds field if non-nil, zero value otherwise.

### GetRemoveUserIdsOk

`func (o *V1Service) GetRemoveUserIdsOk() (*[]string, bool)`

GetRemoveUserIdsOk returns a tuple with the RemoveUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveUserIds

`func (o *V1Service) SetRemoveUserIds(v []string)`

SetRemoveUserIds sets RemoveUserIds field to given value.

### HasRemoveUserIds

`func (o *V1Service) HasRemoveUserIds() bool`

HasRemoveUserIds returns a boolean if a field has been set.

### GetRevision

`func (o *V1Service) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1Service) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1Service) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1Service) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetToken

`func (o *V1Service) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1Service) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1Service) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1Service) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTotalInstanceCount

`func (o *V1Service) GetTotalInstanceCount() int32`

GetTotalInstanceCount returns the TotalInstanceCount field if non-nil, zero value otherwise.

### GetTotalInstanceCountOk

`func (o *V1Service) GetTotalInstanceCountOk() (*int32, bool)`

GetTotalInstanceCountOk returns a tuple with the TotalInstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInstanceCount

`func (o *V1Service) SetTotalInstanceCount(v int32)`

SetTotalInstanceCount sets TotalInstanceCount field to given value.

### HasTotalInstanceCount

`func (o *V1Service) HasTotalInstanceCount() bool`

HasTotalInstanceCount returns a boolean if a field has been set.

### GetUserIds

`func (o *V1Service) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *V1Service) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *V1Service) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *V1Service) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


