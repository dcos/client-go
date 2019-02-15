package dcos

import (
	"net/http"
)

type Config struct {
	// TODO
}

func (c *Config) TLS(roundtripper http.RoundTripper) {
	// if SslVerify is bool and false
	roundtripper.(*http.Transport).TLSClientConfig.InsecureSkipVerify = false

	// otherwise add cert
	// transport.(*http.Transport).TLSClientConfig.RootCAs = x509.NewCertPool()
	// transport.(*http.Transport).TLSClientConfig.RootCAs.AddCert(cert)
}
