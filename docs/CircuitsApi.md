# \CircuitsApi

All URIs are relative to *https://netbox/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CircuitsCircuitTerminationsCreate**](CircuitsApi.md#CircuitsCircuitTerminationsCreate) | **Post** /circuits/circuit-terminations/ | 
[**CircuitsCircuitTerminationsDelete**](CircuitsApi.md#CircuitsCircuitTerminationsDelete) | **Delete** /circuits/circuit-terminations/{id}/ | 
[**CircuitsCircuitTerminationsList**](CircuitsApi.md#CircuitsCircuitTerminationsList) | **Get** /circuits/circuit-terminations/ | 
[**CircuitsCircuitTerminationsPartialUpdate**](CircuitsApi.md#CircuitsCircuitTerminationsPartialUpdate) | **Patch** /circuits/circuit-terminations/{id}/ | 
[**CircuitsCircuitTerminationsRead**](CircuitsApi.md#CircuitsCircuitTerminationsRead) | **Get** /circuits/circuit-terminations/{id}/ | 
[**CircuitsCircuitTerminationsUpdate**](CircuitsApi.md#CircuitsCircuitTerminationsUpdate) | **Put** /circuits/circuit-terminations/{id}/ | 
[**CircuitsCircuitTypesCreate**](CircuitsApi.md#CircuitsCircuitTypesCreate) | **Post** /circuits/circuit-types/ | 
[**CircuitsCircuitTypesDelete**](CircuitsApi.md#CircuitsCircuitTypesDelete) | **Delete** /circuits/circuit-types/{id}/ | 
[**CircuitsCircuitTypesList**](CircuitsApi.md#CircuitsCircuitTypesList) | **Get** /circuits/circuit-types/ | 
[**CircuitsCircuitTypesPartialUpdate**](CircuitsApi.md#CircuitsCircuitTypesPartialUpdate) | **Patch** /circuits/circuit-types/{id}/ | 
[**CircuitsCircuitTypesRead**](CircuitsApi.md#CircuitsCircuitTypesRead) | **Get** /circuits/circuit-types/{id}/ | 
[**CircuitsCircuitTypesUpdate**](CircuitsApi.md#CircuitsCircuitTypesUpdate) | **Put** /circuits/circuit-types/{id}/ | 
[**CircuitsCircuitsCreate**](CircuitsApi.md#CircuitsCircuitsCreate) | **Post** /circuits/circuits/ | 
[**CircuitsCircuitsDelete**](CircuitsApi.md#CircuitsCircuitsDelete) | **Delete** /circuits/circuits/{id}/ | 
[**CircuitsCircuitsList**](CircuitsApi.md#CircuitsCircuitsList) | **Get** /circuits/circuits/ | 
[**CircuitsCircuitsPartialUpdate**](CircuitsApi.md#CircuitsCircuitsPartialUpdate) | **Patch** /circuits/circuits/{id}/ | 
[**CircuitsCircuitsRead**](CircuitsApi.md#CircuitsCircuitsRead) | **Get** /circuits/circuits/{id}/ | 
[**CircuitsCircuitsUpdate**](CircuitsApi.md#CircuitsCircuitsUpdate) | **Put** /circuits/circuits/{id}/ | 
[**CircuitsProvidersCreate**](CircuitsApi.md#CircuitsProvidersCreate) | **Post** /circuits/providers/ | 
[**CircuitsProvidersDelete**](CircuitsApi.md#CircuitsProvidersDelete) | **Delete** /circuits/providers/{id}/ | 
[**CircuitsProvidersGraphs**](CircuitsApi.md#CircuitsProvidersGraphs) | **Get** /circuits/providers/{id}/graphs/ | 
[**CircuitsProvidersList**](CircuitsApi.md#CircuitsProvidersList) | **Get** /circuits/providers/ | 
[**CircuitsProvidersPartialUpdate**](CircuitsApi.md#CircuitsProvidersPartialUpdate) | **Patch** /circuits/providers/{id}/ | 
[**CircuitsProvidersRead**](CircuitsApi.md#CircuitsProvidersRead) | **Get** /circuits/providers/{id}/ | 
[**CircuitsProvidersUpdate**](CircuitsApi.md#CircuitsProvidersUpdate) | **Put** /circuits/providers/{id}/ | 


# **CircuitsCircuitTerminationsCreate**
> CircuitTermination CircuitsCircuitTerminationsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableCircuitTermination**](WritableCircuitTermination.md)|  | 

### Return type

[**CircuitTermination**](CircuitTermination.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTerminationsDelete**
> CircuitsCircuitTerminationsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit termination. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTerminationsList**
> InlineResponse200 CircuitsCircuitTerminationsList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CircuitsApiCircuitsCircuitTerminationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CircuitsApiCircuitsCircuitTerminationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **termSide** | **optional.String**|  | 
 **portSpeed** | **optional.String**|  | 
 **upstreamSpeed** | **optional.String**|  | 
 **xconnectId** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **circuitId** | **optional.String**|  | 
 **siteId** | **optional.String**|  | 
 **site** | **optional.String**|  | 
 **termSideN** | **optional.String**|  | 
 **portSpeedN** | **optional.String**|  | 
 **portSpeedLte** | **optional.String**|  | 
 **portSpeedLt** | **optional.String**|  | 
 **portSpeedGte** | **optional.String**|  | 
 **portSpeedGt** | **optional.String**|  | 
 **upstreamSpeedN** | **optional.String**|  | 
 **upstreamSpeedLte** | **optional.String**|  | 
 **upstreamSpeedLt** | **optional.String**|  | 
 **upstreamSpeedGte** | **optional.String**|  | 
 **upstreamSpeedGt** | **optional.String**|  | 
 **xconnectIdN** | **optional.String**|  | 
 **xconnectIdIc** | **optional.String**|  | 
 **xconnectIdNic** | **optional.String**|  | 
 **xconnectIdIew** | **optional.String**|  | 
 **xconnectIdNiew** | **optional.String**|  | 
 **xconnectIdIsw** | **optional.String**|  | 
 **xconnectIdNisw** | **optional.String**|  | 
 **xconnectIdIe** | **optional.String**|  | 
 **xconnectIdNie** | **optional.String**|  | 
 **circuitIdN** | **optional.String**|  | 
 **siteIdN** | **optional.String**|  | 
 **siteN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTerminationsPartialUpdate**
> CircuitTermination CircuitsCircuitTerminationsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit termination. | 
  **data** | [**WritableCircuitTermination**](WritableCircuitTermination.md)|  | 

### Return type

[**CircuitTermination**](CircuitTermination.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTerminationsRead**
> CircuitTermination CircuitsCircuitTerminationsRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit termination. | 

### Return type

[**CircuitTermination**](CircuitTermination.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTerminationsUpdate**
> CircuitTermination CircuitsCircuitTerminationsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit termination. | 
  **data** | [**WritableCircuitTermination**](WritableCircuitTermination.md)|  | 

### Return type

[**CircuitTermination**](CircuitTermination.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTypesCreate**
> CircuitType CircuitsCircuitTypesCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CircuitType**](CircuitType.md)|  | 

### Return type

[**CircuitType**](CircuitType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTypesDelete**
> CircuitsCircuitTypesDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit type. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTypesList**
> InlineResponse2001 CircuitsCircuitTypesList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CircuitsApiCircuitsCircuitTypesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CircuitsApiCircuitsCircuitTypesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **slug** | **optional.String**|  | 
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
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTypesPartialUpdate**
> CircuitType CircuitsCircuitTypesPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit type. | 
  **data** | [**CircuitType**](CircuitType.md)|  | 

### Return type

[**CircuitType**](CircuitType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTypesRead**
> CircuitType CircuitsCircuitTypesRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit type. | 

### Return type

[**CircuitType**](CircuitType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitTypesUpdate**
> CircuitType CircuitsCircuitTypesUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit type. | 
  **data** | [**CircuitType**](CircuitType.md)|  | 

### Return type

[**CircuitType**](CircuitType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitsCreate**
> Circuit CircuitsCircuitsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableCircuit**](WritableCircuit.md)|  | 

### Return type

[**Circuit**](Circuit.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitsDelete**
> CircuitsCircuitsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitsList**
> InlineResponse2002 CircuitsCircuitsList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CircuitsApiCircuitsCircuitsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CircuitsApiCircuitsCircuitsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **cid** | **optional.String**|  | 
 **installDate** | **optional.String**|  | 
 **commitRate** | **optional.String**|  | 
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
 **providerId** | **optional.String**|  | 
 **provider** | **optional.String**|  | 
 **typeId** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **status** | **optional.String**|  | 
 **siteId** | **optional.String**|  | 
 **site** | **optional.String**|  | 
 **regionId** | **optional.String**|  | 
 **region** | **optional.String**|  | 
 **tag** | **optional.String**|  | 
 **idN** | **optional.String**|  | 
 **idLte** | **optional.String**|  | 
 **idLt** | **optional.String**|  | 
 **idGte** | **optional.String**|  | 
 **idGt** | **optional.String**|  | 
 **cidN** | **optional.String**|  | 
 **cidIc** | **optional.String**|  | 
 **cidNic** | **optional.String**|  | 
 **cidIew** | **optional.String**|  | 
 **cidNiew** | **optional.String**|  | 
 **cidIsw** | **optional.String**|  | 
 **cidNisw** | **optional.String**|  | 
 **cidIe** | **optional.String**|  | 
 **cidNie** | **optional.String**|  | 
 **installDateN** | **optional.String**|  | 
 **installDateLte** | **optional.String**|  | 
 **installDateLt** | **optional.String**|  | 
 **installDateGte** | **optional.String**|  | 
 **installDateGt** | **optional.String**|  | 
 **commitRateN** | **optional.String**|  | 
 **commitRateLte** | **optional.String**|  | 
 **commitRateLt** | **optional.String**|  | 
 **commitRateGte** | **optional.String**|  | 
 **commitRateGt** | **optional.String**|  | 
 **tenantGroupIdN** | **optional.String**|  | 
 **tenantGroupN** | **optional.String**|  | 
 **tenantIdN** | **optional.String**|  | 
 **tenantN** | **optional.String**|  | 
 **providerIdN** | **optional.String**|  | 
 **providerN** | **optional.String**|  | 
 **typeIdN** | **optional.String**|  | 
 **typeN** | **optional.String**|  | 
 **statusN** | **optional.String**|  | 
 **siteIdN** | **optional.String**|  | 
 **siteN** | **optional.String**|  | 
 **regionIdN** | **optional.String**|  | 
 **regionN** | **optional.String**|  | 
 **tagN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitsPartialUpdate**
> Circuit CircuitsCircuitsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit. | 
  **data** | [**WritableCircuit**](WritableCircuit.md)|  | 

### Return type

[**Circuit**](Circuit.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitsRead**
> Circuit CircuitsCircuitsRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit. | 

### Return type

[**Circuit**](Circuit.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsCircuitsUpdate**
> Circuit CircuitsCircuitsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this circuit. | 
  **data** | [**WritableCircuit**](WritableCircuit.md)|  | 

### Return type

[**Circuit**](Circuit.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsProvidersCreate**
> Provider CircuitsProvidersCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Provider**](Provider.md)|  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsProvidersDelete**
> CircuitsProvidersDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this provider. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsProvidersGraphs**
> Provider CircuitsProvidersGraphs(ctx, id)


A convenience method for rendering graphs for a particular provider.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this provider. | 

### Return type

[**Provider**](Provider.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsProvidersList**
> InlineResponse2003 CircuitsProvidersList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CircuitsApiCircuitsProvidersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CircuitsApiCircuitsProvidersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **slug** | **optional.String**|  | 
 **asn** | **optional.String**|  | 
 **account** | **optional.String**|  | 
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
 **asnN** | **optional.String**|  | 
 **asnLte** | **optional.String**|  | 
 **asnLt** | **optional.String**|  | 
 **asnGte** | **optional.String**|  | 
 **asnGt** | **optional.String**|  | 
 **accountN** | **optional.String**|  | 
 **accountIc** | **optional.String**|  | 
 **accountNic** | **optional.String**|  | 
 **accountIew** | **optional.String**|  | 
 **accountNiew** | **optional.String**|  | 
 **accountIsw** | **optional.String**|  | 
 **accountNisw** | **optional.String**|  | 
 **accountIe** | **optional.String**|  | 
 **accountNie** | **optional.String**|  | 
 **regionIdN** | **optional.String**|  | 
 **regionN** | **optional.String**|  | 
 **siteIdN** | **optional.String**|  | 
 **siteN** | **optional.String**|  | 
 **tagN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsProvidersPartialUpdate**
> Provider CircuitsProvidersPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this provider. | 
  **data** | [**Provider**](Provider.md)|  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsProvidersRead**
> Provider CircuitsProvidersRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this provider. | 

### Return type

[**Provider**](Provider.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CircuitsProvidersUpdate**
> Provider CircuitsProvidersUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this provider. | 
  **data** | [**Provider**](Provider.md)|  | 

### Return type

[**Provider**](Provider.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

