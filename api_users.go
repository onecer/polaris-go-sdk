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

// UsersApiService UsersApi service
type UsersApiService service

type ApiCreateUsersRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	body *V1User
}

// create user
func (r ApiCreateUsersRequest) Body(body V1User) ApiCreateUsersRequest {
	r.body = &body
	return r
}

func (r ApiCreateUsersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CreateUsersExecute(r)
}

/*
CreateUsers 创建用户


批量创建用户至北极星

请求示例：

~~~
POST /core/v1/users
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
[
	{
	"name": "polarismesh",
	"password": "polarismesh",
	"comment": "polarismesh",
	"source": "Polaris"
	}
]
~~~

| 参数名   | 类型   | 描述     | 是否必填 |
|----------|--------|--------|---------|
| name     | string | 用户名   | 是       |
| password | string | 用户密码 | 是       |
| comment  | string | 用户备注 | 否       |
| source   | string | 用户来源 | 否       |


应答示例：

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

数据结构：

> user

| 参数名       | 类型   | 描述                 |
|--------------|--------|--------------------|
| id           | string | 用户ID               |
| name         | string | 用户名称             |
| password     | string | 用户密码             |
| source       | string | 用户来源             |
| auth_token   | string | 用户访问凭据         |
| token_enable | bool   | 用户访问凭据是否可用 |
| comment      | string | 用户备注信息         |
| ctime        | string | 用户创建时间         |
| mtime        | string | 用户修改时间         |


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateUsersRequest
*/
func (a *UsersApiService) CreateUsers(ctx _context.Context) ApiCreateUsersRequest {
	return ApiCreateUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) CreateUsersExecute(r ApiCreateUsersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.CreateUsers")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/users"

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

type ApiDeleteUsersRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	body *V1User
}

// delete user
func (r ApiDeleteUsersRequest) Body(body V1User) ApiDeleteUsersRequest {
	r.body = &body
	return r
}

func (r ApiDeleteUsersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteUsersExecute(r)
}

/*
DeleteUsers 删除用户


批量删除北极星中的用户

请求示例

~~~
POST /core/v1/users/delete
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
[
	{
		"id": "xxx"
	}
]
~~~

| 参数名 | 类型   | 描述   | 是否必填 |
|--------|--------|------|---------|
| id     | string | 用户ID | 是       |


响应示例：

~~~json
{
	"code": 200000,
	"info": "execute success",
	"responses": [
		{
			"code": 200000,
			"info": "execute success"
		}
	]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteUsersRequest
*/
func (a *UsersApiService) DeleteUsers(ctx _context.Context) ApiDeleteUsersRequest {
	return ApiDeleteUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) DeleteUsersExecute(r ApiDeleteUsersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.DeleteUsers")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/users/delete"

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

type ApiGetUserTokenRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	id *string
}

// 用户ID
func (r ApiGetUserTokenRequest) Id(id string) ApiGetUserTokenRequest {
	r.id = &id
	return r
}

func (r ApiGetUserTokenRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetUserTokenExecute(r)
}

/*
GetUserToken 获取用户Token


查询用户的Token凭据信息，通过用户ID或者用户名进行查询

请求示例：

~~~
GET /core/v1/user/token?id=xxx
Header X-Polaris-Token: {访问凭据}
~~~

| 参数名 | 类型   | 描述   | 是否必填 |
|--------|--------|------|---------|
| id     | string | 用户ID | 是       |

响应示例：

~~~json
{
	"code": 200000,
	"info": "execute success",
	"user": {
		"id": "xxx",
		"auth_token": "xxx",
		"token_enable": true
	}
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserTokenRequest
*/
func (a *UsersApiService) GetUserToken(ctx _context.Context) ApiGetUserTokenRequest {
	return ApiGetUserTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) GetUserTokenExecute(r ApiGetUserTokenRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.GetUserToken")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/user/token"

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

type ApiGetUsersRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	id *string
	name *string
	source *string
	groupId *string
	offset *int32
	limit *int32
}

// 用户ID
func (r ApiGetUsersRequest) Id(id string) ApiGetUsersRequest {
	r.id = &id
	return r
}
// 用户名称, 当前仅提供全模糊搜索
func (r ApiGetUsersRequest) Name(name string) ApiGetUsersRequest {
	r.name = &name
	return r
}
// 用户来源
func (r ApiGetUsersRequest) Source(source string) ApiGetUsersRequest {
	r.source = &source
	return r
}
// 用户组ID, 用于查询某个用户组下用户列表
func (r ApiGetUsersRequest) GroupId(groupId string) ApiGetUsersRequest {
	r.groupId = &groupId
	return r
}
// 查询偏移量, 默认为0
func (r ApiGetUsersRequest) Offset(offset int32) ApiGetUsersRequest {
	r.offset = &offset
	return r
}
// 本次查询条数, 最大为100
func (r ApiGetUsersRequest) Limit(limit int32) ApiGetUsersRequest {
	r.limit = &limit
	return r
}

func (r ApiGetUsersRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetUsersExecute(r)
}

/*
GetUsers 获取用户


根据相关条件对用户列表进行查询

请求示例：

~~~
GET /core/v1/users?id=xxx&name=xxx&source=xxx&group_id=xxx&offset=xxx&limit=xxx
Header X-Polaris-Token: {访问凭据}
~~~

| 参数名   | 类型   | 描述                                   | 是否必填 |
|----------|--------|--------------------------------------|---------|
| id       | string | 用户ID                                 | 否       |
| name     | string | 用户名称, 当前仅提供全模糊搜索         | 否       |
| source   | string | 用户来源                               | 否       |
| group_id | string | 用户组ID, 用于查询某个用户组下用户列表 | 否       |
| offset   | int    | 查询偏移量, 默认为0                    | 否       |
| limit    | int    | 本次查询条数, 最大为100                | 否       |


响应示例：

~~~json
{
	"code": 200000,
	"info": "execute success",
	"amount": 1,
	"size": 1,
	"users": [
		{
			"id": "xxx",
			"name": "xxx",
			"source": "",
			"auth_token": "",
			"token_enable": true,
			"comment": "",
			"ctime": "2022-02-09 19:48:53",
			"mtime": "2022-02-09 19:48:53",
		}
	]
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUsersRequest
*/
func (a *UsersApiService) GetUsers(ctx _context.Context) ApiGetUsersRequest {
	return ApiGetUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) GetUsersExecute(r ApiGetUsersRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.GetUsers")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.id != nil {
		localVarQueryParams.Add("id", parameterToString(*r.id, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.source != nil {
		localVarQueryParams.Add("source", parameterToString(*r.source, ""))
	}
	if r.groupId != nil {
		localVarQueryParams.Add("group_id", parameterToString(*r.groupId, ""))
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

type ApiLoginRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	body *V1LoginRequest
}

// 用户登录接口
func (r ApiLoginRequest) Body(body V1LoginRequest) ApiLoginRequest {
	r.body = &body
	return r
}

func (r ApiLoginRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.LoginExecute(r)
}

/*
Login 用户登录


用于控制台进行用户登陆操作

请求示例：

~~~
POST /core/v1/user/login
~~~

| 参数名   | 类型   | 描述     | 是否必填 |
|----------|--------|--------|---------|
| name     | string | 用户名   | 是       |
| password | string | 用户密码 | 是       |


应答示例：

~~~json
{
	"code": 200000,
	"info": "execute success",
	"loginResponse": {
		"token": "xxx",
		"name": "xxx",
		"user_id": "xxx",
		"role": "xxx"
	}
}
~~~

| 参数名        |         | 类型   | 描述                                                    |
|---------------|---------|--------|-------------------------------------------------------|
| code          |         | uint32 | 六位返回码                                              |
| info          |         | string | 返回信息                                                |
| loginResponse |         |        | 命名空间                                                |
|               | token   | string | 用户Token, 用于接口请求访问                             |
|               | name    | string | 用户名                                                  |
|               | role    | string | 当前用户角色, (admin:超级账户, main:主账户, sub:子账户) |
|               | user_id | string | 当前用户ID                                              |



 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLoginRequest
*/
func (a *UsersApiService) Login(ctx _context.Context) ApiLoginRequest {
	return ApiLoginRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) LoginExecute(r ApiLoginRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.Login")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/user/login"

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

type ApiResetUserTokenRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	body *V1User
}

// reset user token
func (r ApiResetUserTokenRequest) Body(body V1User) ApiResetUserTokenRequest {
	r.body = &body
	return r
}

func (r ApiResetUserTokenRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ResetUserTokenExecute(r)
}

/*
ResetUserToken 重置用户Token


重置用户的Token, 当Token重置之后，原先的Token失效并且无法进行访问北极星接口以及资源

请求示例：

~~~
PUT /core/v1/user/token/refresh
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
{
	"id": "xxx"
}
~~~

| 参数名 | 类型   | 描述   | 是否必填 |
|--------|--------|------|---------|
| id     | string | 用户ID | 是       |


响应示例：

~~~json
{
	"code": 200000,
	"info": "execute success"
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResetUserTokenRequest
*/
func (a *UsersApiService) ResetUserToken(ctx _context.Context) ApiResetUserTokenRequest {
	return ApiResetUserTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) ResetUserTokenExecute(r ApiResetUserTokenRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.ResetUserToken")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/user/token/refresh"

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

type ApiUpdateUserRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	body *V1User
}

// update user
func (r ApiUpdateUserRequest) Body(body V1User) ApiUpdateUserRequest {
	r.body = &body
	return r
}

func (r ApiUpdateUserRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateUserExecute(r)
}

/*
UpdateUser 更新用户


更新用户的备注信息数据

请求示例：

~~~
PUT /core/v1/user
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
{
	"id": "xxx",
	"comment": "polarismesh"
}
~~~

| 参数名  | 类型   | 描述     | 是否必填 |
|---------|--------|--------|---------|
| id      | string | 用户ID   | 是       |
| comment | string | 用户备注 | 否       |


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateUserRequest
*/
func (a *UsersApiService) UpdateUser(ctx _context.Context) ApiUpdateUserRequest {
	return ApiUpdateUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) UpdateUserExecute(r ApiUpdateUserRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UpdateUser")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/user"

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

type ApiUpdateUserPasswordRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	body *V1ModifyUserPassword
}

// update user password
func (r ApiUpdateUserPasswordRequest) Body(body V1ModifyUserPassword) ApiUpdateUserPasswordRequest {
	r.body = &body
	return r
}

func (r ApiUpdateUserPasswordRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateUserPasswordExecute(r)
}

/*
UpdateUserPassword 更新用户密码


用户重新更新密码, 如果密码更新成功, 则会一并更新对应用户的访问凭据

请求示例：

~~~
PUT /core/v1/user/password
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
{
	"id": "xxx",
	"old_password": "xxx",
	"new_password": "xxx"
}
~~~

| 参数名       | 类型   | 描述   | 是否必填 |
|--------------|--------|------|---------|
| id           | string | 用户ID | 是       |
| old_password | string | 旧密码 | 否       |
| new_password | string | 新密码 | 否       |


响应示例：


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateUserPasswordRequest
*/
func (a *UsersApiService) UpdateUserPassword(ctx _context.Context) ApiUpdateUserPasswordRequest {
	return ApiUpdateUserPasswordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) UpdateUserPasswordExecute(r ApiUpdateUserPasswordRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UpdateUserPassword")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/user/password"

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

type ApiUpdateUserTokenRequest struct {
	ctx _context.Context
	ApiService *UsersApiService
	body *V1User
}

// update user token
func (r ApiUpdateUserTokenRequest) Body(body V1User) ApiUpdateUserTokenRequest {
	r.body = &body
	return r
}

func (r ApiUpdateUserTokenRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateUserTokenExecute(r)
}

/*
UpdateUserToken 更新用户Token


对用户Token的使用状态进行修改, 如果用户的Token被设置为禁用状态, 则该Token后续无法用于访问北极星接口以及资源, 需使用主账户或者超级账户进行解封

请求示例：

~~~
PUT /core/v1/user/token/status
Header X-Polaris-Token: {访问凭据}
~~~

~~~json
{
	"id": "xxx",
	"token_enable": false
}
~~~

| 参数名       | 类型   | 描述                                       | 是否必填 |
|--------------|--------|------------------------------------------|---------|
| id           | string | 用户ID                                     | 是       |
| token_enable | bool   | 当前Token可用状态, true为启用, false为禁用 | 是       |


响应示例：

~~~json
{
	"code": 200000,
	"info": "execute success"
}
~~~


 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUpdateUserTokenRequest
*/
func (a *UsersApiService) UpdateUserToken(ctx _context.Context) ApiUpdateUserTokenRequest {
	return ApiUpdateUserTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UsersApiService) UpdateUserTokenExecute(r ApiUpdateUserTokenRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersApiService.UpdateUserToken")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/core/v1/user/token/status"

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