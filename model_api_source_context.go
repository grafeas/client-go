/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// A SourceContext is a reference to a tree of files. A SourceContext together with a path point to a unique revision of a single file or directory.
type ApiSourceContext struct {
	// A SourceContext referring to a revision in a Google Cloud Source Repo.
	CloudRepo *ApiCloudRepoSourceContext `json:"cloud_repo,omitempty"`

	// A SourceContext referring to a Gerrit project.
	Gerrit *ApiGerritSourceContext `json:"gerrit,omitempty"`

	// A SourceContext referring to any third party Git repo (e.g., GitHub).
	Git *ApiGitSourceContext `json:"git,omitempty"`

	// Labels with user defined metadata.
	Labels map[string]string `json:"labels,omitempty"`
}