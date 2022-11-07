# V1Route

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | Pointer to [**[]V1Destination**](V1Destination.md) |  | [optional] 
**ExtendInfo** | Pointer to **map[string]string** |  | [optional] 
**Sources** | Pointer to [**[]V1Source**](V1Source.md) |  | [optional] 

## Methods

### NewV1Route

`func NewV1Route() *V1Route`

NewV1Route instantiates a new V1Route object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RouteWithDefaults

`func NewV1RouteWithDefaults() *V1Route`

NewV1RouteWithDefaults instantiates a new V1Route object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *V1Route) GetDestinations() []V1Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *V1Route) GetDestinationsOk() (*[]V1Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *V1Route) SetDestinations(v []V1Destination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *V1Route) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetExtendInfo

`func (o *V1Route) GetExtendInfo() map[string]string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *V1Route) GetExtendInfoOk() (*map[string]string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *V1Route) SetExtendInfo(v map[string]string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *V1Route) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetSources

`func (o *V1Route) GetSources() []V1Source`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *V1Route) GetSourcesOk() (*[]V1Source, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *V1Route) SetSources(v []V1Source)`

SetSources sets Sources field to given value.

### HasSources

`func (o *V1Route) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


