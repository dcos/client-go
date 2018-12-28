package dcos

import (
	"context"
	"fmt"
	"net/http"

	apiclient "github.com/mesosphere/dcos-api-go/dcos/client"
	"github.com/mesosphere/dcos-api-go/dcos/config"
	"github.com/mesosphere/dcos-api-go/dcos/dcospackage"
	"github.com/mesosphere/dcos-api-go/dcos/job"
	"github.com/mesosphere/dcos-api-go/dcos/marathon"
)

// A Client manages communication with the DC/OS API.
type Client struct {
	ctx    *context.Context
	client *http.Client // API client used to communicate with the API.

	config *config.Config

	Marathon *marathon.MarathonService
	Package  *dcospackage.PackageService
	Job      *job.JobService
}

func NewClient() (*Client, error) {
	ctx := context.TODO()
	return NewClientWithContext(&ctx)
}

func NewClientWithContext(ctx *context.Context) (*Client, error) {
	c := Client{}
	c.ctx = ctx

	c.config = config.NewConfigFromChain()
	if !c.config.Valid() {
		return &c, fmt.Errorf("Configuration invalid")
	}
	c.client = apiclient.NewHttpClient(c.config)

	c.Marathon = marathon.NewMarathonService(c.config, c.client)
	c.Package = dcospackage.NewPackageService(c.config, c.client)
	c.Job = job.NewJobService(c.config, c.client)

	return &c, nil
}
