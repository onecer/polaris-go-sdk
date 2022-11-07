# V1AuthStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **int32** |  | [optional] 
**AuthToken** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Ctime** | Pointer to **string** |  | [optional] 
**DefaultStrategy** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Principals** | Pointer to [**V1Principals**](V1Principals.md) |  | [optional] 
**Resources** | Pointer to [**V1StrategyResources**](V1StrategyResources.md) |  | [optional] 

## Methods

### NewV1AuthStrategy

`func NewV1AuthStrategy() *V1AuthStrategy`

NewV1AuthStrategy instantiates a new V1AuthStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AuthStrategyWithDefaults

`func NewV1AuthStrategyWithDefaults() *V1AuthStrategy`

NewV1AuthStrategyWithDefaults instantiates a new V1AuthStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *V1AuthStrategy) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V1AuthStrategy) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V1AuthStrategy) SetAction(v int32)`

SetAction sets Action field to given value.

### HasAction

`func (o *V1AuthStrategy) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAuthToken

`func (o *V1AuthStrategy) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *V1AuthStrategy) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *V1AuthStrategy) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *V1AuthStrategy) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetComment

`func (o *V1AuthStrategy) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1AuthStrategy) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1AuthStrategy) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1AuthStrategy) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCtime

`func (o *V1AuthStrategy) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1AuthStrategy) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1AuthStrategy) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1AuthStrategy) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetDefaultStrategy

`func (o *V1AuthStrategy) GetDefaultStrategy() bool`

GetDefaultStrategy returns the DefaultStrategy field if non-nil, zero value otherwise.

### GetDefaultStrategyOk

`func (o *V1AuthStrategy) GetDefaultStrategyOk() (*bool, bool)`

GetDefaultStrategyOk returns a tuple with the DefaultStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStrategy

`func (o *V1AuthStrategy) SetDefaultStrategy(v bool)`

SetDefaultStrategy sets DefaultStrategy field to given value.

### HasDefaultStrategy

`func (o *V1AuthStrategy) HasDefaultStrategy() bool`

HasDefaultStrategy returns a boolean if a field has been set.

### GetId

`func (o *V1AuthStrategy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1AuthStrategy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1AuthStrategy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1AuthStrategy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtime

`func (o *V1AuthStrategy) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1AuthStrategy) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1AuthStrategy) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1AuthStrategy) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetName

`func (o *V1AuthStrategy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1AuthStrategy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1AuthStrategy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1AuthStrategy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *V1AuthStrategy) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1AuthStrategy) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1AuthStrategy) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1AuthStrategy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrincipals

`func (o *V1AuthStrategy) GetPrincipals() V1Principals`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *V1AuthStrategy) GetPrincipalsOk() (*V1Principals, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *V1AuthStrategy) SetPrincipals(v V1Principals)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *V1AuthStrategy) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetResources

`func (o *V1AuthStrategy) GetResources() V1StrategyResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1AuthStrategy) GetResourcesOk() (*V1StrategyResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1AuthStrategy) SetResources(v V1StrategyResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1AuthStrategy) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


