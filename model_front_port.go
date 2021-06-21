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

type FrontPort struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	Device *NestedDevice `json:"device"`
	Name string `json:"name"`
	// Physical label
	Label string `json:"label,omitempty"`
	Type_ *Type1 `json:"type"`
	RearPort *FrontPortRearPort `json:"rear_port"`
	RearPortPosition int32 `json:"rear_port_position,omitempty"`
	Description string `json:"description,omitempty"`
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`
	Cable *NestedCable `json:"cable,omitempty"`
	//  Return the appropriate serializer for the cable termination model. 
	CablePeer map[string]string `json:"cable_peer,omitempty"`
	CablePeerType string `json:"cable_peer_type,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
	Occupied string `json:"_occupied,omitempty"`
}
