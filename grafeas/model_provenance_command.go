/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Command describes a step performed as part of the build pipeline.
type ProvenanceCommand struct {
	// Required. Name of the command, as presented on the command line, or if the command is packaged as a Docker container, as presented to `docker pull`.
	Name string `json:"name,omitempty"`
	// Environment variables set before running this command.
	Env []string `json:"env,omitempty"`
	// Command-line arguments used when executing this command.
	Args []string `json:"args,omitempty"`
	// Working directory (relative to project source root) used when running this command.
	Dir string `json:"dir,omitempty"`
	// Optional unique identifier for this command, used in wait_for to reference this command as a dependency.
	Id string `json:"id,omitempty"`
	// The ID(s) of the command(s) that this command depends on.
	WaitFor []string `json:"waitFor,omitempty"`
}
