# \RoutingRulesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoutings**](RoutingRulesApi.md#CreateRoutings) | **Post** /naming/v1/routings | 创建路由规则
[**DeleteRoutings**](RoutingRulesApi.md#DeleteRoutings) | **Post** /naming/v1/routings/delete | 删除路由规则
[**GetRoutings**](RoutingRulesApi.md#GetRoutings) | **Get** /naming/v1/routings | 查询路由规则
[**UpdateRoutings**](RoutingRulesApi.md#UpdateRoutings) | **Put** /naming/v1/routings | 更新路由规则
[**V2CreateRoutings**](RoutingRulesApi.md#V2CreateRoutings) | **Post** /naming/v2/routings | 创建路由规则
[**V2DeleteRoutings**](RoutingRulesApi.md#V2DeleteRoutings) | **Post** /naming/v2/routings/delete | 删除路由规则
[**V2EnableRoutings**](RoutingRulesApi.md#V2EnableRoutings) | **Put** /naming/v2/routings/enable | 启用路由规则
[**V2GetRoutings**](RoutingRulesApi.md#V2GetRoutings) | **Get** /naming/v2/routings | 获取路由规则
[**V2UpdateRoutings**](RoutingRulesApi.md#V2UpdateRoutings) | **Put** /naming/v2/routings | 更新路由规则



## CreateRoutings

> CreateRoutings(ctx).Body(body).Execute()

创建路由规则



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
    body := []openapiclient.V1Routing{*openapiclient.NewV1Routing()} // []V1Routing | create routing rules

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingRulesApi.CreateRoutings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.CreateRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Routing**](V1Routing.md) | create routing rules | 

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


## DeleteRoutings

> DeleteRoutings(ctx).Body(body).Execute()

删除路由规则



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
    body := []openapiclient.V1Routing{*openapiclient.NewV1Routing()} // []V1Routing | delete routing rules

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingRulesApi.DeleteRoutings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.DeleteRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoutingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Routing**](V1Routing.md) | delete routing rules | 

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


## GetRoutings

> GetRoutings(ctx, service, namespace, offset, limit).Execute()

查询路由规则



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
    service := "service_example" // string | 服务名称
    namespace := "namespace_example" // string | 命名空间
    offset := int32(56) // int32 | 分页的起始位置，默认为0 (default to 0)
    limit := int32(56) // int32 | 每页行数，默认100 (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingRulesApi.GetRoutings(context.Background(), service, namespace, offset, limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.GetRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | 服务名称 | 
**namespace** | **string** | 命名空间 | 
**offset** | **int32** | 分页的起始位置，默认为0 | [default to 0]
**limit** | **int32** | 每页行数，默认100 | [default to 100]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





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


## UpdateRoutings

> UpdateRoutings(ctx).Body(body).Execute()

更新路由规则



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
    body := []openapiclient.V1Routing{*openapiclient.NewV1Routing()} // []V1Routing | update routing rules

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingRulesApi.UpdateRoutings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.UpdateRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoutingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Routing**](V1Routing.md) | update routing rules | 

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


## V2CreateRoutings

> V2CreateRoutings(ctx).Body(body).Execute()

创建路由规则



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
    body := []openapiclient.V2Routing{*openapiclient.NewV2Routing()} // []V2Routing | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingRulesApi.V2CreateRoutings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.V2CreateRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2CreateRoutingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V2Routing**](V2Routing.md) |  | 

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


## V2DeleteRoutings

> V2DeleteRoutings(ctx).Execute()

删除路由规则



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
    resp, r, err := api_client.RoutingRulesApi.V2DeleteRoutings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.V2DeleteRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2DeleteRoutingsRequest struct via the builder pattern


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


## V2EnableRoutings

> V2EnableRoutings(ctx).Execute()

启用路由规则



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
    resp, r, err := api_client.RoutingRulesApi.V2EnableRoutings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.V2EnableRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2EnableRoutingsRequest struct via the builder pattern


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


## V2GetRoutings

> V2GetRoutings(ctx).Execute()

获取路由规则



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
    resp, r, err := api_client.RoutingRulesApi.V2GetRoutings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.V2GetRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiV2GetRoutingsRequest struct via the builder pattern


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


## V2UpdateRoutings

> V2UpdateRoutings(ctx).Body(body).Execute()

更新路由规则



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
    body := []openapiclient.V2Routing{*openapiclient.NewV2Routing()} // []V2Routing | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RoutingRulesApi.V2UpdateRoutings(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesApi.V2UpdateRoutings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV2UpdateRoutingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V2Routing**](V2Routing.md) |  | 

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

