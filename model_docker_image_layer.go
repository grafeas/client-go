/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

// Layer holds metadata specific to a layer of a Docker image.
type DockerImageLayer struct {
	// The recovered Dockerfile directive used to construct this layer.
	Directive *LayerDirective `json:"directive,omitempty"`

	// The recovered arguments to the Dockerfile directive.
	Arguments string `json:"arguments,omitempty"`
}