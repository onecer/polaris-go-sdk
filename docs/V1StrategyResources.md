# V1StrategyResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigGroups** | Pointer to [**[]V1StrategyResourceEntry**](V1StrategyResourceEntry.md) |  | [optional] 
**Namespaces** | Pointer to [**[]V1StrategyResourceEntry**](V1StrategyResourceEntry.md) |  | [optional] 
**Services** | Pointer to [**[]V1StrategyResourceEntry**](V1StrategyResourceEntry.md) |  | [optional] 
**StrategyId** | Pointer to **string** |  | [optional] 

## Methods

### NewV1StrategyResources

`func NewV1StrategyResources() *V1StrategyResources`

NewV1StrategyResources instantiates a new V1StrategyResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1StrategyResourcesWithDefaults

`func NewV1StrategyResourcesWithDefaults() *V1StrategyResources`

NewV1StrategyResourcesWithDefaults instantiates a new V1StrategyResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigGroups

`func (o *V1StrategyResources) GetConfigGroups() []V1StrategyResourceEntry`

GetConfigGroups returns the ConfigGroups field if non-nil, zero value otherwise.

### GetConfigGroupsOk

`func (o *V1StrategyResources) GetConfigGroupsOk() (*[]V1StrategyResourceEntry, bool)`

GetConfigGroupsOk returns a tuple with the ConfigGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigGroups

`func (o *V1StrategyResources) SetConfigGroups(v []V1StrategyResourceEntry)`

SetConfigGroups sets ConfigGroups field to given value.

### HasConfigGroups

`func (o *V1StrategyResources) HasConfigGroups() bool`

HasConfigGroups returns a boolean if a field has been set.

### GetNamespaces

`func (o *V1StrategyResources) GetNamespaces() []V1StrategyResourceEntry`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *V1StrategyResources) GetNamespacesOk() (*[]V1StrategyResourceEntry, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *V1StrategyResources) SetNamespaces(v []V1StrategyResourceEntry)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *V1StrategyResources) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### GetServices

`func (o *V1StrategyResources) GetServices() []V1StrategyResourceEntry`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *V1StrategyResources) GetServicesOk() (*[]V1StrategyResourceEntry, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *V1StrategyResources) SetServices(v []V1StrategyResourceEntry)`

SetServices sets Services field to given value.

### HasServices

`func (o *V1StrategyResources) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStrategyId

`func (o *V1StrategyResources) GetStrategyId() string`

GetStrategyId returns the StrategyId field if non-nil, zero value otherwise.

### GetStrategyIdOk

`func (o *V1StrategyResources) GetStrategyIdOk() (*string, bool)`

GetStrategyIdOk returns a tuple with the StrategyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyId

`func (o *V1StrategyResources) SetStrategyId(v string)`

SetStrategyId sets StrategyId field to given value.

### HasStrategyId

`func (o *V1StrategyResources) HasStrategyId() bool`

HasStrategyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


