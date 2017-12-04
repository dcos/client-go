package config

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/imdario/mergo"
	homedir "github.com/mitchellh/go-homedir"
)

// gets joined with homedir
const DCOS_DIR = ".dcos"

// is relative to DCOS_DIR
const CLUSTER_DIR = "clusters"

// Config reoresents the whike DC/OS client configuration.
type Config struct {
	Core     *Core     `toml:"core"`
	Cluster  *Cluster  `toml:"cluster"`
	Job      *Job      `toml:"job"`
	Marathon *Marathon `toml:"marathon"`
	Package  *Package  `toml:"package"`
}

// Valid returns true if configuration has the least options set.
func (c *Config) Valid() bool {
	if c.Core != nil && c.Core.DcosUrl != "" && c.Core.DcosAcsToken != "" {
		u, err := url.Parse(c.Core.DcosUrl)
		if err != nil {
			return false
		}

		if u.Scheme != "http" && u.Scheme != "https" {
			return false
		}

		return true
	}
	return false
}

func getDCOSDir() (string, error) {
	dir := os.Getenv("DCOS_DIR")
	if dir == "" {
		hd, err := homedir.Dir()
		if err != nil {
			return "", fmt.Errorf("Could not get Homedir")
		}
		dir = path.Join(hd, DCOS_DIR)
	}
	return dir, nil
}

// NewConfigFromFile reads a `.toml` file and unmarshalls it into an Config
// struct. It returns an empty Config and an error if the could not be read.
func NewConfigFromFile(file string) (*Config, error) {
	config := Config{}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return &config, err
	}

	f, err := ioutil.ReadFile(file)
	if err != nil {
		return &config, err
	}

	err = toml.Unmarshal(f, &config)
	return &config, err
}

// NewConfigFromCommonFile expects and `dcos.toml` in `DCOS_DIR` by
// default `DCOS_DIR` is `$HOME/.dcos`
func NewConfigFromCommonFile() *Config {
	config := Config{}
	dir, err := getDCOSDir()
	if err != nil {
		return &config
	}

	cnf, _ := NewConfigFromFile(dir + "/dcos.toml")
	return cnf

}

// NewConfigFromClusterAttached looks for an attached cluster (`dcos cluster attach [...]`)
// and reads it into and Config object
func NewConfigFromClusterAttached() *Config {
	config := Config{}
	dir, err := getDCOSDir()
	if err != nil {
		return &config
	}
	clusterDir := path.Join(dir, CLUSTER_DIR)

	files, err := ioutil.ReadDir(clusterDir)
	if err != nil {
		return &config
	}

	for _, files := range files {
		if files.IsDir() {
			clusterPath := path.Join(clusterDir, "/", files.Name())
			if _, err := os.Stat(path.Join(clusterPath, "attached")); os.IsNotExist(err) {
				continue
			}
			conf, err := NewConfigFromFile(path.Join(clusterPath, "/dcos.toml"))
			if err != nil {
				return &config
			}
			return conf
		}
	}
	return &config
}

// NewConfigFromClusterEnv tries to find a cluster by its ID, Name or URL and
// returns its configuration
func NewConfigFromClusterEnv() *Config {
	config := Config{}
	if cluster := os.Getenv("DCOS_CLUSTER"); cluster != "" {
		dir, err := getDCOSDir()
		if err != nil {
			return &config
		}
		clusterDir := path.Join(dir, CLUSTER_DIR)

		files, err := ioutil.ReadDir(clusterDir)
		if err != nil {
			return &config
		}

		for _, files := range files {
			if files.IsDir() {
				clusterPath := path.Join(clusterDir, "/", files.Name())
				conf, err := NewConfigFromFile(path.Join(clusterPath, "/dcos.toml"))
				if err != nil {
					return &config
				}
				if files.Name() == cluster {
					return conf
				}
				if conf.Cluster != nil && conf.Cluster.Name == cluster {
					return conf
				}
				if conf.Core != nil && conf.Core.DcosUrl == cluster {
					return conf
				}
			}
		}
	}
	return &config
}

// TODO: Module needed to marshall from nested env Vars
// NewConfigFromEnv reads the configuration from Environment variables.
// `DCOS_<section>_<confiration>=<value>`. This is not yet implemented.
func NewConfigFromEnv() *Config {
	config := Config{}
	return &config
}

// NewConfigFromChain tries all configuration methods in an order and
// merges the results. First Common File `~/.dcos/dcos.toml` then attached
// cluster and afterwards the cluster set with `DCOS_CLUSTER`
func NewConfigFromChain() (config *Config) {
	config = NewConfigFromEnv()
	mergo.Merge(config, NewConfigFromClusterEnv())
	mergo.Merge(config, NewConfigFromClusterAttached())
	mergo.Merge(config, NewConfigFromCommonFile())
	return config
}
