# \InstancesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstances**](InstancesApi.md#CreateInstances) | **Post** /naming/v1/instances | 创建实例
[**DeleteInstances**](InstancesApi.md#DeleteInstances) | **Post** /naming/v1/instances/delete | 删除实例(根据实例ID)
[**DeleteInstancesByHost**](InstancesApi.md#DeleteInstancesByHost) | **Post** /naming/v1/instances/delete/host | 删除实例(根据主机)
[**GetInstanceLabels**](InstancesApi.md#GetInstanceLabels) | **Get** /naming/v1/instances/labels | 查询某个服务下所有实例的标签信息
[**GetInstances**](InstancesApi.md#GetInstances) | **Get** /naming/v1/instances | 查询服务实例
[**GetInstancesCount**](InstancesApi.md#GetInstancesCount) | **Get** /naming/v1/instances/count | 查询服务实例数量
[**UpdateInstances**](InstancesApi.md#UpdateInstances) | **Put** /naming/v1/instances | 更新实例
[**UpdateInstancesIsolate**](InstancesApi.md#UpdateInstancesIsolate) | **Put** /naming/v1/instances/isolate/host | 修改服务实例的隔离状态



## CreateInstances

> CreateInstances(ctx).Body(body).Execute()

创建实例



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
    body := []openapiclient.V1Instance{*openapiclient.NewV1Instance()} // []V1Instance | create instances

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.CreateInstances(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.CreateInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Instance**](V1Instance.md) | create instances | 

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


## DeleteInstances

> DeleteInstances(ctx).Body(body).Execute()

删除实例(根据实例ID)



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
    body := []openapiclient.V1Instance{*openapiclient.NewV1Instance()} // []V1Instance | delete instances

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.DeleteInstances(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Instance**](V1Instance.md) | delete instances | 

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


## DeleteInstancesByHost

> DeleteInstancesByHost(ctx).Body(body).Execute()

删除实例(根据主机)



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
    body := []openapiclient.V1Instance{*openapiclient.NewV1Instance()} // []V1Instance | delete instances

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.DeleteInstancesByHost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.DeleteInstancesByHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstancesByHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Instance**](V1Instance.md) | delete instances | 

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


## GetInstanceLabels

> GetInstanceLabels(ctx).Service(service).Namespace(namespace).Execute()

查询某个服务下所有实例的标签信息



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetInstanceLabels(context.Background()).Service(service).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstanceLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **service** | **string** | 服务名称 | 
 **namespace** | **string** | 命名空间 | 

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


## GetInstances

> GetInstances(ctx, service, namespace, host, keys, values, healthy, isolate, protocol, version, cmdbRegion, cmdbZone, cmdbIdc, offset, limit).Execute()

查询服务实例



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
    host := "host_example" // string | 实例IP
    keys := "keys_example" // string | 标签key
    values := "values_example" // string | 标签value
    healthy := "healthy_example" // string | 实例健康状态
    isolate := "isolate_example" // string | 实例隔离状态
    protocol := "protocol_example" // string | 实例端口协议状态
    version := "version_example" // string | 实例版本
    cmdbRegion := "cmdbRegion_example" // string | 实例region信息
    cmdbZone := "cmdbZone_example" // string | 实例zone信息
    cmdbIdc := "cmdbIdc_example" // string | 实例idc信息
    offset := int32(56) // int32 | 查询偏移量
    limit := int32(56) // int32 | 查询条数

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.GetInstances(context.Background(), service, namespace, host, keys, values, healthy, isolate, protocol, version, cmdbRegion, cmdbZone, cmdbIdc, offset, limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstances``: %v\n", err)
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
**host** | **string** | 实例IP | 
**keys** | **string** | 标签key | 
**values** | **string** | 标签value | 
**healthy** | **string** | 实例健康状态 | 
**isolate** | **string** | 实例隔离状态 | 
**protocol** | **string** | 实例端口协议状态 | 
**version** | **string** | 实例版本 | 
**cmdbRegion** | **string** | 实例region信息 | 
**cmdbZone** | **string** | 实例zone信息 | 
**cmdbIdc** | **string** | 实例idc信息 | 
**offset** | **int32** | 查询偏移量 | 
**limit** | **int32** | 查询条数 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancesRequest struct via the builder pattern


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


## GetInstancesCount

> GetInstancesCount(ctx).Execute()

查询服务实例数量



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
    resp, r, err := api_client.InstancesApi.GetInstancesCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.GetInstancesCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancesCountRequest struct via the builder pattern


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


## UpdateInstances

> UpdateInstances(ctx).Body(body).Execute()

更新实例



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
    body := []openapiclient.V1Instance{*openapiclient.NewV1Instance()} // []V1Instance | update instances

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.UpdateInstances(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UpdateInstances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Instance**](V1Instance.md) | update instances | 

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


## UpdateInstancesIsolate

> UpdateInstancesIsolate(ctx).Body(body).Execute()

修改服务实例的隔离状态



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
    body := []openapiclient.V1Instance{*openapiclient.NewV1Instance()} // []V1Instance | update instances

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InstancesApi.UpdateInstancesIsolate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstancesApi.UpdateInstancesIsolate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstancesIsolateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]V1Instance**](V1Instance.md) | update instances | 

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

