/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type PowerPanel struct {
	Id int32 `json:"id,omitempty"`
	Site *NestedSite `json:"site"`
	RackGroup *NestedRackGroup `json:"rack_group,omitempty"`
	Name string `json:"name"`
	PowerfeedCount int32 `json:"powerfeed_count,omitempty"`
}