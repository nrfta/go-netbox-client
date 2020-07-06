/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type Rir struct {
	Id int32 `json:"id,omitempty"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	// IP space managed by this RIR is considered private
	IsPrivate bool `json:"is_private,omitempty"`
	Description string `json:"description,omitempty"`
	AggregateCount int32 `json:"aggregate_count,omitempty"`
}
