# V1ConfigRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitBreaker** | Pointer to [**V1CircuitBreaker**](V1CircuitBreaker.md) |  | [optional] 
**Ctime** | Pointer to **string** |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Service** | Pointer to [**V1Service**](V1Service.md) |  | [optional] 

## Methods

### NewV1ConfigRelease

`func NewV1ConfigRelease() *V1ConfigRelease`

NewV1ConfigRelease instantiates a new V1ConfigRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigReleaseWithDefaults

`func NewV1ConfigReleaseWithDefaults() *V1ConfigRelease`

NewV1ConfigReleaseWithDefaults instantiates a new V1ConfigRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitBreaker

`func (o *V1ConfigRelease) GetCircuitBreaker() V1CircuitBreaker`

GetCircuitBreaker returns the CircuitBreaker field if non-nil, zero value otherwise.

### GetCircuitBreakerOk

`func (o *V1ConfigRelease) GetCircuitBreakerOk() (*V1CircuitBreaker, bool)`

GetCircuitBreakerOk returns a tuple with the CircuitBreaker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreaker

`func (o *V1ConfigRelease) SetCircuitBreaker(v V1CircuitBreaker)`

SetCircuitBreaker sets CircuitBreaker field to given value.

### HasCircuitBreaker

`func (o *V1ConfigRelease) HasCircuitBreaker() bool`

HasCircuitBreaker returns a boolean if a field has been set.

### GetCtime

`func (o *V1ConfigRelease) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1ConfigRelease) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1ConfigRelease) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1ConfigRelease) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetMtime

`func (o *V1ConfigRelease) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1ConfigRelease) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1ConfigRelease) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1ConfigRelease) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetService

`func (o *V1ConfigRelease) GetService() V1Service`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1ConfigRelease) GetServiceOk() (*V1Service, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1ConfigRelease) SetService(v V1Service)`

SetService sets Service field to given value.

### HasService

`func (o *V1ConfigRelease) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


