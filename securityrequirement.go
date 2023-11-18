package goopenapispec

type SecurityRequirement struct {
	// Name the name of the security scheme
	Name string `json:"name" yaml:"name"`
	// Scopes the list of scope names required for the execution
	Scopes []string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}
