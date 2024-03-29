/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Container message for hashes of byte content of files, used in source messages to verify integrity of source input to the build.
type ProvenanceFileHashes struct {
	// Required. Collection of file hashes.
	FileHash []ProvenanceHash `json:"fileHash,omitempty"`
}
