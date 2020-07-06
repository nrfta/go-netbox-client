# WritablePowerPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Device** | **int32** |  | [default to null]
**Name** | **string** |  | [default to null]
**Type_** | **string** | Physical port type | [optional] [default to null]
**MaximumDraw** | **int32** | Maximum power draw (watts) | [optional] [default to null]
**AllocatedDraw** | **int32** | Allocated power draw (watts) | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**ConnectedEndpointType** | **string** |  | [optional] [default to null]
**ConnectedEndpoint** | **map[string]string** |  Return the appropriate serializer for the type of connected object.  | [optional] [default to null]
**ConnectionStatus** | **bool** |  | [optional] [default to null]
**Cable** | [***NestedCable**](NestedCable.md) |  | [optional] [default to null]
**Tags** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


