# IpAddress

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Display** | **string** |  | [optional] [default to null]
**Family** | [***Family**](Family.md) |  | [optional] [default to null]
**Address** | **string** | IPv4 or IPv6 address (with mask) | [default to null]
**Vrf** | [***NestedVrf**](NestedVRF.md) |  | [optional] [default to null]
**Tenant** | [***NestedTenant**](NestedTenant.md) |  | [optional] [default to null]
**Status** | [***Status7**](Status_7.md) |  | [optional] [default to null]
**Role** | [***Role1**](Role_1.md) |  | [optional] [default to null]
**AssignedObjectType** | **string** |  | [optional] [default to null]
**AssignedObjectId** | **int32** |  | [optional] [default to null]
**AssignedObject** | **map[string]string** |  | [optional] [default to null]
**NatInside** | [***NestedIpAddress**](NestedIPAddress.md) |  | [optional] [default to null]
**NatOutside** | [***NestedIpAddress**](NestedIPAddress.md) |  | [optional] [default to null]
**DnsName** | **string** | Hostname or FQDN (not case-sensitive) | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Tags** | [**[]NestedTag**](NestedTag.md) |  | [optional] [default to null]
**CustomFields** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Created** | **string** |  | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


