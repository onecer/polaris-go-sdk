# \AuthApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthStatus**](AuthApi.md#AuthStatus) | **Get** /core/v1/auth/status | 查询鉴权开关信息
[**CreateStrategy**](AuthApi.md#CreateStrategy) | **Post** /core/v1/auth/strategy | 创建鉴权策略
[**DeleteStrategies**](AuthApi.md#DeleteStrategies) | **Post** /core/v1/auth/strategies/delete | 删除鉴权策略
[**GetPrincipalResources**](AuthApi.md#GetPrincipalResources) | **Get** /core/v1/auth/principal/resources | 获取鉴权策略详细
[**GetStrategies**](AuthApi.md#GetStrategies) | **Get** /core/v1/auth/strategies | 查询鉴权策略列表
[**GetStrategy**](AuthApi.md#GetStrategy) | **Get** /core/v1/auth/strategy/detail | 获取鉴权策略详细
[**UpdateStrategies**](AuthApi.md#UpdateStrategies) | **Put** /core/v1/auth/strategies | 更新鉴权策略



## AuthStatus

> AuthStatus(ctx).Execute()

查询鉴权开关信息



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.AuthStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.AuthStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAuthStatusRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStrategy

> CreateStrategy(ctx).Body(body).Execute()

创建鉴权策略



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewV1AuthStrategy() // V1AuthStrategy | create auth strategy

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.CreateStrategy(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.CreateStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1AuthStrategy**](V1AuthStrategy.md) | create auth strategy | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStrategies

> DeleteStrategies(ctx).Body(body).Execute()

删除鉴权策略



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewV1AuthStrategy() // V1AuthStrategy | delete auth strategy

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.DeleteStrategies(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.DeleteStrategies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1AuthStrategy**](V1AuthStrategy.md) | delete auth strategy | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalResources

> GetPrincipalResources(ctx).PrincipalId(principalId).PrincipalType(principalType).Execute()

获取鉴权策略详细



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    principalId := "principalId_example" // string | 策略ID
    principalType := "principalType_example" // string | Principal类别，user/group

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.GetPrincipalResources(context.Background()).PrincipalId(principalId).PrincipalType(principalType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetPrincipalResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **principalId** | **string** | 策略ID | 
 **principalType** | **string** | Principal类别，user/group | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStrategies

> GetStrategies(ctx).Id(id).Name(name).Default_(default_).ResId(resId).ResType(resType).PrincipalId(principalId).PrincipalType(principalType).ShowDetail(showDetail).Offset(offset).Limit(limit).Execute()

查询鉴权策略列表



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 策略ID (optional)
    name := "name_example" // string | 策略名称, 当前仅提供全模糊搜索 (optional)
    default_ := "default__example" // string | “0” 查询自定义策略；“1” 查询默认策略；不填则为查询（默认+自定义）鉴权策略 (optional)
    resId := "resId_example" // string | 资源ID (optional)
    resType := "resType_example" // string | 资源类型, namespace、service、config_group (optional)
    principalId := "principalId_example" // string | 成员ID (optional)
    principalType := "principalType_example" // string | 成员类型, user、group (optional)
    showDetail := true // bool | 是否显示策略详细 (optional)
    offset := int32(56) // int32 | 查询偏移量, 默认为0 (optional) (default to 0)
    limit := int32(56) // int32 | 本次查询条数, 最大为100 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.GetStrategies(context.Background()).Id(id).Name(name).Default_(default_).ResId(resId).ResType(resType).PrincipalId(principalId).PrincipalType(principalType).ShowDetail(showDetail).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetStrategies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 策略ID | 
 **name** | **string** | 策略名称, 当前仅提供全模糊搜索 | 
 **default_** | **string** | “0” 查询自定义策略；“1” 查询默认策略；不填则为查询（默认+自定义）鉴权策略 | 
 **resId** | **string** | 资源ID | 
 **resType** | **string** | 资源类型, namespace、service、config_group | 
 **principalId** | **string** | 成员ID | 
 **principalType** | **string** | 成员类型, user、group | 
 **showDetail** | **bool** | 是否显示策略详细 | 
 **offset** | **int32** | 查询偏移量, 默认为0 | [default to 0]
 **limit** | **int32** | 本次查询条数, 最大为100 | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStrategy

> GetStrategy(ctx).Id(id).Execute()

获取鉴权策略详细



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 策略ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.GetStrategy(context.Background()).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.GetStrategy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 策略ID | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStrategies

> UpdateStrategies(ctx).Body(body).Execute()

更新鉴权策略



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewV1AuthStrategy() // V1AuthStrategy | update auth strategy

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthApi.UpdateStrategies(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.UpdateStrategies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStrategiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1AuthStrategy**](V1AuthStrategy.md) | update auth strategy | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

