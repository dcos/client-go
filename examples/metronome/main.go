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

	mJobs := getJobs(ctx, client)

	initialJobs := len(mJobs)
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
	_, resp, err := client.Metronome.V1CreateJob(ctx, metronomeV1Job)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created Job - %s\n", resp.Status)

	mJobs2 := getJobs(ctx, client)

	if len(mJobs2) == initialJobs {
		log.Fatal("no job created")
	}

	for _, job := range mJobs {
		fmt.Printf("Found Job %s with cmd %s\n", job.Id, job.Run.Cmd)
	}

	resp, err = client.Metronome.V1DeleteJob(ctx, jobID)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Deleted Job %s\n", resp.Status)

}
