package main

import (
	"log"
	"time"

	"github.com/dcos/client-go/dcos"
	marathon "github.com/gambol99/go-marathon"
)

func createApp(marathonClient marathon.Marathon, app *marathon.Application) error {
	_, err := marathonClient.CreateApplication(app)
	if err != nil {
		switch err := err.(type) {
		case *marathon.APIError:
			if err.ErrCode == marathon.ErrCodeDuplicateID {
				log.Printf("Application %s already installed", app.ID)
				return nil
			}
		}
		return err
	}

	log.Printf("Marathon application created: %+v\n", app)

	return nil
}

func listApps(marathonClient marathon.Marathon) error {
	applications, err := marathonClient.Applications(nil)
	if err != nil {
		return err
	}

	log.Printf("Found %d application(s) running", len(applications.Apps))
	for _, application := range applications.Apps {
		log.Printf("Application: %+v\n", application)
	}
	return nil
}

func main() {
	client, err := dcos.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	dcosConfig := client.CurrentDCOSConfig()

	marathonConfig := marathon.NewDefaultConfig()
	marathonConfig.URL = dcosConfig.URL() + "/service/marathon"
	marathonConfig.HTTPClient = client.HTTPClient()
	marathonClient, err := marathon.NewClient(marathonConfig)
	if err != nil {
		log.Fatalf("Couldn't create Marathon client: %s\n", err)
	}
	marathonApp := marathon.NewDockerApplication().
		Name("appzzz").
		CPU(0.1).
		Memory(64).
		Count(2).
		AddArgs("sleep", "36000").
		AddEnv("MY_ENV", "my_val")
	marathonApp.Container.Docker.Container("alpine:latest")

	err = createApp(marathonClient, marathonApp)
	if err != nil {
		log.Fatalf("Creating Marathon app failed: %s\n", err)
	}

	err = listApps(marathonClient)
	if err != nil {
		log.Fatalf("Failed to list Marathon applications: %s", err)
	}

	deploymentID, err := marathonClient.DeleteApplication(marathonApp.ID, true)
	if err != nil {
		log.Fatalf("Failed to delete Marathon application: %s", err)
	}

	err = marathonClient.WaitOnDeployment(deploymentID.DeploymentID, 60*time.Second)
	if err != nil {
		log.Fatalf("Error waiting for deployment %s: %s", deploymentID.DeploymentID, err)
	}

}
