/*
 * v1alpha1/proto/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Note kind that represents a logical attestation \"role\" or \"authority\".  For example, an organization might have one AttestationAuthority for \"QA\" and one for \"build\".  This Note is intended to act strictly as a grouping mechanism for the attached Occurrences (Attestations).  This grouping mechanism also provides a security boundary and provides a single point of lookup to find all attached Attestation Occurrences, even if they don't all live in the same project.
type ApiAttestationAuthority struct {

	Hint *AttestationAuthorityAttestationAuthorityHint `json:"hint,omitempty"`
}