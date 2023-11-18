package goopenapispec

type Server struct {
	// Url represents the URL of the server and is required.
	Url string `json:"url" yaml:"url"`

	// Description provides an optional description of the server.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Variables contain a map of variable names to their values.
	Variables map[string]ServerVariable `json:"variables,omitempty" yaml:"variables,omitempty"`
}
