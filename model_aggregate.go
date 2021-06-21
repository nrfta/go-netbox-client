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

type Aggregate struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Family *Family `json:"family,omitempty"`
	Prefix string `json:"prefix"`
	Rir *NestedRir `json:"rir"`
	Tenant *NestedTenant `json:"tenant,omitempty"`
	DateAdded string `json:"date_added,omitempty"`
	Description string `json:"description,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
