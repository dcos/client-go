# MetronomeV1JobRun

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | **[]string** | An array of strings that represents an alternative mode of specifying the command to run. This was motivated by safe usage of containerizer features like a custom Docker ENTRYPOINT. Either &#x60;cmd&#x60; or &#x60;args&#x60; must be supplied. It is invalid to supply both &#x60;cmd&#x60; and &#x60;args&#x60; in the same job. | [optional] 
**Artifacts** | [**[]MetronomeV1JobRunArtifacts**](MetronomeV1Job_run_artifacts.md) | Provided URIs are passed to Mesos fetcher module and resolved in runtime. | [optional] 
**Cmd** | **string** | The command that is executed.  This value is wrapped by Mesos via &#x60;/bin/sh -c ${job.cmd}&#x60;.  Either &#x60;cmd&#x60; or &#x60;args&#x60; must be supplied. It is invalid to supply both &#x60;cmd&#x60; and &#x60;args&#x60; in the same job. | [optional] 
**Cpus** | **float32** | The number of CPU shares this job needs per instance. This number does not have to be integer, but can be a fraction. | 
**Gpus** | **int32** | The number of GPU shares this job needs per instance. This number does not have to be integer, but can be a fraction. | [optional] 
**Disk** | **float32** | How much disk space is needed for this job. This number does not have to be an integer, but can be a fraction. | 
**Ucr** | [**MetronomeV1JobRunUcr**](MetronomeV1Job_run_ucr.md) |  | [optional] 
**Docker** | [**MetronomeV1JobRunDocker**](MetronomeV1Job_run_docker.md) |  | [optional] 
**Env** | [**map[string]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 
**MaxLaunchDelay** | **int32** | The number of seconds until the job needs to be running. If the deadline is reached without successfully running the job, the job is aborted. | [optional] 
**Mem** | **float32** | The amount of memory in MB that is needed for the job per instance. | 
**Placement** | [**MetronomeV1JobRunPlacement**](MetronomeV1Job_run_placement.md) |  | [optional] 
**User** | **string** | The user to use to run the tasks on the agent. | [optional] 
**TaskKillGracePeriodSeconds** | **float32** | Configures the number of seconds between escalating from SIGTERM to SIGKILL when signalling tasks to terminate. Using this grace period, tasks should perform orderly shut down immediately upon receiving SIGTERM. | [optional] 
**Restart** | [**MetronomeV1JobRunRestart**](MetronomeV1Job_run_restart.md) |  | [optional] 
**Volumes** | [**[]MetronomeV1JobRunVolumes**](MetronomeV1Job_run_volumes.md) | The list of volumes for this job. | [optional] 
**Secrets** | [**map[string]map[string]interface{}**](map[string]interface{}.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


