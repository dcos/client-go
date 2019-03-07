# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DcosSystemHealth**](DefaultApi.md#DcosSystemHealth) | **Get** /system/health/v1 | 
[**DcosSystemNodes**](DefaultApi.md#DcosSystemNodes) | **Get** /system/health/v1/nodes | 
[**HealthchechService**](DefaultApi.md#HealthchechService) | **Get** /service/{appId}/v1/health | 
[**KubernetesAuthData**](DefaultApi.md#KubernetesAuthData) | **Get** /service/{appId}/v1/auth/data | 
[**PackageDescribe**](DefaultApi.md#PackageDescribe) | **Post** /package/describe | 
[**PackageInstall**](DefaultApi.md#PackageInstall) | **Post** /package/install | 
[**PackageRepositoryAdd**](DefaultApi.md#PackageRepositoryAdd) | **Post** /package/repository/add | 
[**PackageRepositoryDelete**](DefaultApi.md#PackageRepositoryDelete) | **Post** /package/repository/delete | 
[**PackageSearch**](DefaultApi.md#PackageSearch) | **Post** /package/search | 
[**PackageUninstall**](DefaultApi.md#PackageUninstall) | **Post** /package/uninstall | 
[**ServiceDescribe**](DefaultApi.md#ServiceDescribe) | **Post** /cosmos/service/describe | 
[**ServicePlan**](DefaultApi.md#ServicePlan) | **Get** /service/{appId}/v1/plans/{plan} | 
[**ServiceUpdate**](DefaultApi.md#ServiceUpdate) | **Post** /cosmos/service/update | 


# **DcosSystemHealth**
> SystemHealth DcosSystemHealth(ctx, )


DC/OS system health.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemHealth**](systemHealth.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DcosSystemNodes**
> SystemNodes DcosSystemNodes(ctx, )


DC/OS nodes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemNodes**](systemNodes.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HealthchechService**
> HealthchechService(ctx, appId)


Healthcheck service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KubernetesAuthData**
> KubernetesAuthDataResponse KubernetesAuthData(ctx, appId)


Kuberentes Auth Data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 

### Return type

[**KubernetesAuthDataResponse**](kubernetesAuthDataResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PackageDescribe**
> V3PackageDescribeResponse PackageDescribe(ctx, optional)


Show information about the package, including the required resources and configuration to start the service, and command line extensions that are included with the package.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PackageDescribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackageDescribeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageDescribeRequest** | [**optional.Interface of PackageDescribeRequest**](PackageDescribeRequest.md)|  | 

### Return type

[**V3PackageDescribeResponse**](v3PackageDescribeResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/vnd.dcos.package.describe-request+json;charset=utf-8;version=v1
 - **Accept**: application/vnd.dcos.package.describe-response+json;charset=utf-8;version=v3

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PackageInstall**
> InstallResponse PackageInstall(ctx, installRequest)


Runs a service from a Universe package.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **installRequest** | [**InstallRequest**](InstallRequest.md)|  | 

### Return type

[**InstallResponse**](installResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/vnd.dcos.package.install-request+json;charset=utf-8;version=v1
 - **Accept**: application/vnd.dcos.package.install-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PackageRepositoryAdd**
> PackageAddRepoResponse PackageRepositoryAdd(ctx, optional)


Adds a package repository (for example Universe) for use by DC/OS. To add a package repository to the beginning of the list set the index to zero (0). To add a package repository to the end of the list do not specify an index. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PackageRepositoryAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackageRepositoryAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageAddRepoRequest** | [**optional.Interface of PackageAddRepoRequest**](PackageAddRepoRequest.md)|  | 

### Return type

[**PackageAddRepoResponse**](packageAddRepoResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/vnd.dcos.package.repository.add-request+json;charset=utf-8;version=v1
 - **Accept**: application/vnd.dcos.package.repository.add-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PackageRepositoryDelete**
> PackageDeleteRepoResponse PackageRepositoryDelete(ctx, optional)


Deletes a package repository (for example Universe) from DC/OS.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PackageRepositoryDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackageRepositoryDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **packageDeleteRepoRequest** | [**optional.Interface of PackageDeleteRepoRequest**](PackageDeleteRepoRequest.md)|  | 

### Return type

[**PackageDeleteRepoResponse**](packageDeleteRepoResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/vnd.dcos.package.repository.delete-request+json;charset=utf-8;version=v1
 - **Accept**: application/vnd.dcos.package.repository.delete-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PackageSearch**
> SearchResponse PackageSearch(ctx, searchRequest, optional)


Lists all matching packages in the repository given a partial pattern. The character \\'\\*\\' can be used to match any number of characters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **searchRequest** | [**SearchRequest**](SearchRequest.md)|  | 
 **optional** | ***PackageSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackageSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **optional.String**|  | 

### Return type

[**SearchResponse**](searchResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/vnd.dcos.package.search-request+json;charset=utf-8;version=v1
 - **Accept**: application/vnd.dcos.package.search-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PackageUninstall**
> UninstallResponse PackageUninstall(ctx, uninstallRequest, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uninstallRequest** | [**UninstallRequest**](UninstallRequest.md)|  | 
 **optional** | ***PackageUninstallOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PackageUninstallOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **optional.String**|  | 

### Return type

[**UninstallResponse**](uninstallResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/vnd.dcos.package.uninstall-request+json;charset=utf-8;version=v1
 - **Accept**: application/vnd.dcos.package.uninstall-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceDescribe**
> ServiceDescribeResponse ServiceDescribe(ctx, optional)


Describes a DC/OS Service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceDescribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceDescribeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceDescribeRequest** | [**optional.Interface of ServiceDescribeRequest**](ServiceDescribeRequest.md)|  | 

### Return type

[**ServiceDescribeResponse**](serviceDescribeResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/vnd.dcos.service.describe-request+json;charset=utf-8;version=v1
 - **Accept**: application/vnd.dcos.service.describe-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServicePlan**
> PlanDefinition ServicePlan(ctx, appId, plan)


Service plan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **appId** | **string**|  | 
  **plan** | **string**|  | 

### Return type

[**PlanDefinition**](planDefinition.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceUpdate**
> ServiceUpdateResponse ServiceUpdate(ctx, serviceUpdateRequest)


Runs a service update.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **serviceUpdateRequest** | [**ServiceUpdateRequest**](ServiceUpdateRequest.md)|  | 

### Return type

[**ServiceUpdateResponse**](serviceUpdateResponse.md)

### Authorization

[token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/vnd.dcos.service.update-request+json;charset=utf-8;version=v1
 - **Accept**: application/vnd.dcos.service.update-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

