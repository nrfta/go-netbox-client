/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type WritableInventoryItem struct {
	Id int32 `json:"id,omitempty"`
	Device int32 `json:"device"`
	Parent int32 `json:"parent,omitempty"`
	Name string `json:"name"`
	Manufacturer int32 `json:"manufacturer,omitempty"`
	// Manufacturer-assigned part identifier
	PartId string `json:"part_id,omitempty"`
	Serial string `json:"serial,omitempty"`
	// A unique tag used to identify this item
	AssetTag string `json:"asset_tag,omitempty"`
	// This item was automatically discovered
	Discovered bool `json:"discovered,omitempty"`
	Description string `json:"description,omitempty"`
	Tags []string `json:"tags,omitempty"`
}
