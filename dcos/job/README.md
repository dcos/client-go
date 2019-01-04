

# job
`import "github.com/mesosphere/dcos-api-go/dcos/job"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [type Artifact](#Artifact)
* [type Constraint](#Constraint)
* [type Docker](#Docker)
* [type Job](#Job)
  * [func (j *Job) Valid() bool](#Job.Valid)
* [type JobService](#JobService)
  * [func NewJobService(config *config.Config, client *http.Client) *JobService](#NewJobService)
  * [func (j *JobService) CreateJob(job *Job) (*Job, error)](#JobService.CreateJob)
  * [func (j *JobService) DeleteJob(jobid string) error](#JobService.DeleteJob)
  * [func (j *JobService) Job(jobid string) (*Job, error)](#JobService.Job)
  * [func (j *JobService) List() ([]*Job, error)](#JobService.List)
  * [func (j *JobService) NewJobCmd(jobid string, cmd string) (*Job, error)](#JobService.NewJobCmd)
  * [func (j *JobService) UpdateJob(jobid string, job *Job) (*Job, error)](#JobService.UpdateJob)
* [type Placement](#Placement)
* [type Restart](#Restart)
* [type Run](#Run)
* [type Schedule](#Schedule)
* [type Volume](#Volume)


#### <a name="pkg-files">Package files</a>
[job.go](/src/github.com/mesosphere/dcos-api-go/dcos/job/job.go) 






## <a name="Artifact">type</a> [Artifact](/src/target/job.go?s=275:470#L17)
``` go
type Artifact struct {
    Cache      bool   `json:"cache,omitempty"`
    Executable bool   `json:"executable,omitempty"`
    Extract    bool   `json:"extract,omitempty"`
    URI        string `json:"uri"`
}
```
Artifact describes the jobs artifact










## <a name="Constraint">type</a> [Constraint](/src/target/job.go?s=970:1112#L43)
``` go
type Constraint struct {
    Attribute string `json:"attribute"`
    Operator  string `json:"operator"`
    Value     string `json:"value,omitempty"`
}
```
Constraint describes










## <a name="Docker">type</a> [Docker](/src/target/job.go?s=520:581#L25)
``` go
type Docker struct {
    Image string `json:"image,omitempty"`
}
```
Docker describes the docker settings for Job










## <a name="Job">type</a> [Job](/src/target/job.go?s=2482:2767#L82)
``` go
type Job struct {
    Description string            `json:"description,omitempty"`
    ID          string            `json:"id"`
    Labels      map[string]string `json:"labels,omitempty"`
    Run         *Run              `json:"run"`
    Schedules   []*Schedule       `json:"schedules,omitempty"`
}
```
Job is an metronom job










### <a name="Job.Valid">func</a> (\*Job) [Valid](/src/target/job.go?s=2896:2922#L92)
``` go
func (j *Job) Valid() bool
```
Valid checks if the Job struct is valid
TODO: better to a response check. Or have abstraction expose required attributes




## <a name="JobService">type</a> [JobService](/src/target/job.go?s=190:233#L12)
``` go
type JobService struct {
    *common.Service
}
```
JobService represents access to package API







### <a name="NewJobService">func</a> [NewJobService](/src/target/job.go?s=3134:3208#L100)
``` go
func NewJobService(config *config.Config, client *http.Client) *JobService
```
NewJobService create an *JobService instance from *config.Config and *http.Client





### <a name="JobService.CreateJob">func</a> (\*JobService) [CreateJob](/src/target/job.go?s=4042:4096#L144)
``` go
func (j *JobService) CreateJob(job *Job) (*Job, error)
```
CreateJob posts the Job definition to the API. Job.ID has to be unique




### <a name="JobService.DeleteJob">func</a> (\*JobService) [DeleteJob](/src/target/job.go?s=4831:4881#L175)
``` go
func (j *JobService) DeleteJob(jobid string) error
```
DeleteJob deletes a job by its jobid from the API




### <a name="JobService.Job">func</a> (\*JobService) [Job](/src/target/job.go?s=3448:3500#L115)
``` go
func (j *JobService) Job(jobid string) (*Job, error)
```
Job gets a Job from the API by its jobid




### <a name="JobService.List">func</a> (\*JobService) [List](/src/target/job.go?s=3338:3381#L110)
``` go
func (j *JobService) List() ([]*Job, error)
```



### <a name="JobService.NewJobCmd">func</a> (\*JobService) [NewJobCmd](/src/target/job.go?s=3772:3842#L129)
``` go
func (j *JobService) NewJobCmd(jobid string, cmd string) (*Job, error)
```
NewJobCmd creates a new Job and expect jobid and cmd as input..




### <a name="JobService.UpdateJob">func</a> (\*JobService) [UpdateJob](/src/target/job.go?s=4386:4454#L157)
``` go
func (j *JobService) UpdateJob(jobid string, job *Job) (*Job, error)
```
UpdateJob puts a new definition for a jobid onto the API




## <a name="Placement">type</a> [Placement](/src/target/job.go?s=1158:1241#L50)
``` go
type Placement struct {
    Constraints []*Constraint `json:"constraints,omitempty"`
}
```
Placement contains a list of Constraints










## <a name="Restart">type</a> [Restart](/src/target/job.go?s=794:944#L37)
``` go
type Restart struct {
    ActiveDeadlineSeconds int    `json:"activeDeadlineSeconds,omitempty"`
    Policy                string `json:"policy,omitempty"`
}
```
Restart defines the restart policy










## <a name="Run">type</a> [Run](/src/target/job.go?s=1270:2038#L55)
``` go
type Run struct {
    Args           []string          `json:"args,omitempty"`
    Artifacts      []*Artifact       `json:"artifacts,omitempty"`
    Cmd            string            `json:"cmd,omitempty"`
    Cpus           float64           `json:"cpus"`
    Disk           float64           `json:"disk"`
    Docker         *Docker           `json:"docker,omitempty"`
    Env            map[string]string `json:"env,omitempty"`
    MaxLaunchDelay int               `json:"maxLaunchDelay,omitempty"`
    Mem            float64           `json:"mem"`
    Placement      *Placement        `json:"placement,omitempty"`
    Restart        *Restart          `json:"restart,omitempty"`
    User           string            `json:"user,omitempty"`
    Volumes        []*Volume         `json:"volumes,omitempty"`
}
```
Run describes a Job Run










## <a name="Schedule">type</a> [Schedule](/src/target/job.go?s=2077:2454#L72)
``` go
type Schedule struct {
    ConcurrencyPolicy       string `json:"concurrencyPolicy,omitempty"`
    Cron                    string `json:"cron"`
    Enabled                 bool   `json:"enabled,omitempty"`
    ID                      string `json:"id"`
    StartingDeadlineSeconds int    `json:"startingDeadlineSeconds,omitempty"`
    Timezone                string `json:"timezone,omitempty"`
}
```
Schedule describes a Job schedule










## <a name="Volume">type</a> [Volume](/src/target/job.go?s=611:754#L30)
``` go
type Volume struct {
    ContainerPath string `json:"containerPath"`
    HostPath      string `json:"hostPath"`
    Mode          string `json:"mode"`
}
```
Volume to be used in Job














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
