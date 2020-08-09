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

type WritableVirtualMachineWithConfigContext struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Status string `json:"status,omitempty"`
	Site string `json:"site,omitempty"`
	Cluster int32 `json:"cluster"`
	Role int32 `json:"role,omitempty"`
	Tenant int32 `json:"tenant,omitempty"`
	Platform int32 `json:"platform,omitempty"`
	PrimaryIp string `json:"primary_ip,omitempty"`
	PrimaryIp4 int32 `json:"primary_ip4,omitempty"`
	PrimaryIp6 int32 `json:"primary_ip6,omitempty"`
	Vcpus int32 `json:"vcpus,omitempty"`
	Memory int32 `json:"memory,omitempty"`
	Disk int32 `json:"disk,omitempty"`
	Comments string `json:"comments,omitempty"`
	LocalContextData string `json:"local_context_data,omitempty"`
	Tags []string `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	ConfigContext *interface{} `json:"config_context,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}