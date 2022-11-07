# V1RateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | Pointer to **string** |  | [optional] 
**Rules** | Pointer to [**[]V1Rule**](V1Rule.md) |  | [optional] 

## Methods

### NewV1RateLimit

`func NewV1RateLimit() *V1RateLimit`

NewV1RateLimit instantiates a new V1RateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RateLimitWithDefaults

`func NewV1RateLimitWithDefaults() *V1RateLimit`

NewV1RateLimitWithDefaults instantiates a new V1RateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *V1RateLimit) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1RateLimit) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1RateLimit) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1RateLimit) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRules

`func (o *V1RateLimit) GetRules() []V1Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *V1RateLimit) GetRulesOk() (*[]V1Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *V1RateLimit) SetRules(v []V1Rule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *V1RateLimit) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


