# V1MatchArgument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Value** | Pointer to [**V1MatchString**](V1MatchString.md) |  | [optional] 

## Methods

### NewV1MatchArgument

`func NewV1MatchArgument() *V1MatchArgument`

NewV1MatchArgument instantiates a new V1MatchArgument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MatchArgumentWithDefaults

`func NewV1MatchArgumentWithDefaults() *V1MatchArgument`

NewV1MatchArgumentWithDefaults instantiates a new V1MatchArgument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *V1MatchArgument) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V1MatchArgument) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V1MatchArgument) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *V1MatchArgument) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *V1MatchArgument) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1MatchArgument) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1MatchArgument) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *V1MatchArgument) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *V1MatchArgument) GetValue() V1MatchString`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1MatchArgument) GetValueOk() (*V1MatchString, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1MatchArgument) SetValue(v V1MatchString)`

SetValue sets Value field to given value.

### HasValue

`func (o *V1MatchArgument) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


