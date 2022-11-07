# V1SourceMatcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**map[string]V1MatchString**](V1MatchString.md) |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewV1SourceMatcher

`func NewV1SourceMatcher() *V1SourceMatcher`

NewV1SourceMatcher instantiates a new V1SourceMatcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SourceMatcherWithDefaults

`func NewV1SourceMatcherWithDefaults() *V1SourceMatcher`

NewV1SourceMatcherWithDefaults instantiates a new V1SourceMatcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *V1SourceMatcher) GetLabels() map[string]V1MatchString`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1SourceMatcher) GetLabelsOk() (*map[string]V1MatchString, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1SourceMatcher) SetLabels(v map[string]V1MatchString)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1SourceMatcher) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNamespace

`func (o *V1SourceMatcher) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1SourceMatcher) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1SourceMatcher) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1SourceMatcher) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetService

`func (o *V1SourceMatcher) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1SourceMatcher) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1SourceMatcher) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1SourceMatcher) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


