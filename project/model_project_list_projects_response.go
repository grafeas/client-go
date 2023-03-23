/*
 * project.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Response for listing projects.
type ProjectListProjectsResponse struct {
	// The projects requested.
	Projects []ProjectProject `json:"projects,omitempty"`
	// The next pagination token in the list response. It should be used as `page_token` for the following request. An empty value means no more results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}
