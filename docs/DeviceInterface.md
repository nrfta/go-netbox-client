# DeviceInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Device** | [***NestedDevice**](NestedDevice.md) |  | [default to null]
**Name** | **string** |  | [default to null]
**Type_** | [***Type2**](Type_2.md) |  | [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Lag** | [***NestedInterface**](NestedInterface.md) |  | [optional] [default to null]
**Mtu** | **int32** |  | [optional] [default to null]
**MacAddress** | **string** |  | [optional] [default to null]
**MgmtOnly** | **bool** | This interface is used only for out-of-band management | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**ConnectedEndpointType** | **string** |  | [optional] [default to null]
**ConnectedEndpoint** | **map[string]string** |  Return the appropriate serializer for the type of connected object.  | [optional] [default to null]
**ConnectionStatus** | [***ConnectionStatus**](Connection status.md) |  | [optional] [default to null]
**Cable** | [***NestedCable**](NestedCable.md) |  | [optional] [default to null]
**Mode** | [***Mode**](Mode.md) |  | [optional] [default to null]
**UntaggedVlan** | [***NestedVlan**](NestedVLAN.md) |  | [optional] [default to null]
**TaggedVlans** | [**[]NestedVlan**](NestedVLAN.md) |  | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]
**CountIpaddresses** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


