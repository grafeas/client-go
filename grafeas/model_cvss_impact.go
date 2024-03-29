/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

type CvssImpact string

// List of CVSSImpact
const (
	UNSPECIFIED_CvssImpact CvssImpact = "IMPACT_UNSPECIFIED"
	HIGH_CvssImpact CvssImpact = "IMPACT_HIGH"
	LOW_CvssImpact CvssImpact = "IMPACT_LOW"
	NONE_CvssImpact CvssImpact = "IMPACT_NONE"
	PARTIAL_CvssImpact CvssImpact = "IMPACT_PARTIAL"
	COMPLETE_CvssImpact CvssImpact = "IMPACT_COMPLETE"
)
