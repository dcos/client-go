# \MetronomeApi

All URIs are relative to *http://your-dcos-cluster.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateJob**](MetronomeApi.md#V1CreateJob) | **Post** /service/metronome/v1/jobs | 
[**V1CreateJobSchedules**](MetronomeApi.md#V1CreateJobSchedules) | **Post** /service/metronome/v1/jobs/{jobId}/schedules | 
[**V1DeleteJob**](MetronomeApi.md#V1DeleteJob) | **Delete** /service/metronome/v1/jobs/{jobId} | 
[**V1DeleteJobSchedulesByScheduleId**](MetronomeApi.md#V1DeleteJobSchedulesByScheduleId) | **Delete** /service/metronome/v1/jobs/{jobId}/schedules/{scheduleId} | 
[**V1GetJob**](MetronomeApi.md#V1GetJob) | **Get** /service/metronome/v1/jobs/{jobId} | 
[**V1GetJobSchedules**](MetronomeApi.md#V1GetJobSchedules) | **Get** /service/metronome/v1/jobs/{jobId}/schedules | 
[**V1GetJobSchedulesByScheduleId**](MetronomeApi.md#V1GetJobSchedulesByScheduleId) | **Get** /service/metronome/v1/jobs/{jobId}/schedules/{scheduleId} | 
[**V1GetJobs**](MetronomeApi.md#V1GetJobs) | **Get** /service/metronome/v1/jobs | 
[**V1PutJobSchedulesByScheduleId**](MetronomeApi.md#V1PutJobSchedulesByScheduleId) | **Put** /service/metronome/v1/jobs/{jobId}/schedules/{scheduleId} | 
[**V1UpdateJob**](MetronomeApi.md#V1UpdateJob) | **Put** /service/metronome/v1/jobs/{jobId} | 



## V1CreateJob

> MetronomeV1Job V1CreateJob(ctx, metronomeV1Job)


Create a new job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metronomeV1Job** | [**MetronomeV1Job**](MetronomeV1Job.md)|  | 

### Return type

[**MetronomeV1Job**](MetronomeV1Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1CreateJobSchedules

> MetronomeV1CreateJobSchedules V1CreateJobSchedules(ctx, jobId, metronomeV1CreateJobSchedules)


Create a new schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**|  | 
**metronomeV1CreateJobSchedules** | [**MetronomeV1CreateJobSchedules**](MetronomeV1CreateJobSchedules.md)|  | 

### Return type

[**MetronomeV1CreateJobSchedules**](MetronomeV1CreateJobSchedules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteJob

> V1DeleteJob(ctx, jobId)


Delete a job. All data about that job will be deleted.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1DeleteJobSchedulesByScheduleId

> V1DeleteJobSchedulesByScheduleId(ctx, jobId, scheduleId)


Destroy a schedule

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**|  | 
**scheduleId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetJob

> MetronomeV1Job V1GetJob(ctx, jobId, optional)


Get the job with id `jobId`. You can specify optional embed arguments to get more embedded information.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**|  | 
 **optional** | ***V1GetJobOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a V1GetJobOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **embeded** | [**optional.Interface of []MetronomeEmbeded**](MetronomeEmbeded.md)|  | 

### Return type

[**MetronomeV1Job**](MetronomeV1Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetJobSchedules

> []map[string]interface{} V1GetJobSchedules(ctx, jobId)


Get the list of all schedules for this jobId.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**|  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetJobSchedulesByScheduleId

> MetronomeV1CreateJobSchedules V1GetJobSchedulesByScheduleId(ctx, jobId, scheduleId)


Get the schedule for jobId with schedule scheduleId

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**|  | 
**scheduleId** | **string**|  | 

### Return type

[**MetronomeV1CreateJobSchedules**](MetronomeV1CreateJobSchedules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetJobs

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1PutJobSchedulesByScheduleId

> V1PutJobSchedulesByScheduleId(ctx, jobId, scheduleId, metronomeV1CreateJobSchedules)


Replaces an existing schedule.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**|  | 
**scheduleId** | **string**|  | 
**metronomeV1CreateJobSchedules** | [**MetronomeV1CreateJobSchedules**](MetronomeV1CreateJobSchedules.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1UpdateJob

> MetronomeV1Job V1UpdateJob(ctx, jobId, metronomeV1Job)


Update an existing job.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string**|  | 
**metronomeV1Job** | [**MetronomeV1Job**](MetronomeV1Job.md)|  | 

### Return type

[**MetronomeV1Job**](MetronomeV1Job.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

