package marathon

import (
	"log"
	"net/http"
	"net/url"
	"path"

	marathon "github.com/gambol99/go-marathon"
	"github.com/mesosphere/dcos-api-go/dcos/config"
)

type MarathonService struct {
	config         *config.Config
	client         *http.Client
	MarathonClient marathon.Marathon
}

func (m *MarathonService) getMarathonURL() *url.URL {
	u, _ := m.config.GetDCOSUrl()
	u.Path = path.Join(u.Path, "/marathon")
	return u
}

// NewMarathonService implements a marathon service from *config.Confg and apiclient.APIClient
func NewMarathonService(config *config.Config, client *http.Client) *MarathonService {
	m := MarathonService{}
	m.client = client
	m.config = config

	mconfig := marathon.NewDefaultConfig()
	mconfig.HTTPClient = m.client
	mconfig.HTTPSSEClient = m.client
	mconfig.URL = m.getMarathonURL().String()
	mclient, err := marathon.NewClient(mconfig)
	if err != nil {
		log.Fatalf("Failed to create a client for marathon, error: %s", err)
	}

	m.MarathonClient = mclient

	return &m
}
