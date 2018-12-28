package main

import (
	"fmt"
	"os"
	"time"

	marathon "github.com/gambol99/go-marathon"
	"github.com/mesosphere/dcos-api-go/dcos"
)

func main() {
	appName := "test/foobar"
	d, err := dcos.NewClient()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	//create a marathon job
	cmd := "while true;echo test;sleep 30;done"
	mem := 128.0
	application := marathon.Application{
		ID:   appName,
		CPUs: 0.1,
		Mem:  &mem,
		Cmd:  &cmd,
	}
	createdApp, err := d.Marathon.MarathonClient.CreateApplication(&application)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("created app: %v", createdApp)

	a, err := d.Marathon.MarathonClient.Application(appName)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("app definition %v", a)

	deployment, err := d.Marathon.MarathonClient.DeleteApplication(appName, false)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	err = d.Marathon.MarathonClient.WaitOnDeployment(deployment.DeploymentID, 30*time.Second)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

}
