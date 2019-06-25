# V1beta1Occurrence

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Output only. The name of the occurrence in the form of &#x60;projects/[PROJECT_ID]/occurrences/[OCCURRENCE_ID]&#x60;. | [optional] [default to null]
**Resource** | [***V1beta1Resource**](v1beta1Resource.md) | Required. Immutable. The resource for which the occurrence applies. | [optional] [default to null]
**NoteName** | **string** | Required. Immutable. The analysis note associated with this occurrence, in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. This field can be used as a filter in list requests. | [optional] [default to null]
**Kind** | [***V1beta1NoteKind**](v1beta1NoteKind.md) | Output only. This explicitly denotes which of the occurrence details are specified. This field can be used as a filter in list requests. | [optional] [default to null]
**Remediation** | **string** | A description of actions that can be taken to remedy the note. | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | Output only. The time this occurrence was created. | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | Output only. The time this occurrence was last updated. | [optional] [default to null]
**Vulnerability** | [***V1beta1vulnerabilityDetails**](v1beta1vulnerabilityDetails.md) | Describes a security vulnerability. | [optional] [default to null]
**Build** | [***V1beta1buildDetails**](v1beta1buildDetails.md) | Describes a verifiable build. | [optional] [default to null]
**DerivedImage** | [***V1beta1imageDetails**](v1beta1imageDetails.md) | Describes how this resource derives from the basis in the associated note. | [optional] [default to null]
**Installation** | [***V1beta1packageDetails**](v1beta1packageDetails.md) | Describes the installation of a package on the linked resource. | [optional] [default to null]
**Deployment** | [***V1beta1deploymentDetails**](v1beta1deploymentDetails.md) | Describes the deployment of an artifact on a runtime. | [optional] [default to null]
**Discovered** | [***V1beta1discoveryDetails**](v1beta1discoveryDetails.md) | Describes when a resource was discovered. | [optional] [default to null]
**Attestation** | [***V1beta1attestationDetails**](v1beta1attestationDetails.md) | Describes an attestation of an artifact. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


