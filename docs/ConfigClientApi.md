# \ConfigClientApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigFile**](ConfigClientApi.md#GetConfigFile) | **Get** /config/v1/GetConfigFile | 拉取配置
[**WatchConfigFile**](ConfigClientApi.md#WatchConfigFile) | **Post** /config/v1/WatchConfigFile | 监听配置



## GetConfigFile

> GetConfigFile(ctx).Namespace(namespace).Group(group).FileName(fileName).Version(version).Execute()

拉取配置

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
    namespace := "namespace_example" // string | 命名空间
    group := "group_example" // string | 配置文件分组
    fileName := "fileName_example" // string | 配置文件名
    version := int32(56) // int32 | 配置文件客户端版本号，刚启动时设置为 0

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigClientApi.GetConfigFile(context.Background()).Namespace(namespace).Group(group).FileName(fileName).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigClientApi.GetConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | 命名空间 | 
 **group** | **string** | 配置文件分组 | 
 **fileName** | **string** | 配置文件名 | 
 **version** | **int32** | 配置文件客户端版本号，刚启动时设置为 0 | 

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


## WatchConfigFile

> WatchConfigFile(ctx).Body(body).Execute()

监听配置

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
    body := *openapiclient.NewV1ClientWatchConfigFileRequest() // V1ClientWatchConfigFileRequest | 通过 Http LongPolling 机制订阅配置变更。

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigClientApi.WatchConfigFile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigClientApi.WatchConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatchConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ClientWatchConfigFileRequest**](V1ClientWatchConfigFileRequest.md) | 通过 Http LongPolling 机制订阅配置变更。 | 

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

