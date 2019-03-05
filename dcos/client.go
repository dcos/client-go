package dcos

import (
	"fmt"
	"net/http"
)

// Client is a client for DC/OS.
type Client struct {
	HTTPClient *http.Client
	Config     *Config
	UserAgent  string
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

	return &Client{
		HTTPClient: httpClient,
		Config:     config,
		UserAgent:  fmt.Sprintf("%s(%s)", ClientName, Version),
	}, nil
}
