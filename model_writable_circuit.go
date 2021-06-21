/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.11
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

import (
	"time"
)

type WritableCircuit struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Cid string `json:"cid"`
	Provider int32 `json:"provider"`
	Type_ int32 `json:"type"`
	Status string `json:"status,omitempty"`
	Tenant int32 `json:"tenant,omitempty"`
	InstallDate string `json:"install_date,omitempty"`
	CommitRate int32 `json:"commit_rate,omitempty"`
	Description string `json:"description,omitempty"`
	TerminationA int32 `json:"termination_a,omitempty"`
	TerminationZ int32 `json:"termination_z,omitempty"`
	Comments string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}