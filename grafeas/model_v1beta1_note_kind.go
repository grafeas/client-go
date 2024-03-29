/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// V1beta1NoteKind : Kind represents the kinds of notes supported.   - NOTE_KIND_UNSPECIFIED: Default value. This value is unused.  - VULNERABILITY: The note and occurrence represent a package vulnerability.  - BUILD: The note and occurrence assert build provenance.  - IMAGE: This represents an image basis relationship.  - PACKAGE: This represents a package installed via a package manager.  - DEPLOYMENT: The note and occurrence track deployment events.  - DISCOVERY: The note and occurrence track the initial discovery status of a resource.  - ATTESTATION: This represents a logical \"role\" that can attest to artifacts.  - INTOTO: This represents an in-toto link.  - SBOM: This represents a software bill of materials.  - SPDX_PACKAGE: This represents an SPDX Package.  - SPDX_FILE: This represents an SPDX File.  - SPDX_RELATIONSHIP: This represents an SPDX Relationship.  - VULNERABILITY_ASSESSMENT: This represents a Vulnerability Assessment.
type V1beta1NoteKind string

// List of v1beta1NoteKind
const (
	NOTE_KIND_UNSPECIFIED_V1beta1NoteKind V1beta1NoteKind = "NOTE_KIND_UNSPECIFIED"
	VULNERABILITY_V1beta1NoteKind V1beta1NoteKind = "VULNERABILITY"
	BUILD_V1beta1NoteKind V1beta1NoteKind = "BUILD"
	IMAGE_V1beta1NoteKind V1beta1NoteKind = "IMAGE"
	PACKAGE__V1beta1NoteKind V1beta1NoteKind = "PACKAGE"
	DEPLOYMENT_V1beta1NoteKind V1beta1NoteKind = "DEPLOYMENT"
	DISCOVERY_V1beta1NoteKind V1beta1NoteKind = "DISCOVERY"
	ATTESTATION_V1beta1NoteKind V1beta1NoteKind = "ATTESTATION"
	INTOTO_V1beta1NoteKind V1beta1NoteKind = "INTOTO"
	SBOM_V1beta1NoteKind V1beta1NoteKind = "SBOM"
	SPDX_PACKAGE_V1beta1NoteKind V1beta1NoteKind = "SPDX_PACKAGE"
	SPDX_FILE_V1beta1NoteKind V1beta1NoteKind = "SPDX_FILE"
	SPDX_RELATIONSHIP_V1beta1NoteKind V1beta1NoteKind = "SPDX_RELATIONSHIP"
	VULNERABILITY_ASSESSMENT_V1beta1NoteKind V1beta1NoteKind = "VULNERABILITY_ASSESSMENT"
)
