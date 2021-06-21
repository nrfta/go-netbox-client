/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.11
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type RackUnit struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Face *Face `json:"face,omitempty"`
	Device *NestedDevice `json:"device,omitempty"`
	Occupied bool `json:"occupied,omitempty"`
}