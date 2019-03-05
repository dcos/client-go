package dcos

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

const (
	// defaultTransportDialTimeout specifies the maximum amount of time
	// waiting for a connection to be established.
	defaultTransportDialTimeout = 10 * time.Second

	// defaultTransportTLSHandshakeTimeout specifies the maximum
	// amount of time waiting to wait for a TLS handshake.
	defaultTransportTLSHandshakeTimeout = 10 * time.Second

	// defaultTransportMaxIdleConns specifies the maximum number of idle connections.
	defaultTransportMaxIdleConns = 30
)

// DefaultTransport is a http.RoundTripper that adds authentication based on Config
type DefaultTransport struct {
	Config *Config
	Base   http.RoundTripper
}

func (t *DefaultTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

// RoundTrip authorizes requests to DC/OS by adding dcos_acs_token to Authorization header
func (t *DefaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// meet the requirements of RoundTripper and only modify a copy
	req2 := cloneRequest(req)
	req2.Header.Set("Authorization", fmt.Sprintf("token=%s", t.Config.ACSToken()))

	return t.base().RoundTrip(req2)
}

func cloneRequest(req *http.Request) *http.Request {
	req2 := new(http.Request)
	*req2 = *req

	// until now we only clone headers as we only modify those.
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}

	return req2
}

// NewHTTPClient provides a http.Client able to communicate to dcos in an authenticated way
func NewHTTPClient(config *Config) (*http.Client, error) {
	if config == nil {
		return nil, fmt.Errorf("Config should not be nil")
	}
	client := &http.Client{}
	client.Transport = &http.Transport{

		// Allow http_proxy, https_proxy, and no_proxy.
		Proxy: http.ProxyFromEnvironment,

		DialContext: (&net.Dialer{
			Timeout: defaultTransportDialTimeout,
		}).DialContext,

		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.TLS().Insecure,
			RootCAs:            config.TLS().RootCAs,
		},

		TLSHandshakeTimeout: defaultTransportTLSHandshakeTimeout,

		// As the client is dealing with a single host (the DC/OS master node),
		// set both MaxIdleConns and MaxIdleConnsPerHost to the same value.
		MaxIdleConns:        defaultTransportMaxIdleConns,
		MaxIdleConnsPerHost: defaultTransportMaxIdleConns,
	}

	return AddTransportHTTPClient(client, config), nil
}

// AddTransportHTTPClient adds dcos.DefaultTransport to http.Client to add dcos authentication
func AddTransportHTTPClient(client *http.Client, config *Config) *http.Client {
	transport := DefaultTransport{
		Config: config,
		Base:   client.Transport,
	}

	client.Transport = &transport

	return client
}
