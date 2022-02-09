# \GrafeasV1Beta1Api

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GrafeasV1Beta1BatchCreateNotes**](GrafeasV1Beta1Api.md#GrafeasV1Beta1BatchCreateNotes) | **Post** /v1beta1/{parent}/notes:batchCreate | Creates new notes in batch.
[**GrafeasV1Beta1BatchCreateOccurrences**](GrafeasV1Beta1Api.md#GrafeasV1Beta1BatchCreateOccurrences) | **Post** /v1beta1/{parent}/occurrences:batchCreate | Creates new occurrences in batch.
[**GrafeasV1Beta1CreateNote**](GrafeasV1Beta1Api.md#GrafeasV1Beta1CreateNote) | **Post** /v1beta1/{parent}/notes | Creates a new note.
[**GrafeasV1Beta1CreateOccurrence**](GrafeasV1Beta1Api.md#GrafeasV1Beta1CreateOccurrence) | **Post** /v1beta1/{parent}/occurrences | Creates a new occurrence.
[**GrafeasV1Beta1DeleteNote**](GrafeasV1Beta1Api.md#GrafeasV1Beta1DeleteNote) | **Delete** /v1beta1/{name_1} | Deletes the specified note.
[**GrafeasV1Beta1DeleteOccurrence**](GrafeasV1Beta1Api.md#GrafeasV1Beta1DeleteOccurrence) | **Delete** /v1beta1/{name} | Deletes the specified occurrence. For example, use this method to delete an occurrence when the occurrence is no longer applicable for the given resource.
[**GrafeasV1Beta1GetNote**](GrafeasV1Beta1Api.md#GrafeasV1Beta1GetNote) | **Get** /v1beta1/{name_1} | Gets the specified note.
[**GrafeasV1Beta1GetOccurrence**](GrafeasV1Beta1Api.md#GrafeasV1Beta1GetOccurrence) | **Get** /v1beta1/{name} | Gets the specified occurrence.
[**GrafeasV1Beta1GetOccurrenceNote**](GrafeasV1Beta1Api.md#GrafeasV1Beta1GetOccurrenceNote) | **Get** /v1beta1/{name}/notes | Gets the note attached to the specified occurrence. Consumer projects can use this method to get a note that belongs to a provider project.
[**GrafeasV1Beta1GetVulnerabilityOccurrencesSummary**](GrafeasV1Beta1Api.md#GrafeasV1Beta1GetVulnerabilityOccurrencesSummary) | **Get** /v1beta1/{parent}/occurrences:vulnerabilitySummary | Gets a summary of the number and severity of occurrences.
[**GrafeasV1Beta1ListNoteOccurrences**](GrafeasV1Beta1Api.md#GrafeasV1Beta1ListNoteOccurrences) | **Get** /v1beta1/{name}/occurrences | Lists occurrences referencing the specified note. Provider projects can use this method to get all occurrences across consumer projects referencing the specified note.
[**GrafeasV1Beta1ListNotes**](GrafeasV1Beta1Api.md#GrafeasV1Beta1ListNotes) | **Get** /v1beta1/{parent}/notes | Lists notes for the specified project.
[**GrafeasV1Beta1ListOccurrences**](GrafeasV1Beta1Api.md#GrafeasV1Beta1ListOccurrences) | **Get** /v1beta1/{parent}/occurrences | Lists occurrences for the specified project.
[**GrafeasV1Beta1UpdateNote**](GrafeasV1Beta1Api.md#GrafeasV1Beta1UpdateNote) | **Patch** /v1beta1/{name_1} | Updates the specified note.
[**GrafeasV1Beta1UpdateOccurrence**](GrafeasV1Beta1Api.md#GrafeasV1Beta1UpdateOccurrence) | **Patch** /v1beta1/{name} | Updates the specified occurrence.


# **GrafeasV1Beta1BatchCreateNotes**
> V1beta1BatchCreateNotesResponse GrafeasV1Beta1BatchCreateNotes(ctx, parent, body)
Creates new notes in batch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**| The name of the project in the form of &#x60;projects/[PROJECT_ID]&#x60;, under which the notes are to be created. | 
  **body** | [**Body**](Body.md)|  | 

### Return type

[**V1beta1BatchCreateNotesResponse**](v1beta1BatchCreateNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1BatchCreateOccurrences**
> V1beta1BatchCreateOccurrencesResponse GrafeasV1Beta1BatchCreateOccurrences(ctx, parent, body)
Creates new occurrences in batch.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**| The name of the project in the form of &#x60;projects/[PROJECT_ID]&#x60;, under which the occurrences are to be created. | 
  **body** | [**Body1**](Body1.md)|  | 

### Return type

[**V1beta1BatchCreateOccurrencesResponse**](v1beta1BatchCreateOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1CreateNote**
> V1beta1Note GrafeasV1Beta1CreateNote(ctx, parent, body, noteId)
Creates a new note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**| The name of the project in the form of &#x60;projects/[PROJECT_ID]&#x60;, under which the note is to be created. | 
  **body** | [**V1beta1Note**](V1beta1Note.md)| The note to create. | 
  **noteId** | **string**| The ID to use for this note. | 

### Return type

[**V1beta1Note**](v1beta1Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1CreateOccurrence**
> V1beta1Occurrence GrafeasV1Beta1CreateOccurrence(ctx, parent, body)
Creates a new occurrence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**| The name of the project in the form of &#x60;projects/[PROJECT_ID]&#x60;, under which the occurrence is to be created. | 
  **body** | [**V1beta1Occurrence**](V1beta1Occurrence.md)| The occurrence to create. | 

### Return type

[**V1beta1Occurrence**](v1beta1Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1DeleteNote**
> interface{} GrafeasV1Beta1DeleteNote(ctx, name1)
Deletes the specified note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name1** | **string**| The name of the note in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1DeleteOccurrence**
> interface{} GrafeasV1Beta1DeleteOccurrence(ctx, name)
Deletes the specified occurrence. For example, use this method to delete an occurrence when the occurrence is no longer applicable for the given resource.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the occurrence in the form of &#x60;projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]&#x60;. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1GetNote**
> V1beta1Note GrafeasV1Beta1GetNote(ctx, name1)
Gets the specified note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name1** | **string**| The name of the note in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 

### Return type

[**V1beta1Note**](v1beta1Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1GetOccurrence**
> V1beta1Occurrence GrafeasV1Beta1GetOccurrence(ctx, name)
Gets the specified occurrence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the occurrence in the form of &#x60;projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]&#x60;. | 

### Return type

[**V1beta1Occurrence**](v1beta1Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1GetOccurrenceNote**
> V1beta1Note GrafeasV1Beta1GetOccurrenceNote(ctx, name)
Gets the note attached to the specified occurrence. Consumer projects can use this method to get a note that belongs to a provider project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the occurrence in the form of &#x60;projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]&#x60;. | 

### Return type

[**V1beta1Note**](v1beta1Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1GetVulnerabilityOccurrencesSummary**
> V1beta1VulnerabilityOccurrencesSummary GrafeasV1Beta1GetVulnerabilityOccurrencesSummary(ctx, parent, optional)
Gets a summary of the number and severity of occurrences.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**| The name of the project to get a vulnerability summary for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **optional** | ***GrafeasV1Beta1GetVulnerabilityOccurrencesSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GrafeasV1Beta1GetVulnerabilityOccurrencesSummaryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter expression. | 

### Return type

[**V1beta1VulnerabilityOccurrencesSummary**](v1beta1VulnerabilityOccurrencesSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1ListNoteOccurrences**
> V1beta1ListNoteOccurrencesResponse GrafeasV1Beta1ListNoteOccurrences(ctx, name, optional)
Lists occurrences referencing the specified note. Provider projects can use this method to get all occurrences across consumer projects referencing the specified note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the note to list occurrences for in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 
 **optional** | ***GrafeasV1Beta1ListNoteOccurrencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GrafeasV1Beta1ListNoteOccurrencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter expression. | 
 **pageSize** | **optional.Int32**| Number of occurrences to return in the list. | 
 **pageToken** | **optional.String**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**V1beta1ListNoteOccurrencesResponse**](v1beta1ListNoteOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1ListNotes**
> V1beta1ListNotesResponse GrafeasV1Beta1ListNotes(ctx, parent, optional)
Lists notes for the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**| The name of the project to list notes for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **optional** | ***GrafeasV1Beta1ListNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GrafeasV1Beta1ListNotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter expression. | 
 **pageSize** | **optional.Int32**| Number of notes to return in the list. Must be positive. Max allowed page size is 1000. If not specified, page size defaults to 20. | 
 **pageToken** | **optional.String**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**V1beta1ListNotesResponse**](v1beta1ListNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1ListOccurrences**
> V1beta1ListOccurrencesResponse GrafeasV1Beta1ListOccurrences(ctx, parent, optional)
Lists occurrences for the specified project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **parent** | **string**| The name of the project to list occurrences for in the form of &#x60;projects/[PROJECT_ID]&#x60;. | 
 **optional** | ***GrafeasV1Beta1ListOccurrencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GrafeasV1Beta1ListOccurrencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **optional.String**| The filter expression. | 
 **pageSize** | **optional.Int32**| Number of occurrences to return in the list. Must be positive. Max allowed page size is 1000. If not specified, page size defaults to 20. | 
 **pageToken** | **optional.String**| Token to provide to skip to a particular spot in the list. | 

### Return type

[**V1beta1ListOccurrencesResponse**](v1beta1ListOccurrencesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1UpdateNote**
> V1beta1Note GrafeasV1Beta1UpdateNote(ctx, name1, body, optional)
Updates the specified note.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name1** | **string**| The name of the note in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | 
  **body** | [**V1beta1Note**](V1beta1Note.md)| The updated note. | 
 **optional** | ***GrafeasV1Beta1UpdateNoteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GrafeasV1Beta1UpdateNoteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **optional.String**| The fields to update. | 

### Return type

[**V1beta1Note**](v1beta1Note.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GrafeasV1Beta1UpdateOccurrence**
> V1beta1Occurrence GrafeasV1Beta1UpdateOccurrence(ctx, name, body, optional)
Updates the specified occurrence.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the occurrence in the form of &#x60;projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]&#x60;. | 
  **body** | [**V1beta1Occurrence**](V1beta1Occurrence.md)| The updated occurrence. | 
 **optional** | ***GrafeasV1Beta1UpdateOccurrenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GrafeasV1Beta1UpdateOccurrenceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateMask** | **optional.String**| The fields to update. | 

### Return type

[**V1beta1Occurrence**](v1beta1Occurrence.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

