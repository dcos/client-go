/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type MetronomeV1Job struct {
	// Unique identifier for the job consisting of a series of names separated by dots. Each name must be at least 1 character and may only contain digits (`0-9`), dashes (`-`), and lowercase letters (`a-z`). The name may not begin or end with a dash.
	Id string `json:"id"`
	// A description of this job.
	Description string `json:"description,omitempty"`
	// Attaching metadata to jobs can be useful to expose additional information to other services, so we added the ability to place labels on jobs (for example, you could label jobs staging and production to mark services by their position in the pipeline).
	Labels map[string]string `json:"labels,omitempty"`
	Run    MetronomeV1JobRun `json:"run"`
}
