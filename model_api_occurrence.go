/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

import (
	"time"
)

// `Occurrence` includes information about analysis occurrences for an image.
type ApiOccurrence struct {
	Name string `json:"name,omitempty"`

	// The unique URL of the image or the container for which the `Occurrence` applies. For example, https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
	ResourceUrl string `json:"resource_url,omitempty"`

	// An analysis note associated with this image, in the form \"providers/{provider_id}/notes/{NOTE_ID}\" This field can be used as a filter in list requests.
	NoteName string `json:"note_name,omitempty"`

	// Output only. This explicitly denotes which of the `Occurrence` details are specified. This field can be used as a filter in list requests.
	Kind *ApiNoteKind `json:"kind,omitempty"`

	// Details of a security vulnerability note.
	VulnerabilityDetails *VulnerabilityTypeVulnerabilityDetails `json:"vulnerability_details,omitempty"`

	// Build details for a verifiable build.
	BuildDetails *ApiBuildDetails `json:"build_details,omitempty"`

	// Describes how this resource derives from the basis in the associated note.
	DerivedImageDetails *DockerImageDerivedDetails `json:"derived_image_details,omitempty"`

	// Describes the installation of a package on the linked resource.
	InstallationDetails *PackageManagerInstallationDetails `json:"installation_details,omitempty"`

	// Describes the deployment of an artifact on a runtime.
	DeploymentDetails *DeployableDeploymentDetails `json:"deployment_details,omitempty"`

	// Describes the initial scan status for this resource.
	DiscoveredDetails *DiscoveryDiscoveredDetails `json:"discovered_details,omitempty"`

	// Describes an attestation of an artifact.
	AttestationDetails *AttestationAuthorityAttestationDetails `json:"attestation_details,omitempty"`

	Remediation string `json:"remediation,omitempty"`

	// Output only. The time this `Occurrence` was created.
	CreateTime time.Time `json:"create_time,omitempty"`

	// Output only. The time this `Occurrence` was last updated.
	UpdateTime time.Time `json:"update_time,omitempty"`

	OperationName string `json:"operation_name,omitempty"`
}
