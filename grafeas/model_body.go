/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Request to create notes in batch.
type Body struct {
	// The notes to create, the key is expected to be the note ID. Max allowed length is 1000.
	Notes map[string]V1beta1Note `json:"notes"`
}