# \RateLimitsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRateLimits**](RateLimitsApi.md#CreateRateLimits) | **Post** /naming/v1/ratelimits | 创建限流规则
[**DeleteRateLimits**](RateLimitsApi.md#DeleteRateLimits) | **Post** /naming/v1/ratelimits/delete | 删除限流规则
[**EnableRateLimits**](RateLimitsApi.md#EnableRateLimits) | **Put** /naming/v1/ratelimits/enable | 启用限流规则
[**GetRateLimits**](RateLimitsApi.md#GetRateLimits) | **Get** /naming/v1/ratelimits | 查询限流规则
[**UpdateRateLimits**](RateLimitsApi.md#UpdateRateLimits) | **Put** /naming/v1/ratelimits | 更新限流规则



## CreateRateLimits

> CreateRateLimits(ctx).Body(body).Execute()

创建限流规则



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
    body := []openapiclient.V1RateLimit{*openapiclient.NewV1RateLimit()} // []V1RateLimit | create rate limits

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RateLimitsApi.CreateRateLimits(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsApi.CreateRateLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1RateLimit**](V1RateLimit.md) | create rate limits | 

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


## DeleteRateLimits

> DeleteRateLimits(ctx).Body(body).Execute()

删除限流规则



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
    body := []openapiclient.V1RateLimit{*openapiclient.NewV1RateLimit()} // []V1RateLimit | delete rate limits

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RateLimitsApi.DeleteRateLimits(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsApi.DeleteRateLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1RateLimit**](V1RateLimit.md) | delete rate limits | 

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


## EnableRateLimits

> EnableRateLimits(ctx).Body(body).Execute()

启用限流规则



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
    body := []openapiclient.V1RateLimit{*openapiclient.NewV1RateLimit()} // []V1RateLimit | enable rate limits

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RateLimitsApi.EnableRateLimits(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsApi.EnableRateLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1RateLimit**](V1RateLimit.md) | enable rate limits | 

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


## GetRateLimits

> GetRateLimits(ctx, id, name, service, namespace, method, disable, brief, offset, limit).Execute()

查询限流规则



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
    id := "id_example" // string | 规则ID
    name := "name_example" // string | 规则名称
    service := "service_example" // string | 服务名称
    namespace := "namespace_example" // string | 命名空间
    method := "method_example" // string | 限流接口名，默认为模糊匹配 
    disable := true // bool | 规则是否启用，true为不启用，false为启用
    brief := true // bool | 是否只显示概要信息，brief=true时，则不返回规则详情，只返回规则列表概要信息，默认为false (default to false)
    offset := int32(56) // int32 | 分页的起始位置，默认为0 (default to 0)
    limit := int32(56) // int32 | 每页行数，默认100   (default to 100)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RateLimitsApi.GetRateLimits(context.Background(), id, name, service, namespace, method, disable, brief, offset, limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsApi.GetRateLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 规则ID | 
**name** | **string** | 规则名称 | 
**service** | **string** | 服务名称 | 
**namespace** | **string** | 命名空间 | 
**method** | **string** | 限流接口名，默认为模糊匹配  | 
**disable** | **bool** | 规则是否启用，true为不启用，false为启用 | 
**brief** | **bool** | 是否只显示概要信息，brief&#x3D;true时，则不返回规则详情，只返回规则列表概要信息，默认为false | [default to false]
**offset** | **int32** | 分页的起始位置，默认为0 | [default to 0]
**limit** | **int32** | 每页行数，默认100   | [default to 100]

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateLimitsRequest struct via the builder pattern


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


## UpdateRateLimits

> UpdateRateLimits(ctx).Body(body).Execute()

更新限流规则



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
    body := []openapiclient.V1RateLimit{*openapiclient.NewV1RateLimit()} // []V1RateLimit | update rate limits

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RateLimitsApi.UpdateRateLimits(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RateLimitsApi.UpdateRateLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRateLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1RateLimit**](V1RateLimit.md) | update rate limits | 

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

