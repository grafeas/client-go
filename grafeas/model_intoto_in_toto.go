/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// This contains the fields corresponding to the definition of a software supply chain step in an in-toto layout. This information goes into a Grafeas note.
type IntotoInToto struct {
	// This field identifies the name of the step in the supply chain.
	StepName string `json:"stepName,omitempty"`
	// This field contains the public keys that can be used to verify the signatures on the step metadata.
	SigningKeys []IntotoSigningKey `json:"signingKeys,omitempty"`
	// The following fields contain in-toto artifact rules identifying the artifacts that enter this supply chain step, and exit the supply chain step, i.e. materials and products of the step.
	ExpectedMaterials []InTotoArtifactRule `json:"expectedMaterials,omitempty"`
	ExpectedProducts []InTotoArtifactRule `json:"expectedProducts,omitempty"`
	// This field contains the expected command used to perform the step.
	ExpectedCommand []string `json:"expectedCommand,omitempty"`
	// This field contains a value that indicates the minimum number of keys that need to be used to sign the step's in-toto link.
	Threshold string `json:"threshold,omitempty"`
}
