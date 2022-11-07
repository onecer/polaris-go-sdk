# V1DestinationSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCodes** | Pointer to [**[]WrapperspbInt64Value**](WrapperspbInt64Value.md) |  | [optional] 
**Metadata** | Pointer to [**map[string]V1MatchString**](V1MatchString.md) |  | [optional] 
**Method** | Pointer to [**V1MatchString**](V1MatchString.md) |  | [optional] 
**MetricPrecision** | Pointer to **int32** |  | [optional] 
**MetricWindow** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Policy** | Pointer to [**V1CbPolicy**](V1CbPolicy.md) |  | [optional] 
**Recover** | Pointer to [**V1RecoverConfig**](V1RecoverConfig.md) |  | [optional] 
**Resource** | Pointer to **int32** |  | [optional] 
**Scope** | Pointer to **int32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**UpdateInterval** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 

## Methods

### NewV1DestinationSet

`func NewV1DestinationSet() *V1DestinationSet`

NewV1DestinationSet instantiates a new V1DestinationSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DestinationSetWithDefaults

`func NewV1DestinationSetWithDefaults() *V1DestinationSet`

NewV1DestinationSetWithDefaults instantiates a new V1DestinationSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCodes

`func (o *V1DestinationSet) GetErrorCodes() []WrapperspbInt64Value`

GetErrorCodes returns the ErrorCodes field if non-nil, zero value otherwise.

### GetErrorCodesOk

`func (o *V1DestinationSet) GetErrorCodesOk() (*[]WrapperspbInt64Value, bool)`

GetErrorCodesOk returns a tuple with the ErrorCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCodes

`func (o *V1DestinationSet) SetErrorCodes(v []WrapperspbInt64Value)`

SetErrorCodes sets ErrorCodes field to given value.

### HasErrorCodes

`func (o *V1DestinationSet) HasErrorCodes() bool`

HasErrorCodes returns a boolean if a field has been set.

### GetMetadata

`func (o *V1DestinationSet) GetMetadata() map[string]V1MatchString`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1DestinationSet) GetMetadataOk() (*map[string]V1MatchString, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1DestinationSet) SetMetadata(v map[string]V1MatchString)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1DestinationSet) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMethod

`func (o *V1DestinationSet) GetMethod() V1MatchString`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V1DestinationSet) GetMethodOk() (*V1MatchString, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V1DestinationSet) SetMethod(v V1MatchString)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V1DestinationSet) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMetricPrecision

`func (o *V1DestinationSet) GetMetricPrecision() int32`

GetMetricPrecision returns the MetricPrecision field if non-nil, zero value otherwise.

### GetMetricPrecisionOk

`func (o *V1DestinationSet) GetMetricPrecisionOk() (*int32, bool)`

GetMetricPrecisionOk returns a tuple with the MetricPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricPrecision

`func (o *V1DestinationSet) SetMetricPrecision(v int32)`

SetMetricPrecision sets MetricPrecision field to given value.

### HasMetricPrecision

`func (o *V1DestinationSet) HasMetricPrecision() bool`

HasMetricPrecision returns a boolean if a field has been set.

### GetMetricWindow

`func (o *V1DestinationSet) GetMetricWindow() DurationpbDuration`

GetMetricWindow returns the MetricWindow field if non-nil, zero value otherwise.

### GetMetricWindowOk

`func (o *V1DestinationSet) GetMetricWindowOk() (*DurationpbDuration, bool)`

GetMetricWindowOk returns a tuple with the MetricWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricWindow

`func (o *V1DestinationSet) SetMetricWindow(v DurationpbDuration)`

SetMetricWindow sets MetricWindow field to given value.

### HasMetricWindow

`func (o *V1DestinationSet) HasMetricWindow() bool`

HasMetricWindow returns a boolean if a field has been set.

### GetNamespace

`func (o *V1DestinationSet) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1DestinationSet) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1DestinationSet) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1DestinationSet) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPolicy

`func (o *V1DestinationSet) GetPolicy() V1CbPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1DestinationSet) GetPolicyOk() (*V1CbPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1DestinationSet) SetPolicy(v V1CbPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *V1DestinationSet) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRecover

`func (o *V1DestinationSet) GetRecover() V1RecoverConfig`

GetRecover returns the Recover field if non-nil, zero value otherwise.

### GetRecoverOk

`func (o *V1DestinationSet) GetRecoverOk() (*V1RecoverConfig, bool)`

GetRecoverOk returns a tuple with the Recover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecover

`func (o *V1DestinationSet) SetRecover(v V1RecoverConfig)`

SetRecover sets Recover field to given value.

### HasRecover

`func (o *V1DestinationSet) HasRecover() bool`

HasRecover returns a boolean if a field has been set.

### GetResource

`func (o *V1DestinationSet) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *V1DestinationSet) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *V1DestinationSet) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *V1DestinationSet) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetScope

`func (o *V1DestinationSet) GetScope() int32`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *V1DestinationSet) GetScopeOk() (*int32, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *V1DestinationSet) SetScope(v int32)`

SetScope sets Scope field to given value.

### HasScope

`func (o *V1DestinationSet) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetService

`func (o *V1DestinationSet) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1DestinationSet) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1DestinationSet) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1DestinationSet) HasService() bool`

HasService returns a boolean if a field has been set.

### GetType

`func (o *V1DestinationSet) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1DestinationSet) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1DestinationSet) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *V1DestinationSet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateInterval

`func (o *V1DestinationSet) GetUpdateInterval() DurationpbDuration`

GetUpdateInterval returns the UpdateInterval field if non-nil, zero value otherwise.

### GetUpdateIntervalOk

`func (o *V1DestinationSet) GetUpdateIntervalOk() (*DurationpbDuration, bool)`

GetUpdateIntervalOk returns a tuple with the UpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInterval

`func (o *V1DestinationSet) SetUpdateInterval(v DurationpbDuration)`

SetUpdateInterval sets UpdateInterval field to given value.

### HasUpdateInterval

`func (o *V1DestinationSet) HasUpdateInterval() bool`

HasUpdateInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


