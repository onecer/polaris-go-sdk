# V2Routing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctime** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**Etime** | Pointer to **string** |  | [optional] 
**ExtendInfo** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**RoutingConfig** | Pointer to [**AnypbAny**](AnypbAny.md) |  | [optional] 
**RoutingPolicy** | Pointer to **int32** |  | [optional] 

## Methods

### NewV2Routing

`func NewV2Routing() *V2Routing`

NewV2Routing instantiates a new V2Routing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2RoutingWithDefaults

`func NewV2RoutingWithDefaults() *V2Routing`

NewV2RoutingWithDefaults instantiates a new V2Routing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtime

`func (o *V2Routing) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V2Routing) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V2Routing) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V2Routing) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetDescription

`func (o *V2Routing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V2Routing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V2Routing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V2Routing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnable

`func (o *V2Routing) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *V2Routing) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *V2Routing) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *V2Routing) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEtime

`func (o *V2Routing) GetEtime() string`

GetEtime returns the Etime field if non-nil, zero value otherwise.

### GetEtimeOk

`func (o *V2Routing) GetEtimeOk() (*string, bool)`

GetEtimeOk returns a tuple with the Etime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtime

`func (o *V2Routing) SetEtime(v string)`

SetEtime sets Etime field to given value.

### HasEtime

`func (o *V2Routing) HasEtime() bool`

HasEtime returns a boolean if a field has been set.

### GetExtendInfo

`func (o *V2Routing) GetExtendInfo() map[string]string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *V2Routing) GetExtendInfoOk() (*map[string]string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *V2Routing) SetExtendInfo(v map[string]string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *V2Routing) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetId

`func (o *V2Routing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V2Routing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V2Routing) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V2Routing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtime

`func (o *V2Routing) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V2Routing) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V2Routing) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V2Routing) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetName

`func (o *V2Routing) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V2Routing) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V2Routing) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V2Routing) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V2Routing) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V2Routing) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V2Routing) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V2Routing) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPriority

`func (o *V2Routing) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V2Routing) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V2Routing) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V2Routing) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRevision

`func (o *V2Routing) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V2Routing) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V2Routing) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V2Routing) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRoutingConfig

`func (o *V2Routing) GetRoutingConfig() AnypbAny`

GetRoutingConfig returns the RoutingConfig field if non-nil, zero value otherwise.

### GetRoutingConfigOk

`func (o *V2Routing) GetRoutingConfigOk() (*AnypbAny, bool)`

GetRoutingConfigOk returns a tuple with the RoutingConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingConfig

`func (o *V2Routing) SetRoutingConfig(v AnypbAny)`

SetRoutingConfig sets RoutingConfig field to given value.

### HasRoutingConfig

`func (o *V2Routing) HasRoutingConfig() bool`

HasRoutingConfig returns a boolean if a field has been set.

### GetRoutingPolicy

`func (o *V2Routing) GetRoutingPolicy() int32`

GetRoutingPolicy returns the RoutingPolicy field if non-nil, zero value otherwise.

### GetRoutingPolicyOk

`func (o *V2Routing) GetRoutingPolicyOk() (*int32, bool)`

GetRoutingPolicyOk returns a tuple with the RoutingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingPolicy

`func (o *V2Routing) SetRoutingPolicy(v int32)`

SetRoutingPolicy sets RoutingPolicy field to given value.

### HasRoutingPolicy

`func (o *V2Routing) HasRoutingPolicy() bool`

HasRoutingPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


