package dcos

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultTransportBaseNil(t *testing.T) {
	c := &DefaultTransport{
		Base: nil,
	}
	rt := c.base()

	assert.ObjectsAreEqual(http.DefaultTransport, rt)
}

func TestDefaultTransportBase(t *testing.T) {
	c := &DefaultTransport{
		Base: &DefaultTransport{Base: nil},
	}
	rt := c.base()

	assert.IsType(t, &DefaultTransport{}, rt)
}

func TestDefaultHTTPClientAuth(t *testing.T) {
	tokenValue := "TestDefaultHTTPClientAuth-token"
	config := &Config{
		Authentication: &TestAuthentication{Token: tokenValue},
	}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "token="+tokenValue {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("TestDefaultHTTPClientAuth - Success"))
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("TestDefaultHTTPClientAuth - Forbidden"))
		}
	}))
	defer s.Close()

	c := NewHTTPClient(config)

	resp, err := c.Get(s.URL)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(t, string(respBody), "Success")

	respDflt, err := http.DefaultClient.Get(s.URL)
	assert.NoError(t, err)
	assert.Equal(t, 401, respDflt.StatusCode)
	respDfltbody, _ := ioutil.ReadAll(respDflt.Body)
	assert.Contains(t, string(respDfltbody), "Forbidden")

}
