package main

import (
	"context"
	"log"
	"time"

	"github.com/dcos/client-go/dcos"
)

func main() {
	client, err := dcos.NewClient()
	ctx := context.TODO()
	if err != nil {
		log.Fatal(err)
	}

	// Create job
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

	mv1job1, resp, err1 := client.Metronome.V1CreateJob(ctx, metronomeV1Job)
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Printf("V1CreateJob object: %+v\n", mv1job1)
	log.Printf("V1CreateJob Status code: %s\n", resp.Status)

	time.Sleep(5 * time.Second)

	// Create job schedule
	v1mcjs := dcos.MetronomeV1JobSchedule{
		Id:                      jobID,
		Cron:                    "1 1 1 1 1",
		ConcurrencyPolicy:       "ALLOW",
		Enabled:                 true,
		StartingDeadlineSeconds: 60,
		Timezone:                "America/Chicago",
	}
	mjs, resp1, err2 := client.Metronome.V1CreateJobSchedules(ctx, jobID, v1mcjs)
	if err2 != nil {
		log.Println(err2)
	}
	log.Printf("V1CreateJobSchedules object: %+v\n", mjs)
	log.Printf("V1CreateJobSchedules Status code: %s\n", resp1.Status)

	// Get job schedule
	mv1js, resp3, err3 := client.Metronome.V1GetJobSchedules(ctx, jobID)
	if err3 != nil {
		log.Println(err3)
	}
	log.Printf("V1GetJobSchedules: %+v\n", mv1js)
	log.Printf("V1GetJobSchedules Status code: %s\n", resp3.Status)

	// Start a run
	mv1job2, resp4, err4 := client.Metronome.V1StartJobRun(ctx, jobID)
	if err4 != nil {
		log.Println(err4)
	}
	log.Printf("V1StartJobRun object: %+v\n", mv1job2)
	log.Printf("V1StartJobRun Status code: %s\n", resp4.Status)

	time.Sleep(5 * time.Second)

	// Get job by run ID
	mv1job3, resp5, err6 := client.Metronome.V1GetJobRunByRunId(ctx, jobID, mv1job2.Id)
	if err6 != nil {
		log.Println(err6)
	}
	log.Printf("V1GetJobRunByRunId: %+v\n", mv1job3)
	log.Printf("V1GetJobRunByRunId Status code: %s\n", resp5.Status)

	// Stop job by run ID
	resp6, e := client.Metronome.V1StopJobRunByRunId(ctx, jobID, mv1job2.Id)
	if e != nil {
		log.Println(e)
	}
	log.Printf("V1StopJobRunByRunId Status code: %s\n", resp6.Status)

	time.Sleep(5 * time.Second)

	// Get Schedule
	mv1js2, resp7, err9 := client.Metronome.V1GetJobSchedulesByScheduleId(ctx, metronomeV1Job.Id, v1mcjs.Id)
	if err9 != nil {
		log.Println(err9)
	}
	log.Printf("V1GetJobSchedulesByScheduleId object: %+v\n", mv1js2)
	log.Printf("V1GetJobSchedulesByScheduleId Status code: %s\n", resp7.Status)

	// Delete schedule
	resp7, err8 := client.Metronome.V1DeleteJobSchedulesByScheduleId(ctx, metronomeV1Job.Id, v1mcjs.Id)
	if err8 != nil {
		log.Println(err8)
	}
	log.Printf("V1DeleteJobSchedulesByScheduleId Status code: %s\n", resp7.Status)

	// Delete job
	resp5, err5 := client.Metronome.V1DeleteJob(ctx, jobID)
	if err5 != nil {
		log.Fatal(err5)
	}
	log.Printf("V1DeleteJob Status code: %s\n", resp5.Status)
}
