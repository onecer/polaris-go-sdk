# V1ClimbConfigMetricConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Precision** | Pointer to **int32** |  | [optional] 
**ReportInterval** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 
**Window** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 

## Methods

### NewV1ClimbConfigMetricConfig

`func NewV1ClimbConfigMetricConfig() *V1ClimbConfigMetricConfig`

NewV1ClimbConfigMetricConfig instantiates a new V1ClimbConfigMetricConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClimbConfigMetricConfigWithDefaults

`func NewV1ClimbConfigMetricConfigWithDefaults() *V1ClimbConfigMetricConfig`

NewV1ClimbConfigMetricConfigWithDefaults instantiates a new V1ClimbConfigMetricConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrecision

`func (o *V1ClimbConfigMetricConfig) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *V1ClimbConfigMetricConfig) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *V1ClimbConfigMetricConfig) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *V1ClimbConfigMetricConfig) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetReportInterval

`func (o *V1ClimbConfigMetricConfig) GetReportInterval() DurationpbDuration`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *V1ClimbConfigMetricConfig) GetReportIntervalOk() (*DurationpbDuration, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *V1ClimbConfigMetricConfig) SetReportInterval(v DurationpbDuration)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *V1ClimbConfigMetricConfig) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetWindow

`func (o *V1ClimbConfigMetricConfig) GetWindow() DurationpbDuration`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *V1ClimbConfigMetricConfig) GetWindowOk() (*DurationpbDuration, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *V1ClimbConfigMetricConfig) SetWindow(v DurationpbDuration)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *V1ClimbConfigMetricConfig) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


