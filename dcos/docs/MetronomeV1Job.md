# MetronomeV1Job

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the job consisting of a series of names separated by dots. Each name must be at least 1 character and may only contain digits (&#x60;0-9&#x60;), dashes (&#x60;-&#x60;), and lowercase letters (&#x60;a-z&#x60;). The name may not begin or end with a dash. | 
**Description** | **string** | A description of this job. | [optional] 
**Labels** | **map[string]string** | Attaching metadata to jobs can be useful to expose additional information to other services, so we added the ability to place labels on jobs (for example, you could label jobs staging and production to mark services by their position in the pipeline). | [optional] 
**Run** | [**MetronomeV1JobRun**](MetronomeV1Job_run.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


