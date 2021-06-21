# PowerFeed

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Display** | **string** |  | [optional] [default to null]
**PowerPanel** | [***NestedPowerPanel**](NestedPowerPanel.md) |  | [default to null]
**Rack** | [***NestedRack**](NestedRack.md) |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Status** | [***Status3**](Status_3.md) |  | [optional] [default to null]
**Type_** | [***Type4**](Type_4.md) |  | [optional] [default to null]
**Supply** | [***Supply**](Supply.md) |  | [optional] [default to null]
**Phase** | [***Phase**](Phase.md) |  | [optional] [default to null]
**Voltage** | **int32** |  | [optional] [default to null]
**Amperage** | **int32** |  | [optional] [default to null]
**MaxUtilization** | **int32** | Maximum permissible draw (percentage) | [optional] [default to null]
**Comments** | **string** |  | [optional] [default to null]
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


