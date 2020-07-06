/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type NestedVlanGroup struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	VlanCount int32 `json:"vlan_count,omitempty"`
}
