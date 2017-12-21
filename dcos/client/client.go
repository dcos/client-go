package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

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

type APIClient struct {
	HTTPClient *http.Client
}

// NewAPIClient creates a new APIClient and adds a HTTPclient from Config
func NewAPIClient(conf *config.Config) *APIClient {
	c := APIClient{}
	c.HTTPClient = NewHttpClient(conf)
	return &c
}

// NewAPIClient creates a new APIClient and adds a HTTPclient from Config
func NewAPIClientWithHTTPClient(hc *http.Client) *APIClient {
	c := APIClient{}
	c.HTTPClient = hc
	return &c
}

func (a *APIClient) Post(url, contentType, accept string, input, output interface{}) (err error) {
	in, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(in))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", accept)
	req.Header.Add("Content-Type", contentType)

	resp, err := a.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("Unauthorized. You're not allowed to request this enpoint")
	}

	if resp.StatusCode >= 400 {
		b, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Request error got: %d - body: %s", resp.StatusCode, b)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, output)
	if err != nil {
		return err
	}

	return nil
}

func (a *APIClient) Get(url, output interface{}) (err error) {
	return nil
}

// NewHttpClient creates an *http.Client from an DCOS configuration *config.Config
func NewHttpClient(conf *config.Config) *http.Client {
	// Taken from http.DefaultTransport
	var transport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{},
	}

	if conf.Core != nil && conf.Core.SslVerify != "" {

		sslVerify, err := strconv.ParseBool(conf.Core.SslVerify)
		if err != nil {
			// we expect a file
			if _, err := os.Stat(conf.Core.SslVerify); os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "dcos-go fatal: could not find specied certificate at %s", conf.Core.SslVerify)
				return nil
			}
			certfile, err := ioutil.ReadFile(conf.Core.SslVerify)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dcos-go fatal: read certificate from %s", conf.Core.SslVerify)
				return nil
			}
			p, _ := pem.Decode(certfile)
			if p == nil {
				fmt.Fprintf(os.Stderr, "dcos-go fatal: not a pem file %s", conf.Core.SslVerify)
				return nil
			}
			cert, err := x509.ParseCertificate(p.Bytes)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dcos-go fatal: parse certificate from %s - %s", conf.Core.SslVerify, err)
				return nil
			}

			transport.(*http.Transport).TLSClientConfig.RootCAs = x509.NewCertPool()
			transport.(*http.Transport).TLSClientConfig.RootCAs.AddCert(cert)
		} else {
			transport.(*http.Transport).TLSClientConfig.RootCAs = nil
			transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = !sslVerify
		}
	}

	return &http.Client{
		Transport: &Transport{
			Config: conf,
			Base:   transport,
		},
	}
}
