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

type Rack struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Name string `json:"name"`
	// Locally-assigned identifier
	FacilityId string `json:"facility_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Site *NestedSite `json:"site"`
	Location *NestedLocation `json:"location,omitempty"`
	Tenant *NestedTenant `json:"tenant,omitempty"`
	Status *Status4 `json:"status,omitempty"`
	Role *NestedRackRole `json:"role,omitempty"`
	Serial string `json:"serial,omitempty"`
	// A unique tag used to identify this rack
	AssetTag string `json:"asset_tag,omitempty"`
	Type_ *Type6 `json:"type,omitempty"`
	Width *Width `json:"width,omitempty"`
	// Height in rack units
	UHeight int32 `json:"u_height,omitempty"`
	// Units are numbered top-to-bottom
	DescUnits bool `json:"desc_units,omitempty"`
	// Outer dimension of rack (width)
	OuterWidth int32 `json:"outer_width,omitempty"`
	// Outer dimension of rack (depth)
	OuterDepth int32 `json:"outer_depth,omitempty"`
	OuterUnit *OuterUnit `json:"outer_unit,omitempty"`
	Comments string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	DeviceCount int32 `json:"device_count,omitempty"`
	PowerfeedCount int32 `json:"powerfeed_count,omitempty"`
}
