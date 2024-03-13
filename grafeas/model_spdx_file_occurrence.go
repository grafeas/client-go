/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

type SpdxFileOccurrence struct {
	Id string `json:"id,omitempty"`
	LicenseConcluded *V1beta1License `json:"licenseConcluded,omitempty"`
	FilesLicenseInfo []string `json:"filesLicenseInfo,omitempty"`
	Copyright string `json:"copyright,omitempty"`
	Comment string `json:"comment,omitempty"`
	Notice string `json:"notice,omitempty"`
	Contributors []string `json:"contributors,omitempty"`
	Attributions []string `json:"attributions,omitempty"`
}