# V1Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**map[string]V1MatchString**](V1MatchString.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewV1Source

`func NewV1Source() *V1Source`

NewV1Source instantiates a new V1Source object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SourceWithDefaults

`func NewV1SourceWithDefaults() *V1Source`

NewV1SourceWithDefaults instantiates a new V1Source object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *V1Source) GetMetadata() map[string]V1MatchString`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Source) GetMetadataOk() (*map[string]V1MatchString, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Source) SetMetadata(v map[string]V1MatchString)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Source) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNamespace

`func (o *V1Source) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1Source) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1Source) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1Source) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetService

`func (o *V1Source) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1Source) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1Source) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1Source) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


