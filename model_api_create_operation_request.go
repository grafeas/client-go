/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

type ApiCreateOperationRequest struct {
	// The projectId that this operation should be created under.
	Parent string `json:"parent,omitempty"`

	// The ID to use for this operation.
	OperationId string `json:"operation_id,omitempty"`

	// The operation to create.
	Operation *LongrunningOperation `json:"operation,omitempty"`
}