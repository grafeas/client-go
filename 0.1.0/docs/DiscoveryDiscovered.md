# DiscoveryDiscovered

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousAnalysis** | [***DiscoveredContinuousAnalysis**](DiscoveredContinuousAnalysis.md) | Whether the resource is continuously analyzed. | [optional] [default to null]
**LastAnalysisTime** | [**time.Time**](time.Time.md) | The last time continuous analysis was done for this resource. Deprecated, do not use. | [optional] [default to null]
**AnalysisStatus** | [***DiscoveredAnalysisStatus**](DiscoveredAnalysisStatus.md) | The status of discovery for the resource. | [optional] [default to null]
**AnalysisStatusError** | [***RpcStatus**](rpcStatus.md) | When an error is encountered this will contain a LocalizedMessage under details to show to the user. The LocalizedMessage is output only and populated by the API. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


