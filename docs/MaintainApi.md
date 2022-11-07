# \MaintainApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanInstance**](MaintainApi.md#CleanInstance) | **Post** /maintain/v1/instance/clean | 彻底清理flag&#x3D;1的实例
[**CloseConnections**](MaintainApi.md#CloseConnections) | **Post** /maintain/v1/apiserver/conn/close | 关闭指定client ip的连接
[**FreeOSMemory**](MaintainApi.md#FreeOSMemory) | **Post** /maintain/v1/memory/free | 释放系统内存
[**GetLastHeartbeat**](MaintainApi.md#GetLastHeartbeat) | **Get** /maintain/v1/instance/heartbeat | 获取上一次心跳的时间
[**GetLogOutputLevel**](MaintainApi.md#GetLogOutputLevel) | **Get** /maintain/v1/log/outputlevel | 获取日志输出级别
[**GetServerConnStats**](MaintainApi.md#GetServerConnStats) | **Get** /maintain/v1/apiserver/conn/stats | 获取服务端连接统计信息
[**GetServerConnections**](MaintainApi.md#GetServerConnections) | **Get** /maintain/v1/apiserver/conn | 获取服务端连接数
[**SetLogOutputLevel**](MaintainApi.md#SetLogOutputLevel) | **Put** /maintain/v1/log/outputlevel | 设置日志输出级别



## CleanInstance

> CleanInstance(ctx).Body(body).Execute()

彻底清理flag=1的实例



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
    body := *openapiclient.NewV1Instance() // V1Instance | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintainApi.CleanInstance(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainApi.CleanInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCleanInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1Instance**](V1Instance.md) |  | 

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


## CloseConnections

> CloseConnections(ctx).Body(body).Execute()

关闭指定client ip的连接



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
    body := []openapiclient.MaintainConnReq{*openapiclient.NewMaintainConnReq(int32(123), "Host_example", int32(123), "Protocol_example")} // []MaintainConnReq | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintainApi.CloseConnections(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainApi.CloseConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]MaintainConnReq**](MaintainConnReq.md) |  | 

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


## FreeOSMemory

> FreeOSMemory(ctx).Execute()

释放系统内存



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
    resp, r, err := api_client.MaintainApi.FreeOSMemory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainApi.FreeOSMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFreeOSMemoryRequest struct via the builder pattern


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


## GetLastHeartbeat

> GetLastHeartbeat(ctx).Id(id).Service(service).Namespace(namespace).Host(host).Port(port).VpvId(vpvId).Execute()

获取上一次心跳的时间



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
    id := "id_example" // string | 实例ID 如果存在则其它参数可不填 (optional)
    service := "service_example" // string | 服务名 (optional)
    namespace := "namespace_example" // string | 命名空间 (optional)
    host := "host_example" // string | 主机名 (optional)
    port := int32(56) // int32 | 端口 (optional)
    vpvId := "vpvId_example" // string | VPC ID (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintainApi.GetLastHeartbeat(context.Background()).Id(id).Service(service).Namespace(namespace).Host(host).Port(port).VpvId(vpvId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainApi.GetLastHeartbeat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLastHeartbeatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 实例ID 如果存在则其它参数可不填 | 
 **service** | **string** | 服务名 | 
 **namespace** | **string** | 命名空间 | 
 **host** | **string** | 主机名 | 
 **port** | **int32** | 端口 | 
 **vpvId** | **string** | VPC ID | 

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


## GetLogOutputLevel

> GetLogOutputLevel(ctx).Execute()

获取日志输出级别



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
    resp, r, err := api_client.MaintainApi.GetLogOutputLevel(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainApi.GetLogOutputLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogOutputLevelRequest struct via the builder pattern


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


## GetServerConnStats

> GetServerConnStats(ctx).Protocol(protocol).Host(host).Amount(amount).Execute()

获取服务端连接统计信息



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
    protocol := "protocol_example" // string | 查看指定协议
    host := "host_example" // string | 查看指定host (optional)
    amount := int32(56) // int32 | 总数 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintainApi.GetServerConnStats(context.Background()).Protocol(protocol).Host(host).Amount(amount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainApi.GetServerConnStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerConnStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protocol** | **string** | 查看指定协议 | 
 **host** | **string** | 查看指定host | 
 **amount** | **int32** | 总数 | 

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


## GetServerConnections

> GetServerConnections(ctx).Protocol(protocol).Host(host).Execute()

获取服务端连接数



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
    protocol := "protocol_example" // string | 查看指定协议
    host := "host_example" // string | 查看指定host (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MaintainApi.GetServerConnections(context.Background()).Protocol(protocol).Host(host).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainApi.GetServerConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServerConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **protocol** | **string** | 查看指定协议 | 
 **host** | **string** | 查看指定host | 

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


## SetLogOutputLevel

> SetLogOutputLevel(ctx).Execute()

设置日志输出级别



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
    resp, r, err := api_client.MaintainApi.SetLogOutputLevel(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintainApi.SetLogOutputLevel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSetLogOutputLevelRequest struct via the builder pattern


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

