# V1CbPolicyErrRateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**ErrorRateToOpen** | Pointer to **int32** |  | [optional] 
**ErrorRateToPreserved** | Pointer to **int32** |  | [optional] 
**RequestVolumeThreshold** | Pointer to **int32** |  | [optional] 
**Specials** | Pointer to [**[]V1CbPolicyErrRateConfigSpecialConfig**](V1CbPolicyErrRateConfigSpecialConfig.md) |  | [optional] 

## Methods

### NewV1CbPolicyErrRateConfig

`func NewV1CbPolicyErrRateConfig() *V1CbPolicyErrRateConfig`

NewV1CbPolicyErrRateConfig instantiates a new V1CbPolicyErrRateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CbPolicyErrRateConfigWithDefaults

`func NewV1CbPolicyErrRateConfigWithDefaults() *V1CbPolicyErrRateConfig`

NewV1CbPolicyErrRateConfigWithDefaults instantiates a new V1CbPolicyErrRateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *V1CbPolicyErrRateConfig) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *V1CbPolicyErrRateConfig) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *V1CbPolicyErrRateConfig) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *V1CbPolicyErrRateConfig) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetErrorRateToOpen

`func (o *V1CbPolicyErrRateConfig) GetErrorRateToOpen() int32`

GetErrorRateToOpen returns the ErrorRateToOpen field if non-nil, zero value otherwise.

### GetErrorRateToOpenOk

`func (o *V1CbPolicyErrRateConfig) GetErrorRateToOpenOk() (*int32, bool)`

GetErrorRateToOpenOk returns a tuple with the ErrorRateToOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRateToOpen

`func (o *V1CbPolicyErrRateConfig) SetErrorRateToOpen(v int32)`

SetErrorRateToOpen sets ErrorRateToOpen field to given value.

### HasErrorRateToOpen

`func (o *V1CbPolicyErrRateConfig) HasErrorRateToOpen() bool`

HasErrorRateToOpen returns a boolean if a field has been set.

### GetErrorRateToPreserved

`func (o *V1CbPolicyErrRateConfig) GetErrorRateToPreserved() int32`

GetErrorRateToPreserved returns the ErrorRateToPreserved field if non-nil, zero value otherwise.

### GetErrorRateToPreservedOk

`func (o *V1CbPolicyErrRateConfig) GetErrorRateToPreservedOk() (*int32, bool)`

GetErrorRateToPreservedOk returns a tuple with the ErrorRateToPreserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRateToPreserved

`func (o *V1CbPolicyErrRateConfig) SetErrorRateToPreserved(v int32)`

SetErrorRateToPreserved sets ErrorRateToPreserved field to given value.

### HasErrorRateToPreserved

`func (o *V1CbPolicyErrRateConfig) HasErrorRateToPreserved() bool`

HasErrorRateToPreserved returns a boolean if a field has been set.

### GetRequestVolumeThreshold

`func (o *V1CbPolicyErrRateConfig) GetRequestVolumeThreshold() int32`

GetRequestVolumeThreshold returns the RequestVolumeThreshold field if non-nil, zero value otherwise.

### GetRequestVolumeThresholdOk

`func (o *V1CbPolicyErrRateConfig) GetRequestVolumeThresholdOk() (*int32, bool)`

GetRequestVolumeThresholdOk returns a tuple with the RequestVolumeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestVolumeThreshold

`func (o *V1CbPolicyErrRateConfig) SetRequestVolumeThreshold(v int32)`

SetRequestVolumeThreshold sets RequestVolumeThreshold field to given value.

### HasRequestVolumeThreshold

`func (o *V1CbPolicyErrRateConfig) HasRequestVolumeThreshold() bool`

HasRequestVolumeThreshold returns a boolean if a field has been set.

### GetSpecials

`func (o *V1CbPolicyErrRateConfig) GetSpecials() []V1CbPolicyErrRateConfigSpecialConfig`

GetSpecials returns the Specials field if non-nil, zero value otherwise.

### GetSpecialsOk

`func (o *V1CbPolicyErrRateConfig) GetSpecialsOk() (*[]V1CbPolicyErrRateConfigSpecialConfig, bool)`

GetSpecialsOk returns a tuple with the Specials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecials

`func (o *V1CbPolicyErrRateConfig) SetSpecials(v []V1CbPolicyErrRateConfigSpecialConfig)`

SetSpecials sets Specials field to given value.

### HasSpecials

`func (o *V1CbPolicyErrRateConfig) HasSpecials() bool`

HasSpecials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


