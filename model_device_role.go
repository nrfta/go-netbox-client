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

type DeviceRole struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Color string `json:"color,omitempty"`
	// Virtual machines may be assigned to this role
	VmRole bool `json:"vm_role,omitempty"`
	Description string `json:"description,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	DeviceCount int32 `json:"device_count,omitempty"`
	VirtualmachineCount int32 `json:"virtualmachine_count,omitempty"`
}
