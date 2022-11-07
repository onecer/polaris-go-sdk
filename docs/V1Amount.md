# V1Amount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAmount** | Pointer to **int32** |  | [optional] 
**MinAmount** | Pointer to **int32** |  | [optional] 
**Precision** | Pointer to **int32** |  | [optional] 
**StartAmount** | Pointer to **int32** |  | [optional] 
**ValidDuration** | Pointer to [**DurationpbDuration**](DurationpbDuration.md) |  | [optional] 

## Methods

### NewV1Amount

`func NewV1Amount() *V1Amount`

NewV1Amount instantiates a new V1Amount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AmountWithDefaults

`func NewV1AmountWithDefaults() *V1Amount`

NewV1AmountWithDefaults instantiates a new V1Amount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAmount

`func (o *V1Amount) GetMaxAmount() int32`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *V1Amount) GetMaxAmountOk() (*int32, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *V1Amount) SetMaxAmount(v int32)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *V1Amount) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetMinAmount

`func (o *V1Amount) GetMinAmount() int32`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *V1Amount) GetMinAmountOk() (*int32, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *V1Amount) SetMinAmount(v int32)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *V1Amount) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetPrecision

`func (o *V1Amount) GetPrecision() int32`

GetPrecision returns the Precision field if non-nil, zero value otherwise.

### GetPrecisionOk

`func (o *V1Amount) GetPrecisionOk() (*int32, bool)`

GetPrecisionOk returns a tuple with the Precision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecision

`func (o *V1Amount) SetPrecision(v int32)`

SetPrecision sets Precision field to given value.

### HasPrecision

`func (o *V1Amount) HasPrecision() bool`

HasPrecision returns a boolean if a field has been set.

### GetStartAmount

`func (o *V1Amount) GetStartAmount() int32`

GetStartAmount returns the StartAmount field if non-nil, zero value otherwise.

### GetStartAmountOk

`func (o *V1Amount) GetStartAmountOk() (*int32, bool)`

GetStartAmountOk returns a tuple with the StartAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAmount

`func (o *V1Amount) SetStartAmount(v int32)`

SetStartAmount sets StartAmount field to given value.

### HasStartAmount

`func (o *V1Amount) HasStartAmount() bool`

HasStartAmount returns a boolean if a field has been set.

### GetValidDuration

`func (o *V1Amount) GetValidDuration() DurationpbDuration`

GetValidDuration returns the ValidDuration field if non-nil, zero value otherwise.

### GetValidDurationOk

`func (o *V1Amount) GetValidDurationOk() (*DurationpbDuration, bool)`

GetValidDurationOk returns a tuple with the ValidDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidDuration

`func (o *V1Amount) SetValidDuration(v DurationpbDuration)`

SetValidDuration sets ValidDuration field to given value.

### HasValidDuration

`func (o *V1Amount) HasValidDuration() bool`

HasValidDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


