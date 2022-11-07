# V1ClimbConfigClimbThrottling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ColdAboveTuneDownRate** | Pointer to [**WrapperspbInt32Value**](WrapperspbInt32Value.md) |  | [optional] 
**ColdAboveTuneUpRate** | Pointer to [**WrapperspbInt32Value**](WrapperspbInt32Value.md) |  | [optional] 
**ColdBelowTuneDownRate** | Pointer to [**WrapperspbInt32Value**](WrapperspbInt32Value.md) |  | [optional] 
**ColdBelowTuneUpRate** | Pointer to [**WrapperspbInt32Value**](WrapperspbInt32Value.md) |  | [optional] 
**JudgeDuration** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 
**LimitThresholdToTuneUp** | Pointer to [**WrapperspbInt32Value**](WrapperspbInt32Value.md) |  | [optional] 
**TuneDownPeriod** | Pointer to [**WrapperspbInt32Value**](WrapperspbInt32Value.md) |  | [optional] 
**TuneUpPeriod** | Pointer to [**WrapperspbInt32Value**](WrapperspbInt32Value.md) |  | [optional] 

## Methods

### NewV1ClimbConfigClimbThrottling

`func NewV1ClimbConfigClimbThrottling() *V1ClimbConfigClimbThrottling`

NewV1ClimbConfigClimbThrottling instantiates a new V1ClimbConfigClimbThrottling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClimbConfigClimbThrottlingWithDefaults

`func NewV1ClimbConfigClimbThrottlingWithDefaults() *V1ClimbConfigClimbThrottling`

NewV1ClimbConfigClimbThrottlingWithDefaults instantiates a new V1ClimbConfigClimbThrottling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColdAboveTuneDownRate

`func (o *V1ClimbConfigClimbThrottling) GetColdAboveTuneDownRate() WrapperspbInt32Value`

GetColdAboveTuneDownRate returns the ColdAboveTuneDownRate field if non-nil, zero value otherwise.

### GetColdAboveTuneDownRateOk

`func (o *V1ClimbConfigClimbThrottling) GetColdAboveTuneDownRateOk() (*WrapperspbInt32Value, bool)`

GetColdAboveTuneDownRateOk returns a tuple with the ColdAboveTuneDownRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdAboveTuneDownRate

`func (o *V1ClimbConfigClimbThrottling) SetColdAboveTuneDownRate(v WrapperspbInt32Value)`

SetColdAboveTuneDownRate sets ColdAboveTuneDownRate field to given value.

### HasColdAboveTuneDownRate

`func (o *V1ClimbConfigClimbThrottling) HasColdAboveTuneDownRate() bool`

HasColdAboveTuneDownRate returns a boolean if a field has been set.

### GetColdAboveTuneUpRate

`func (o *V1ClimbConfigClimbThrottling) GetColdAboveTuneUpRate() WrapperspbInt32Value`

GetColdAboveTuneUpRate returns the ColdAboveTuneUpRate field if non-nil, zero value otherwise.

### GetColdAboveTuneUpRateOk

`func (o *V1ClimbConfigClimbThrottling) GetColdAboveTuneUpRateOk() (*WrapperspbInt32Value, bool)`

GetColdAboveTuneUpRateOk returns a tuple with the ColdAboveTuneUpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdAboveTuneUpRate

`func (o *V1ClimbConfigClimbThrottling) SetColdAboveTuneUpRate(v WrapperspbInt32Value)`

SetColdAboveTuneUpRate sets ColdAboveTuneUpRate field to given value.

### HasColdAboveTuneUpRate

`func (o *V1ClimbConfigClimbThrottling) HasColdAboveTuneUpRate() bool`

HasColdAboveTuneUpRate returns a boolean if a field has been set.

### GetColdBelowTuneDownRate

`func (o *V1ClimbConfigClimbThrottling) GetColdBelowTuneDownRate() WrapperspbInt32Value`

GetColdBelowTuneDownRate returns the ColdBelowTuneDownRate field if non-nil, zero value otherwise.

### GetColdBelowTuneDownRateOk

`func (o *V1ClimbConfigClimbThrottling) GetColdBelowTuneDownRateOk() (*WrapperspbInt32Value, bool)`

GetColdBelowTuneDownRateOk returns a tuple with the ColdBelowTuneDownRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdBelowTuneDownRate

`func (o *V1ClimbConfigClimbThrottling) SetColdBelowTuneDownRate(v WrapperspbInt32Value)`

SetColdBelowTuneDownRate sets ColdBelowTuneDownRate field to given value.

### HasColdBelowTuneDownRate

`func (o *V1ClimbConfigClimbThrottling) HasColdBelowTuneDownRate() bool`

HasColdBelowTuneDownRate returns a boolean if a field has been set.

### GetColdBelowTuneUpRate

`func (o *V1ClimbConfigClimbThrottling) GetColdBelowTuneUpRate() WrapperspbInt32Value`

GetColdBelowTuneUpRate returns the ColdBelowTuneUpRate field if non-nil, zero value otherwise.

### GetColdBelowTuneUpRateOk

`func (o *V1ClimbConfigClimbThrottling) GetColdBelowTuneUpRateOk() (*WrapperspbInt32Value, bool)`

GetColdBelowTuneUpRateOk returns a tuple with the ColdBelowTuneUpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColdBelowTuneUpRate

`func (o *V1ClimbConfigClimbThrottling) SetColdBelowTuneUpRate(v WrapperspbInt32Value)`

SetColdBelowTuneUpRate sets ColdBelowTuneUpRate field to given value.

### HasColdBelowTuneUpRate

`func (o *V1ClimbConfigClimbThrottling) HasColdBelowTuneUpRate() bool`

HasColdBelowTuneUpRate returns a boolean if a field has been set.

### GetJudgeDuration

`func (o *V1ClimbConfigClimbThrottling) GetJudgeDuration() DurationpbDuration`

GetJudgeDuration returns the JudgeDuration field if non-nil, zero value otherwise.

### GetJudgeDurationOk

`func (o *V1ClimbConfigClimbThrottling) GetJudgeDurationOk() (*DurationpbDuration, bool)`

GetJudgeDurationOk returns a tuple with the JudgeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgeDuration

`func (o *V1ClimbConfigClimbThrottling) SetJudgeDuration(v DurationpbDuration)`

SetJudgeDuration sets JudgeDuration field to given value.

### HasJudgeDuration

`func (o *V1ClimbConfigClimbThrottling) HasJudgeDuration() bool`

HasJudgeDuration returns a boolean if a field has been set.

### GetLimitThresholdToTuneUp

`func (o *V1ClimbConfigClimbThrottling) GetLimitThresholdToTuneUp() WrapperspbInt32Value`

GetLimitThresholdToTuneUp returns the LimitThresholdToTuneUp field if non-nil, zero value otherwise.

### GetLimitThresholdToTuneUpOk

`func (o *V1ClimbConfigClimbThrottling) GetLimitThresholdToTuneUpOk() (*WrapperspbInt32Value, bool)`

GetLimitThresholdToTuneUpOk returns a tuple with the LimitThresholdToTuneUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitThresholdToTuneUp

`func (o *V1ClimbConfigClimbThrottling) SetLimitThresholdToTuneUp(v WrapperspbInt32Value)`

SetLimitThresholdToTuneUp sets LimitThresholdToTuneUp field to given value.

### HasLimitThresholdToTuneUp

`func (o *V1ClimbConfigClimbThrottling) HasLimitThresholdToTuneUp() bool`

HasLimitThresholdToTuneUp returns a boolean if a field has been set.

### GetTuneDownPeriod

`func (o *V1ClimbConfigClimbThrottling) GetTuneDownPeriod() WrapperspbInt32Value`

GetTuneDownPeriod returns the TuneDownPeriod field if non-nil, zero value otherwise.

### GetTuneDownPeriodOk

`func (o *V1ClimbConfigClimbThrottling) GetTuneDownPeriodOk() (*WrapperspbInt32Value, bool)`

GetTuneDownPeriodOk returns a tuple with the TuneDownPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuneDownPeriod

`func (o *V1ClimbConfigClimbThrottling) SetTuneDownPeriod(v WrapperspbInt32Value)`

SetTuneDownPeriod sets TuneDownPeriod field to given value.

### HasTuneDownPeriod

`func (o *V1ClimbConfigClimbThrottling) HasTuneDownPeriod() bool`

HasTuneDownPeriod returns a boolean if a field has been set.

### GetTuneUpPeriod

`func (o *V1ClimbConfigClimbThrottling) GetTuneUpPeriod() WrapperspbInt32Value`

GetTuneUpPeriod returns the TuneUpPeriod field if non-nil, zero value otherwise.

### GetTuneUpPeriodOk

`func (o *V1ClimbConfigClimbThrottling) GetTuneUpPeriodOk() (*WrapperspbInt32Value, bool)`

GetTuneUpPeriodOk returns a tuple with the TuneUpPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuneUpPeriod

`func (o *V1ClimbConfigClimbThrottling) SetTuneUpPeriod(v WrapperspbInt32Value)`

SetTuneUpPeriod sets TuneUpPeriod field to given value.

### HasTuneUpPeriod

`func (o *V1ClimbConfigClimbThrottling) HasTuneUpPeriod() bool`

HasTuneUpPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


