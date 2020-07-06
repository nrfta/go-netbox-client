/*
 * NetBox API
 *
 * API to access NetBox
 *
 * API version: 2.8
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package netbox_client

type InterfaceConnection struct {
	InterfaceA *NestedInterface `json:"interface_a,omitempty"`
	InterfaceB *NestedInterface `json:"interface_b"`
	ConnectionStatus *ConnectionStatus `json:"connection_status,omitempty"`
}
