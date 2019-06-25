/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Response for listing occurrences for a note.
type V1beta1ListNoteOccurrencesResponse struct {
	// The occurrences attached to the specified note.
	Occurrences []V1beta1Occurrence `json:"occurrences,omitempty"`
	// Token to provide to skip to a particular spot in the list.
	NextPageToken string `json:"next_page_token,omitempty"`
}