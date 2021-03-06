/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.11
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type WritablePowerPanel struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Site int32 `json:"site"`
	Location int32 `json:"location,omitempty"`
	Name string `json:"name"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	PowerfeedCount int32 `json:"powerfeed_count,omitempty"`
}
