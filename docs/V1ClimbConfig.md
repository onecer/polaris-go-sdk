# V1ClimbConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Metric** | Pointer to [**V1ClimbConfigMetricConfig**](V1ClimbConfigMetricConfig.md) |  | [optional] 
**Policy** | Pointer to [**V1ClimbConfigTriggerPolicy**](V1ClimbConfigTriggerPolicy.md) |  | [optional] 
**Throttling** | Pointer to [**V1ClimbConfigClimbThrottling**](V1ClimbConfigClimbThrottling.md) |  | [optional] 

## Methods

### NewV1ClimbConfig

`func NewV1ClimbConfig() *V1ClimbConfig`

NewV1ClimbConfig instantiates a new V1ClimbConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClimbConfigWithDefaults

`func NewV1ClimbConfigWithDefaults() *V1ClimbConfig`

NewV1ClimbConfigWithDefaults instantiates a new V1ClimbConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *V1ClimbConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *V1ClimbConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *V1ClimbConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *V1ClimbConfig) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMetric

`func (o *V1ClimbConfig) GetMetric() V1ClimbConfigMetricConfig`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *V1ClimbConfig) GetMetricOk() (*V1ClimbConfigMetricConfig, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *V1ClimbConfig) SetMetric(v V1ClimbConfigMetricConfig)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *V1ClimbConfig) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetPolicy

`func (o *V1ClimbConfig) GetPolicy() V1ClimbConfigTriggerPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *V1ClimbConfig) GetPolicyOk() (*V1ClimbConfigTriggerPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *V1ClimbConfig) SetPolicy(v V1ClimbConfigTriggerPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *V1ClimbConfig) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetThrottling

`func (o *V1ClimbConfig) GetThrottling() V1ClimbConfigClimbThrottling`

GetThrottling returns the Throttling field if non-nil, zero value otherwise.

### GetThrottlingOk

`func (o *V1ClimbConfig) GetThrottlingOk() (*V1ClimbConfigClimbThrottling, bool)`

GetThrottlingOk returns a tuple with the Throttling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottling

`func (o *V1ClimbConfig) SetThrottling(v V1ClimbConfigClimbThrottling)`

SetThrottling sets Throttling field to given value.

### HasThrottling

`func (o *V1ClimbConfig) HasThrottling() bool`

HasThrottling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


