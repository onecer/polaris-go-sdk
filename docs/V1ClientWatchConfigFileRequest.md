# V1ClientWatchConfigFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIp** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**WatchFiles** | Pointer to [**[]V1ClientConfigFileInfo**](V1ClientConfigFileInfo.md) |  | [optional] 

## Methods

### NewV1ClientWatchConfigFileRequest

`func NewV1ClientWatchConfigFileRequest() *V1ClientWatchConfigFileRequest`

NewV1ClientWatchConfigFileRequest instantiates a new V1ClientWatchConfigFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClientWatchConfigFileRequestWithDefaults

`func NewV1ClientWatchConfigFileRequestWithDefaults() *V1ClientWatchConfigFileRequest`

NewV1ClientWatchConfigFileRequestWithDefaults instantiates a new V1ClientWatchConfigFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIp

`func (o *V1ClientWatchConfigFileRequest) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *V1ClientWatchConfigFileRequest) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *V1ClientWatchConfigFileRequest) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *V1ClientWatchConfigFileRequest) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetServiceName

`func (o *V1ClientWatchConfigFileRequest) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *V1ClientWatchConfigFileRequest) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *V1ClientWatchConfigFileRequest) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *V1ClientWatchConfigFileRequest) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetWatchFiles

`func (o *V1ClientWatchConfigFileRequest) GetWatchFiles() []V1ClientConfigFileInfo`

GetWatchFiles returns the WatchFiles field if non-nil, zero value otherwise.

### GetWatchFilesOk

`func (o *V1ClientWatchConfigFileRequest) GetWatchFilesOk() (*[]V1ClientConfigFileInfo, bool)`

GetWatchFilesOk returns a tuple with the WatchFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchFiles

`func (o *V1ClientWatchConfigFileRequest) SetWatchFiles(v []V1ClientConfigFileInfo)`

SetWatchFiles sets WatchFiles field to given value.

### HasWatchFiles

`func (o *V1ClientWatchConfigFileRequest) HasWatchFiles() bool`

HasWatchFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


