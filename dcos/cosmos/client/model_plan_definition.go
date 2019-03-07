/*
 * Cosmos
 *
 * DC/OS Package and Service
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

type PlanDefinition struct {
	Status string      `json:"status"`
	Phases []PlanPhase `json:"phases,omitempty"`
	Errors []string    `json:"errors,omitempty"`
}
