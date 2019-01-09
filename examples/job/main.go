package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mesosphere/dcos-api-go/dcos"
	"github.com/mesosphere/dcos-api-go/dcos/job"
)

func errExit(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

func main() {
	d, err := dcos.NewClient()

	errExit(err)

	schedule := &job.Schedule{
		ID:      "default",
		Cron:    "0/1 * * * *",
		Enabled: true,
	}

	j, err := job.NewJobWithCmd("testjob1", "echo test")

	errExit(err)

	_, err = d.Job.CreateJob(j)
	errExit(err)
	time.Sleep(20 * time.Second)

	job, err := d.Job.Job(j.ID)
	errExit(err)
	fmt.Printf("%#v", job)
	time.Sleep(20 * time.Second)

	s, err := d.Job.CreateSchedule(job.ID, schedule)
	errExit(err)
	fmt.Printf("%#v", s)
	time.Sleep(20 * time.Second)

	scheds, err := d.Job.GetSchedules(job.ID)
	errExit(err)
	fmt.Printf("%#v", scheds)
	time.Sleep(20 * time.Second)

	schedule.Enabled = false
	_, err = d.Job.UpdateSchedule(job.ID, "default", schedule)
	errExit(err)

	sched, err := d.Job.GetSchedule(job.ID, "default")
	errExit(err)
	fmt.Printf("%#v", sched)
	time.Sleep(20 * time.Second)

	err = d.Job.DeleteSchedule(job.ID, "default")
	errExit(err)
	fmt.Printf("%#v", sched)
	time.Sleep(20 * time.Second)

	err = d.Job.DeleteJob(j.ID)
	errExit(err)

}
