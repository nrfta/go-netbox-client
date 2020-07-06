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

type WritableVlan struct {
	Id int32 `json:"id,omitempty"`
	Site int32 `json:"site,omitempty"`
	Group int32 `json:"group,omitempty"`
	Vid int32 `json:"vid"`
	Name string `json:"name"`
	Tenant int32 `json:"tenant,omitempty"`
	Status string `json:"status,omitempty"`
	Role int32 `json:"role,omitempty"`
	Description string `json:"description,omitempty"`
	Tags []string `json:"tags,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	PrefixCount int32 `json:"prefix_count,omitempty"`
}