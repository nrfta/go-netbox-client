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

type VirtualMachineWithConfigContext struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Status *Status9 `json:"status,omitempty"`
	Site *NestedSite `json:"site,omitempty"`
	Cluster *NestedCluster `json:"cluster"`
	Role *NestedDeviceRole `json:"role,omitempty"`
	Tenant *NestedTenant `json:"tenant,omitempty"`
	Platform *NestedPlatform `json:"platform,omitempty"`
	PrimaryIp *NestedIpAddress `json:"primary_ip,omitempty"`
	PrimaryIp4 *NestedIpAddress `json:"primary_ip4,omitempty"`
	PrimaryIp6 *NestedIpAddress `json:"primary_ip6,omitempty"`
	Vcpus int32 `json:"vcpus,omitempty"`
	Memory int32 `json:"memory,omitempty"`
	Disk int32 `json:"disk,omitempty"`
	Comments string `json:"comments,omitempty"`
	LocalContextData string `json:"local_context_data,omitempty"`
	Tags []string `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	ConfigContext map[string]string `json:"config_context,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
