/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type MetronomeV1JobSchedules struct {
	Id                      string `json:"id,omitempty"`
	Cron                    string `json:"cron,omitempty"`
	ConcurrencyPolicy       string `json:"concurrencyPolicy,omitempty"`
	Enabled                 bool   `json:"enabled,omitempty"`
	StartingDeadlineSeconds int32  `json:"startingDeadlineSeconds,omitempty"`
	Timezone                string `json:"timezone,omitempty"`
}
