

# config
`import "github.com/mesosphere/dcos-api-go/dcos/config"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>



## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type Cluster](#Cluster)
* [type Config](#Config)
  * [func NewConfigFromChain() (config *Config)](#NewConfigFromChain)
  * [func NewConfigFromClusterAttached() *Config](#NewConfigFromClusterAttached)
  * [func NewConfigFromClusterEnv() *Config](#NewConfigFromClusterEnv)
  * [func NewConfigFromCommonFile() *Config](#NewConfigFromCommonFile)
  * [func NewConfigFromEnv() *Config](#NewConfigFromEnv)
  * [func NewConfigFromFile(file string) (*Config, error)](#NewConfigFromFile)
  * [func (c *Config) GetDCOSUrl() (*url.URL, error)](#Config.GetDCOSUrl)
  * [func (c *Config) SetAuthHeader(req *http.Request)](#Config.SetAuthHeader)
  * [func (c *Config) Valid() bool](#Config.Valid)
* [type Core](#Core)
* [type Job](#Job)
* [type Marathon](#Marathon)
* [type Package](#Package)


#### <a name="pkg-files">Package files</a>
[cluster.go](/src/github.com/mesosphere/dcos-api-go/dcos/config/cluster.go) [config.go](/src/github.com/mesosphere/dcos-api-go/dcos/config/config.go) [core.go](/src/github.com/mesosphere/dcos-api-go/dcos/config/core.go) [job.go](/src/github.com/mesosphere/dcos-api-go/dcos/config/job.go) [marathon.go](/src/github.com/mesosphere/dcos-api-go/dcos/config/marathon.go) [package.go](/src/github.com/mesosphere/dcos-api-go/dcos/config/package.go) 


## <a name="pkg-constants">Constants</a>
``` go
const CLUSTER_DIR = "clusters"
```
is relative to DCOS_DIR

``` go
const DCOS_DIR = ".dcos"
```
gets joined with homedir





## <a name="Cluster">type</a> [Cluster](/src/target/cluster.go?s=27:87#L4)
``` go
type Cluster struct {
    Name string `toml:"name,omitempty"`
}
```
Cluster










## <a name="Config">type</a> [Config](/src/target/config.go?s=359:560#L23)
``` go
type Config struct {
    Core     *Core     `toml:"core"`
    Cluster  *Cluster  `toml:"cluster"`
    Job      *Job      `toml:"job"`
    Marathon *Marathon `toml:"marathon"`
    Package  *Package  `toml:"package"`
}
```
Config represents the whole DC/OS client configuration.







### <a name="NewConfigFromChain">func</a> [NewConfigFromChain](/src/target/config.go?s=4541:4583#L188)
``` go
func NewConfigFromChain() (config *Config)
```
NewConfigFromChain tries all configuration methods in an order and
merges the results. First Common File `~/.dcos/dcos.toml` then attached
cluster and afterwards the cluster set with `DCOS_CLUSTER`


### <a name="NewConfigFromClusterAttached">func</a> [NewConfigFromClusterAttached](/src/target/config.go?s=2534:2577#L110)
``` go
func NewConfigFromClusterAttached() *Config
```
NewConfigFromClusterAttached looks for an attached cluster (`dcos cluster attach [...]`)
and reads it into and Config object


### <a name="NewConfigFromClusterEnv">func</a> [NewConfigFromClusterEnv](/src/target/config.go?s=3261:3299#L141)
``` go
func NewConfigFromClusterEnv() *Config
```
NewConfigFromClusterEnv tries to find a cluster by its ID, Name or URL and
returns its configuration


### <a name="NewConfigFromCommonFile">func</a> [NewConfigFromCommonFile](/src/target/config.go?s=2213:2251#L96)
``` go
func NewConfigFromCommonFile() *Config
```
NewConfigFromCommonFile expects and `dcos.toml` in `DCOS_DIR` by
default `DCOS_DIR` is `$HOME/.dcos`


### <a name="NewConfigFromEnv">func</a> [NewConfigFromEnv](/src/target/config.go?s=4261:4292#L180)
``` go
func NewConfigFromEnv() *Config
```
TODO: Module needed to marshall from nested env Vars
NewConfigFromEnv reads the configuration from Environment variables.
`DCOS_<section>_<confiration>=<value>`. This is not yet implemented.


### <a name="NewConfigFromFile">func</a> [NewConfigFromFile](/src/target/config.go?s=1820:1872#L78)
``` go
func NewConfigFromFile(file string) (*Config, error)
```
NewConfigFromFile reads a `.toml` file and unmarshalls it into an Config
struct. It returns an empty Config and an error if the could not be read.





### <a name="Config.GetDCOSUrl">func</a> (\*Config) [GetDCOSUrl](/src/target/config.go?s=1344:1391#L60)
``` go
func (c *Config) GetDCOSUrl() (*url.URL, error)
```



### <a name="Config.SetAuthHeader">func</a> (\*Config) [SetAuthHeader](/src/target/config.go?s=1190:1239#L53)
``` go
func (c *Config) SetAuthHeader(req *http.Request)
```
SetAuthHeader adds the token header to the provided *http.Request

TODO: check if it makes sense to have an config refresh method befor adding
the header. The best case would be using refresh tokens and oauth2 in here
but thats currently not supported by DC/OS.




### <a name="Config.Valid">func</a> (\*Config) [Valid](/src/target/config.go?s=628:657#L32)
``` go
func (c *Config) Valid() bool
```
Valid returns true if configuration has the least options set.




## <a name="Core">type</a> [Core](/src/target/core.go?s=24:474#L4)
``` go
type Core struct {
    DcosAcsToken   string `toml:"dcos_acs_token,omitempty"`
    DcosUrl        string `toml:"dcos_url,omitempty"`
    MesosMasterUrl string `toml:"mesos_master_url,omitempty"`
    Pagination     bool   `toml:"pagination,omitempty"`
    PromptLogin    bool   `toml:"prompt_login,omitempty"`
    Reporting      bool   `toml:"reporting,omitempty"`
    SslVerify      string `toml:"ssl_verify,omitempty"`
    Timeout        int    `toml:"timeout,omitempty"`
}
```
Core










## <a name="Job">type</a> [Job](/src/target/job.go?s=23:137#L4)
``` go
type Job struct {
    ServiceName string `toml:"service_name,omitempty"`
    Url         string `toml:"url,omitempty"`
}
```
Job










## <a name="Marathon">type</a> [Marathon](/src/target/marathon.go?s=28:87#L4)
``` go
type Marathon struct {
    Url string `toml:"url,omitempty"`
}
```
Marathon










## <a name="Package">type</a> [Package](/src/target/package.go?s=27:98#L4)
``` go
type Package struct {
    CosmosUrl string `toml:"cosmos_url,omitempty"`
}
```
Package














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
