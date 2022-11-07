# V1CbRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destinations** | Pointer to [**[]V1DestinationSet**](V1DestinationSet.md) |  | [optional] 
**Sources** | Pointer to [**[]V1SourceMatcher**](V1SourceMatcher.md) |  | [optional] 

## Methods

### NewV1CbRule

`func NewV1CbRule() *V1CbRule`

NewV1CbRule instantiates a new V1CbRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1CbRuleWithDefaults

`func NewV1CbRuleWithDefaults() *V1CbRule`

NewV1CbRuleWithDefaults instantiates a new V1CbRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinations

`func (o *V1CbRule) GetDestinations() []V1DestinationSet`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *V1CbRule) GetDestinationsOk() (*[]V1DestinationSet, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *V1CbRule) SetDestinations(v []V1DestinationSet)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *V1CbRule) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.

### GetSources

`func (o *V1CbRule) GetSources() []V1SourceMatcher`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *V1CbRule) GetSourcesOk() (*[]V1SourceMatcher, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *V1CbRule) SetSources(v []V1SourceMatcher)`

SetSources sets Sources field to given value.

### HasSources

`func (o *V1CbRule) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


