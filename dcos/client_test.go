package dcos

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	homedir.DisableCache = true
	wd, err := os.Getwd()
	require.NoError(t, err)
	testdir := filepath.Join(wd, "testdata", "config", "single_config_attached")
	os.Setenv("HOME", testdir)

	c, err := NewClient()

	require.NoErrorf(t, err, "NewClient returned unexpected error: %s - testdir: %s", err, testdir)
	require.Equal(t, "single_config_attached", c.Config.Name())
	require.IsTypef(t, &DefaultTransport{}, c.HTTPClient.Transport, "HTTPClient.Transport type different")
}

func TestNewClientWithOptions(t *testing.T) {
	timeout := 7 * time.Second
	baseClient := &http.Client{Timeout: timeout}
	config := NewConfig(nil)
	config.SetACSToken("test")
	c, err := NewClientWithOptions(baseClient, config)

	require.NoErrorf(t, err, "NewClientWithOptions returned unexpected error: %s", err)
	require.Equal(t, baseClient, c.HTTPClient, "NewClientWithOptions should leave HTTPClient unchanged")

	_, err = NewClientWithOptions(nil, config)
	require.Error(t, err)

	_, err = NewClientWithOptions(baseClient, nil)
	require.Error(t, err)
}
