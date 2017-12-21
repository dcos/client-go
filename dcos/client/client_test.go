package client

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mesosphere/dcos-go/dcos/config"
	"github.com/stretchr/testify/assert"
)

// Lets do an samle request and see if we find the token header
func TestNewClient(t *testing.T) {
	testtoken := "Testtoken"
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "token="+testtoken {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("TestNewClient - Success!"))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("TestNewClient - Forbidden!"))
		}
	}))
	defer s.Close()
	testurl := s.URL
	conf := config.Config{
		Core: &config.Core{
			DcosUrl:      testurl,
			DcosAcsToken: testtoken,
		},
	}
	c := NewHttpClient(&conf)
	resp, err := c.Get(testurl)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(respBody), "Success")

	respDflt, err := http.DefaultClient.Get(testurl)
	assert.NoError(t, err)
	assert.Equal(t, 401, respDflt.StatusCode)
	respDfltbody, _ := ioutil.ReadAll(respDflt.Body)
	assert.Contains(t, string(respDfltbody), "Forbidden")
}

func TestNewClientInvalidConf(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))
	testurl := s.URL

	conf := config.Config{}

	assert.False(t, conf.Valid())

	c := NewHttpClient(&conf)
	resp, err := c.Get(testurl)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
func TestNewClientNotExistingCert(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "dcos-go-not-existing")
	assert.NoError(t, err)
	os.Remove(file.Name())

	conf := config.Config{
		Core: &config.Core{
			SslVerify:    file.Name(),
			DcosUrl:      "http://foo.bar.baz",
			DcosAcsToken: "foobar",
		},
	}

	assert.True(t, conf.Valid())

	c := NewHttpClient(&conf)
	assert.Nil(t, c)
}

func TestNewClientInvalidPermissions(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "dcos-go-invalid-permissions")
	assert.NoError(t, err)
	file.Chmod(0222)

	conf := config.Config{
		Core: &config.Core{
			SslVerify:    file.Name(),
			DcosUrl:      "http://foo.bar.baz",
			DcosAcsToken: "foobar",
		},
	}

	assert.True(t, conf.Valid())

	c := NewHttpClient(&conf)
	assert.Nil(t, c)
}

func TestNewClientInvalidPEM(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "dcos-go-invalid-pem")
	assert.NoError(t, err)

	conf := config.Config{
		Core: &config.Core{
			SslVerify:    file.Name(),
			DcosUrl:      "http://foo.bar.baz",
			DcosAcsToken: "foobar",
		},
	}

	assert.True(t, conf.Valid())

	c := NewHttpClient(&conf)
	assert.Nil(t, c)
}

func TestNewClientInvalidCert(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "dcos-go-invalid-cert")
	assert.NoError(t, err)
	b := pem.Block{
		Type:  "foobar",
		Bytes: []byte("Invalid Cert"),
	}
	out := pem.EncodeToMemory(&b)

	file.Write(out)

	conf := config.Config{
		Core: &config.Core{
			SslVerify:    file.Name(),
			DcosUrl:      "http://foo.bar.baz",
			DcosAcsToken: "foobar",
		},
	}

	assert.True(t, conf.Valid())

	c := NewHttpClient(&conf)
	assert.Nil(t, c)
}

func TestNewClientWithandWithoutVerify(t *testing.T) {
	cVerify := NewHttpClient(&config.Config{
		Core: &config.Core{
			SslVerify: "true",
		},
	})
	cNoVerify := NewHttpClient(&config.Config{
		Core: &config.Core{
			SslVerify: "false",
		},
	})

	assert.False(t, cVerify.Transport.(*Transport).Base.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
	assert.True(t, cNoVerify.Transport.(*Transport).Base.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
}

func TestNewClientWithSelfSignedSSL(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TestNewClientWithSelfSignedSSL")
	}))
	defer ts.Close()

	cert := ts.Certificate()

	file, err := ioutil.TempFile(os.TempDir(), "dcos-go-ssl-ca")
	assert.NoError(t, err)
	b := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})

	_, err = file.Write(b)
	assert.NoError(t, err)

	testurl := "http://foobar.com/baz"
	testtoken := "Testtoken"
	conf := config.Config{
		Core: &config.Core{
			DcosUrl:      testurl,
			DcosAcsToken: testtoken,
			SslVerify:    file.Name(),
		},
	}
	c := NewHttpClient(&conf)

	resp, err := c.Get(ts.URL)
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)

	// lets check we really get an error
	conf1 := config.Config{
		Core: &config.Core{
			DcosUrl:      testurl,
			DcosAcsToken: testtoken,
			SslVerify:    "true",
		},
	}

	c1 := NewHttpClient(&conf1)
	_, err1 := c1.Get(ts.URL)
	assert.Error(t, err1)
	fmt.Printf(ts.URL)
	ts.Close()
}

func TestNewClientWithoutSSLVerify(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "TestNewClientWithoutSSLVerify")
	}))
	defer ts.Close()

	testurl := "http://foobar.com/baz"
	testtoken := "Testtoken"

	conf := config.Config{
		Core: &config.Core{
			DcosUrl:      testurl,
			DcosAcsToken: testtoken,
			SslVerify:    "false",
		},
	}

	conf1 := config.Config{
		Core: &config.Core{
			DcosUrl:      testurl,
			DcosAcsToken: testtoken,
			SslVerify:    "true",
		},
	}

	c := NewHttpClient(&conf)
	c1 := NewHttpClient(&conf1)

	resp, err := c.Get(ts.URL)
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode)

	_, err = c1.Get(ts.URL)
	assert.Error(t, err)
}
