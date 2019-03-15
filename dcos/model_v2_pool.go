/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

// The pool contains information on resources that the pool needs. Changes make to this section will relaunch the tasks.
type V2Pool struct {
	// Pool tasks healthcheck grace period (in seconds)
	PoolHealthcheckGracePeriod int32 `json:"poolHealthcheckGracePeriod,omitempty"`
	// Pool tasks healthcheck interval (in seconds)
	PoolHealthcheckInterval int32 `json:"poolHealthcheckInterval,omitempty"`
	// Pool tasks healthcheck maximum number of consecutive failures before declaring as unhealthy
	PoolHealthcheckMaxFail int32 `json:"poolHealthcheckMaxFail,omitempty"`
	// Maximum amount of time that Mesos will wait for the healthcheck container to finish executing
	PoolHealthcheckTimeout int32      `json:"poolHealthcheckTimeout,omitempty"`
	ApiVersion             ApiVersion `json:"apiVersion,omitempty"`
	// The pool name.
	Name string `json:"name,omitempty"`
	// The DC/OS space (sometimes also referred to as a \"group\").
	Namespace      string `json:"namespace,omitempty"`
	PackageName    string `json:"packageName,omitempty"`
	PackageVersion string `json:"packageVersion,omitempty"`
	// Mesos role for load balancers. Defaults to \"slave_public\" so that load balancers will be run on public agents. Use \"*\" to run load balancers on private agents. Read more about Mesos roles at http://mesos.apache.org/documentation/latest/roles/
	Role string `json:"role,omitempty"`
	// Mesos principal for pool framework authentication. If omitted or left blank, the service account used to install Edge-LB will be used if present.
	Principal string `json:"principal,omitempty"`
	// Service account secret name for pool framework authentication. If omitted or left blank, the service account used to install Edge-LB will be used if present.
	SecretName        string  `json:"secretName,omitempty"`
	Cpus              float32 `json:"cpus,omitempty"`
	CpusAdminOverhead float32 `json:"cpusAdminOverhead,omitempty"`
	// Memory requirements (in MB)
	Mem int32 `json:"mem,omitempty"`
	// Memory requirements (in MB)
	MemAdminOverhead int32 `json:"memAdminOverhead,omitempty"`
	// Disk size (in MB)
	Disk int32 `json:"disk,omitempty"`
	// Number of load balancer instances in the pool.
	Count int32 `json:"count,omitempty"`
	// Marathon style constraints for load balancer instance placement.
	Constraints string `json:"constraints,omitempty"`
	// Override ports to allocate for each load balancer instance. Defaults to {{haproxy.frontends[].bindPort}} and   {{haproxy.stats.bindPort}}. Use this field to pre-allocate all needed ports with or   without the frontends present. For example: [80, 443, 9090]. If the length of the ports array is not zero, only the   ports specified will be allocated by the pool scheduler.
	Ports []int32 `json:"ports,omitempty"`
	// DC/OS secrets.
	Secrets []V2PoolSecrets `json:"secrets,omitempty"`
	// Environment variables to pass to tasks. Prefix with \"ELB_FILE_\" and it will be written to a file. For example, the contents of \"ELB_FILE_MYENV\" will be written to \"$ENVFILE/ELB_FILE_MYENV\".
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
	// Autogenerate a self-signed SSL/TLS certificate. It is not generated by default. It will be written to \"$AUTOCERT\".
	AutoCertificate bool `json:"autoCertificate,omitempty"`
	// Virtual networks to join.
	VirtualNetworks []V2PoolVirtualNetworks `json:"virtualNetworks,omitempty"`
	Haproxy         V2Haproxy               `json:"haproxy,omitempty"`
}
