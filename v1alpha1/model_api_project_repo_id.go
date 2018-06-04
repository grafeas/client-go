/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Selects a repo using a Google Cloud Platform project ID (e.g., winged-cargo-31) and a repo name within that project.
type ApiProjectRepoId struct {
	// The ID of the project.
	ProjectId string `json:"project_id,omitempty"`

	// The name of the repo. Leave empty for the default repo.
	RepoName string `json:"repo_name,omitempty"`
}