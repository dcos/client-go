package dcos

import (
	"fmt"
	"net/http"
)

// Client
type Client struct {
	HTTPClient *http.Client
	Config     *Config
	UserAgent  string
}

// NewClient returns a Client which will detect its Config through the well known
// process and use a http.Client with DC/OS authentication.
func NewClient() (*Client, error) {
	return NewClientWithOptions(nil, nil)
}

// NewClientWithOptions creates a Client from a given *http.Client and *Config
// if nil is given for one or both options defaults will be taken
//
//   dcosClient, err := NewClientWithOptions(nil,nil)
//
func NewClientWithOptions(httpClient *http.Client, config *Config) (*Client, error) {
	if config == nil {
		currentConfig, err := NewConfigManager(DefaultConfigManagerOpts).Current()
		if err != nil {
			return nil, err
		}
		config = currentConfig
	}

	if httpClient == nil {
		hc, err := NewHTTPClient(config)
		if err != nil {
			return nil, err
		}
		httpClient = hc
	}
	return &Client{
		HTTPClient: httpClient,
		Config:     config,
		UserAgent:  fmt.Sprintf("%s(%s)", ClientName, Version),
	}, nil
}
