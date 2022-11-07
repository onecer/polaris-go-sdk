# V1CbPolicySlowRateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**MaxRt** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 
**SlowRateToOpen** | Pointer to **int32** |  | [optional] 
**SlowRateToPreserved** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1CbPolicySlowRateConfig

`func NewV1CbPolicySlowRateConfig() *V1CbPolicySlowRateConfig`

NewV1CbPolicySlowRateConfig instantiates a new V1CbPolicySlowRateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CbPolicySlowRateConfigWithDefaults

`func NewV1CbPolicySlowRateConfigWithDefaults() *V1CbPolicySlowRateConfig`

NewV1CbPolicySlowRateConfigWithDefaults instantiates a new V1CbPolicySlowRateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *V1CbPolicySlowRateConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *V1CbPolicySlowRateConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *V1CbPolicySlowRateConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *V1CbPolicySlowRateConfig) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMaxRt

`func (o *V1CbPolicySlowRateConfig) GetMaxRt() DurationpbDuration`

GetMaxRt returns the MaxRt field if non-nil, zero value otherwise.

### GetMaxRtOk

`func (o *V1CbPolicySlowRateConfig) GetMaxRtOk() (*DurationpbDuration, bool)`

GetMaxRtOk returns a tuple with the MaxRt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRt

`func (o *V1CbPolicySlowRateConfig) SetMaxRt(v DurationpbDuration)`

SetMaxRt sets MaxRt field to given value.

### HasMaxRt

`func (o *V1CbPolicySlowRateConfig) HasMaxRt() bool`

HasMaxRt returns a boolean if a field has been set.

### GetSlowRateToOpen

`func (o *V1CbPolicySlowRateConfig) GetSlowRateToOpen() int32`

GetSlowRateToOpen returns the SlowRateToOpen field if non-nil, zero value otherwise.

### GetSlowRateToOpenOk

`func (o *V1CbPolicySlowRateConfig) GetSlowRateToOpenOk() (*int32, bool)`

GetSlowRateToOpenOk returns a tuple with the SlowRateToOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowRateToOpen

`func (o *V1CbPolicySlowRateConfig) SetSlowRateToOpen(v int32)`

SetSlowRateToOpen sets SlowRateToOpen field to given value.

### HasSlowRateToOpen

`func (o *V1CbPolicySlowRateConfig) HasSlowRateToOpen() bool`

HasSlowRateToOpen returns a boolean if a field has been set.

### GetSlowRateToPreserved

`func (o *V1CbPolicySlowRateConfig) GetSlowRateToPreserved() int32`

GetSlowRateToPreserved returns the SlowRateToPreserved field if non-nil, zero value otherwise.

### GetSlowRateToPreservedOk

`func (o *V1CbPolicySlowRateConfig) GetSlowRateToPreservedOk() (*int32, bool)`

GetSlowRateToPreservedOk returns a tuple with the SlowRateToPreserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowRateToPreserved

`func (o *V1CbPolicySlowRateConfig) SetSlowRateToPreserved(v int32)`

SetSlowRateToPreserved sets SlowRateToPreserved field to given value.

### HasSlowRateToPreserved

`func (o *V1CbPolicySlowRateConfig) HasSlowRateToPreserved() bool`

HasSlowRateToPreserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


