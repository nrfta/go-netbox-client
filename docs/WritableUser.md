# WritableUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Url** | **string** |  | [optional] [default to null]
**Display** | **string** |  | [optional] [default to null]
**Username** | **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | [default to null]
**Password** | **string** |  | [default to null]
**FirstName** | **string** |  | [optional] [default to null]
**LastName** | **string** |  | [optional] [default to null]
**Email** | **string** |  | [optional] [default to null]
**IsStaff** | **bool** | Designates whether the user can log into this admin site. | [optional] [default to null]
**IsActive** | **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] [default to null]
**DateJoined** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Groups** | **[]int32** | The groups this user belongs to. A user will get all permissions granted to each of their groups. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


