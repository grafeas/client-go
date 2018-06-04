/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"time"
)

// Provides a detailed description of a `Note`.
type ApiNote struct {
	Name string `json:"name,omitempty"`

	// A one sentence description of this `Note`.
	ShortDescription string `json:"short_description,omitempty"`

	// A detailed description of this `Note`.
	LongDescription string `json:"long_description,omitempty"`

	// Output only. This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
	Kind *ApiNoteKind `json:"kind,omitempty"`

	// A package vulnerability type of note.
	VulnerabilityType *ApiVulnerabilityType `json:"vulnerability_type,omitempty"`

	// Build provenance type for a verifiable build.
	BuildType *ApiBuildType `json:"build_type,omitempty"`

	// A note describing a base image.
	BaseImage *DockerImageBasis `json:"base_image,omitempty"`

	// A note describing a package hosted by various package managers.
	Package_ *PackageManagerPackage `json:"package,omitempty"`

	// A note describing something that can be deployed.
	Deployable *ApiDeployable `json:"deployable,omitempty"`

	// A note describing a provider/analysis type.
	Discovery *ApiDiscovery `json:"discovery,omitempty"`

	RelatedUrl []NoteRelatedUrl `json:"related_url,omitempty"`

	// Time of expiration for this note, null if note does not expire.
	ExpirationTime time.Time `json:"expiration_time,omitempty"`

	// Output only. The time this note was created. This field can be used as a filter in list requests.
	CreateTime time.Time `json:"create_time,omitempty"`

	// Output only. The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime time.Time `json:"update_time,omitempty"`
}