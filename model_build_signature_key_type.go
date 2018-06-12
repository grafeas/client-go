/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// BuildSignatureKeyType : - KEY_TYPE_UNSPECIFIED: `KeyType` is not set.  - PGP_ASCII_ARMORED: `PGP ASCII Armored` public key.  - PKIX_PEM: `PKIX PEM` public key.
type BuildSignatureKeyType string

// List of BuildSignatureKeyType
const (
	KEY_TYPE_UNSPECIFIED BuildSignatureKeyType = "KEY_TYPE_UNSPECIFIED"

	PGP_ASCII_ARMORED BuildSignatureKeyType = "PGP_ASCII_ARMORED"

	PKIX_PEM BuildSignatureKeyType = "PKIX_PEM"
)