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

type Prefix struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Family *Family `json:"family,omitempty"`
	// IPv4 or IPv6 network with mask
	Prefix string `json:"prefix"`
	Site *NestedSite `json:"site,omitempty"`
	Vrf *NestedVrf `json:"vrf,omitempty"`
	Tenant *NestedTenant `json:"tenant,omitempty"`
	Vlan *NestedVlan `json:"vlan,omitempty"`
	Status *Status8 `json:"status,omitempty"`
	Role *NestedRole `json:"role,omitempty"`
	// All IP addresses within this prefix are considered usable
	IsPool bool `json:"is_pool,omitempty"`
	Description string `json:"description,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	Children int32 `json:"children,omitempty"`
	Depth int32 `json:"_depth,omitempty"`
}