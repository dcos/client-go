package dcos

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	require.IsType(t, &DefaultTransport{}, rt)
}

func TestDefaultHTTPClientAuth(t *testing.T) {
	tokenValue := "TestDefaultHTTPClientAuth-token"
	config := NewConfig(nil)
	config.SetACSToken(tokenValue)

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
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode, "using the dcos.NewHTTPClient should respond with 200")

	respDflt, err := http.DefaultClient.Get(s.URL)
	require.NoError(t, err)
	require.Equal(t, 401, respDflt.StatusCode, "expect a forbidden state with http.DefaultClient")
}
