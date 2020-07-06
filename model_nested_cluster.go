/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type NestedCluster struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Name string `json:"name"`
	VirtualmachineCount int32 `json:"virtualmachine_count,omitempty"`
}
