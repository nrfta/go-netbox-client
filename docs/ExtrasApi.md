# \ExtrasApi

All URIs are relative to *https://netbox/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtrasConfigContextsCreate**](ExtrasApi.md#ExtrasConfigContextsCreate) | **Post** /extras/config-contexts/ | 
[**ExtrasConfigContextsDelete**](ExtrasApi.md#ExtrasConfigContextsDelete) | **Delete** /extras/config-contexts/{id}/ | 
[**ExtrasConfigContextsList**](ExtrasApi.md#ExtrasConfigContextsList) | **Get** /extras/config-contexts/ | 
[**ExtrasConfigContextsPartialUpdate**](ExtrasApi.md#ExtrasConfigContextsPartialUpdate) | **Patch** /extras/config-contexts/{id}/ | 
[**ExtrasConfigContextsRead**](ExtrasApi.md#ExtrasConfigContextsRead) | **Get** /extras/config-contexts/{id}/ | 
[**ExtrasConfigContextsUpdate**](ExtrasApi.md#ExtrasConfigContextsUpdate) | **Put** /extras/config-contexts/{id}/ | 
[**ExtrasCustomFieldChoicesList**](ExtrasApi.md#ExtrasCustomFieldChoicesList) | **Get** /extras/_custom_field_choices/ | 
[**ExtrasCustomFieldChoicesRead**](ExtrasApi.md#ExtrasCustomFieldChoicesRead) | **Get** /extras/_custom_field_choices/{id}/ | 
[**ExtrasExportTemplatesCreate**](ExtrasApi.md#ExtrasExportTemplatesCreate) | **Post** /extras/export-templates/ | 
[**ExtrasExportTemplatesDelete**](ExtrasApi.md#ExtrasExportTemplatesDelete) | **Delete** /extras/export-templates/{id}/ | 
[**ExtrasExportTemplatesList**](ExtrasApi.md#ExtrasExportTemplatesList) | **Get** /extras/export-templates/ | 
[**ExtrasExportTemplatesPartialUpdate**](ExtrasApi.md#ExtrasExportTemplatesPartialUpdate) | **Patch** /extras/export-templates/{id}/ | 
[**ExtrasExportTemplatesRead**](ExtrasApi.md#ExtrasExportTemplatesRead) | **Get** /extras/export-templates/{id}/ | 
[**ExtrasExportTemplatesUpdate**](ExtrasApi.md#ExtrasExportTemplatesUpdate) | **Put** /extras/export-templates/{id}/ | 
[**ExtrasGraphsCreate**](ExtrasApi.md#ExtrasGraphsCreate) | **Post** /extras/graphs/ | 
[**ExtrasGraphsDelete**](ExtrasApi.md#ExtrasGraphsDelete) | **Delete** /extras/graphs/{id}/ | 
[**ExtrasGraphsList**](ExtrasApi.md#ExtrasGraphsList) | **Get** /extras/graphs/ | 
[**ExtrasGraphsPartialUpdate**](ExtrasApi.md#ExtrasGraphsPartialUpdate) | **Patch** /extras/graphs/{id}/ | 
[**ExtrasGraphsRead**](ExtrasApi.md#ExtrasGraphsRead) | **Get** /extras/graphs/{id}/ | 
[**ExtrasGraphsUpdate**](ExtrasApi.md#ExtrasGraphsUpdate) | **Put** /extras/graphs/{id}/ | 
[**ExtrasImageAttachmentsCreate**](ExtrasApi.md#ExtrasImageAttachmentsCreate) | **Post** /extras/image-attachments/ | 
[**ExtrasImageAttachmentsDelete**](ExtrasApi.md#ExtrasImageAttachmentsDelete) | **Delete** /extras/image-attachments/{id}/ | 
[**ExtrasImageAttachmentsList**](ExtrasApi.md#ExtrasImageAttachmentsList) | **Get** /extras/image-attachments/ | 
[**ExtrasImageAttachmentsPartialUpdate**](ExtrasApi.md#ExtrasImageAttachmentsPartialUpdate) | **Patch** /extras/image-attachments/{id}/ | 
[**ExtrasImageAttachmentsRead**](ExtrasApi.md#ExtrasImageAttachmentsRead) | **Get** /extras/image-attachments/{id}/ | 
[**ExtrasImageAttachmentsUpdate**](ExtrasApi.md#ExtrasImageAttachmentsUpdate) | **Put** /extras/image-attachments/{id}/ | 
[**ExtrasObjectChangesList**](ExtrasApi.md#ExtrasObjectChangesList) | **Get** /extras/object-changes/ | 
[**ExtrasObjectChangesRead**](ExtrasApi.md#ExtrasObjectChangesRead) | **Get** /extras/object-changes/{id}/ | 
[**ExtrasReportsList**](ExtrasApi.md#ExtrasReportsList) | **Get** /extras/reports/ | 
[**ExtrasReportsRead**](ExtrasApi.md#ExtrasReportsRead) | **Get** /extras/reports/{id}/ | 
[**ExtrasReportsRun**](ExtrasApi.md#ExtrasReportsRun) | **Post** /extras/reports/{id}/run/ | 
[**ExtrasScriptsList**](ExtrasApi.md#ExtrasScriptsList) | **Get** /extras/scripts/ | 
[**ExtrasScriptsRead**](ExtrasApi.md#ExtrasScriptsRead) | **Get** /extras/scripts/{id}/ | 
[**ExtrasTagsCreate**](ExtrasApi.md#ExtrasTagsCreate) | **Post** /extras/tags/ | 
[**ExtrasTagsDelete**](ExtrasApi.md#ExtrasTagsDelete) | **Delete** /extras/tags/{id}/ | 
[**ExtrasTagsList**](ExtrasApi.md#ExtrasTagsList) | **Get** /extras/tags/ | 
[**ExtrasTagsPartialUpdate**](ExtrasApi.md#ExtrasTagsPartialUpdate) | **Patch** /extras/tags/{id}/ | 
[**ExtrasTagsRead**](ExtrasApi.md#ExtrasTagsRead) | **Get** /extras/tags/{id}/ | 
[**ExtrasTagsUpdate**](ExtrasApi.md#ExtrasTagsUpdate) | **Put** /extras/tags/{id}/ | 


# **ExtrasConfigContextsCreate**
> ConfigContext ExtrasConfigContextsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableConfigContext**](WritableConfigContext.md)|  | 

### Return type

[**ConfigContext**](ConfigContext.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasConfigContextsDelete**
> ExtrasConfigContextsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this config context. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasConfigContextsList**
> InlineResponse20037 ExtrasConfigContextsList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtrasApiExtrasConfigContextsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtrasApiExtrasConfigContextsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **isActive** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **regionId** | **optional.String**|  | 
 **region** | **optional.String**|  | 
 **siteId** | **optional.String**|  | 
 **site** | **optional.String**|  | 
 **roleId** | **optional.String**|  | 
 **role** | **optional.String**|  | 
 **platformId** | **optional.String**|  | 
 **platform** | **optional.String**|  | 
 **clusterGroupId** | **optional.String**|  | 
 **clusterGroup** | **optional.String**|  | 
 **clusterId** | **optional.String**|  | 
 **tenantGroupId** | **optional.String**|  | 
 **tenantGroup** | **optional.String**|  | 
 **tenantId** | **optional.String**|  | 
 **tenant** | **optional.String**|  | 
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
 **regionIdN** | **optional.String**|  | 
 **regionN** | **optional.String**|  | 
 **siteIdN** | **optional.String**|  | 
 **siteN** | **optional.String**|  | 
 **roleIdN** | **optional.String**|  | 
 **roleN** | **optional.String**|  | 
 **platformIdN** | **optional.String**|  | 
 **platformN** | **optional.String**|  | 
 **clusterGroupIdN** | **optional.String**|  | 
 **clusterGroupN** | **optional.String**|  | 
 **clusterIdN** | **optional.String**|  | 
 **tenantGroupIdN** | **optional.String**|  | 
 **tenantGroupN** | **optional.String**|  | 
 **tenantIdN** | **optional.String**|  | 
 **tenantN** | **optional.String**|  | 
 **tagN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasConfigContextsPartialUpdate**
> ConfigContext ExtrasConfigContextsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this config context. | 
  **data** | [**WritableConfigContext**](WritableConfigContext.md)|  | 

### Return type

[**ConfigContext**](ConfigContext.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasConfigContextsRead**
> ConfigContext ExtrasConfigContextsRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this config context. | 

### Return type

[**ConfigContext**](ConfigContext.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasConfigContextsUpdate**
> ConfigContext ExtrasConfigContextsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this config context. | 
  **data** | [**WritableConfigContext**](WritableConfigContext.md)|  | 

### Return type

[**ConfigContext**](ConfigContext.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasCustomFieldChoicesList**
> ExtrasCustomFieldChoicesList(ctx, )




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

# **ExtrasCustomFieldChoicesRead**
> ExtrasCustomFieldChoicesRead(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasExportTemplatesCreate**
> ExportTemplate ExtrasExportTemplatesCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableExportTemplate**](WritableExportTemplate.md)|  | 

### Return type

[**ExportTemplate**](ExportTemplate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasExportTemplatesDelete**
> ExtrasExportTemplatesDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this export template. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasExportTemplatesList**
> InlineResponse20038 ExtrasExportTemplatesList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtrasApiExtrasExportTemplatesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtrasApiExtrasExportTemplatesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **contentType** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **templateLanguage** | **optional.String**|  | 
 **idN** | **optional.String**|  | 
 **idLte** | **optional.String**|  | 
 **idLt** | **optional.String**|  | 
 **idGte** | **optional.String**|  | 
 **idGt** | **optional.String**|  | 
 **contentTypeN** | **optional.String**|  | 
 **nameN** | **optional.String**|  | 
 **nameIc** | **optional.String**|  | 
 **nameNic** | **optional.String**|  | 
 **nameIew** | **optional.String**|  | 
 **nameNiew** | **optional.String**|  | 
 **nameIsw** | **optional.String**|  | 
 **nameNisw** | **optional.String**|  | 
 **nameIe** | **optional.String**|  | 
 **nameNie** | **optional.String**|  | 
 **templateLanguageN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasExportTemplatesPartialUpdate**
> ExportTemplate ExtrasExportTemplatesPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this export template. | 
  **data** | [**WritableExportTemplate**](WritableExportTemplate.md)|  | 

### Return type

[**ExportTemplate**](ExportTemplate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasExportTemplatesRead**
> ExportTemplate ExtrasExportTemplatesRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this export template. | 

### Return type

[**ExportTemplate**](ExportTemplate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasExportTemplatesUpdate**
> ExportTemplate ExtrasExportTemplatesUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this export template. | 
  **data** | [**WritableExportTemplate**](WritableExportTemplate.md)|  | 

### Return type

[**ExportTemplate**](ExportTemplate.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasGraphsCreate**
> Graph ExtrasGraphsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Graph**](Graph.md)|  | 

### Return type

[**Graph**](Graph.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasGraphsDelete**
> ExtrasGraphsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this graph. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasGraphsList**
> InlineResponse20039 ExtrasGraphsList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtrasApiExtrasGraphsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtrasApiExtrasGraphsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **type_** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **templateLanguage** | **optional.String**|  | 
 **idN** | **optional.String**|  | 
 **idLte** | **optional.String**|  | 
 **idLt** | **optional.String**|  | 
 **idGte** | **optional.String**|  | 
 **idGt** | **optional.String**|  | 
 **typeN** | **optional.String**|  | 
 **nameN** | **optional.String**|  | 
 **nameIc** | **optional.String**|  | 
 **nameNic** | **optional.String**|  | 
 **nameIew** | **optional.String**|  | 
 **nameNiew** | **optional.String**|  | 
 **nameIsw** | **optional.String**|  | 
 **nameNisw** | **optional.String**|  | 
 **nameIe** | **optional.String**|  | 
 **nameNie** | **optional.String**|  | 
 **templateLanguageN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasGraphsPartialUpdate**
> Graph ExtrasGraphsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this graph. | 
  **data** | [**Graph**](Graph.md)|  | 

### Return type

[**Graph**](Graph.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasGraphsRead**
> Graph ExtrasGraphsRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this graph. | 

### Return type

[**Graph**](Graph.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasGraphsUpdate**
> Graph ExtrasGraphsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this graph. | 
  **data** | [**Graph**](Graph.md)|  | 

### Return type

[**Graph**](Graph.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasImageAttachmentsCreate**
> ImageAttachment ExtrasImageAttachmentsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ImageAttachment**](ImageAttachment.md)|  | 

### Return type

[**ImageAttachment**](ImageAttachment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasImageAttachmentsDelete**
> ExtrasImageAttachmentsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this image attachment. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasImageAttachmentsList**
> InlineResponse20040 ExtrasImageAttachmentsList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtrasApiExtrasImageAttachmentsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtrasApiExtrasImageAttachmentsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasImageAttachmentsPartialUpdate**
> ImageAttachment ExtrasImageAttachmentsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this image attachment. | 
  **data** | [**ImageAttachment**](ImageAttachment.md)|  | 

### Return type

[**ImageAttachment**](ImageAttachment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasImageAttachmentsRead**
> ImageAttachment ExtrasImageAttachmentsRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this image attachment. | 

### Return type

[**ImageAttachment**](ImageAttachment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasImageAttachmentsUpdate**
> ImageAttachment ExtrasImageAttachmentsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this image attachment. | 
  **data** | [**ImageAttachment**](ImageAttachment.md)|  | 

### Return type

[**ImageAttachment**](ImageAttachment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasObjectChangesList**
> InlineResponse20041 ExtrasObjectChangesList(ctx, optional)


Retrieve a list of recent changes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtrasApiExtrasObjectChangesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtrasApiExtrasObjectChangesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **user** | **optional.String**|  | 
 **userName** | **optional.String**|  | 
 **requestId** | **optional.String**|  | 
 **action** | **optional.String**|  | 
 **changedObjectType** | **optional.String**|  | 
 **changedObjectId** | **optional.String**|  | 
 **objectRepr** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **time** | **optional.String**|  | 
 **idN** | **optional.String**|  | 
 **idLte** | **optional.String**|  | 
 **idLt** | **optional.String**|  | 
 **idGte** | **optional.String**|  | 
 **idGt** | **optional.String**|  | 
 **userN** | **optional.String**|  | 
 **userNameN** | **optional.String**|  | 
 **userNameIc** | **optional.String**|  | 
 **userNameNic** | **optional.String**|  | 
 **userNameIew** | **optional.String**|  | 
 **userNameNiew** | **optional.String**|  | 
 **userNameIsw** | **optional.String**|  | 
 **userNameNisw** | **optional.String**|  | 
 **userNameIe** | **optional.String**|  | 
 **userNameNie** | **optional.String**|  | 
 **actionN** | **optional.String**|  | 
 **changedObjectTypeN** | **optional.String**|  | 
 **changedObjectIdN** | **optional.String**|  | 
 **changedObjectIdLte** | **optional.String**|  | 
 **changedObjectIdLt** | **optional.String**|  | 
 **changedObjectIdGte** | **optional.String**|  | 
 **changedObjectIdGt** | **optional.String**|  | 
 **objectReprN** | **optional.String**|  | 
 **objectReprIc** | **optional.String**|  | 
 **objectReprNic** | **optional.String**|  | 
 **objectReprIew** | **optional.String**|  | 
 **objectReprNiew** | **optional.String**|  | 
 **objectReprIsw** | **optional.String**|  | 
 **objectReprNisw** | **optional.String**|  | 
 **objectReprIe** | **optional.String**|  | 
 **objectReprNie** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasObjectChangesRead**
> ObjectChange ExtrasObjectChangesRead(ctx, id)


Retrieve a list of recent changes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this object change. | 

### Return type

[**ObjectChange**](ObjectChange.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasReportsList**
> ExtrasReportsList(ctx, )


Compile all reports and their related results (if any). Result data is deferred in the list view.

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

# **ExtrasReportsRead**
> ExtrasReportsRead(ctx, id)


Retrieve a single Report identified as \"<module>.<report>\".

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasReportsRun**
> ExtrasReportsRun(ctx, id)


Run a Report and create a new ReportResult, overwriting any previous result for the Report.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasScriptsList**
> ExtrasScriptsList(ctx, )




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

# **ExtrasScriptsRead**
> ExtrasScriptsRead(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasTagsCreate**
> Tag ExtrasTagsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**Tag**](Tag.md)|  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasTagsDelete**
> ExtrasTagsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tag. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasTagsList**
> InlineResponse20042 ExtrasTagsList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtrasApiExtrasTagsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtrasApiExtrasTagsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **slug** | **optional.String**|  | 
 **color** | **optional.String**|  | 
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
 **colorN** | **optional.String**|  | 
 **colorIc** | **optional.String**|  | 
 **colorNic** | **optional.String**|  | 
 **colorIew** | **optional.String**|  | 
 **colorNiew** | **optional.String**|  | 
 **colorIsw** | **optional.String**|  | 
 **colorNisw** | **optional.String**|  | 
 **colorIe** | **optional.String**|  | 
 **colorNie** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20042**](inline_response_200_42.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasTagsPartialUpdate**
> Tag ExtrasTagsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tag. | 
  **data** | [**Tag**](Tag.md)|  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasTagsRead**
> Tag ExtrasTagsRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tag. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtrasTagsUpdate**
> Tag ExtrasTagsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this tag. | 
  **data** | [**Tag**](Tag.md)|  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

