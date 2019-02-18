package dcos

import (
	"net"
	"net/http"
	"time"
)

const (
	DefaultHTTPClientDialContextTimeout  = 10 * time.Second
	DefaultHTTPClientTLSHandshakeTimeout = 10 * time.Second
	DefaultHTTPClientMaxIdleConns        = 30
	DefaultHTTPClientMaxIdleConnsPerHost = 30
)

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

func (t *DefaultTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// meet the requirements of RoundTripper and only modify a copy
	req2 := cloneRequest(req)

	// TODO: we might need to check Expiry before.
	t.Config.Authentication.AddHeaders(req2)

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
func NewHTTPClient(config *Config) *http.Client {
	client := &http.Client{}
	client.Transport = &http.Transport{

		// Allow http_proxy, https_proxy, and no_proxy.
		Proxy: http.ProxyFromEnvironment,

		// Set a 10 seconds timeout for the connection to be established.
		DialContext: (&net.Dialer{
			Timeout: DefaultHTTPClientDialContextTimeout,
		}).DialContext,

		// Set it to 10 seconds as well for the TLS handshake when using HTTPS.
		TLSHandshakeTimeout: DefaultHTTPClientTLSHandshakeTimeout,

		// The client will be dealing with a single host (the one in baseURL),
		// set max idle connections to 30 regardless of the host.
		MaxIdleConns:        DefaultHTTPClientMaxIdleConns,
		MaxIdleConnsPerHost: DefaultHTTPClientMaxIdleConnsPerHost,
	}

	// Set the TLS configuration as specified in the context.
	config.TLS(client.Transport)

	return ConfigureHTTPClient(config, client)
}

// ConfigureHTTPClient adds dcos.DefaultTransport to http.Client to add dcos authentication
func ConfigureHTTPClient(config *Config, client *http.Client) *http.Client {
	transport := DefaultTransport{
		Config: config,
		Base:   client.Transport,
	}

	client.Transport = &transport

	return client
}
