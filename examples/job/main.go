package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mesosphere/dcos-api-go/dcos"
	"github.com/mesosphere/dcos-api-go/dcos/job"
)

func main() {
	d, err := dcos.NewClient()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	schedule := &job.Schedule{
		ID:      "default",
		Cron:    "0/1 * * * *",
		Enabled: true,
	}

	j := job.Job{
		ID: "testjob1",
		Run: &job.Run{
			Cmd:  "echo test",
			Cpus: 0.1,
			Mem:  128,
			Disk: 50,
		},
		Schedules: []*job.Schedule{schedule},
	}

	_, err = d.Job.CreateJob(&j)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	time.Sleep(20 * time.Second)

	job, err := d.Job.Job(j.ID)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("%v", job)

	time.Sleep(20 * time.Second)

	err = d.Job.DeleteJob(j.ID)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

}
