# PackageVersion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epoch** | **int32** | Used to correct mistakes in the version numbering scheme. | [optional] [default to null]
**Name** | **string** | Required only when version kind is NORMAL. The main part of the version name. | [optional] [default to null]
**Revision** | **string** | The iteration of the package build from the above version. | [optional] [default to null]
**Inclusive** | **bool** | Whether this version is specifying part of an inclusive range. Grafeas does not have the capability to specify version ranges; instead we have fields that specify start version and end versions. At times this is insufficient - we also need to specify whether the version is included in the range or is excluded from the range. This boolean is expected to be set to true when the version is included in a range. | [optional] [default to null]
**Kind** | [***VersionVersionKind**](VersionVersionKind.md) | Required. Distinguishes between sentinel MIN/MAX versions and normal versions. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


