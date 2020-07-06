/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type InlineResponse20048 struct {
	Count int32 `json:"count"`
	Next string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	Results []Service `json:"results"`
}
