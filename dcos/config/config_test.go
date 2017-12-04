package config

import (
	"os"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	homedir.DisableCache = true
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestConfigValid(t *testing.T) {
	conf := Config{}

	assert.Equal(t, false, conf.Valid())

	conf.Core = &Core{DcosUrl: "https://mydcostest.dcos.io", DcosAcsToken: "foo"}
	assert.Equal(t, true, conf.Valid())

	conf.Core.DcosUrl = "wss://mydcostest.dcos.io"
	assert.Equal(t, false, conf.Valid())

	conf.Core.DcosUrl = "http://mydcostest.dcos.io"
	assert.Equal(t, true, conf.Valid())

	conf.Core.DcosUrl = "äöü://foo.bar"
	assert.Equal(t, false, conf.Valid())
}

func TestGetDCOSDir(t *testing.T) {
	os.Clearenv()
	os.Setenv("HOME", "./foobar")
	assert.Equal(t, "./foobar", os.Getenv("HOME"))
	dir, err := getDCOSDir()

	assert.NoError(t, err)

	assert.Equal(t, "foobar/.dcos", dir)

	os.Clearenv()
	os.Setenv("HOME", "./foobar2")
	dir, err = getDCOSDir()

	assert.NoError(t, err)

	assert.Equal(t, "foobar2/.dcos", dir)
}

func TestgetDCOSDirDCOSDIROverwrite(t *testing.T) {
	os.Clearenv()
	os.Setenv("HOME", "./foobar")
	os.Setenv("DCOS_DIR", "./baz")
	dir, err := getDCOSDir()

	assert.NoError(t, err)

	assert.Equal(t, "./baz/.dcos", dir)
}

func TestNewConfigFromFile(t *testing.T) {
	config, err := NewConfigFromFile("./test/dcos.toml")

	assert.NoError(t, err)

	assert.Equal(t, "https://mydcoscluster.mesosphere.io", config.Core.DcosUrl)
	assert.NotEqual(t, config.Core.DcosAcsToken, ".dcos-folder")

	assert.Equal(t, "False", config.Core.SslVerify)
	assert.Equal(t, "MainDCOSCluster", config.Cluster.Name)
	assert.Equal(t, "Job Service Name", config.Job.ServiceName)
	assert.Equal(t, "https://mydcoscluster.mesosphere.io/jobservice", config.Job.Url)
	assert.Equal(t, "https://mydcoscluster.mesosphere.io/marathon", config.Marathon.Url)
	assert.Equal(t, "https://localhost:7070", config.Package.CosmosUrl)
}

func TestNewConfigFromInvalidFile(t *testing.T) {
	config, err := NewConfigFromFile("./test/dcos-invalid.toml")

	assert.Error(t, err)

	assert.NotEqual(t, config.Core.DcosUrl, "https://mydcoscluster.mesosphere.io")
	assert.NotEqual(t, config.Core.SslVerify, "False")
}

func TestNewConfigFromNotExistentFile(t *testing.T) {
	config, err := NewConfigFromFile("./test/dcos-notexistent.toml")

	assert.Error(t, err)

	assert.Nil(t, config.Cluster)
	assert.Nil(t, config.Core)
	assert.Nil(t, config.Job)
	assert.Nil(t, config.Marathon)
	assert.Nil(t, config.Package)
}

func TestNewConfigFromNotReadableFile(t *testing.T) {
	filename := "./test/dcos-notreadable.toml"
	os.Chmod(filename, 0200)
	config, err := NewConfigFromFile(filename)

	assert.Error(t, err)

	assert.Nil(t, config.Cluster)
	assert.Nil(t, config.Core)
	assert.Nil(t, config.Job)
	assert.Nil(t, config.Marathon)
	assert.Nil(t, config.Package)
	os.Chmod(filename, 0644)
}

func TestNewConfigFromCommonFile(t *testing.T) {
	os.Clearenv()
	os.Setenv("HOME", "./test")

	dir, err := getDCOSDir()
	assert.NoError(t, err)

	assert.Equal(t, "test/.dcos", dir)
	config := NewConfigFromCommonFile()

	assert.NotNil(t, config.Cluster)
	assert.Equal(t, config.Core.DcosUrl, "https://mydcoscluster.mesosphere.io")
	assert.Equal(t, config.Core.DcosAcsToken, ".dcos-folder")
}

func TestNewConfigFromCommonFileMissing(t *testing.T) {
	os.Setenv("HOME", "./test2")
	config := NewConfigFromCommonFile()

	assert.Nil(t, config.Cluster)
	assert.Nil(t, config.Core)
	assert.Nil(t, config.Job)
	assert.Nil(t, config.Marathon)
	assert.Nil(t, config.Package)
}

func TestNewConfigFromCommonFileDCOS_DIR(t *testing.T) {
	os.Setenv("DCOS_DIR", "./test")
	config := NewConfigFromCommonFile()

	assert.NotNil(t, config.Cluster)
	assert.Equal(t, config.Core.DcosUrl, "https://mydcoscluster.mesosphere.io")
	assert.NotEqual(t, config.Core.DcosAcsToken, ".dcos-folder")
	os.Unsetenv("DCOS_DIR")
}

func TestNewConfigFromClusterAttached(t *testing.T) {
	os.Setenv("HOME", "./test")
	config := NewConfigFromClusterAttached()

	assert.NotNil(t, config.Cluster)
	assert.NotNil(t, config.Core)
	assert.Equal(t, "https://myattacheddcoscluster1.mesosphere.io", config.Core.DcosUrl)
	assert.Equal(t, "AttachedCluster", config.Cluster.Name)
}

func TestNewConfigFromClusterAttachedWrongFilename(t *testing.T) {
	os.Setenv("HOME", "./clusterwrongname")
	config := NewConfigFromClusterAttached()
	assert.Nil(t, config.Cluster)
	assert.Nil(t, config.Core)
	assert.Nil(t, config.Job)
	assert.Nil(t, config.Marathon)
	assert.Nil(t, config.Package)
}

func TestNewConfigFromClusterEnvUnattached(t *testing.T) {
	os.Setenv("HOME", "./test")
	os.Setenv("DCOS_CLUSTER", "a223535e-b3e8-48ca-99b8-4b6df8079a4e")
	config := NewConfigFromClusterEnv()
	assert.NotNil(t, config.Cluster)
	assert.NotNil(t, config.Core)
	assert.Equal(t, "https://mydcoscluster2.mesosphere.io", config.Core.DcosUrl)
	assert.Equal(t, "UnAttachedCluster", config.Cluster.Name)
}

func TestNewConfigFromClusterEnvEmpty(t *testing.T) {
	os.Setenv("DCOS_CLUSTER", "")
	config := NewConfigFromClusterEnv()
	assert.Nil(t, config.Cluster)
	assert.Nil(t, config.Core)
	assert.Nil(t, config.Job)
	assert.Nil(t, config.Marathon)
	assert.Nil(t, config.Package)
}

func TestNewConfigFromClusterEnvClusterName(t *testing.T) {
	os.Setenv("HOME", "./test")
	os.Setenv("DCOS_CLUSTER", "UnAttachedCluster2")
	config := NewConfigFromClusterEnv()
	assert.NotNil(t, config.Cluster)
	assert.NotNil(t, config.Core)
	assert.Equal(t, "https://mydcosclusterunattached2.mesosphere.io", config.Core.DcosUrl)
	assert.Equal(t, "UnAttachedCluster2", config.Cluster.Name)
}

func TestNewConfigFromClusterEnvClusterURL(t *testing.T) {
	os.Setenv("HOME", "./test")
	os.Setenv("DCOS_CLUSTER", "https://mydcosclusterunattached2.mesosphere.io")
	config := NewConfigFromClusterEnv()
	assert.NotNil(t, config.Cluster)
	assert.NotNil(t, config.Core)
	assert.Equal(t, "https://mydcosclusterunattached2.mesosphere.io", config.Core.DcosUrl)
	assert.Equal(t, "UnAttachedCluster2", config.Cluster.Name)
}

func TestNewConfigFromChain(t *testing.T) {
	os.Setenv("HOME", "./test/chainconfig")
	os.Setenv("DCOS_CLUSTER", "376ac56b-7bc4-4bb1-a6d9-0e9f5e1f5b50")
	config := NewConfigFromChain()
	assert.Equal(t, ".dcos.toml file", config.Core.DcosUrl)
	assert.Equal(t, "attached cluster", config.Core.DcosAcsToken)
	assert.Equal(t, "unattached cluster", config.Core.SslVerify)
	assert.NotNil(t, config.Job)
	assert.NotNil(t, config.Marathon)
	assert.NotNil(t, config.Package)
}
