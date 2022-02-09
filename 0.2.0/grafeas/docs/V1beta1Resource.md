# V1beta1Resource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Deprecated, do not use. Use uri instead.  The name of the resource. For example, the name of a Docker image - \&quot;Debian\&quot;. | [optional] [default to null]
**Uri** | **string** | Required. The unique URI of the resource. For example, &#x60;https://gcr.io/project/image@sha256:foo&#x60; for a Docker image. | [optional] [default to null]
**ContentHash** | [***ProvenanceHash**](provenanceHash.md) | Deprecated, do not use. Use uri instead.  The hash of the resource content. For example, the Docker digest. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


