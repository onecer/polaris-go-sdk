# V1HealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Heartbeat** | Pointer to [**V1HeartbeatHealthCheck**](V1HeartbeatHealthCheck.md) |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1HealthCheck

`func NewV1HealthCheck() *V1HealthCheck`

NewV1HealthCheck instantiates a new V1HealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1HealthCheckWithDefaults

`func NewV1HealthCheckWithDefaults() *V1HealthCheck`

NewV1HealthCheckWithDefaults instantiates a new V1HealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeartbeat

`func (o *V1HealthCheck) GetHeartbeat() V1HeartbeatHealthCheck`

GetHeartbeat returns the Heartbeat field if non-nil, zero value otherwise.

### GetHeartbeatOk

`func (o *V1HealthCheck) GetHeartbeatOk() (*V1HeartbeatHealthCheck, bool)`

GetHeartbeatOk returns a tuple with the Heartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeat

`func (o *V1HealthCheck) SetHeartbeat(v V1HeartbeatHealthCheck)`

SetHeartbeat sets Heartbeat field to given value.

### HasHeartbeat

`func (o *V1HealthCheck) HasHeartbeat() bool`

HasHeartbeat returns a boolean if a field has been set.

### GetType

`func (o *V1HealthCheck) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1HealthCheck) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1HealthCheck) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *V1HealthCheck) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


