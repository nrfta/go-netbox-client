/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type Platform struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Manufacturer *NestedManufacturer `json:"manufacturer,omitempty"`
	// The name of the NAPALM driver to use when interacting with devices
	NapalmDriver string `json:"napalm_driver,omitempty"`
	// Additional arguments to pass when initiating the NAPALM driver (JSON format)
	NapalmArgs string `json:"napalm_args,omitempty"`
	Description string `json:"description,omitempty"`
	DeviceCount int32 `json:"device_count,omitempty"`
	VirtualmachineCount int32 `json:"virtualmachine_count,omitempty"`
}