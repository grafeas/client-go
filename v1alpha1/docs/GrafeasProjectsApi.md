# \GrafeasProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](GrafeasProjectsApi.md#CreateProject) | **Post** /v1alpha1/projects | Returns the requested &#x60;Occurrence&#x60;.
[**DeleteProject**](GrafeasProjectsApi.md#DeleteProject) | **Delete** /v1alpha1/{name} | Creates a new &#x60;Occurrence&#x60;. Use this method to create &#x60;Occurrences&#x60; for a resource.
[**GetProject**](GrafeasProjectsApi.md#GetProject) | **Get** /v1alpha1/{name} | Lists active &#x60;Occurrences&#x60; for a given project matching the filters.
[**ListProjects**](GrafeasProjectsApi.md#ListProjects) | **Get** /v1alpha1/projects | Deletes the given &#x60;Occurrence&#x60; from the system. Use this when an &#x60;Occurrence&#x60; is no longer applicable for the given resource.


# **CreateProject**
> ProtobufEmpty CreateProject(ctx, body)
Returns the requested `Occurrence`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ApiProject**](ApiProject.md)| The project to be inserted | 

### Return type

[**ProtobufEmpty**](protobufEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProject**
> ProtobufEmpty DeleteProject(ctx, name)
Creates a new `Occurrence`. Use this method to create `Occurrences` for a resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the project of the form \&quot;projects/{project_id}\&quot; | 

### Return type

[**ProtobufEmpty**](protobufEmpty.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProject**
> ApiProject GetProject(ctx, name)
Lists active `Occurrences` for a given project matching the filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| The name of the project of the form \&quot;projects/{project_id}\&quot; | 

### Return type

[**ApiProject**](apiProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProjects**
> ApiListProjectsResponse ListProjects(ctx, optional)
Deletes the given `Occurrence` from the system. Use this when an `Occurrence` is no longer applicable for the given resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string**| The filter expression. | 
 **pageSize** | **int32**| Number of projects to return in the list. | 
 **pageToken** | **string**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**ApiListProjectsResponse**](apiListProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

