# MetronomeV1JobRunRestart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | The policy to use if a job fails. NEVER will never try to relaunch a job. ON_FAILURE will try to start a job in case of failure. | 
**ActiveDeadlineSeconds** | **int32** | If the job fails, how long should we try to restart the job. If no value is set, this means forever. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


