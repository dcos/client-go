/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type V2Service struct {
	Marathon V2ServiceMarathon `json:"marathon,omitempty"`
	Mesos    V2ServiceMesos    `json:"mesos,omitempty"`
	Endpoint V2Endpoint        `json:"endpoint,omitempty"`
}