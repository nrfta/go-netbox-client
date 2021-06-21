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

type WritablePowerOutletTemplate struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	DeviceType int32 `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label string `json:"label,omitempty"`
	Type_ string `json:"type,omitempty"`
	PowerPort int32 `json:"power_port,omitempty"`
	// Phase (for three-phase feeds)
	FeedLeg string `json:"feed_leg,omitempty"`
	Description string `json:"description,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}