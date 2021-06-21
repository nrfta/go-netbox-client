/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type WritableDeviceInterface struct {
	Id int32 `json:"id,omitempty"`
	Device int32 `json:"device"`
	Name string `json:"name"`
	Type_ string `json:"type"`
	Enabled bool `json:"enabled,omitempty"`
	Lag int32 `json:"lag,omitempty"`
	Mtu int32 `json:"mtu,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`
	Description string `json:"description,omitempty"`
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`
	//  Return the appropriate serializer for the type of connected object. 
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`
	ConnectionStatus bool `json:"connection_status,omitempty"`
	Cable *NestedCable `json:"cable,omitempty"`
	Mode string `json:"mode,omitempty"`
	UntaggedVlan int32 `json:"untagged_vlan,omitempty"`
	TaggedVlans []int32 `json:"tagged_vlans,omitempty"`
	Tags []string `json:"tags,omitempty"`
	CountIpaddresses int32 `json:"count_ipaddresses,omitempty"`
}