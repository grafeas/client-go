# \ProjectsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsCreateProject**](ProjectsApi.md#ProjectsCreateProject) | **Post** /v1beta1/projects | Creates a new project.
[**ProjectsDeleteProject**](ProjectsApi.md#ProjectsDeleteProject) | **Delete** /v1beta1/{name} | Deletes the specified project.
[**ProjectsGetProject**](ProjectsApi.md#ProjectsGetProject) | **Get** /v1beta1/{name} | Gets the specified project.
[**ProjectsListProjects**](ProjectsApi.md#ProjectsListProjects) | **Get** /v1beta1/projects | Lists projects.


# **ProjectsCreateProject**
> ProjectProject ProjectsCreateProject(ctx, body)
Creates a new project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectProject**](ProjectProject.md)| The project to create. | 

### Return type

[**ProjectProject**](projectProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsDeleteProject**
> interface{} ProjectsDeleteProject(ctx, name)
Deletes the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the project in the form of &#x60;projects/{PROJECT_ID}&#x60;. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsGetProject**
> ProjectProject ProjectsGetProject(ctx, name)
Gets the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the project in the form of &#x60;projects/{PROJECT_ID}&#x60;. | 

### Return type

[**ProjectProject**](projectProject.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsListProjects**
> ProjectListProjectsResponse ProjectsListProjects(ctx, optional)
Lists projects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsListProjectsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsListProjectsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| The filter expression. | 
 **pageSize** | **optional.Int32**| Number of projects to return in the list. | 
 **pageToken** | **optional.String**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**ProjectListProjectsResponse**](projectListProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

