/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type DeviceRole struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Color string `json:"color,omitempty"`
	// Virtual machines may be assigned to this role
	VmRole bool `json:"vm_role,omitempty"`
	Description string `json:"description,omitempty"`
	DeviceCount int32 `json:"device_count,omitempty"`
	VirtualmachineCount int32 `json:"virtualmachine_count,omitempty"`
}