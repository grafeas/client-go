/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Response including listed active occurrences.
type ApiListOccurrencesResponse struct {
	// The occurrences requested.
	Occurrences []ApiOccurrence `json:"occurrences,omitempty"`

	// The next pagination token in the list response. It should be used as `page_token` for the following request. An empty value means no more results.
	NextPageToken string `json:"next_page_token,omitempty"`
}
