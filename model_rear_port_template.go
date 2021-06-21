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

type RearPortTemplate struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	DeviceType *NestedDeviceType `json:"device_type"`
	Name string `json:"name"`
	// Physical label
	Label string `json:"label,omitempty"`
	Type_ *Type1 `json:"type"`
	Positions int32 `json:"positions,omitempty"`
	Description string `json:"description,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
