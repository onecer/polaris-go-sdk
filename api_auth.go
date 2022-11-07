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
)

// Linger please
var (
	_ _context.Context
)

// AuthApiService AuthApi service
type AuthApiService service

type ApiAuthStatusRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
}


func (r ApiAuthStatusRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AuthStatusExecute(r)
}

/*
AuthStatus 查询鉴权开关信息



请求示例：

~~~
GET /core/v1/auth/status
~~~



响应示例：

~~~
{
	"code": 200000,
	"info": "execute success",
	"optionSwitch": 
	{
		"auth": "true"
	}
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthStatusRequest
*/
func (a *AuthApiService) AuthStatus(ctx _context.Context) ApiAuthStatusRequest {
	return ApiAuthStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) AuthStatusExecute(r ApiAuthStatusRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.AuthStatus")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/auth/status"

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

type ApiCreateStrategyRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	body *V1AuthStrategy
}

// create auth strategy
func (r ApiCreateStrategyRequest) Body(body V1AuthStrategy) ApiCreateStrategyRequest {
	r.body = &body
	return r
}

func (r ApiCreateStrategyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CreateStrategyExecute(r)
}

/*
CreateStrategy 创建鉴权策略


请求示例：

~~~
POST /core/v1/auth/strategy
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
{
  "name": "xxx",
  "comment": "xxx",
  "principals": {
    "users": [
      {
          "id": "xxx"
      }
    ],
    "groups": [
      {
          "id": "xxx"
      }
    ]
  },
  "resources": {
    "namespaces": [
      {
          "id": "Polaris"
      }
    ],
    "services": [
      {
          "id": "Polaris"
      }
    ],
    "config_groups": [
      {
          "id": "Polaris"
      }
    ]
  }
}
~~~

响应示例：

~~~json
{
    "code": 200000,
    "info": "execute success"
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateStrategyRequest
*/
func (a *AuthApiService) CreateStrategy(ctx _context.Context) ApiCreateStrategyRequest {
	return ApiCreateStrategyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) CreateStrategyExecute(r ApiCreateStrategyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.CreateStrategy")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/auth/strategy"

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

type ApiDeleteStrategiesRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	body *V1AuthStrategy
}

// delete auth strategy
func (r ApiDeleteStrategiesRequest) Body(body V1AuthStrategy) ApiDeleteStrategiesRequest {
	r.body = &body
	return r
}

func (r ApiDeleteStrategiesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteStrategiesExecute(r)
}

/*
DeleteStrategies 删除鉴权策略


请求示例：

~~~
POST /core/v1/auth/strategies/delete
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
[
    {
        "id": "xxx"
    }
]
~~~

响应示例：

~~~json
{
    "code": 200000,
    "info": "execute success",
    "size": 1,
    "responses": [
        {
            "code": 200000,
            "info": "execute success"
        }
    ]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteStrategiesRequest
*/
func (a *AuthApiService) DeleteStrategies(ctx _context.Context) ApiDeleteStrategiesRequest {
	return ApiDeleteStrategiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) DeleteStrategiesExecute(r ApiDeleteStrategiesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.DeleteStrategies")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/auth/strategies/delete"

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

type ApiGetPrincipalResourcesRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	principalId *string
	principalType *string
}

// 策略ID
func (r ApiGetPrincipalResourcesRequest) PrincipalId(principalId string) ApiGetPrincipalResourcesRequest {
	r.principalId = &principalId
	return r
}
// Principal类别，user/group
func (r ApiGetPrincipalResourcesRequest) PrincipalType(principalType string) ApiGetPrincipalResourcesRequest {
	r.principalType = &principalType
	return r
}

func (r ApiGetPrincipalResourcesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetPrincipalResourcesExecute(r)
}

/*
GetPrincipalResources 获取鉴权策略详细


请求示例：

~~~
GET /core/v1/auth/principal/resources?principal_id=xxx&principal_type=user
Header X-Polaris-Token: {访问凭据}
~~~

| 参数名         | 类型   | 描述                     | 是否必填 |
|----------------|--------|------------------------|---------|
| principal_id   | string | 策略ID                   | 是       |
| principal_type | string | Principal类别，user/group | 是       |


响应示例：

~~~json
{
    "code": 200000,
    "info": "execute success",
    "resources": {
        "namespaces": [
            {
                "id": "xxx",
                "namespace": "xxx",
                "name": "xxx"
            }
        ],
        "services": [
            {
                "id": "xxx",
                "namespace": "Polaris",
                "name": "xxx"
            }
        ],
        "config_groups": [{
                "id": "xxx",
                "namespace": "xxx",
                "name": "xxx"
            }
        ]
    }
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPrincipalResourcesRequest
*/
func (a *AuthApiService) GetPrincipalResources(ctx _context.Context) ApiGetPrincipalResourcesRequest {
	return ApiGetPrincipalResourcesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) GetPrincipalResourcesExecute(r ApiGetPrincipalResourcesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.GetPrincipalResources")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/auth/principal/resources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.principalId == nil {
		return nil, reportError("principalId is required and must be specified")
	}
	if r.principalType == nil {
		return nil, reportError("principalType is required and must be specified")
	}

	localVarQueryParams.Add("principal_id", parameterToString(*r.principalId, ""))
	localVarQueryParams.Add("principal_type", parameterToString(*r.principalType, ""))
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

type ApiGetStrategiesRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	id *string
	name *string
	default_ *string
	resId *string
	resType *string
	principalId *string
	principalType *string
	showDetail *bool
	offset *int32
	limit *int32
}

// 策略ID
func (r ApiGetStrategiesRequest) Id(id string) ApiGetStrategiesRequest {
	r.id = &id
	return r
}
// 策略名称, 当前仅提供全模糊搜索
func (r ApiGetStrategiesRequest) Name(name string) ApiGetStrategiesRequest {
	r.name = &name
	return r
}
// “0” 查询自定义策略；“1” 查询默认策略；不填则为查询（默认+自定义）鉴权策略
func (r ApiGetStrategiesRequest) Default_(default_ string) ApiGetStrategiesRequest {
	r.default_ = &default_
	return r
}
// 资源ID
func (r ApiGetStrategiesRequest) ResId(resId string) ApiGetStrategiesRequest {
	r.resId = &resId
	return r
}
// 资源类型, namespace、service、config_group
func (r ApiGetStrategiesRequest) ResType(resType string) ApiGetStrategiesRequest {
	r.resType = &resType
	return r
}
// 成员ID
func (r ApiGetStrategiesRequest) PrincipalId(principalId string) ApiGetStrategiesRequest {
	r.principalId = &principalId
	return r
}
// 成员类型, user、group
func (r ApiGetStrategiesRequest) PrincipalType(principalType string) ApiGetStrategiesRequest {
	r.principalType = &principalType
	return r
}
// 是否显示策略详细
func (r ApiGetStrategiesRequest) ShowDetail(showDetail bool) ApiGetStrategiesRequest {
	r.showDetail = &showDetail
	return r
}
// 查询偏移量, 默认为0
func (r ApiGetStrategiesRequest) Offset(offset int32) ApiGetStrategiesRequest {
	r.offset = &offset
	return r
}
// 本次查询条数, 最大为100
func (r ApiGetStrategiesRequest) Limit(limit int32) ApiGetStrategiesRequest {
	r.limit = &limit
	return r
}

func (r ApiGetStrategiesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetStrategiesExecute(r)
}

/*
GetStrategies 查询鉴权策略列表


请求示例：

~~~
GET /core/v1/auth/strategies?{key}={value}
Header X-Polaris-Token: {访问凭据}
~~~

支持的URL参数

| 参数名         | 类型   | 描述                                                                  | 是否必填 |
|----------------|--------|---------------------------------------------------------------------|---------|
| id             | string | 策略ID                                                                | 否       |
| name           | string | 策略名称, 当前仅提供全模糊搜索                                        | 否       |
| default        | string | “0” 查询自定义策略；“1” 查询默认策略；不填则为查询（默认+自定义）鉴权策略 | 否       |
| res_id         | string | 资源ID                                                                | 否       |
| res_type       | string | 资源类型, namespace、service、config_group                              | 否       |
| principal_id   | string | 成员ID                                                                | 否       |
| principal_type | string | 成员类型, user、group                                                  | 否       |
| show_detail    | bool   | 是否显示策略详细                                                      | 否       |
| offset         | int    | 查询偏移量, 默认为0                                                   | 否       |
| limit          | int    | 本次查询条数, 最大为100                                               | 否       |


响应示例：

~~~json
{
    "code": 200000,
    "info": "execute success",
    "amount": 1,
    "size": 1,
    "authStrategies": [
        {
            "id": "xxx",
            "name": "xxx",
            "principals": {
                "users": [
                    {
                        "id": "xxx",
                        "name": "xxx"
                    }
                ],
                "groups": [
                    {
                        "id": "xxx",
                        "name": "xxx"
                    }
                ]
            },
            "resources": {
                "strategy_id": null,
                "namespaces": [],
                "services": [
                    {
                        "id": "xxx",
                        "namespace": "default",
                        "name": "Demo-1"
                    },
                    {
                        "id": "xxx",
                        "namespace": "default",
                        "name": "Demo-2"
                    }
                ],
                "config_groups": []
            },
            "action": "READ_WRITE",
            "comment": "Default Strategy",
            "ctime": "2022-02-09 19:48:53",
            "mtime": "2022-02-09 19:48:53",
            "default_strategy": true
        },
    ]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStrategiesRequest
*/
func (a *AuthApiService) GetStrategies(ctx _context.Context) ApiGetStrategiesRequest {
	return ApiGetStrategiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) GetStrategiesExecute(r ApiGetStrategiesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.GetStrategies")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/auth/strategies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.default_ != nil {
		localVarQueryParams.Add("default", parameterToString(*r.default_, ""))
	}
	if r.resId != nil {
		localVarQueryParams.Add("res_id", parameterToString(*r.resId, ""))
	}
	if r.resType != nil {
		localVarQueryParams.Add("res_type", parameterToString(*r.resType, ""))
	}
	if r.principalId != nil {
		localVarQueryParams.Add("principal_id", parameterToString(*r.principalId, ""))
	}
	if r.principalType != nil {
		localVarQueryParams.Add("principal_type", parameterToString(*r.principalType, ""))
	}
	if r.showDetail != nil {
		localVarQueryParams.Add("show_detail", parameterToString(*r.showDetail, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
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

type ApiGetStrategyRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	id *string
}

// 策略ID
func (r ApiGetStrategyRequest) Id(id string) ApiGetStrategyRequest {
	r.id = &id
	return r
}

func (r ApiGetStrategyRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetStrategyExecute(r)
}

/*
GetStrategy 获取鉴权策略详细


根据策略ID查询该策略的具体详细信息

请求示例：

~~~
GET /core/v1/auth/strategy/detail?id=xxx
Header X-Polaris-Token: {访问凭据}
~~~

| 参数名 | 类型   | 描述   | 是否必填 |
|--------|--------|------|---------|
| id     | string | 策略ID | 是       |



响应示例：

~~~json
{
    "code": 200000,
    "info": "execute success",
    "authStrategy": {
        "id": "xxx",
        "name": "xxx",
        "principals": {
            "users": [
                {
                    "id": "xxx",
                    "name": "xxx"
                }
            ],
            "groups": []
        },
        "resources": {
            "namespaces": [],
            "services": [
                {
                    "id": "xxx",
                    "namespace": "default",
                    "name": "Demo-1"
                },
                {
                    "id": "xxx",
                    "namespace": "default",
                    "name": "Demo-2"
                }
            ],
            "config_groups": []
        },
        "action": "READ_WRITE",
        "comment": "Default Strategy",
        "ctime": "2022-02-09 19:43:26",
        "mtime": "2022-02-15 23:20:48",
        "default_strategy": true
    }
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStrategyRequest
*/
func (a *AuthApiService) GetStrategy(ctx _context.Context) ApiGetStrategyRequest {
	return ApiGetStrategyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) GetStrategyExecute(r ApiGetStrategyRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.GetStrategy")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/auth/strategy/detail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.id == nil {
		return nil, reportError("id is required and must be specified")
	}

	localVarQueryParams.Add("id", parameterToString(*r.id, ""))
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

type ApiUpdateStrategiesRequest struct {
	ctx _context.Context
	ApiService *AuthApiService
	body *V1AuthStrategy
}

// update auth strategy
func (r ApiUpdateStrategiesRequest) Body(body V1AuthStrategy) ApiUpdateStrategiesRequest {
	r.body = &body
	return r
}

func (r ApiUpdateStrategiesRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateStrategiesExecute(r)
}

/*
UpdateStrategies 更新鉴权策略


请求示例：

~~~
PUT /core/v1/auth/strategies
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
[
    {
        "id": "xxx",
        "comment": "xxx",
        "add_principals": {
            "users": [{
                "id": "xxx"
            },{
                "id": "xxx"
            }],
            "groups": [{
                "id": "xxx"
            }]
        },
        "remove_principals": {
            "users": [{"id": "xxx"}],
            "groups": [{
                "id": "xxx"
            }]
        },
        "add_resources": {
          "namespaces": [
            {
                "id":"xxx"
            }],
          "services": [{
                "id":"xxx"
           }],
          "config_groups": [
              {
                "id":"xxx"
           }]
        },
        "remove_resources": {
          "namespaces": [
            {
                "id": "xxx"
            }
          ],
          "services": [],
          "config_groups": []
        }
    }
]
~~~

响应示例：

~~~json
{
    "code": 200000,
    "info": "execute success",
    "size": 1,
    "responses": [
        {
            "code": 200000,
            "info": "execute success"
        }
    ]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateStrategiesRequest
*/
func (a *AuthApiService) UpdateStrategies(ctx _context.Context) ApiUpdateStrategiesRequest {
	return ApiUpdateStrategiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) UpdateStrategiesExecute(r ApiUpdateStrategiesRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.UpdateStrategies")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/auth/strategies"

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
