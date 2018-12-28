package client

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mesosphere/dcos-api-go/dcos/config"
)

const (
	libraryVersion = "1"
	userAgent      = "dcos-go/" + libraryVersion
)

// Transport is an http.RoundTripper that makes Authorized Requests,
// wrapping a base RoundTripper and adding an Authorization header
// with a token from the config.
// This basic logic is taken from https://github.com/google/go-github
type Transport struct {
	// The DC/OS config to get the token from
	Config *config.Config

	UserAgent string

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

	if t.UserAgent != "" {
		req.Header.Set("User-Agent", t.UserAgent)
	}

	res, err := t.base().RoundTrip(req)
	if err != nil {
		return nil, err
	}

	return res, nil
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
			UserAgent: userAgent,
			Config:    conf,
			Base:      transport,
		},
	}
}
