/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

type CvssScope string

// List of CVSSScope
const (
	UNSPECIFIED_CvssScope CvssScope = "SCOPE_UNSPECIFIED"
	UNCHANGED_CvssScope CvssScope = "SCOPE_UNCHANGED"
	CHANGED_CvssScope CvssScope = "SCOPE_CHANGED"
)
