# DetailsVexAssessment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cve** | **string** | Holds the MITRE standard Common Vulnerabilities and Exposures (CVE) tracking number for the vulnerability. | [optional] [default to null]
**RelatedUris** | [**[]V1beta1RelatedUrl**](v1beta1RelatedUrl.md) | Holds a list of references associated with this vulnerability item and assessment. | [optional] [default to null]
**NoteName** | **string** |  | [optional] [default to null]
**State** | [***AssessmentState**](AssessmentState.md) | Provides the state of this Vulnerability assessment. | [optional] [default to null]
**Impacts** | **[]string** | Contains information about the impact of this vulnerability, this will change with time. | [optional] [default to null]
**Remediations** | [**[]AssessmentRemediation**](AssessmentRemediation.md) | Specifies details on how to handle (and presumably, fix) a vulnerability. | [optional] [default to null]
**Justification** | [***AssessmentJustification**](AssessmentJustification.md) | Justification provides the justification when the state of the assessment if NOT_AFFECTED. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


