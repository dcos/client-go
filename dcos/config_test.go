package dcos

import (
	"crypto/x509"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	toml "github.com/pelletier/go-toml"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestConfigGetters(t *testing.T) {
	store := NewConfigStore(nil)

	store.Set("core.dcos_url", "https://dcos.example.com")
	store.Set("core.dcos_acs_token", "token_zj8Tb0vhQw")
	store.Set("core.ssl_verify", "/path/to/dcos_ca.crt")
	store.Set("core.timeout", 15)
	store.Set("cluster.name", "mr-cluster")

	config := NewConfig(store)

	require.Equal(t, "https://dcos.example.com", config.URL())
	require.Equal(t, "token_zj8Tb0vhQw", config.ACSToken())
	require.Equal(t, true, config.TLS().Insecure)
	require.Equal(t, 15*time.Second, config.Timeout())
	require.Equal(t, "mr-cluster", config.Name())
}

func TestConfigSetters(t *testing.T) {
	store := NewConfigStore(nil)

	config := NewConfig(store)
	config.SetURL("https://dcos.example.com")
	config.SetACSToken("token_XYZ")
	config.SetTLS(TLS{})
	config.SetTimeout(15 * time.Second)
	config.SetName("custom-cluster-name")

	require.Equal(t, "https://dcos.example.com", store.Get("core.dcos_url"))
	require.Equal(t, "token_XYZ", store.Get("core.dcos_acs_token"))
	require.Equal(t, "true", store.Get("core.ssl_verify"))
	require.EqualValues(t, 15, store.Get("core.timeout"))
	require.Equal(t, "custom-cluster-name", store.Get("cluster.name"))
}

func TestConfigTLSToString(t *testing.T) {
	expectedTLSStrings := []struct {
		tls TLS
		str string
	}{
		{TLS{}, "true"},
		{TLS{Insecure: true}, "false"},
		{TLS{RootCAsPath: "/path/to/ca"}, "/path/to/ca"},
	}

	for _, exp := range expectedTLSStrings {
		require.Equal(t, exp.str, exp.tls.String())
	}
}

func TestConfigGetTLS(t *testing.T) {

	expectedTLSInsecureSkipVerify := []struct {
		val      interface{}
		insecure bool
	}{
		{"True", false},
		{"1", false},
		{"true", false},
		{"False", true},
		{"false", true},
		{"0", true},
		{"/path/to/unexisting/ca", true},
	}

	for _, exp := range expectedTLSInsecureSkipVerify {
		t.Run(exp.val.(string), func(t *testing.T) {
			store := NewConfigStore(nil)
			store.Set("core.ssl_verify", exp.val)
			config := NewConfig(store)
			require.Equal(t, exp.insecure, config.TLS().Insecure)
		})
	}
}

func TestConfigGetTLSWithInvalidCA(t *testing.T) {
	store := NewConfigStore(&ConfigStoreOpts{
		Fs: afero.NewMemMapFs(),
	})

	ca := []byte(`
-----BEGIN CERTIFICATE-----
I am no authority.
-----END CERTIFICATE-----
`)
	f, _ := afero.TempFile(store.Fs(), "/", "ca")
	f.Write(ca)

	store.Set("core.ssl_verify", f.Name())
	config := NewConfig(store)

	tlsConfig := config.TLS()
	require.Equal(t, true, tlsConfig.Insecure)
	require.Equal(t, f.Name(), tlsConfig.RootCAsPath)
	require.Nil(t, tlsConfig.RootCAs)
}

func TestConfigGetTLSWithValidCA(t *testing.T) {
	store := NewConfigStore(&ConfigStoreOpts{
		Fs: afero.NewMemMapFs(),
	})

	ca := []byte(`
-----BEGIN CERTIFICATE-----
MIIDszCCApugAwIBAgIQcaz0cEq1THqqPyMRUq6YADANBgkqhkiG9w0BAQsFADCB
ijELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkNBMRYwFAYDVQQHDA1TYW4gRnJhbmNp
c2NvMRkwFwYDVQQKDBBNZXNvc3BoZXJlLCBJbmMuMTswOQYDVQQDDDJEQy9PUyBS
b290IENBIDIwMWMzZTlhLWUzZWUtNDc0MC1hZWUzLWJmMTJjMDBjYWUzMzAeFw0x
ODAzMDcwNzE1MDFaFw0yODAzMDQwNzE1MDFaMIGKMQswCQYDVQQGEwJVUzELMAkG
A1UECAwCQ0ExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xGTAXBgNVBAoMEE1lc29z
cGhlcmUsIEluYy4xOzA5BgNVBAMMMkRDL09TIFJvb3QgQ0EgMjAxYzNlOWEtZTNl
ZS00NzQwLWFlZTMtYmYxMmMwMGNhZTMzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEA36+QojAQvnJhCVeaRCeA/kC70j62OvnecaE/VYI9DkaQAgibTO/z
V99Vz+tUQMfJRuWVd4BSImZ7vOixewU5+jKgcqjuAq7lDuNNX+yz4uYkz67oaBrV
cFvHLmrqDKWeAc8EhsJIscquUCvqGuXdemjvaN22hh09GUqs1N+fMbiZUYmifOJY
+xVE62hD1U2h0FJdPHbnVmIvev7LZ7R90qoFHmNTqdkkT8Q2QkdHpcm3+bxxrtVI
7Mgf50xvLwd7pHMs60caduPIAX5PHs8BFcwYwTSt2QTWlcT2bNsBxCxwFlspCOqb
+VdKTWftDJbzTvarmo3U07i5C0BJJ0MjUQIDAQABoxMwETAPBgNVHRMBAf8EBTAD
AQH/MA0GCSqGSIb3DQEBCwUAA4IBAQB+iINvE921xcAqHo4s/4NfByNhH/XYnAyn
vfGPvb2I8ijoaj2Iab2FCrt9SdCcYQd+RPwwbe3ByPKvgSSzw/IpmniSlJ6+cKo+
cmwLp2NVvFE73YDsq5mbo3T7Zb5E7SMTWWq7fZsWFOVMA2AML6n2DcQzzjjDRdBQ
ItsQvDefqA5fNDB5LepftYbCNuk65ONGyCjpIoAw+reyzYMJorkG5Sb7AJIyFh2/
XG3O73Yy5lml6cOyz0iaX46ZaMdm+YEvisSdYGG75uX/ilEOvQObi0vUfM5f6asL
NT4Sf75bbjkawxsKnddRgK2dILw//sQdOXmSJboaStNrHS5joczy
-----END CERTIFICATE-----
`)
	f, _ := afero.TempFile(store.Fs(), "/", "ca")
	f.Write(ca)

	store.Set("core.ssl_verify", f.Name())
	config := NewConfig(store)

	tlsConfig := config.TLS()
	require.Equal(t, false, tlsConfig.Insecure)
	require.Equal(t, f.Name(), tlsConfig.RootCAsPath)

	certPool := x509.NewCertPool()
	require.True(t, certPool.AppendCertsFromPEM(ca))
	require.Equal(t, certPool, tlsConfig.RootCAs)
}

func TestConfigLoadPath(t *testing.T) {
	fs := afero.NewMemMapFs()

	cfg := `
[core]
dcos_url = "https://dcos.example.com"
dcos_acs_token = "token_zj8Tb0vhQw"
`
	f, _ := afero.TempFile(fs, "/", "config")
	f.Write([]byte(cfg))

	store := NewConfigStore(&ConfigStoreOpts{
		Fs: fs,
	})
	err := store.LoadPath(f.Name())
	require.NoError(t, err)
	require.Equal(t, "https://dcos.example.com", store.Get(configKeyURL).(string))
	require.Equal(t, "token_zj8Tb0vhQw", store.Get(configKeyACSToken).(string))
	require.Equal(t, f.Name(), store.Path())
}

func TestConfigLoadJsonPath(t *testing.T) {
	fs := afero.NewMemMapFs()

	cfg := `{
		"core": { "dcos_url": "https://toml.example.com" }
}`
	f, _ := afero.TempFile(fs, "/", "config")
	f.Write([]byte(cfg))

	store := NewConfigStore(&ConfigStoreOpts{
		Fs: fs,
	})
	err := store.LoadPath(f.Name())
	require.Error(t, err)
}

func TestConfigLoadInvalidPath(t *testing.T) {
	// Memory FS can't be used here as it creates missing parent directories with O_CREATE.
	fs := afero.NewOsFs()
	store := NewConfigStore(&ConfigStoreOpts{Fs: fs})

	tempDir, err := afero.TempDir(store.Fs(), "", "dcos_cli_tests")
	require.NoError(t, err)
	defer fs.Remove(tempDir)

	// Try to load a config path with a missing paraent directory.
	err = store.LoadPath(filepath.Join(tempDir, "unexisting_dir", "dcos.toml"))
	require.Error(t, err)
}

func TestConfigLoadEmptyString(t *testing.T) {
	store := NewConfigStore(&ConfigStoreOpts{
		Fs: afero.NewMemMapFs(),
	})
	err := store.LoadReader(strings.NewReader(""))
	require.NoError(t, err)

	require.Equal(t, make(map[string]interface{}), store.tree.ToMap())
}

func TestConfigLoadString(t *testing.T) {
	store := NewConfigStore(nil)
	err := store.LoadReader(strings.NewReader(`
[core]

dcos_url = "https://dcos.example.com"
dcos_acs_token = "token_zj8Tb0vhQw"
ssl_verify = "/path/to/dcos_ca.crt"
timeout = 15
ssh_user = "myuser"
ssh_proxy_ip = "192.0.2.1"
pagination = true
reporting = true
mesos_master_url = "https://mesos.example.com"
prompt_login = true

[cluster]

name = "mr-cluster"
`))

	require.NoError(t, err)

	require.Equal(t, "https://dcos.example.com", store.Get(configKeyURL).(string))
	require.Equal(t, "token_zj8Tb0vhQw", store.Get(configKeyACSToken).(string))
	require.Equal(t, "/path/to/dcos_ca.crt", store.Get(configKeyTLS).(string))
	require.EqualValues(t, 15, store.Get(configKeyTimeout).(int64))
	require.Equal(t, "myuser", store.Get(configKeySSHUser).(string))
	require.Equal(t, "192.0.2.1", store.Get(configKeySSHProxyHost).(string))
	require.Equal(t, true, store.Get(configKeyPagination).(bool))
	require.Equal(t, true, store.Get(configKeyReporting).(bool))
	require.Equal(t, "https://mesos.example.com", store.Get(configKeyMesosMasterURL).(string))
	require.Equal(t, true, store.Get(configKeyPromptLogin).(bool))
	require.Equal(t, "mr-cluster", store.Get(configKeyClusterName).(string))
}

func TestConfigGet(t *testing.T) {
	treeMap := map[string]interface{}{
		"core": map[string]interface{}{
			"dcos_url":   "https://dcos.example.com",
			"ssl_verify": "false",
			"timeout":    15,
			"reporting":  true,
		},
		"cluster": map[string]interface{}{
			"name": "mr-cluster",
		},
	}

	tree, err := toml.TreeFromMap(treeMap)
	require.NoError(t, err)

	store := NewConfigStore(&ConfigStoreOpts{
		EnvLookup: func(key string) (string, bool) {
			switch key {
			case "DCOS_URL":
				return "https://dcos-env.example.com", true
			case "DCOS_SSL_VERIFY":
				return "true", true
			case "DCOS_CLUSTER_NAME":
				return "dummy", true
			}
			return "", false
		},
	})
	store.LoadTree(tree)

	val := store.Get(configKeyURL)
	require.Equal(t, "https://dcos.example.com", val)

	val = store.Get(configKeyTLS)
	require.Equal(t, "true", val)

	val = store.Get(configKeyReporting)
	require.Equal(t, true, val)

	val = store.Get(configKeyTimeout)
	require.EqualValues(t, 15, val)

	val = store.Get(configKeyClusterName)
	require.Equal(t, "mr-cluster", val)
}

func TestConfigSetAndUnset(t *testing.T) {
	store := NewConfigStore(nil)

	store.Set(configKeyURL, "https://dcos.example.com")
	store.Set(configKeyTimeout, "30")
	store.Set(configKeyReporting, "1")

	require.Equal(t, nil, store.Get("core"))
	store.Unset("core")
	store.Unset("core.dcos_url.unexisting")

	require.Equal(t, "https://dcos.example.com", store.Get(configKeyURL))
	require.Equal(t, true, store.Get(configKeyReporting))
	require.EqualValues(t, 30, store.Get(configKeyTimeout))

	store.Unset(configKeyURL)
	require.Nil(t, store.Get(configKeyURL))
}

func TestConfigSetTLS(t *testing.T) {
	store := NewConfigStore(&ConfigStoreOpts{
		Fs: afero.NewMemMapFs(),
	})

	require.Error(t, store.Set(configKeyTLS, "no"))
	require.Error(t, store.Set(configKeyTLS, "/invalid/path"))

	f, _ := afero.TempFile(store.Fs(), "/", "ca")
	f.Write([]byte("Let's say I am an authority."))
	require.NoError(t, store.Set("core.ssl_verify", f.Name()))
	require.NoError(t, store.Set(configKeyTLS, "true"))
	require.NoError(t, store.Set(configKeyTLS, "false"))
	require.NoError(t, store.Set(configKeyTLS, "1"))
	require.NoError(t, store.Set(configKeyTLS, "0"))
}

func TestConfigACSTokenIsUnsetOnURLUpdate(t *testing.T) {
	store := NewConfigStore(&ConfigStoreOpts{})

	store.Set(configKeyURL, "https://dcos.example.com")
	store.Set(configKeyACSToken, "token_123")

	require.Equal(t, "https://dcos.example.com", store.Get(configKeyURL))
	require.Equal(t, "token_123", store.Get(configKeyACSToken))

	// Updating the DC/OS URL should remove the ACS token property
	store.Set(configKeyURL, "https://dcos2.example.com")
	require.Nil(t, store.Get(configKeyACSToken))

	// Unsetting the DC/OS URL should remove the ACS token property
	store.Set(configKeyACSToken, "token_123")
	store.Unset(configKeyURL)
	require.Nil(t, store.Get(configKeyURL))
}

func TestConfigKeys(t *testing.T) {
	treeMap := map[string]interface{}{
		"core": map[string]interface{}{
			"dcos_url": "https://dcos.example.com",
			"timeout":  15,
		},
		"cluster": map[string]interface{}{
			"name": "mr-cluster",
		},
		"marathon": map[string]interface{}{
			"url": "https://marathon.example.com",
		},
	}

	tree, err := toml.TreeFromMap(treeMap)
	require.NoError(t, err)

	store := NewConfigStore(&ConfigStoreOpts{
		EnvLookup: func(key string) (string, bool) {
			switch key {
			case "DCOS_URL":
				return "https://dcos-env.example.com", true
			case "DCOS_SSL_VERIFY":
				return "true", true
			}
			return "", false
		},
	})
	store.LoadTree(tree)

	expectedKeys := []string{
		"cluster.name",
		"core.dcos_url",
		"core.ssl_verify",
		"core.timeout",
		"marathon.url",
	}
	require.Equal(t, expectedKeys, store.Keys())
}

func TestConfigPersistWithoutPath(t *testing.T) {
	store := NewConfigStore(&ConfigStoreOpts{})
	require.Equal(t, ErrNoConfigPath, store.Persist())
}

func TestConfigPersist(t *testing.T) {
	fs := afero.NewMemMapFs()
	store := NewConfigStore(&ConfigStoreOpts{
		Fs: fs,
	})
	store.Set(configKeyURL, "https://dcos.example.com")

	f, _ := afero.TempFile(fs, "/", "config")
	store.SetPath(f.Name())

	require.NoError(t, store.Persist())

	contents, err := ioutil.ReadAll(f)
	require.NoError(t, err)

	expectedTOML := []byte(`
[core]
  dcos_url = "https://dcos.example.com"
`)
	require.Equal(t, expectedTOML, contents)
}

func TestConfigCurrent(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	fixturesDir := filepath.Join(wd, "testdata", "config")

	fixtures := []struct {
		name       string
		envLookup  func(key string) (val string, ok bool)
		shouldFail bool
	}{
		{"conflicting_.dcos_file", nil, true},
		{"multi_config_with_none_attached", nil, true},
		{"multiple_config_attached", nil, true},
		{"conflicting_clusters_file", nil, true},
		{"single_config_attached", nil, false},
		{"single_config_unattached", nil, false},
		{"multi_config_with_one_attached", nil, false},
		{"multi_config_with_none_attached", func(key string) (val string, ok bool) {
			if key == "DCOS_CLUSTER" {
				return "97193161-f7f1-2295-2514-a6b3918043b6", true
			}
			return
		}, false},
		{"multi_config_with_none_attached", func(key string) (val string, ok bool) {
			if key == "DCOS_CLUSTER" {
				return "multi_config_with_none_attached", true
			}
			return
		}, false},
		{"multi_config_with_none_attached", func(key string) (val string, ok bool) {
			if key == "DCOS_CLUSTER" {
				return "multi_config_with", true
			}
			return
		}, true},
	}

	for _, fixture := range fixtures {
		dcosDir := filepath.Join(fixturesDir, fixture.name, ".dcos")
		manager := NewConfigManager(&ConfigManagerOpts{
			Dir:       dcosDir,
			EnvLookup: fixture.envLookup,
		})

		t.Run(fixture.name, func(t *testing.T) {
			conf, err := manager.Current()
			if fixture.shouldFail {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, fixture.name, conf.Name())
			}
		})
	}
}

func TestConfigFind(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)

	fixturesDir := filepath.Join(wd, "testdata", "config")

	fixtures := []struct {
		name       string
		search     string
		shouldFail bool
	}{
		{"multi_config_with_same_name", "multi_config_with_same_name", true},
		{"multi_config_with_same_name", "multi", true},
		{"multi_config_with_same_name", "97193161-f7f1-2295-2514-a6b3918043b6", false},
		{"multi_config_with_same_name", "97193161", false},
	}

	for _, fixture := range fixtures {
		dcosDir := filepath.Join(fixturesDir, fixture.name, ".dcos")
		manager := NewConfigManager(&ConfigManagerOpts{
			Dir: dcosDir,
		})

		t.Run(fixture.name, func(t *testing.T) {
			conf, err := manager.Find(fixture.search, false)
			if fixture.shouldFail {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, fixture.name, conf.Name())
				require.Equal(t, "https://success.example.com", conf.URL())
			}
		})
	}
}

func TestConfigAttach(t *testing.T) {
	fs := afero.NewMemMapFs()
	clusterID := "97193161-f7f1-2295-2514-a6b3918043b6"
	fs.Create(filepath.Join(".dcos", "clusters", clusterID, "dcos.toml"))

	manager := NewConfigManager(&ConfigManagerOpts{
		Dir: ".dcos",
		Fs:  fs,
	})

	conf, err := manager.Find(clusterID, true)
	require.NoError(t, err)

	attachedFilePath := manager.attachedFilePath(conf)

	require.False(t, manager.fileExists(attachedFilePath))
	require.NoError(t, manager.Attach(conf))
	require.True(t, manager.fileExists(attachedFilePath))
}

func TestExpandHomedir(t *testing.T) {
	homedir.DisableCache = true
	os.Setenv("HOME", "/home/testuser")
	dir := expandHomeDir()

	require.Equal(t, dir, "/home/testuser/.dcos", "wrong return")
}
