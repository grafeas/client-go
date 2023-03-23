# ImageDerived

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | [***ImageFingerprint**](imageFingerprint.md) | Required. The fingerprint of the derived image. | [optional] [default to null]
**Distance** | **int32** | Output only. The number of layers by which this image differs from the associated image basis. | [optional] [default to null]
**LayerInfo** | [**[]ImageLayer**](imageLayer.md) | This contains layer-specific metadata, if populated it has length \&quot;distance\&quot; and is ordered with [distance] being the layer immediately following the base image and [1] being the final layer. | [optional] [default to null]
**BaseResourceUrl** | **string** | Output only. This contains the base image URL for the derived image occurrence. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


