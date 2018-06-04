/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// ApiNoteKind : This must be 1:1 with members of our oneofs, it can be used for filtering Note and Occurrence on their kind.   - KIND_UNSPECIFIED: Unknown  - PACKAGE_VULNERABILITY: The note and occurrence represent a package vulnerability.  - BUILD_DETAILS: The note and occurrence assert build provenance.  - IMAGE_BASIS: This represents an image basis relationship.  - PACKAGE_MANAGER: This represents a package installed via a package manager.  - DEPLOYABLE: The note and occurrence track deployment events.  - DISCOVERY: The note and occurrence track the initial discovery status of a resource.  - ATTESTATION_AUTHORITY: This represents a logical \"role\" that can attest to artifacts.
type ApiNoteKind string

// List of apiNoteKind
const (
	KIND_UNSPECIFIED ApiNoteKind = "KIND_UNSPECIFIED"

	PACKAGE_VULNERABILITY ApiNoteKind = "PACKAGE_VULNERABILITY"

	BUILD_DETAILS ApiNoteKind = "BUILD_DETAILS"

	IMAGE_BASIS ApiNoteKind = "IMAGE_BASIS"

	PACKAGE_MANAGER ApiNoteKind = "PACKAGE_MANAGER"

	DEPLOYABLE ApiNoteKind = "DEPLOYABLE"

	DISCOVERY ApiNoteKind = "DISCOVERY"

	ATTESTATION_AUTHORITY ApiNoteKind = "ATTESTATION_AUTHORITY"
)
