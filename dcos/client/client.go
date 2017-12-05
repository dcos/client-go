package client

import (
	"errors"
	"net/http"

	"github.com/mesosphere/dcos-go/dcos/config"
)

// Transport is an http.RoundTripper that makes Authorized Requests,
// wrapping a base RoundTripper and adding an Authorization header
// with a token from the config.
// This basic logic is taken from https://github.com/google/go-github
type Transport struct {
	// The DC/OS config to get the token from
	Config *config.Config

	// Base is the base RoundTripper used to make HTTP requests.
	// If nil, http.DefaultTransport is used.
	Base http.RoundTripper
}

func (t *Transport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// RoundTrip authorizes and authenticates the request with an
// access token.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Config == nil || !t.Config.Valid() {
		return nil, errors.New("dcos-go client: Transport's Config is nil")
	}

	//set the auth header
	t.Config.SetAuthHeader(req)

	res, err := t.base().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// NewClient creates an *http.Client from an DCOS configuration *config.Config
func NewClient(conf *config.Config) *http.Client {
	c := http.Client{
		Transport: &Transport{
			Config: conf,
		},
	}
	return &c
}
