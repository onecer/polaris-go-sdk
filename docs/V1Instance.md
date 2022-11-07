# V1Instance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctime** | Pointer to **string** |  | [optional] 
**EnableHealthCheck** | Pointer to **bool** |  | [optional] 
**HealthCheck** | Pointer to [**V1HealthCheck**](V1HealthCheck.md) |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Isolate** | Pointer to **bool** |  | [optional] 
**Location** | Pointer to [**V1Location**](V1Location.md) |  | [optional] 
**LogicSet** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VpcId** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1Instance

`func NewV1Instance() *V1Instance`

NewV1Instance instantiates a new V1Instance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1InstanceWithDefaults

`func NewV1InstanceWithDefaults() *V1Instance`

NewV1InstanceWithDefaults instantiates a new V1Instance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtime

`func (o *V1Instance) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1Instance) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1Instance) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1Instance) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetEnableHealthCheck

`func (o *V1Instance) GetEnableHealthCheck() bool`

GetEnableHealthCheck returns the EnableHealthCheck field if non-nil, zero value otherwise.

### GetEnableHealthCheckOk

`func (o *V1Instance) GetEnableHealthCheckOk() (*bool, bool)`

GetEnableHealthCheckOk returns a tuple with the EnableHealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHealthCheck

`func (o *V1Instance) SetEnableHealthCheck(v bool)`

SetEnableHealthCheck sets EnableHealthCheck field to given value.

### HasEnableHealthCheck

`func (o *V1Instance) HasEnableHealthCheck() bool`

HasEnableHealthCheck returns a boolean if a field has been set.

### GetHealthCheck

`func (o *V1Instance) GetHealthCheck() V1HealthCheck`

GetHealthCheck returns the HealthCheck field if non-nil, zero value otherwise.

### GetHealthCheckOk

`func (o *V1Instance) GetHealthCheckOk() (*V1HealthCheck, bool)`

GetHealthCheckOk returns a tuple with the HealthCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheck

`func (o *V1Instance) SetHealthCheck(v V1HealthCheck)`

SetHealthCheck sets HealthCheck field to given value.

### HasHealthCheck

`func (o *V1Instance) HasHealthCheck() bool`

HasHealthCheck returns a boolean if a field has been set.

### GetHealthy

`func (o *V1Instance) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *V1Instance) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *V1Instance) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *V1Instance) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetHost

`func (o *V1Instance) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1Instance) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1Instance) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1Instance) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *V1Instance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Instance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Instance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Instance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsolate

`func (o *V1Instance) GetIsolate() bool`

GetIsolate returns the Isolate field if non-nil, zero value otherwise.

### GetIsolateOk

`func (o *V1Instance) GetIsolateOk() (*bool, bool)`

GetIsolateOk returns a tuple with the Isolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolate

`func (o *V1Instance) SetIsolate(v bool)`

SetIsolate sets Isolate field to given value.

### HasIsolate

`func (o *V1Instance) HasIsolate() bool`

HasIsolate returns a boolean if a field has been set.

### GetLocation

`func (o *V1Instance) GetLocation() V1Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1Instance) GetLocationOk() (*V1Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1Instance) SetLocation(v V1Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1Instance) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLogicSet

`func (o *V1Instance) GetLogicSet() string`

GetLogicSet returns the LogicSet field if non-nil, zero value otherwise.

### GetLogicSetOk

`func (o *V1Instance) GetLogicSetOk() (*string, bool)`

GetLogicSetOk returns a tuple with the LogicSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicSet

`func (o *V1Instance) SetLogicSet(v string)`

SetLogicSet sets LogicSet field to given value.

### HasLogicSet

`func (o *V1Instance) HasLogicSet() bool`

HasLogicSet returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Instance) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Instance) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Instance) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Instance) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMtime

`func (o *V1Instance) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1Instance) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1Instance) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1Instance) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetNamespace

`func (o *V1Instance) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1Instance) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1Instance) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1Instance) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPort

`func (o *V1Instance) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1Instance) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1Instance) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V1Instance) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPriority

`func (o *V1Instance) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V1Instance) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V1Instance) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V1Instance) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *V1Instance) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1Instance) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1Instance) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1Instance) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRevision

`func (o *V1Instance) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1Instance) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1Instance) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1Instance) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetService

`func (o *V1Instance) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1Instance) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1Instance) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1Instance) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceToken

`func (o *V1Instance) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *V1Instance) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *V1Instance) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *V1Instance) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetVersion

`func (o *V1Instance) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1Instance) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1Instance) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1Instance) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVpcId

`func (o *V1Instance) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *V1Instance) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *V1Instance) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *V1Instance) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetWeight

`func (o *V1Instance) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V1Instance) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V1Instance) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V1Instance) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


