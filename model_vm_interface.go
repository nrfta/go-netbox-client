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

type VmInterface struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	VirtualMachine *NestedVirtualMachine `json:"virtual_machine"`
	Name string `json:"name"`
	Enabled bool `json:"enabled,omitempty"`
	Parent *NestedVmInterface `json:"parent,omitempty"`
	Mtu int32 `json:"mtu,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	Description string `json:"description,omitempty"`
	Mode *Mode `json:"mode,omitempty"`
	UntaggedVlan *NestedVlan `json:"untagged_vlan,omitempty"`
	TaggedVlans []NestedVlan `json:"tagged_vlans,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	CountIpaddresses int32 `json:"count_ipaddresses,omitempty"`
}
