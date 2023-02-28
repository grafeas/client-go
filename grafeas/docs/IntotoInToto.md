# IntotoInToto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StepName** | **string** | This field identifies the name of the step in the supply chain. | [optional] [default to null]
**SigningKeys** | [**[]IntotoSigningKey**](intotoSigningKey.md) | This field contains the public keys that can be used to verify the signatures on the step metadata. | [optional] [default to null]
**ExpectedMaterials** | [**[]InTotoArtifactRule**](InTotoArtifactRule.md) | The following fields contain in-toto artifact rules identifying the artifacts that enter this supply chain step, and exit the supply chain step, i.e. materials and products of the step. | [optional] [default to null]
**ExpectedProducts** | [**[]InTotoArtifactRule**](InTotoArtifactRule.md) |  | [optional] [default to null]
**ExpectedCommand** | **[]string** | This field contains the expected command used to perform the step. | [optional] [default to null]
**Threshold** | **string** | This field contains a value that indicates the minimum number of keys that need to be used to sign the step&#39;s in-toto link. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


