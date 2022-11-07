# \CircuitBreakersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCircuitBreakerVersions**](CircuitBreakersApi.md#CreateCircuitBreakerVersions) | **Post** /naming/v1/circuitbreakers/version | 创建熔断规则版本
[**CreateCircuitBreakers**](CircuitBreakersApi.md#CreateCircuitBreakers) | **Post** /naming/v1/circuitbreakers | 创建熔断规则
[**DeleteCircuitBreakers**](CircuitBreakersApi.md#DeleteCircuitBreakers) | **Post** /naming/v1/circuitbreakers/delete | 删除熔断规则
[**GetCircuitBreaker**](CircuitBreakersApi.md#GetCircuitBreaker) | **Get** /naming/v1/circuitbreaker | 查询熔断规则
[**GetCircuitBreakerToken**](CircuitBreakersApi.md#GetCircuitBreakerToken) | **Get** /naming/v1/circuitbreaker/token | 查询熔断规则Token
[**GetCircuitBreakerVersions**](CircuitBreakersApi.md#GetCircuitBreakerVersions) | **Get** /naming/v1/circuitbreaker/versions | 查询熔断规则版本
[**GetMasterCircuitBreakers**](CircuitBreakersApi.md#GetMasterCircuitBreakers) | **Get** /naming/v1/circuitbreakers/master | 查询熔断规则Master版本
[**GetReleaseCircuitBreakers**](CircuitBreakersApi.md#GetReleaseCircuitBreakers) | **Get** /naming/v1/circuitbreakers/release | 根据规则id查询已发布的熔断规则
[**ReleaseCircuitBreakers**](CircuitBreakersApi.md#ReleaseCircuitBreakers) | **Post** /naming/v1/circuitbreakers/release | 发布熔断规则
[**UnBindCircuitBreakers**](CircuitBreakersApi.md#UnBindCircuitBreakers) | **Post** /naming/v1/circuitbreakers/unbind | 解绑熔断规则
[**UpdateCircuitBreakers**](CircuitBreakersApi.md#UpdateCircuitBreakers) | **Put** /naming/v1/circuitbreakers | 更新熔断规则



## CreateCircuitBreakerVersions

> CreateCircuitBreakerVersions(ctx).Body(body).Execute()

创建熔断规则版本



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
    body := []openapiclient.V1CircuitBreaker{*openapiclient.NewV1CircuitBreaker()} // []V1CircuitBreaker | create circuit breaker versions

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.CreateCircuitBreakerVersions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.CreateCircuitBreakerVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCircuitBreakerVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1CircuitBreaker**](V1CircuitBreaker.md) | create circuit breaker versions | 

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


## CreateCircuitBreakers

> CreateCircuitBreakers(ctx).Body(body).Execute()

创建熔断规则



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
    body := []openapiclient.V1CircuitBreaker{*openapiclient.NewV1CircuitBreaker()} // []V1CircuitBreaker | create circuit breakers

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.CreateCircuitBreakers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.CreateCircuitBreakers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCircuitBreakersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1CircuitBreaker**](V1CircuitBreaker.md) | create circuit breakers | 

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


## DeleteCircuitBreakers

> DeleteCircuitBreakers(ctx).Body(body).Execute()

删除熔断规则



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
    body := []openapiclient.V1CircuitBreaker{*openapiclient.NewV1CircuitBreaker()} // []V1CircuitBreaker | delete circuit breakers

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.DeleteCircuitBreakers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.DeleteCircuitBreakers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCircuitBreakersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1CircuitBreaker**](V1CircuitBreaker.md) | delete circuit breakers | 

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


## GetCircuitBreaker

> GetCircuitBreaker(ctx, id, version).Execute()

查询熔断规则



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
    version := "version_example" // string | 版本

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.GetCircuitBreaker(context.Background(), id, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.GetCircuitBreaker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 规则ID | 
**version** | **string** | 版本 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCircuitBreakerRequest struct via the builder pattern


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


## GetCircuitBreakerToken

> GetCircuitBreakerToken(ctx).Execute()

查询熔断规则Token

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
    resp, r, err := api_client.CircuitBreakersApi.GetCircuitBreakerToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.GetCircuitBreakerToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCircuitBreakerTokenRequest struct via the builder pattern


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


## GetCircuitBreakerVersions

> GetCircuitBreakerVersions(ctx, id).Execute()

查询熔断规则版本



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.GetCircuitBreakerVersions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.GetCircuitBreakerVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 规则ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCircuitBreakerVersionsRequest struct via the builder pattern


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


## GetMasterCircuitBreakers

> GetMasterCircuitBreakers(ctx, id).Execute()

查询熔断规则Master版本



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.GetMasterCircuitBreakers(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.GetMasterCircuitBreakers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 规则ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMasterCircuitBreakersRequest struct via the builder pattern


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


## GetReleaseCircuitBreakers

> GetReleaseCircuitBreakers(ctx, id).Execute()

根据规则id查询已发布的熔断规则



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.GetReleaseCircuitBreakers(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.GetReleaseCircuitBreakers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | 规则ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReleaseCircuitBreakersRequest struct via the builder pattern


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


## ReleaseCircuitBreakers

> ReleaseCircuitBreakers(ctx).Body(body).Execute()

发布熔断规则



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
    body := []openapiclient.V1ConfigRelease{*openapiclient.NewV1ConfigRelease()} // []V1ConfigRelease | release circuit breakers

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.ReleaseCircuitBreakers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.ReleaseCircuitBreakers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReleaseCircuitBreakersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1ConfigRelease**](V1ConfigRelease.md) | release circuit breakers | 

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


## UnBindCircuitBreakers

> UnBindCircuitBreakers(ctx).Body(body).Execute()

解绑熔断规则



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
    body := []openapiclient.V1ConfigRelease{*openapiclient.NewV1ConfigRelease()} // []V1ConfigRelease | unbind circuit breakers

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.UnBindCircuitBreakers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.UnBindCircuitBreakers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnBindCircuitBreakersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1ConfigRelease**](V1ConfigRelease.md) | unbind circuit breakers | 

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


## UpdateCircuitBreakers

> UpdateCircuitBreakers(ctx).Body(body).Execute()

更新熔断规则



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
    body := []openapiclient.V1CircuitBreaker{*openapiclient.NewV1CircuitBreaker()} // []V1CircuitBreaker | update circuit breakers

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CircuitBreakersApi.UpdateCircuitBreakers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CircuitBreakersApi.UpdateCircuitBreakers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCircuitBreakersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1CircuitBreaker**](V1CircuitBreaker.md) | update circuit breakers | 

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

