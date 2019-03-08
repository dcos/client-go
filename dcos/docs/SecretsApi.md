# \SecretsApi

All URIs are relative to *http://your-dcos-cluster.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecret**](SecretsApi.md#CreateSecret) | **Put** /secrets/v1/secret/{store}/{path-to-secret} | Create a secret in the store at the path.
[**DeleteSecret**](SecretsApi.md#DeleteSecret) | **Delete** /secrets/v1/secret/{store}/{path-to-secret} | Delete a secret.
[**GetSecret**](SecretsApi.md#GetSecret) | **Get** /secrets/v1/secret/{store}/{path-to-secret} | Read or list a secret from the store by its path.
[**UpdateSecret**](SecretsApi.md#UpdateSecret) | **Patch** /secrets/v1/secret/{store}/{path-to-secret} | Update secret.


# **CreateSecret**
> CreateSecret(ctx, store, pathToSecret, secret)
Create a secret in the store at the path.

Create a secret in the store at the path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **store** | **string**| The backend to store the secret in. | 
  **pathToSecret** | **string**| The path to store the secret in. | 
  **secret** | [**Secret**](Secret.md)| Secret value. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSecret**
> DeleteSecret(ctx, store, pathToSecret)
Delete a secret.

Delete a secret.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **store** | **string**| The backend to delete the secret from. | 
  **pathToSecret** | **string**| The path to the secret to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecret**
> Secret GetSecret(ctx, store, pathToSecret, optional)
Read or list a secret from the store by its path.

Read or list a secret from the store by its path.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **store** | **string**| The backend store from which to get the secret. | 
  **pathToSecret** | **string**| The path to store the secret in. | 
 **optional** | ***GetSecretOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSecretOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **list** | **optional.String**| If set to true, list only secret keys.  | 

### Return type

[**Secret**](Secret.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSecret**
> UpdateSecret(ctx, store, pathToSecret, secret)
Update secret.

Update existing secret in the specified store.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **store** | **string**| The backend to store the secret in. | 
  **pathToSecret** | **string**| The path for the secret to update. | 
  **secret** | [**Secret**](Secret.md)| Secret value. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

