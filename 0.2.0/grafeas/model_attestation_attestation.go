/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Occurrence that represents a single \"attestation\". The authenticity of an attestation can be verified using the attached signature. If the verifier trusts the public key of the signer, then verifying the signature is sufficient to establish trust. In this circumstance, the authority to which this attestation is attached is primarily useful for look-up (how to find this attestation if you already know the authority and artifact to be verified) and intent (which authority was this attestation intended to sign for).
type AttestationAttestation struct {
	// A PGP signed attestation.
	PgpSignedAttestation *AttestationPgpSignedAttestation `json:"pgpSignedAttestation,omitempty"`
	// An attestation that supports multiple `Signature`s over the same attestation payload. The signatures (defined in common.proto) support a superset of public key types and IDs compared to PgpSignedAttestation.
	GenericSignedAttestation *AttestationGenericSignedAttestation `json:"genericSignedAttestation,omitempty"`
}
