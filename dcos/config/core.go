package config

// Core
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
