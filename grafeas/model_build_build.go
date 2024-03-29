/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Note holding the version of the provider's builder and the signature of the provenance message in the build details occurrence.
type BuildBuild struct {
	// Required. Immutable. Version of the builder which produced this build.
	BuilderVersion string `json:"builderVersion,omitempty"`
	// Signature of the build in occurrences pointing to this build note containing build details.
	Signature *BuildBuildSignature `json:"signature,omitempty"`
}
