# V1Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Adjuster** | Pointer to [**V1AmountAdjuster**](V1AmountAdjuster.md) |  | [optional] 
**AmountMode** | Pointer to **int32** |  | [optional] 
**Amounts** | Pointer to [**[]V1Amount**](V1Amount.md) |  | [optional] 
**Arguments** | Pointer to [**[]V1MatchArgument**](V1MatchArgument.md) |  | [optional] 
**Cluster** | Pointer to [**V1RateLimitCluster**](V1RateLimitCluster.md) |  | [optional] 
**Ctime** | Pointer to **string** |  | [optional] 
**Disable** | Pointer to **bool** |  | [optional] 
**Etime** | Pointer to **string** |  | [optional] 
**Failover** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to [**map[string]V1MatchString**](V1MatchString.md) |  | [optional] 
**MaxQueueDelay** | Pointer to **int32** |  | [optional] 
**Method** | Pointer to [**V1MatchString**](V1MatchString.md) |  | [optional] 
**Mtime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**RegexCombine** | Pointer to **bool** |  | [optional] 
**Report** | Pointer to [**V1Report**](V1Report.md) |  | [optional] 
**Resource** | Pointer to **int32** |  | [optional] 
**Revision** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **string** |  | [optional] 
**Subset** | Pointer to [**map[string]V1MatchString**](V1MatchString.md) |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1Rule

`func NewV1Rule() *V1Rule`

NewV1Rule instantiates a new V1Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RuleWithDefaults

`func NewV1RuleWithDefaults() *V1Rule`

NewV1RuleWithDefaults instantiates a new V1Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *V1Rule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *V1Rule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *V1Rule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *V1Rule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAdjuster

`func (o *V1Rule) GetAdjuster() V1AmountAdjuster`

GetAdjuster returns the Adjuster field if non-nil, zero value otherwise.

### GetAdjusterOk

`func (o *V1Rule) GetAdjusterOk() (*V1AmountAdjuster, bool)`

GetAdjusterOk returns a tuple with the Adjuster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjuster

`func (o *V1Rule) SetAdjuster(v V1AmountAdjuster)`

SetAdjuster sets Adjuster field to given value.

### HasAdjuster

`func (o *V1Rule) HasAdjuster() bool`

HasAdjuster returns a boolean if a field has been set.

### GetAmountMode

`func (o *V1Rule) GetAmountMode() int32`

GetAmountMode returns the AmountMode field if non-nil, zero value otherwise.

### GetAmountModeOk

`func (o *V1Rule) GetAmountModeOk() (*int32, bool)`

GetAmountModeOk returns a tuple with the AmountMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountMode

`func (o *V1Rule) SetAmountMode(v int32)`

SetAmountMode sets AmountMode field to given value.

### HasAmountMode

`func (o *V1Rule) HasAmountMode() bool`

HasAmountMode returns a boolean if a field has been set.

### GetAmounts

`func (o *V1Rule) GetAmounts() []V1Amount`

GetAmounts returns the Amounts field if non-nil, zero value otherwise.

### GetAmountsOk

`func (o *V1Rule) GetAmountsOk() (*[]V1Amount, bool)`

GetAmountsOk returns a tuple with the Amounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmounts

`func (o *V1Rule) SetAmounts(v []V1Amount)`

SetAmounts sets Amounts field to given value.

### HasAmounts

`func (o *V1Rule) HasAmounts() bool`

HasAmounts returns a boolean if a field has been set.

### GetArguments

`func (o *V1Rule) GetArguments() []V1MatchArgument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *V1Rule) GetArgumentsOk() (*[]V1MatchArgument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *V1Rule) SetArguments(v []V1MatchArgument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *V1Rule) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetCluster

`func (o *V1Rule) GetCluster() V1RateLimitCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *V1Rule) GetClusterOk() (*V1RateLimitCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *V1Rule) SetCluster(v V1RateLimitCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *V1Rule) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCtime

`func (o *V1Rule) GetCtime() string`

GetCtime returns the Ctime field if non-nil, zero value otherwise.

### GetCtimeOk

`func (o *V1Rule) GetCtimeOk() (*string, bool)`

GetCtimeOk returns a tuple with the Ctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtime

`func (o *V1Rule) SetCtime(v string)`

SetCtime sets Ctime field to given value.

### HasCtime

`func (o *V1Rule) HasCtime() bool`

HasCtime returns a boolean if a field has been set.

### GetDisable

`func (o *V1Rule) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *V1Rule) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *V1Rule) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *V1Rule) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetEtime

`func (o *V1Rule) GetEtime() string`

GetEtime returns the Etime field if non-nil, zero value otherwise.

### GetEtimeOk

`func (o *V1Rule) GetEtimeOk() (*string, bool)`

GetEtimeOk returns a tuple with the Etime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtime

`func (o *V1Rule) SetEtime(v string)`

SetEtime sets Etime field to given value.

### HasEtime

`func (o *V1Rule) HasEtime() bool`

HasEtime returns a boolean if a field has been set.

### GetFailover

`func (o *V1Rule) GetFailover() int32`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *V1Rule) GetFailoverOk() (*int32, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *V1Rule) SetFailover(v int32)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *V1Rule) HasFailover() bool`

HasFailover returns a boolean if a field has been set.

### GetId

`func (o *V1Rule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1Rule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1Rule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1Rule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabels

`func (o *V1Rule) GetLabels() map[string]V1MatchString`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1Rule) GetLabelsOk() (*map[string]V1MatchString, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1Rule) SetLabels(v map[string]V1MatchString)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1Rule) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMaxQueueDelay

`func (o *V1Rule) GetMaxQueueDelay() int32`

GetMaxQueueDelay returns the MaxQueueDelay field if non-nil, zero value otherwise.

### GetMaxQueueDelayOk

`func (o *V1Rule) GetMaxQueueDelayOk() (*int32, bool)`

GetMaxQueueDelayOk returns a tuple with the MaxQueueDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueueDelay

`func (o *V1Rule) SetMaxQueueDelay(v int32)`

SetMaxQueueDelay sets MaxQueueDelay field to given value.

### HasMaxQueueDelay

`func (o *V1Rule) HasMaxQueueDelay() bool`

HasMaxQueueDelay returns a boolean if a field has been set.

### GetMethod

`func (o *V1Rule) GetMethod() V1MatchString`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V1Rule) GetMethodOk() (*V1MatchString, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V1Rule) SetMethod(v V1MatchString)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V1Rule) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMtime

`func (o *V1Rule) GetMtime() string`

GetMtime returns the Mtime field if non-nil, zero value otherwise.

### GetMtimeOk

`func (o *V1Rule) GetMtimeOk() (*string, bool)`

GetMtimeOk returns a tuple with the Mtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtime

`func (o *V1Rule) SetMtime(v string)`

SetMtime sets Mtime field to given value.

### HasMtime

`func (o *V1Rule) HasMtime() bool`

HasMtime returns a boolean if a field has been set.

### GetName

`func (o *V1Rule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Rule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Rule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1Rule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *V1Rule) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1Rule) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1Rule) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *V1Rule) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPriority

`func (o *V1Rule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *V1Rule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *V1Rule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *V1Rule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRegexCombine

`func (o *V1Rule) GetRegexCombine() bool`

GetRegexCombine returns the RegexCombine field if non-nil, zero value otherwise.

### GetRegexCombineOk

`func (o *V1Rule) GetRegexCombineOk() (*bool, bool)`

GetRegexCombineOk returns a tuple with the RegexCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexCombine

`func (o *V1Rule) SetRegexCombine(v bool)`

SetRegexCombine sets RegexCombine field to given value.

### HasRegexCombine

`func (o *V1Rule) HasRegexCombine() bool`

HasRegexCombine returns a boolean if a field has been set.

### GetReport

`func (o *V1Rule) GetReport() V1Report`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *V1Rule) GetReportOk() (*V1Report, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *V1Rule) SetReport(v V1Report)`

SetReport sets Report field to given value.

### HasReport

`func (o *V1Rule) HasReport() bool`

HasReport returns a boolean if a field has been set.

### GetResource

`func (o *V1Rule) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *V1Rule) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *V1Rule) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *V1Rule) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRevision

`func (o *V1Rule) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1Rule) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1Rule) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1Rule) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetService

`func (o *V1Rule) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *V1Rule) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *V1Rule) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *V1Rule) HasService() bool`

HasService returns a boolean if a field has been set.

### GetServiceToken

`func (o *V1Rule) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *V1Rule) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *V1Rule) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *V1Rule) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetSubset

`func (o *V1Rule) GetSubset() map[string]V1MatchString`

GetSubset returns the Subset field if non-nil, zero value otherwise.

### GetSubsetOk

`func (o *V1Rule) GetSubsetOk() (*map[string]V1MatchString, bool)`

GetSubsetOk returns a tuple with the Subset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubset

`func (o *V1Rule) SetSubset(v map[string]V1MatchString)`

SetSubset sets Subset field to given value.

### HasSubset

`func (o *V1Rule) HasSubset() bool`

HasSubset returns a boolean if a field has been set.

### GetType

`func (o *V1Rule) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1Rule) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1Rule) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *V1Rule) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


