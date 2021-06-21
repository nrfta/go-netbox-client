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

type WritableTenant struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Group int32 `json:"group,omitempty"`
	Description string `json:"description,omitempty"`
	Comments string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	CircuitCount int32 `json:"circuit_count,omitempty"`
	DeviceCount int32 `json:"device_count,omitempty"`
	IpaddressCount int32 `json:"ipaddress_count,omitempty"`
	PrefixCount int32 `json:"prefix_count,omitempty"`
	RackCount int32 `json:"rack_count,omitempty"`
	SiteCount int32 `json:"site_count,omitempty"`
	VirtualmachineCount int32 `json:"virtualmachine_count,omitempty"`
	VlanCount int32 `json:"vlan_count,omitempty"`
	VrfCount int32 `json:"vrf_count,omitempty"`
	ClusterCount int32 `json:"cluster_count,omitempty"`
}
