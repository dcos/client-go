/*
 * Cosmos
 *
 * DC/OS Package and Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

type Binaries struct {
	Darwin  Os `json:"darwin,omitempty"`
	Linux   Os `json:"linux,omitempty"`
	Windows Os `json:"windows,omitempty"`
}
