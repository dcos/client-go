package dcos

import "crypto/tls"

type Config struct {
	// TODO
	Authentication Authentication
}

func (c *Config) TLS() *tls.Config {
	// if SslVerify is bool and false
	// roundtripper.(*http.Transport).TLSClientConfig.InsecureSkipVerify = false
	return &tls.Config{
		InsecureSkipVerify: false,
	}

	// otherwise add cert
	// transport.(*http.Transport).TLSClientConfig.RootCAs = x509.NewCertPool()
	// transport.(*http.Transport).TLSClientConfig.RootCAs.AddCert(cert)
}
