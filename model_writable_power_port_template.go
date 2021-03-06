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

type WritablePowerPortTemplate struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	DeviceType int32 `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label string `json:"label,omitempty"`
	Type_ string `json:"type,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw int32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw int32 `json:"allocated_draw,omitempty"`
	Description string `json:"description,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
