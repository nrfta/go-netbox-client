# WritableConsolePort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Display** | **string** |  | [optional] [default to null]
**Device** | **int32** |  | [default to null]
**Name** | **string** |  | [default to null]
**Label** | **string** | Physical label | [optional] [default to null]
**Type_** | **string** | Physical port type | [optional] [default to null]
**Speed** | **int32** | Port speed in bits per second | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**MarkConnected** | **bool** | Treat as if a cable is connected | [optional] [default to null]
**Cable** | [***NestedCable**](NestedCable.md) |  | [optional] [default to null]
**CablePeer** | **map[string]string** |  Return the appropriate serializer for the cable termination model.  | [optional] [default to null]
**CablePeerType** | **string** |  | [optional] [default to null]
**ConnectedEndpoint** | **map[string]string** |  Return the appropriate serializer for the type of connected object.  | [optional] [default to null]
**ConnectedEndpointType** | **string** |  | [optional] [default to null]
**ConnectedEndpointReachable** | **bool** |  | [optional] [default to null]
**Tags** | [**[]NestedTag**](NestedTag.md) |  | [optional] [default to null]
**CustomFields** | [***interface{}**](interface{}.md) |  | [optional] [default to null]
**Created** | **string** |  | [optional] [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Occupied** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


