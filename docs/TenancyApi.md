# \TenancyApi

All URIs are relative to *http://localhost:8000/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TenancyTenantGroupsBulkDelete**](TenancyApi.md#TenancyTenantGroupsBulkDelete) | **Delete** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkPartialUpdate**](TenancyApi.md#TenancyTenantGroupsBulkPartialUpdate) | **Patch** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsBulkUpdate**](TenancyApi.md#TenancyTenantGroupsBulkUpdate) | **Put** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsCreate**](TenancyApi.md#TenancyTenantGroupsCreate) | **Post** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsDelete**](TenancyApi.md#TenancyTenantGroupsDelete) | **Delete** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsList**](TenancyApi.md#TenancyTenantGroupsList) | **Get** /tenancy/tenant-groups/ | 
[**TenancyTenantGroupsPartialUpdate**](TenancyApi.md#TenancyTenantGroupsPartialUpdate) | **Patch** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsRead**](TenancyApi.md#TenancyTenantGroupsRead) | **Get** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantGroupsUpdate**](TenancyApi.md#TenancyTenantGroupsUpdate) | **Put** /tenancy/tenant-groups/{id}/ | 
[**TenancyTenantsBulkDelete**](TenancyApi.md#TenancyTenantsBulkDelete) | **Delete** /tenancy/tenants/ | 
[**TenancyTenantsBulkPartialUpdate**](TenancyApi.md#TenancyTenantsBulkPartialUpdate) | **Patch** /tenancy/tenants/ | 
[**TenancyTenantsBulkUpdate**](TenancyApi.md#TenancyTenantsBulkUpdate) | **Put** /tenancy/tenants/ | 
[**TenancyTenantsCreate**](TenancyApi.md#TenancyTenantsCreate) | **Post** /tenancy/tenants/ | 
[**TenancyTenantsDelete**](TenancyApi.md#TenancyTenantsDelete) | **Delete** /tenancy/tenants/{id}/ | 
[**TenancyTenantsList**](TenancyApi.md#TenancyTenantsList) | **Get** /tenancy/tenants/ | 
[**TenancyTenantsPartialUpdate**](TenancyApi.md#TenancyTenantsPartialUpdate) | **Patch** /tenancy/tenants/{id}/ | 
[**TenancyTenantsRead**](TenancyApi.md#TenancyTenantsRead) | **Get** /tenancy/tenants/{id}/ | 
[**TenancyTenantsUpdate**](TenancyApi.md#TenancyTenantsUpdate) | **Put** /tenancy/tenants/{id}/ | 


# **TenancyTenantGroupsBulkDelete**
> TenancyTenantGroupsBulkDelete(ctx, )




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

# **TenancyTenantGroupsBulkPartialUpdate**
> TenantGroup TenancyTenantGroupsBulkPartialUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableTenantGroup**](WritableTenantGroup.md)|  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsBulkUpdate**
> TenantGroup TenancyTenantGroupsBulkUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableTenantGroup**](WritableTenantGroup.md)|  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsCreate**
> TenantGroup TenancyTenantGroupsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableTenantGroup**](WritableTenantGroup.md)|  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsDelete**
> TenancyTenantGroupsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tenant group. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsList**
> InlineResponse20062 TenancyTenantGroupsList(ctx, optional)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenancyApiTenancyTenantGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantGroupsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **slug** | **optional.String**|  | 
 **description** | **optional.String**|  | 
 **created** | **optional.String**|  | 
 **createdGte** | **optional.String**|  | 
 **createdLte** | **optional.String**|  | 
 **lastUpdated** | **optional.String**|  | 
 **lastUpdatedGte** | **optional.String**|  | 
 **lastUpdatedLte** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **parentId** | **optional.String**|  | 
 **parent** | **optional.String**|  | 
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
 **slugN** | **optional.String**|  | 
 **slugIc** | **optional.String**|  | 
 **slugNic** | **optional.String**|  | 
 **slugIew** | **optional.String**|  | 
 **slugNiew** | **optional.String**|  | 
 **slugIsw** | **optional.String**|  | 
 **slugNisw** | **optional.String**|  | 
 **slugIe** | **optional.String**|  | 
 **slugNie** | **optional.String**|  | 
 **descriptionN** | **optional.String**|  | 
 **descriptionIc** | **optional.String**|  | 
 **descriptionNic** | **optional.String**|  | 
 **descriptionIew** | **optional.String**|  | 
 **descriptionNiew** | **optional.String**|  | 
 **descriptionIsw** | **optional.String**|  | 
 **descriptionNisw** | **optional.String**|  | 
 **descriptionIe** | **optional.String**|  | 
 **descriptionNie** | **optional.String**|  | 
 **parentIdN** | **optional.String**|  | 
 **parentN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20062**](inline_response_200_62.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsPartialUpdate**
> TenantGroup TenancyTenantGroupsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tenant group. | 
  **data** | [**WritableTenantGroup**](WritableTenantGroup.md)|  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsRead**
> TenantGroup TenancyTenantGroupsRead(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tenant group. | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantGroupsUpdate**
> TenantGroup TenancyTenantGroupsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tenant group. | 
  **data** | [**WritableTenantGroup**](WritableTenantGroup.md)|  | 

### Return type

[**TenantGroup**](TenantGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsBulkDelete**
> TenancyTenantsBulkDelete(ctx, )




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

# **TenancyTenantsBulkPartialUpdate**
> Tenant TenancyTenantsBulkPartialUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableTenant**](WritableTenant.md)|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsBulkUpdate**
> Tenant TenancyTenantsBulkUpdate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableTenant**](WritableTenant.md)|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsCreate**
> Tenant TenancyTenantsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableTenant**](WritableTenant.md)|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsDelete**
> TenancyTenantsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tenant. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsList**
> InlineResponse20063 TenancyTenantsList(ctx, optional)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TenancyApiTenancyTenantsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TenancyApiTenancyTenantsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **slug** | **optional.String**|  | 
 **created** | **optional.String**|  | 
 **createdGte** | **optional.String**|  | 
 **createdLte** | **optional.String**|  | 
 **lastUpdated** | **optional.String**|  | 
 **lastUpdatedGte** | **optional.String**|  | 
 **lastUpdatedLte** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **groupId** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **tag** | **optional.String**|  | 
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
 **slugN** | **optional.String**|  | 
 **slugIc** | **optional.String**|  | 
 **slugNic** | **optional.String**|  | 
 **slugIew** | **optional.String**|  | 
 **slugNiew** | **optional.String**|  | 
 **slugIsw** | **optional.String**|  | 
 **slugNisw** | **optional.String**|  | 
 **slugIe** | **optional.String**|  | 
 **slugNie** | **optional.String**|  | 
 **groupIdN** | **optional.String**|  | 
 **groupN** | **optional.String**|  | 
 **tagN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20063**](inline_response_200_63.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsPartialUpdate**
> Tenant TenancyTenantsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tenant. | 
  **data** | [**WritableTenant**](WritableTenant.md)|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsRead**
> Tenant TenancyTenantsRead(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tenant. | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TenancyTenantsUpdate**
> Tenant TenancyTenantsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tenant. | 
  **data** | [**WritableTenant**](WritableTenant.md)|  | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

