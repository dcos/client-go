package config

// Job
type Job struct {
	ServiceName string `toml:"service_name,omitempty"`
	Url         string `toml:"url,omitempty"`
}
