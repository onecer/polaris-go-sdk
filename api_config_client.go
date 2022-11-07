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

// ConfigClientApiService ConfigClientApi service
type ConfigClientApiService service

type ApiGetConfigFileRequest struct {
	ctx _context.Context
	ApiService *ConfigClientApiService
	namespace *string
	group *string
	fileName *string
	version *int32
}

// 命名空间
func (r ApiGetConfigFileRequest) Namespace(namespace string) ApiGetConfigFileRequest {
	r.namespace = &namespace
	return r
}
// 配置文件分组
func (r ApiGetConfigFileRequest) Group(group string) ApiGetConfigFileRequest {
	r.group = &group
	return r
}
// 配置文件名
func (r ApiGetConfigFileRequest) FileName(fileName string) ApiGetConfigFileRequest {
	r.fileName = &fileName
	return r
}
// 配置文件客户端版本号，刚启动时设置为 0
func (r ApiGetConfigFileRequest) Version(version int32) ApiGetConfigFileRequest {
	r.version = &version
	return r
}

func (r ApiGetConfigFileRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.GetConfigFileExecute(r)
}

/*
GetConfigFile 拉取配置

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetConfigFileRequest
*/
func (a *ConfigClientApiService) GetConfigFile(ctx _context.Context) ApiGetConfigFileRequest {
	return ApiGetConfigFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigClientApiService) GetConfigFileExecute(r ApiGetConfigFileRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigClientApiService.GetConfigFile")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/config/v1/GetConfigFile"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.namespace == nil {
		return nil, reportError("namespace is required and must be specified")
	}
	if r.group == nil {
		return nil, reportError("group is required and must be specified")
	}
	if r.fileName == nil {
		return nil, reportError("fileName is required and must be specified")
	}
	if r.version == nil {
		return nil, reportError("version is required and must be specified")
	}

	localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	localVarQueryParams.Add("group", parameterToString(*r.group, ""))
	localVarQueryParams.Add("fileName", parameterToString(*r.fileName, ""))
	localVarQueryParams.Add("version", parameterToString(*r.version, ""))
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

type ApiWatchConfigFileRequest struct {
	ctx _context.Context
	ApiService *ConfigClientApiService
	body *V1ClientWatchConfigFileRequest
}

// 通过 Http LongPolling 机制订阅配置变更。
func (r ApiWatchConfigFileRequest) Body(body V1ClientWatchConfigFileRequest) ApiWatchConfigFileRequest {
	r.body = &body
	return r
}

func (r ApiWatchConfigFileRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.WatchConfigFileExecute(r)
}

/*
WatchConfigFile 监听配置

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWatchConfigFileRequest
*/
func (a *ConfigClientApiService) WatchConfigFile(ctx _context.Context) ApiWatchConfigFileRequest {
	return ApiWatchConfigFileRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigClientApiService) WatchConfigFileExecute(r ApiWatchConfigFileRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigClientApiService.WatchConfigFile")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/config/v1/WatchConfigFile"

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
