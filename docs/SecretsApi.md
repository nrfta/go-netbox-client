# \SecretsApi

All URIs are relative to *https://netbox/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretsGenerateRsaKeyPairList**](SecretsApi.md#SecretsGenerateRsaKeyPairList) | **Get** /secrets/generate-rsa-key-pair/ | This endpoint can be used to generate a new RSA key pair. The keys are returned in PEM format.
[**SecretsGetSessionKeyCreate**](SecretsApi.md#SecretsGetSessionKeyCreate) | **Post** /secrets/get-session-key/ | 
[**SecretsSecretRolesCreate**](SecretsApi.md#SecretsSecretRolesCreate) | **Post** /secrets/secret-roles/ | 
[**SecretsSecretRolesDelete**](SecretsApi.md#SecretsSecretRolesDelete) | **Delete** /secrets/secret-roles/{id}/ | 
[**SecretsSecretRolesList**](SecretsApi.md#SecretsSecretRolesList) | **Get** /secrets/secret-roles/ | 
[**SecretsSecretRolesPartialUpdate**](SecretsApi.md#SecretsSecretRolesPartialUpdate) | **Patch** /secrets/secret-roles/{id}/ | 
[**SecretsSecretRolesRead**](SecretsApi.md#SecretsSecretRolesRead) | **Get** /secrets/secret-roles/{id}/ | 
[**SecretsSecretRolesUpdate**](SecretsApi.md#SecretsSecretRolesUpdate) | **Put** /secrets/secret-roles/{id}/ | 
[**SecretsSecretsCreate**](SecretsApi.md#SecretsSecretsCreate) | **Post** /secrets/secrets/ | 
[**SecretsSecretsDelete**](SecretsApi.md#SecretsSecretsDelete) | **Delete** /secrets/secrets/{id}/ | 
[**SecretsSecretsList**](SecretsApi.md#SecretsSecretsList) | **Get** /secrets/secrets/ | 
[**SecretsSecretsPartialUpdate**](SecretsApi.md#SecretsSecretsPartialUpdate) | **Patch** /secrets/secrets/{id}/ | 
[**SecretsSecretsRead**](SecretsApi.md#SecretsSecretsRead) | **Get** /secrets/secrets/{id}/ | 
[**SecretsSecretsUpdate**](SecretsApi.md#SecretsSecretsUpdate) | **Put** /secrets/secrets/{id}/ | 


# **SecretsGenerateRsaKeyPairList**
> SecretsGenerateRsaKeyPairList(ctx, )
This endpoint can be used to generate a new RSA key pair. The keys are returned in PEM format.

{         \"public_key\": \"<public key>\",         \"private_key\": \"<private key>\"     }

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

# **SecretsGetSessionKeyCreate**
> SecretsGetSessionKeyCreate(ctx, )


Retrieve a temporary session key to use for encrypting and decrypting secrets via the API. The user's private RSA key is POSTed with the name `private_key`. An example:      curl -v -X POST -H \"Authorization: Token <token>\" -H \"Accept: application/json; indent=4\" \\     --data-urlencode \"private_key@<filename>\" https://netbox/api/secrets/get-session-key/  This request will yield a base64-encoded session key to be included in an `X-Session-Key` header in future requests:      {         \"session_key\": \"+8t4SI6XikgVmB5+/urhozx9O5qCQANyOk1MNe6taRf=\"     }  This endpoint accepts one optional parameter: `preserve_key`. If True and a session key exists, the existing session key will be returned instead of a new one.

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

# **SecretsSecretRolesCreate**
> SecretRole SecretsSecretRolesCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SecretRole**](SecretRole.md)|  | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretRolesDelete**
> SecretsSecretRolesDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this secret role. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretRolesList**
> InlineResponse20052 SecretsSecretRolesList(ctx, optional)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecretsApiSecretsSecretRolesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiSecretsSecretRolesListOpts struct

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

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretRolesPartialUpdate**
> SecretRole SecretsSecretRolesPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this secret role. | 
  **data** | [**SecretRole**](SecretRole.md)|  | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretRolesRead**
> SecretRole SecretsSecretRolesRead(ctx, id)


Call to super to allow for caching

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this secret role. | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretRolesUpdate**
> SecretRole SecretsSecretRolesUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this secret role. | 
  **data** | [**SecretRole**](SecretRole.md)|  | 

### Return type

[**SecretRole**](SecretRole.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretsCreate**
> Secret SecretsSecretsCreate(ctx, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WritableSecret**](WritableSecret.md)|  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretsDelete**
> SecretsSecretsDelete(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this secret. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretsList**
> InlineResponse20053 SecretsSecretsList(ctx, optional)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecretsApiSecretsSecretsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiSecretsSecretsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.String**|  | 
 **name** | **optional.String**|  | 
 **created** | **optional.String**|  | 
 **createdGte** | **optional.String**|  | 
 **createdLte** | **optional.String**|  | 
 **lastUpdated** | **optional.String**|  | 
 **lastUpdatedGte** | **optional.String**|  | 
 **lastUpdatedLte** | **optional.String**|  | 
 **q** | **optional.String**|  | 
 **roleId** | **optional.String**|  | 
 **role** | **optional.String**|  | 
 **deviceId** | **optional.String**|  | 
 **device** | **optional.String**|  | 
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
 **roleIdN** | **optional.String**|  | 
 **roleN** | **optional.String**|  | 
 **deviceIdN** | **optional.String**|  | 
 **deviceN** | **optional.String**|  | 
 **tagN** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to return per page. | 
 **offset** | **optional.Int32**| The initial index from which to return the results. | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretsPartialUpdate**
> Secret SecretsSecretsPartialUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this secret. | 
  **data** | [**WritableSecret**](WritableSecret.md)|  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretsRead**
> Secret SecretsSecretsRead(ctx, id)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this secret. | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretsSecretsUpdate**
> Secret SecretsSecretsUpdate(ctx, id, data)




### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| A unique integer value identifying this secret. | 
  **data** | [**WritableSecret**](WritableSecret.md)|  | 

### Return type

[**Secret**](Secret.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

