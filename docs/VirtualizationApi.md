# \VirtualizationApi

All URIs are relative to *https://netbox/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VirtualizationClusterGroupsCreate**](VirtualizationApi.md#VirtualizationClusterGroupsCreate) | **Post** /virtualization/cluster-groups/ | 
[**VirtualizationClusterGroupsDelete**](VirtualizationApi.md#VirtualizationClusterGroupsDelete) | **Delete** /virtualization/cluster-groups/{id}/ | 
[**VirtualizationClusterGroupsList**](VirtualizationApi.md#VirtualizationClusterGroupsList) | **Get** /virtualization/cluster-groups/ | 
[**VirtualizationClusterGroupsPartialUpdate**](VirtualizationApi.md#VirtualizationClusterGroupsPartialUpdate) | **Patch** /virtualization/cluster-groups/{id}/ | 
[**VirtualizationClusterGroupsRead**](VirtualizationApi.md#VirtualizationClusterGroupsRead) | **Get** /virtualization/cluster-groups/{id}/ | 
[**VirtualizationClusterGroupsUpdate**](VirtualizationApi.md#VirtualizationClusterGroupsUpdate) | **Put** /virtualization/cluster-groups/{id}/ | 
[**VirtualizationClusterTypesCreate**](VirtualizationApi.md#VirtualizationClusterTypesCreate) | **Post** /virtualization/cluster-types/ | 
[**VirtualizationClusterTypesDelete**](VirtualizationApi.md#VirtualizationClusterTypesDelete) | **Delete** /virtualization/cluster-types/{id}/ | 
[**VirtualizationClusterTypesList**](VirtualizationApi.md#VirtualizationClusterTypesList) | **Get** /virtualization/cluster-types/ | 
[**VirtualizationClusterTypesPartialUpdate**](VirtualizationApi.md#VirtualizationClusterTypesPartialUpdate) | **Patch** /virtualization/cluster-types/{id}/ | 
[**VirtualizationClusterTypesRead**](VirtualizationApi.md#VirtualizationClusterTypesRead) | **Get** /virtualization/cluster-types/{id}/ | 
[**VirtualizationClusterTypesUpdate**](VirtualizationApi.md#VirtualizationClusterTypesUpdate) | **Put** /virtualization/cluster-types/{id}/ | 
[**VirtualizationClustersCreate**](VirtualizationApi.md#VirtualizationClustersCreate) | **Post** /virtualization/clusters/ | 
[**VirtualizationClustersDelete**](VirtualizationApi.md#VirtualizationClustersDelete) | **Delete** /virtualization/clusters/{id}/ | 
[**VirtualizationClustersList**](VirtualizationApi.md#VirtualizationClustersList) | **Get** /virtualization/clusters/ | 
[**VirtualizationClustersPartialUpdate**](VirtualizationApi.md#VirtualizationClustersPartialUpdate) | **Patch** /virtualization/clusters/{id}/ | 
[**VirtualizationClustersRead**](VirtualizationApi.md#VirtualizationClustersRead) | **Get** /virtualization/clusters/{id}/ | 
[**VirtualizationClustersUpdate**](VirtualizationApi.md#VirtualizationClustersUpdate) | **Put** /virtualization/clusters/{id}/ | 
[**VirtualizationInterfacesCreate**](VirtualizationApi.md#VirtualizationInterfacesCreate) | **Post** /virtualization/interfaces/ | 
[**VirtualizationInterfacesDelete**](VirtualizationApi.md#VirtualizationInterfacesDelete) | **Delete** /virtualization/interfaces/{id}/ | 
[**VirtualizationInterfacesList**](VirtualizationApi.md#VirtualizationInterfacesList) | **Get** /virtualization/interfaces/ | 
[**VirtualizationInterfacesPartialUpdate**](VirtualizationApi.md#VirtualizationInterfacesPartialUpdate) | **Patch** /virtualization/interfaces/{id}/ | 
[**VirtualizationInterfacesRead**](VirtualizationApi.md#VirtualizationInterfacesRead) | **Get** /virtualization/interfaces/{id}/ | 
[**VirtualizationInterfacesUpdate**](VirtualizationApi.md#VirtualizationInterfacesUpdate) | **Put** /virtualization/interfaces/{id}/ | 
[**VirtualizationVirtualMachinesCreate**](VirtualizationApi.md#VirtualizationVirtualMachinesCreate) | **Post** /virtualization/virtual-machines/ | 
[**VirtualizationVirtualMachinesDelete**](VirtualizationApi.md#VirtualizationVirtualMachinesDelete) | **Delete** /virtualization/virtual-machines/{id}/ | 
[**VirtualizationVirtualMachinesList**](VirtualizationApi.md#VirtualizationVirtualMachinesList) | **Get** /virtualization/virtual-machines/ | 
[**VirtualizationVirtualMachinesPartialUpdate**](VirtualizationApi.md#VirtualizationVirtualMachinesPartialUpdate) | **Patch** /virtualization/virtual-machines/{id}/ | 
[**VirtualizationVirtualMachinesRead**](VirtualizationApi.md#VirtualizationVirtualMachinesRead) | **Get** /virtualization/virtual-machines/{id}/ | 
[**VirtualizationVirtualMachinesUpdate**](VirtualizationApi.md#VirtualizationVirtualMachinesUpdate) | **Put** /virtualization/virtual-machines/{id}/ | 


# **VirtualizationClusterGroupsCreate**
> ClusterGroup VirtualizationClusterGroupsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ClusterGroup**](ClusterGroup.md)|  | 

### Return type

[**ClusterGroup**](ClusterGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterGroupsDelete**
> VirtualizationClusterGroupsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster group. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterGroupsList**
> InlineResponse20056 VirtualizationClusterGroupsList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualizationApiVirtualizationClusterGroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualizationApiVirtualizationClusterGroupsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **slug** | **optional.String**|  | 
 **description** | **optional.String**|  | 
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
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20056**](inline_response_200_56.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterGroupsPartialUpdate**
> ClusterGroup VirtualizationClusterGroupsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster group. | 
  **data** | [**ClusterGroup**](ClusterGroup.md)|  | 

### Return type

[**ClusterGroup**](ClusterGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterGroupsRead**
> ClusterGroup VirtualizationClusterGroupsRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster group. | 

### Return type

[**ClusterGroup**](ClusterGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterGroupsUpdate**
> ClusterGroup VirtualizationClusterGroupsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster group. | 
  **data** | [**ClusterGroup**](ClusterGroup.md)|  | 

### Return type

[**ClusterGroup**](ClusterGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterTypesCreate**
> ClusterType VirtualizationClusterTypesCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ClusterType**](ClusterType.md)|  | 

### Return type

[**ClusterType**](ClusterType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterTypesDelete**
> VirtualizationClusterTypesDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster type. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterTypesList**
> InlineResponse20057 VirtualizationClusterTypesList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualizationApiVirtualizationClusterTypesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualizationApiVirtualizationClusterTypesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **slug** | **optional.String**|  | 
 **description** | **optional.String**|  | 
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
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterTypesPartialUpdate**
> ClusterType VirtualizationClusterTypesPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster type. | 
  **data** | [**ClusterType**](ClusterType.md)|  | 

### Return type

[**ClusterType**](ClusterType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterTypesRead**
> ClusterType VirtualizationClusterTypesRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster type. | 

### Return type

[**ClusterType**](ClusterType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClusterTypesUpdate**
> ClusterType VirtualizationClusterTypesUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster type. | 
  **data** | [**ClusterType**](ClusterType.md)|  | 

### Return type

[**ClusterType**](ClusterType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClustersCreate**
> Cluster VirtualizationClustersCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableCluster**](WritableCluster.md)|  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClustersDelete**
> VirtualizationClustersDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClustersList**
> InlineResponse20058 VirtualizationClustersList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualizationApiVirtualizationClustersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualizationApiVirtualizationClustersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **tenantGroupId** | **optional.String**|  | 
 **tenantGroup** | **optional.String**|  | 
 **tenantId** | **optional.String**|  | 
 **tenant** | **optional.String**|  | 
 **created** | **optional.String**|  | 
 **createdGte** | **optional.String**|  | 
 **createdLte** | **optional.String**|  | 
 **lastUpdated** | **optional.String**|  | 
 **lastUpdatedGte** | **optional.String**|  | 
 **lastUpdatedLte** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **regionId** | **optional.String**|  | 
 **region** | **optional.String**|  | 
 **siteId** | **optional.String**|  | 
 **site** | **optional.String**|  | 
 **groupId** | **optional.String**|  | 
 **group** | **optional.String**|  | 
 **typeId** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
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
 **tenantGroupIdN** | **optional.String**|  | 
 **tenantGroupN** | **optional.String**|  | 
 **tenantIdN** | **optional.String**|  | 
 **tenantN** | **optional.String**|  | 
 **regionIdN** | **optional.String**|  | 
 **regionN** | **optional.String**|  | 
 **siteIdN** | **optional.String**|  | 
 **siteN** | **optional.String**|  | 
 **groupIdN** | **optional.String**|  | 
 **groupN** | **optional.String**|  | 
 **typeIdN** | **optional.String**|  | 
 **typeN** | **optional.String**|  | 
 **tagN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClustersPartialUpdate**
> Cluster VirtualizationClustersPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster. | 
  **data** | [**WritableCluster**](WritableCluster.md)|  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClustersRead**
> Cluster VirtualizationClustersRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster. | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationClustersUpdate**
> Cluster VirtualizationClustersUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this cluster. | 
  **data** | [**WritableCluster**](WritableCluster.md)|  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationInterfacesCreate**
> VirtualMachineInterface VirtualizationInterfacesCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableVirtualMachineInterface**](WritableVirtualMachineInterface.md)|  | 

### Return type

[**VirtualMachineInterface**](VirtualMachineInterface.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationInterfacesDelete**
> VirtualizationInterfacesDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this interface. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationInterfacesList**
> InlineResponse20059 VirtualizationInterfacesList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualizationApiVirtualizationInterfacesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualizationApiVirtualizationInterfacesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **enabled** | **optional.String**|  | 
 **mtu** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **virtualMachineId** | **optional.String**|  | 
 **virtualMachine** | **optional.String**|  | 
 **macAddress** | **optional.String**|  | 
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
 **mtuN** | **optional.String**|  | 
 **mtuLte** | **optional.String**|  | 
 **mtuLt** | **optional.String**|  | 
 **mtuGte** | **optional.String**|  | 
 **mtuGt** | **optional.String**|  | 
 **virtualMachineIdN** | **optional.String**|  | 
 **virtualMachineN** | **optional.String**|  | 
 **macAddressN** | **optional.String**|  | 
 **macAddressIc** | **optional.String**|  | 
 **macAddressNic** | **optional.String**|  | 
 **macAddressIew** | **optional.String**|  | 
 **macAddressNiew** | **optional.String**|  | 
 **macAddressIsw** | **optional.String**|  | 
 **macAddressNisw** | **optional.String**|  | 
 **macAddressIe** | **optional.String**|  | 
 **macAddressNie** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20059**](inline_response_200_59.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationInterfacesPartialUpdate**
> VirtualMachineInterface VirtualizationInterfacesPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this interface. | 
  **data** | [**WritableVirtualMachineInterface**](WritableVirtualMachineInterface.md)|  | 

### Return type

[**VirtualMachineInterface**](VirtualMachineInterface.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationInterfacesRead**
> VirtualMachineInterface VirtualizationInterfacesRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this interface. | 

### Return type

[**VirtualMachineInterface**](VirtualMachineInterface.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationInterfacesUpdate**
> VirtualMachineInterface VirtualizationInterfacesUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this interface. | 
  **data** | [**WritableVirtualMachineInterface**](WritableVirtualMachineInterface.md)|  | 

### Return type

[**VirtualMachineInterface**](VirtualMachineInterface.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationVirtualMachinesCreate**
> VirtualMachineWithConfigContext VirtualizationVirtualMachinesCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableVirtualMachineWithConfigContext**](WritableVirtualMachineWithConfigContext.md)|  | 

### Return type

[**VirtualMachineWithConfigContext**](VirtualMachineWithConfigContext.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationVirtualMachinesDelete**
> VirtualizationVirtualMachinesDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this virtual machine. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationVirtualMachinesList**
> InlineResponse20060 VirtualizationVirtualMachinesList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VirtualizationApiVirtualizationVirtualMachinesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VirtualizationApiVirtualizationVirtualMachinesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **cluster** | **optional.String**|  | 
 **vcpus** | **optional.String**|  | 
 **memory** | **optional.String**|  | 
 **disk** | **optional.String**|  | 
 **localContextData** | **optional.String**|  | 
 **tenantGroupId** | **optional.String**|  | 
 **tenantGroup** | **optional.String**|  | 
 **tenantId** | **optional.String**|  | 
 **tenant** | **optional.String**|  | 
 **created** | **optional.String**|  | 
 **createdGte** | **optional.String**|  | 
 **createdLte** | **optional.String**|  | 
 **lastUpdated** | **optional.String**|  | 
 **lastUpdatedGte** | **optional.String**|  | 
 **lastUpdatedLte** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **status** | **optional.String**|  | 
 **clusterGroupId** | **optional.String**|  | 
 **clusterGroup** | **optional.String**|  | 
 **clusterTypeId** | **optional.String**|  | 
 **clusterType** | **optional.String**|  | 
 **clusterId** | **optional.String**|  | 
 **regionId** | **optional.String**|  | 
 **region** | **optional.String**|  | 
 **siteId** | **optional.String**|  | 
 **site** | **optional.String**|  | 
 **roleId** | **optional.String**|  | 
 **role** | **optional.String**|  | 
 **platformId** | **optional.String**|  | 
 **platform** | **optional.String**|  | 
 **macAddress** | **optional.String**|  | 
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
 **clusterN** | **optional.String**|  | 
 **vcpusN** | **optional.String**|  | 
 **vcpusLte** | **optional.String**|  | 
 **vcpusLt** | **optional.String**|  | 
 **vcpusGte** | **optional.String**|  | 
 **vcpusGt** | **optional.String**|  | 
 **memoryN** | **optional.String**|  | 
 **memoryLte** | **optional.String**|  | 
 **memoryLt** | **optional.String**|  | 
 **memoryGte** | **optional.String**|  | 
 **memoryGt** | **optional.String**|  | 
 **diskN** | **optional.String**|  | 
 **diskLte** | **optional.String**|  | 
 **diskLt** | **optional.String**|  | 
 **diskGte** | **optional.String**|  | 
 **diskGt** | **optional.String**|  | 
 **tenantGroupIdN** | **optional.String**|  | 
 **tenantGroupN** | **optional.String**|  | 
 **tenantIdN** | **optional.String**|  | 
 **tenantN** | **optional.String**|  | 
 **statusN** | **optional.String**|  | 
 **clusterGroupIdN** | **optional.String**|  | 
 **clusterGroupN** | **optional.String**|  | 
 **clusterTypeIdN** | **optional.String**|  | 
 **clusterTypeN** | **optional.String**|  | 
 **clusterIdN** | **optional.String**|  | 
 **regionIdN** | **optional.String**|  | 
 **regionN** | **optional.String**|  | 
 **siteIdN** | **optional.String**|  | 
 **siteN** | **optional.String**|  | 
 **roleIdN** | **optional.String**|  | 
 **roleN** | **optional.String**|  | 
 **platformIdN** | **optional.String**|  | 
 **platformN** | **optional.String**|  | 
 **macAddressN** | **optional.String**|  | 
 **macAddressIc** | **optional.String**|  | 
 **macAddressNic** | **optional.String**|  | 
 **macAddressIew** | **optional.String**|  | 
 **macAddressNiew** | **optional.String**|  | 
 **macAddressIsw** | **optional.String**|  | 
 **macAddressNisw** | **optional.String**|  | 
 **macAddressIe** | **optional.String**|  | 
 **macAddressNie** | **optional.String**|  | 
 **tagN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationVirtualMachinesPartialUpdate**
> VirtualMachineWithConfigContext VirtualizationVirtualMachinesPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this virtual machine. | 
  **data** | [**WritableVirtualMachineWithConfigContext**](WritableVirtualMachineWithConfigContext.md)|  | 

### Return type

[**VirtualMachineWithConfigContext**](VirtualMachineWithConfigContext.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationVirtualMachinesRead**
> VirtualMachineWithConfigContext VirtualizationVirtualMachinesRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this virtual machine. | 

### Return type

[**VirtualMachineWithConfigContext**](VirtualMachineWithConfigContext.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VirtualizationVirtualMachinesUpdate**
> VirtualMachineWithConfigContext VirtualizationVirtualMachinesUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this virtual machine. | 
  **data** | [**WritableVirtualMachineWithConfigContext**](WritableVirtualMachineWithConfigContext.md)|  | 

### Return type

[**VirtualMachineWithConfigContext**](VirtualMachineWithConfigContext.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

