/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: master
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

// Identifies a specific gate and trigger match from a policy against an image and indicates it should be ignored in final policy decisions
type WhitelistItem struct {
	Id string `json:"id,omitempty"`
	Gate string `json:"gate"`
	TriggerId string `json:"trigger_id"`
}
