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

type DeviceType struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Manufacturer *NestedManufacturer `json:"manufacturer"`
	Model string `json:"model"`
	Slug string `json:"slug"`
	DisplayName string `json:"display_name,omitempty"`
	// Discrete part number (optional)
	PartNumber string `json:"part_number,omitempty"`
	UHeight int32 `json:"u_height,omitempty"`
	// Device consumes both front and rear rack faces
	IsFullDepth bool `json:"is_full_depth,omitempty"`
	SubdeviceRole *SubdeviceRole `json:"subdevice_role,omitempty"`
	FrontImage string `json:"front_image,omitempty"`
	RearImage string `json:"rear_image,omitempty"`
	Comments string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	DeviceCount int32 `json:"device_count,omitempty"`
}
