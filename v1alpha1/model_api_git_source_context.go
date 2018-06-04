/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// A GitSourceContext denotes a particular revision in a third party Git repository (e.g., GitHub).
type ApiGitSourceContext struct {
	// Git repository URL.
	Url string `json:"url,omitempty"`

	// Required. Git commit hash.
	RevisionId string `json:"revision_id,omitempty"`
}
