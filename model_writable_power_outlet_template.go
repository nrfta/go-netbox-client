/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type WritablePowerOutletTemplate struct {
	Id int32 `json:"id,omitempty"`
	DeviceType int32 `json:"device_type"`
	Name string `json:"name"`
	Type_ string `json:"type,omitempty"`
	PowerPort int32 `json:"power_port,omitempty"`
	// Phase (for three-phase feeds)
	FeedLeg string `json:"feed_leg,omitempty"`
}
