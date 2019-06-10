package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dcos/client-go/dcos"
)

func getJobs(ctx context.Context, client *dcos.APIClient) []dcos.MetronomeV1Job {
	mJobs, _, err := client.Metronome.V1GetJobs(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	return mJobs
}

func main() {
	client, err := dcos.NewClient()
	ctx := context.TODO()
	if err != nil {
		log.Fatal(err)
	}
	// Get Jobs probably empty list
	mJobs := getJobs(ctx, client)

	//initialJobs := len(mJobs)
	for _, job := range mJobs {
		fmt.Printf("Found Job %s with cmd %s\n", job.Id, job.Run.Cmd)
	}

	jobID := "testjob1"

	metronomeV1Job := dcos.MetronomeV1Job{
		Id: jobID,
		Run: dcos.MetronomeV1JobRun{
			Cmd:  "echo foobar",
			Cpus: 0.1,
			Mem:  128,
			Disk: 10.0,
		},
	}
	// Create Job
	_, resp, err := client.Metronome.V1CreateJob(ctx, metronomeV1Job)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created Job - %s\n", resp.Status)

	resp1, e := client.Metronome.V1StopJobRunByRunId(ctx, jobID, "20190607000515a584s")
	if e != nil {
		log.Println(e)
	}
	log.Printf("V1StopJobRunByRunId Status code: %s\n", resp1.Status)

	/*
		val, resp1, e := client.Metronome.V1GetJobRunByRunId(ctx, jobID, "201906062355158xfU7")
		if e != nil {
			log.Println(e)
		}
		log.Printf("V1GetJobRunByRunId: %+v\n", val)
		log.Printf("V1GetJobRunByRunId Status code: %s\n", resp1.Status)
	*/

	/*
		val, resp1, e := client.Metronome.V1StartJobRun(ctx, jobID)
		if e != nil {
			log.Println(e)
		}
		log.Printf("V1StartJobRun runs: %+v\n", val)
		log.Printf("V1StartJobRun Status code: %s\n", resp1.Status)
	*/

	/*
		resp1, e := client.Metronome.V1DeleteJobSchedulesByScheduleId(ctx, jobID, jobID)
		if e != nil {
			log.Println(e)
		}
		log.Printf("V1DeleteJobSchedulesByScheduleId Status code: %s\n", resp1.Status)
	*/

	/*
		v1mcjs := dcos.MetronomeV1CreateJobSchedules{
			Id:                      jobID,
			Cron:                    "1 1 1 1 1",
			ConcurrencyPolicy:       "ALLOW",
			Enabled:                 true,
			StartingDeadlineSeconds: 60,
			Timezone:                "America/Chicago",
		}
		resp1, e := client.Metronome.V1PutJobSchedulesByScheduleId(ctx, jobID, jobID, v1mcjs)
		if e != nil {
			log.Println(e)
		}
		log.Printf("V1PutJobSchedulesByScheduleId Status code: %s\n", resp1.Status)
	*/

	/*
		m, resp1, e := client.Metronome.V1GetJobSchedulesByScheduleId(ctx, jobID, jobID)
		if e != nil {
			log.Println(e)
		}
		log.Printf("V1GetJobSchedulesByScheduleId: %+v\n", m)
		log.Printf("Status code: %s\n", resp1.Status)
	*/

	/*
		v1mcjs := dcos.MetronomeV1CreateJobSchedules{
			Id:                      jobID,
			Cron:                    "* * * * *",
			ConcurrencyPolicy:       "ALLOW",
			Enabled:                 true,
			StartingDeadlineSeconds: 60,
			Timezone:                "America/Chicago",
		}
		m, resp1, e := client.Metronome.V1CreateJobSchedules(ctx, jobID, v1mcjs)
		if e != nil {
			log.Println(e)
		}
		log.Printf("V1CreateJobSchedules: %+v\n", m)
		log.Printf("Status code: %s\n", resp1.Status)
	*/

	/*
		m2, resp2, err := client.Metronome.V1GetJobSchedules(ctx, jobID)
		if err != nil {
			log.Println(err)
		}
		log.Printf("V1GetJobSchedules: %+v\n", m2)
		log.Printf("Status code: %s\n", resp2.Status)
	*/

	/*
		mJobs2 := getJobs(ctx, client)

		if len(mJobs2) == initialJobs {
			log.Fatal("no job created")
		}
		// Get Jobs 2nd time
		for _, job := range mJobs {
			fmt.Printf("Found Job %s with cmd %s\n", job.Id, job.Run.Cmd)
		}
		// Get Job
		retJob, _, err := client.Metronome.V1GetJob(ctx, jobID, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Job with ID: %s found - %#v", jobID, retJob)

		updateJobDescription := "Updated jobs description"
		// Update Job
		retJob.Description = updateJobDescription
		_, _, err = client.Metronome.V1UpdateJob(ctx, jobID, retJob)
		if err != nil {
			fmt.Printf("Error: update did not work %v\n", err)
		} else {
			fmt.Printf("Updated job with id %s\n", jobID)

		}

		// Get job again and compare Description
		retupdatedJob, _, err := client.Metronome.V1GetJob(ctx, jobID, nil)
		if err != nil {
			log.Fatal(err)
		}

		if retupdatedJob.Description != updateJobDescription {
			fmt.Printf("Error: update did not work\n")
		} else {
			fmt.Printf("Update successfull\n")
		}
	*/

	/*
		resp, err = client.Metronome.V1DeleteJob(ctx, jobID)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Deleted Job %s\n", resp.Status)
	*/
}
