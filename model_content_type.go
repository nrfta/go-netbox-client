/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.11
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type ContentType struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	AppLabel string `json:"app_label"`
	Model string `json:"model"`
	DisplayName string `json:"display_name,omitempty"`
}
