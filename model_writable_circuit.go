/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

import (
	"time"
)

type WritableCircuit struct {
	Id int32 `json:"id,omitempty"`
	Cid string `json:"cid"`
	Provider int32 `json:"provider"`
	Type_ int32 `json:"type"`
	Status string `json:"status,omitempty"`
	Tenant int32 `json:"tenant,omitempty"`
	InstallDate string `json:"install_date,omitempty"`
	CommitRate int32 `json:"commit_rate,omitempty"`
	Description string `json:"description,omitempty"`
	TerminationA string `json:"termination_a,omitempty"`
	TerminationZ string `json:"termination_z,omitempty"`
	Comments string `json:"comments,omitempty"`
	Tags []string `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}