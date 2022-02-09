/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// A unique identifier for a Cloud Repo.
type SourceRepoId struct {
	// A combination of a project ID and a repo name.
	ProjectRepoId *SourceProjectRepoId `json:"projectRepoId,omitempty"`
	// A server-assigned, globally unique identifier.
	Uid string `json:"uid,omitempty"`
}
