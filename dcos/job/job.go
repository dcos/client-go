package job

import (
	"fmt"
	"net/http"

	"github.com/mesosphere/dcos-api-go/dcos/common"
	"github.com/mesosphere/dcos-api-go/dcos/config"
)

// PackageService represents access to package API
type JobService struct {
	*common.Service
}

// Job is an metronom job
type Job struct {
	Description string            `json:"description,omitempty"`
	ID          string            `json:"id"`
	Labels      map[string]string `json:"labels,omitempty"`
	Run         struct {
		Args      []string `json:"args,omitempty"`
		Artifacts []struct {
			Cache      bool   `json:"cache,omitempty"`
			Executable bool   `json:"executable,omitempty"`
			Extract    bool   `json:"extract,omitempty"`
			URI        string `json:"uri"`
		} `json:"artifacts,omitempty"`
		Cmd    string  `json:"cmd,omitempty"`
		Cpus   float64 `json:"cpus"`
		Disk   float64 `json:"disk"`
		Docker struct {
			Image string `json:"image"`
		} `json:"docker,omitempty"`
		Env            map[string]string `json:"env,omitempty"`
		MaxLaunchDelay int               `json:"maxLaunchDelay,omitempty"`
		Mem            float64           `json:"mem"`
		Placement      *struct {
			Constraints []struct {
				Attribute string `json:"attribute"`
				Operator  string `json:"operator"`
				Value     string `json:"value,omitempty"`
			} `json:"constraints,omitempty"`
		} `json:"placement,omitempty"`
		Restart struct {
			ActiveDeadlineSeconds int    `json:"activeDeadlineSeconds,omitempty"`
			Policy                string `json:"policy"`
		} `json:"restart,omitempty"`
		User    string `json:"user,omitempty"`
		Volumes []struct {
			ContainerPath string `json:"containerPath"`
			HostPath      string `json:"hostPath"`
			Mode          string `json:"mode"`
		} `json:"volumes,omitempty"`
	} `json:"run"`
	Schedules []struct {
		ConcurrencyPolicy       string `json:"concurrencyPolicy,omitempty"`
		Cron                    string `json:"cron"`
		Enabled                 bool   `json:"enabled,omitempty"`
		ID                      string `json:"id"`
		StartingDeadlineSeconds int    `json:"startingDeadlineSeconds,omitempty"`
		Timezone                string `json:"timezone,omitempty"`
	} `json:"schedules,omitempty"`
}

// Valid checks if the Job struct is valid
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

func (j *JobService) List() {

}

func (j *JobService) Create(job *Job) (*Job, error) {
	if !job.Valid() {
		return nil, fmt.Errorf("Job %v is invalid", job)
	}
	receiveObject := &Job{}
	err := j.PostJSON("/v1/jobs", job, receiveObject)
	if err != nil {
		return nil, err
	}
	return receiveObject, nil
}
