/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// SpdxRelationshipType : - RELATIONSHIP_TYPE_UNSPECIFIED: Unspecified  - DESCRIBES: Is to be used when SPDXRef-DOCUMENT describes SPDXRef-A  - DESCRIBED_BY: Is to be used when SPDXRef-A is described by SPDXREF-Document  - CONTAINS: Is to be used when SPDXRef-A contains SPDXRef-B  - CONTAINED_BY: Is to be used when SPDXRef-A is contained by SPDXRef-B  - DEPENDS_ON: Is to be used when SPDXRef-A depends on SPDXRef-B  - DEPENDENCY_OF: Is to be used when SPDXRef-A is dependency of SPDXRef-B  - DEPENDENCY_MANIFEST_OF: Is to be used when SPDXRef-A is a manifest file that lists a set of dependencies for SPDXRef-B  - BUILD_DEPENDENCY_OF: Is to be used when SPDXRef-A is a build dependency of SPDXRef-B  - DEV_DEPENDENCY_OF: Is to be used when SPDXRef-A is a development dependency of SPDXRef-B  - OPTIONAL_DEPENDENCY_OF: Is to be used when SPDXRef-A is an optional dependency of SPDXRef-B  - PROVIDED_DEPENDENCY_OF: Is to be used when SPDXRef-A is a to be provided dependency of SPDXRef-B  - TEST_DEPENDENCY_OF: Is to be used when SPDXRef-A is a test dependency of SPDXRef-B  - RUNTIME_DEPENDENCY_OF: Is to be used when SPDXRef-A is a dependency required for the execution of SPDXRef-B  - EXAMPLE_OF: Is to be used when SPDXRef-A is an example of SPDXRef-B  - GENERATES: Is to be used when SPDXRef-A generates SPDXRef-B  - GENERATED_FROM: Is to be used when SPDXRef-A was generated from SPDXRef-B  - ANCESTOR_OF: Is to be used when SPDXRef-A is an ancestor (same lineage but pre-dates) SPDXRef-B  - DESCENDANT_OF: Is to be used when SPDXRef-A is a descendant of (same lineage but postdates) SPDXRef-B  - VARIANT_OF: Is to be used when SPDXRef-A is a variant of (same lineage but not clear which came first) SPDXRef-B  - DISTRIBUTION_ARTIFACT: Is to be used when distributing SPDXRef-A requires that SPDXRef-B also be distributed  - PATCH_FOR: Is to be used when SPDXRef-A is a patch file for (to be applied to) SPDXRef-B  - PATCH_APPLIED: Is to be used when SPDXRef-A is a patch file that has been applied to SPDXRef-B  - COPY_OF: Is to be used when SPDXRef-A is an exact copy of SPDXRef-B  - FILE_ADDED: Is to be used when SPDXRef-A is a file that was added to SPDXRef-B  - FILE_DELETED: Is to be used when SPDXRef-A is a file that was deleted from SPDXRef-B  - FILE_MODIFIED: Is to be used when SPDXRef-A is a file that was modified from SPDXRef-B  - EXPANDED_FROM_ARCHIVE: Is to be used when SPDXRef-A is expanded from the archive SPDXRef-B  - DYNAMIC_LINK: Is to be used when SPDXRef-A dynamically links to SPDXRef-B  - STATIC_LINK: Is to be used when SPDXRef-A statically links to SPDXRef-B  - DATA_FILE_OF: Is to be used when SPDXRef-A is a data file used in SPDXRef-B  - TEST_CASE_OF: Is to be used when SPDXRef-A is a test case used in testing SPDXRef-B  - BUILD_TOOL_OF: Is to be used when SPDXRef-A is used to build SPDXRef-B  - DEV_TOOL_OF: Is to be used when SPDXRef-A is used as a development tool for SPDXRef-B  - TEST_OF: Is to be used when SPDXRef-A is used for testing SPDXRef-B  - TEST_TOOL_OF: Is to be used when SPDXRef-A is used as a test tool for SPDXRef-B  - DOCUMENTATION_OF: Is to be used when SPDXRef-A provides documentation of SPDXRef-B  - OPTIONAL_COMPONENT_OF: Is to be used when SPDXRef-A is an optional component of SPDXRef-B  - METAFILE_OF: Is to be used when SPDXRef-A is a metafile of SPDXRef-B  - PACKAGE_OF: Is to be used when SPDXRef-A is used as a package as part of SPDXRef-B  - AMENDS: Is to be used when (current) SPDXRef-DOCUMENT amends the SPDX information in SPDXRef-B  - PREREQUISITE_FOR: Is to be used when SPDXRef-A is a prerequisite for SPDXRef-B  - HAS_PREREQUISITE: Is to be used when SPDXRef-A has as a prerequisite SPDXRef-B  - OTHER: Is to be used for a relationship which has not been defined in the formal SPDX specification. A description of the relationship should be included in the Relationship comments field
type SpdxRelationshipType string

// List of spdxRelationshipType
const (
	RELATIONSHIP_TYPE_UNSPECIFIED_SpdxRelationshipType SpdxRelationshipType = "RELATIONSHIP_TYPE_UNSPECIFIED"
	DESCRIBES_SpdxRelationshipType SpdxRelationshipType = "DESCRIBES"
	DESCRIBED_BY_SpdxRelationshipType SpdxRelationshipType = "DESCRIBED_BY"
	CONTAINS_SpdxRelationshipType SpdxRelationshipType = "CONTAINS"
	CONTAINED_BY_SpdxRelationshipType SpdxRelationshipType = "CONTAINED_BY"
	DEPENDS_ON_SpdxRelationshipType SpdxRelationshipType = "DEPENDS_ON"
	DEPENDENCY_OF_SpdxRelationshipType SpdxRelationshipType = "DEPENDENCY_OF"
	DEPENDENCY_MANIFEST_OF_SpdxRelationshipType SpdxRelationshipType = "DEPENDENCY_MANIFEST_OF"
	BUILD_DEPENDENCY_OF_SpdxRelationshipType SpdxRelationshipType = "BUILD_DEPENDENCY_OF"
	DEV_DEPENDENCY_OF_SpdxRelationshipType SpdxRelationshipType = "DEV_DEPENDENCY_OF"
	OPTIONAL_DEPENDENCY_OF_SpdxRelationshipType SpdxRelationshipType = "OPTIONAL_DEPENDENCY_OF"
	PROVIDED_DEPENDENCY_OF_SpdxRelationshipType SpdxRelationshipType = "PROVIDED_DEPENDENCY_OF"
	TEST_DEPENDENCY_OF_SpdxRelationshipType SpdxRelationshipType = "TEST_DEPENDENCY_OF"
	RUNTIME_DEPENDENCY_OF_SpdxRelationshipType SpdxRelationshipType = "RUNTIME_DEPENDENCY_OF"
	EXAMPLE_OF_SpdxRelationshipType SpdxRelationshipType = "EXAMPLE_OF"
	GENERATES_SpdxRelationshipType SpdxRelationshipType = "GENERATES"
	GENERATED_FROM_SpdxRelationshipType SpdxRelationshipType = "GENERATED_FROM"
	ANCESTOR_OF_SpdxRelationshipType SpdxRelationshipType = "ANCESTOR_OF"
	DESCENDANT_OF_SpdxRelationshipType SpdxRelationshipType = "DESCENDANT_OF"
	VARIANT_OF_SpdxRelationshipType SpdxRelationshipType = "VARIANT_OF"
	DISTRIBUTION_ARTIFACT_SpdxRelationshipType SpdxRelationshipType = "DISTRIBUTION_ARTIFACT"
	PATCH_FOR_SpdxRelationshipType SpdxRelationshipType = "PATCH_FOR"
	PATCH_APPLIED_SpdxRelationshipType SpdxRelationshipType = "PATCH_APPLIED"
	COPY_OF_SpdxRelationshipType SpdxRelationshipType = "COPY_OF"
	FILE_ADDED_SpdxRelationshipType SpdxRelationshipType = "FILE_ADDED"
	FILE_DELETED_SpdxRelationshipType SpdxRelationshipType = "FILE_DELETED"
	FILE_MODIFIED_SpdxRelationshipType SpdxRelationshipType = "FILE_MODIFIED"
	EXPANDED_FROM_ARCHIVE_SpdxRelationshipType SpdxRelationshipType = "EXPANDED_FROM_ARCHIVE"
	DYNAMIC_LINK_SpdxRelationshipType SpdxRelationshipType = "DYNAMIC_LINK"
	STATIC_LINK_SpdxRelationshipType SpdxRelationshipType = "STATIC_LINK"
	DATA_FILE_OF_SpdxRelationshipType SpdxRelationshipType = "DATA_FILE_OF"
	TEST_CASE_OF_SpdxRelationshipType SpdxRelationshipType = "TEST_CASE_OF"
	BUILD_TOOL_OF_SpdxRelationshipType SpdxRelationshipType = "BUILD_TOOL_OF"
	DEV_TOOL_OF_SpdxRelationshipType SpdxRelationshipType = "DEV_TOOL_OF"
	TEST_OF_SpdxRelationshipType SpdxRelationshipType = "TEST_OF"
	TEST_TOOL_OF_SpdxRelationshipType SpdxRelationshipType = "TEST_TOOL_OF"
	DOCUMENTATION_OF_SpdxRelationshipType SpdxRelationshipType = "DOCUMENTATION_OF"
	OPTIONAL_COMPONENT_OF_SpdxRelationshipType SpdxRelationshipType = "OPTIONAL_COMPONENT_OF"
	METAFILE_OF_SpdxRelationshipType SpdxRelationshipType = "METAFILE_OF"
	PACKAGE_OF_SpdxRelationshipType SpdxRelationshipType = "PACKAGE_OF"
	AMENDS_SpdxRelationshipType SpdxRelationshipType = "AMENDS"
	PREREQUISITE_FOR_SpdxRelationshipType SpdxRelationshipType = "PREREQUISITE_FOR"
	HAS_PREREQUISITE_SpdxRelationshipType SpdxRelationshipType = "HAS_PREREQUISITE"
	OTHER_SpdxRelationshipType SpdxRelationshipType = "OTHER"
)