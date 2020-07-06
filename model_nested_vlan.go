/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type NestedVlan struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Vid int32 `json:"vid"`
	Name string `json:"name"`
	DisplayName string `json:"display_name,omitempty"`
}
