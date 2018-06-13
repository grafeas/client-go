# \GrafeasProjectsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](GrafeasProjectsApi.md#CreateProject) | **Post** /v1alpha1/projects | Creates a new &#x60;Project&#x60;.
[**DeleteProject**](GrafeasProjectsApi.md#DeleteProject) | **Delete** /v1alpha1/{name} | Deletes the given &#x60;Project&#x60; from the system.
[**GetProject**](GrafeasProjectsApi.md#GetProject) | **Get** /v1alpha1/{name} | Returns the requested &#x60;Project&#x60;.
[**ListProjects**](GrafeasProjectsApi.md#ListProjects) | **Get** /v1alpha1/projects | Lists &#x60;Projects&#x60;


# **CreateProject**
> ProtobufEmpty CreateProject(ctx, body)
Creates a new `Project`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApiProject**](ApiProject.md)|  | 

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
Deletes the given `Project` from the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

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
Returns the requested `Project`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

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
Lists `Projects`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| The filter expression. | 
 **pageSize** | **optional.Int32**| Number of projects to return in the list. | 
 **pageToken** | **optional.String**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**ApiListProjectsResponse**](apiListProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

