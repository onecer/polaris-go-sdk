# V1CircuitBreaker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Business** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Ctime** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Inbounds** | Pointer to [**[]V1CbRule**](V1CbRule.md) |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Outbounds** | Pointer to [**[]V1CbRule**](V1CbRule.md) |  | [optional] 
**Owners** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**ServiceNamespace** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewV1CircuitBreaker

`func NewV1CircuitBreaker() *V1CircuitBreaker`

NewV1CircuitBreaker instantiates a new V1CircuitBreaker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CircuitBreakerWithDefaults

`func NewV1CircuitBreakerWithDefaults() *V1CircuitBreaker`

NewV1CircuitBreakerWithDefaults instantiates a new V1CircuitBreaker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusiness

`func (o *V1CircuitBreaker) GetBusiness() string`

GetBusiness returns the Business field if non-nil, zero value otherwise.

### GetBusinessOk

`func (o *V1CircuitBreaker) GetBusinessOk() (*string, bool)`

GetBusinessOk returns a tuple with the Business field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusiness

`func (o *V1CircuitBreaker) SetBusiness(v string)`

SetBusiness sets Business field to given value.

### HasBusiness

`func (o *V1CircuitBreaker) HasBusiness() bool`

HasBusiness returns a boolean if a field has been set.

### GetComment

`func (o *V1CircuitBreaker) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1CircuitBreaker) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1CircuitBreaker) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1CircuitBreaker) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCtime

`func (o *V1CircuitBreaker) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1CircuitBreaker) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1CircuitBreaker) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1CircuitBreaker) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetDepartment

`func (o *V1CircuitBreaker) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *V1CircuitBreaker) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *V1CircuitBreaker) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *V1CircuitBreaker) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetId

`func (o *V1CircuitBreaker) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1CircuitBreaker) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1CircuitBreaker) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1CircuitBreaker) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInbounds

`func (o *V1CircuitBreaker) GetInbounds() []V1CbRule`

GetInbounds returns the Inbounds field if non-nil, zero value otherwise.

### GetInboundsOk

`func (o *V1CircuitBreaker) GetInboundsOk() (*[]V1CbRule, bool)`

GetInboundsOk returns a tuple with the Inbounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInbounds

`func (o *V1CircuitBreaker) SetInbounds(v []V1CbRule)`

SetInbounds sets Inbounds field to given value.

### HasInbounds

`func (o *V1CircuitBreaker) HasInbounds() bool`

HasInbounds returns a boolean if a field has been set.

### GetMtime

`func (o *V1CircuitBreaker) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1CircuitBreaker) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1CircuitBreaker) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1CircuitBreaker) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetName

`func (o *V1CircuitBreaker) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1CircuitBreaker) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1CircuitBreaker) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1CircuitBreaker) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1CircuitBreaker) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1CircuitBreaker) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1CircuitBreaker) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1CircuitBreaker) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetOutbounds

`func (o *V1CircuitBreaker) GetOutbounds() []V1CbRule`

GetOutbounds returns the Outbounds field if non-nil, zero value otherwise.

### GetOutboundsOk

`func (o *V1CircuitBreaker) GetOutboundsOk() (*[]V1CbRule, bool)`

GetOutboundsOk returns a tuple with the Outbounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutbounds

`func (o *V1CircuitBreaker) SetOutbounds(v []V1CbRule)`

SetOutbounds sets Outbounds field to given value.

### HasOutbounds

`func (o *V1CircuitBreaker) HasOutbounds() bool`

HasOutbounds returns a boolean if a field has been set.

### GetOwners

`func (o *V1CircuitBreaker) GetOwners() string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *V1CircuitBreaker) GetOwnersOk() (*string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *V1CircuitBreaker) SetOwners(v string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *V1CircuitBreaker) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetRevision

`func (o *V1CircuitBreaker) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1CircuitBreaker) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1CircuitBreaker) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1CircuitBreaker) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetService

`func (o *V1CircuitBreaker) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1CircuitBreaker) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1CircuitBreaker) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1CircuitBreaker) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceNamespace

`func (o *V1CircuitBreaker) GetServiceNamespace() string`

GetServiceNamespace returns the ServiceNamespace field if non-nil, zero value otherwise.

### GetServiceNamespaceOk

`func (o *V1CircuitBreaker) GetServiceNamespaceOk() (*string, bool)`

GetServiceNamespaceOk returns a tuple with the ServiceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNamespace

`func (o *V1CircuitBreaker) SetServiceNamespace(v string)`

SetServiceNamespace sets ServiceNamespace field to given value.

### HasServiceNamespace

`func (o *V1CircuitBreaker) HasServiceNamespace() bool`

HasServiceNamespace returns a boolean if a field has been set.

### GetToken

`func (o *V1CircuitBreaker) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *V1CircuitBreaker) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *V1CircuitBreaker) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *V1CircuitBreaker) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetVersion

`func (o *V1CircuitBreaker) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1CircuitBreaker) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1CircuitBreaker) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1CircuitBreaker) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


