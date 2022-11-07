# V1Principals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]V1Principal**](V1Principal.md) |  | [optional] 
**Users** | Pointer to [**[]V1Principal**](V1Principal.md) |  | [optional] 

## Methods

### NewV1Principals

`func NewV1Principals() *V1Principals`

NewV1Principals instantiates a new V1Principals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PrincipalsWithDefaults

`func NewV1PrincipalsWithDefaults() *V1Principals`

NewV1PrincipalsWithDefaults instantiates a new V1Principals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *V1Principals) GetGroups() []V1Principal`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V1Principals) GetGroupsOk() (*[]V1Principal, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V1Principals) SetGroups(v []V1Principal)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V1Principals) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *V1Principals) GetUsers() []V1Principal`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V1Principals) GetUsersOk() (*[]V1Principal, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V1Principals) SetUsers(v []V1Principal)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *V1Principals) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


