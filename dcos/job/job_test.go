package job

import (
	"net/http"
	"testing"

	"github.com/mesosphere/dcos-api-go/dcos/common"
	"github.com/mesosphere/dcos-api-go/dcos/config"
	"github.com/stretchr/testify/assert"
	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func newJobService() *JobService {
	testurl := "https://foobar.dcos.io"
	testtoken := "token"
	conf := &config.Config{
		Core: &config.Core{
			DcosUrl:      testurl,
			DcosAcsToken: testtoken,
		},
	}
	j := NewJobService(conf, http.DefaultClient)

	return j
}

func TestNewJobService(t *testing.T) {
	j := newJobService()
	assert.Equal(t, "https://foobar.dcos.io/service/metronome", j.GetServiceURL("").String())
	assert.Equal(t, "https://foobar.dcos.io/service/metronome/foobar", j.GetServiceURL("foobar").String())

}

func TestJobValid(t *testing.T) {
	var job, job2 Job
	job.ID = "test"
	job.Run = &Run{
		Cmd:  "echo test",
		Cpus: 0.1,
		Mem:  128,
		Disk: 50,
	}

	assert.True(t, job.Valid())
	assert.False(t, job2.Valid())
}

func TestNewJobCmd(t *testing.T) {
	job, err := NewJobWithCmd("test", "echo foo")

	assert.NoError(t, err)
	assert.True(t, job.Valid())

}

func TestJobEmptyName(t *testing.T) {
	j := newJobService()

	_, err := j.Job("")
	assert.Error(t, err)
}

func TestJobRetreive(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://foobar.dcos.io/service/metronome/v1/jobs/test",
		httpmock.NewStringResponder(200, `{"id": "test","labels": {},"run": {"cpus": 0.01,"mem": 128,"disk": 0,"cmd": "echo foo","env": {},"placement": {"constraints": []},"artifacts": [],"maxLaunchDelay": 3600,"volumes": [],"restart": {"policy": "NEVER"},"secrets": {}}}`))
	j := newJobService()

	job, err := j.Job("test")

	assert.NoError(t, err)

	assert.Equal(t, job.ID, "test")
	assert.Equal(t, job.Run.Cmd, "echo foo")

}

func TestJobNotFound(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://foobar.dcos.io/service/metronome/v1/jobs/test2",
		httpmock.NewStringResponder(404, `{"id":"test2","message":"Job not found"}`))
	j := newJobService()

	_, err := j.Job("test2")
	assert.Error(t, err)
	assert.IsType(t, &common.RequestError{}, err)
	assert.Equal(t, err.(*common.RequestError).StatusCode(), 404)
}

func TestGetSchedules(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", "https://foobar.dcos.io/service/metronome/v1/jobs/test",
		httpmock.NewStringResponder(200, `{"id": "test","labels": {},"run": {"cpus": 0.01,"mem": 128,"disk": 0,"cmd": "echo foo","env": {},"placement": {"constraints": []},"artifacts": [],"maxLaunchDelay": 3600,"volumes": [],"restart": {"policy": "NEVER"},"secrets": {}}}`))
	httpmock.RegisterResponder("GET", "https://foobar.dcos.io/service/metronome/v1/jobs/test/schedules",
		httpmock.NewStringResponder(200, `[
  {
    "id": "everyminute",
    "cron": "* * * * *",
    "concurrencyPolicy": "ALLOW",
    "enabled": true,
    "startingDeadlineSeconds": 60,
    "timezone": "America/Chicago"
  },
  {
    "id": "christmas",
    "cron": "* * * * *",
    "concurrencyPolicy": "ALLOW",
    "enabled": true,
    "startingDeadlineSeconds": 60,
    "timezone": "America/Chicago"
  }
]`))
	j := newJobService()
	schedules, err := j.GetSchedules("test")

	assert.NoError(t, err)
	assert.Equal(t, 2, len(schedules))
	assert.Equal(t, "everyminute", schedules[0].ID)

}
