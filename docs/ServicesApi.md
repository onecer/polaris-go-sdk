# \ServicesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAlias**](ServicesApi.md#CreateServiceAlias) | **Post** /naming/v1/service/alias | 创建服务别名
[**CreateServices**](ServicesApi.md#CreateServices) | **Post** /naming/v1/services | 创建服务
[**DeleteServiceAliases**](ServicesApi.md#DeleteServiceAliases) | **Post** /naming/v1/service/aliases/delete | 删除服务别名
[**DeleteServices**](ServicesApi.md#DeleteServices) | **Post** /naming/v1/services/delete | 删除服务
[**GetCircuitBreakerByService**](ServicesApi.md#GetCircuitBreakerByService) | **Get** /naming/v1/service/circuitbreaker | 根据服务查询熔断规则
[**GetServiceAliases**](ServicesApi.md#GetServiceAliases) | **Get** /naming/v1/service/aliases | 查询服务别名
[**GetServiceOwner**](ServicesApi.md#GetServiceOwner) | **Post** /naming/v1/service/owner | 根据服务获取服务负责人
[**GetServiceToken**](ServicesApi.md#GetServiceToken) | **Get** /naming/v1/service/token | 查询服务Token
[**GetServices**](ServicesApi.md#GetServices) | **Get** /naming/v1/services | 获取服务列表
[**GetServicesCount**](ServicesApi.md#GetServicesCount) | **Get** /naming/v1/services/count | 获取服务数量
[**UpdateServiceAlias**](ServicesApi.md#UpdateServiceAlias) | **Put** /naming/v1/service/alias | 更新服务别名
[**UpdateServiceToken**](ServicesApi.md#UpdateServiceToken) | **Put** /naming/v1/service/token | 更新服务Token
[**UpdateServices**](ServicesApi.md#UpdateServices) | **Put** /naming/v1/services | 更新服务



## CreateServiceAlias

> CreateServiceAlias(ctx).Body(body).Execute()

创建服务别名



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
    body := *openapiclient.NewV1ServiceAlias() // V1ServiceAlias | create service alias

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.CreateServiceAlias(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.CreateServiceAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ServiceAlias**](V1ServiceAlias.md) | create service alias | 

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


## CreateServices

> CreateServices(ctx).Body(body).Execute()

创建服务



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
    body := []openapiclient.V1Service{*openapiclient.NewV1Service()} // []V1Service | create services

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.CreateServices(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.CreateServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Service**](V1Service.md) | create services | 

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


## DeleteServiceAliases

> DeleteServiceAliases(ctx).Body(body).Execute()

删除服务别名



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
    body := []openapiclient.V1ServiceAlias{*openapiclient.NewV1ServiceAlias()} // []V1ServiceAlias | delete service aliases

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.DeleteServiceAliases(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeleteServiceAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1ServiceAlias**](V1ServiceAlias.md) | delete service aliases | 

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


## DeleteServices

> DeleteServices(ctx).Body(body).Execute()

删除服务



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
    body := []openapiclient.V1Service{*openapiclient.NewV1Service()} // []V1Service | delete services

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.DeleteServices(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.DeleteServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Service**](V1Service.md) | delete services | 

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


## GetCircuitBreakerByService

> GetCircuitBreakerByService(ctx, service, namespace).Execute()

根据服务查询熔断规则



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
    service := "service_example" // string | 服务名
    namespace := "namespace_example" // string | 命名空间

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.GetCircuitBreakerByService(context.Background(), service, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetCircuitBreakerByService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**service** | **string** | 服务名 | 
**namespace** | **string** | 命名空间 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCircuitBreakerByServiceRequest struct via the builder pattern


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


## GetServiceAliases

> GetServiceAliases(ctx).Execute()

查询服务别名



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
    resp, r, err := api_client.ServicesApi.GetServiceAliases(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServiceAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAliasesRequest struct via the builder pattern


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


## GetServiceOwner

> GetServiceOwner(ctx).Execute()

根据服务获取服务负责人

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
    resp, r, err := api_client.ServicesApi.GetServiceOwner(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServiceOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceOwnerRequest struct via the builder pattern


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


## GetServiceToken

> GetServiceToken(ctx).Execute()

查询服务Token

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
    resp, r, err := api_client.ServicesApi.GetServiceToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServiceToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceTokenRequest struct via the builder pattern


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


## GetServices

> GetServices(ctx).Name(name).Namespace(namespace).Business(business).Department(department).Host(host).Port(port).Keys(keys).Values(values).Offset(offset).Limit(limit).Execute()

获取服务列表



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
    name := "name_example" // string | 服务名 (optional) (default to "demo-service")
    namespace := "namespace_example" // string | 命名空间 (optional) (default to "default")
    business := "business_example" // string | 业务，默认模糊查询 (optional)
    department := "department_example" // string | 部门 (optional)
    host := "host_example" // string | 实例IP，**多个IP以英文逗号分隔** (optional)
    port := "port_example" // string | **实例端口**，**多个端口以英文逗号分隔**  (optional)
    keys := "keys_example" // string | 服务元数据名，keys和values需要同时填写，目前只支持查询一组元数据。 (optional)
    values := "values_example" // string | 服务元数据名，keys和values需要同时填写，目前只支持查询一组元数据。 (optional)
    offset := int32(56) // int32 | 查询偏移量 (optional) (default to 0)
    limit := int32(56) // int32 | 查询条数，**最多查询100条** (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.GetServices(context.Background()).Name(name).Namespace(namespace).Business(business).Department(department).Host(host).Port(port).Keys(keys).Values(values).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 服务名 | [default to &quot;demo-service&quot;]
 **namespace** | **string** | 命名空间 | [default to &quot;default&quot;]
 **business** | **string** | 业务，默认模糊查询 | 
 **department** | **string** | 部门 | 
 **host** | **string** | 实例IP，**多个IP以英文逗号分隔** | 
 **port** | **string** | **实例端口**，**多个端口以英文逗号分隔**  | 
 **keys** | **string** | 服务元数据名，keys和values需要同时填写，目前只支持查询一组元数据。 | 
 **values** | **string** | 服务元数据名，keys和values需要同时填写，目前只支持查询一组元数据。 | 
 **offset** | **int32** | 查询偏移量 | [default to 0]
 **limit** | **int32** | 查询条数，**最多查询100条** | 

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


## GetServicesCount

> GetServicesCount(ctx).Execute()

获取服务数量



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
    resp, r, err := api_client.ServicesApi.GetServicesCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.GetServicesCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServicesCountRequest struct via the builder pattern


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


## UpdateServiceAlias

> UpdateServiceAlias(ctx).Body(body).Execute()

更新服务别名



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
    body := *openapiclient.NewV1ServiceAlias() // V1ServiceAlias | update service alias

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.UpdateServiceAlias(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateServiceAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ServiceAlias**](V1ServiceAlias.md) | update service alias | 

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


## UpdateServiceToken

> UpdateServiceToken(ctx).Execute()

更新服务Token

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
    resp, r, err := api_client.ServicesApi.UpdateServiceToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateServiceToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceTokenRequest struct via the builder pattern


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


## UpdateServices

> UpdateServices(ctx).Body(body).Execute()

更新服务



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
    body := []openapiclient.V1Service{*openapiclient.NewV1Service()} // []V1Service | update services

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServicesApi.UpdateServices(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServicesApi.UpdateServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Service**](V1Service.md) | update services | 

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

