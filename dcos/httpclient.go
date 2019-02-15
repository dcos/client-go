package dcos

import (
	"net"
	"net/http"
	"time"
)

const (
	userAgent = ClientName + " - " + Version
)

type DCOSTransport struct {
	Config *Config
	Base   http.RoundTripper
}

func (t *DCOSTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}

func (t *DCOSTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := t.base().RoundTrip(req)

	req.Header.Set("User-Agent", userAgent)

	return res, err
}

func NewHTTPClient(config *Config) *http.Client {
	client := http.DefaultClient
	client.Transport = &http.Transport{

		// Allow http_proxy, https_proxy, and no_proxy.
		Proxy: http.ProxyFromEnvironment,

		// Set a 10 seconds timeout for the connection to be established.
		DialContext: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).DialContext,

		// Set it to 10 seconds as well for the TLS handshake when using HTTPS.
		TLSHandshakeTimeout: 10 * time.Second,

		// The client will be dealing with a single host (the one in baseURL),
		// set max idle connections to 30 regardless of the host.
		MaxIdleConns:        30,
		MaxIdleConnsPerHost: 30,
	}

	// Set the TLS configuration as specified in the context.
	config.TLS(client.Transport)

	return NewHttpClientWithHTTPClient(config, client)
}

func NewHttpClientWithHTTPClient(config *Config, client *http.Client) *http.Client {
	transport := DCOSTransport{
		Config: config,
		Base:   client.Transport,
	}

	client.Transport = &transport

	return client
}
