# \GrafeasApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNote**](GrafeasApi.md#CreateNote) | **Post** /v1alpha1/{parent}/notes | Creates a new &#x60;Note&#x60;.
[**CreateOccurrence**](GrafeasApi.md#CreateOccurrence) | **Post** /v1alpha1/{parent}/occurrences | Creates a new &#x60;Occurrence&#x60;. Use this method to create &#x60;Occurrences&#x60; for a resource.
[**CreateOperation**](GrafeasApi.md#CreateOperation) | **Post** /v1alpha1/{parent}/operations | Creates a new &#x60;Operation&#x60;.
[**GetOccurrenceNote**](GrafeasApi.md#GetOccurrenceNote) | **Get** /v1alpha1/{name}/notes | Gets the &#x60;Note&#x60; attached to the given &#x60;Occurrence&#x60;.
[**ListNoteOccurrences**](GrafeasApi.md#ListNoteOccurrences) | **Get** /v1alpha1/{name}/occurrences | Lists &#x60;Occurrences&#x60; referencing the specified &#x60;Note&#x60;. Use this method to get all occurrences referencing your &#x60;Note&#x60; across all your customer projects.
[**ListNotes**](GrafeasApi.md#ListNotes) | **Get** /v1alpha1/{parent}/notes | Lists all &#x60;Notes&#x60; for a given project.
[**ListOccurrences**](GrafeasApi.md#ListOccurrences) | **Get** /v1alpha1/{parent}/occurrences | Lists active &#x60;Occurrences&#x60; for a given project matching the filters.
[**UpdateNote**](GrafeasApi.md#UpdateNote) | **Patch** /v1alpha1/{name} | Updates an existing &#x60;Note&#x60;.


# **CreateNote**
> ApiNote CreateNote(ctx, parent, body)
Creates a new `Note`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**|  | 
  **body** | [**ApiNote**](ApiNote.md)|  | 

### Return type

[**ApiNote**](apiNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOccurrence**
> ApiOccurrence CreateOccurrence(ctx, parent, body)
Creates a new `Occurrence`. Use this method to create `Occurrences` for a resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**|  | 
  **body** | [**ApiOccurrence**](ApiOccurrence.md)|  | 

### Return type

[**ApiOccurrence**](apiOccurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOperation**
> LongrunningOperation CreateOperation(ctx, parent, body)
Creates a new `Operation`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**|  | 
  **body** | [**ApiCreateOperationRequest**](ApiCreateOperationRequest.md)|  | 

### Return type

[**LongrunningOperation**](longrunningOperation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOccurrenceNote**
> ApiNote GetOccurrenceNote(ctx, name)
Gets the `Note` attached to the given `Occurrence`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**ApiNote**](apiNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNoteOccurrences**
> ApiListNoteOccurrencesResponse ListNoteOccurrences(ctx, name, optional)
Lists `Occurrences` referencing the specified `Note`. Use this method to get all occurrences referencing your `Note` across all your customer projects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
 **optional** | ***ListNoteOccurrencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListNoteOccurrencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter expression. | 
 **pageSize** | **optional.Int32**| Number of notes to return in the list. | 
 **pageToken** | **optional.String**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**ApiListNoteOccurrencesResponse**](apiListNoteOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNotes**
> ApiListNotesResponse ListNotes(ctx, parent, optional)
Lists all `Notes` for a given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**|  | 
 **optional** | ***ListNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListNotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter expression. | 
 **pageSize** | **optional.Int32**| Number of notes to return in the list. | 
 **pageToken** | **optional.String**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**ApiListNotesResponse**](apiListNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOccurrences**
> ApiListOccurrencesResponse ListOccurrences(ctx, parent, optional)
Lists active `Occurrences` for a given project matching the filters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**|  | 
 **optional** | ***ListOccurrencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOccurrencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter expression. | 
 **pageSize** | **optional.Int32**| Number of occurrences to return in the list. | 
 **pageToken** | **optional.String**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**ApiListOccurrencesResponse**](apiListOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNote**
> ApiNote UpdateNote(ctx, name, body)
Updates an existing `Note`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 
  **body** | [**ApiNote**](ApiNote.md)|  | 

### Return type

[**ApiNote**](apiNote.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

