# V1beta1Note

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Output only. The name of the note in the form of &#x60;projects/[PROVIDER_ID]/notes/[NOTE_ID]&#x60;. | [optional] [default to null]
**ShortDescription** | **string** | A one sentence description of this note. | [optional] [default to null]
**LongDescription** | **string** | A detailed description of this note. | [optional] [default to null]
**Kind** | [***V1beta1NoteKind**](v1beta1NoteKind.md) | Output only. The type of analysis. This field can be used as a filter in list requests. | [optional] [default to null]
**RelatedUrl** | [**[]V1beta1RelatedUrl**](v1beta1RelatedUrl.md) | URLs associated with this note. | [optional] [default to null]
**ExpirationTime** | [**time.Time**](time.Time.md) | Time of expiration for this note. Empty if note does not expire. | [optional] [default to null]
**CreateTime** | [**time.Time**](time.Time.md) | Output only. The time this note was created. This field can be used as a filter in list requests. | [optional] [default to null]
**UpdateTime** | [**time.Time**](time.Time.md) | Output only. The time this note was last updated. This field can be used as a filter in list requests. | [optional] [default to null]
**RelatedNoteNames** | **[]string** | Other notes related to this note. | [optional] [default to null]
**Vulnerability** | [***VulnerabilityVulnerability**](vulnerabilityVulnerability.md) | A note describing a package vulnerability. | [optional] [default to null]
**Build** | [***BuildBuild**](buildBuild.md) | A note describing build provenance for a verifiable build. | [optional] [default to null]
**BaseImage** | [***ImageBasis**](imageBasis.md) | A note describing a base image. | [optional] [default to null]
**Package_** | [***PackagePackage**](packagePackage.md) | A note describing a package hosted by various package managers. | [optional] [default to null]
**Deployable** | [***DeploymentDeployable**](deploymentDeployable.md) | A note describing something that can be deployed. | [optional] [default to null]
**Discovery** | [***DiscoveryDiscovery**](discoveryDiscovery.md) | A note describing the initial analysis of a resource. | [optional] [default to null]
**AttestationAuthority** | [***AttestationAuthority**](attestationAuthority.md) | A note describing an attestation role. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


