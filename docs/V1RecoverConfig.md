# V1RecoverConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxRetryAfterHalfOpen** | Pointer to **int32** |  | [optional] 
**OutlierDetectWhen** | Pointer to **int32** |  | [optional] 
**RequestCountAfterHalfOpen** | Pointer to **int32** |  | [optional] 
**RequestRateAfterHalfOpen** | Pointer to **[]int32** |  | [optional] 
**SleepWindow** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 
**SuccessRateToClose** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1RecoverConfig

`func NewV1RecoverConfig() *V1RecoverConfig`

NewV1RecoverConfig instantiates a new V1RecoverConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RecoverConfigWithDefaults

`func NewV1RecoverConfigWithDefaults() *V1RecoverConfig`

NewV1RecoverConfigWithDefaults instantiates a new V1RecoverConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxRetryAfterHalfOpen

`func (o *V1RecoverConfig) GetMaxRetryAfterHalfOpen() int32`

GetMaxRetryAfterHalfOpen returns the MaxRetryAfterHalfOpen field if non-nil, zero value otherwise.

### GetMaxRetryAfterHalfOpenOk

`func (o *V1RecoverConfig) GetMaxRetryAfterHalfOpenOk() (*int32, bool)`

GetMaxRetryAfterHalfOpenOk returns a tuple with the MaxRetryAfterHalfOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetryAfterHalfOpen

`func (o *V1RecoverConfig) SetMaxRetryAfterHalfOpen(v int32)`

SetMaxRetryAfterHalfOpen sets MaxRetryAfterHalfOpen field to given value.

### HasMaxRetryAfterHalfOpen

`func (o *V1RecoverConfig) HasMaxRetryAfterHalfOpen() bool`

HasMaxRetryAfterHalfOpen returns a boolean if a field has been set.

### GetOutlierDetectWhen

`func (o *V1RecoverConfig) GetOutlierDetectWhen() int32`

GetOutlierDetectWhen returns the OutlierDetectWhen field if non-nil, zero value otherwise.

### GetOutlierDetectWhenOk

`func (o *V1RecoverConfig) GetOutlierDetectWhenOk() (*int32, bool)`

GetOutlierDetectWhenOk returns a tuple with the OutlierDetectWhen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutlierDetectWhen

`func (o *V1RecoverConfig) SetOutlierDetectWhen(v int32)`

SetOutlierDetectWhen sets OutlierDetectWhen field to given value.

### HasOutlierDetectWhen

`func (o *V1RecoverConfig) HasOutlierDetectWhen() bool`

HasOutlierDetectWhen returns a boolean if a field has been set.

### GetRequestCountAfterHalfOpen

`func (o *V1RecoverConfig) GetRequestCountAfterHalfOpen() int32`

GetRequestCountAfterHalfOpen returns the RequestCountAfterHalfOpen field if non-nil, zero value otherwise.

### GetRequestCountAfterHalfOpenOk

`func (o *V1RecoverConfig) GetRequestCountAfterHalfOpenOk() (*int32, bool)`

GetRequestCountAfterHalfOpenOk returns a tuple with the RequestCountAfterHalfOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCountAfterHalfOpen

`func (o *V1RecoverConfig) SetRequestCountAfterHalfOpen(v int32)`

SetRequestCountAfterHalfOpen sets RequestCountAfterHalfOpen field to given value.

### HasRequestCountAfterHalfOpen

`func (o *V1RecoverConfig) HasRequestCountAfterHalfOpen() bool`

HasRequestCountAfterHalfOpen returns a boolean if a field has been set.

### GetRequestRateAfterHalfOpen

`func (o *V1RecoverConfig) GetRequestRateAfterHalfOpen() []int32`

GetRequestRateAfterHalfOpen returns the RequestRateAfterHalfOpen field if non-nil, zero value otherwise.

### GetRequestRateAfterHalfOpenOk

`func (o *V1RecoverConfig) GetRequestRateAfterHalfOpenOk() (*[]int32, bool)`

GetRequestRateAfterHalfOpenOk returns a tuple with the RequestRateAfterHalfOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRateAfterHalfOpen

`func (o *V1RecoverConfig) SetRequestRateAfterHalfOpen(v []int32)`

SetRequestRateAfterHalfOpen sets RequestRateAfterHalfOpen field to given value.

### HasRequestRateAfterHalfOpen

`func (o *V1RecoverConfig) HasRequestRateAfterHalfOpen() bool`

HasRequestRateAfterHalfOpen returns a boolean if a field has been set.

### GetSleepWindow

`func (o *V1RecoverConfig) GetSleepWindow() DurationpbDuration`

GetSleepWindow returns the SleepWindow field if non-nil, zero value otherwise.

### GetSleepWindowOk

`func (o *V1RecoverConfig) GetSleepWindowOk() (*DurationpbDuration, bool)`

GetSleepWindowOk returns a tuple with the SleepWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepWindow

`func (o *V1RecoverConfig) SetSleepWindow(v DurationpbDuration)`

SetSleepWindow sets SleepWindow field to given value.

### HasSleepWindow

`func (o *V1RecoverConfig) HasSleepWindow() bool`

HasSleepWindow returns a boolean if a field has been set.

### GetSuccessRateToClose

`func (o *V1RecoverConfig) GetSuccessRateToClose() int32`

GetSuccessRateToClose returns the SuccessRateToClose field if non-nil, zero value otherwise.

### GetSuccessRateToCloseOk

`func (o *V1RecoverConfig) GetSuccessRateToCloseOk() (*int32, bool)`

GetSuccessRateToCloseOk returns a tuple with the SuccessRateToClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRateToClose

`func (o *V1RecoverConfig) SetSuccessRateToClose(v int32)`

SetSuccessRateToClose sets SuccessRateToClose field to given value.

### HasSuccessRateToClose

`func (o *V1RecoverConfig) HasSuccessRateToClose() bool`

HasSuccessRateToClose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


