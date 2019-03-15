# \EdgelbApi

All URIs are relative to *http://your-dcos-cluster.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigContainer**](EdgelbApi.md#GetConfigContainer) | **Get** /service/edgelb/config | 
[**GetPoolContainer**](EdgelbApi.md#GetPoolContainer) | **Get** /service/edgelb/pools/{name} | 
[**Ping**](EdgelbApi.md#Ping) | **Get** /service/edgelb/ping | 
[**V2CreatePool**](EdgelbApi.md#V2CreatePool) | **Post** /service/edgelb/v2/pools | 
[**V2DeletePool**](EdgelbApi.md#V2DeletePool) | **Delete** /service/edgelb/v2/pools/{name} | 
[**V2GetPool**](EdgelbApi.md#V2GetPool) | **Get** /service/edgelb/v2/pools/{name} | 
[**V2GetPools**](EdgelbApi.md#V2GetPools) | **Get** /service/edgelb/v2/pools | 
[**V2UpdatePool**](EdgelbApi.md#V2UpdatePool) | **Put** /service/edgelb/v2/pools/{name} | 
[**Version**](EdgelbApi.md#Version) | **Get** /service/edgelb/version | 


# **GetConfigContainer**
> ConfigContainer GetConfigContainer(ctx, )


Get the entire configuration object including v1 and v2 pools.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigContainer**](ConfigContainer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPoolContainer**
> PoolContainer GetPoolContainer(ctx, name)


Returns a v1 or v2 load balancer pool based on a single name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**PoolContainer**](PoolContainer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> string Ping(ctx, )


Healthcheck endpoint.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2CreatePool**
> V2Pool V2CreatePool(ctx, v2Pool, optional)


Creates a new load balancer pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **v2Pool** | [**V2Pool**](V2Pool.md)|  | 
 **optional** | ***V2CreatePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V2CreatePoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **optional.String**| DCOS Auth Token | 

### Return type

[**V2Pool**](V2Pool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2DeletePool**
> V2DeletePool(ctx, name, optional)


Deletes a single load balancer pool based on the name supplied.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***V2DeletePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V2DeletePoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **token** | **optional.String**| DCOS Auth Token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2GetPool**
> V2Pool V2GetPool(ctx, name)


Returns a v2 load balancer pool based on a single name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**V2Pool**](V2Pool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2GetPools**
> []V2Pool V2GetPools(ctx, )


Get all load balancer pools.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]V2Pool**](V2Pool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2UpdatePool**
> V2Pool V2UpdatePool(ctx, name, v2Pool, optional)


Updates a new load balancer pool.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **v2Pool** | [**V2Pool**](V2Pool.md)|  | 
 **optional** | ***V2UpdatePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V2UpdatePoolOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **token** | **optional.String**| DCOS Auth Token | 

### Return type

[**V2Pool**](V2Pool.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Version**
> string Version(ctx, )


Returns the installed Edge-LB package version.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

