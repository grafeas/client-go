/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Container message for hashes of byte content of files, used in Source messages to verify integrity of source input to the build.
type ApiFileHashes struct {
	// Collection of file hashes.
	FileHash []ApiHash `json:"file_hash,omitempty"`
}