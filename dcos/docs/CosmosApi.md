# \CosmosApi

All URIs are relative to *http://your-dcos-cluster.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PackageDescribe**](CosmosApi.md#PackageDescribe) | **Post** /package/describe | 
[**PackageInstall**](CosmosApi.md#PackageInstall) | **Post** /package/install | 
[**PackageList**](CosmosApi.md#PackageList) | **Post** /package/list | 
[**PackageRepositoryAdd**](CosmosApi.md#PackageRepositoryAdd) | **Post** /package/repository/add | 
[**PackageRepositoryDelete**](CosmosApi.md#PackageRepositoryDelete) | **Post** /package/repository/delete | 
[**PackageRepositoryList**](CosmosApi.md#PackageRepositoryList) | **Post** /package/repository/list | 
[**PackageSearch**](CosmosApi.md#PackageSearch) | **Post** /package/search | 
[**PackageUninstall**](CosmosApi.md#PackageUninstall) | **Post** /package/uninstall | 
[**ServiceDescribe**](CosmosApi.md#ServiceDescribe) | **Post** /cosmos/service/describe | 
[**ServiceUpdate**](CosmosApi.md#ServiceUpdate) | **Post** /cosmos/service/update | 



## PackageDescribe

> CosmosPackageDescribeV3Response PackageDescribe(ctx, optional)


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
 **cosmosPackageDescribeV1Request** | [**optional.Interface of CosmosPackageDescribeV1Request**](CosmosPackageDescribeV1Request.md)|  | 

### Return type

[**CosmosPackageDescribeV3Response**](CosmosPackageDescribeV3Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.package.describe-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.package.describe-response+json;charset=utf-8;version=v3, application/vnd.dcos.package.error+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageInstall

> CosmosPackageInstallV1Response PackageInstall(ctx, cosmosPackageInstallV1Request)


Runs a service from a Universe package.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cosmosPackageInstallV1Request** | [**CosmosPackageInstallV1Request**](CosmosPackageInstallV1Request.md)|  | 

### Return type

[**CosmosPackageInstallV1Response**](CosmosPackageInstallV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.package.install-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.package.install-response+json;charset=utf-8;version=v1, application/vnd.dcos.package.error+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageList

> CosmosPackageListV1Response PackageList(ctx, optional)


Lists all of the running DC/OS services started from a DC/OS package..

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PackageListOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackageListOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cosmosPackageListV1Request** | [**optional.Interface of CosmosPackageListV1Request**](CosmosPackageListV1Request.md)|  | 

### Return type

[**CosmosPackageListV1Response**](CosmosPackageListV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.package.list-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.package.list-response+json;charset=utf-8;version=v1, application/vnd.dcos.package.error+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageRepositoryAdd

> CosmosPackageAddRepoV1Response PackageRepositoryAdd(ctx, optional)


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
 **cosmosPackageAddRepoV1Request** | [**optional.Interface of CosmosPackageAddRepoV1Request**](CosmosPackageAddRepoV1Request.md)|  | 

### Return type

[**CosmosPackageAddRepoV1Response**](CosmosPackageAddRepoV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.package.repository.add-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.package.repository.add-response+json;charset=utf-8;version=v1, application/vnd.dcos.package.error+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageRepositoryDelete

> CosmosPackageDeleteRepoV1Response PackageRepositoryDelete(ctx, optional)


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
 **cosmosPackageDeleteRepoV1Request** | [**optional.Interface of CosmosPackageDeleteRepoV1Request**](CosmosPackageDeleteRepoV1Request.md)|  | 

### Return type

[**CosmosPackageDeleteRepoV1Response**](CosmosPackageDeleteRepoV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.package.repository.delete-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.package.repository.delete-response+json;charset=utf-8;version=v1, application/vnd.dcos.package.error+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageRepositoryList

> []CosmosPackageListRepoV1Response PackageRepositoryList(ctx, body)


Enumerates the package repositories (for example Universe) that are already installed and in-use by DC/OS. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | **map[string]interface{}**|  | 

### Return type

[**[]CosmosPackageListRepoV1Response**](CosmosPackageListRepoV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.package.repository.list-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.package.repository.list-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageSearch

> CosmosPackageSearchV1Response PackageSearch(ctx, cosmosPackageSearchV1Request)


Lists all matching packages in the repository given a partial pattern. The character \\'\\*\\' can be used to match any number of characters. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cosmosPackageSearchV1Request** | [**CosmosPackageSearchV1Request**](CosmosPackageSearchV1Request.md)|  | 

### Return type

[**CosmosPackageSearchV1Response**](CosmosPackageSearchV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.package.search-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.package.search-response+json;charset=utf-8;version=v1, application/vnd.dcos.package.error+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PackageUninstall

> CosmosPackageUninstallV1Response PackageUninstall(ctx, cosmosPackageUninstallV1Request, optional)


### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cosmosPackageUninstallV1Request** | [**CosmosPackageUninstallV1Request**](CosmosPackageUninstallV1Request.md)|  | 
 **optional** | ***PackageUninstallOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PackageUninstallOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **optional.String**|  | 

### Return type

[**CosmosPackageUninstallV1Response**](CosmosPackageUninstallV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.package.uninstall-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.package.uninstall-response+json;charset=utf-8;version=v1, application/vnd.dcos.package.error+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceDescribe

> CosmosServiceDescribeV1Response ServiceDescribe(ctx, optional)


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
 **cosmosServiceDescribeV1Request** | [**optional.Interface of CosmosServiceDescribeV1Request**](CosmosServiceDescribeV1Request.md)|  | 

### Return type

[**CosmosServiceDescribeV1Response**](CosmosServiceDescribeV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.service.describe-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.service.describe-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceUpdate

> CosmosServiceUpdateV1Response ServiceUpdate(ctx, cosmosServiceUpdateV1Request)


Runs a service update.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cosmosServiceUpdateV1Request** | [**CosmosServiceUpdateV1Request**](CosmosServiceUpdateV1Request.md)|  | 

### Return type

[**CosmosServiceUpdateV1Response**](CosmosServiceUpdateV1Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/vnd.dcos.service.update-request+json;charset=utf-8;version=v1
- **Accept**: application/vnd.dcos.service.update-response+json;charset=utf-8;version=v1

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

