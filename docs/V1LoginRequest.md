# V1LoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewV1LoginRequest

`func NewV1LoginRequest() *V1LoginRequest`

NewV1LoginRequest instantiates a new V1LoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LoginRequestWithDefaults

`func NewV1LoginRequestWithDefaults() *V1LoginRequest`

NewV1LoginRequestWithDefaults instantiates a new V1LoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1LoginRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1LoginRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1LoginRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1LoginRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *V1LoginRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1LoginRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1LoginRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1LoginRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPassword

`func (o *V1LoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V1LoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V1LoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V1LoginRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


