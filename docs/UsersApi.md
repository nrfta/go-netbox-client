# \UsersApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersConfigList**](UsersApi.md#UsersConfigList) | **Get** /users/config/ | 
[**UsersGroupsBulkDelete**](UsersApi.md#UsersGroupsBulkDelete) | **Delete** /users/groups/ | 
[**UsersGroupsBulkPartialUpdate**](UsersApi.md#UsersGroupsBulkPartialUpdate) | **Patch** /users/groups/ | 
[**UsersGroupsBulkUpdate**](UsersApi.md#UsersGroupsBulkUpdate) | **Put** /users/groups/ | 
[**UsersGroupsCreate**](UsersApi.md#UsersGroupsCreate) | **Post** /users/groups/ | 
[**UsersGroupsDelete**](UsersApi.md#UsersGroupsDelete) | **Delete** /users/groups/{id}/ | 
[**UsersGroupsList**](UsersApi.md#UsersGroupsList) | **Get** /users/groups/ | 
[**UsersGroupsPartialUpdate**](UsersApi.md#UsersGroupsPartialUpdate) | **Patch** /users/groups/{id}/ | 
[**UsersGroupsRead**](UsersApi.md#UsersGroupsRead) | **Get** /users/groups/{id}/ | 
[**UsersGroupsUpdate**](UsersApi.md#UsersGroupsUpdate) | **Put** /users/groups/{id}/ | 
[**UsersPermissionsBulkDelete**](UsersApi.md#UsersPermissionsBulkDelete) | **Delete** /users/permissions/ | 
[**UsersPermissionsBulkPartialUpdate**](UsersApi.md#UsersPermissionsBulkPartialUpdate) | **Patch** /users/permissions/ | 
[**UsersPermissionsBulkUpdate**](UsersApi.md#UsersPermissionsBulkUpdate) | **Put** /users/permissions/ | 
[**UsersPermissionsCreate**](UsersApi.md#UsersPermissionsCreate) | **Post** /users/permissions/ | 
[**UsersPermissionsDelete**](UsersApi.md#UsersPermissionsDelete) | **Delete** /users/permissions/{id}/ | 
[**UsersPermissionsList**](UsersApi.md#UsersPermissionsList) | **Get** /users/permissions/ | 
[**UsersPermissionsPartialUpdate**](UsersApi.md#UsersPermissionsPartialUpdate) | **Patch** /users/permissions/{id}/ | 
[**UsersPermissionsRead**](UsersApi.md#UsersPermissionsRead) | **Get** /users/permissions/{id}/ | 
[**UsersPermissionsUpdate**](UsersApi.md#UsersPermissionsUpdate) | **Put** /users/permissions/{id}/ | 
[**UsersUsersBulkDelete**](UsersApi.md#UsersUsersBulkDelete) | **Delete** /users/users/ | 
[**UsersUsersBulkPartialUpdate**](UsersApi.md#UsersUsersBulkPartialUpdate) | **Patch** /users/users/ | 
[**UsersUsersBulkUpdate**](UsersApi.md#UsersUsersBulkUpdate) | **Put** /users/users/ | 
[**UsersUsersCreate**](UsersApi.md#UsersUsersCreate) | **Post** /users/users/ | 
[**UsersUsersDelete**](UsersApi.md#UsersUsersDelete) | **Delete** /users/users/{id}/ | 
[**UsersUsersList**](UsersApi.md#UsersUsersList) | **Get** /users/users/ | 
[**UsersUsersPartialUpdate**](UsersApi.md#UsersUsersPartialUpdate) | **Patch** /users/users/{id}/ | 
[**UsersUsersRead**](UsersApi.md#UsersUsersRead) | **Get** /users/users/{id}/ | 
[**UsersUsersUpdate**](UsersApi.md#UsersUsersUpdate) | **Put** /users/users/{id}/ | 


# **UsersConfigList**
> UsersConfigList(ctx, )


Return the UserConfig for the currently authenticated User.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsBulkDelete**
> UsersGroupsBulkDelete(ctx, )




### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsBulkPartialUpdate**
> Group UsersGroupsBulkPartialUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Group**](Group.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsBulkUpdate**
> Group UsersGroupsBulkUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Group**](Group.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsCreate**
> Group UsersGroupsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Group**](Group.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsDelete**
> UsersGroupsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this group. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsList**
> InlineResponse20064 UsersGroupsList(ctx, optional)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersGroupsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **idN** | **optional.String**|  | 
 **idLte** | **optional.String**|  | 
 **idLt** | **optional.String**|  | 
 **idGte** | **optional.String**|  | 
 **idGt** | **optional.String**|  | 
 **nameN** | **optional.String**|  | 
 **nameIc** | **optional.String**|  | 
 **nameNic** | **optional.String**|  | 
 **nameIew** | **optional.String**|  | 
 **nameNiew** | **optional.String**|  | 
 **nameIsw** | **optional.String**|  | 
 **nameNisw** | **optional.String**|  | 
 **nameIe** | **optional.String**|  | 
 **nameNie** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20064**](inline_response_200_64.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsPartialUpdate**
> Group UsersGroupsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this group. | 
  **data** | [**Group**](Group.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsRead**
> Group UsersGroupsRead(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this group. | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGroupsUpdate**
> Group UsersGroupsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this group. | 
  **data** | [**Group**](Group.md)|  | 

### Return type

[**Group**](Group.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsBulkDelete**
> UsersPermissionsBulkDelete(ctx, )




### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsBulkPartialUpdate**
> ObjectPermission UsersPermissionsBulkPartialUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableObjectPermission**](WritableObjectPermission.md)|  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsBulkUpdate**
> ObjectPermission UsersPermissionsBulkUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableObjectPermission**](WritableObjectPermission.md)|  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsCreate**
> ObjectPermission UsersPermissionsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableObjectPermission**](WritableObjectPermission.md)|  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsDelete**
> UsersPermissionsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this permission. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsList**
> InlineResponse20065 UsersPermissionsList(ctx, optional)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersPermissionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersPermissionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **enabled** | **optional.String**|  | 
 **objectTypes** | **optional.String**|  | 
 **userId** | **optional.String**|  | 
 **user** | **optional.String**|  | 
 **groupId** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **idN** | **optional.String**|  | 
 **idLte** | **optional.String**|  | 
 **idLt** | **optional.String**|  | 
 **idGte** | **optional.String**|  | 
 **idGt** | **optional.String**|  | 
 **nameN** | **optional.String**|  | 
 **nameIc** | **optional.String**|  | 
 **nameNic** | **optional.String**|  | 
 **nameIew** | **optional.String**|  | 
 **nameNiew** | **optional.String**|  | 
 **nameIsw** | **optional.String**|  | 
 **nameNisw** | **optional.String**|  | 
 **nameIe** | **optional.String**|  | 
 **nameNie** | **optional.String**|  | 
 **objectTypesN** | **optional.String**|  | 
 **userIdN** | **optional.String**|  | 
 **userN** | **optional.String**|  | 
 **groupIdN** | **optional.String**|  | 
 **groupN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20065**](inline_response_200_65.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsPartialUpdate**
> ObjectPermission UsersPermissionsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this permission. | 
  **data** | [**WritableObjectPermission**](WritableObjectPermission.md)|  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsRead**
> ObjectPermission UsersPermissionsRead(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this permission. | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPermissionsUpdate**
> ObjectPermission UsersPermissionsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this permission. | 
  **data** | [**WritableObjectPermission**](WritableObjectPermission.md)|  | 

### Return type

[**ObjectPermission**](ObjectPermission.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersBulkDelete**
> UsersUsersBulkDelete(ctx, )




### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersBulkPartialUpdate**
> User UsersUsersBulkPartialUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableUser**](WritableUser.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersBulkUpdate**
> User UsersUsersBulkUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableUser**](WritableUser.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersCreate**
> User UsersUsersCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableUser**](WritableUser.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersDelete**
> UsersUsersDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this user. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersList**
> InlineResponse20066 UsersUsersList(ctx, optional)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUsersUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUsersUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **firstName** | **optional.String**|  | 
 **lastName** | **optional.String**|  | 
 **email** | **optional.String**|  | 
 **isStaff** | **optional.String**|  | 
 **isActive** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **groupId** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **idN** | **optional.String**|  | 
 **idLte** | **optional.String**|  | 
 **idLt** | **optional.String**|  | 
 **idGte** | **optional.String**|  | 
 **idGt** | **optional.String**|  | 
 **usernameN** | **optional.String**|  | 
 **usernameIc** | **optional.String**|  | 
 **usernameNic** | **optional.String**|  | 
 **usernameIew** | **optional.String**|  | 
 **usernameNiew** | **optional.String**|  | 
 **usernameIsw** | **optional.String**|  | 
 **usernameNisw** | **optional.String**|  | 
 **usernameIe** | **optional.String**|  | 
 **usernameNie** | **optional.String**|  | 
 **firstNameN** | **optional.String**|  | 
 **firstNameIc** | **optional.String**|  | 
 **firstNameNic** | **optional.String**|  | 
 **firstNameIew** | **optional.String**|  | 
 **firstNameNiew** | **optional.String**|  | 
 **firstNameIsw** | **optional.String**|  | 
 **firstNameNisw** | **optional.String**|  | 
 **firstNameIe** | **optional.String**|  | 
 **firstNameNie** | **optional.String**|  | 
 **lastNameN** | **optional.String**|  | 
 **lastNameIc** | **optional.String**|  | 
 **lastNameNic** | **optional.String**|  | 
 **lastNameIew** | **optional.String**|  | 
 **lastNameNiew** | **optional.String**|  | 
 **lastNameIsw** | **optional.String**|  | 
 **lastNameNisw** | **optional.String**|  | 
 **lastNameIe** | **optional.String**|  | 
 **lastNameNie** | **optional.String**|  | 
 **emailN** | **optional.String**|  | 
 **emailIc** | **optional.String**|  | 
 **emailNic** | **optional.String**|  | 
 **emailIew** | **optional.String**|  | 
 **emailNiew** | **optional.String**|  | 
 **emailIsw** | **optional.String**|  | 
 **emailNisw** | **optional.String**|  | 
 **emailIe** | **optional.String**|  | 
 **emailNie** | **optional.String**|  | 
 **groupIdN** | **optional.String**|  | 
 **groupN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20066**](inline_response_200_66.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersPartialUpdate**
> User UsersUsersPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this user. | 
  **data** | [**WritableUser**](WritableUser.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersRead**
> User UsersUsersRead(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this user. | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUsersUpdate**
> User UsersUsersUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this user. | 
  **data** | [**WritableUser**](WritableUser.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

