/*
 * Cosmos
 *
 * DC/OS Package and Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

type SystemNode struct {
	HostIp string `json:"host_ip,omitempty"`
	Health int32  `json:"health,omitempty"`
	Role   string `json:"role,omitempty"`
}
