/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

type CvssPrivilegesRequired string

// List of CVSSPrivilegesRequired
const (
	UNSPECIFIED_CvssPrivilegesRequired CvssPrivilegesRequired = "PRIVILEGES_REQUIRED_UNSPECIFIED"
	NONE_CvssPrivilegesRequired CvssPrivilegesRequired = "PRIVILEGES_REQUIRED_NONE"
	LOW_CvssPrivilegesRequired CvssPrivilegesRequired = "PRIVILEGES_REQUIRED_LOW"
	HIGH_CvssPrivilegesRequired CvssPrivilegesRequired = "PRIVILEGES_REQUIRED_HIGH"
)
