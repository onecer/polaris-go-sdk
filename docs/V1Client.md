# V1Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ctime** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**V1Location**](V1Location.md) |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Stat** | Pointer to [**[]V1StatInfo**](V1StatInfo.md) |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewV1Client

`func NewV1Client() *V1Client`

NewV1Client instantiates a new V1Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClientWithDefaults

`func NewV1ClientWithDefaults() *V1Client`

NewV1ClientWithDefaults instantiates a new V1Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtime

`func (o *V1Client) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1Client) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1Client) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1Client) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetHost

`func (o *V1Client) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1Client) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1Client) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1Client) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *V1Client) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Client) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Client) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Client) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *V1Client) GetLocation() V1Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1Client) GetLocationOk() (*V1Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1Client) SetLocation(v V1Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1Client) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMtime

`func (o *V1Client) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1Client) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1Client) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1Client) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetStat

`func (o *V1Client) GetStat() []V1StatInfo`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *V1Client) GetStatOk() (*[]V1StatInfo, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *V1Client) SetStat(v []V1StatInfo)`

SetStat sets Stat field to given value.

### HasStat

`func (o *V1Client) HasStat() bool`

HasStat returns a boolean if a field has been set.

### GetType

`func (o *V1Client) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1Client) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1Client) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *V1Client) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *V1Client) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V1Client) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V1Client) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V1Client) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


