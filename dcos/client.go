package dcos

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dcos/client-go/cosmos"
	"github.com/go-openapi/runtime"

	httptransport "github.com/go-openapi/runtime/client"
)

// Client
type Client struct {
	HTTPClient *http.Client
	Config     *Config
	UserAgent  string
	Cosmos     *cosmos.Client
}

// NewClient returns a Client which will detect its Config through the well
// known process and use a default http.Client with DC/OS authentication.
func NewClient() (*Client, error) {
	config, err := NewConfigManager(nil).Current()
	if err != nil {
		return nil, err
	}

	httpClient, err := NewHTTPClient(config)
	if err != nil {
		return nil, err
	}

	return NewClientWithOptions(httpClient, config)
}

// NewClientWithOptions creates a Client from a given *http.Client and *Config
func NewClientWithOptions(httpClient *http.Client, config *Config) (*Client, error) {
	if httpClient == nil {
		return nil, fmt.Errorf("httpClient cannot be nil")
	}

	if config == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	transport, _ := newTransport(httpClient, config)

	return &Client{
		HTTPClient: httpClient,
		Config:     config,
		UserAgent:  fmt.Sprintf("%s(%s)", ClientName, Version),
		Cosmos:     cosmos.New(transport, nil),
	}, nil
}

func newTransport(httpClient *http.Client, config *Config) (runtime.ClientTransport, error) {
	configurl, _ := url.Parse(config.URL())
	return httptransport.NewWithClient(configurl.Hostname(), configurl.Path, nil, httpClient), nil
}
