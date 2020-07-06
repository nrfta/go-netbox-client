/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type RearPortTemplate struct {
	Id int32 `json:"id,omitempty"`
	DeviceType *NestedDeviceType `json:"device_type"`
	Name string `json:"name"`
	Type_ *Type1 `json:"type"`
	Positions int32 `json:"positions,omitempty"`
}
