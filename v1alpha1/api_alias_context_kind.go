/*
 * v1alpha1/proto/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// ApiAliasContextKind : The type of an alias.   - KIND_UNSPECIFIED: Unknown.  - FIXED: Git tag.  - MOVABLE: Git branch.  - OTHER: Used to specify non-standard aliases. For example, if a Git repo has a ref named \"refs/foo/bar\".
type apiAliasContextKind string

// List of apiAliasContextKind
const (
	KIND_UNSPECIFIED ApiAliasContextKind = "KIND_UNSPECIFIED"
	FIXED ApiAliasContextKind = "FIXED"
	MOVABLE ApiAliasContextKind = "MOVABLE"
	OTHER ApiAliasContextKind = "OTHER"
)