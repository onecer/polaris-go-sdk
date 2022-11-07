# V1ServiceAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**AliasNamespace** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Ctime** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Owners** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1ServiceAlias

`func NewV1ServiceAlias() *V1ServiceAlias`

NewV1ServiceAlias instantiates a new V1ServiceAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ServiceAliasWithDefaults

`func NewV1ServiceAliasWithDefaults() *V1ServiceAlias`

NewV1ServiceAliasWithDefaults instantiates a new V1ServiceAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *V1ServiceAlias) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *V1ServiceAlias) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *V1ServiceAlias) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *V1ServiceAlias) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAliasNamespace

`func (o *V1ServiceAlias) GetAliasNamespace() string`

GetAliasNamespace returns the AliasNamespace field if non-nil, zero value otherwise.

### GetAliasNamespaceOk

`func (o *V1ServiceAlias) GetAliasNamespaceOk() (*string, bool)`

GetAliasNamespaceOk returns a tuple with the AliasNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasNamespace

`func (o *V1ServiceAlias) SetAliasNamespace(v string)`

SetAliasNamespace sets AliasNamespace field to given value.

### HasAliasNamespace

`func (o *V1ServiceAlias) HasAliasNamespace() bool`

HasAliasNamespace returns a boolean if a field has been set.

### GetComment

`func (o *V1ServiceAlias) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1ServiceAlias) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1ServiceAlias) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1ServiceAlias) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCtime

`func (o *V1ServiceAlias) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1ServiceAlias) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1ServiceAlias) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1ServiceAlias) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetEditable

`func (o *V1ServiceAlias) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *V1ServiceAlias) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *V1ServiceAlias) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *V1ServiceAlias) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetId

`func (o *V1ServiceAlias) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ServiceAlias) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ServiceAlias) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1ServiceAlias) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtime

`func (o *V1ServiceAlias) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1ServiceAlias) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1ServiceAlias) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1ServiceAlias) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetNamespace

`func (o *V1ServiceAlias) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1ServiceAlias) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1ServiceAlias) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1ServiceAlias) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOwners

`func (o *V1ServiceAlias) GetOwners() string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *V1ServiceAlias) GetOwnersOk() (*string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *V1ServiceAlias) SetOwners(v string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *V1ServiceAlias) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetService

`func (o *V1ServiceAlias) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1ServiceAlias) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1ServiceAlias) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1ServiceAlias) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceToken

`func (o *V1ServiceAlias) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *V1ServiceAlias) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *V1ServiceAlias) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *V1ServiceAlias) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetType

`func (o *V1ServiceAlias) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ServiceAlias) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ServiceAlias) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *V1ServiceAlias) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


