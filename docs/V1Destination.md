# V1Destination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isolate** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to [**map[string]V1MatchString**](V1MatchString.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Transfer** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1Destination

`func NewV1Destination() *V1Destination`

NewV1Destination instantiates a new V1Destination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DestinationWithDefaults

`func NewV1DestinationWithDefaults() *V1Destination`

NewV1DestinationWithDefaults instantiates a new V1Destination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsolate

`func (o *V1Destination) GetIsolate() bool`

GetIsolate returns the Isolate field if non-nil, zero value otherwise.

### GetIsolateOk

`func (o *V1Destination) GetIsolateOk() (*bool, bool)`

GetIsolateOk returns a tuple with the Isolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolate

`func (o *V1Destination) SetIsolate(v bool)`

SetIsolate sets Isolate field to given value.

### HasIsolate

`func (o *V1Destination) HasIsolate() bool`

HasIsolate returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Destination) GetMetadata() map[string]V1MatchString`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Destination) GetMetadataOk() (*map[string]V1MatchString, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Destination) SetMetadata(v map[string]V1MatchString)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Destination) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNamespace

`func (o *V1Destination) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1Destination) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1Destination) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1Destination) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPriority

`func (o *V1Destination) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V1Destination) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V1Destination) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V1Destination) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetService

`func (o *V1Destination) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1Destination) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1Destination) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1Destination) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTransfer

`func (o *V1Destination) GetTransfer() string`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *V1Destination) GetTransferOk() (*string, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *V1Destination) SetTransfer(v string)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *V1Destination) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetWeight

`func (o *V1Destination) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V1Destination) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V1Destination) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V1Destination) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


