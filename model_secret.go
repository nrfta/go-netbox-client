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

type Secret struct {
	Id int32 `json:"id,omitempty"`
	Url string `json:"url,omitempty"`
	Display string `json:"display,omitempty"`
	AssignedObjectType string `json:"assigned_object_type"`
	AssignedObjectId int32 `json:"assigned_object_id"`
	AssignedObject map[string]string `json:"assigned_object,omitempty"`
	Role *NestedSecretRole `json:"role"`
	Name string `json:"name,omitempty"`
	Plaintext string `json:"plaintext"`
	Hash string `json:"hash,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields *interface{} `json:"custom_fields,omitempty"`
	Created string `json:"created,omitempty"`
	LastUpdated time.Time `json:"last_updated,omitempty"`
}
