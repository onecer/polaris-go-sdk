# V1Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountPercent** | Pointer to **int32** |  | [optional] 
**Interval** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 

## Methods

### NewV1Report

`func NewV1Report() *V1Report`

NewV1Report instantiates a new V1Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ReportWithDefaults

`func NewV1ReportWithDefaults() *V1Report`

NewV1ReportWithDefaults instantiates a new V1Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountPercent

`func (o *V1Report) GetAmountPercent() int32`

GetAmountPercent returns the AmountPercent field if non-nil, zero value otherwise.

### GetAmountPercentOk

`func (o *V1Report) GetAmountPercentOk() (*int32, bool)`

GetAmountPercentOk returns a tuple with the AmountPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountPercent

`func (o *V1Report) SetAmountPercent(v int32)`

SetAmountPercent sets AmountPercent field to given value.

### HasAmountPercent

`func (o *V1Report) HasAmountPercent() bool`

HasAmountPercent returns a boolean if a field has been set.

### GetInterval

`func (o *V1Report) GetInterval() DurationpbDuration`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *V1Report) GetIntervalOk() (*DurationpbDuration, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *V1Report) SetInterval(v DurationpbDuration)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *V1Report) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


