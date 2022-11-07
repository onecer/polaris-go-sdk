# V1CbPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consecutive** | Pointer to [**V1CbPolicyConsecutiveErrConfig**](V1CbPolicyConsecutiveErrConfig.md) |  | [optional] 
**ErrorRate** | Pointer to [**V1CbPolicyErrRateConfig**](V1CbPolicyErrRateConfig.md) |  | [optional] 
**JudgeDuration** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 
**MaxEjectionPercent** | Pointer to **int32** |  | [optional] 
**SlowRate** | Pointer to [**V1CbPolicySlowRateConfig**](V1CbPolicySlowRateConfig.md) |  | [optional] 

## Methods

### NewV1CbPolicy

`func NewV1CbPolicy() *V1CbPolicy`

NewV1CbPolicy instantiates a new V1CbPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CbPolicyWithDefaults

`func NewV1CbPolicyWithDefaults() *V1CbPolicy`

NewV1CbPolicyWithDefaults instantiates a new V1CbPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsecutive

`func (o *V1CbPolicy) GetConsecutive() V1CbPolicyConsecutiveErrConfig`

GetConsecutive returns the Consecutive field if non-nil, zero value otherwise.

### GetConsecutiveOk

`func (o *V1CbPolicy) GetConsecutiveOk() (*V1CbPolicyConsecutiveErrConfig, bool)`

GetConsecutiveOk returns a tuple with the Consecutive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutive

`func (o *V1CbPolicy) SetConsecutive(v V1CbPolicyConsecutiveErrConfig)`

SetConsecutive sets Consecutive field to given value.

### HasConsecutive

`func (o *V1CbPolicy) HasConsecutive() bool`

HasConsecutive returns a boolean if a field has been set.

### GetErrorRate

`func (o *V1CbPolicy) GetErrorRate() V1CbPolicyErrRateConfig`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *V1CbPolicy) GetErrorRateOk() (*V1CbPolicyErrRateConfig, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *V1CbPolicy) SetErrorRate(v V1CbPolicyErrRateConfig)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *V1CbPolicy) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.

### GetJudgeDuration

`func (o *V1CbPolicy) GetJudgeDuration() DurationpbDuration`

GetJudgeDuration returns the JudgeDuration field if non-nil, zero value otherwise.

### GetJudgeDurationOk

`func (o *V1CbPolicy) GetJudgeDurationOk() (*DurationpbDuration, bool)`

GetJudgeDurationOk returns a tuple with the JudgeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeDuration

`func (o *V1CbPolicy) SetJudgeDuration(v DurationpbDuration)`

SetJudgeDuration sets JudgeDuration field to given value.

### HasJudgeDuration

`func (o *V1CbPolicy) HasJudgeDuration() bool`

HasJudgeDuration returns a boolean if a field has been set.

### GetMaxEjectionPercent

`func (o *V1CbPolicy) GetMaxEjectionPercent() int32`

GetMaxEjectionPercent returns the MaxEjectionPercent field if non-nil, zero value otherwise.

### GetMaxEjectionPercentOk

`func (o *V1CbPolicy) GetMaxEjectionPercentOk() (*int32, bool)`

GetMaxEjectionPercentOk returns a tuple with the MaxEjectionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEjectionPercent

`func (o *V1CbPolicy) SetMaxEjectionPercent(v int32)`

SetMaxEjectionPercent sets MaxEjectionPercent field to given value.

### HasMaxEjectionPercent

`func (o *V1CbPolicy) HasMaxEjectionPercent() bool`

HasMaxEjectionPercent returns a boolean if a field has been set.

### GetSlowRate

`func (o *V1CbPolicy) GetSlowRate() V1CbPolicySlowRateConfig`

GetSlowRate returns the SlowRate field if non-nil, zero value otherwise.

### GetSlowRateOk

`func (o *V1CbPolicy) GetSlowRateOk() (*V1CbPolicySlowRateConfig, bool)`

GetSlowRateOk returns a tuple with the SlowRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowRate

`func (o *V1CbPolicy) SetSlowRate(v V1CbPolicySlowRateConfig)`

SetSlowRate sets SlowRate field to given value.

### HasSlowRate

`func (o *V1CbPolicy) HasSlowRate() bool`

HasSlowRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


