# CustomField

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Display** | **string** |  | [optional] [default to null]
**ContentTypes** | **[]string** |  | [default to null]
**Type_** | [***Type7**](Type_7.md) |  | [default to null]
**Name** | **string** | Internal field name | [default to null]
**Label** | **string** | Name of the field as displayed to users (if not provided, the field&#39;s name will be used) | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Required** | **bool** | If true, this field is required when creating new objects or editing an existing object. | [optional] [default to null]
**FilterLogic** | [***FilterLogic**](Filter logic.md) |  | [optional] [default to null]
**Default_** | **string** | Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] [default to null]
**Weight** | **int32** | Fields with higher weights appear lower in a form. | [optional] [default to null]
**ValidationMinimum** | **int32** | Minimum allowed value (for numeric fields) | [optional] [default to null]
**ValidationMaximum** | **int32** | Maximum allowed value (for numeric fields) | [optional] [default to null]
**ValidationRegex** | **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. | [optional] [default to null]
**Choices** | **[]string** | Comma-separated list of available choices (for selection fields) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


