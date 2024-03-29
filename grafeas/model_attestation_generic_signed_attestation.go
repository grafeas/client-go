/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// An attestation wrapper that uses the Grafeas `Signature` message. This attestation must define the `serialized_payload` that the `signatures` verify and any metadata necessary to interpret that plaintext.  The signatures should always be over the `serialized_payload` bytestring.
type AttestationGenericSignedAttestation struct {
	// Type (for example schema) of the attestation payload that was signed. The verifier must ensure that the provided type is one that the verifier supports, and that the attestation payload is a valid instantiation of that type (for example by validating a JSON schema).
	ContentType *AttestationGenericSignedAttestationContentType `json:"contentType,omitempty"`
	// The serialized payload that is verified by one or more `signatures`. The encoding and semantic meaning of this payload must match what is set in `content_type`.
	SerializedPayload string `json:"serializedPayload,omitempty"`
	// One or more signatures over `serialized_payload`.  Verifier implementations should consider this attestation message verified if at least one `signature` verifies `serialized_payload`.  See `Signature` in common.proto for more details on signature structure and verification.
	Signatures []Grafeasv1beta1Signature `json:"signatures,omitempty"`
}
