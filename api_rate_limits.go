/*
Polaris Server

一个支持多语言、多框架的云原生服务发现和治理中心  提供高性能SDK和无侵入Sidecar两种接入方式  

API version: v0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package polaris

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// RateLimitsApiService RateLimitsApi service
type RateLimitsApiService service

type ApiCreateRateLimitsRequest struct {
	ctx _context.Context
	ApiService *RateLimitsApiService
	body *[]V1RateLimit
}

// create rate limits
func (r ApiCreateRateLimitsRequest) Body(body []V1RateLimit) ApiCreateRateLimitsRequest {
	r.body = &body
	return r
}

func (r ApiCreateRateLimitsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CreateRateLimitsExecute(r)
}

/*
CreateRateLimits 创建限流规则


为服务创建多个限流规则，以对服务进行流量限制，按优先级顺序进行匹配，匹配到一个则执行该规则。

请求示例：
~~~
POST /naming/v1/ratelimits

# 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header
Header X-Polaris-Token: {访问凭据}

[
	{
		"name": "rule1",
		"service": "testsvc1",
		"namespace": "default",
		"method": {
			"type": "EXACT",
			"value": "/getsomething"
		},
		"arguments": [
			{
				"type": "HEADER",
				"key": "host",
				"value": {
					"type": "EXACT",
					"value": "www.baidu.com"
				}
			},
			{
				"type": "CALLER_SERVICE",
				"key": "default",
				"value": {
					"type": "IN",
					"value": "testsvc1,testsvc2"
				}
			}
		],
		"resource": "QPS",
		"type": "LOCAL",
		"amounts": [
			{
				"maxAmount": 1000,
				"validDuration": "1s"
			},
			{
				"maxAmount": 2000,
				"validDuration": "1m"
			}	
		],
		"regex_combine": false,
		"disable": false,
		"failover": "FAILOVER_LOCAL"
	}
]
~~~
回复示例：
~~~
{
 "code": 200000,
 "info": "execute success",
 "size": 1,
 "responses": [
  {
   "code": 200000,
   "info": "execute success",
   "rateLimit": {
    "id": "e04f201e7b7e4599b42a9b6631a7ba08", //规则ID
    "service": "testsvc1",
    "namespace": "default",
    "name": "rule2"
   }
  }
 ]
}
~~~

数据结构：

> Ratelimit结构参数

| 参数名          | 类型          | 描述                                                         | 是否必填 |
| --------------- | ------------- | ------------------------------------------------------------ | -------- |
| name            | string        | 规则名                                                       | 是       |
| service         | string        | 规则所属的服务名，创建规则时，如果服务不存在，则会自动创建服务。 | 是       |
| namespace       | string        | 规则所属的命名空间                                           | 是       |
| method          | MatchString   | 规则所针对的服务接口                                         | 否       |
| arguments       | MatchArgument | 请求参数匹配条件，需全匹配才通过                             | 否       |
| resource        | string        | 限流资源，默认为QPS(针对QPS进行限流)                         | 否       |
| type            | string        | 限流类型，支持LOCAL（单机限流）, GLOBAL（分布式限流）        | 是       |
| amounts         | Amount[]      | 限流配额，包含限流周期和配额总数，可配置多个                 | 是       |
| regex_combine   | bool          | 合并计算配额，对于匹配到同一条正则表达式规则的多个不同的请求进行合并计算，默认为false | 否       |
| disable         | bool          | 是否启用该限流规则，默认为false（启用）                      | 否       |
| action          | string        | 限流效果，支持REJECT（直接拒绝）,UNIRATE（匀速排队），默认REJECT | 否       |
| failover        | string        | 失败降级措施，仅分布式限流有效，当远程token服务出现故障时，本地如何降级。支持FAILOVER_LOCAL（降级到单机限流），FAILOVER_PASS（直接通过）。默认FAILOVER_LOCAL | 否       |
| max_queue_delay | int           | 最大排队时长，单位秒，仅对匀速排队生效。默认1秒              | 否       |

> Amount结构参数

| 参数名        | 类型   | 描述                                                 | 是否必填 |
| ------------- | ------ | ---------------------------------------------------- | -------- |
| maxAmount     | uint32 | 周期内最大配额数                                     | 是       |
| validDuration | string | 周期描述，支持duration类型的字符串，比如1s, 1m, 1h等 | 是       |

> MatchString结构参数

| 参数名 | 类型   | 描述                                                         | 是否必填 |
| ------ | ------ | ------------------------------------------------------------ | -------- |
| type   | string | 匹配类型，枚举，支持：EXACT（全匹配，默认），REGEX（正则表达式匹配），NOT_EQUALS（不等于），IN（包含），NOT_IN（不包含） | 是       |
| value  | string | 匹配的目标值，如果选择的是包含和不包含，则通过逗号分割多个值 | 是       |

> MatchArgument结构参数

| 参数名 | 类型        | 描述                                                         | 是否必填 |
| ------ | ----------- | ------------------------------------------------------------ | -------- |
| type   | string      | 参数类型，枚举，支持：CUSTOM（自定义，默认），METHOD（方法），HEADER（请求头），QUERY（请求参数），CALLER_SERVICE（主调方服务），CALLER_IP（主调方IP） | 是       |
| key    | string      | 参数键，对于HEADER、QUERY、CUSTOM，对应的是key值；对于CALLER_SERVICE，对应的是服务的命名空间值 | 是       |
| value  | MatchString | 参数值，对于HEADER、QUERY、CUSTOM，对应的是key所关联的value；对于CALLER_SERVICE，对应的是服务名，其他类型则是具体的值，支持多种匹配模式（见MatchString的定义） | 是       |


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateRateLimitsRequest
*/
func (a *RateLimitsApiService) CreateRateLimits(ctx _context.Context) ApiCreateRateLimitsRequest {
	return ApiCreateRateLimitsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RateLimitsApiService) CreateRateLimitsExecute(r ApiCreateRateLimitsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitsApiService.CreateRateLimits")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/naming/v1/ratelimits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Polaris-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteRateLimitsRequest struct {
	ctx _context.Context
	ApiService *RateLimitsApiService
	body *[]V1RateLimit
}

// delete rate limits
func (r ApiDeleteRateLimitsRequest) Body(body []V1RateLimit) ApiDeleteRateLimitsRequest {
	r.body = &body
	return r
}

func (r ApiDeleteRateLimitsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteRateLimitsExecute(r)
}

/*
DeleteRateLimits 删除限流规则


请求示例：

~~~
POST /naming/v1/ratelimits/delete

# 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header
Header X-Polaris-Token: {访问凭据}

[
	{
		"id": "6942526fbac545848cd8fb32a3a55bb6" //规则ID，必填
	}
]
~~~

应答示例：

~~~
{
 "code": 200000,
 "info": "execute success",
 "size": 1,
 "responses": [
  {
   "code": 200000,
   "info": "execute success",
   "rateLimit": {
    "id": "6942526fbac545848cd8fb32a3a55bb6"
   }
  }
 ]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteRateLimitsRequest
*/
func (a *RateLimitsApiService) DeleteRateLimits(ctx _context.Context) ApiDeleteRateLimitsRequest {
	return ApiDeleteRateLimitsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RateLimitsApiService) DeleteRateLimitsExecute(r ApiDeleteRateLimitsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitsApiService.DeleteRateLimits")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/naming/v1/ratelimits/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Polaris-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEnableRateLimitsRequest struct {
	ctx _context.Context
	ApiService *RateLimitsApiService
	body *[]V1RateLimit
}

// enable rate limits
func (r ApiEnableRateLimitsRequest) Body(body []V1RateLimit) ApiEnableRateLimitsRequest {
	r.body = &body
	return r
}

func (r ApiEnableRateLimitsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.EnableRateLimitsExecute(r)
}

/*
EnableRateLimits 启用限流规则


请求示例：

~~~
PUT /naming/v1/ratelimits/enable
[
	{
		"id": "6942526fbac545848cd8fb32a3a55bb6", //规则ID，必填
		"disable": true // 是否禁用，true为不启用，false为启用
	}
]
~~~

应答示例：

~~~
{
 "code": 200000,
 "info": "execute success",
 "size": 1,
 "responses": [
  {
   "code": 200000,
   "info": "execute success",
   "rateLimit": {
    "id": "e04f201e7b7e4599b42a9b6631a7ba08",
    "disable": false
   }
  }
 ]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEnableRateLimitsRequest
*/
func (a *RateLimitsApiService) EnableRateLimits(ctx _context.Context) ApiEnableRateLimitsRequest {
	return ApiEnableRateLimitsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RateLimitsApiService) EnableRateLimitsExecute(r ApiEnableRateLimitsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitsApiService.EnableRateLimits")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/naming/v1/ratelimits/enable"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Polaris-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetRateLimitsRequest struct {
	ctx _context.Context
	ApiService *RateLimitsApiService
	id string
	name string
	service string
	namespace string
	method string
	disable bool
	brief bool
	offset int32
	limit int32
}


func (r ApiGetRateLimitsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetRateLimitsExecute(r)
}

/*
GetRateLimits 查询限流规则


请求示例：

~~~
GET /naming/v1/ratelimits?参数名=参数值

# 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header
Header X-Polaris-Token: {访问凭据}
~~~

| 参数名    | 类型   | 描述                                                         | 是否必填 |
| --------- | ------ | ------------------------------------------------------------ | -------- |
| id        | string | 规则ID                                                       | 否       |
| name      | string | 规则名                                                       | 否       |
| service   | string | 服务名                                                       | 否       |
| namespace | string | 命名空间                                                     | 否       |
| method    | string | 限流接口名，默认为模糊匹配                                   | 否       |
| disable   | bool   | 规则是否启用，true为不启用，false为启用                      | 否       |
| brief     | bool   | 是否只显示概要信息，brief=true时，则不返回规则详情，只返回规则列表概要信息，默认为false | 否       |
| offset    | int    | 分页的起始位置，默认为0                                      | 否       |
| limit     | int    | 每页行数，默认100                                            | 否       |

应答示例：

~~~
{
 "code": 200000,
 "info": "execute success",
 "amount": 2,
 "size": 2,
 "rateLimits": [
  {
   "id": "e04f201e7b7e4599b42a9b6631a7ba08",
   "service": "testsvc1",
   "namespace": "default",
   "priority": 0,
   "disable": false,
   "ctime": "2022-07-26 21:03:50",
   "mtime": "2022-07-26 21:03:50",
   "revision": "",
   "method": {
    "value": "/getsomething2"
   },
   "name": "rule2",
   "etime": "2022-07-26 21:03:50"
  },
  {
   "id": "6942526fbac545848cd8fb32a3a55bb6",
   "service": "testsvc1",
   "namespace": "default",
   "priority": 0,
   "disable": false,
   "ctime": "2022-07-26 10:09:49",
   "mtime": "2022-07-26 11:46:07",
   "revision": "",
   "method": {
    "value": "/getsomething"
   },
   "name": "rule1",
   "etime": "2022-07-26 11:46:07"
  }
 ]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id 规则ID
 @param name 规则名称
 @param service 服务名称
 @param namespace 命名空间
 @param method 限流接口名，默认为模糊匹配 
 @param disable 规则是否启用，true为不启用，false为启用
 @param brief 是否只显示概要信息，brief=true时，则不返回规则详情，只返回规则列表概要信息，默认为false
 @param offset 分页的起始位置，默认为0
 @param limit 每页行数，默认100  
 @return ApiGetRateLimitsRequest
*/
func (a *RateLimitsApiService) GetRateLimits(ctx _context.Context, id string, name string, service string, namespace string, method string, disable bool, brief bool, offset int32, limit int32) ApiGetRateLimitsRequest {
	return ApiGetRateLimitsRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
		name: name,
		service: service,
		namespace: namespace,
		method: method,
		disable: disable,
		brief: brief,
		offset: offset,
		limit: limit,
	}
}

// Execute executes the request
func (a *RateLimitsApiService) GetRateLimitsExecute(r ApiGetRateLimitsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitsApiService.GetRateLimits")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/naming/v1/ratelimits"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.PathEscape(parameterToString(r.name, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"service"+"}", _neturl.PathEscape(parameterToString(r.service, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", _neturl.PathEscape(parameterToString(r.namespace, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"method"+"}", _neturl.PathEscape(parameterToString(r.method, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"disable"+"}", _neturl.PathEscape(parameterToString(r.disable, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"brief"+"}", _neturl.PathEscape(parameterToString(r.brief, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"offset"+"}", _neturl.PathEscape(parameterToString(r.offset, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"limit"+"}", _neturl.PathEscape(parameterToString(r.limit, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Polaris-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateRateLimitsRequest struct {
	ctx _context.Context
	ApiService *RateLimitsApiService
	body *[]V1RateLimit
}

// update rate limits
func (r ApiUpdateRateLimitsRequest) Body(body []V1RateLimit) ApiUpdateRateLimitsRequest {
	r.body = &body
	return r
}

func (r ApiUpdateRateLimitsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateRateLimitsExecute(r)
}

/*
UpdateRateLimits 更新限流规则


更新服务下的限流规则的相关信息

请求示例：

~~~
PUT /naming/v1/ratelimits

# 开启北极星服务端针对控制台接口鉴权开关后，需要添加下面的 header
Header X-Polaris-Token: {访问凭据}

[
	{
	    "id":   "e04f201e7b7e4599b42a9b6631a7ba08",
		"name": "rule1",
		"service": "testsvc1",
		"namespace": "default",
		"method": {
			"type": "EXACT",
			"value": "/getsomething"
		},
		"arguments": [
			{
				"type": "HEADER",
				"key": "host",
				"value": {
					"type": "EXACT",
					"value": "www.baidu.com"
				}
			},
			{
				"type": "CALLER_SERVICE",
				"key": "default",
				"value": {
					"type": "IN",
					"value": "testsvc1,testsvc2"
				}
			}
		],
		"resource": "QPS",
		"type": "LOCAL",
		"amounts": [
			{
				"maxAmount": 1000,
				"validDuration": "1s"
			},
			{
				"maxAmount": 2000,
				"validDuration": "1m"
			}	
		],
		"regex_combine": false,
		"disable": true,
		"failover": "FAILOVER_LOCAL"
	}
]
~~~

应答示例：

~~~
{
 "code": 200000,
 "info": "execute success",
 "size": 1,
 "responses": [
  {
   "code": 200000,
   "info": "execute success",
   "rateLimit": {
    "id": "e04f201e7b7e4599b42a9b6631a7ba08", //规则ID
    "service": "testsvc1",
    "namespace": "default",
    "name": "rule2"
   }
  }
 ]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateRateLimitsRequest
*/
func (a *RateLimitsApiService) UpdateRateLimits(ctx _context.Context) ApiUpdateRateLimitsRequest {
	return ApiUpdateRateLimitsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RateLimitsApiService) UpdateRateLimitsExecute(r ApiUpdateRateLimitsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RateLimitsApiService.UpdateRateLimits")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/naming/v1/ratelimits"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Polaris-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
