/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type RackGroup struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Site *NestedSite `json:"site"`
	Parent *NestedRackGroup `json:"parent,omitempty"`
	Description string `json:"description,omitempty"`
	RackCount int32 `json:"rack_count,omitempty"`
}
