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

type WritableLocation struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Site int32 `json:"site"`
	Parent int32 `json:"parent,omitempty"`
	Description string `json:"description,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	RackCount int32 `json:"rack_count,omitempty"`
	DeviceCount int32 `json:"device_count,omitempty"`
	Depth int32 `json:"_depth,omitempty"`
}
