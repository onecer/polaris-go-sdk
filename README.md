# Go API client for polaris

一个支持多语言、多框架的云原生服务发现和治理中心

提供高性能SDK和无侵入Sidecar两种接入方式



## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://polarismesh.cn/](https://polarismesh.cn/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./polaris"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthApi* | [**AuthStatus**](docs/AuthApi.md#authstatus) | **Get** /core/v1/auth/status | 查询鉴权开关信息
*AuthApi* | [**CreateStrategy**](docs/AuthApi.md#createstrategy) | **Post** /core/v1/auth/strategy | 创建鉴权策略
*AuthApi* | [**DeleteStrategies**](docs/AuthApi.md#deletestrategies) | **Post** /core/v1/auth/strategies/delete | 删除鉴权策略
*AuthApi* | [**GetPrincipalResources**](docs/AuthApi.md#getprincipalresources) | **Get** /core/v1/auth/principal/resources | 获取鉴权策略详细
*AuthApi* | [**GetStrategies**](docs/AuthApi.md#getstrategies) | **Get** /core/v1/auth/strategies | 查询鉴权策略列表
*AuthApi* | [**GetStrategy**](docs/AuthApi.md#getstrategy) | **Get** /core/v1/auth/strategy/detail | 获取鉴权策略详细
*AuthApi* | [**UpdateStrategies**](docs/AuthApi.md#updatestrategies) | **Put** /core/v1/auth/strategies | 更新鉴权策略
*CircuitBreakersApi* | [**CreateCircuitBreakerVersions**](docs/CircuitBreakersApi.md#createcircuitbreakerversions) | **Post** /naming/v1/circuitbreakers/version | 创建熔断规则版本
*CircuitBreakersApi* | [**CreateCircuitBreakers**](docs/CircuitBreakersApi.md#createcircuitbreakers) | **Post** /naming/v1/circuitbreakers | 创建熔断规则
*CircuitBreakersApi* | [**DeleteCircuitBreakers**](docs/CircuitBreakersApi.md#deletecircuitbreakers) | **Post** /naming/v1/circuitbreakers/delete | 删除熔断规则
*CircuitBreakersApi* | [**GetCircuitBreaker**](docs/CircuitBreakersApi.md#getcircuitbreaker) | **Get** /naming/v1/circuitbreaker | 查询熔断规则
*CircuitBreakersApi* | [**GetCircuitBreakerToken**](docs/CircuitBreakersApi.md#getcircuitbreakertoken) | **Get** /naming/v1/circuitbreaker/token | 查询熔断规则Token
*CircuitBreakersApi* | [**GetCircuitBreakerVersions**](docs/CircuitBreakersApi.md#getcircuitbreakerversions) | **Get** /naming/v1/circuitbreaker/versions | 查询熔断规则版本
*CircuitBreakersApi* | [**GetMasterCircuitBreakers**](docs/CircuitBreakersApi.md#getmastercircuitbreakers) | **Get** /naming/v1/circuitbreakers/master | 查询熔断规则Master版本
*CircuitBreakersApi* | [**GetReleaseCircuitBreakers**](docs/CircuitBreakersApi.md#getreleasecircuitbreakers) | **Get** /naming/v1/circuitbreakers/release | 根据规则id查询已发布的熔断规则
*CircuitBreakersApi* | [**ReleaseCircuitBreakers**](docs/CircuitBreakersApi.md#releasecircuitbreakers) | **Post** /naming/v1/circuitbreakers/release | 发布熔断规则
*CircuitBreakersApi* | [**UnBindCircuitBreakers**](docs/CircuitBreakersApi.md#unbindcircuitbreakers) | **Post** /naming/v1/circuitbreakers/unbind | 解绑熔断规则
*CircuitBreakersApi* | [**UpdateCircuitBreakers**](docs/CircuitBreakersApi.md#updatecircuitbreakers) | **Put** /naming/v1/circuitbreakers | 更新熔断规则
*ConfigClientApi* | [**GetConfigFile**](docs/ConfigClientApi.md#getconfigfile) | **Get** /config/v1/GetConfigFile | 拉取配置
*ConfigClientApi* | [**WatchConfigFile**](docs/ConfigClientApi.md#watchconfigfile) | **Post** /config/v1/WatchConfigFile | 监听配置
*ConfigConsoleApi* | [**BatchDeleteConfigFile**](docs/ConfigConsoleApi.md#batchdeleteconfigfile) | **Post** /config/v1/configfiles/batchdelete | 批量删除配置文件
*ConfigConsoleApi* | [**CreateConfigFile**](docs/ConfigConsoleApi.md#createconfigfile) | **Post** /config/v1/configfiles | 创建配置文件
*ConfigConsoleApi* | [**CreateConfigFileGroup**](docs/ConfigConsoleApi.md#createconfigfilegroup) | **Post** /config/v1/configfilegroups | 创建配置文件组
*ConfigConsoleApi* | [**CreateConfigFileTemplate**](docs/ConfigConsoleApi.md#createconfigfiletemplate) | **Post** /config/v1/configfiletemplates | 创建配置模板
*ConfigConsoleApi* | [**DeleteConfigFile**](docs/ConfigConsoleApi.md#deleteconfigfile) | **Delete** /config/v1/configfiles | 创建配置文件
*ConfigConsoleApi* | [**DeleteConfigFileGroup**](docs/ConfigConsoleApi.md#deleteconfigfilegroup) | **Delete** /config/v1/configfilegroups | 删除配置文件
*ConfigConsoleApi* | [**GetAllConfigFileTemplates**](docs/ConfigConsoleApi.md#getallconfigfiletemplates) | **Get** /config/v1/configfiletemplates | 获取配置模板
*ConfigConsoleApi* | [**GetConfigFile**](docs/ConfigConsoleApi.md#getconfigfile) | **Get** /config/v1/configfiles | 拉取配置
*ConfigConsoleApi* | [**GetConfigFileRelease**](docs/ConfigConsoleApi.md#getconfigfilerelease) | **Get** /config/v1/configfiles/release | 获取配置文件最后一次全量发布信息
*ConfigConsoleApi* | [**GetConfigFileReleaseHistory**](docs/ConfigConsoleApi.md#getconfigfilereleasehistory) | **Get** /config/v1/configfiles/releasehistory | 获取配置文件发布历史记录
*ConfigConsoleApi* | [**PublishConfigFile**](docs/ConfigConsoleApi.md#publishconfigfile) | **Post** /config/v1/configfiles/release | 发布配置文件
*ConfigConsoleApi* | [**QueryConfigFileGroups**](docs/ConfigConsoleApi.md#queryconfigfilegroups) | **Get** /config/v1/configfilegroups | 搜索配置文件组
*ConfigConsoleApi* | [**QueryConfigFilesByGroup**](docs/ConfigConsoleApi.md#queryconfigfilesbygroup) | **Get** /config/v1/configfiles/by-group | 搜索配置文件
*ConfigConsoleApi* | [**SearchConfigFile**](docs/ConfigConsoleApi.md#searchconfigfile) | **Get** /config/v1/configfiles/search | 搜索配置文件
*ConfigConsoleApi* | [**UpdateConfigFile**](docs/ConfigConsoleApi.md#updateconfigfile) | **Put** /config/v1/configfiles | 创建配置文件
*ConfigConsoleApi* | [**UpdateConfigFileGroup**](docs/ConfigConsoleApi.md#updateconfigfilegroup) | **Put** /config/v1/configfilegroups | 更新配置文件组
*InstancesApi* | [**CreateInstances**](docs/InstancesApi.md#createinstances) | **Post** /naming/v1/instances | 创建实例
*InstancesApi* | [**DeleteInstances**](docs/InstancesApi.md#deleteinstances) | **Post** /naming/v1/instances/delete | 删除实例(根据实例ID)
*InstancesApi* | [**DeleteInstancesByHost**](docs/InstancesApi.md#deleteinstancesbyhost) | **Post** /naming/v1/instances/delete/host | 删除实例(根据主机)
*InstancesApi* | [**GetInstanceLabels**](docs/InstancesApi.md#getinstancelabels) | **Get** /naming/v1/instances/labels | 查询某个服务下所有实例的标签信息
*InstancesApi* | [**GetInstances**](docs/InstancesApi.md#getinstances) | **Get** /naming/v1/instances | 查询服务实例
*InstancesApi* | [**GetInstancesCount**](docs/InstancesApi.md#getinstancescount) | **Get** /naming/v1/instances/count | 查询服务实例数量
*InstancesApi* | [**UpdateInstances**](docs/InstancesApi.md#updateinstances) | **Put** /naming/v1/instances | 更新实例
*InstancesApi* | [**UpdateInstancesIsolate**](docs/InstancesApi.md#updateinstancesisolate) | **Put** /naming/v1/instances/isolate/host | 修改服务实例的隔离状态
*MaintainApi* | [**CleanInstance**](docs/MaintainApi.md#cleaninstance) | **Post** /maintain/v1/instance/clean | 彻底清理flag&#x3D;1的实例
*MaintainApi* | [**CloseConnections**](docs/MaintainApi.md#closeconnections) | **Post** /maintain/v1/apiserver/conn/close | 关闭指定client ip的连接
*MaintainApi* | [**FreeOSMemory**](docs/MaintainApi.md#freeosmemory) | **Post** /maintain/v1/memory/free | 释放系统内存
*MaintainApi* | [**GetLastHeartbeat**](docs/MaintainApi.md#getlastheartbeat) | **Get** /maintain/v1/instance/heartbeat | 获取上一次心跳的时间
*MaintainApi* | [**GetLogOutputLevel**](docs/MaintainApi.md#getlogoutputlevel) | **Get** /maintain/v1/log/outputlevel | 获取日志输出级别
*MaintainApi* | [**GetServerConnStats**](docs/MaintainApi.md#getserverconnstats) | **Get** /maintain/v1/apiserver/conn/stats | 获取服务端连接统计信息
*MaintainApi* | [**GetServerConnections**](docs/MaintainApi.md#getserverconnections) | **Get** /maintain/v1/apiserver/conn | 获取服务端连接数
*MaintainApi* | [**SetLogOutputLevel**](docs/MaintainApi.md#setlogoutputlevel) | **Put** /maintain/v1/log/outputlevel | 设置日志输出级别
*NamespacesApi* | [**CoreCreateNamespaces**](docs/NamespacesApi.md#corecreatenamespaces) | **Post** /core/v1/namespaces | 创建命名空间
*NamespacesApi* | [**CoreDeleteNamespaces**](docs/NamespacesApi.md#coredeletenamespaces) | **Post** /core/v1/namespaces/delete | 删除命名空间
*NamespacesApi* | [**CoreGetNamespaceToken**](docs/NamespacesApi.md#coregetnamespacetoken) | **Get** /core/v1/namespaces/token | 查询命名空间Token
*NamespacesApi* | [**CoreGetNamespaces**](docs/NamespacesApi.md#coregetnamespaces) | **Get** /core/v1/namespaces | 查询命名空间列表
*NamespacesApi* | [**CoreUpdateNamespaceToken**](docs/NamespacesApi.md#coreupdatenamespacetoken) | **Put** /core/v1/namespaces/token | 更新命名空间Token
*NamespacesApi* | [**CoreUpdateNamespaces**](docs/NamespacesApi.md#coreupdatenamespaces) | **Put** /core/v1/namespaces | 更新命名空间
*NamespacesApi* | [**CreateNamespaces**](docs/NamespacesApi.md#createnamespaces) | **Post** /naming/v1/namespaces | 创建命名空间
*NamespacesApi* | [**DeleteNamespaces**](docs/NamespacesApi.md#deletenamespaces) | **Post** /naming/v1/namespaces/delete | 删除命名空间
*NamespacesApi* | [**GetNamespaceToken**](docs/NamespacesApi.md#getnamespacetoken) | **Get** /naming/v1/namespace/token | 查询命名空间Token
*NamespacesApi* | [**GetNamespaces**](docs/NamespacesApi.md#getnamespaces) | **Get** /naming/v1/namespaces | 获取命名空间列表
*NamespacesApi* | [**UpdateNamespaceToken**](docs/NamespacesApi.md#updatenamespacetoken) | **Put** /naming/v1/namespace/token | 更新命名空间Token
*NamespacesApi* | [**UpdateNamespaces**](docs/NamespacesApi.md#updatenamespaces) | **Put** /naming/v1/namespaces | 更新命名空间
*RateLimitsApi* | [**CreateRateLimits**](docs/RateLimitsApi.md#createratelimits) | **Post** /naming/v1/ratelimits | 创建限流规则
*RateLimitsApi* | [**DeleteRateLimits**](docs/RateLimitsApi.md#deleteratelimits) | **Post** /naming/v1/ratelimits/delete | 删除限流规则
*RateLimitsApi* | [**EnableRateLimits**](docs/RateLimitsApi.md#enableratelimits) | **Put** /naming/v1/ratelimits/enable | 启用限流规则
*RateLimitsApi* | [**GetRateLimits**](docs/RateLimitsApi.md#getratelimits) | **Get** /naming/v1/ratelimits | 查询限流规则
*RateLimitsApi* | [**UpdateRateLimits**](docs/RateLimitsApi.md#updateratelimits) | **Put** /naming/v1/ratelimits | 更新限流规则
*RegisterInstanceApi* | [**DeregisterInstance**](docs/RegisterInstanceApi.md#deregisterinstance) | **Post** /v1/DeregisterInstance | 注销实例
*RegisterInstanceApi* | [**Discover**](docs/RegisterInstanceApi.md#discover) | **Post** /v1/Discover | 服务发现
*RegisterInstanceApi* | [**Heartbeat**](docs/RegisterInstanceApi.md#heartbeat) | **Post** /v1/Heartbeat | 上报心跳
*RegisterInstanceApi* | [**RegisterInstance**](docs/RegisterInstanceApi.md#registerinstance) | **Post** /v1/RegisterInstance | 注册实例
*RegisterInstanceApi* | [**ReportClient**](docs/RegisterInstanceApi.md#reportclient) | **Post** /v1/ReportClient | 上报客户端
*RegisterInstanceApi* | [**V2Discover**](docs/RegisterInstanceApi.md#v2discover) | **Post** /v2/Discover | 服务发现
*RoutingRulesApi* | [**CreateRoutings**](docs/RoutingRulesApi.md#createroutings) | **Post** /naming/v1/routings | 创建路由规则
*RoutingRulesApi* | [**DeleteRoutings**](docs/RoutingRulesApi.md#deleteroutings) | **Post** /naming/v1/routings/delete | 删除路由规则
*RoutingRulesApi* | [**GetRoutings**](docs/RoutingRulesApi.md#getroutings) | **Get** /naming/v1/routings | 查询路由规则
*RoutingRulesApi* | [**UpdateRoutings**](docs/RoutingRulesApi.md#updateroutings) | **Put** /naming/v1/routings | 更新路由规则
*RoutingRulesApi* | [**V2CreateRoutings**](docs/RoutingRulesApi.md#v2createroutings) | **Post** /naming/v2/routings | 创建路由规则
*RoutingRulesApi* | [**V2DeleteRoutings**](docs/RoutingRulesApi.md#v2deleteroutings) | **Post** /naming/v2/routings/delete | 删除路由规则
*RoutingRulesApi* | [**V2EnableRoutings**](docs/RoutingRulesApi.md#v2enableroutings) | **Put** /naming/v2/routings/enable | 启用路由规则
*RoutingRulesApi* | [**V2GetRoutings**](docs/RoutingRulesApi.md#v2getroutings) | **Get** /naming/v2/routings | 获取路由规则
*RoutingRulesApi* | [**V2UpdateRoutings**](docs/RoutingRulesApi.md#v2updateroutings) | **Put** /naming/v2/routings | 更新路由规则
*ServicesApi* | [**CreateServiceAlias**](docs/ServicesApi.md#createservicealias) | **Post** /naming/v1/service/alias | 创建服务别名
*ServicesApi* | [**CreateServices**](docs/ServicesApi.md#createservices) | **Post** /naming/v1/services | 创建服务
*ServicesApi* | [**DeleteServiceAliases**](docs/ServicesApi.md#deleteservicealiases) | **Post** /naming/v1/service/aliases/delete | 删除服务别名
*ServicesApi* | [**DeleteServices**](docs/ServicesApi.md#deleteservices) | **Post** /naming/v1/services/delete | 删除服务
*ServicesApi* | [**GetCircuitBreakerByService**](docs/ServicesApi.md#getcircuitbreakerbyservice) | **Get** /naming/v1/service/circuitbreaker | 根据服务查询熔断规则
*ServicesApi* | [**GetServiceAliases**](docs/ServicesApi.md#getservicealiases) | **Get** /naming/v1/service/aliases | 查询服务别名
*ServicesApi* | [**GetServiceOwner**](docs/ServicesApi.md#getserviceowner) | **Post** /naming/v1/service/owner | 根据服务获取服务负责人
*ServicesApi* | [**GetServiceToken**](docs/ServicesApi.md#getservicetoken) | **Get** /naming/v1/service/token | 查询服务Token
*ServicesApi* | [**GetServices**](docs/ServicesApi.md#getservices) | **Get** /naming/v1/services | 获取服务列表
*ServicesApi* | [**GetServicesCount**](docs/ServicesApi.md#getservicescount) | **Get** /naming/v1/services/count | 获取服务数量
*ServicesApi* | [**UpdateServiceAlias**](docs/ServicesApi.md#updateservicealias) | **Put** /naming/v1/service/alias | 更新服务别名
*ServicesApi* | [**UpdateServiceToken**](docs/ServicesApi.md#updateservicetoken) | **Put** /naming/v1/service/token | 更新服务Token
*ServicesApi* | [**UpdateServices**](docs/ServicesApi.md#updateservices) | **Put** /naming/v1/services | 更新服务
*UserGroupApi* | [**CreateGroup**](docs/UserGroupApi.md#creategroup) | **Post** /core/v1/usergroup | 创建用户组
*UserGroupApi* | [**DeleteGroups**](docs/UserGroupApi.md#deletegroups) | **Post** /core/v1/usergroups/delete | 删除用户组
*UserGroupApi* | [**GetGroup**](docs/UserGroupApi.md#getgroup) | **Get** /core/v1/usergroup/detail | 获取用户组详情
*UserGroupApi* | [**GetGroupToken**](docs/UserGroupApi.md#getgrouptoken) | **Get** /core/v1/usergroup/token | 获取用户组 token
*UserGroupApi* | [**GetGroups**](docs/UserGroupApi.md#getgroups) | **Get** /core/v1/usergroups | 查询用户组列表
*UserGroupApi* | [**ResetGroupToken**](docs/UserGroupApi.md#resetgrouptoken) | **Put** /core/v1/usergroup/token/refresh | 重置用户组 token
*UserGroupApi* | [**UpdateGroupToken**](docs/UserGroupApi.md#updategrouptoken) | **Put** /core/v1/usergroup/token/status | 更新用户组 token
*UserGroupApi* | [**UpdateGroups**](docs/UserGroupApi.md#updategroups) | **Put** /core/v1/usergroups | 更新用户组
*UsersApi* | [**CreateUsers**](docs/UsersApi.md#createusers) | **Post** /core/v1/users | 创建用户
*UsersApi* | [**DeleteUsers**](docs/UsersApi.md#deleteusers) | **Post** /core/v1/users/delete | 删除用户
*UsersApi* | [**GetUserToken**](docs/UsersApi.md#getusertoken) | **Get** /core/v1/user/token | 获取用户Token
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /core/v1/users | 获取用户
*UsersApi* | [**Login**](docs/UsersApi.md#login) | **Post** /core/v1/user/login | 用户登录
*UsersApi* | [**ResetUserToken**](docs/UsersApi.md#resetusertoken) | **Put** /core/v1/user/token/refresh | 重置用户Token
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /core/v1/user | 更新用户
*UsersApi* | [**UpdateUserPassword**](docs/UsersApi.md#updateuserpassword) | **Put** /core/v1/user/password | 更新用户密码
*UsersApi* | [**UpdateUserToken**](docs/UsersApi.md#updateusertoken) | **Put** /core/v1/user/token/status | 更新用户Token


## Documentation For Models

 - [AnypbAny](docs/AnypbAny.md)
 - [DurationpbDuration](docs/DurationpbDuration.md)
 - [MaintainConnReq](docs/MaintainConnReq.md)
 - [V1Amount](docs/V1Amount.md)
 - [V1AmountAdjuster](docs/V1AmountAdjuster.md)
 - [V1AuthStrategy](docs/V1AuthStrategy.md)
 - [V1CbPolicy](docs/V1CbPolicy.md)
 - [V1CbPolicyConsecutiveErrConfig](docs/V1CbPolicyConsecutiveErrConfig.md)
 - [V1CbPolicyErrRateConfig](docs/V1CbPolicyErrRateConfig.md)
 - [V1CbPolicyErrRateConfigSpecialConfig](docs/V1CbPolicyErrRateConfigSpecialConfig.md)
 - [V1CbPolicySlowRateConfig](docs/V1CbPolicySlowRateConfig.md)
 - [V1CbRule](docs/V1CbRule.md)
 - [V1CircuitBreaker](docs/V1CircuitBreaker.md)
 - [V1Client](docs/V1Client.md)
 - [V1ClientConfigFileInfo](docs/V1ClientConfigFileInfo.md)
 - [V1ClientWatchConfigFileRequest](docs/V1ClientWatchConfigFileRequest.md)
 - [V1ClimbConfig](docs/V1ClimbConfig.md)
 - [V1ClimbConfigClimbThrottling](docs/V1ClimbConfigClimbThrottling.md)
 - [V1ClimbConfigMetricConfig](docs/V1ClimbConfigMetricConfig.md)
 - [V1ClimbConfigTriggerPolicy](docs/V1ClimbConfigTriggerPolicy.md)
 - [V1ClimbConfigTriggerPolicyErrorRate](docs/V1ClimbConfigTriggerPolicyErrorRate.md)
 - [V1ClimbConfigTriggerPolicyErrorRateSpecialConfig](docs/V1ClimbConfigTriggerPolicyErrorRateSpecialConfig.md)
 - [V1ClimbConfigTriggerPolicySlowRate](docs/V1ClimbConfigTriggerPolicySlowRate.md)
 - [V1ConfigFile](docs/V1ConfigFile.md)
 - [V1ConfigFileGroup](docs/V1ConfigFileGroup.md)
 - [V1ConfigFileRelease](docs/V1ConfigFileRelease.md)
 - [V1ConfigFileTag](docs/V1ConfigFileTag.md)
 - [V1ConfigRelease](docs/V1ConfigRelease.md)
 - [V1Destination](docs/V1Destination.md)
 - [V1DestinationSet](docs/V1DestinationSet.md)
 - [V1HealthCheck](docs/V1HealthCheck.md)
 - [V1HeartbeatHealthCheck](docs/V1HeartbeatHealthCheck.md)
 - [V1Instance](docs/V1Instance.md)
 - [V1Location](docs/V1Location.md)
 - [V1LoginRequest](docs/V1LoginRequest.md)
 - [V1MatchArgument](docs/V1MatchArgument.md)
 - [V1MatchString](docs/V1MatchString.md)
 - [V1ModifyUserPassword](docs/V1ModifyUserPassword.md)
 - [V1Namespace](docs/V1Namespace.md)
 - [V1Principal](docs/V1Principal.md)
 - [V1Principals](docs/V1Principals.md)
 - [V1RateLimit](docs/V1RateLimit.md)
 - [V1RateLimitCluster](docs/V1RateLimitCluster.md)
 - [V1RecoverConfig](docs/V1RecoverConfig.md)
 - [V1Report](docs/V1Report.md)
 - [V1Route](docs/V1Route.md)
 - [V1Routing](docs/V1Routing.md)
 - [V1Rule](docs/V1Rule.md)
 - [V1Service](docs/V1Service.md)
 - [V1ServiceAlias](docs/V1ServiceAlias.md)
 - [V1Source](docs/V1Source.md)
 - [V1SourceMatcher](docs/V1SourceMatcher.md)
 - [V1StatInfo](docs/V1StatInfo.md)
 - [V1StrategyResourceEntry](docs/V1StrategyResourceEntry.md)
 - [V1StrategyResources](docs/V1StrategyResources.md)
 - [V1User](docs/V1User.md)
 - [V1UserGroup](docs/V1UserGroup.md)
 - [V1UserGroupRelation](docs/V1UserGroupRelation.md)
 - [V2Routing](docs/V2Routing.md)
 - [WrapperspbInt32Value](docs/WrapperspbInt32Value.md)
 - [WrapperspbInt64Value](docs/WrapperspbInt64Value.md)


## Documentation For Authorization



### api_key

- **Type**: API key
- **API key parameter name**: X-Polaris-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Polaris-Token and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



