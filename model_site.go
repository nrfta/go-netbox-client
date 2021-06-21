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

type Site struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Status *Status5 `json:"status,omitempty"`
	Region *NestedRegion `json:"region,omitempty"`
	Group *NestedSiteGroup `json:"group,omitempty"`
	Tenant *NestedTenant `json:"tenant,omitempty"`
	// Local facility ID or description
	Facility string `json:"facility,omitempty"`
	// 32-bit autonomous system number
	Asn int32 `json:"asn,omitempty"`
	TimeZone string `json:"time_zone,omitempty"`
	Description string `json:"description,omitempty"`
	PhysicalAddress string `json:"physical_address,omitempty"`
	ShippingAddress string `json:"shipping_address,omitempty"`
	// GPS coordinate (latitude)
	Latitude string `json:"latitude,omitempty"`
	// GPS coordinate (longitude)
	Longitude string `json:"longitude,omitempty"`
	ContactName string `json:"contact_name,omitempty"`
	ContactPhone string `json:"contact_phone,omitempty"`
	ContactEmail string `json:"contact_email,omitempty"`
	Comments string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	CircuitCount int32 `json:"circuit_count,omitempty"`
	DeviceCount int32 `json:"device_count,omitempty"`
	PrefixCount int32 `json:"prefix_count,omitempty"`
	RackCount int32 `json:"rack_count,omitempty"`
	VirtualmachineCount int32 `json:"virtualmachine_count,omitempty"`
	VlanCount int32 `json:"vlan_count,omitempty"`
}