# \ConfigConsoleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteConfigFile**](ConfigConsoleApi.md#BatchDeleteConfigFile) | **Post** /config/v1/configfiles/batchdelete | 批量删除配置文件
[**CreateConfigFile**](ConfigConsoleApi.md#CreateConfigFile) | **Post** /config/v1/configfiles | 创建配置文件
[**CreateConfigFileGroup**](ConfigConsoleApi.md#CreateConfigFileGroup) | **Post** /config/v1/configfilegroups | 创建配置文件组
[**CreateConfigFileTemplate**](ConfigConsoleApi.md#CreateConfigFileTemplate) | **Post** /config/v1/configfiletemplates | 创建配置模板
[**DeleteConfigFile**](ConfigConsoleApi.md#DeleteConfigFile) | **Delete** /config/v1/configfiles | 创建配置文件
[**DeleteConfigFileGroup**](ConfigConsoleApi.md#DeleteConfigFileGroup) | **Delete** /config/v1/configfilegroups | 删除配置文件
[**GetAllConfigFileTemplates**](ConfigConsoleApi.md#GetAllConfigFileTemplates) | **Get** /config/v1/configfiletemplates | 获取配置模板
[**GetConfigFile**](ConfigConsoleApi.md#GetConfigFile) | **Get** /config/v1/configfiles | 拉取配置
[**GetConfigFileRelease**](ConfigConsoleApi.md#GetConfigFileRelease) | **Get** /config/v1/configfiles/release | 获取配置文件最后一次全量发布信息
[**GetConfigFileReleaseHistory**](ConfigConsoleApi.md#GetConfigFileReleaseHistory) | **Get** /config/v1/configfiles/releasehistory | 获取配置文件发布历史记录
[**PublishConfigFile**](ConfigConsoleApi.md#PublishConfigFile) | **Post** /config/v1/configfiles/release | 发布配置文件
[**QueryConfigFileGroups**](ConfigConsoleApi.md#QueryConfigFileGroups) | **Get** /config/v1/configfilegroups | 搜索配置文件组
[**QueryConfigFilesByGroup**](ConfigConsoleApi.md#QueryConfigFilesByGroup) | **Get** /config/v1/configfiles/by-group | 搜索配置文件
[**SearchConfigFile**](ConfigConsoleApi.md#SearchConfigFile) | **Get** /config/v1/configfiles/search | 搜索配置文件
[**UpdateConfigFile**](ConfigConsoleApi.md#UpdateConfigFile) | **Put** /config/v1/configfiles | 创建配置文件
[**UpdateConfigFileGroup**](ConfigConsoleApi.md#UpdateConfigFileGroup) | **Put** /config/v1/configfilegroups | 更新配置文件组



## BatchDeleteConfigFile

> BatchDeleteConfigFile(ctx).Body(body).DeleteBy(deleteBy).Execute()

批量删除配置文件

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
    body := *openapiclient.NewV1ConfigFile() // V1ConfigFile | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据} ```[      {          \"name\":\"application.properties\",          \"namespace\":\"someNamespace\",          \"group\":\"someGroup\"      } ] ```
    deleteBy := "deleteBy_example" // string | 操作人 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.BatchDeleteConfigFile(context.Background()).Body(body).DeleteBy(deleteBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.BatchDeleteConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchDeleteConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ConfigFile**](V1ConfigFile.md) | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据} &#x60;&#x60;&#x60;[      {          \&quot;name\&quot;:\&quot;application.properties\&quot;,          \&quot;namespace\&quot;:\&quot;someNamespace\&quot;,          \&quot;group\&quot;:\&quot;someGroup\&quot;      } ] &#x60;&#x60;&#x60; | 
 **deleteBy** | **string** | 操作人 | 

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


## CreateConfigFile

> CreateConfigFile(ctx).Body(body).Execute()

创建配置文件

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
    body := *openapiclient.NewV1ConfigFile() // V1ConfigFile | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据}  ```{     \"name\":\"application.properties\",     \"namespace\":\"someNamespace\",     \"group\":\"someGroup\",     \"content\":\"redis.cache.age=10\",     \"comment\":\"第一个配置文件\",     \"tags\":[{\"key\":\"service\", \"value\":\"helloService\"}],     \"createBy\":\"ledou\",     \"format\":\"properties\" } ``` 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.CreateConfigFile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.CreateConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ConfigFile**](V1ConfigFile.md) | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据}  &#x60;&#x60;&#x60;{     \&quot;name\&quot;:\&quot;application.properties\&quot;,     \&quot;namespace\&quot;:\&quot;someNamespace\&quot;,     \&quot;group\&quot;:\&quot;someGroup\&quot;,     \&quot;content\&quot;:\&quot;redis.cache.age&#x3D;10\&quot;,     \&quot;comment\&quot;:\&quot;第一个配置文件\&quot;,     \&quot;tags\&quot;:[{\&quot;key\&quot;:\&quot;service\&quot;, \&quot;value\&quot;:\&quot;helloService\&quot;}],     \&quot;createBy\&quot;:\&quot;ledou\&quot;,     \&quot;format\&quot;:\&quot;properties\&quot; } &#x60;&#x60;&#x60;  | 

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


## CreateConfigFileGroup

> CreateConfigFileGroup(ctx).Body(body).Execute()

创建配置文件组

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
    body := *openapiclient.NewV1ConfigFileGroup() // V1ConfigFileGroup | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据}  ``` {     \"name\":\"someGroup\",     \"namespace\":\"someNamespace\",     \"comment\":\"some comment\",     \"createBy\":\"ledou\" } ```

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.CreateConfigFileGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.CreateConfigFileGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigFileGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ConfigFileGroup**](V1ConfigFileGroup.md) | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据}  &#x60;&#x60;&#x60; {     \&quot;name\&quot;:\&quot;someGroup\&quot;,     \&quot;namespace\&quot;:\&quot;someNamespace\&quot;,     \&quot;comment\&quot;:\&quot;some comment\&quot;,     \&quot;createBy\&quot;:\&quot;ledou\&quot; } &#x60;&#x60;&#x60; | 

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


## CreateConfigFileTemplate

> CreateConfigFileTemplate(ctx).Execute()

创建配置模板

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
    resp, r, err := api_client.ConfigConsoleApi.CreateConfigFileTemplate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.CreateConfigFileTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigFileTemplateRequest struct via the builder pattern


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


## DeleteConfigFile

> DeleteConfigFile(ctx).Namespace(namespace).Group(group).Name(name).DeleteBy(deleteBy).Execute()

创建配置文件

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
    name := "name_example" // string | 配置文件
    deleteBy := "deleteBy_example" // string | 操作人 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.DeleteConfigFile(context.Background()).Namespace(namespace).Group(group).Name(name).DeleteBy(deleteBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.DeleteConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | 命名空间 | 
 **group** | **string** | 配置文件分组 | 
 **name** | **string** | 配置文件 | 
 **deleteBy** | **string** | 操作人 | 

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


## DeleteConfigFileGroup

> DeleteConfigFileGroup(ctx).Namespace(namespace).Group(group).Execute()

删除配置文件

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.DeleteConfigFileGroup(context.Background()).Namespace(namespace).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.DeleteConfigFileGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigFileGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | 命名空间 | 
 **group** | **string** | 配置文件分组 | 

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


## GetAllConfigFileTemplates

> GetAllConfigFileTemplates(ctx).Execute()

获取配置模板

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
    resp, r, err := api_client.ConfigConsoleApi.GetAllConfigFileTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.GetAllConfigFileTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConfigFileTemplatesRequest struct via the builder pattern


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


## GetConfigFile

> GetConfigFile(ctx).Namespace(namespace).Group(group).Name(name).Execute()

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
    name := "name_example" // string | 配置文件名

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.GetConfigFile(context.Background()).Namespace(namespace).Group(group).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.GetConfigFile``: %v\n", err)
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
 **name** | **string** | 配置文件名 | 

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


## GetConfigFileRelease

> GetConfigFileRelease(ctx).Namespace(namespace).Group(group).Name(name).Execute()

获取配置文件最后一次全量发布信息

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
    name := "name_example" // string | 配置文件

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.GetConfigFileRelease(context.Background()).Namespace(namespace).Group(group).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.GetConfigFileRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigFileReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | 命名空间 | 
 **group** | **string** | 配置文件分组 | 
 **name** | **string** | 配置文件 | 

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


## GetConfigFileReleaseHistory

> GetConfigFileReleaseHistory(ctx).Namespace(namespace).Limit(limit).Group(group).Name(name).Offset(offset).Execute()

获取配置文件发布历史记录

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
    limit := int32(56) // int32 | 一页大小，最大为 100 (default to 100)
    group := "group_example" // string | 配置文件分组 (optional)
    name := "name_example" // string | 配置文件 (optional)
    offset := int32(56) // int32 | 翻页偏移量 默认为 0 (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.GetConfigFileReleaseHistory(context.Background()).Namespace(namespace).Limit(limit).Group(group).Name(name).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.GetConfigFileReleaseHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigFileReleaseHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | 命名空间 | 
 **limit** | **int32** | 一页大小，最大为 100 | [default to 100]
 **group** | **string** | 配置文件分组 | 
 **name** | **string** | 配置文件 | 
 **offset** | **int32** | 翻页偏移量 默认为 0 | [default to 0]

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


## PublishConfigFile

> PublishConfigFile(ctx).Body(body).Execute()

发布配置文件

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
    body := *openapiclient.NewV1ConfigFileRelease() // V1ConfigFileRelease | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据} ```{     \"name\":\"release-002\",     \"fileName\":\"application.properties\",     \"namespace\":\"someNamespace\",     \"group\":\"someGroup\",     \"comment\":\"发布第一个配置文件\",     \"createBy\":\"ledou\" } ```

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.PublishConfigFile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.PublishConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPublishConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ConfigFileRelease**](V1ConfigFileRelease.md) | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据} &#x60;&#x60;&#x60;{     \&quot;name\&quot;:\&quot;release-002\&quot;,     \&quot;fileName\&quot;:\&quot;application.properties\&quot;,     \&quot;namespace\&quot;:\&quot;someNamespace\&quot;,     \&quot;group\&quot;:\&quot;someGroup\&quot;,     \&quot;comment\&quot;:\&quot;发布第一个配置文件\&quot;,     \&quot;createBy\&quot;:\&quot;ledou\&quot; } &#x60;&#x60;&#x60; | 

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


## QueryConfigFileGroups

> QueryConfigFileGroups(ctx).Limit(limit).Namespace(namespace).Group(group).FileName(fileName).Offset(offset).Execute()

搜索配置文件组

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
    limit := int32(56) // int32 | 一页大小，最大为 100 (default to 100)
    namespace := "namespace_example" // string | 命名空间，不填表示全部命名空间 (optional)
    group := "group_example" // string | 配置文件分组名，模糊搜索 (optional)
    fileName := "fileName_example" // string | 配置文件名称，模糊搜索 (optional)
    offset := int32(56) // int32 | 翻页偏移量 默认为 0 (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.QueryConfigFileGroups(context.Background()).Limit(limit).Namespace(namespace).Group(group).FileName(fileName).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.QueryConfigFileGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryConfigFileGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | 一页大小，最大为 100 | [default to 100]
 **namespace** | **string** | 命名空间，不填表示全部命名空间 | 
 **group** | **string** | 配置文件分组名，模糊搜索 | 
 **fileName** | **string** | 配置文件名称，模糊搜索 | 
 **offset** | **int32** | 翻页偏移量 默认为 0 | [default to 0]

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


## QueryConfigFilesByGroup

> QueryConfigFilesByGroup(ctx).Limit(limit).Namespace(namespace).Group(group).Offset(offset).Execute()

搜索配置文件

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
    limit := int32(56) // int32 | 一页大小，最大为 100 (default to 100)
    namespace := "namespace_example" // string | 命名空间 (optional)
    group := "group_example" // string | 配置文件分组 (optional)
    offset := int32(56) // int32 | 翻页偏移量 默认为 0 (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.QueryConfigFilesByGroup(context.Background()).Limit(limit).Namespace(namespace).Group(group).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.QueryConfigFilesByGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryConfigFilesByGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | 一页大小，最大为 100 | [default to 100]
 **namespace** | **string** | 命名空间 | 
 **group** | **string** | 配置文件分组 | 
 **offset** | **int32** | 翻页偏移量 默认为 0 | [default to 0]

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


## SearchConfigFile

> SearchConfigFile(ctx).Limit(limit).Namespace(namespace).Group(group).Name(name).Tags(tags).Offset(offset).Execute()

搜索配置文件

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
    limit := int32(56) // int32 | 一页大小，最大为 100 (default to 100)
    namespace := "namespace_example" // string | 命名空间 (optional)
    group := "group_example" // string | 配置文件分组 (optional)
    name := "name_example" // string | 配置文件 (optional)
    tags := "tags_example" // string | 格式：key1,value1,key2,value2 (optional)
    offset := int32(56) // int32 | 翻页偏移量 默认为 0 (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.SearchConfigFile(context.Background()).Limit(limit).Namespace(namespace).Group(group).Name(name).Tags(tags).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.SearchConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | 一页大小，最大为 100 | [default to 100]
 **namespace** | **string** | 命名空间 | 
 **group** | **string** | 配置文件分组 | 
 **name** | **string** | 配置文件 | 
 **tags** | **string** | 格式：key1,value1,key2,value2 | 
 **offset** | **int32** | 翻页偏移量 默认为 0 | [default to 0]

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


## UpdateConfigFile

> UpdateConfigFile(ctx).Body(body).Execute()

创建配置文件

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
    body := *openapiclient.NewV1ConfigFile() // V1ConfigFile | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据}  ```{     \"name\":\"application.properties\",     \"namespace\":\"someNamespace\",     \"group\":\"someGroup\",     \"content\":\"redis.cache.age=10\",     \"comment\":\"第一个配置文件\",     \"tags\":[{\"key\":\"service\", \"value\":\"helloService\"}],     \"createBy\":\"ledou\",     \"format\":\"properties\" } ``` 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.UpdateConfigFile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.UpdateConfigFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ConfigFile**](V1ConfigFile.md) | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据}  &#x60;&#x60;&#x60;{     \&quot;name\&quot;:\&quot;application.properties\&quot;,     \&quot;namespace\&quot;:\&quot;someNamespace\&quot;,     \&quot;group\&quot;:\&quot;someGroup\&quot;,     \&quot;content\&quot;:\&quot;redis.cache.age&#x3D;10\&quot;,     \&quot;comment\&quot;:\&quot;第一个配置文件\&quot;,     \&quot;tags\&quot;:[{\&quot;key\&quot;:\&quot;service\&quot;, \&quot;value\&quot;:\&quot;helloService\&quot;}],     \&quot;createBy\&quot;:\&quot;ledou\&quot;,     \&quot;format\&quot;:\&quot;properties\&quot; } &#x60;&#x60;&#x60;  | 

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


## UpdateConfigFileGroup

> UpdateConfigFileGroup(ctx).Body(body).Execute()

更新配置文件组

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
    body := *openapiclient.NewV1ConfigFileGroup() // V1ConfigFileGroup | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据}  ``` {     \"name\":\"someGroup\",     \"namespace\":\"someNamespace\",     \"comment\":\"some comment\",     \"createBy\":\"ledou\" } ```

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConfigConsoleApi.UpdateConfigFileGroup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigConsoleApi.UpdateConfigFileGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigFileGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1ConfigFileGroup**](V1ConfigFileGroup.md) | 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header Header X-Polaris-Token: {访问凭据}  &#x60;&#x60;&#x60; {     \&quot;name\&quot;:\&quot;someGroup\&quot;,     \&quot;namespace\&quot;:\&quot;someNamespace\&quot;,     \&quot;comment\&quot;:\&quot;some comment\&quot;,     \&quot;createBy\&quot;:\&quot;ledou\&quot; } &#x60;&#x60;&#x60; | 

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

