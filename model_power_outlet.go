/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.11
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

import (
	"time"
)

type PowerOutlet struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Device *NestedDevice `json:"device"`
	Name string `json:"name"`
	// Physical label
	Label string `json:"label,omitempty"`
	Type_ *Type5 `json:"type,omitempty"`
	PowerPort *NestedPowerPort `json:"power_port,omitempty"`
	FeedLeg *FeedLeg `json:"feed_leg,omitempty"`
	Description string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`
	Cable *NestedCable `json:"cable,omitempty"`
	//  Return the appropriate serializer for the cable termination model. 
	CablePeer map[string]string `json:"cable_peer,omitempty"`
	CablePeerType string `json:"cable_peer_type,omitempty"`
	//  Return the appropriate serializer for the type of connected object. 
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`
	ConnectedEndpointReachable bool `json:"connected_endpoint_reachable,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	Occupied string `json:"_occupied,omitempty"`
}
