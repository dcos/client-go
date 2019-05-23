# \MetronomeApi

All URIs are relative to *http://your-dcos-cluster.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GetJobs**](MetronomeApi.md#V1GetJobs) | **Get** /service/metronome/v1/jobs | 


# **V1GetJobs**
> []MetronomeV1Job V1GetJobs(ctx, optional)


Get the list of all jobs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1GetJobsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1GetJobsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **embeded** | [**optional.Interface of []MetronomeEmbeded**](MetronomeEmbeded.md)|  | 

### Return type

[**[]MetronomeV1Job**](MetronomeV1Job.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

