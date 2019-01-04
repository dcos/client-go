package job

import (
	"fmt"
	"net/http"

	"github.com/mesosphere/dcos-api-go/dcos/common"
	"github.com/mesosphere/dcos-api-go/dcos/config"
)

// JobService represents access to package API
type JobService struct {
	*common.Service
}

// Artifact describes the jobs artifact
type Artifact struct {
	Cache      bool   `json:"cache,omitempty"`
	Executable bool   `json:"executable,omitempty"`
	Extract    bool   `json:"extract,omitempty"`
	URI        string `json:"uri"`
}

// Docker describes the docker settings for Job
type Docker struct {
	Image string `json:"image,omitempty"`
}

// Volume to be used in Job
type Volume struct {
	ContainerPath string `json:"containerPath"`
	HostPath      string `json:"hostPath"`
	Mode          string `json:"mode"`
}

// Restart defines the restart policy
type Restart struct {
	ActiveDeadlineSeconds int    `json:"activeDeadlineSeconds,omitempty"`
	Policy                string `json:"policy,omitempty"`
}

// Constraint describes
type Constraint struct {
	Attribute string `json:"attribute"`
	Operator  string `json:"operator"`
	Value     string `json:"value,omitempty"`
}

// Placement contains a list of Constraints
type Placement struct {
	Constraints []*Constraint `json:"constraints,omitempty"`
}

// Run describes a Job Run
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

// Schedule describes a Job schedule
type Schedule struct {
	ConcurrencyPolicy       string `json:"concurrencyPolicy,omitempty"`
	Cron                    string `json:"cron"`
	Enabled                 bool   `json:"enabled,omitempty"`
	ID                      string `json:"id"`
	StartingDeadlineSeconds int    `json:"startingDeadlineSeconds,omitempty"`
	Timezone                string `json:"timezone,omitempty"`
}

// Job is an metronom job
type Job struct {
	Description string            `json:"description,omitempty"`
	ID          string            `json:"id"`
	Labels      map[string]string `json:"labels,omitempty"`
	Run         *Run              `json:"run"`
	Schedules   []*Schedule       `json:"schedules,omitempty"`
}

// Valid checks if the Job struct is valid
// TODO: better to a response check. Or have abstraction expose required attributes
func (j *Job) Valid() bool {
	if j.ID != "" && j.Run.Cmd != "" && j.Run.Cpus > 0 && j.Run.Mem > 0 && j.Run.Disk >= 0 {
		return true
	}
	return false
}

// NewJobService create an *JobService instance from *config.Config and *http.Client
func NewJobService(config *config.Config, client *http.Client) *JobService {
	j := JobService{
		&common.Service{
			Config: config,
			Client: client,
			Path:   "/service/metronome",
		}}
	return &j
}

func (j *JobService) List() ([]*Job, error) {
	return nil, nil
}

// Job gets a Job from the API by its jobid
func (j *JobService) Job(jobid string) (*Job, error) {
	var job Job
	if jobid == "" {
		return nil, fmt.Errorf("jobid can't be empty")
	}
	err := j.GetJSON(fmt.Sprintf("/v1/jobs/%s", jobid), &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

// NewJobCmd creates a new Job and expect jobid and cmd as input..
func (j *JobService) NewJobCmd(jobid string, cmd string) (*Job, error) {
	job := Job{
		ID: jobid,
		Run: &Run{
			Cmd:  cmd,
			Cpus: 0.1,
			Disk: 0,
			Mem:  128,
		},
	}

	return &job, nil
}

// CreateJob posts the Job definition to the API. Job.ID has to be unique
func (j *JobService) CreateJob(job *Job) (*Job, error) {
	if !job.Valid() {
		return nil, fmt.Errorf("Job %v is invalid", job)
	}
	receiveObject := &Job{}
	err := j.PostJSON("/v0/scheduled-jobs", job, receiveObject)
	if err != nil {
		return nil, err
	}
	return receiveObject, nil
}

// UpdateJob puts a new definition for a jobid onto the API
func (j *JobService) UpdateJob(jobid string, job *Job) (*Job, error) {
	if !job.Valid() {
		return nil, fmt.Errorf("Job %v is invalid", job)
	}

	if jobid != job.ID {
		return nil, fmt.Errorf("Job.ID is different to the target jobid")
	}

	receiveObject := &Job{}
	err := j.PutJSON("/v0/scheduled-jobs", job, receiveObject)
	if err != nil {
		return nil, err
	}
	return receiveObject, nil
}

// DeleteJob deletes a job by its jobid from the API
func (j *JobService) DeleteJob(jobid string) error {
	err := j.Service.Delete(fmt.Sprintf("/v1/jobs/%s", jobid))

	if reqErr, ok := err.(*common.RequestError); ok {
		if reqErr.StatusCode() == 404 {
			return fmt.Errorf("Job ID unknown")
		}
		return reqErr
	}
	if err != nil {
		return err
	}
	return nil
}
