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

type DeviceWithConfigContext struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Name string `json:"name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	DeviceType *NestedDeviceType `json:"device_type"`
	DeviceRole *NestedDeviceRole `json:"device_role"`
	Tenant *NestedTenant `json:"tenant,omitempty"`
	Platform *NestedPlatform `json:"platform,omitempty"`
	Serial string `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag string `json:"asset_tag,omitempty"`
	Site *NestedSite `json:"site"`
	Location *NestedLocation `json:"location,omitempty"`
	Rack *NestedRack `json:"rack,omitempty"`
	// The lowest-numbered unit occupied by the device
	Position int32 `json:"position,omitempty"`
	Face *Face `json:"face,omitempty"`
	ParentDevice *NestedDevice `json:"parent_device,omitempty"`
	Status *Status2 `json:"status,omitempty"`
	PrimaryIp *NestedIpAddress `json:"primary_ip,omitempty"`
	PrimaryIp4 *NestedIpAddress `json:"primary_ip4,omitempty"`
	PrimaryIp6 *NestedIpAddress `json:"primary_ip6,omitempty"`
	Cluster *NestedCluster `json:"cluster,omitempty"`
	VirtualChassis *NestedVirtualChassis `json:"virtual_chassis,omitempty"`
	VcPosition int32 `json:"vc_position,omitempty"`
	VcPriority int32 `json:"vc_priority,omitempty"`
	Comments string `json:"comments,omitempty"`
	LocalContextData string `json:"local_context_data,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	ConfigContext map[string]interface{} `json:"config_context,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
