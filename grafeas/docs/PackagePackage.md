# PackagePackage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the package. | [default to null]
**Distribution** | [**[]PackageDistribution**](packageDistribution.md) | The various channels by which a package is distributed. | [optional] [default to null]
**PackageType** | **string** | The type of package; whether native or non native (e.g., ruby gems, node.js packages, etc.). | [optional] [default to null]
**CpeUri** | **string** | The cpe_uri in [CPE format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package. The cpe_uri will be blank for language packages. | [optional] [default to null]
**Architecture** | [***PackageArchitecture**](packageArchitecture.md) | The CPU architecture for which packages in this distribution channel were built. Architecture will be blank for language packages. | [optional] [default to null]
**Version** | [***PackageVersion**](packageVersion.md) | The version of the package. | [optional] [default to null]
**Maintainer** | **string** | A freeform text denoting the maintainer of this package. | [optional] [default to null]
**Url** | **string** | The homepage for this package. | [optional] [default to null]
**Description** | **string** | The description of this package. | [optional] [default to null]
**License** | [***V1beta1License**](v1beta1License.md) | Licenses that have been declared by the authors of the package. | [optional] [default to null]
**Digest** | [**[]V1beta1Digest**](v1beta1Digest.md) | Hash value, typically a file digest, that allows unique identification a specific package. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


